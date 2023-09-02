// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewMSNotReachableReason creates a new MSNotReachableReason IE.
func NewMSNotReachableReason(reason uint8) *IE {
	return newUint8ValIE(MSNotReachableReason, reason)
}

// MSNotReachableReason returns MSNotReachableReason value if type matches.
func (i *IE) MSNotReachableReason() (uint8, error) {
	if i.Type != MSNotReachableReason {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustMSNotReachableReason returns MSNotReachableReason in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustMSNotReachableReason() uint8 {
	v, _ := i.MSNotReachableReason()
	return v
}
