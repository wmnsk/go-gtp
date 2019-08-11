// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewDetachType creates a new DetachType IE.
func NewDetachType(dtype uint8) *IE {
	return newUint8ValIE(DetachType, dtype)
}

// DetachType returns DetachType in uint8 if the type of IE matches.
func (i *IE) DetachType() uint8 {
	if i.Type != DetachType {
		return 0
	}
	if len(i.Payload) == 0 {
		return 0
	}

	return i.Payload[0]
}
