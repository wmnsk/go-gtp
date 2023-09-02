// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewSelectionMode creates a new SelectionMode IE.
// Note that exactly one of the parameters should be set to true.
// Otherwise, you'll get the unexpected result.
func NewSelectionMode(mode uint8) *IE {
	return newUint8ValIE(SelectionMode, mode)
}

// SelectionMode returns SelectionMode value if type matches.
func (i *IE) SelectionMode() (uint8, error) {
	if i.Type != SelectionMode {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustSelectionMode returns SelectionMode in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustSelectionMode() uint8 {
	v, _ := i.SelectionMode()
	return v
}
