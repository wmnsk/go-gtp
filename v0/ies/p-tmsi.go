// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "encoding/binary"

// NewPacketTMSI creates a new PacketTMSI IE.
func NewPacketTMSI(ptmsi uint32) *IE {
	return newUint32ValIE(PacketTMSI, ptmsi)
}

// PacketTMSI returns PacketTMSI value in uint32 if type matches.
func (i *IE) PacketTMSI() uint32 {
	if i.Type != PacketTMSI {
		return 0
	}
	return binary.BigEndian.Uint32(i.Payload)
}
