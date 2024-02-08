package chain

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	smblock "github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/hypersdk/executor"
	"github.com/ava-labs/hypersdk/keys"
	"github.com/ava-labs/hypersdk/state"
	"github.com/ava-labs/hypersdk/tstate"
	"github.com/sourcegraph/conc/stream"
	"go.uber.org/zap"
)

const numTxs = 50000 // TODO: somehow estimate this (needed to ensure no backlog)

type Processor struct {
	vm  VM
	eng *Engine

	authStream *stream.Stream

	blkCtx     *smblock.Context
	timestamp  int64
	im         state.Immutable
	feeManager *FeeManager
	r          Rules
	sm         StateManager
	cacheLock  sync.RWMutex
	cache      map[string]*fetchData
	exectutor  *executor.Executor
	ts         *tstate.TState

	txs     set.Set[ids.ID]
	results [][]*Result
}

type fetchData struct {
	v      []byte
	exists bool

	chunks uint16
}

// Only run one processor at once
func NewProcessor(
	vm VM, eng *Engine,
	blkCtx *smblock.Context, chunks int, timestamp int64, im state.Immutable, feeManager *FeeManager, r Rules,
) *Processor {
	stream := stream.New()
	stream.WithMaxGoroutines(10) // TODO: use config
	return &Processor{
		vm:  vm,
		eng: eng,

		authStream: stream,

		blkCtx:     blkCtx,
		timestamp:  timestamp,
		im:         im,
		feeManager: feeManager,
		r:          r,
		sm:         vm.StateManager(),
		cache:      make(map[string]*fetchData, numTxs),
		// Executor is shared across all chunks, this means we don't need to "wait" at the end of each chunk to continue
		// processing transactions.
		exectutor: executor.New(numTxs, vm.GetTransactionExecutionCores(), vm.GetExecutorVerifyRecorder()),
		ts:        tstate.New(numTxs * 2),

		txs:     set.NewSet[ids.ID](numTxs),
		results: make([][]*Result, chunks),
	}
}

func (p *Processor) process(ctx context.Context, chunkIndex int, txIndex int, tx *Transaction) {
	stateKeys, err := tx.StateKeys(p.sm)
	if err != nil {
		p.vm.Logger().Warn("could not compute state keys", zap.Stringer("txID", tx.ID()), zap.Error(err))
		p.results[chunkIndex][txIndex] = &Result{Valid: false}
		return
	}
	p.exectutor.Run(stateKeys, func() error {
		// Fetch keys from cache
		var (
			reads    = make(map[string]uint16, len(stateKeys))
			storage  = make(map[string][]byte, len(stateKeys))
			toLookup = make([]string, 0, len(stateKeys))
		)
		p.cacheLock.RLock()
		for k := range stateKeys {
			if v, ok := p.cache[k]; ok {
				reads[k] = v.chunks
				if v.exists {
					storage[k] = v.v
				}
				continue
			}
			toLookup = append(toLookup, k)
		}
		p.cacheLock.RUnlock()

		// Fetch keys from disk
		var toCache map[string]*fetchData
		if len(toLookup) > 0 {
			toCache = make(map[string]*fetchData, len(toLookup))
			for _, k := range toLookup {
				v, err := p.im.GetValue(ctx, []byte(k))
				if errors.Is(err, database.ErrNotFound) {
					reads[k] = 0
					toCache[k] = &fetchData{nil, false, 0}
					continue
				} else if err != nil {
					return err
				}
				// We verify that the [NumChunks] is already less than the number
				// added on the write path, so we don't need to do so again here.
				numChunks, ok := keys.NumChunks(v)
				if !ok {
					return ErrInvalidKeyValue
				}
				reads[k] = numChunks
				toCache[k] = &fetchData{v, true, numChunks}
				storage[k] = v
			}
		}

		// Execute transaction
		//
		// It is critical we explicitly set the scope before each transaction is
		// processed
		tsv := p.ts.NewView(stateKeys, storage)

		// Ensure we have enough funds to pay fees
		if err := tx.PreExecute(ctx, p.feeManager, p.sm, p.r, tsv, p.timestamp); err != nil {
			// TODO: freeze account and pay bond
			p.vm.Logger().Warn("pre-execution failure", zap.Stringer("txID", tx.ID()), zap.Error(err))
			p.results[chunkIndex][txIndex] = &Result{Valid: false}
			return nil
		}

		// Wait to perform warp verification until we know the transaction can pay fees
		warpVerified := p.verifyWarpMessage(ctx, tx)

		// Execute transaction
		result, err := tx.Execute(ctx, p.feeManager, reads, p.sm, p.r, tsv, p.timestamp, warpVerified)
		if err != nil {
			p.vm.Logger().Warn("execution failure", zap.Stringer("txID", tx.ID()), zap.Error(err))
			p.results[chunkIndex][txIndex] = &Result{Valid: false}
			return nil
		}
		result.Valid = true
		p.txs.Add(tx.ID()) // only track duplicates on valid txs
		result.WarpVerified = warpVerified
		p.results[chunkIndex][txIndex] = result

		// Update block metadata with units actually consumed (if more is consumed than block allows, we will non-deterministically
		// exit with an error based on which tx over the limit is processed first)
		//
		// TODO: we won't know this when just including certs?
		if ok, d := p.feeManager.Consume(result.Consumed, p.r.GetMaxBlockUnits()); !ok {
			return fmt.Errorf("%w: %d too large", ErrInvalidUnitsConsumed, d)
		}

		// Commit results to parent [TState]
		tsv.Commit()

		// TODO: pay portion of fees to validator that included

		// Update key cache
		if len(toCache) > 0 {
			p.cacheLock.Lock()
			for k := range toCache {
				p.cache[k] = toCache[k]
			}
			p.cacheLock.Unlock()
		}
		return nil
	})
}

