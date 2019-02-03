// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"github.com/wmnsk/go-gtp/gtp/utils"
)

// NewIMSI creates a new IMSI IE.
func NewIMSI(imsi string) *IE {
	i, err := utils.StrToSwappedBytes(imsi, "f")
	if err != nil {
		return nil
	}
	return New(IMSI, i)
}

// IMSI returns IMSI value if type matches.
func (i *IE) IMSI() string {
	if i.Type != IMSI {
		return ""
	}
	return utils.SwappedBytesToStr(i.Payload, true)
}

// NewMSISDN creates a new MSISDN IE.
func NewMSISDN(msisdn string) *IE {
	i, err := utils.StrToSwappedBytes("19"+msisdn, "f")
	if err != nil {
		return nil
	}
	return New(MSISDN, i)
}

// MSISDN returns MSISDN value if type matches.
func (i *IE) MSISDN() string {
	if i.Type != MSISDN {
		return ""
	}
	return utils.SwappedBytesToStr(i.Payload[1:], false)
}

// NewIMEISV creates a new IMEISV IE.
func NewIMEISV(imei string) *IE {
	i, err := utils.StrToSwappedBytes(imei, "f")
	if err != nil {
		return nil
	}
	return New(IMEISV, i)
}

// IMEISV returns IMEISV value if type matches.
func (i *IE) IMEISV() string {
	if i.Type != IMEISV {
		return ""
	}
	return utils.SwappedBytesToStr(i.Payload, true)
}
