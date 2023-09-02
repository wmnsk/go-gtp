// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"fmt"
	"io"
)

// NewBearerFlags creates a new BearerFlags IE.
func NewBearerFlags(asi, vInd, vb, ppc uint8) *IE {
	i := New(BearerFlags, 0x00, make([]byte, 1))
	i.Payload[0] |= ((asi << 3 & 0x08) | (vInd << 2 & 0x04) | (vb << 1 & 0x2) | ppc&0x01)
	return i
}

// BearerFlags returns BearerFlags in uint8(=as it is) if the type of IE matches.
func (i *IE) BearerFlags() (uint8, error) {
	switch i.Type {
	case BearerFlags:
		if len(i.Payload) < 1 {
			return 0, io.ErrUnexpectedEOF
		}

		return i.Payload[0], nil
	case BearerContext:
		ies, err := i.BearerContext()
		if err != nil {
			return 0, fmt.Errorf("failed to retrieve BearerFlags: %w", err)
		}

		for _, child := range ies {
			if child.Type == BearerFlags {
				return child.BearerFlags()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustBearerFlags returns BearerFlags in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustBearerFlags() uint8 {
	v, _ := i.BearerFlags()
	return v
}

// HasPPC reports whether an IE has PPC bit.
func (i *IE) HasPPC() bool {
	v, err := i.BearerFlags()
	if err != nil {
		return false
	}

	return has1stBit(v)
}

// HasVB reports whether an IE has VB bit.
func (i *IE) HasVB() bool {
	v, err := i.BearerFlags()
	if err != nil {
		return false
	}

	return has2ndBit(v)
}

// HasVind reports whether an IE has Vind bit.
func (i *IE) HasVind() bool {
	v, err := i.BearerFlags()
	if err != nil {
		return false
	}

	return has3rdBit(v)
}

// HasASI reports whether an IE has ASI bit.
func (i *IE) HasASI() bool {
	v, err := i.BearerFlags()
	if err != nil {
		return false
	}

	return has4thBit(v)
}

// ActivityStatusIndicator reports whether the bearer context is preserved in
// the CN without corresponding Radio Access Bearer established.
func (i *IE) ActivityStatusIndicator() bool {
	if len(i.Payload) < 1 {
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
	if len(i.Payload) < 1 {
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
	if len(i.Payload) < 1 {
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
	if len(i.Payload) < 1 {
		return false
	}
	switch i.Type {
	case BearerFlags:
		return i.Payload[0]&0x01 == 1
	default:
		return false
	}
}
