// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPacketTMSI creates a new PacketTMSI IE.
func NewPacketTMSI(ptmsi uint32) *IE {
	return NewUint32IE(PacketTMSI, ptmsi)
}

// PacketTMSI returns PacketTMSI value in uint32 if type matches.
func (i *IE) PacketTMSI() (uint32, error) {
	if i.Type != PacketTMSI {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	return i.ValueAsUint32()
}

// MustPacketTMSI returns PacketTMSI in uint32, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustPacketTMSI() uint32 {
	v, _ := i.PacketTMSI()
	return v
}
