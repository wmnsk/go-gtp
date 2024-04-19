// Copyright 2019-2024 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewChargingCharacteristics creates a new ChargingCharacteristics IE.
func NewChargingCharacteristics(chr uint16) *IE {
	return NewUint16IE(ChargingCharacteristics, chr)
}

// ChargingCharacteristics returns the ChargingCharacteristics value in uint16 if the type of IE matches.
func (i *IE) ChargingCharacteristics() (uint16, error) {
	if i.Type != ChargingCharacteristics {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	return i.ValueAsUint16()
}

// MustChargingCharacteristics returns ChargingCharacteristics in uint16, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustChargingCharacteristics() uint16 {
	v, _ := i.ChargingCharacteristics()
	return v
}
