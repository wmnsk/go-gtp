// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"
	"io"
	"net"
)

// NewFullyQualifiedTEID creates a new FullyQualifiedTEID IE.
func NewFullyQualifiedTEID(ifType uint8, teid uint32, v4, v6 string) *IE {
	i := New(FullyQualifiedTEID, 0x00, make([]byte, 5))
	i.Payload[0] = ifType
	binary.BigEndian.PutUint32(i.Payload[1:5], teid)

	if v4addr := net.ParseIP(v4); v4addr != nil {
		i.Payload[0] |= 0x80
		i.Payload = append(i.Payload, []byte(v4addr.To4())...)
	}
	if v6addr := net.ParseIP(v6); v6addr != nil {
		i.Payload[0] |= 0x40
		i.Payload = append(i.Payload, []byte(v6addr.To16())...)
	}
	i.SetLength()

	return i
}

// HasIPv4 reports whether the IE has IPv4 address in its payload or not.
func (i *IE) HasIPv4() bool {
	if i.Type != FullyQualifiedTEID {
		return false
	}
	if len(i.Payload) == 0 {
		return false
	}

	return i.Payload[0]&0x80>>7 == 1
}

// HasIPv6 reports whether the IE has IPv6 address in its payload or not.
func (i *IE) HasIPv6() bool {
	if i.Type != FullyQualifiedTEID {
		return false
	}
	if len(i.Payload) == 0 {
		return false
	}

	return i.Payload[0]&0x48>>6 == 1
}

// InterfaceType returns InterfaceType in uint8 if the type of IE matches.
func (i *IE) InterfaceType() (uint8, error) {
	if i.Type != FullyQualifiedTEID {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0] & 0x3f, nil
}

// MustInterfaceType returns InterfaceType in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustInterfaceType() uint8 {
	v, _ := i.InterfaceType()
	return v
}

// GREKey returns GREKey in uint32 if the type of IE matches.
func (i *IE) GREKey() (uint32, error) {
	if len(i.Payload) < 6 {
		return 0, io.ErrUnexpectedEOF
	}
	switch i.Type {
	case FullyQualifiedTEID:
		return binary.BigEndian.Uint32(i.Payload[1:5]), nil
	case S103PDNDataForwardingInfo:
		switch i.Payload[0] {
		case 4:
			if len(i.Payload) < 9 {
				return 0, io.ErrUnexpectedEOF
			}
			return binary.BigEndian.Uint32(i.Payload[5:9]), nil
		case 16:
			if len(i.Payload) < 21 {
				return 0, io.ErrUnexpectedEOF
			}
			return binary.BigEndian.Uint32(i.Payload[17:21]), nil
		default:
			return 0, ErrMalformed
		}
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustGREKey returns GREKey in uint32, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustGREKey() uint32 {
	v, _ := i.GREKey()
	return v
}

// TEID returns TEID in uint32 if the type of IE matches.
func (i *IE) TEID() (uint32, error) {
	if len(i.Payload) < 5 {
		return 0, io.ErrUnexpectedEOF
	}
	switch i.Type {
	case FullyQualifiedTEID:
		return binary.BigEndian.Uint32(i.Payload[1:5]), nil
	case S1UDataForwarding:
		switch i.Payload[0] {
		case 4:
			if len(i.Payload) < 9 {
				return 0, io.ErrUnexpectedEOF
			}
			return binary.BigEndian.Uint32(i.Payload[5:9]), nil
		case 16:
			if len(i.Payload) < 21 {
				return 0, io.ErrUnexpectedEOF
			}
			return binary.BigEndian.Uint32(i.Payload[17:21]), nil
		default:
			return 0, ErrMalformed
		}
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}

}

// MustTEID returns TEID in uint32, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustTEID() uint32 {
	v, _ := i.TEID()
	return v
}
