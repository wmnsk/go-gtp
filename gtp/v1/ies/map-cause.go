// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewMAPCause creates a new MAPCause IE.
func NewMAPCause(cause uint8) *IE {
	return newUint8ValIE(MAPCause, cause)
}

// MAPCause returns MAPCause in uint8 if type matches.
func (i *IE) MAPCause() uint8 {
	if i.Type != AuthenticationTriplet {
		return 0
	}
	return i.Payload[0]
}
