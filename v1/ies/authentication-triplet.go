// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewAuthenticationTriplet creates a new AuthenticationTriplet IE.
func NewAuthenticationTriplet(rand, sres, kc []byte) *IE {
	i := New(AuthenticationTriplet, make([]byte, 28))

	copy(i.Payload[0:16], rand)
	copy(i.Payload[16:20], sres)
	copy(i.Payload[20:28], kc)
	return i
}

// AuthenticationTriplet returns AuthenticationTriplet in []byte if type matches.
func (i *IE) AuthenticationTriplet() []byte {
	if i.Type != AuthenticationTriplet {
		return nil
	}
	return i.Payload
}

// RAND returns RAND in []byte if type matches.
func (i *IE) RAND() []byte {
	switch i.Type {
	case AuthenticationTriplet, AuthenticationQuintuplet:
		return i.Payload[0:16]
	default:
		return nil
	}
}

// SRES returns SRES in []byte if type matches.
func (i *IE) SRES() []byte {
	switch i.Type {
	case AuthenticationTriplet:
		return i.Payload[16:20]
	default:
		return nil
	}
}

// Kc returns Kc in []byte if type matches.
func (i *IE) Kc() []byte {
	switch i.Type {
	case AuthenticationTriplet:
		return i.Payload[20:28]
	default:
		return nil
	}
}
