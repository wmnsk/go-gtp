// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"

	"github.com/wmnsk/go-gtp/utils"
)

// NewUserCSGInformation creates a new UserCSGInformation IE.
func NewUserCSGInformation(mcc, mnc string, csgID uint32, mode, lcsg, cmi uint8) *IE {
	v := NewUserCSGInformationFields(mcc, mnc, csgID, mode, lcsg, cmi)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(UserCSGInformation, 0x00, b)
}

// UserCSGInformation returns UserCSGInformation in UserCSGInformationFields type if the type of IE matches.
func (i *IE) UserCSGInformation() (*UserCSGInformationFields, error) {
	switch i.Type {
	case UserCSGInformation:
		return ParseUserCSGInformationFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// UserCSGInformationFields is a set of fields in UserCSGInformation IE.
type UserCSGInformationFields struct {
	MCC, MNC   string
	CSGID      uint32 // 27-bit
	AccessMode uint8  // 7-8th bit, in the same octet as Flags
	Flags      uint8  // 1-2th bit, in the same octet as AccessMode
}

// NewUserCSGInformationFields creates a new UserCSGInformationFields.
func NewUserCSGInformationFields(mcc, mnc string, csgID uint32, mode, lcsg, cmi uint8) *UserCSGInformationFields {
	return &UserCSGInformationFields{
		MCC:        mcc,
		MNC:        mnc,
		CSGID:      csgID,
		AccessMode: mode,
		Flags:      ((lcsg << 1) & 0x02) | (cmi & 0x01),
	}
}

// Marshal serializes UserCSGInformationFields.
func (f *UserCSGInformationFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes UserCSGInformationFields.
func (f *UserCSGInformationFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 3 {
		return io.ErrUnexpectedEOF
	}

	plmn, err := utils.EncodePLMN(f.MCC, f.MNC)
	if err != nil {
		return err
	}
	copy(b[0:3], plmn)

	if l < 8 {
		return io.ErrUnexpectedEOF
	}

	binary.BigEndian.PutUint32(b[3:7], f.CSGID&0x7ffffff)
	b[7] = ((f.AccessMode & 0x03) << 6) | f.Flags

	return nil
}

// ParseUserCSGInformationFields decodes UserCSGInformationFields.
func ParseUserCSGInformationFields(b []byte) (*UserCSGInformationFields, error) {
	f := &UserCSGInformationFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into UserCSGInformationFields.
func (f *UserCSGInformationFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 3 {
		return io.ErrUnexpectedEOF
	}

	var err error
	f.MCC, f.MNC, err = utils.DecodePLMN(b[0:3])
	if err != nil {
		return err
	}

	if l <= 7 {
		return io.ErrUnexpectedEOF
	}

	f.CSGID = binary.BigEndian.Uint32(b[3:7]) & 0x7ffffff
	f.AccessMode = b[7] >> 6
	f.Flags = b[7] & 0x03

	return nil
}

// MarshalLen returns the serial length of UserCSGInformationFields in int.
func (f *UserCSGInformationFields) MarshalLen() int {
	return 8
}

// AccessMode returns AccessMode in uint8 if the type of IE matches.
func (i *IE) AccessMode() (uint8, error) {
	switch i.Type {
	case UserCSGInformation:
		if len(i.Payload) < 8 {
			return 0, io.ErrUnexpectedEOF
		}
		return i.Payload[7] >> 6, nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustAccessMode returns AccessMode in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustAccessMode() uint8 {
	v, _ := i.AccessMode()
	return v
}
