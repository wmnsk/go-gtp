// Copyright 2019-2024 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPortNumber creates a new PortNumber IE.
func NewPortNumber(port uint16) *IE {
	return NewUint16IE(PortNumber, port)
}

// PortNumber returns PortNumber in uint16 if the type of IE matches.
func (i *IE) PortNumber() (uint16, error) {
	switch i.Type {
	case PortNumber:
		return i.ValueAsUint16()
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustPortNumber returns PortNumber in uint16, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustPortNumber() uint16 {
	v, _ := i.PortNumber()
	return v
}
