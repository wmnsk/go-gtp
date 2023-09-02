// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"

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

// MCC returns MCC value if type matches.
func (i *IE) MCC() (string, error) {
	switch i.Type {
	case RouteingAreaIdentity:
		if len(i.Payload) < 2 {
			return "", io.ErrUnexpectedEOF
		}
		return utils.SwappedBytesToStr(i.Payload[0:2], false), nil
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}

// MustMCC returns MCC in string if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustMCC() string {
	v, _ := i.MCC()
	return v
}

// MNC returns MNC value if type matches.
func (i *IE) MNC() (string, error) {
	switch i.Type {
	case RouteingAreaIdentity:
		if len(i.Payload) < 2 {
			return "", io.ErrUnexpectedEOF
		}
		return utils.SwappedBytesToStr(i.Payload[1:2], true), nil
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}

// MustMNC returns MNC in string if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustMNC() string {
	v, _ := i.MNC()
	return v
}

// LAC returns LAC value if type matches.
func (i *IE) LAC() (uint16, error) {
	switch i.Type {
	case RouteingAreaIdentity:
		if len(i.Payload) < 5 {
			return 0, io.ErrUnexpectedEOF
		}
		return binary.BigEndian.Uint16(i.Payload[3:5]), nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustLAC returns LAC in uint16 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustLAC() uint16 {
	v, _ := i.LAC()
	return v
}

// RAC returns RAC value if type matches.
func (i *IE) RAC() (uint8, error) {
	switch i.Type {
	case RouteingAreaIdentity:
		if len(i.Payload) < 6 {
			return 0, io.ErrUnexpectedEOF
		}
		return i.Payload[5], nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustRAC returns RAC in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustRAC() uint8 {
	v, _ := i.RAC()
	return v
}
