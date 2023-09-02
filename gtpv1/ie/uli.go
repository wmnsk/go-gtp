// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"

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
	plmn, err := utils.EncodePLMN(mcc, mnc)
	if err != nil {
		return nil
	}

	uli := New(
		UserLocationInformation,
		make([]byte, 8),
	)
	uli.Payload[0] = locTypeCGI
	copy(uli.Payload[1:4], plmn)
	binary.BigEndian.PutUint16(uli.Payload[4:6], lac)
	binary.BigEndian.PutUint16(uli.Payload[6:8], cgi)

	return uli
}

// NewUserLocationInformationWithSAI creates a new UserLocationInformation IE with LAC.
func NewUserLocationInformationWithSAI(mcc, mnc string, lac, sac uint16) *IE {
	plmn, err := utils.EncodePLMN(mcc, mnc)
	if err != nil {
		return nil
	}

	uli := New(
		UserLocationInformation,
		make([]byte, 8),
	)
	uli.Payload[0] = locTypeSAI
	copy(uli.Payload[1:4], plmn)
	binary.BigEndian.PutUint16(uli.Payload[4:6], lac)
	binary.BigEndian.PutUint16(uli.Payload[6:8], sac)

	return uli
}

// NewUserLocationInformationWithRAI creates a new UserLocationInformation IE with LAC.
func NewUserLocationInformationWithRAI(mcc, mnc string, lac uint16, rac uint8) *IE {
	plmn, err := utils.EncodePLMN(mcc, mnc)
	if err != nil {
		return nil
	}

	uli := New(
		UserLocationInformation,
		make([]byte, 7),
	)
	uli.Payload[0] = locTypeRAI
	copy(uli.Payload[1:4], plmn)
	binary.BigEndian.PutUint16(uli.Payload[4:6], lac)
	uli.Payload[6] = rac

	return uli
}

// UserLocationInformation returns UserLocationInformation value if type matches.
func (i *IE) UserLocationInformation() ([]byte, error) {
	if i.Type != UserLocationInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	return i.Payload, nil
}

// MustUserLocationInformation returns UserLocationInformation in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustUserLocationInformation() []byte {
	v, _ := i.UserLocationInformation()
	return v
}

// CGI returns CGI value if type matches.
func (i *IE) CGI() (uint16, error) {
	if i.Type != UserLocationInformation {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Payload[0] {
	case locTypeCGI:
		if len(i.Payload) < 8 {
			return 0, io.ErrUnexpectedEOF
		}
		return binary.BigEndian.Uint16(i.Payload[6:8]), nil
	}
	return 0, &InvalidTypeError{Type: i.Type}
}

// MustCGI returns CGI in uint16 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustCGI() uint16 {
	v, _ := i.CGI()
	return v
}

// SAC returns SAC value if type matches.
func (i *IE) SAC() (uint16, error) {
	if i.Type != UserLocationInformation {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Payload[0] {
	case locTypeSAI:
		if len(i.Payload) < 8 {
			return 0, io.ErrUnexpectedEOF
		}
		return binary.BigEndian.Uint16(i.Payload[6:8]), nil
	}
	return 0, &InvalidTypeError{Type: i.Type}
}

// MustSAC returns SAC in uint16 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustSAC() uint16 {
	v, _ := i.SAC()
	return v
}
