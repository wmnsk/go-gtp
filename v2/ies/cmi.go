// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "io"

// NewCSGMembershipIndication creates a new CSGMembershipIndication IE.
func NewCSGMembershipIndication(cmi uint8) *IE {
	return newUint8ValIE(CSGMembershipIndication, cmi)
}

// CMI returns CMI in uint8 if the type of IE matches.
func (i *IE) CMI() (uint8, error) {
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case CSGMembershipIndication:
		return i.Payload[0] & 0x01, nil
	case UserCSGInformation:
		if len(i.Payload) < 8 {
			return 0, io.ErrUnexpectedEOF
		}
		return i.Payload[7] & 0x01, nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustCMI returns CMI in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustCMI() uint8 {
	v, _ := i.CMI()
	return v
}
