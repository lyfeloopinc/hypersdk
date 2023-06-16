// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package chain

import (
	"context"
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	smblock "github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/hypersdk/consts"
	"go.uber.org/zap"
)

const txBatchSize = 4_096

func HandlePreExecute(
	err error,
) (bool /* continue */, bool /* restore */, bool /* remove account */) {
	switch {
	case errors.Is(err, ErrInsufficientPrice):
		return true, true, false
	case errors.Is(err, ErrTimestampTooEarly):
		return true, true, false
	case errors.Is(err, ErrTimestampTooLate):
		return true, false, false
	case errors.Is(err, ErrInvalidBalance):
		return true, false, true
	case errors.Is(err, ErrAuthNotActivated):
		return true, false, false
	case errors.Is(err, ErrAuthFailed):
		return true, false, false
	case errors.Is(err, ErrActionNotActivated):
		return true, false, false
	default:
		// If unknown error, drop
		return true, false, false
	}
}

// TODO: add a min build time where we just listen for txs
func BuildBlock(
	ctx context.Context,
	vm VM,
	preferred ids.ID,
	blockContext *smblock.Context,
) (*StatelessRootBlock, error) {
	ctx, span := vm.Tracer().Start(ctx, "chain.BuildBlock")
	defer span.End()
	log := vm.Logger()

	// TODO: migrate to milli
	// nextTime := time.Now().UnixMilli()
	nextTime := time.Now().Unix()
	r := vm.Rules(nextTime)
	parent, err := vm.GetStatelessRootBlock(ctx, preferred)
	if err != nil {
		log.Warn("block building failed: couldn't get parent", zap.Error(err))
		return nil, err
	}
	parentTxBlock, err := parent.LastTxBlock()
	if err != nil {
		log.Warn("block building failed: couldn't get parent tx block", zap.Error(err))
		return nil, err
	}

	b := NewRootBlock(vm, parent, nextTime)
	var (
		oldestAllowed = nextTime - r.GetValidityWindow()
		mempool       = vm.Mempool()

		txBlocks    = []*StatelessTxBlock{}
		txBlock     = NewTxBlock(vm, parentTxBlock, nextTime)
		txBlockSize = 0

		txsAttempted = 0
		txsAdded     = 0
		warpCount    = 0

		start = time.Now()
	)
	b.MinTxHght = txBlock.Hght

	mempool.StartBuild(ctx)
	for txBlock != nil && (time.Since(start) < vm.GetMinBuildTime() || mempool.Len(ctx) > 0) && time.Since(start) < vm.GetMaxBuildTime() {
		txs := mempool.LeaseItems(ctx, txBatchSize)
		if len(txs) == 0 {
			mempool.ClearLease(ctx, nil, nil)
			time.Sleep(25 * time.Millisecond)
			continue
		}

		restorable := make([]*Transaction, 0, txBatchSize)
		exempt := make([]*Transaction, 0, 10)
		var execErr error
		for i, next := range txs {
			txsAttempted++

			// Ensure we can process if transaction includes a warp message
			if next.WarpMessage != nil && blockContext == nil {
				log.Info(
					"dropping pending warp message because no context provided",
					zap.Stringer("txID", next.ID()),
				)
				if next.Base.Timestamp > oldestAllowed {
					restorable = append(restorable, next)
					continue
				}
			}

			// Skip warp message if at max
			if next.WarpMessage != nil && warpCount == MaxWarpMessages {
				log.Info(
					"dropping pending warp message because already have MaxWarpMessages",
					zap.Stringer("txID", next.ID()),
				)
				exempt = append(exempt, next)
				continue
			}

			// Check for repeats
			//
			// TODO: check a bunch at once during pre-fetch to avoid re-walking blocks
			// for every tx
			dup, err := parentTxBlock.IsRepeat(ctx, oldestAllowed, []*Transaction{next})
			if err != nil {
				restorable = append(restorable, txs[i:]...)
				execErr = err
				break
			}
			if dup {
				continue
			}

			// TODO: verify units space
			nextSize := next.Size()

			// Determine if we need to create a new TxBlock
			//
			// TODO: handle case where tx is larger than max size of TxBlock
			if txBlockSize+nextSize > consts.NetworkSizeLimit-4_096 {
				txBlock.Issued = time.Now().UnixMilli()
				if err := txBlock.initializeBuilt(ctx); err != nil {
					restorable = append(restorable, txs[i:]...)
					execErr = err
					break
				}
				b.TxBlocks = append(b.TxBlocks, txBlock.ID())
				txBlocks = append(txBlocks, txBlock)
				vm.IssueTxBlock(ctx, txBlock)

				if len(txBlocks)+1 /* account for current */ >= r.GetMaxTxBlocks() {
					txBlock = nil
					restorable = append(restorable, txs[i:]...)
					break
				}
				parentTxBlock = txBlock
				txBlockSize = 0
				txBlock = NewTxBlock(vm, parentTxBlock, nextTime)
			}

			// Update block with new transaction
			txBlock.Txs = append(txBlock.Txs, next)
			if next.WarpMessage != nil {
				warpCount++
			}
			txBlockSize++
			txsAdded++
		}
		mempool.ClearLease(ctx, restorable, exempt)
		if execErr != nil {
			for _, block := range txBlocks {
				b.vm.Mempool().Add(ctx, block.Txs)
			}
			if txBlock != nil {
				b.vm.Mempool().Add(ctx, txBlock.Txs)
			}
			mempool.FinishBuild(ctx)
			b.vm.Logger().Warn("build failed", zap.Error(execErr))
			return nil, execErr
		}
	}
	mempool.FinishBuild(ctx)

	// Record if went to the limit
	if time.Since(start) >= vm.GetMaxBuildTime() {
		vm.RecordEarlyBuildStop()
	}

	// Create last tx block
	//
	// TODO: unify this logic with inner block tracker
	if txBlock != nil && len(txBlock.Txs) > 0 {
		txBlock.Issued = time.Now().UnixMilli()
		if err := txBlock.initializeBuilt(ctx); err != nil {
			return nil, err
		}
		b.TxBlocks = append(b.TxBlocks, txBlock.ID())
		txBlocks = append(txBlocks, txBlock)
		vm.IssueTxBlock(ctx, txBlock)
	}

	// Perform basic validity checks to make sure the block is well-formatted
	if len(b.TxBlocks) == 0 {
		return nil, ErrNoTxs
	}
	b.ContainsWarp = warpCount > 0
	b.Issued = time.Now().UnixMilli()
	if err := b.initializeBuilt(ctx, txBlocks); err != nil {
		return nil, err
	}
	mempoolSize := b.vm.Mempool().Len(ctx)
	vm.RecordMempoolSizeAfterBuild(mempoolSize)
	log.Info(
		"built block",
		zap.Uint64("hght", b.Hght),
		zap.Int("attempted", txsAttempted),
		zap.Int("added", len(b.TxBlocks)),
		zap.Int("mempool size", mempoolSize),
		zap.Bool("context", blockContext != nil),
	)
	return b, nil
}
