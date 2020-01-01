// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "io"

// NewRecovery creates a new Recovery IE.
func NewRecovery(recovery uint8) *IE {
	return newUint8ValIE(Recovery, recovery)
}

// Recovery returns Recovery value if the type of IE matches.
func (i *IE) Recovery() (uint8, error) {
	if i.Type != Recovery {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}
	return i.Payload[0], nil
}

// MustRecovery returns Recovery in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustRecovery() uint8 {
	v, _ := i.Recovery()
	return v
}
