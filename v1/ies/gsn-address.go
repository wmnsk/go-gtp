// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "net"

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
