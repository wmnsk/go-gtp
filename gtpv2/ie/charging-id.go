// Copyright 2019-2024 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewChargingID creates a new ChargingID IE.
func NewChargingID(id uint32) *IE {
	return NewUint32IE(ChargingID, id)
}

// ChargingID returns the ChargingID value in uint32 if the type of IE matches.
func (i *IE) ChargingID() (uint32, error) {
	switch i.Type {
	case ChargingID:
		return i.ValueAsUint32()
	case BearerContext:
		ies, err := i.BearerContext()
		if err != nil {
			return 0, err
		}

		for _, child := range ies {
			if child.Type == ChargingID {
				return child.ChargingID()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustChargingID returns ChargingID in uint32, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustChargingID() uint32 {
	v, _ := i.ChargingID()
	return v
}
