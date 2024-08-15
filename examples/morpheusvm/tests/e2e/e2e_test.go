// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package e2e_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
	"github.com/stretchr/testify/require"

	"github.com/ava-labs/hypersdk/auth"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/codec"
	"github.com/ava-labs/hypersdk/crypto/ed25519"
	"github.com/ava-labs/hypersdk/examples/morpheusvm/actions"
	"github.com/ava-labs/hypersdk/examples/morpheusvm/consts"
	"github.com/ava-labs/hypersdk/rpc"
	"github.com/ava-labs/hypersdk/tests/fixture"
	"github.com/ava-labs/hypersdk/tests/workload"

	lrpc "github.com/ava-labs/hypersdk/examples/morpheusvm/rpc"
	le2e "github.com/ava-labs/hypersdk/examples/morpheusvm/tests/e2e"
	he2e "github.com/ava-labs/hypersdk/tests/e2e"
	ginkgo "github.com/onsi/ginkgo/v2"
)

const owner = "morpheusvm-e2e-tests"

var (
	_            workload.TxWorkloadFactory  = (*workloadFactory)(nil)
	_            workload.TxWorkloadIterator = (*simpleTxWorkload)(nil)
	flagVars     *e2e.FlagVars
	genesisBytes []byte
	factory      *auth.ED25519Factory
)

func TestE2e(t *testing.T) {
	ginkgo.RunSpecs(t, "morpheusvm e2e test suites")
}

func init() {
	flagVars = e2e.RegisterFlags()
	require := require.New(ginkgo.GinkgoT())

	genBytes, err := le2e.DefaultGenesisValues()
	require.NoError(err)
	genesisBytes = genBytes

	// Load default pk
	privBytes, err := codec.LoadHex(
		"323b1d8f4eed5f0da9da93071b034f2dce9d2d22692c172f3cb252a64ddfafd01b057de320297c29ad0c1f589ea216869cf1938d88c9fbd70d6748323dbf2fa7", //nolint:lll
		ed25519.PrivateKeyLen,
	)
	require.NoError(err)
	priv := ed25519.PrivateKey(privBytes)
	factory = auth.NewED25519Factory(priv)

	he2e.SetWorkload(consts.Name, &workloadFactory{factory})
}

var _ = ginkgo.SynchronizedBeforeSuite(func() []byte {
	// Run only once in the first ginkgo process
	nodes := tmpnet.NewNodesOrPanic(flagVars.NodeCount())
	subnet := fixture.NewHyperVMSubnet(
		consts.Name,
		consts.ID,
		genesisBytes,
		nodes...,
	)

	network := fixture.NewTmpnetNetwork(owner, nodes, subnet)
	return e2e.NewTestEnvironment(
		e2e.NewTestContext(),
		flagVars,
		network,
	).Marshal()
}, func(envBytes []byte) {
	// Run in every ginkgo process

	// Initialize the local test environment from the global state
	e2e.InitSharedTestEnvironment(ginkgo.GinkgoT(), envBytes)
})

type workloadFactory struct {
	factory *auth.ED25519Factory
}

func (f *workloadFactory) NewWorkloads(uri string) ([]workload.TxWorkloadIterator, error) {
	basicTxWorkload, err := f.NewSizedTxWorkload(uri, 1)
	return []workload.TxWorkloadIterator{basicTxWorkload}, err
}

func (f *workloadFactory) NewSizedTxWorkload(uri string, size int) (workload.TxWorkloadIterator, error) {
	cli := rpc.NewJSONRPCClient(uri)
	networkID, _, blockchainID, err := cli.Network(context.Background())
	if err != nil {
		return nil, err
	}
	lcli := lrpc.NewJSONRPCClient(uri, networkID, blockchainID)
	return &simpleTxWorkload{
		factory: f.factory,
		cli:     cli,
		lcli:    lcli,
		size:    size,
	}, nil
}

type simpleTxWorkload struct {
	factory   *auth.ED25519Factory
	cli       *rpc.JSONRPCClient
	lcli      *lrpc.JSONRPCClient
	networkID uint32
	chainID   ids.ID
	count     int
	size      int
}

func (g *simpleTxWorkload) Next() bool {
	return g.count < g.size
}

func (g *simpleTxWorkload) GenerateTxWithAssertion(ctx context.Context) (*chain.Transaction, func(ctx context.Context, uri string) error, error) {
	g.count++
	other, err := ed25519.GeneratePrivateKey()
	if err != nil {
		return nil, nil, err
	}

	aother := auth.NewED25519Address(other.PublicKey())
	aotherStr := codec.MustAddressBech32(consts.HRP, aother)
	parser, err := g.lcli.Parser(ctx)
	if err != nil {
		return nil, nil, err
	}
	_, tx, _, err := g.cli.GenerateTransaction(
		ctx,
		parser,
		[]chain.Action{&actions.Transfer{
			To:    aother,
			Value: 1,
		}},
		g.factory,
	)
	if err != nil {
		return nil, nil, err
	}

	return tx, func(ctx context.Context, uri string) error {
		lcli := lrpc.NewJSONRPCClient(uri, g.networkID, g.chainID)
		success, _, err := lcli.WaitForTransaction(ctx, tx.ID())
		if err != nil {
			return fmt.Errorf("failed to wait for tx %s: %w", tx.ID(), err)
		}
		if !success {
			return fmt.Errorf("tx %s not accepted", tx.ID())
		}
		balance, err := lcli.Balance(ctx, aotherStr)
		if err != nil {
			return fmt.Errorf("failed to get balance of %s: %w", aotherStr, err)
		}
		if balance != 1 {
			return fmt.Errorf("expected balance of 1, got %d", balance)
		}
		return nil
	}, nil
}
