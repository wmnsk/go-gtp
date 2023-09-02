// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewServiceIndicator creates a new ServiceIndicator IE.
func NewServiceIndicator(ind uint8) *IE {
	return newUint8ValIE(ServiceIndicator, ind)
}

// ServiceIndicator returns ServiceIndicator in uint8 if the type of IE matches.
func (i *IE) ServiceIndicator() (uint8, error) {
	if i.Type != ServiceIndicator {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustServiceIndicator returns ServiceIndicator in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustServiceIndicator() uint8 {
	v, _ := i.ServiceIndicator()
	return v
}
