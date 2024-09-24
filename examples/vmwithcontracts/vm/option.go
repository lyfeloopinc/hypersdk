// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vm

import (
	"github.com/ava-labs/hypersdk/state/tstate"
	"github.com/ava-labs/hypersdk/vm"
	"github.com/ava-labs/hypersdk/x/contracts/runtime"
)

const Namespace = "controller"

type Config struct {
	Enabled bool `json:"enabled"`
}

func NewDefaultConfig() Config {
	return Config{
		Enabled: true,
	}
}

func With() vm.Option[*tstate.TStateView] {
	return vm.NewOption[*tstate.TStateView](Namespace, NewDefaultConfig(), func(v *vm.VM[*tstate.TStateView], config Config) error {
		if !config.Enabled {
			return nil
		}
		vm.WithVMAPIs[*tstate.TStateView](jsonRPCServerFactory{})(v)
		return nil
	})
}

func WithRuntime() vm.Option[*tstate.TStateView] {
	return vm.NewOption[*tstate.TStateView](Namespace+"runtime", *runtime.NewConfig(), func(v *vm.VM[*tstate.TStateView], cfg runtime.Config) error {
		wasmRuntime = runtime.NewRuntime(&cfg, v.Logger())
		return nil
	})
}
