// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package runtime

import (
	"context"
	"errors"
	"slices"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/bytecodealliance/wasmtime-go/v14"

	"github.com/ava-labs/hypersdk/codec"
)

type (
	ProgramCallErrorCode byte
	DeployErrorCode      byte
)

const (
	callProgramCost   = 10000
	setCallResultCost = 10000
	remainingFuelCost = 10000
	deployCost        = 10000
)

const (
	CallPanicked ProgramCallErrorCode = iota
	ExecutionFailure
	OutOfFuel
)

const (
	DeployFailed DeployErrorCode = iota
)

func extractProgramCallErrorCode(err error) (ProgramCallErrorCode, bool) {
	var trap *wasmtime.Trap
	if errors.As(err, &trap) {
		switch *trap.Code() {
		case wasmtime.UnreachableCodeReached:
			return CallPanicked, true
		case wasmtime.OutOfFuel:
			return OutOfFuel, true
		default:
			return ExecutionFailure, true
		}
	}
	return 0, false
}

type callProgramInput struct {
	Program      codec.Address
	FunctionName string
	Params       []byte
	Fuel         uint64
}

type deployProgramInput struct {
	ProgramID           ids.ID
	AccountCreationData []byte
}

func NewProgramModule(r *WasmRuntime) *ImportModule {
	return &ImportModule{
		Name: "program",
		HostFunctions: map[string]HostFunction{
			"call_program": {FuelCost: callProgramCost, Function: Function[callProgramInput, Result[RawBytes, ProgramCallErrorCode]](func(callInfo *CallInfo, input callProgramInput) (Result[RawBytes, ProgramCallErrorCode], error) {
				newInfo := *callInfo

				if err := callInfo.ConsumeFuel(input.Fuel); err != nil {
					return Err[RawBytes, ProgramCallErrorCode](OutOfFuel), nil
				}

				newInfo.Actor = callInfo.Program
				newInfo.Program = input.Program
				newInfo.FunctionName = input.FunctionName
				newInfo.Params = input.Params
				newInfo.Fuel = input.Fuel

				result, err := r.CallProgram(
					context.Background(),
					&newInfo)
				if err != nil {
					if code, ok := extractProgramCallErrorCode(err); ok {
						return Err[RawBytes, ProgramCallErrorCode](code), nil
					}
					return Err[RawBytes, ProgramCallErrorCode](ExecutionFailure), err
				}

				// return any remaining fuel to the calling program
				callInfo.AddFuel(newInfo.RemainingFuel())

				return Ok[RawBytes, ProgramCallErrorCode](result), nil
			})},
			"set_call_result": {FuelCost: setCallResultCost, Function: FunctionNoOutput[RawBytes](func(callInfo *CallInfo, input RawBytes) error {
				// needs to clone because this points into the current store's linear memory which may be gone when this is read
				callInfo.inst.result = slices.Clone(input)
				return nil
			})},
			"remaining_fuel": {FuelCost: remainingFuelCost, Function: FunctionNoInput[uint64](func(callInfo *CallInfo) (uint64, error) {
				return callInfo.RemainingFuel(), nil
			})},
			"deploy": {
				FuelCost: deployCost,
				Function: Function[deployProgramInput, Result[codec.Address, DeployErrorCode]](
					func(_ *CallInfo, input deployProgramInput) (Result[codec.Address, DeployErrorCode], error) {
						ctx, cancel := context.WithCancel(context.Background())
						defer cancel()
						address, err := r.programStore.NewAccountWithProgram(ctx, input.ProgramID, input.AccountCreationData)
						if err != nil {
							return Err[codec.Address, DeployErrorCode](DeployFailed), nil
						}
						return Ok[codec.Address, DeployErrorCode](address), nil
					}),
			},
		},
	}
}
