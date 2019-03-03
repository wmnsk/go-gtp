// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"

	"github.com/wmnsk/go-gtp/utils"
)

// UserLocationInformation GeographicLocationType definitions.
const (
	locTypeCGI uint8 = iota
	locTypeSAI
	locTypeRAI
)

// NewUserLocationInformationWithCGI creates a new UserLocationInformation IE with LAC.
func NewUserLocationInformationWithCGI(mcc, mnc string, lac, cgi uint16) *IE {
	mc, err := utils.StrToSwappedBytes(mcc, "f")
	if err != nil {
		return nil
	}
	mn, err := utils.StrToSwappedBytes(mnc, "f")
	if err != nil {
		return nil
	}

	uli := New(
		UserLocationInformation,
		make([]byte, 8),
	)
	uli.Payload[0] = locTypeCGI
	copy(uli.Payload[1:3], mc)
	uli.Payload[3] = mn[0]
	binary.BigEndian.PutUint16(uli.Payload[4:6], lac)
	binary.BigEndian.PutUint16(uli.Payload[6:8], cgi)

	return uli
}

// NewUserLocationInformationWithSAI creates a new UserLocationInformation IE with LAC.
func NewUserLocationInformationWithSAI(mcc, mnc string, lac, sac uint16) *IE {
	mc, err := utils.StrToSwappedBytes(mcc, "f")
	if err != nil {
		return nil
	}
	mn, err := utils.StrToSwappedBytes(mnc, "f")
	if err != nil {
		return nil
	}

	uli := New(
		UserLocationInformation,
		make([]byte, 8),
	)
	uli.Payload[0] = locTypeSAI
	copy(uli.Payload[1:3], mc)
	uli.Payload[3] = mn[0]
	binary.BigEndian.PutUint16(uli.Payload[4:6], lac)
	binary.BigEndian.PutUint16(uli.Payload[6:8], sac)

	return uli
}

// NewUserLocationInformationWithRAI creates a new UserLocationInformation IE with LAC.
func NewUserLocationInformationWithRAI(mcc, mnc string, lac uint16, rac uint8) *IE {
	mc, err := utils.StrToSwappedBytes(mcc, "f")
	if err != nil {
		return nil
	}
	mn, err := utils.StrToSwappedBytes(mnc, "f")
	if err != nil {
		return nil
	}

	uli := New(
		UserLocationInformation,
		make([]byte, 7),
	)
	uli.Payload[0] = locTypeRAI
	copy(uli.Payload[1:3], mc)
	uli.Payload[3] = mn[0]
	binary.BigEndian.PutUint16(uli.Payload[4:6], lac)
	uli.Payload[6] = rac

	return uli
}

// UserLocationInformation returns UserLocationInformation value if type matches.
func (i *IE) UserLocationInformation() []byte {
	if i.Type != UserLocationInformation {
		return nil
	}
	return i.Payload
}

// MCC returns MCC value if type matches.
func (i *IE) MCC() string {
	switch i.Type {
	case RouteingAreaIdentity:
		return utils.SwappedBytesToStr(i.Payload[0:2], false)
	case UserLocationInformation:
		return utils.SwappedBytesToStr(i.Payload[1:3], false)
	default:
		return ""
	}
}

// MNC returns MNC value if type matches.
func (i *IE) MNC() string {
	switch i.Type {
	case RouteingAreaIdentity:
		return utils.SwappedBytesToStr(i.Payload[1:2], true)
	case UserLocationInformation:
		return utils.SwappedBytesToStr(i.Payload[2:3], true)
	default:
		return ""
	}
}

// LAC returns LAC value if type matches.
func (i *IE) LAC() uint16 {
	switch i.Type {
	case RouteingAreaIdentity:
		return binary.BigEndian.Uint16(i.Payload[3:5])
	case UserLocationInformation:
		return binary.BigEndian.Uint16(i.Payload[4:6])
	default:
		return 0
	}
}

// CGI returns CGI value if type matches.
func (i *IE) CGI() uint16 {
	if i.Type == UserLocationInformation && i.Payload[0] == locTypeCGI {
		return binary.BigEndian.Uint16(i.Payload[6:7])
	}
	return 0
}

// SAC returns SAC value if type matches.
func (i *IE) SAC() uint16 {
	if i.Type == UserLocationInformation && i.Payload[0] == locTypeSAI {
		return binary.BigEndian.Uint16(i.Payload[6:8])
	}
	return 0
}

// RAC returns RAC value if type matches.
func (i *IE) RAC() uint8 {
	switch i.Type {
	case RouteingAreaIdentity:
		return i.Payload[5]
	case UserLocationInformation:
		if i.Payload[0] == locTypeRAI {
			return i.Payload[6]
		}
	}
	return 0
}
