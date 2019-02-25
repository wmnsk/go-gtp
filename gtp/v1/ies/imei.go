// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "github.com/wmnsk/go-gtp/gtp/utils"

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
