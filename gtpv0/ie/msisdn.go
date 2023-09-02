// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"

	"github.com/wmnsk/go-gtp/utils"
)

// NewMSISDN creates a new MSISDN IE.
func NewMSISDN(msisdn string) *IE {
	i, err := utils.StrToSwappedBytes("19"+msisdn, "f")
	if err != nil {
		return nil
	}
	return New(MSISDN, i)
}

// MSISDN returns MSISDN value if type matches.
func (i *IE) MSISDN() (string, error) {
	if i.Type != MSISDN {
		return "", &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 2 {
		return "", io.ErrUnexpectedEOF
	}

	return utils.SwappedBytesToStr(i.Payload[1:], false), nil
}

// MustMSISDN returns MSISDN in string if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustMSISDN() string {
	v, _ := i.MSISDN()
	return v
}
