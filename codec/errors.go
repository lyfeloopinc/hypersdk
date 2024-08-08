// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package codec

import "errors"

var (
	ErrTooManyItems         = errors.New("too many items")
	ErrDuplicateItem        = errors.New("duplicate item")
	ErrFieldNotPopulated    = errors.New("field is not populated")
	ErrInvalidBitset        = errors.New("invalid bitset")
	ErrIncorrectHRP         = errors.New("incorrect hrp")
	ErrInsufficientLength   = errors.New("insufficient length")
	ErrInvalidSize          = errors.New("invalid size")
	ErrStringTooLong        = errors.New("string length exceeds maximum allowed")
	ErrUnsupportedFieldType = errors.New("unsupported field type")
	ErrEmptyAddress         = errors.New("empty address is not allowed during marshal")
)
