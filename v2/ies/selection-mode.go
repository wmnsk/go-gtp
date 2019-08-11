// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewSelectionMode creates a new SelectionMode IE.
//
// Note that exactly one of the parameters should be set to true.
// Otherwise, you'll get the unexpected result.
func NewSelectionMode(mode uint8) *IE {
	return newUint8ValIE(SelectionMode, mode)
}

// SelectionMode returns SelectionMode value if the type of IE matches.
func (i *IE) SelectionMode() uint8 {
	if i.Type != SelectionMode {
		return 0
	}
	return i.Payload[0]
}
