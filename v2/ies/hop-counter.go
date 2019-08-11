// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewHopCounter creates a new HopCounter IE.
func NewHopCounter(hop uint8) *IE {
	return newUint8ValIE(HopCounter, hop)
}

// HopCounter returns HopCounter in uint8 if the type of IE matches.
func (i *IE) HopCounter() uint8 {
	if i.Type != HopCounter {
		return 0
	}
	if len(i.Payload) == 0 {
		return 0
	}

	return i.Payload[0]
}
