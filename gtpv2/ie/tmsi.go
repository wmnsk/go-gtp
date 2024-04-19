// Copyright 2019-2024 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewTMSI creates a new TMSI IE.
func NewTMSI(tmsi uint32) *IE {
	return NewUint32IE(TMSI, tmsi)
}

// TMSI returns TMSI in uint32 if the type of IE matches.
func (i *IE) TMSI() (uint32, error) {
	switch i.Type {
	case TMSI:
		return i.ValueAsUint32()
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustTMSI returns TMSI in uint32, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustTMSI() uint32 {
	v, _ := i.TMSI()
	return v
}
