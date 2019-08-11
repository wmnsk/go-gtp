// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"github.com/wmnsk/go-gtp/utils"
)

// NewPLMNID creates a PLMNID IE.
func NewPLMNID(mcc, mnc string) *IE {
	encoded, err := utils.EncodePLMN(mcc, mnc)
	if err != nil {
		return nil
	}

	return New(PLMNID, 0x00, encoded)
}

// PLMNID returns PLMNID(MCC and MNC) in string if the type of IE matches.
func (i *IE) PLMNID() string {
	if i.Type != PLMNID {
		return ""
	}
	if len(i.Payload) < 3 {
		return ""
	}

	mcc, mnc, err := utils.DecodePLMN(i.Payload)
	if err != nil {
		return ""
	}

	return mcc + mnc
}
