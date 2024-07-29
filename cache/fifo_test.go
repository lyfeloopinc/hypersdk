// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cache

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFIFOCacheInsertion(t *testing.T) {
	type put struct {
		i      int
		exists bool
	}

	type get struct {
		i  int
		ok bool
	}

	tests := []struct {
		name string
		ops  []interface{}
	}{
		{
			name: "inserting up to limit works",
			ops: []interface{}{
				put{
					i:      0,
					exists: false,
				},
				put{
					i:      1,
					exists: false,
				},
			},
		},
		{
			name: "inserting after limit cleans first",
			ops: []interface{}{
				put{
					i:      0,
					exists: false,
				},
				put{
					i:      1,
					exists: false,
				},
				put{
					i:      2,
					exists: false,
				},
				get{
					i:  0,
					ok: false,
				},
				get{
					i:  1,
					ok: true,
				},
			},
		},
		{
			name: "no elements removed when cache is exactly full",
			ops: []interface{}{
				put{
					i:      0,
					exists: false,
				},
				put{
					i:      1,
					exists: false,
				},
				get{
					i:  0,
					ok: true,
				},
				get{
					i:  1,
					ok: true,
				},
			},
		},
		{
			name: "no elements removed when the cache is less than full",
			ops: []interface{}{
				put{
					i:      0,
					exists: false,
				},
				get{
					i:  0,
					ok: true,
				},
			},
		},
		{
			name: "inserting existing value when full doesn't free value",
			ops: []interface{}{
				put{
					i:      0,
					exists: false,
				},
				put{
					i:      0,
					exists: true,
				},
				get{
					i:  0,
					ok: true,
				},
			},
		},
		{
			name: "elements removed in FIFO order when cache overfills",
			ops: []interface{}{
				put{
					i:      0,
					exists: false,
				},
				put{
					i:      1,
					exists: false,
				},
				put{
					i:      2,
					exists: false,
				},
				get{
					i:  0,
					ok: false,
				},
				put{
					i:      3,
					exists: false,
				},
				get{
					i:  1,
					ok: false,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			cache, err := NewFIFO[int, int](2)
			require.NoError(err)

			for _, op := range tt.ops {
				if put, ok := op.(put); ok {
					exists := cache.Put(put.i, put.i)
					require.Equal(put.exists, exists)
				} else if get, ok := op.(get); ok {
					val, ok := cache.Get(get.i)
					require.Equal(get.ok, ok)
					if ok {
						require.Equal(get.i, val)
					}
				} else {
					require.Fail("op can only be a put or a get")
				}
			}
		})
	}
}

func TestEmptyCacheFails(t *testing.T) {
	require := require.New(t)
	_, err := NewFIFO[int, int](0)
	expectedErr := errors.New("maxSize must be greater than 0")
	require.Equal(expectedErr, err)
}
