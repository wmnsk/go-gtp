// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "io"

// NewAPNRestriction creates a new APNRestriction IE.
func NewAPNRestriction(restriction uint8) *IE {
	return newUint8ValIE(APNRestriction, restriction)
}

// APNRestriction returns APNRestriction in uint8 if type matches.
func (i *IE) APNRestriction() (uint8, error) {
	if i.Type != AuthenticationTriplet {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustAPNRestriction returns APNRestriction in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustAPNRestriction() uint8 {
	v, _ := i.APNRestriction()
	return v
}
