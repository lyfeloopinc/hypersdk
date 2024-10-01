// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package layout

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLayoutPrefixConflicts(t *testing.T) {
	require := require.New(t)

	// Test that default starting prefix works
	require.NoError(IsValidLayout([]byte{lowestAvailablePrefix}))

	// Test that a conflicting prefix raises an error
	require.Error(IsValidLayout([]byte{defaultHeightStatePrefix}))
}

func TestIsConflictingPrefix(t *testing.T) {
	require := require.New(t)

	// Test that default starting prefix works
	require.False(IsConflictingPrefix(lowestAvailablePrefix))

	// Test that a conflicting prefix raises an error
	require.True(IsConflictingPrefix(defaultHeightStatePrefix))
}
