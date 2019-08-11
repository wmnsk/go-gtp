// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "encoding/binary"

// NewPortNumber creates a new PortNumber IE.
func NewPortNumber(port uint16) *IE {
	return newUint16ValIE(PortNumber, port)
}

// PortNumber returns PortNumber in uint16 if the type of IE matches.
func (i *IE) PortNumber() uint16 {
	if len(i.Payload) == 0 {
		return 0
	}

	switch i.Type {
	case PortNumber:
		if len(i.Payload) < 2 {
			return 0
		}
		return binary.BigEndian.Uint16(i.Payload[0:2])
	default:
		return 0
	}
}
