// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewCause creates a new Cause IE.
func NewCause(cause uint8) *IE {
	return newUint8ValIE(Cause, cause)
}

// Cause returns Cause value if type matches.
func (i *IE) Cause() (uint8, error) {
	if i.Type != Cause {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustCause returns Cause in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustCause() uint8 {
	v, _ := i.Cause()
	return v
}
