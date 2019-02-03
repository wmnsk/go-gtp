// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"net"
	"strings"
)

// NewRecovery creates a new Recovery IE.
func NewRecovery(recovery uint8) *IE {
	return newUint8ValIE(Recovery, recovery)
}

// Recovery returns Recovery value if type matches.
func (i *IE) Recovery() uint8 {
	if i.Type != Recovery {
		return 0
	}
	return i.Payload[0]
}

// NewAccessPointName creates a new AccessPointName IE.
func NewAccessPointName(apn string) *IE {
	i := New(AccessPointName, make([]byte, len(apn)+1))
	var offset = 0
	for _, label := range strings.Split(apn, ".") {
		l := len(label)
		i.Payload[offset] = uint8(l)
		copy(i.Payload[offset+1:], []byte(label))
		offset += l + 1
	}

	return i
}

// AccessPointName returns AccessPointName in string if type of IE matches.
func (i *IE) AccessPointName() string {
	if i.Type != AccessPointName {
		return ""
	}

	var (
		apn    []string
		offset int
	)

	max := len(i.Payload)
	for {
		if offset >= max {
			break
		}
		l := int(i.Payload[offset])
		apn = append(apn, string(i.Payload[offset+1:offset+l+1]))
		offset += l + 1
	}

	return strings.Join(apn, ".")
}

// NewGSNAddress creates a new GSNAddress IE.
func NewGSNAddress(addr string) *IE {
	ip := net.ParseIP(addr)
	v4 := ip.To4()

	// IPv4
	if v4 != nil {
		return New(GSNAddress, v4)
	}
	//IPv6
	return New(GSNAddress, ip)
}

// GSNAddress returns GSNAddress value if type matches.
func (i *IE) GSNAddress() string {
	if i.Type != GSNAddress {
		return ""
	}
	return net.IP(i.Payload).String()
}

// NewRATType creates a new RATType IE.
func NewRATType(ratType uint8) *IE {
	return New(
		RATType,
		[]byte{ratType},
	)
}

// RATType returns RATType value if type matches.
func (i *IE) RATType() uint8 {
	if i.Type != RATType {
		return 0
	}
	return i.Payload[0]
}
