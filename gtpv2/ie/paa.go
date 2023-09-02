// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
	"net"
)

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
//
// NOTE: IPv6 Prefix Length will be set to 0 when an IPv6 address is given to this.
// Use NewPDNAddressAllocationIPv6 instead to set the correct value.
func NewPDNAddressAllocation(addr string) *IE {
	return NewPDNAddressAllocationNetIP(net.ParseIP(addr), 0)
}

// NewPDNAddressAllocation creates a new PDNAddressAllocation IE with IPv6 value.
func NewPDNAddressAllocationIPv6(addr string, prefix uint8) *IE {
	return NewPDNAddressAllocationNetIP(net.ParseIP(addr), prefix)
}

// NewPDNAddressAllocationDual creates a new PDNAddressAllocation IE with
// IPv4 address and IPv6 address given.
//
// If they cannot be converted as IPv4/IPv6, PDN Type will be Non-IP.
func NewPDNAddressAllocationDual(v4addr, v6addr string, v6prefix uint8) *IE {
	return NewPDNAddressAllocationDualNetIP(net.ParseIP(v4addr), net.ParseIP(v6addr), v6prefix)
}

// NewPDNAddressAllocationNetIP creates a new PDNAddressAllocation IE from net.IP.
func NewPDNAddressAllocationNetIP(ip net.IP, v6prefix uint8) *IE {
	var v *PDNAddressAllocationFields
	if v4 := ip.To4(); v4 != nil {
		v = NewPDNAddressAllocationFields(pdnTypeIPv4, v4, nil, 0)
	} else if v6 := ip.To16(); v6 != nil {
		v = NewPDNAddressAllocationFields(pdnTypeIPv6, nil, v6, v6prefix)
	} else {
		v = NewPDNAddressAllocationFields(pdnTypeNonIP, nil, nil, 0)
	}

	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(PDNAddressAllocation, 0x00, b)
}

// NewPDNAddressAllocationDualNetIP creates a new PDNAddressAllocation IE from
// IPv4 and IPv6 in net.IP.
func NewPDNAddressAllocationDualNetIP(v4, v6 net.IP, v6prefix uint8) *IE {
	v := NewPDNAddressAllocationFields(pdnTypeIPv4v6, v4, v6, v6prefix)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(PDNAddressAllocation, 0x00, b)
}

// PDNAddressAllocationFields is a set of fields in PDNAddressAllocation IE.
type PDNAddressAllocationFields struct {
	PDNType          uint8 // 3-bit
	IPv4Address      net.IP
	IPv6PrefixLength uint8
	IPv6Address      net.IP
}

// NewPDNAddressAllocationFields creates a new PDNAddressAllocationFields.
func NewPDNAddressAllocationFields(pType uint8, v4, v6 net.IP, prefix uint8) *PDNAddressAllocationFields {
	return &PDNAddressAllocationFields{
		PDNType:          pType,
		IPv4Address:      v4,
		IPv6PrefixLength: prefix,
		IPv6Address:      v6,
	}
}

// Marshal serializes PDNAddressAllocationFields.
func (f *PDNAddressAllocationFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes PDNAddressAllocationFields.
func (f *PDNAddressAllocationFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	b[0] = f.PDNType & 0x07
	offset := 1

	switch f.PDNType {
	case pdnTypeIPv4:
		if l < offset+4 {
			return io.ErrUnexpectedEOF
		}

		if f.IPv4Address != nil {
			copy(b[offset:offset+4], f.IPv4Address.To4())
		}

		return nil
	case pdnTypeIPv6:
		if l < offset+17 {
			return io.ErrUnexpectedEOF
		}

		b[offset] = f.IPv6PrefixLength
		copy(b[offset+1:offset+17], f.IPv6Address.To16())

		return nil
	case pdnTypeIPv4v6:
		if l < offset+21 {
			return io.ErrUnexpectedEOF
		}

		b[offset] = f.IPv6PrefixLength
		copy(b[offset+1:offset+17], f.IPv6Address.To16())
		copy(b[offset+17:offset+21], f.IPv4Address.To4())

		return nil
	case pdnTypeNonIP: // no payload
		return nil
	default:
		return nil
	}

}

// ParsePDNAddressAllocationFields decodes PDNAddressAllocationFields.
func ParsePDNAddressAllocationFields(b []byte) (*PDNAddressAllocationFields, error) {
	f := &PDNAddressAllocationFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into PDNAddressAllocationFields.
func (f *PDNAddressAllocationFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	f.PDNType = b[0] & 0x07
	offset := 1

	switch f.PDNType {
	case pdnTypeIPv4:
		if l < offset+4 {
			return io.ErrUnexpectedEOF
		}

		f.IPv4Address = b[offset : offset+4]

		return nil
	case pdnTypeIPv6:
		if l < offset+17 {
			return io.ErrUnexpectedEOF
		}

		f.IPv6PrefixLength = b[offset]
		f.IPv6Address = b[offset+1 : offset+17]

		return nil
	case pdnTypeIPv4v6:
		if l < offset+21 {
			return io.ErrUnexpectedEOF
		}

		f.IPv6PrefixLength = b[offset]
		f.IPv6Address = b[offset+1 : offset+17]
		f.IPv4Address = b[offset+17 : offset+21]

		return nil
	case pdnTypeNonIP: // no payload
		return nil
	default:
		return nil
	}
}

// MarshalLen returns the serial length of PDNAddressAllocationFields in int.
func (f *PDNAddressAllocationFields) MarshalLen() int {
	l := 1
	switch f.PDNType {
	case pdnTypeIPv4:
		l += 4
	case pdnTypeIPv6:
		l += 17
	case pdnTypeIPv4v6:
		l += 21
	case pdnTypeNonIP:
		// no payload, do nothing
	}

	return l
}
