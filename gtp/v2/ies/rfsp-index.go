// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewRFSPIndex creates a new RFSPIndex IE.
func NewRFSPIndex(idx uint8) *IE {
	return newUint8ValIE(RFSPIndex, idx)
}

// RFSPIndex returns RFSPIndex in uint8 if the type of IE matches.
func (i *IE) RFSPIndex() uint8 {
	if i.Type != RFSPIndex {
		return 0
	}

	return i.Payload[0]
}
