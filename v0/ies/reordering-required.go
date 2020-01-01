// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewReorderingRequired creates a new ReorderingRequired IE.
func NewReorderingRequired(required bool) *IE {
	if required {
		return New(ReorderingRequired, []byte{0xff})
	}
	return New(ReorderingRequired, []byte{0xfe})
}

// ReorderingRequired returns ReorderingRequired value in bool if type matches.
func (i *IE) ReorderingRequired() bool {
	if i.Type != ReorderingRequired {
		return false
	}
	if len(i.Payload) == 0 {
		return false
	}

	return i.Payload[0]&0x01 == 1
}
