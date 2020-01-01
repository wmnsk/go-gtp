// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"
	"io"
)

// NewTEIDDataI creates a new TEIDDataI IE.
func NewTEIDDataI(teid uint32) *IE {
	return newUint32ValIE(TEIDDataI, teid)
}

// NewTEIDCPlane creates a new TEID C-Plane IE.
func NewTEIDCPlane(teid uint32) *IE {
	return newUint32ValIE(TEIDCPlane, teid)
}

// NewTEIDDataII creates a new TEIDDataII IE.
func NewTEIDDataII(teid uint32) *IE {
	return newUint32ValIE(TEIDDataII, teid)
}

// TEID returns TEID value if type matches.
func (i *IE) TEID() (uint32, error) {
	if len(i.Payload) < 4 {
		return 0, io.ErrUnexpectedEOF
	}
	switch i.Type {
	case TEIDCPlane, TEIDDataI, TEIDDataII:
		return binary.BigEndian.Uint32(i.Payload), nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustTEID returns TEID in uint32 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustTEID() uint32 {
	v, _ := i.TEID()
	return v
}
