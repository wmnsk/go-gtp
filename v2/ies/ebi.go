// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewEPSBearerID creates a new EPSBearerID IE.
func NewEPSBearerID(ebi uint8) *IE {
	ebi &= 0x0f
	return newUint8ValIE(EPSBearerID, ebi)
}

// EPSBearerID returns EPSBearerID if the type of IE matches.
func (i *IE) EPSBearerID() uint8 {
	if i.Type != EPSBearerID {
		return 0
	}

	return i.Payload[0]
}
