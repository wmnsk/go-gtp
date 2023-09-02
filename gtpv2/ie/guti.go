// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"

	"github.com/wmnsk/go-gtp/utils"
)

// NewGUTI creates a new GUTI IE.
func NewGUTI(mcc, mnc string, groupID uint16, code uint8, mTMSI uint32) *IE {
	v := NewGUTIFields(mcc, mnc, groupID, code, mTMSI)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(GUTI, 0x00, b)
}

// GUTI returns GUTI in GUTIFields type if the type of IE matches.
func (i *IE) GUTI() (*GUTIFields, error) {
	switch i.Type {
	case GUTI:
		return ParseGUTIFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// GUTIFields is a set of fields in GUTI IE.
type GUTIFields struct {
	MCC, MNC   string
	MMEGroupID uint16
	MMECode    uint8
	MTMSI      uint32
}

// NewGUTIFields creates a new GUTIFields.
func NewGUTIFields(mcc, mnc string, groupID uint16, code uint8, mTMSI uint32) *GUTIFields {
	return &GUTIFields{
		MCC:        mcc,
		MNC:        mnc,
		MMEGroupID: groupID,
		MMECode:    code,
		MTMSI:      mTMSI,
	}
}

// Marshal serializes GUTIFields.
func (f *GUTIFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes GUTIFields.
func (f *GUTIFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 3 {
		return io.ErrUnexpectedEOF
	}

	plmn, err := utils.EncodePLMN(f.MCC, f.MNC)
	if err != nil {
		return nil
	}
	copy(b[0:3], plmn)

	if l < 10 {
		return io.ErrUnexpectedEOF
	}
	binary.BigEndian.PutUint16(b[3:5], f.MMEGroupID)
	b[5] = f.MMECode
	binary.BigEndian.PutUint32(b[6:10], f.MTMSI)

	return nil
}

// ParseGUTIFields decodes GUTIFields.
func ParseGUTIFields(b []byte) (*GUTIFields, error) {
	f := &GUTIFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into GUTIFields.
func (f *GUTIFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 3 {
		return io.ErrUnexpectedEOF
	}

	var err error
	f.MCC, f.MNC, err = utils.DecodePLMN(b[1:3])
	if err != nil {
		return err
	}

	if l < 10 {
		return io.ErrUnexpectedEOF
	}

	f.MMEGroupID = binary.BigEndian.Uint16(b[3:5])
	f.MMECode = b[5]
	f.MTMSI = binary.BigEndian.Uint32(b[6:10])

	return nil
}

// MarshalLen returns the serial length of GUTIFields in int.
func (f *GUTIFields) MarshalLen() int {
	return 10
}

// MMEGroupID returns MMEGroupID in uint16 if the type of IE matches.
func (i *IE) MMEGroupID() (uint16, error) {
	switch i.Type {
	case GUTI:
		if len(i.Payload) < 5 {
			return 0, io.ErrUnexpectedEOF
		}
		return binary.BigEndian.Uint16(i.Payload[3:5]), nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustMMEGroupID returns MMEGroupID in uint16, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustMMEGroupID() uint16 {
	v, _ := i.MMEGroupID()
	return v
}

// MMECode returns MMECode in uint8 if the type of IE matches.
func (i *IE) MMECode() (uint8, error) {
	switch i.Type {
	case GUTI:
		if len(i.Payload) < 6 {
			return 0, io.ErrUnexpectedEOF
		}
		return i.Payload[5], nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustMMECode returns MMECode in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustMMECode() uint8 {
	v, _ := i.MMECode()
	return v
}

// MTMSI returns MTMSI in uint32 if the type of IE matches.
func (i *IE) MTMSI() (uint32, error) {
	switch i.Type {
	case GUTI:
		if len(i.Payload) < 10 {
			return 0, io.ErrUnexpectedEOF
		}
		return binary.BigEndian.Uint32(i.Payload[6:10]), nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustMTMSI returns MTMSI in uint32, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustMTMSI() uint32 {
	v, _ := i.MTMSI()
	return v
}
