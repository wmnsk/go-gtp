// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

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

//ULI decoded info block
type ULI struct {
	Cgi    CGI
	Sai    SAI
	Rai    RAI
	Tai    TAI
	Ecgi   ECGI
	Lai    LAI
	Menbi  MENBI
	Emenbi EMENBI
}

//PLMN decoded info fields
type PLMN struct {
	Mcc string
	Mnc string
}

//CGI field
type CGI struct {
	Plmn PLMN
	Lac  uint16
	Ci   uint16
}

//SAI field
type SAI struct {
	Plmn PLMN
	Lac  uint16
	Sac  uint16
}

//RAI field
type RAI struct {
	Plmn PLMN
	Lac  uint16
	Rac  uint16
}

//TAI field
type TAI struct {
	Plmn PLMN
	Tac  uint16
}

//ECGI field
type ECGI struct {
	Plmn PLMN
	Eci  uint32
}

//LAI field
type LAI struct {
	Plmn PLMN
	Lac  uint16
}

//MENBI field
type MENBI struct {
	Plmn  PLMN
	Menbi uint32
}

//EMENBI field
type EMENBI struct {
	Plmn   PLMN
	Emenbi uint32
}

// NewUserLocationInformationLazy creates a new UserLocationInformation IE.
//
// The flags and corresponding fields are automatically set depending on the values given in int.
// If a value is less than 0, the field is considered as missing.
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

// UserLocationInfo is a getter function to parse ULI
func (i *IE) UserLocationInfo() (ULI, error) {
	var uli ULI
	if len(i.Payload) == 0 {
		return uli, io.ErrUnexpectedEOF
	}
	offset := 1
	if i.Payload[0]&0x01 == 1 {
		uli.Cgi.Plmn.Mcc, uli.Cgi.Plmn.Mnc, _ = utils.DecodePLMN(i.Payload[offset : offset+3])
		uli.Cgi.Lac = binary.BigEndian.Uint16(i.Payload[offset+3 : offset+5])
		uli.Cgi.Ci = binary.BigEndian.Uint16(i.Payload[offset+5 : offset+7])
		offset += cgilen
	}
	if i.Payload[0]>>1&0x01 == 1 {
		uli.Sai.Plmn.Mcc, uli.Sai.Plmn.Mnc, _ = utils.DecodePLMN(i.Payload[offset : offset+3])
		uli.Sai.Lac = binary.BigEndian.Uint16(i.Payload[offset+3 : offset+5])
		uli.Sai.Sac = binary.BigEndian.Uint16(i.Payload[offset+5 : offset+7])
		offset += sailen
	}
	if i.Payload[0]>>2&0x01 == 1 {
		uli.Rai.Plmn.Mcc, uli.Rai.Plmn.Mnc, _ = utils.DecodePLMN(i.Payload[offset : offset+3])
		uli.Rai.Lac = binary.BigEndian.Uint16(i.Payload[offset+3 : offset+5])
		uli.Rai.Rac = binary.BigEndian.Uint16(i.Payload[offset+5 : offset+7])
		offset += railen
	}
	if i.Payload[0]>>3&0x01 == 1 {
		uli.Tai.Plmn.Mcc, uli.Tai.Plmn.Mnc, _ = utils.DecodePLMN(i.Payload[offset : offset+3])
		uli.Tai.Tac = binary.BigEndian.Uint16(i.Payload[offset+3 : offset+5])
		offset += tailen
	}
	if i.Payload[0]>>4&0x01 == 1 {
		uli.Ecgi.Plmn.Mcc, uli.Ecgi.Plmn.Mnc, _ = utils.DecodePLMN(i.Payload[offset : offset+3])
		uli.Ecgi.Eci = binary.BigEndian.Uint32(i.Payload[offset+3 : offset+7])
		offset += ecgilen
	}
	if i.Payload[0]>>5&0x01 == 1 {
		uli.Lai.Plmn.Mcc, uli.Lai.Plmn.Mnc, _ = utils.DecodePLMN(i.Payload[offset : offset+3])
		uli.Lai.Lac = binary.BigEndian.Uint16(i.Payload[offset+3 : offset+5])
		offset += lailen
	}
	if i.Payload[0]>>6&0x01 == 1 {
		uli.Menbi.Plmn.Mcc, uli.Menbi.Plmn.Mnc, _ = utils.DecodePLMN(i.Payload[offset : offset+3])
		uli.Menbi.Menbi = utils.Uint24To32(i.Payload[offset+3 : offset+6])
		offset += menbilen
	}
	if i.Payload[0]>>7&0x01 == 1 {
		uli.Emenbi.Plmn.Mcc, uli.Emenbi.Plmn.Mnc, _ = utils.DecodePLMN(i.Payload[offset : offset+3])
		uli.Emenbi.Emenbi = utils.Uint24To32(i.Payload[offset+3 : offset+6])
	}
	return uli, nil
}
