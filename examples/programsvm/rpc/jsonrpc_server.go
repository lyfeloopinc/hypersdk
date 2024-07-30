// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"context"
	"github.com/ava-labs/hypersdk/examples/programsvm/actions"
	"github.com/ava-labs/hypersdk/state"
	"net/http"

	"github.com/ava-labs/avalanchego/ids"

	"github.com/ava-labs/hypersdk/codec"
	"github.com/ava-labs/hypersdk/examples/programsvm/consts"
	"github.com/ava-labs/hypersdk/examples/programsvm/genesis"
	"github.com/ava-labs/hypersdk/fees"
)

type JSONRPCServer struct {
	c        Controller
	simulate func(context.Context, actions.CallProgram, codec.Address) (state.Keys, error)
}

func NewJSONRPCServer(c Controller, simulate func(context.Context, actions.CallProgram, codec.Address) (state.Keys, error)) *JSONRPCServer {
	return &JSONRPCServer{c, simulate}
}

type GenesisReply struct {
	Genesis *genesis.Genesis `json:"genesis"`
}

func (j *JSONRPCServer) Genesis(_ *http.Request, _ *struct{}, reply *GenesisReply) (err error) {
	reply.Genesis = j.c.Genesis()
	return nil
}

type SimulateCallProgramTxArgs struct {
	CallProgramTx actions.CallProgram `json:"callProgramTx"`
	Actor         codec.Address       `json:"actor"`
}

type SimulateCallProgramTxReply struct {
	StateKeys state.Keys `json:"stateKeys"`
}

func (j *JSONRPCServer) SimulateCallProgramTx(req *http.Request, args *SimulateCallProgramTxArgs, reply *SimulateCallProgramTxReply) (err error) {
	reply.StateKeys, err = j.simulate(req.Context(), args.CallProgramTx, args.Actor)
	return err
}

type TxArgs struct {
	TxID ids.ID `json:"txId"`
}

type TxReply struct {
	Timestamp int64           `json:"timestamp"`
	Success   bool            `json:"success"`
	Units     fees.Dimensions `json:"units"`
	Fee       uint64          `json:"fee"`
}

func (j *JSONRPCServer) Tx(req *http.Request, args *TxArgs, reply *TxReply) error {
	_, span := j.c.Tracer().Start(req.Context(), "Server.Tx")
	defer span.End()

	found, t, success, units, fee, err := j.c.GetTransaction(args.TxID)
	if err != nil {
		return err
	}
	if !found {
		return ErrTxNotFound
	}
	reply.Timestamp = t
	reply.Success = success
	reply.Units = units
	reply.Fee = fee
	return nil
}

type BalanceArgs struct {
	Address string `json:"address"`
}

type BalanceReply struct {
	Amount uint64 `json:"amount"`
}

func (j *JSONRPCServer) Balance(req *http.Request, args *BalanceArgs, reply *BalanceReply) error {
	ctx, span := j.c.Tracer().Start(req.Context(), "Server.Balance")
	defer span.End()

	addr, err := codec.ParseAddressBech32(consts.HRP, args.Address)
	if err != nil {
		return err
	}
	balance, err := j.c.GetBalanceFromState(ctx, addr)
	if err != nil {
		return err
	}
	reply.Amount = balance
	return err
}
