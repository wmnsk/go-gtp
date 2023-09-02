// Copyright 2019-2023 go-gtp authors. All rights reservev.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"errors"

	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// VersionNotSupportedIndication is a VersionNotSupportedIndication Header and its IEs above.
type VersionNotSupportedIndication struct {
	*Header
	AdditionalIEs []*ie.IE
}

// NewVersionNotSupportedIndication creates a new VersionNotSupportedIndication.
func NewVersionNotSupportedIndication(teid, seq uint32, ie ...*ie.IE) *VersionNotSupportedIndication {
	v := &VersionNotSupportedIndication{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeVersionNotSupportedIndication, teid, seq, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		v.AdditionalIEs = append(v.AdditionalIEs, i)
	}

	v.SetLength()
	return v
}

// Marshal serializes VersionNotSupportedIndication into bytes.
func (v *VersionNotSupportedIndication) Marshal() ([]byte, error) {
	b := make([]byte, v.MarshalLen())
	if err := v.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes VersionNotSupportedIndication into bytes.
func (v *VersionNotSupportedIndication) MarshalTo(b []byte) error {
	if v.Header.Payload != nil {
		v.Header.Payload = nil
	}
	v.Header.Payload = make([]byte, v.MarshalLen()-v.Header.MarshalLen())

	offset := 0
	for _, ie := range v.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.MarshalTo(v.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	v.Header.SetLength()
	return v.Header.MarshalTo(b)
}

// ParseVersionNotSupportedIndication decodes given bytes as VersionNotSupportedIndication.
func ParseVersionNotSupportedIndication(b []byte) (*VersionNotSupportedIndication, error) {
	v := &VersionNotSupportedIndication{}
	if err := v.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return v, nil
}

// UnmarshalBinary decodes given bytes as VersionNotSupportedIndication.
func (v *VersionNotSupportedIndication) UnmarshalBinary(b []byte) error {
	var err error
	v.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(v.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ie.ParseMultiIEs(v.Header.Payload)
	if err != nil {
		if errors.Is(err, ErrTooShortToParse) {
			return nil
		}
		return err
	}

	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		v.AdditionalIEs = append(v.AdditionalIEs, i)
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (v *VersionNotSupportedIndication) MarshalLen() int {
	l := v.Header.MarshalLen() - len(v.Header.Payload)
	for _, ie := range v.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.MarshalLen()
	}
	return l
}

// SetLength sets the length in Length fielv.
func (v *VersionNotSupportedIndication) SetLength() {
	v.Header.Length = uint16(v.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (v *VersionNotSupportedIndication) MessageTypeName() string {
	return "Version Not Supported Indication"
}

// TEID returns the TEID in uint32.
func (v *VersionNotSupportedIndication) TEID() uint32 {
	return v.Header.teid()
}
