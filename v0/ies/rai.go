// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"

	"github.com/wmnsk/go-gtp/utils"
)

// NewRouteingAreaIdentity creates a new RouteingAreaIdentity IE.
func NewRouteingAreaIdentity(mcc, mnc string, lac uint16, rac uint8) *IE {
	plmn, err := utils.EncodePLMN(mcc, mnc)
	if err != nil {
		return nil
	}

	rai := New(
		RouteingAreaIdentity,
		make([]byte, 6),
	)
	copy(rai.Payload[0:3], plmn)
	binary.BigEndian.PutUint16(rai.Payload[3:5], lac)
	rai.Payload[5] = rac

	return rai
}

// RouteingAreaIdentity returns RouteingAreaIdentity value if type matches.
func (i *IE) RouteingAreaIdentity() []byte {
	if i.Type != RouteingAreaIdentity {
		return nil
	}
	return i.Payload
}

// MCC returns MCC value if type matches.
func (i *IE) MCC() string {
	switch i.Type {
	case RouteingAreaIdentity:
		return utils.SwappedBytesToStr(i.Payload[0:2], false)
	default:
		return ""
	}
}

// MNC returns MNC value if type matches.
func (i *IE) MNC() string {
	switch i.Type {
	case RouteingAreaIdentity:
		return utils.SwappedBytesToStr(i.Payload[1:2], true)
	default:
		return ""
	}
}

// LAC returns LAC value if type matches.
func (i *IE) LAC() uint16 {
	switch i.Type {
	case RouteingAreaIdentity:
		return binary.BigEndian.Uint16(i.Payload[3:5])
	default:
		return 0
	}
}

// RAC returns RAC value if type matches.
func (i *IE) RAC() uint8 {
	switch i.Type {
	case RouteingAreaIdentity:
		return i.Payload[5]
	default:
		return 0
	}
}
