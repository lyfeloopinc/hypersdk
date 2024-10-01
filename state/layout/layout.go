// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package layout

import (
	"errors"

	"github.com/ava-labs/avalanchego/utils/set"
)

const (
	defaultHeightStatePrefix    byte = 0x0
	defaultTimestampStatePrefix byte = 0x1
	defaultFeeStatePrefix       byte = 0x2

	LowestAvailablePrefix byte = 0x3
)

var ErrConflictingPrefix = errors.New("conflicting prefix")

func IsValidLayout(vmSpecificPrefixes []byte) error {
	prefixes := []byte{
		defaultHeightStatePrefix,
		defaultTimestampStatePrefix,
		defaultFeeStatePrefix,
	}

	prefixes = append(prefixes, vmSpecificPrefixes...)

	verifiedPrefixes := set.Set[string]{}

	for _, k := range prefixes {
		keyString := string(k)

		for prefix := range verifiedPrefixes {
			if prefix == keyString {
				return ErrConflictingPrefix
			}
		}

		verifiedPrefixes.Add(keyString)
	}

	return nil
}

func HeightPrefix() []byte {
	return []byte{defaultHeightStatePrefix}
}

func TimestampPrefix() []byte {
	return []byte{defaultTimestampStatePrefix}
}

func FeePrefix() []byte {
	return []byte{defaultFeeStatePrefix}
}

func IsConflictingPrefix(prefix byte) bool {
	return defaultHeightStatePrefix == prefix || defaultTimestampStatePrefix == prefix || defaultFeeStatePrefix == prefix
}
