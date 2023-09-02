// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
	"net"
)

// NewIPAddress creates a new IPAddress IE from string.
func NewIPAddress(addr string) *IE {
	ip := net.ParseIP(addr)
	v4 := ip.To4()

	// IPv4
	if v4 != nil {
		return New(IPAddress, 0x00, v4)
	}
	// IPv6
	return New(IPAddress, 0x00, ip)
}

// NewIPAddressNetIP creates a new IPAddress IE from net.IP.
func NewIPAddressNetIP(ip net.IP) *IE {
	if v := ip.To4(); v != nil {
		return New(IPAddress, 0x00, v)
	}

	if v := ip.To16(); v != nil {
		return New(IPAddress, 0x00, v)
	}

	// return IE w/ "something" anyway, to avoid crash
	return New(IPAddress, 0x00, ip)
}

// IPAddress returns IPAddress value if the type of IE matches.
func (i *IE) IPAddress() (string, error) {
	ip, err := i.IP()
	if err != nil {
		return "", err
	}

	return ip.String(), nil
}

// MustIPAddress returns IPAddress in string, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustIPAddress() string {
	v, _ := i.IPAddress()
	return v
}

// IP returns IP value in net.IP if the type of IE matches.
//
// This does not validate the values. Even if the IP address fields in an
// IE contains some invalid byte sequence, it just try to return as it is.
// To retrieve only if it's valid, use IPv4 or IPv6 instead.
//
// If the IE has both IPv4 and IPv6, this returns IPv4. To retrieve IPv6
// value in that case, use IPv6() instead.
func (i *IE) IP() (net.IP, error) {
	if len(i.Payload) < 4 {
		return nil, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case IPAddress:
		return net.IP(i.Payload), nil
	case PDNAddressAllocation:
		if len(i.Payload) < 5 {
			return nil, io.ErrUnexpectedEOF
		}
		paa, err := ParsePDNAddressAllocationFields(i.Payload)
		if err != nil {
			return nil, err
		}
		switch paa.PDNType {
		case pdnTypeIPv4, pdnTypeIPv4v6:
			return paa.IPv4Address, nil
		case pdnTypeIPv6:
			return paa.IPv6Address, nil
		default:
			return nil, ErrIEValueNotFound
		}
	case S103PDNDataForwardingInfo:
		switch i.Payload[0] {
		case 4:
			if len(i.Payload) < 5 {
				return nil, io.ErrUnexpectedEOF
			}
			return net.IP(i.Payload[1:5]), nil
		case 16:
			if len(i.Payload) < 17 {
				return nil, io.ErrUnexpectedEOF
			}
			return net.IP(i.Payload[1:17]), nil
		default:
			return nil, ErrMalformed
		}
	case S1UDataForwarding:
		switch i.Payload[1] {
		case 4:
			if len(i.Payload) < 6 {
				return nil, io.ErrUnexpectedEOF
			}
			return net.IP(i.Payload[2:6]), nil
		case 16:
			if len(i.Payload) < 18 {
				return nil, io.ErrUnexpectedEOF
			}
			return net.IP(i.Payload[2:18]), nil
		default:
			return nil, ErrMalformed
		}
	case FullyQualifiedTEID:
		fteid, err := ParseFullyQualifiedTEIDFields(i.Payload)
		if err != nil {
			return nil, err
		}
		if i.HasIPv4() {
			return fteid.IPv4Address, nil
		} else if i.HasIPv6() {
			return fteid.IPv6Address, nil
		}
		return nil, ErrIEValueNotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MustIP returns IP in net.IP, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustIP() net.IP {
	v, _ := i.IP()
	return v
}

// IPv4 returns IPv4 value in net.IP if the type of IE matches.
func (i *IE) IPv4() (net.IP, error) {
	ip, err := i.IP()
	if err != nil {
		return nil, err
	}

	if v := ip.To4(); v != nil {
		return v, nil
	}
	return nil, ErrIEValueNotFound
}

// MustIPv4 returns IPv4 in net.IP, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustIPv4() net.IP {
	v, _ := i.IPv4()
	return v
}

// IPv6 returns IPv6 value in net.IP if the type of IE matches.
func (i *IE) IPv6() (net.IP, error) {
	if len(i.Payload) < 4 {
		return nil, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case IPAddress, S103PDNDataForwardingInfo, S1UDataForwarding:
		ip, err := i.IP()
		if err != nil {
			return nil, err
		}

		if v := ip.To16(); v != nil {
			return v, nil
		}
		return nil, ErrIEValueNotFound
	case PDNAddressAllocation:
		paa, err := ParsePDNAddressAllocationFields(i.Payload)
		if err != nil {
			return nil, err
		}
		if v := paa.IPv6Address.To16(); v != nil {
			return v, nil
		}
		return nil, ErrIEValueNotFound
	case FullyQualifiedTEID:
		fteid, err := ParseFullyQualifiedTEIDFields(i.Payload)
		if err != nil {
			return nil, err
		}

		if i.HasIPv6() {
			if v := fteid.IPv6Address.To16(); v != nil {
				return v, nil
			}
			return nil, ErrIEValueNotFound
		}
		return nil, ErrIEValueNotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MustIPv6 returns IPv6 in net.IP, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustIPv6() net.IP {
	v, _ := i.IPv6()
	return v
}
