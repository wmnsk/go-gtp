// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"errors"
	"io"

	"github.com/wmnsk/go-gtp/utils"
)

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
	case UserLocationInformation:
		uliFields, err := i.UserLocationInformation()
		if err != nil {
			return "", err
		}
		switch {
		case uliFields.HasCGI():
			return uliFields.CGI.MCCFromPLMN(), nil
		case uliFields.HasSAI():
			return uliFields.SAI.MCCFromPLMN(), nil
		case uliFields.HasRAI():
			return uliFields.RAI.MCCFromPLMN(), nil
		case uliFields.HasTAI():
			return uliFields.TAI.MCCFromPLMN(), nil
		case uliFields.HasECGI():
			return uliFields.ECGI.MCCFromPLMN(), nil
		case uliFields.HasLAI():
			return uliFields.LAI.MCCFromPLMN(), nil
		case uliFields.HasMENBI():
			return uliFields.MENBI.MCCFromPLMN(), nil
		case uliFields.HasEMENBI():
			return uliFields.EMENBI.MCCFromPLMN(), nil
		}
		return "", errors.New("MCC is not present in ULI")
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
	case UserLocationInformation:
		uliFields, err := i.UserLocationInformation()
		if err != nil {
			return "", err
		}
		switch {
		case uliFields.HasCGI():
			return uliFields.CGI.MNCFromPLMN(), nil
		case uliFields.HasSAI():
			return uliFields.SAI.MNCFromPLMN(), nil
		case uliFields.HasRAI():
			return uliFields.RAI.MNCFromPLMN(), nil
		case uliFields.HasTAI():
			return uliFields.TAI.MNCFromPLMN(), nil
		case uliFields.HasECGI():
			return uliFields.ECGI.MNCFromPLMN(), nil
		case uliFields.HasLAI():
			return uliFields.LAI.MNCFromPLMN(), nil
		case uliFields.HasMENBI():
			return uliFields.MENBI.MNCFromPLMN(), nil
		case uliFields.HasEMENBI():
			return uliFields.EMENBI.MNCFromPLMN(), nil
		}
		return "", errors.New("MNC is not present in ULI")
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
