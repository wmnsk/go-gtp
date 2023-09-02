// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewRATType creates a new RATType IE.
func NewRATType(rat uint8) *IE {
	return newUint8ValIE(RATType, rat)
}

// RATType returns RATType in uint8 if the type of IE matches.
func (i *IE) RATType() (uint8, error) {
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case RATType:
		return i.Payload[0], nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustRATType returns RATType in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustRATType() uint8 {
	v, _ := i.RATType()
	return v
}
