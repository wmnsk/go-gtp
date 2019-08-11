// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "encoding/binary"

// NewCSGID creates a new CSGID IE.
func NewCSGID(id uint32) *IE {
	return newUint32ValIE(CSGID, id&0x7ffffff)
}

// CSGID returns CSGID in uint32 if the type of IE matches.
func (i *IE) CSGID() uint32 {
	if len(i.Payload) < 4 {
		return 0
	}

	switch i.Type {
	case CSGID:
		return binary.BigEndian.Uint32(i.Payload[0:4]) & 0x7ffffff
	case UserCSGInformation:
		if len(i.Payload) < 7 {
			return 0
		}
		return binary.BigEndian.Uint32(i.Payload[3:7]) & 0x7ffffff
	default:
		return 0
	}
}
