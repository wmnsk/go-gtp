// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
	"net"
)

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

	return NewEndUserAddressByIP(net.ParseIP(addr))
}

// NewEndUserAddressByIP creates a new EndUserAddress IE from net.IP.
func NewEndUserAddressByIP(ip net.IP) *IE {
	if ip == nil {
		return nil
	}

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
func (i *IE) EndUserAddress() ([]byte, error) {
	if i.Type != EndUserAddress {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	return i.Payload, nil
}

// MustEndUserAddress returns EndUserAddress in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustEndUserAddress() []byte {
	v, _ := i.EndUserAddress()
	return v
}

// PDPTypeOrganization returns PDPTypeOrganization if type matches.
func (i *IE) PDPTypeOrganization() (uint8, error) {
	if i.Type != EndUserAddress {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustPDPTypeOrganization returns PDPTypeOrganization in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustPDPTypeOrganization() uint8 {
	v, _ := i.PDPTypeOrganization()
	return v
}

// PDPTypeNumber returns PDPTypeNumber if type matches.
func (i *IE) PDPTypeNumber() (uint8, error) {
	if i.Type != EndUserAddress {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[1], nil
}

// MustPDPTypeNumber returns PDPTypeNumber in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustPDPTypeNumber() uint8 {
	v, _ := i.PDPTypeNumber()
	return v
}
