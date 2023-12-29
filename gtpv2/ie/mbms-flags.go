// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewMBMSFlags creates a new MBMSFlags IE.
func NewMBMSFlags(lmri, msri uint8) *IE {
	i := New(MBMSFlags, 0x00, make([]byte, 1))
	i.Payload[0] |= (lmri << 1 & 0x02) | (msri & 0x01)
	return i
}

// MBMSFlags returns MBMSFlags in uint8 if the type of IE matches.
func (i *IE) MBMSFlags() (uint8, error) {
	if i.Type != MBMSFlags {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	return i.ValueAsUint8()
}

// MustMBMSFlags returns MBMSFlags in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustMBMSFlags() uint8 {
	v, _ := i.MBMSFlags()
	return v
}

// HasMSRI reports whether an IE has MSRI bit.
func (i *IE) HasMSRI() bool {
	v, err := i.MBMSFlags()
	if err != nil {
		return false
	}

	return has1stBit(v)
}

// HasLMRI reports whether an IE has LMRI bit.
func (i *IE) HasLMRI() bool {
	v, err := i.MBMSFlags()
	if err != nil {
		return false
	}

	return has2ndBit(v)
}

// LocalMBMSBearerContextRelease reports whether the MBMS Session Stop Request
// message is used to release the MBMS Bearer Context locally in the MME/SGSN.
func (i *IE) LocalMBMSBearerContextRelease() bool {
	v, err := i.MBMSFlags()
	if err != nil {
		return false
	}
	return v&0x02 == 1
}

// MBMSSessionReEstablishment reports whether the MBMS Session Start Request
// message is used to re-establish an MBMS session.
func (i *IE) MBMSSessionReEstablishment() bool {
	v, err := i.MBMSFlags()
	if err != nil {
		return false
	}
	return v&0x01 == 1
}
