// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

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
func (i *IE) AuthenticationQuintuplet() []byte {
	if i.Type != AuthenticationQuintuplet {
		return nil
	}
	return i.Payload
}

// XRES returns XRES in []byte if type matches.
func (i *IE) XRES() []byte {
	switch i.Type {
	case AuthenticationQuintuplet:
		xresLen := i.Payload[16]
		return i.Payload[17 : 17+xresLen]
	default:
		return nil
	}
}

// CK returns CK in []byte if type matches.
func (i *IE) CK() []byte {
	switch i.Type {
	case AuthenticationQuintuplet:
		offset := 17 + i.Payload[16]
		return i.Payload[offset : offset+16]
	default:
		return nil
	}
}

// IK returns IK in []byte if type matches.
func (i *IE) IK() []byte {
	switch i.Type {
	case AuthenticationQuintuplet:
		offset := 33 + i.Payload[16]
		return i.Payload[offset : offset+16]
	default:
		return nil
	}
}

// AUTN returns AUTN in []byte if type matches.
func (i *IE) AUTN() []byte {
	switch i.Type {
	case AuthenticationQuintuplet:
		offset := 49 + i.Payload[16]
		autnLen := i.Payload[50]
		return i.Payload[offset : offset+autnLen]
	default:
		return nil
	}
}
