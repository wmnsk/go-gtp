// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"github.com/wmnsk/go-gtp/gtp/utils"
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
func (i *IE) ServingNetwork() string {
	if i.Type != ServingNetwork {
		return ""
	}

	mcc, mnc, err := utils.DecodePLMN(i.Payload)
	if err != nil {
		return ""
	}

	return mcc + mnc
}

// MCC returns MCC in string if the type of IE matches.
func (i *IE) MCC() string {
	switch i.Type {
	case ServingNetwork, PLMNID:
		mcc, _, err := utils.DecodePLMN(i.Payload)
		if err != nil {
			return ""
		}
		return mcc
	case GlobalCNID, TraceReference, GUTI, UserCSGInformation:
		mcc, _, err := utils.DecodePLMN(i.Payload[:3])
		if err != nil {
			return ""
		}
		return mcc
	default:
		return ""
	}
}

// MNC returns MNC in string if the type of IE matches.
func (i *IE) MNC() string {
	switch i.Type {
	case ServingNetwork, PLMNID:
		_, mnc, err := utils.DecodePLMN(i.Payload)
		if err != nil {
			return ""
		}
		return mnc
	case GlobalCNID, TraceReference, GUTI, UserCSGInformation:
		_, mnc, err := utils.DecodePLMN(i.Payload[:3])
		if err != nil {
			return ""
		}
		return mnc
	default:
		return ""
	}
}
