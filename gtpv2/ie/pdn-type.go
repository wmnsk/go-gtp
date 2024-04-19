// Copyright 2019-2024 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPDNType creates a new PDNType IE.
func NewPDNType(pdn uint8) *IE {
	return NewUint8IE(PDNType, pdn)
}

// PDNType returns the PDNType value in uint8 if the type of IE matches.
func (i *IE) PDNType() (uint8, error) {
	switch i.Type {
	case PDNType, PDNAddressAllocation:
		return i.ValueAsUint8()
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustPDNType returns PDNType in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustPDNType() uint8 {
	v, _ := i.PDNType()
	return v
}
