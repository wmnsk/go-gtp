// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewRATType creates a new RATType IE.
func NewRATType(ratType uint8) *IE {
	return New(
		RATType,
		[]byte{ratType},
	)
}

// RATType returns RATType value if type matches.
func (i *IE) RATType() uint8 {
	if i.Type != RATType {
		return 0
	}
	if len(i.Payload) == 0 {
		return 0
	}

	return i.Payload[0]
}
