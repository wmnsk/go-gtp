// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"
	"io"
	"net"
)

// NewS103PDNDataForwardingInfo creates a new S103PDNDataForwardingInfo IE.
func NewS103PDNDataForwardingInfo(hsgwAddr string, greKey uint32, ebis ...uint8) *IE {
	addr := net.ParseIP(hsgwAddr)
	if addr == nil {
		return nil
	}

	// HSGW Address: IPv4
	if v4 := addr.To4(); v4 != nil {
		i := New(S103PDNDataForwardingInfo, 0x00, make([]byte, 1+4+4+1+len(ebis)))
		i.Payload[0] = 4
		copy(i.Payload[1:5], v4)
		binary.BigEndian.PutUint32(i.Payload[5:9], greKey)
		i.Payload[9] = uint8(len(ebis))
		for n, e := range ebis {
			i.Payload[10+n] = e & 0x0f
		}
		return i
	}

	// HSGW Address: IPv6
	i := New(S103PDNDataForwardingInfo, 0x00, make([]byte, 1+16+4+1+len(ebis)))
	i.Payload[0] = 16
	copy(i.Payload[1:17], addr)
	binary.BigEndian.PutUint32(i.Payload[17:21], greKey)
	i.Payload[21] = uint8(len(ebis))
	for n, e := range ebis {
		i.Payload[22+n] = e & 0x0f
	}
	return i
}

// HSGWAddress returns IP address of HSGW in string if the type of IE matches.
func (i *IE) HSGWAddress() (string, error) {
	if i.Type != S103PDNDataForwardingInfo {
		return "", &InvalidTypeError{Type: i.Type}
	}

	return i.IPAddress()
}

// MustHSGWAddress returns HSGWAddress in string, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustHSGWAddress() string {
	v, _ := i.HSGWAddress()
	return v
}

// EBIs returns the EBIs in []uint8 if the type of IE matches.
func (i *IE) EBIs() ([]uint8, error) {
	if i.Type != S103PDNDataForwardingInfo {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return nil, io.ErrUnexpectedEOF
	}

	var n, offset int
	switch i.Payload[0] {
	case 4:
		if len(i.Payload) < 9 {
			return nil, io.ErrUnexpectedEOF
		}
		n = int(i.Payload[9])
		offset = 10
	case 16:
		if len(i.Payload) < 21 {
			return nil, io.ErrUnexpectedEOF
		}
		n = int(i.Payload[21])
		offset = 22
	default:
		return nil, ErrMalformed
	}

	var ebis []uint8
	for x := 0; x < n; x++ {
		if len(i.Payload) <= offset+x {
			break
		}
		ebis = append(ebis, i.Payload[offset+x])
	}
	return ebis, nil
}

// MustEBIs returns EBIs in []uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustEBIs() []uint8 {
	v, _ := i.EBIs()
	return v
}
