// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "io"

// NewRFSPIndex creates a new RFSPIndex IE.
func NewRFSPIndex(idx uint8) *IE {
	return newUint8ValIE(RFSPIndex, idx)
}

// RFSPIndex returns RFSPIndex in uint8 if the type of IE matches.
func (i *IE) RFSPIndex() (uint8, error) {
	if i.Type != RFSPIndex {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustRFSPIndex returns RFSPIndex in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustRFSPIndex() uint8 {
	v, _ := i.RFSPIndex()
	return v
}
