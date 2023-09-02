// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewNSAPI creates a new NSAPI IE.
func NewNSAPI(nsapi uint8) *IE {
	return newUint8ValIE(NSAPI, nsapi)
}

// NSAPI returns NSAPI value if type matches.
func (i *IE) NSAPI() (uint8, error) {
	if i.Type != NSAPI {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustNSAPI returns NSAPI in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustNSAPI() uint8 {
	v, _ := i.NSAPI()
	return v
}
