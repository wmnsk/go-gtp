// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "io"

// NewPDNType creates a new PDNType IE.
func NewPDNType(pdn uint8) *IE {
	return newUint8ValIE(PDNType, pdn)
}

// PDNType returns the PDNType value in uint8 if the type of IE matches.
func (i *IE) PDNType() (uint8, error) {
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case PDNType, PDNAddressAllocation:
		return i.Payload[0], nil
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
