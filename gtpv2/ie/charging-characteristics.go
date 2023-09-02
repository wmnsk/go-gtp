// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewChargingCharacteristics creates a new ChargingCharacteristics IE.
func NewChargingCharacteristics(chr uint16) *IE {
	return newUint16ValIE(ChargingCharacteristics, chr)
}

// ChargingCharacteristics returns the ChargingCharacteristics value in uint16 if the type of IE matches.
func (i *IE) ChargingCharacteristics() (uint16, error) {
	if i.Type != ChargingCharacteristics {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint16(i.Payload), nil
}

// MustChargingCharacteristics returns ChargingCharacteristics in uint16, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustChargingCharacteristics() uint16 {
	v, _ := i.ChargingCharacteristics()
	return v
}
