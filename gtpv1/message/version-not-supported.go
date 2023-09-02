// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv1/ie"
)

// VersionNotSupported is a VersionNotSupported Header and its IEs above.
type VersionNotSupported struct {
	*Header
	AdditionalIEs []*ie.IE
}

// NewVersionNotSupported creates a new GTPv1 VersionNotSupported.
func NewVersionNotSupported(teid uint32, seq uint16, ie ...*ie.IE) *VersionNotSupported {
	v := &VersionNotSupported{
		Header: NewHeader(0x32, MsgTypeVersionNotSupported, teid, seq, nil),
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

// Marshal returns the byte sequence generated from a VersionNotSupported.
func (v *VersionNotSupported) Marshal() ([]byte, error) {
	b := make([]byte, v.MarshalLen())
	if err := v.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (v *VersionNotSupported) MarshalTo(b []byte) error {
	if len(b) < v.MarshalLen() {
		return ErrTooShortToMarshal
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

// ParseVersionNotSupported decodes a given byte sequence as a VersionNotSupported.
func ParseVersionNotSupported(b []byte) (*VersionNotSupported, error) {
	v := &VersionNotSupported{}
	if err := v.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return v, nil
}

// UnmarshalBinary decodes a given byte sequence as a VersionNotSupported.
func (v *VersionNotSupported) UnmarshalBinary(b []byte) error {
	var err error
	v.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(v.Header.Payload) < 2 {
		return nil
	}

	ie, err := ie.ParseMultiIEs(v.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		v.AdditionalIEs = append(v.AdditionalIEs, i)
	}
	return nil
}

// MarshalLen returns the serial length of Data.
func (v *VersionNotSupported) MarshalLen() int {
	l := v.Header.MarshalLen() - len(v.Header.Payload)

	for _, ie := range v.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.MarshalLen()
	}
	return l
}

// SetLength sets the length in Length field.
func (v *VersionNotSupported) SetLength() {
	v.Length = uint16(v.MarshalLen() - 8)
}

// MessageTypeName returns the name of protocol.
func (v *VersionNotSupported) MessageTypeName() string {
	return "Version Not Supported"
}

// TEID returns the TEID in human-readable string.
func (v *VersionNotSupported) TEID() uint32 {
	return v.Header.TEID
}
