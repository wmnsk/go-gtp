// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewRATType creates a new RATType IE.
func NewRATType(ratType uint8) *IE {
	return New(
		RATType,
		[]byte{ratType},
	)
}

// RATType returns RATType value if type matches.
func (i *IE) RATType() (uint8, error) {
	if i.Type != RATType {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustRATType returns RATType in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustRATType() uint8 {
	v, _ := i.RATType()
	return v
}
