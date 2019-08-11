// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "github.com/wmnsk/go-gtp/utils"

// NewIMSI creates a new IMSI IE.
func NewIMSI(imsi string) *IE {
	i, err := utils.StrToSwappedBytes(imsi, "f")
	if err != nil {
		return New(IMSI, nil)
	}
	return New(IMSI, i)
}

// IMSI returns IMSI value in string if type matches.
func (i *IE) IMSI() string {
	if i.Type != IMSI {
		return ""
	}
	if len(i.Payload) == 0 {
		return ""
	}

	return utils.SwappedBytesToStr(i.Payload, true)
}
