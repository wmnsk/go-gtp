// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"
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

	return i.Payload[0]&0x80>>7 == 1
}

// HasIPv6 reports whether the IE has IPv6 address in its payload or not.
func (i *IE) HasIPv6() bool {
	if i.Type != FullyQualifiedTEID {
		return false
	}

	return i.Payload[0]&0x48>>6 == 1
}

// InterfaceType returns InterfaceType in uint8 if the type of IE matches.
func (i *IE) InterfaceType() uint8 {
	if i.Type != FullyQualifiedTEID {
		return 0
	}

	return i.Payload[0] & 0x3f
}

// GREKey returns GREKey in uint32 if the type of IE matches.
func (i *IE) GREKey() uint32 {
	switch i.Type {
	case FullyQualifiedTEID:
		return binary.BigEndian.Uint32(i.Payload[1:5])
	case S103PDNDataForwardingInfo:
		switch i.Payload[0] {
		case 4:
			return binary.BigEndian.Uint32(i.Payload[5:9])
		case 16:
			return binary.BigEndian.Uint32(i.Payload[17:21])
		default:
			return 0
		}
	default:
		return 0
	}
}

// TEID returns TEID in uint32 if the type of IE matches.
func (i *IE) TEID() uint32 {
	switch i.Type {
	case FullyQualifiedTEID:
		return binary.BigEndian.Uint32(i.Payload[1:5])
	case S1UDataForwarding:
		switch i.Payload[0] {
		case 4:
			return binary.BigEndian.Uint32(i.Payload[5:9])
		case 16:
			return binary.BigEndian.Uint32(i.Payload[17:21])
		default:
			return 0
		}
	default:
		return 0
	}

}
