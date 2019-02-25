// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "github.com/wmnsk/go-gtp/gtp/utils"

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
