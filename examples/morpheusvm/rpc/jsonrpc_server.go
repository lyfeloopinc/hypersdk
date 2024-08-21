// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"context"
	"errors"
	"net/http"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/trace"

	"github.com/ava-labs/hypersdk/codec"
	"github.com/ava-labs/hypersdk/examples/morpheusvm/consts"
	"github.com/ava-labs/hypersdk/fees"
	"github.com/ava-labs/hypersdk/vm"
)

const JSONRPCEndpoint = "/morpheusapi"

var ErrTxNotFound = errors.New("tx not found")

type Controller interface {
	Tracer() trace.Tracer
	GetTransaction(ids.ID) (bool, int64, bool, fees.Dimensions, uint64, error)
	GetBalanceFromState(context.Context, codec.Address) (uint64, error)
}

type JSONRPCServer struct {
	c Controller
	g *vm.Bech32Genesis
	r vm.RuleFactory
}

func NewJSONRPCServer(c Controller, g *vm.Bech32Genesis, r vm.RuleFactory) *JSONRPCServer {
	return &JSONRPCServer{c: c, g: g, r: r}
}

type GenesisReply struct {
	Genesis *vm.Bech32Genesis `json:"genesis"`
}

func (j *JSONRPCServer) Genesis(_ *http.Request, _ *struct{}, reply *GenesisReply) (err error) {
	reply.Genesis = j.g
	return nil
}

type RulesReply struct {
	Rules *vm.Rules `json:"rules"`
}

func (j *JSONRPCServer) Rules(_ *http.Request, _ *struct{}, reply *RulesReply) (err error) {
	reply.Rules = j.r.GetRules(0).(*vm.Rules)
	return nil
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
