// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"io"

	"github.com/wmnsk/go-gtp/utils"
)

// NewMobileEquipmentIdentity creates a new MobileEquipmentIdentity IE.
func NewMobileEquipmentIdentity(mei string) *IE {
	m, err := utils.StrToSwappedBytes(mei, "f")
	if err != nil {
		return nil
	}
	return New(MobileEquipmentIdentity, 0x00, m)
}

// MobileEquipmentIdentity returns MobileEquipmentIdentity in string if the
// type of IE matches.
func (i *IE) MobileEquipmentIdentity() (string, error) {
	if i.Type != MobileEquipmentIdentity {
		return "", &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return "", io.ErrUnexpectedEOF
	}
	return utils.SwappedBytesToStr(i.Payload, true), nil
}

// MustMobileEquipmentIdentity returns MobileEquipmentIdentity in string, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustMobileEquipmentIdentity() string {
	v, _ := i.MobileEquipmentIdentity()
	return v
}
