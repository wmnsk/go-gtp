// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewTeardownInd creates a new TeardownInd IE.
func NewTeardownInd(teardown bool) *IE {
	if teardown {
		return newUint8ValIE(TeardownInd, 0xff)
	}
	return newUint8ValIE(TeardownInd, 0xfe)
}

// TeardownInd returns TeardownInd in bool if type matches.
func (i *IE) TeardownInd() bool {
	if i.Type != TeardownInd {
		return false
	}
	if len(i.Payload) == 0 {
		return false
	}

	return i.Payload[0]%2 == 1
}
