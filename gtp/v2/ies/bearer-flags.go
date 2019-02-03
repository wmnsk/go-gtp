// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewBearerFlags creates a new BearerFlags IE.
func NewBearerFlags(asi, vind, vb, ppc uint8) *IE {
	i := New(BearerFlags, 0x00, make([]byte, 1))
	i.Payload[0] |= ((asi << 3 & 0x08) | (vind << 2 & 0x04) | (vb << 1 & 0x2) | ppc&0x01)
	return i
}

// BearerFlags returns BearerFlags in uint8(=as it is) if the type of IE matches.
func (i *IE) BearerFlags() uint8 {
	if i.Type != BearerFlags {
		return 0
	}

	return i.Payload[0]
}

// ActivityStatusIndicator reports whether the bearer context is preserved in
// the CN without corresponding Radio Access Bearer established.
func (i *IE) ActivityStatusIndicator() bool {
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
	switch i.Type {
	case BearerFlags:
		return i.Payload[0]&0x04 == 1
	default:
		return false
	}
}

// VoiceBearer reports whether a voice bearer when doing PS-to-CS (v)SRVCC handover.
func (i *IE) VoiceBearer() bool {
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
	switch i.Type {
	case BearerFlags:
		return i.Payload[0]&0x01 == 1
	default:
		return false
	}
}
