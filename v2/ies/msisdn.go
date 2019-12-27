// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"io"

	"github.com/ErvinsK/go-gtp/utils"
)

// NewMSISDN creates a new MSISDN IE.
func NewMSISDN(mei string) *IE {
	m, err := utils.StrToSwappedBytes(mei, "f")
	if err != nil {
		return nil
	}
	return New(MSISDN, 0x00, m)
}

// MSISDN returns MSISDN in string if the type of IE matches.
func (i *IE) MSISDN() (string, error) {
	if i.Type != MSISDN {
		return "", &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return "", io.ErrUnexpectedEOF
	}
	return utils.SwappedBytesToStr(i.Payload, true), nil
}

// MustMSISDN returns MSISDN in string, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustMSISDN() string {
	v, _ := i.MSISDN()
	return v
}
