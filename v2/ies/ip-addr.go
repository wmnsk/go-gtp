// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

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
	//IPv6
	return New(IPAddress, 0x00, ip)
}

// IPAddress returns IPAddress value if the type of IE matches.
func (i *IE) IPAddress() (string, error) {
	if len(i.Payload) == 0 {
		return "", io.ErrUnexpectedEOF
	}

	switch i.Type {
	case IPAddress:
		return net.IP(i.Payload).String(), nil
	case PDNAddressAllocation:
		if len(i.Payload) < 5 {
			return "", io.ErrUnexpectedEOF
		}
		pdnType, err := i.PDNType()
		if err != nil {
			return "", err
		}
		switch pdnType {
		case 0x01:
			return net.IP(i.Payload[1:]).String(), nil
		case 0x02:
			return net.IP(i.Payload[2:]).String(), nil
		default:
			return "", ErrMalformed
		}
	case S103PDNDataForwardingInfo, S1UDataForwarding:
		switch i.Payload[0] {
		case 4:
			if len(i.Payload) < 5 {
				return "", io.ErrUnexpectedEOF
			}
			return net.IP(i.Payload[1:5]).String(), nil
		case 16:
			if len(i.Payload) < 17 {
				return "", io.ErrUnexpectedEOF
			}
			return net.IP(i.Payload[1:17]).String(), nil
		default:
			return "", ErrMalformed
		}
	case FullyQualifiedTEID:
		if i.HasIPv4() {
			if len(i.Payload) < 9 {
				return "", io.ErrUnexpectedEOF
			}
			return net.IP(i.Payload[5:9]).String(), nil
		} else if i.HasIPv6() {
			if len(i.Payload) < 21 {
				return "", io.ErrUnexpectedEOF
			}
			return net.IP(i.Payload[5:21]).String(), nil
		} else {
			return "", ErrMalformed
		}
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}

// MustIPAddress returns IPAddress in string, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustIPAddress() string {
	v, _ := i.IPAddress()
	return v
}
