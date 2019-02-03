// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"

	"github.com/wmnsk/go-gtp/gtp/utils"
)

// NewGUTI creates a new GUTI IE.
func NewGUTI(mcc, mnc string, groupID uint16, code uint8, mTMSI uint32) *IE {
	i := New(GUTI, 0x00, make([]byte, 10))
	plmn, err := utils.EncodePLMN(mcc, mnc)
	if err != nil {
		return nil
	}
	copy(i.Payload[0:3], plmn)
	binary.BigEndian.PutUint16(i.Payload[3:5], groupID)
	i.Payload[5] = code
	binary.BigEndian.PutUint32(i.Payload[6:10], mTMSI)
	return i
}

// MMEGroupID returns MMEGroupID in uint16 if the type of IE matches.
func (i *IE) MMEGroupID() uint16 {
	switch i.Type {
	case GUTI:
		return binary.BigEndian.Uint16(i.Payload[3:5])
	default:
		return 0
	}
}

// MMECode returns MMECode in uint8 if the type of IE matches.
func (i *IE) MMECode() uint8 {
	switch i.Type {
	case GUTI:
		return i.Payload[5]
	default:
		return 0
	}
}

// MTMSI returns MTMSI in uint32 if the type of IE matches.
func (i *IE) MTMSI() uint32 {
	switch i.Type {
	case GUTI:
		return binary.BigEndian.Uint32(i.Payload[6:10])
	default:
		return 0
	}
}
