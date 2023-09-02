// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewReorderingRequired creates a new ReorderingRequired IE.
func NewReorderingRequired(required bool) *IE {
	if required {
		return newUint8ValIE(ReorderingRequired, 0xff)
	}
	return newUint8ValIE(ReorderingRequired, 0xfe)
}

// ReorderingRequired returns ReorderingRequired or not if type matches.
func (i *IE) ReorderingRequired() bool {
	if i.Type != ReorderingRequired {
		return false
	}
	if len(i.Payload) == 0 {
		return false
	}

	return i.Payload[0]&0x01 == 1
}
