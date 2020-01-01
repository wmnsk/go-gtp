// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "io"

// NewHopCounter creates a new HopCounter IE.
func NewHopCounter(hop uint8) *IE {
	return newUint8ValIE(HopCounter, hop)
}

// HopCounter returns HopCounter in uint8 if the type of IE matches.
func (i *IE) HopCounter() (uint8, error) {
	if i.Type != HopCounter {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustHopCounter returns HopCounter in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustHopCounter() uint8 {
	v, _ := i.HopCounter()
	return v
}
