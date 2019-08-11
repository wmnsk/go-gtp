// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "encoding/binary"

// NewTMSI creates a new TMSI IE.
func NewTMSI(tmsi uint32) *IE {
	return newUint32ValIE(TMSI, tmsi)
}

// TMSI returns TMSI in uint32 if the type of IE matches.
func (i *IE) TMSI() uint32 {
	switch i.Type {
	case TMSI:
		return binary.BigEndian.Uint32(i.Payload)
	default:
		return 0
	}
}
