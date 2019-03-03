// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewQualityOfServiceProfile creates a new QualityOfServiceProfile IE.
func NewQualityOfServiceProfile(delay, reliability, peak, precedence, mean uint8) *IE {
	i := New(QualityOfServiceProfile, make([]byte, 3))
	i.Payload[0] = ((delay & 0x07) << 3) | (reliability & 0x07)
	i.Payload[1] = ((peak & 0x0f) << 4) | (precedence & 0x07)
	i.Payload[2] = mean & 0x1f

	return i
}

// QualityOfServiceProfile returns QualityOfServiceProfile if type matches.
func (i *IE) QualityOfServiceProfile() []byte {
	if i.Type != QualityOfServiceProfile {
		return nil
	}
	return i.Payload
}

// QoSDelay returns QoS Delay value in uint8 if type matches.
func (i *IE) QoSDelay() uint8 {
	if i.Type != QualityOfServiceProfile {
		return 0
	}
	return i.Payload[0] & 0x38
}

// QoSReliability returns QoS Reliability value in uint8 if type matches.
func (i *IE) QoSReliability() uint8 {
	if i.Type != QualityOfServiceProfile {
		return 0
	}
	return i.Payload[0] & 0x07
}

// QoSPeak returns QoS Peak value in uint8 if type matches.
func (i *IE) QoSPeak() uint8 {
	if i.Type != QualityOfServiceProfile {
		return 0
	}
	return i.Payload[1] & 0xf0
}

// QoSPrecedence returns QoS Precedence value in uint8 if type matches.
func (i *IE) QoSPrecedence() uint8 {
	if i.Type != QualityOfServiceProfile {
		return 0
	}
	return i.Payload[1] & 0x07
}

// QoSMean returns QoS Mean value in uint8 if type matches.
func (i *IE) QoSMean() uint8 {
	if i.Type != QualityOfServiceProfile {
		return 0
	}
	return i.Payload[2] & 0x0f
}
