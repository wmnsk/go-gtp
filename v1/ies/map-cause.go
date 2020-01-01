// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "io"

// NewMAPCause creates a new MAPCause IE.
func NewMAPCause(cause uint8) *IE {
	return newUint8ValIE(MAPCause, cause)
}

// MAPCause returns MAPCause in uint8 if type matches.
func (i *IE) MAPCause() (uint8, error) {
	if i.Type != MAPCause {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustMAPCause returns MAPCause in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustMAPCause() uint8 {
	v, _ := i.MAPCause()
	return v
}
