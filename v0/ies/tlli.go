// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"
	"io"
)

// NewTemporaryLogicalLinkIdentity creates a new TemporaryLogicalLinkIdentity IE.
func NewTemporaryLogicalLinkIdentity(tlli uint32) *IE {
	return newUint32ValIE(TemporaryLogicalLinkIdentity, tlli)
}

// TemporaryLogicalLinkIdentity returns TemporaryLogicalLinkIdentity value in uint32 if type matches.
func (i *IE) TemporaryLogicalLinkIdentity() (uint32, error) {
	if i.Type != TemporaryLogicalLinkIdentity {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint32(i.Payload), nil
}

// MustTemporaryLogicalLinkIdentity returns TemporaryLogicalLinkIdentity in uint32 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustTemporaryLogicalLinkIdentity() uint32 {
	v, _ := i.TemporaryLogicalLinkIdentity()
	return v
}
