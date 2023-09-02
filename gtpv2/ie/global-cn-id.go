// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"

	"github.com/wmnsk/go-gtp/utils"
)

// NewGlobalCNID creates a new GlobalCNID IE.
func NewGlobalCNID(mcc, mnc string, cnid uint16) *IE {
	i := New(GlobalCNID, 0x00, make([]byte, 5))
	plmn, err := utils.EncodePLMN(mcc, mnc)
	if err != nil {
		return nil
	}
	copy(i.Payload[0:3], plmn)

	cnid &= 0xfff
	binary.BigEndian.PutUint16(i.Payload[3:5], cnid)
	return i
}

// CNID returns CNID in uinte16 if the type of IE matches.
func (i *IE) CNID() (uint16, error) {
	if len(i.Payload) < 5 {
		return 0, io.ErrUnexpectedEOF
	}
	switch i.Type {
	case GlobalCNID:
		return binary.BigEndian.Uint16(i.Payload[3:5]), nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustCNID returns CNID in uint16, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustCNID() uint16 {
	v, _ := i.CNID()
	return v
}
