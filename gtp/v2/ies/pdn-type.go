// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewPDNType creates a new PDNType IE.
func NewPDNType(pdn uint8) *IE {
	return newUint8ValIE(PDNType, pdn)
}

// PDNType returns the PDNType value in uint8 if the type of IE matches.
func (i *IE) PDNType() uint8 {
	switch i.Type {
	case PDNType, PDNAddressAllocation:
		return i.Payload[0]
	default:
		return 0
	}
}
