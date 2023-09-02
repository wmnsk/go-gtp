// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"

	"github.com/wmnsk/go-gtp/utils"
)

// NewRouteingAreaIdentity creates a new RouteingAreaIdentity IE.
func NewRouteingAreaIdentity(mcc, mnc string, lac uint16, rac uint8) *IE {
	mc, err := utils.StrToSwappedBytes(mcc, "f")
	if err != nil {
		return nil
	}
	mn, err := utils.StrToSwappedBytes(mnc, "f")
	if err != nil {
		return nil
	}

	rai := New(
		RouteingAreaIdentity,
		make([]byte, 6),
	)
	copy(rai.Payload[0:2], mc)
	rai.Payload[2] = mn[0]
	binary.BigEndian.PutUint16(rai.Payload[3:5], lac)
	rai.Payload[5] = rac

	return rai
}

// RouteingAreaIdentity returns RouteingAreaIdentity value if type matches.
func (i *IE) RouteingAreaIdentity() ([]byte, error) {
	if i.Type != RouteingAreaIdentity {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	return i.Payload, nil
}

// MustRouteingAreaIdentity returns RouteingAreaIdentity in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustRouteingAreaIdentity() []byte {
	v, _ := i.RouteingAreaIdentity()
	return v
}
