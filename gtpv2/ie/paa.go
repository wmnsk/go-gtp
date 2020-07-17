// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

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
	return NewPDNAddressAllocationNetIP(net.ParseIP(addr))
}

// NewPDNAddressAllocationDual creates a new PDNAddressAllocation IE with
// IPv4 address and IPv6 address given.
//
// If they cannot be converted as IPv4/IPv6, PDN Type will be Non-IP.
func NewPDNAddressAllocationDual(v4addr, v6addr string) *IE {
	return NewPDNAddressAllocationDualNetIP(net.ParseIP(v4addr), net.ParseIP(v6addr))
}

// NewPDNAddressAllocationNetIP creates a new PDNAddressAllocation IE from net.IP.
func NewPDNAddressAllocationNetIP(ip net.IP) *IE {
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
	if ip.To16() != nil {
		i := New(PDNAddressAllocation, 0x00, make([]byte, 18))
		i.Payload[0] = pdnTypeIPv6
		i.Payload[1] = 0x00
		copy(i.Payload[2:], ip)
		return i
	}

	// Non-IP
	return New(PDNAddressAllocation, 0x00, []byte{pdnTypeNonIP})
}

// NewPDNAddressAllocationDualNetIP creates a new PDNAddressAllocation IE from
// IPv4 and IPv6 in net.IP.
func NewPDNAddressAllocationDualNetIP(v4, v6 net.IP) *IE {
	if v4.To4() == nil {
		return New(PDNAddressAllocation, 0x00, []byte{pdnTypeNonIP})
	}

	if v6.To16() == nil {
		return New(PDNAddressAllocation, 0x00, []byte{pdnTypeNonIP})
	}

	i := New(PDNAddressAllocation, 0x00, make([]byte, 23))
	i.Payload[0] = pdnTypeIPv4v6
	copy(i.Payload[1:5], v4)
	i.Payload[5] = 0x00
	copy(i.Payload[6:], v6)

	return i
}
