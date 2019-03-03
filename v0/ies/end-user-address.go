// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "net"

// PDP Type Organization definitions.
const (
	pdpTypeETSI uint8 = iota | 0xf0
	pdpTypeIETF
)

// NewEndUserAddress creates a new EndUserAddress IE from the given IP Address in string.
//
// The addr can be either IPv4 or IPv6. If the address type is PPP,
// just put "ppp" in addr parameter or use NewEndUserAddressPPP func instead.
func NewEndUserAddress(addr string) *IE {
	if addr == "ppp" {
		return NewEndUserAddressPPP()
	}
	ip := net.ParseIP(addr)
	v4 := ip.To4()

	// IPv4
	if v4 != nil {
		return newEUAddrV4(v4)
	}

	return newEUAddrV6(ip)
}

// NewEndUserAddressIPv4 creates a new EndUserAddress IE with IPv4.
func NewEndUserAddressIPv4(addr string) *IE {
	v4 := net.ParseIP(addr).To4()
	if v4 == nil {
		return New(EndUserAddress, []byte{0xf1, 0x21})
	}

	return newEUAddrV4(v4)
}

// NewEndUserAddressIPv6 creates a new EndUserAddress IE with IPv6.
func NewEndUserAddressIPv6(addr string) *IE {
	v6 := net.ParseIP(addr).To16()
	if v6 == nil {
		return New(EndUserAddress, []byte{0xf1, 0x57})
	}

	return newEUAddrV6(v6)
}

func newEUAddrV4(v4 []byte) *IE {
	e := New(
		EndUserAddress,
		make([]byte, 6),
	)
	e.Payload[0] = pdpTypeIETF
	e.Payload[1] = 0x21
	copy(e.Payload[2:], v4)

	return e
}

func newEUAddrV6(v6 []byte) *IE {
	e := New(
		EndUserAddress,
		make([]byte, 18),
	)
	e.Payload = make([]byte, 18)
	e.Payload[0] = pdpTypeIETF
	e.Payload[1] = 0x57
	copy(e.Payload[2:], v6)

	return e
}

// NewEndUserAddressPPP creates a new EndUserAddress IE with PPP.
func NewEndUserAddressPPP() *IE {
	e := New(EndUserAddress, make([]byte, 2))
	e.Payload[0] = pdpTypeETSI
	e.Payload[1] = pdpTypeIETF

	e.SetLength()
	return e
}

// EndUserAddress returns EndUserAddress value if type matches.
func (i *IE) EndUserAddress() []byte {
	if i.Type != EndUserAddress {
		return nil
	}
	return i.Payload
}

// PDPTypeOrganization returns PDPTypeOrganization if type matches.
func (i *IE) PDPTypeOrganization() uint8 {
	if i.Type != EndUserAddress {
		return 0
	}
	return i.Payload[0]
}

// PDPTypeNumber returns PDPTypeNumber if type matches.
func (i *IE) PDPTypeNumber() uint8 {
	if i.Type != EndUserAddress {
		return 0
	}
	return i.Payload[1]
}

// IPAddress returns IPAddress if type matches.
func (i *IE) IPAddress() string {
	switch i.Type {
	case EndUserAddress:
		if i.PDPTypeOrganization() != pdpTypeIETF {
			return ""
		}
		if len(i.Payload) < 3 {
			return ""
		}
		return net.IP(i.Payload[2:]).String()
	case GSNAddress:
		return net.IP(i.Payload).String()
	default:
		return ""
	}
}
