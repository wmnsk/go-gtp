// Copyright 2019-2024 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewServiceIndicator creates a new ServiceIndicator IE.
func NewServiceIndicator(ind uint8) *IE {
	return NewUint8IE(ServiceIndicator, ind)
}

// ServiceIndicator returns ServiceIndicator in uint8 if the type of IE matches.
func (i *IE) ServiceIndicator() (uint8, error) {
	if i.Type != ServiceIndicator {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	return i.ValueAsUint8()
}

// MustServiceIndicator returns ServiceIndicator in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustServiceIndicator() uint8 {
	v, _ := i.ServiceIndicator()
	return v
}
