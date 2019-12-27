// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"io"

	"github.com/ErvinsK/go-gtp/utils"
)

// NewServingNetwork creates a ServingNetwork IE.
func NewServingNetwork(mcc, mnc string) *IE {
	encoded, err := utils.EncodePLMN(mcc, mnc)
	if err != nil {
		return nil
	}

	return New(ServingNetwork, 0x00, encoded)
}

// ServingNetwork returns ServingNetwork(MCC and MNC) in string if the type of IE matches.
func (i *IE) ServingNetwork() (string, error) {
	if i.Type != ServingNetwork {
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

// MustServingNetwork returns ServingNetwork in string, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustServingNetwork() string {
	v, _ := i.ServingNetwork()
	return v
}

// MCC returns MCC in string if the type of IE matches.
func (i *IE) MCC() (string, error) {
	if len(i.Payload) < 3 {
		return "", io.ErrUnexpectedEOF
	}

	switch i.Type {
	case ServingNetwork, PLMNID:
		mcc, _, err := utils.DecodePLMN(i.Payload)
		if err != nil {
			return "", err
		}
		return mcc, nil
	case GlobalCNID, TraceReference, GUTI, UserCSGInformation:
		mcc, _, err := utils.DecodePLMN(i.Payload[:3])
		if err != nil {
			return "", err
		}
		return mcc, nil
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}

// MustMCC returns MCC in string, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustMCC() string {
	v, _ := i.MCC()
	return v
}

// MNC returns MNC in string if the type of IE matches.
func (i *IE) MNC() (string, error) {
	if len(i.Payload) < 3 {
		return "", io.ErrUnexpectedEOF
	}

	switch i.Type {
	case ServingNetwork, PLMNID:
		_, mnc, err := utils.DecodePLMN(i.Payload)
		if err != nil {
			return "", err
		}
		return mnc, nil
	case GlobalCNID, TraceReference, GUTI, UserCSGInformation:
		_, mnc, err := utils.DecodePLMN(i.Payload[:3])
		if err != nil {
			return "", err
		}
		return mnc, nil
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}

// MustMNC returns MNC in string, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustMNC() string {
	v, _ := i.MNC()
	return v
}
