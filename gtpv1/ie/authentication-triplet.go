// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewAuthenticationTriplet creates a new AuthenticationTriplet IE.
func NewAuthenticationTriplet(rand, sres, kc []byte) *IE {
	i := New(AuthenticationTriplet, make([]byte, 28))

	copy(i.Payload[0:16], rand)
	copy(i.Payload[16:20], sres)
	copy(i.Payload[20:28], kc)
	return i
}

// AuthenticationTriplet returns AuthenticationTriplet in []byte if type matches.
func (i *IE) AuthenticationTriplet() ([]byte, error) {
	if i.Type != AuthenticationTriplet {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	return i.Payload, nil
}

// MustAuthenticationTriplet returns AuthenticationTriplet in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustAuthenticationTriplet() []byte {
	v, _ := i.AuthenticationTriplet()
	return v
}

// RAND returns RAND in []byte if type matches.
func (i *IE) RAND() ([]byte, error) {
	switch i.Type {
	case AuthenticationTriplet, AuthenticationQuintuplet:
		if len(i.Payload) < 16 {
			return nil, io.ErrUnexpectedEOF
		}
		return i.Payload[0:16], nil
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MustRAND returns RAND in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustRAND() []byte {
	v, _ := i.RAND()
	return v
}

// SRES returns SRES in []byte if type matches.
func (i *IE) SRES() ([]byte, error) {
	switch i.Type {
	case AuthenticationTriplet:
		if len(i.Payload) < 20 {
			return nil, io.ErrUnexpectedEOF
		}
		return i.Payload[16:20], nil
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MustSRES returns SRES in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustSRES() []byte {
	v, _ := i.SRES()
	return v
}

// Kc returns Kc in []byte if type matches.
func (i *IE) Kc() ([]byte, error) {
	switch i.Type {
	case AuthenticationTriplet:
		if len(i.Payload) < 28 {
			return nil, io.ErrUnexpectedEOF
		}
		return i.Payload[20:28], nil
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MustKc returns Kc in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustKc() []byte {
	v, _ := i.Kc()
	return v
}
