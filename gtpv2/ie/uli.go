// Copyright 2019-2021 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"

	"github.com/wmnsk/go-gtp/utils"
)

const (
	cgilen    int = 7
	sailen    int = 7
	railen    int = 7
	tailen    int = 5
	ecgilen   int = 7
	lailen    int = 5
	menbilen  int = 6
	emenbilen int = 6
)

// PLMN represents a PLMN-ID(MCC and MNC).
type PLMN struct {
	MCC string
	MNC string
}

// CGI represents a CGI, which is defined to be used as a field of UserLocationInformation IE.
type CGI struct {
	*PLMN
	LAC uint16
	CI  uint16
}

// NewCGI creates a new CGI.
func NewCGI(mcc, mnc string, lac, ci uint16) *CGI {
	return &CGI{
		PLMN: &PLMN{MCC: mcc, MNC: mnc},
		LAC:  lac,
		CI:   ci,
	}
}

// SAI represents a SAI, which is defined to be used as a field of UserLocationInformation IE.
type SAI struct {
	*PLMN
	LAC uint16
	SAC uint16
}

// NewSAI creates a new SAI.
func NewSAI(mcc, mnc string, lac, sac uint16) *SAI {
	return &SAI{
		PLMN: &PLMN{MCC: mcc, MNC: mnc},
		LAC:  lac,
		SAC:  sac,
	}
}

// RAI represents a RAI, which is defined to be used as a field of UserLocationInformation IE.
type RAI struct {
	*PLMN
	LAC uint16
	RAC uint16
}

// NewRAI creates a new RAI.
func NewRAI(mcc, mnc string, lac, rac uint16) *RAI {
	return &RAI{
		PLMN: &PLMN{MCC: mcc, MNC: mnc},
		LAC:  lac,
		RAC:  rac,
	}
}

// TAI represents a TAI, which is defined to be used as a field of UserLocationInformation IE.
type TAI struct {
	*PLMN
	TAC uint16
}

// NewTAI creates a new TAI.
func NewTAI(mcc, mnc string, tac uint16) *TAI {
	return &TAI{
		PLMN: &PLMN{MCC: mcc, MNC: mnc},
		TAC:  tac,
	}
}

// ECGI represents a ECGI, which is defined to be used as a field of UserLocationInformation IE.
type ECGI struct {
	*PLMN
	ECI uint32
}

// NewECGI creates a new ECGI.
func NewECGI(mcc, mnc string, eci uint32) *ECGI {
	return &ECGI{
		PLMN: &PLMN{MCC: mcc, MNC: mnc},
		ECI:  eci & 0xfffffff,
	}
}

// LAI represents a LAI, which is defined to be used as a field of UserLocationInformation IE.
type LAI struct {
	*PLMN
	LAC uint16
}

// NewLAI creates a new LAI.
func NewLAI(mcc, mnc string, lac uint16) *LAI {
	return &LAI{
		PLMN: &PLMN{MCC: mcc, MNC: mnc},
		LAC:  lac,
	}
}

// MENBI represents a MENBI, which is defined to be used as a field of UserLocationInformation IE.
type MENBI struct {
	*PLMN
	MENBI uint32
}

// NewMENBI creates a new MENBI.
func NewMENBI(mcc, mnc string, menbi uint32) *MENBI {
	return &MENBI{
		PLMN:  &PLMN{MCC: mcc, MNC: mnc},
		MENBI: menbi,
	}
}

// EMENBI represents a EMENBI, which is defined to be used as a field of UserLocationInformation IE.
type EMENBI struct {
	*PLMN
	EMENBI uint32
}

// NewEMENBI creates a new EMENBI.
func NewEMENBI(mcc, mnc string, menbi uint32) *EMENBI {
	return &EMENBI{
		PLMN:   &PLMN{MCC: mcc, MNC: mnc},
		EMENBI: menbi,
	}
}

// NewUserLocationInformationStruct creates a new UserLocationInformation IE from
// the structs defined in gtpv2/ie package. Give nil for unnecessary values.
func NewUserLocationInformationStruct(cgi *CGI, sai *SAI, rai *RAI, tai *TAI, ecgi *ECGI, lai *LAI, menbi *MENBI, emenbi *EMENBI) *IE {
	v := NewUserLocationInformationFields(cgi, sai, rai, tai, ecgi, lai, menbi, emenbi)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(UserLocationInformation, 0x00, b)
}

