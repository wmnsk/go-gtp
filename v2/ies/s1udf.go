// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"
	"net"
)

// NewS1UDataForwarding creates a new S1UDataForwarding IE.
func NewS1UDataForwarding(sgwAddr string, sgwTEID uint32) *IE {
	addr := net.ParseIP(sgwAddr)
	if addr == nil {
		return nil
	}

	// SGW Address: IPv4
	if v4 := addr.To4(); v4 != nil {
		i := New(S1UDataForwarding, 0x00, make([]byte, 1+4+4))
		i.Payload[0] = 4
		copy(i.Payload[1:5], v4)
		binary.BigEndian.PutUint32(i.Payload[5:9], sgwTEID)
		return i
	}

	// SGW Address: IPv6
	i := New(S1UDataForwarding, 0x00, make([]byte, 1+16+4))
	i.Payload[0] = 16
	copy(i.Payload[1:17], addr)
	binary.BigEndian.PutUint32(i.Payload[17:21], sgwTEID)
	return i
}

// SGWAddress returns IP address of SGW in string if the type of IE matches.
func (i *IE) SGWAddress() (string, error) {
	if i.Type != S1UDataForwarding {
		return "", &InvalidTypeError{Type: i.Type}
	}

	return i.IPAddress()
}

// MustSGWAddress returns SGWAddress in string, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustSGWAddress() string {
	v, _ := i.SGWAddress()
	return v
}
