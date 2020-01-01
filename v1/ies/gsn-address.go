// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"io"
	"net"
)

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
func (i *IE) GSNAddress() (string, error) {
	if i.Type != GSNAddress {
		return "", &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 4 {
		return "", io.ErrUnexpectedEOF
	}

	return net.IP(i.Payload).String(), nil
}

// MustGSNAddress returns GSNAddress in string if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustGSNAddress() string {
	v, _ := i.GSNAddress()
	return v
}
