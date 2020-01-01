// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"
	"io"
)

// NewPacketTMSI creates a new PacketTMSI IE.
func NewPacketTMSI(ptmsi uint32) *IE {
	return newUint32ValIE(PacketTMSI, ptmsi)
}

// PacketTMSI returns PacketTMSI value in uint32 if type matches.
func (i *IE) PacketTMSI() (uint32, error) {
	if i.Type != PacketTMSI {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint32(i.Payload), nil
}

// MustPacketTMSI returns PacketTMSI in uint32 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustPacketTMSI() uint32 {
	v, _ := i.PacketTMSI()
	return v
}
