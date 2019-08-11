// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "encoding/binary"

// NewChargingCharacteristics creates a new ChargingCharacteristics IE.
func NewChargingCharacteristics(chr uint16) *IE {
	return newUint16ValIE(ChargingCharacteristics, chr)
}

// ChargingCharacteristics returns the ChargingCharacteristics value in uint16 if the type of IE matches.
func (i *IE) ChargingCharacteristics() uint16 {
	if i.Type != ChargingCharacteristics {
		return 0
	}

	return binary.BigEndian.Uint16(i.Payload)
}
