// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewRecovery creates a new Recovery IE.
func NewRecovery(recovery uint8) *IE {
	return newUint8ValIE(Recovery, recovery)
}

// Recovery returns Recovery value if the type of IE matches.
func (i *IE) Recovery() uint8 {
	if i.Type != Recovery {
		return 0
	}
	return i.Payload[0]
}
