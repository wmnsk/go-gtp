// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "io"

// NewQualityOfServiceProfile creates a new QualityOfServiceProfile IE.
func NewQualityOfServiceProfile(delay, reliability, peak, precedence, mean uint8) *IE {
	i := New(QualityOfServiceProfile, make([]byte, 3))
	i.Payload[0] = ((delay & 0x07) << 3) | (reliability & 0x07)
	i.Payload[1] = ((peak & 0x0f) << 4) | (precedence & 0x07)
	i.Payload[2] = mean & 0x1f

	return i
}

// QualityOfServiceProfile returns QualityOfServiceProfile if type matches.
func (i *IE) QualityOfServiceProfile() ([]byte, error) {
	if i.Type != QualityOfServiceProfile {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	return i.Payload, nil
}

// MustQualityOfServiceProfile returns QualityOfServiceProfile in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustQualityOfServiceProfile() []byte {
	v, _ := i.QualityOfServiceProfile()
	return v
}

// QoSDelay returns QoS Delay value in uint8 if type matches.
func (i *IE) QoSDelay() (uint8, error) {
	if i.Type != QualityOfServiceProfile {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0] & 0x38, nil
}

// MustQoSDelay returns QoSDelay in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustQoSDelay() uint8 {
	v, _ := i.QoSDelay()
	return v
}

// QoSReliability returns QoS Reliability value in uint8 if type matches.
func (i *IE) QoSReliability() (uint8, error) {
	if i.Type != QualityOfServiceProfile {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0] & 0x07, nil
}

// MustQoSReliability returns QoSReliability in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustQoSReliability() uint8 {
	v, _ := i.QoSReliability()
	return v
}

// QoSPeak returns QoS Peak value in uint8 if type matches.
func (i *IE) QoSPeak() (uint8, error) {
	if i.Type != QualityOfServiceProfile {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[1] & 0xf0, nil
}

// MustQoSPeak returns QoSPeak in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustQoSPeak() uint8 {
	v, _ := i.QoSPeak()
	return v
}

// QoSPrecedence returns QoS Precedence value in uint8 if type matches.
func (i *IE) QoSPrecedence() (uint8, error) {
	if i.Type != QualityOfServiceProfile {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[1] & 0x07, nil
}

// MustQoSPrecedence returns QoSPrecedence in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustQoSPrecedence() uint8 {
	v, _ := i.QoSPrecedence()
	return v
}

// QoSMean returns QoS Mean value in uint8 if type matches.
func (i *IE) QoSMean() (uint8, error) {
	if i.Type != QualityOfServiceProfile {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 3 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[2] & 0x0f, nil
}

// MustQoSMean returns QoSMean in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustQoSMean() uint8 {
	v, _ := i.QoSMean()
	return v
}
