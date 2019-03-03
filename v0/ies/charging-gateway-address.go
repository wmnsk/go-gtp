// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "net"

// NewChargingGatewayAddress creates a new ChargingGatewayAddress IE from string.
func NewChargingGatewayAddress(addr string) *IE {
	ip := net.ParseIP(addr)
	v4 := ip.To4()

	// IPv4
	if v4 != nil {
		return New(ChargingGatewayAddress, v4)
	}
	//IPv6
	return New(ChargingGatewayAddress, ip)
}

// ChargingGatewayAddress returns ChargingGatewayAddress value if type matches.
func (i *IE) ChargingGatewayAddress() string {
	if i.Type != ChargingGatewayAddress {
		return ""
	}
	return net.IP(i.Payload).String()
}
