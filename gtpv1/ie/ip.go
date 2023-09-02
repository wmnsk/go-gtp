// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
	"net"
)

// IP returns IP in net.IP if type matches.
func (i *IE) IP() (net.IP, error) {
	if len(i.Payload) < 4 {
		return nil, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case EndUserAddress:
		if i.MustPDPTypeOrganization() != pdpTypeIETF {
			return nil, ErrMalformed
		}
		return net.IP(i.Payload[2:]), nil
	case GSNAddress:
		return net.IP(i.Payload), nil
	default:
		return nil, &InvalidTypeError{i.Type}
	}
}

func (i *IE) MustIP() net.IP {
	v, _ := i.IP()
	return v
}

// IPAddress returns IPAddress in string if type matches.
func (i *IE) IPAddress() (string, error) {
	ip, err := i.IP()
	if err != nil {
		return "", err
	}

	return ip.String(), nil
}

// MustIPAddress returns IPAddress in string if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustIPAddress() string {
	v, _ := i.IPAddress()
	return v
}
