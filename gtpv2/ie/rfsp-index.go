// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewRFSPIndex creates a new RFSPIndex IE.
func NewRFSPIndex(idx uint8) *IE {
	return NewUint8IE(RFSPIndex, idx)
}

// RFSPIndex returns RFSPIndex in uint8 if the type of IE matches.
func (i *IE) RFSPIndex() (uint8, error) {
	if i.Type != RFSPIndex {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	return i.ValueAsUint8()
}

// MustRFSPIndex returns RFSPIndex in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustRFSPIndex() uint8 {
	v, _ := i.RFSPIndex()
	return v
}
