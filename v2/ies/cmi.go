// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewCSGMembershipIndication creates a new CSGMembershipIndication IE.
func NewCSGMembershipIndication(cmi uint8) *IE {
	return newUint8ValIE(CSGMembershipIndication, cmi)
}

// CMI returns CMI in uint8 if the type of IE matches.
func (i *IE) CMI() uint8 {
	switch i.Type {
	case CSGMembershipIndication:
		return i.Payload[0] & 0x01
	case UserCSGInformation:
		return i.Payload[7] & 0x01
	default:
		return 0
	}
}
