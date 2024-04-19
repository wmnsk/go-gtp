// Copyright 2019-2024 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewIntegerNumber creates a new IntegerNumber IE.
func NewIntegerNumber(num uint16) *IE {
	return NewUint16IE(IntegerNumber, num)
}

// IntegerNumber returns IntegerNumber in uint16 if the type of IE matches.
func (i *IE) IntegerNumber() (uint16, error) {
	switch i.Type {
	case IntegerNumber:
		return i.ValueAsUint16()
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustIntegerNumber returns IntegerNumber in uint16, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustIntegerNumber() uint16 {
	v, _ := i.IntegerNumber()
	return v
}