func (p *Processor) verifyWarpMessage(ctx context.Context, tx *Transaction) bool {
	if tx.WarpMessage == nil {
		return false
	}

	allowed, num, denom := p.r.GetWarpConfig(tx.WarpMessage.SourceChainID)
	if !allowed {
		p.vm.Logger().
			Warn("unable to verify warp message", zap.Stringer("warpID", tx.WarpMessage.ID()), zap.Error(ErrDisabledChainID))

	}
	if err := tx.WarpMessage.Signature.Verify(
		ctx,
		&tx.WarpMessage.UnsignedMessage,
		p.r.NetworkID(),
		p.vm.ValidatorState(),
		p.blkCtx.PChainHeight,
		num,
		denom,
	); err != nil {
		p.vm.Logger().
			Warn("unable to verify warp message", zap.Stringer("warpID", tx.WarpMessage.ID()), zap.Error(err))
		return false
	}
	return true
}

// Allows processing to start before all chunks are acquired.
//
// Chunks MUST be added in order.
//
// Add must not be called concurrently
func (p *Processor) Add(ctx context.Context, chunkIndex int, chunk *Chunk) error {
	ctx, span := p.vm.Tracer().Start(ctx, "Processor.Add")
	defer span.End()

	// Kickoff async signature verification (auth + warp)
	//
	// Wait to start any disk lookup until signature verification is done for that transaction.
	//
	// We can't use batch verification because we don't know which transactions
	// may fail auth.
	//
	// Don't wait for all transactions to finish verification to kickoff execution (should
	// be interleaved).
	p.results[chunkIndex] = make([]*Result, len(chunk.Txs))
	repeats, err := p.eng.IsRepeatTx(ctx, chunk.Txs, set.NewBits())
	if err != nil {
		return err
	}
	for ri, rtx := range chunk.Txs {
		txIndex := ri
		tx := rtx

		// Perform basic verification (also performed inside of PreExecute)
		if err := tx.Base.Execute(p.r.ChainID(), p.r, p.timestamp); err != nil {
			p.vm.Logger().Warn("base transaction is invalid", zap.Stringer("txID", tx.ID()), zap.Error(err))
			p.results[chunkIndex][txIndex] = &Result{Valid: false}
			continue
		}
		if tx.Base.Timestamp > chunk.Slot {
			p.vm.Logger().Warn("base transaction has timestamp after slot", zap.Stringer("txID", tx.ID()))
			p.results[chunkIndex][txIndex] = &Result{Valid: false}
			continue
		}

		// Check that transaction isn't a duplicate
		//
		// TODO: this contains is concurrent map read/write? ...shouldn't be checked here, this isn't safe
		if repeats.Contains(txIndex) || p.txs.Contains(tx.ID()) {
			p.vm.Logger().Warn("transaction is a duplicate", zap.Stringer("txID", tx.ID()))
			p.results[chunkIndex][txIndex] = &Result{Valid: false}
			continue
		}

		// TODO: Check that transaction included in right partition

		// Enqueue transaction for execution
		p.authStream.Go(func() stream.Callback {
			msg, err := tx.Digest()
			if err != nil {
				p.vm.Logger().Warn("could not compute tx digest", zap.Stringer("txID", tx.ID()), zap.Error(err))
				p.results[chunkIndex][txIndex] = &Result{Valid: false}
				return func() {}
			}
			if p.vm.GetVerifyAuth() {
				if err := tx.Auth.Verify(ctx, msg); err != nil {
					p.vm.Logger().Warn("auth verification failed", zap.Stringer("txID", tx.ID()), zap.Error(err))
					p.results[chunkIndex][txIndex] = &Result{Valid: false}
					return func() {}
				}
			}
			return func() { p.process(ctx, chunkIndex, txIndex, tx) }
		})
	}
	return nil
}

func (p *Processor) Wait() (set.Set[ids.ID], *tstate.TState, [][]*Result, error) {
	p.authStream.Wait()
	return p.txs, p.ts, p.results, p.exectutor.Wait()
}
