// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewMSValidated creates a new MSValidated IE.
func NewMSValidated(validated bool) *IE {
	if validated {
		return newUint8ValIE(MSValidated, 0xff)
	}
	return newUint8ValIE(MSValidated, 0xfe)
}

// MSValidated returns MSValidated in bool if type matches.
func (i *IE) MSValidated() bool {
	if i.Type != MSValidated {
		return false
	}
	if len(i.Payload) == 0 {
		return false
	}

	return i.Payload[0]%2 == 1
}
