// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"

	"github.com/wmnsk/go-gtp/utils"
)

// NewUserCSGInformation creates a new UserCSGInformation IE.
func NewUserCSGInformation(mcc, mnc string, csgID uint32, mode, lcsg, cmi uint8) *IE {
	i := New(UserCSGInformation, 0x00, make([]byte, 8))
	plmn, err := utils.EncodePLMN(mcc, mnc)
	if err != nil {
		return nil
	}
	copy(i.Payload[0:3], plmn)
	binary.BigEndian.PutUint32(i.Payload[3:7], csgID&0x7ffffff)
	i.Payload[7] = (mode << 6) | ((lcsg << 1) & 0x02) | (cmi & 0x01)
	return i
}

// AccessMode returns AccessMode in uint8 if the type of IE matches.
func (i *IE) AccessMode() uint8 {
	switch i.Type {
	case UserCSGInformation:
		if len(i.Payload) < 8 {
			return 0
		}
		return i.Payload[7] >> 6
	default:
		return 0
	}
}
