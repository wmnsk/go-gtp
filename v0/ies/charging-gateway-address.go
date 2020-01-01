// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"io"
	"net"
)

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
func (i *IE) ChargingGatewayAddress() (string, error) {
	if i.Type != ChargingGatewayAddress {
		return "", &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 4 {
		return "", io.ErrUnexpectedEOF
	}

	return net.IP(i.Payload).String(), nil
}

// MustChargingGatewayAddress returns ChargingGatewayAddress in string if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustChargingGatewayAddress() string {
	v, _ := i.ChargingGatewayAddress()
	return v
}
