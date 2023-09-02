// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewTMSI creates a new TMSI IE.
func NewTMSI(tmsi uint32) *IE {
	return newUint32ValIE(TMSI, tmsi)
}

// TMSI returns TMSI in uint32 if the type of IE matches.
func (i *IE) TMSI() (uint32, error) {
	if len(i.Payload) < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case TMSI:
		return binary.BigEndian.Uint32(i.Payload), nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustTMSI returns TMSI in uint32, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustTMSI() uint32 {
	v, _ := i.TMSI()
	return v
}
