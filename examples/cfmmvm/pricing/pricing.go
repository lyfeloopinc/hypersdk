// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package pricing

// IDs for pricing models
const (
	InvalidModelID uint8 = iota
	ConstantProductID
)

type Model interface {
	AddLiquidity(amountX uint64, amountY uint64, lpTokenSupply uint64) (uint64, uint64, uint64, error)
	RemoveLiquidity(amount uint64) (uint64, uint64, error)
	Swap(amountX uint64, amountY uint64) (uint64, uint64, error)
	GetState() (uint64, uint64, uint64)
}

type NewModel func(uint64, uint64, uint64, uint64) Model

var Models map[uint8]NewModel

func init() {
	Models = make(map[uint8]NewModel)

	// Append any additional pricing models here
	Models[ConstantProductID] = NewConstantProduct
}
