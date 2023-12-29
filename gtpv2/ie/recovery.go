// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewRecovery creates a new Recovery IE.
func NewRecovery(recovery uint8) *IE {
	return NewUint8IE(Recovery, recovery)
}

// Recovery returns Recovery value if the type of IE matches.
func (i *IE) Recovery() (uint8, error) {
	if i.Type != Recovery {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	return i.ValueAsUint8()
}

// MustRecovery returns Recovery in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustRecovery() uint8 {
	v, _ := i.Recovery()
	return v
}
