// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewIntegerNumber creates a new IntegerNumber IE.
func NewIntegerNumber(port uint16) *IE {
	return newUint16ValIE(IntegerNumber, port)
}

// IntegerNumber returns IntegerNumber in uint16 if the type of IE matches.
func (i *IE) IntegerNumber() (uint16, error) {
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case IntegerNumber:
		if len(i.Payload) < 2 {
			return 0, io.ErrUnexpectedEOF
		}
		return binary.BigEndian.Uint16(i.Payload[0:2]), nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustIntegerNumber returns IntegerNumber in uint16, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustIntegerNumber() uint16 {
	v, _ := i.IntegerNumber()
	return v
}
