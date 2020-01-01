// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "io"

// NewBearerFlags creates a new BearerFlags IE.
func NewBearerFlags(asi, vInd, vb, ppc uint8) *IE {
	i := New(BearerFlags, 0x00, make([]byte, 1))
	i.Payload[0] |= ((asi << 3 & 0x08) | (vInd << 2 & 0x04) | (vb << 1 & 0x2) | ppc&0x01)
	return i
}

// BearerFlags returns BearerFlags in uint8(=as it is) if the type of IE matches.
func (i *IE) BearerFlags() (uint8, error) {
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}
	if i.Type != BearerFlags {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload[0], nil
}

// MustBearerFlags returns BearerFlags in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustBearerFlags() uint8 {
	v, _ := i.BearerFlags()
	return v
}

// ActivityStatusIndicator reports whether the bearer context is preserved in
// the CN without corresponding Radio Access Bearer established.
func (i *IE) ActivityStatusIndicator() bool {
	if len(i.Payload) == 0 {
		return false
	}
	switch i.Type {
	case BearerFlags:
		return i.Payload[0]&0x08 == 1
	default:
		return false
	}
}

// VSRVCC reports whether this bearer is an IMS video bearer and is candidate
// for PS-to-CS vSRVCC handover.
func (i *IE) VSRVCC() bool {
	if len(i.Payload) == 0 {
		return false
	}
	switch i.Type {
	case BearerFlags:
		return i.Payload[0]&0x04 == 1
	default:
		return false
	}
}

// VoiceBearer reports whether a voice bearer when doing PS-to-CS (v)SRVCC handover.
func (i *IE) VoiceBearer() bool {
	if len(i.Payload) == 0 {
		return false
	}
	switch i.Type {
	case BearerFlags:
		return i.Payload[0]&0x02 == 1
	default:
		return false
	}
}

// ProhibitPayloadCompression reports whether an SGSN should attempt to
// compress the payload of user data when the users asks for it to be compressed.
func (i *IE) ProhibitPayloadCompression() bool {
	if len(i.Payload) == 0 {
		return false
	}
	switch i.Type {
	case BearerFlags:
		return i.Payload[0]&0x01 == 1
	default:
		return false
	}
}
