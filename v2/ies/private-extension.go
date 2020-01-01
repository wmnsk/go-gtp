// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"
	"io"
)

// NewPrivateExtension creates a new PrivateExtension IE.
func NewPrivateExtension(id uint16, value []byte) *IE {
	i := New(PrivateExtension, 0x00, make([]byte, 2+len(value)))
	binary.BigEndian.PutUint16(i.Payload[0:2], id)
	copy(i.Payload[2:], value)
	return i
}

// EnterpriseID returns EnterpriseID in uint16 if the type of IE matches.
func (i *IE) EnterpriseID() (uint16, error) {
	if i.Type != PrivateExtension {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint16(i.Payload[0:2]), nil

}

// MustEnterpriseID returns EnterpriseID in uint16, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustEnterpriseID() uint16 {
	v, _ := i.EnterpriseID()
	return v
}

// PrivateExtension returns PrivateExtension value in []byte if the type of IE matches.
func (i *IE) PrivateExtension() ([]byte, error) {
	if i.Type != PrivateExtension {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 3 {
		return nil, io.ErrUnexpectedEOF
	}

	return i.Payload[2:], nil
}

// MustPrivateExtension returns PrivateExtension in []byte, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustPrivateExtension() []byte {
	v, _ := i.PrivateExtension()
	return v
}
