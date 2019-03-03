// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "github.com/wmnsk/go-gtp/utils"

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
func (i *IE) MobileEquipmentIdentity() string {
	return utils.SwappedBytesToStr(i.Payload, true)
}
