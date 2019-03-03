// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewAPNRestriction creates a new APNRestriction IE.
func NewAPNRestriction(restriction uint8) *IE {
	return newUint8ValIE(APNRestriction, restriction)
}

// APNRestriction returns APNRestriction in uint8 if the type of IE matches.
func (i *IE) APNRestriction() uint8 {
	if i.Type != APNRestriction {
		return 0
	}

	return i.Payload[0]
}
