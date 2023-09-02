// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewAuthenticationQuintuplet creates a new AuthenticationQuintuplet IE.
func NewAuthenticationQuintuplet(rand, xres, ck, ik, autn []byte) *IE {
	xresLen := len(xres)
	autnLen := len(autn)
	i := New(AuthenticationQuintuplet, make([]byte, 16+1+xresLen+16+16+1+autnLen))

	copy(i.Payload[0:16], rand)
	i.Payload[16] = uint8(xresLen)
	offset := 17 // variable length appears from here.
	copy(i.Payload[offset:offset+xresLen], xres)
	offset += xresLen
	copy(i.Payload[offset:offset+16], ck)
	offset += 16
	copy(i.Payload[offset:offset+16], ik)
	offset += 16
	i.Payload[offset] = uint8(autnLen)
	offset++
	copy(i.Payload[offset:offset+autnLen], autn)

	return i
}

// AuthenticationQuintuplet returns AuthenticationQuintuplet in []byte if type matches.
func (i *IE) AuthenticationQuintuplet() ([]byte, error) {
	if i.Type != AuthenticationQuintuplet {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return nil, io.ErrUnexpectedEOF
	}

	return i.Payload, nil
}

// MustAuthenticationQuintuplet returns AuthenticationQuintuplet in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustAuthenticationQuintuplet() []byte {
	v, _ := i.AuthenticationQuintuplet()
	return v
}

// XRES returns XRES in []byte if type matches.
func (i *IE) XRES() ([]byte, error) {
	if len(i.Payload) == 0 {
		return nil, nil
	}

	switch i.Type {
	case AuthenticationQuintuplet:
		if len(i.Payload) < 17 {
			return nil, io.ErrUnexpectedEOF
		}

		xresLen := i.Payload[16]
		if len(i.Payload) < 17+int(xresLen) {
			return nil, io.ErrUnexpectedEOF
		}
		return i.Payload[17 : 17+int(xresLen)], nil
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MustXRES returns XRES in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustXRES() []byte {
	v, _ := i.XRES()
	return v
}

// CK returns CK in []byte if type matches.
func (i *IE) CK() ([]byte, error) {
	switch i.Type {
	case AuthenticationQuintuplet:
		if len(i.Payload) < 18 {
			return nil, io.ErrUnexpectedEOF
		}

		offset := 17 + int(i.Payload[16])
		if len(i.Payload) < offset+16 {
			return nil, io.ErrUnexpectedEOF
		}
		return i.Payload[offset : offset+16], nil
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MustCK returns CK in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustCK() []byte {
	v, _ := i.CK()
	return v
}

// IK returns IK in []byte if type matches.
func (i *IE) IK() ([]byte, error) {
	switch i.Type {
	case AuthenticationQuintuplet:
		if len(i.Payload) < 34 {
			return nil, io.ErrUnexpectedEOF
		}

		offset := 33 + int(i.Payload[16])
		if len(i.Payload) < offset+16 {
			return nil, io.ErrUnexpectedEOF
		}
		return i.Payload[offset : offset+16], nil
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MustIK returns IK in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustIK() []byte {
	v, _ := i.IK()
	return v
}

// AUTN returns AUTN in []byte if type matches.
func (i *IE) AUTN() ([]byte, error) {
	switch i.Type {
	case AuthenticationQuintuplet:
		if len(i.Payload) < 50 {
			return nil, io.ErrUnexpectedEOF
		}
		offset := 49 + int(i.Payload[16])
		autnLen := i.Payload[50]
		if len(i.Payload) < offset+int(autnLen) {
			return nil, io.ErrUnexpectedEOF
		}
		return i.Payload[offset : offset+int(autnLen)], nil
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MustAUTN returns AUTN in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustAUTN() []byte {
	v, _ := i.AUTN()
	return v
}
