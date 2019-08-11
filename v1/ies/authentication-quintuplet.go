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
	if len(i.Payload) == 0 {
		return nil
	}

	return i.Payload
}

// XRES returns XRES in []byte if type matches.
func (i *IE) XRES() []byte {
	if len(i.Payload) == 0 {
		return nil
	}

	switch i.Type {
	case AuthenticationQuintuplet:
		if len(i.Payload) < 17 {
			return nil
		}

		xresLen := i.Payload[16]
		if len(i.Payload) < 17+int(xresLen) {
			return nil
		}
		return i.Payload[17 : 17+int(xresLen)]
	default:
		return nil
	}
}

// CK returns CK in []byte if type matches.
func (i *IE) CK() []byte {
	switch i.Type {
	case AuthenticationQuintuplet:
		if len(i.Payload) < 18 {
			return nil
		}

		offset := 17 + int(i.Payload[16])
		if len(i.Payload) < offset+16 {
			return nil
		}
		return i.Payload[offset : offset+16]
	default:
		return nil
	}
}

// IK returns IK in []byte if type matches.
func (i *IE) IK() []byte {
	switch i.Type {
	case AuthenticationQuintuplet:
		if len(i.Payload) < 34 {
			return nil
		}

		offset := 33 + int(i.Payload[16])
		if len(i.Payload) < offset+16 {
			return nil
		}
		return i.Payload[offset : offset+16]
	default:
		return nil
	}
}

// AUTN returns AUTN in []byte if type matches.
func (i *IE) AUTN() []byte {
	switch i.Type {
	case AuthenticationQuintuplet:
		if len(i.Payload) < 50 {
			return nil
		}
		offset := 49 + int(i.Payload[16])
		autnLen := i.Payload[50]
		if len(i.Payload) < offset+int(autnLen) {
			return nil
		}
		return i.Payload[offset : offset+int(autnLen)]
	default:
		return nil
	}
}
