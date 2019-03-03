// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewRATType creates a new RATType IE.
func NewRATType(rat uint8) *IE {
	return newUint8ValIE(RATType, rat)
}

// RATType returns RATType in uint8 if the type of IE matches.
func (i *IE) RATType() uint8 {
	switch i.Type {
	case RATType:
		return i.Payload[0]
	default:
		return 0
	}
}