// UserLocationInformation returns UserLocationInformation in UserLocationInformationFields type
// if the type of IE matches.
func (i *IE) UserLocationInformation() (*UserLocationInformationFields, error) {
	switch i.Type {
	case UserLocationInformation:
		return ParseUserLocationInformationFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// UserLocationInformationFields is a set of fields in UserLocationInformation IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UserLocationInformationFields struct {
	Flags uint8
	*CGI
	*SAI
	*RAI
	*TAI
	*ECGI
	*LAI
	*MENBI
	*EMENBI
}

// NewUserLocationInformationFields creates a new NewUserLocationInformationFields.
func NewUserLocationInformationFields(cgi *CGI, sai *SAI, rai *RAI, tai *TAI, ecgi *ECGI, lai *LAI, menbi *MENBI, emenbi *EMENBI) *UserLocationInformationFields {
	f := &UserLocationInformationFields{}

	if cgi != nil {
		f.Flags |= 0x01
		f.CGI = cgi
	}
	if sai != nil {
		f.Flags |= 0x02
		f.SAI = sai
	}
	if rai != nil {
		f.Flags |= 0x04
		f.RAI = rai
	}
	if tai != nil {
		f.Flags |= 0x08
		f.TAI = tai
	}
	if ecgi != nil {
		f.Flags |= 0x10
		f.ECGI = ecgi
	}
	if lai != nil {
		f.Flags |= 0x20
		f.LAI = lai
	}
	if menbi != nil {
		f.Flags |= 0x40
		f.MENBI = menbi
	}
	if emenbi != nil {
		f.Flags |= 0x80
		f.EMENBI = emenbi
	}

	return f
}

// Marshal serializes UserLocationInformationFields.
func (f *UserLocationInformationFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes UserLocationInformationFields.
func (f *UserLocationInformationFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	b[0] = f.Flags
	offset := 1

	if has1stBit(f.Flags) { // CGI
		if l < offset+7 {
			return io.ErrUnexpectedEOF
		}

		plmn, err := utils.EncodePLMN(f.CGI.MCC, f.CGI.MNC)
		if err != nil {
			return err
		}
		copy(b[offset:offset+3], plmn)
		binary.BigEndian.PutUint16(b[offset+3:offset+5], f.CGI.LAC)
		binary.BigEndian.PutUint16(b[offset+5:offset+7], f.CGI.CI)
		offset += 7
	}
	if has2ndBit(f.Flags) { // SAI
		if l < offset+7 {
			return io.ErrUnexpectedEOF
		}

		plmn, err := utils.EncodePLMN(f.SAI.MCC, f.SAI.MNC)
		if err != nil {
			return err
		}
		copy(b[offset:offset+3], plmn)
		binary.BigEndian.PutUint16(b[offset+3:offset+5], f.SAI.LAC)
		binary.BigEndian.PutUint16(b[offset+5:offset+7], f.SAI.SAC)
		offset += 7
	}
	if has3rdBit(f.Flags) { // RAI
		if l < offset+7 {
			return io.ErrUnexpectedEOF
		}

		plmn, err := utils.EncodePLMN(f.RAI.MCC, f.RAI.MNC)
		if err != nil {
			return err
		}
		copy(b[offset:offset+3], plmn)
		binary.BigEndian.PutUint16(b[offset+3:offset+5], f.RAI.LAC)
		binary.BigEndian.PutUint16(b[offset+5:offset+7], f.RAI.RAC)
		offset += 7
	}
	if has4thBit(f.Flags) { // TAI
		if l < offset+5 {
			return io.ErrUnexpectedEOF
		}

		plmn, err := utils.EncodePLMN(f.TAI.MCC, f.TAI.MNC)
		if err != nil {
			return err
		}
		copy(b[offset:offset+3], plmn)
		binary.BigEndian.PutUint16(b[offset+3:offset+5], f.TAI.TAC)
		offset += 5
	}
	if has5thBit(f.Flags) { // ECGI
		if l < offset+7 {
			return io.ErrUnexpectedEOF
		}

		plmn, err := utils.EncodePLMN(f.ECGI.MCC, f.ECGI.MNC)
		if err != nil {
			return err
		}
		copy(b[offset:offset+3], plmn)
		binary.BigEndian.PutUint32(b[offset+3:offset+7], f.ECGI.ECI)
		offset += 7
	}
	if has6thBit(f.Flags) { // LAI
		if l < offset+5 {
			return io.ErrUnexpectedEOF
		}

		plmn, err := utils.EncodePLMN(f.LAI.MCC, f.LAI.MNC)
		if err != nil {
			return err
		}
		copy(b[offset:offset+3], plmn)
		binary.BigEndian.PutUint16(b[offset+3:offset+5], f.LAI.LAC)
		offset += 5
	}
	if has7thBit(f.Flags) { // MENBI
		if l < offset+6 {
			return io.ErrUnexpectedEOF
		}

		plmn, err := utils.EncodePLMN(f.MENBI.MCC, f.MENBI.MNC)
		if err != nil {
			return err
		}
		copy(b[offset:offset+3], plmn)
		copy(b[offset+3:offset+6], utils.Uint32To24(f.MENBI.MENBI))
		offset += 6
	}
	if has8thBit(f.Flags) { // EMENBI
		if l < offset+6 {
			return io.ErrUnexpectedEOF
		}

		plmn, err := utils.EncodePLMN(f.EMENBI.MCC, f.EMENBI.MNC)
		if err != nil {
			return err
		}
		copy(b[offset:offset+3], plmn)
		copy(b[offset+3:offset+6], utils.Uint32To24(f.EMENBI.EMENBI))
	}

	return nil
}

// ParseUserLocationInformationFields decodes UserLocationInformationFields.
func ParseUserLocationInformationFields(b []byte) (*UserLocationInformationFields, error) {
	f := &UserLocationInformationFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into UserLocationInformationFields.
func (f *UserLocationInformationFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}
	f.Flags = b[0]
	offset := 1

	var err error
	if has1stBit(f.Flags) {
		if l < offset+7 {
			return io.ErrUnexpectedEOF
		}

		f.CGI = &CGI{PLMN: &PLMN{}}
		f.CGI.PLMN.MCC, f.CGI.PLMN.MNC, err = utils.DecodePLMN(b[offset : offset+3])
		if err != nil {
			return err
		}
		f.CGI.LAC = binary.BigEndian.Uint16(b[offset+3 : offset+5])
		f.CGI.CI = binary.BigEndian.Uint16(b[offset+5 : offset+7])
		offset += 7
	}
	if has2ndBit(f.Flags) {
		if l < offset+7 {
			return io.ErrUnexpectedEOF
		}

		f.SAI = &SAI{PLMN: &PLMN{}}
		f.SAI.PLMN.MCC, f.SAI.PLMN.MNC, err = utils.DecodePLMN(b[offset : offset+3])
		if err != nil {
			return err
		}
		f.SAI.LAC = binary.BigEndian.Uint16(b[offset+3 : offset+5])
		f.SAI.SAC = binary.BigEndian.Uint16(b[offset+5 : offset+7])
		offset += 7
	}
	if has3rdBit(f.Flags) {
		if l < offset+7 {
			return io.ErrUnexpectedEOF
		}

		f.RAI = &RAI{PLMN: &PLMN{}}
		f.RAI.PLMN.MCC, f.RAI.PLMN.MNC, err = utils.DecodePLMN(b[offset : offset+3])
		if err != nil {
			return err
		}
		f.RAI.LAC = binary.BigEndian.Uint16(b[offset+3 : offset+5])
		f.RAI.RAC = binary.BigEndian.Uint16(b[offset+5 : offset+7])
		offset += 7
	}
	if has4thBit(f.Flags) {
		if l < offset+5 {
			return io.ErrUnexpectedEOF
		}

		f.TAI = &TAI{PLMN: &PLMN{}}
		f.TAI.PLMN.MCC, f.TAI.PLMN.MNC, err = utils.DecodePLMN(b[offset : offset+3])
		if err != nil {
			return err
		}
		f.TAI.TAC = binary.BigEndian.Uint16(b[offset+3 : offset+5])
		offset += 5
	}
	if has5thBit(f.Flags) {
		if l < offset+7 {
			return io.ErrUnexpectedEOF
		}

		f.ECGI = &ECGI{PLMN: &PLMN{}}
		f.ECGI.PLMN.MCC, f.ECGI.PLMN.MNC, err = utils.DecodePLMN(b[offset : offset+3])
		if err != nil {
			return err
		}
		f.ECGI.ECI = binary.BigEndian.Uint32(b[offset+3 : offset+7])
		offset += 7
	}
	if has6thBit(f.Flags) {
		if l < offset+5 {
			return io.ErrUnexpectedEOF
		}

		f.LAI = &LAI{PLMN: &PLMN{}}
		f.LAI.PLMN.MCC, f.LAI.PLMN.MNC, err = utils.DecodePLMN(b[offset : offset+3])
		if err != nil {
			return err
		}
		f.LAI.LAC = binary.BigEndian.Uint16(b[offset+3 : offset+5])
		offset += 5
	}
	if has7thBit(f.Flags) {
		if l < offset+6 {
			return io.ErrUnexpectedEOF
		}

		f.MENBI = &MENBI{PLMN: &PLMN{}}
		f.MENBI.PLMN.MCC, f.MENBI.PLMN.MNC, err = utils.DecodePLMN(b[offset : offset+3])
		if err != nil {
			return err
		}
		f.MENBI.MENBI = utils.Uint24To32(b[offset+3 : offset+6])
		offset += 6
	}
	if has8thBit(f.Flags) {
		if l < offset+6 {
			return io.ErrUnexpectedEOF
		}

		f.EMENBI = &EMENBI{PLMN: &PLMN{}}
		f.EMENBI.PLMN.MCC, f.EMENBI.PLMN.MNC, err = utils.DecodePLMN(b[offset : offset+3])
		if err != nil {
			return err
		}
		f.EMENBI.EMENBI = utils.Uint24To32(b[offset+3 : offset+6])
	}

	return nil
}

// MarshalLen returns the serial length of UserLocationInformationFields in int.
func (f *UserLocationInformationFields) MarshalLen() int {
	l := 1

	if has1stBit(f.Flags) {
		l += cgilen
	}
	if has2ndBit(f.Flags) {
		l += sailen
	}
	if has3rdBit(f.Flags) {
		l += railen
	}
	if has4thBit(f.Flags) {
		l += tailen
	}
	if has5thBit(f.Flags) {
		l += ecgilen
	}
	if has6thBit(f.Flags) {
		l += lailen
	}
	if has7thBit(f.Flags) {
		l += menbilen
	}
	if has8thBit(f.Flags) {
		l += emenbilen
	}

	return l
}

// NewUserLocationInformationLazy creates a new UserLocationInformation IE.
//
// The flags and corresponding fields are automatically set depending on the values given in int.
// If a value is less than 0, the field is considered as missing.
//
// Deprecated: use NewUserLocationInformationStruct instead. At some point this will be removed.
func NewUserLocationInformationLazy(mcc, mnc string, lac, ci, sac, rac, tac, eci, menbi, emenbi int) *IE {
	var hasCGI, hasSAI, hasRAI, hasTAI, hasECGI, hasLAI, hasMENBI, hasEMENBI uint8
	if ci >= 0 {
		hasCGI = 1
	}
	if sac >= 0 {
		hasSAI = 1
	}
	if rac >= 0 {
		hasRAI = 1
	}
	if tac >= 0 {
		hasTAI = 1
	}
	if eci >= 0 {
		hasECGI = 1
	}
	if lac >= 0 {
		hasLAI = 1
	}
	if menbi >= 0 {
		hasMENBI = 1
	}
	if emenbi >= 0 {
		hasEMENBI = 1
	}

	return NewUserLocationInformation(
		hasCGI, hasSAI, hasRAI, hasTAI, hasECGI, hasLAI, hasMENBI, hasEMENBI,
		mcc, mnc, uint16(lac), uint16(ci), uint16(sac), uint16(rac), uint16(tac),
		uint32(eci), uint32(menbi), uint32(emenbi),
	)
}

// NewUserLocationInformation creates a new UserLocationInformation IE.
//
// Deprecated: use NewUserLocationInformationStruct instead. At some point this will be removed.
func NewUserLocationInformation(
	hasCGI, hasSAI, hasRAI, hasTAI, hasECGI, hasLAI, hasMENBI, hasEMENBI uint8,
	mcc, mnc string, lac, ci, sac, rac, tac uint16, eci, menbi, emenbi uint32,
) *IE {
	flags := ((hasEMENBI & 0x01) << 7) |
		((hasMENBI & 0x01) << 6) |
		((hasLAI & 0x01) << 5) |
		((hasECGI & 0x01) << 4) |
		((hasTAI & 0x01) << 3) |
		((hasRAI & 0x01) << 2) |
		((hasSAI & 0x01) << 1) |
		(hasCGI & 0x01)

	i := New(UserLocationInformation, 0x00, make([]byte, uliPayloadLen(flags)))
	i.Payload[0] = flags

	plmn, err := utils.EncodePLMN(mcc, mnc)
	if err != nil {
		return nil
	}

	offset := 1
	if flags&0x01 == 1 {
		copy(i.Payload[offset:offset+3], plmn)
		binary.BigEndian.PutUint16(i.Payload[offset+3:offset+5], lac)
		binary.BigEndian.PutUint16(i.Payload[offset+5:offset+7], ci)
		offset += cgilen
	}
	if flags>>1&0x01 == 1 {
		copy(i.Payload[offset:offset+3], plmn)
		binary.BigEndian.PutUint16(i.Payload[offset+3:offset+5], lac)
		binary.BigEndian.PutUint16(i.Payload[offset+5:offset+7], sac)
		offset += sailen
	}
	if flags>>2&0x01 == 1 {
		copy(i.Payload[offset:offset+3], plmn)
		binary.BigEndian.PutUint16(i.Payload[offset+3:offset+5], lac)
		binary.BigEndian.PutUint16(i.Payload[offset+5:offset+7], rac)
		offset += railen
	}
	if flags>>3&0x01 == 1 {
		copy(i.Payload[offset:offset+3], plmn)
		binary.BigEndian.PutUint16(i.Payload[offset+3:offset+5], tac)
		offset += tailen
	}
	if flags>>4&0x01 == 1 {
		copy(i.Payload[offset:offset+3], plmn)
		eci &= 0xfffff
		binary.BigEndian.PutUint32(i.Payload[offset+3:offset+7], eci)
		offset += ecgilen
	}
	if flags>>5&0x01 == 1 {
		copy(i.Payload[offset:offset+3], plmn)
		binary.BigEndian.PutUint16(i.Payload[offset+3:offset+5], lac)
		offset += lailen
	}
	if flags>>6&0x01 == 1 {
		copy(i.Payload[offset:offset+3], plmn)
		copy(i.Payload[offset+3:offset+6], utils.Uint32To24(menbi))
		offset += menbilen
	}
	if flags>>7&0x01 == 1 {
		copy(i.Payload[offset:offset+3], plmn)
		copy(i.Payload[offset+3:offset+6], utils.Uint32To24(emenbi))
	}
	return i
}

func uliPayloadLen(flags uint8) int {
	l := 1
	if flags&0x01 == 1 {
		l += cgilen
	}
	if flags>>1&0x01 == 1 {
		l += sailen
	}
	if flags>>2&0x01 == 1 {
		l += railen
	}
	if flags>>3&0x01 == 1 {
		l += tailen
	}
	if flags>>4&0x01 == 1 {
		l += ecgilen
	}
	if flags>>5&0x01 == 1 {
		l += lailen
	}
	if flags>>6&0x01 == 1 {
		l += menbilen
	}
	if flags>>7&0x01 == 1 {
		l += emenbilen
	}
	return l
}

// UserLocationInfo is a getter function to parse UserLocationInformationFields
//
// Deprecated: use UserLocationInformation instead. At some point this will be removed.
func (i *IE) UserLocationInfo() (*UserLocationInformationFields, error) {
	var uli UserLocationInformationFields
	var plmn PLMN
	l := len(i.Payload)
	if l == 0 {
		return &uli, io.ErrUnexpectedEOF
	}
	offset := 1
	if i.Payload[0]&0x01 == 1 {
		if l < (offset + cgilen) {
			return &uli, io.ErrUnexpectedEOF
		}
		var cgi CGI
		uli.CGI = &cgi
		uli.CGI.PLMN = &plmn
		uli.CGI.PLMN.MCC, uli.CGI.PLMN.MNC, _ = utils.DecodePLMN(i.Payload[offset : offset+3])
		uli.CGI.LAC = binary.BigEndian.Uint16(i.Payload[offset+3 : offset+5])
		uli.CGI.CI = binary.BigEndian.Uint16(i.Payload[offset+5 : offset+7])
		offset += cgilen
	}
	if i.Payload[0]>>1&0x01 == 1 {
		if l < (offset + sailen) {
			return &uli, io.ErrUnexpectedEOF
		}
		var sai SAI
		uli.SAI = &sai
		uli.SAI.PLMN = &plmn
		uli.SAI.PLMN.MCC, uli.SAI.PLMN.MNC, _ = utils.DecodePLMN(i.Payload[offset : offset+3])
		uli.SAI.LAC = binary.BigEndian.Uint16(i.Payload[offset+3 : offset+5])
		uli.SAI.SAC = binary.BigEndian.Uint16(i.Payload[offset+5 : offset+7])
		offset += sailen
	}
	if i.Payload[0]>>2&0x01 == 1 {
		if l < (offset + railen) {
			return &uli, io.ErrUnexpectedEOF
		}
		var rai RAI
		uli.RAI = &rai
		uli.RAI.PLMN = &plmn
		uli.RAI.PLMN.MCC, uli.RAI.PLMN.MNC, _ = utils.DecodePLMN(i.Payload[offset : offset+3])
		uli.RAI.LAC = binary.BigEndian.Uint16(i.Payload[offset+3 : offset+5])
		uli.RAI.RAC = binary.BigEndian.Uint16(i.Payload[offset+5 : offset+7])
		offset += railen
	}
	if i.Payload[0]>>3&0x01 == 1 {
		if l < (offset + tailen) {
			return &uli, io.ErrUnexpectedEOF
		}
		var tai TAI
		uli.TAI = &tai
		uli.TAI.PLMN = &plmn
		uli.TAI.PLMN.MCC, uli.TAI.PLMN.MNC, _ = utils.DecodePLMN(i.Payload[offset : offset+3])
		uli.TAI.TAC = binary.BigEndian.Uint16(i.Payload[offset+3 : offset+5])
		offset += tailen

	}
	if i.Payload[0]>>4&0x01 == 1 {
		if l < (offset + ecgilen) {
			return &uli, io.ErrUnexpectedEOF
		}
		var ecgi ECGI
		uli.ECGI = &ecgi
		uli.ECGI.PLMN = &plmn
		uli.ECGI.PLMN.MCC, uli.ECGI.PLMN.MNC, _ = utils.DecodePLMN(i.Payload[offset : offset+3])
		uli.ECGI.ECI = binary.BigEndian.Uint32(i.Payload[offset+3 : offset+7])
		offset += ecgilen

	}
	if i.Payload[0]>>5&0x01 == 1 {
		if l < (offset + lailen) {
			return &uli, io.ErrUnexpectedEOF
		}
		var lai LAI
		uli.LAI = &lai
		uli.LAI.PLMN = &plmn
		uli.LAI.PLMN.MCC, uli.LAI.PLMN.MNC, _ = utils.DecodePLMN(i.Payload[offset : offset+3])
		uli.LAI.LAC = binary.BigEndian.Uint16(i.Payload[offset+3 : offset+5])
		offset += lailen
	}
	if i.Payload[0]>>6&0x01 == 1 {
		if l < (offset + menbilen) {
			return &uli, io.ErrUnexpectedEOF
		}
		var menbi MENBI
		uli.MENBI = &menbi
		uli.MENBI.PLMN = &plmn
		uli.MENBI.PLMN.MCC, uli.MENBI.PLMN.MNC, _ = utils.DecodePLMN(i.Payload[offset : offset+3])
		uli.MENBI.MENBI = utils.Uint24To32(i.Payload[offset+3 : offset+6])
		offset += menbilen
	}
	if i.Payload[0]>>7&0x01 == 1 {
		if l < (offset + emenbilen) {
			return &uli, io.ErrUnexpectedEOF
		}
		var emenbi EMENBI
		uli.EMENBI = &emenbi
		uli.EMENBI.PLMN = &plmn
		uli.EMENBI.PLMN.MCC, uli.EMENBI.PLMN.MNC, _ = utils.DecodePLMN(i.Payload[offset : offset+3])
		uli.EMENBI.EMENBI = utils.Uint24To32(i.Payload[offset+3 : offset+6])
	}
	return &uli, nil
}
