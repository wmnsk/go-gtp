// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"
	"io"
)

// NewPortNumber creates a new PortNumber IE.
func NewPortNumber(port uint16) *IE {
	return newUint16ValIE(PortNumber, port)
}

// PortNumber returns PortNumber in uint16 if the type of IE matches.
func (i *IE) PortNumber() (uint16, error) {
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case PortNumber:
		if len(i.Payload) < 2 {
			return 0, io.ErrUnexpectedEOF
		}
		return binary.BigEndian.Uint16(i.Payload[0:2]), nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustPortNumber returns PortNumber in uint16, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustPortNumber() uint16 {
	v, _ := i.PortNumber()
	return v
}
