// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"

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
func (i *IE) PLMNID() (string, error) {
	if i.Type != PLMNID {
		return "", &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 3 {
		return "", io.ErrUnexpectedEOF
	}

	mcc, mnc, err := utils.DecodePLMN(i.Payload)
	if err != nil {
		return "", err
	}

	return mcc + mnc, nil
}

// MustPLMNID returns PLMNID in string, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustPLMNID() string {
	v, _ := i.PLMNID()
	return v
}
