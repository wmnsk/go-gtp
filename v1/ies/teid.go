// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "encoding/binary"

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
func (i *IE) TEID() uint32 {
	switch i.Type {
	case TEIDCPlane, TEIDDataI, TEIDDataII:
		return binary.BigEndian.Uint32(i.Payload)
	default:
		return 0
	}
}
