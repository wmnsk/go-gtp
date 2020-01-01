// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "net"

// PDN Type definitions.
const (
	_ uint8 = iota
	pdnTypeIPv4
	pdnTypeIPv6
	pdnTypeIPv4v6
	pdnTypeNonIP
)

// NewPDNAddressAllocation creates a new PDNAddressAllocation IE.
//
// The PDN Type field is automatically judged by the format of given addr,
// If it cannot be converted as neither IPv4 nor IPv6, PDN Type will be Non-IP.
func NewPDNAddressAllocation(addr string) *IE {
	ip := net.ParseIP(addr)
	v4 := ip.To4()

	// IPv4
	if v4 != nil {
		i := New(PDNAddressAllocation, 0x00, make([]byte, 5))
		i.Payload[0] = pdnTypeIPv4
		copy(i.Payload[1:], v4)
		return i
	}

	// IPv6
	// XXX - prefix value should be handled properly.
	if ip != nil {
		i := New(PDNAddressAllocation, 0x00, make([]byte, 18))
		i.Payload[0] = pdnTypeIPv6
		i.Payload[1] = 0x00
		copy(i.Payload[2:], ip)
		return i
	}

	// Non-IP
	return New(PDNAddressAllocation, 0x00, []byte{pdnTypeNonIP})
}

// NewPDNAddressAllocationDual creates a new PDNAddressAllocation IE with
// IPv4 address and IPv6 address given.
//
// If they cannot be converted as IPv4/IPv6, PDN Type will be Non-IP.
func NewPDNAddressAllocationDual(v4addr, v6addr string) *IE {
	v4 := net.ParseIP(v4addr).To4()
	if v4 == nil {
		return New(PDNAddressAllocation, 0x00, []byte{pdnTypeNonIP})
	}

	v6 := net.ParseIP(v6addr).To16()
	if v6 == nil {
		return New(PDNAddressAllocation, 0x00, []byte{pdnTypeNonIP})
	}

	i := New(PDNAddressAllocation, 0x00, make([]byte, 23))
	i.Payload[0] = pdnTypeIPv4v6
	copy(i.Payload[1:5], v4)
	i.Payload[5] = 0x00
	copy(i.Payload[6:], v6)

	return i
}
