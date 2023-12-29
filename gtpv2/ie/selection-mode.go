// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewSelectionMode creates a new SelectionMode IE.
func NewSelectionMode(mode uint8) *IE {
	return NewUint8IE(SelectionMode, mode)
}

// SelectionMode returns SelectionMode value if the type of IE matches.
func (i *IE) SelectionMode() (uint8, error) {
	if i.Type != SelectionMode {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	return i.ValueAsUint8()
}

// MustSelectionMode returns SelectionMode in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustSelectionMode() uint8 {
	v, _ := i.SelectionMode()
	return v
}
