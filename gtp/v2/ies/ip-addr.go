// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
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
func (i *IE) IPAddress() string {
	switch i.Type {
	case IPAddress:
		return net.IP(i.Payload).String()
	case PDNAddressAllocation:
		switch i.PDNType() {
		case 0x01:
			return net.IP(i.Payload[1:]).String()
		case 0x02:
			return net.IP(i.Payload[2:]).String()
		default:
			return ""
		}
	case S103PDNDataForwardingInfo, S1UDataForwarding:
		switch i.Payload[0] {
		case 4:
			return net.IP(i.Payload[1:5]).String()
		case 16:
			return net.IP(i.Payload[1:17]).String()
		default:
			return ""
		}
	case FullyQualifiedTEID:
		if i.HasIPv4() {
			return net.IP(i.Payload[5:9]).String()
		} else if i.HasIPv6() {
			return net.IP(i.Payload[5:21]).String()
		} else {
			return ""
		}
	default:
		return ""
	}
}
