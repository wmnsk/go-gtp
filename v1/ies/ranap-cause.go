// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "io"

// NewRANAPCause creates a new RANAPCause IE.
func NewRANAPCause(cause uint8) *IE {
	return newUint8ValIE(RANAPCause, cause)
}

// RANAPCause returns RANAPCause in uint8 if type matches.
func (i *IE) RANAPCause() (uint8, error) {
	if i.Type != RANAPCause {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustRANAPCause returns RANAPCause in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustRANAPCause() uint8 {
	v, _ := i.RANAPCause()
	return v
}
