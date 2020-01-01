// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "io"

// NewEPSBearerID creates a new EPSBearerID IE.
func NewEPSBearerID(ebi uint8) *IE {
	ebi &= 0x0f
	return newUint8ValIE(EPSBearerID, ebi)
}

// EPSBearerID returns EPSBearerID if the type of IE matches.
func (i *IE) EPSBearerID() (uint8, error) {
	if i.Type != EPSBearerID {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustEPSBearerID returns EPSBearerID in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustEPSBearerID() uint8 {
	v, _ := i.EPSBearerID()
	return v
}
