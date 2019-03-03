// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/v1/ies"
)

// VersionNotSupported is a VersionNotSupported Header and its IEs above.
type VersionNotSupported struct {
	*Header
	AdditionalIEs []*ies.IE
}

// NewVersionNotSupported creates a new GTPv1 VersionNotSupported.
func NewVersionNotSupported(teid uint32, seq uint16, ie ...*ies.IE) *VersionNotSupported {
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

// Serialize returns the byte sequence generated from a VersionNotSupported.
func (v *VersionNotSupported) Serialize() ([]byte, error) {
	b := make([]byte, v.Len())
	if err := v.SerializeTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// SerializeTo puts the byte sequence in the byte array given as b.
func (v *VersionNotSupported) SerializeTo(b []byte) error {
	if len(b) < v.Len() {
		return ErrTooShortToSerialize
	}
	v.Header.Payload = make([]byte, v.Len()-v.Header.Len())

	offset := 0
	for _, ie := range v.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.SerializeTo(v.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	v.Header.SetLength()
	return v.Header.SerializeTo(b)
}

// DecodeVersionNotSupported decodes a given byte sequence as a VersionNotSupported.
func DecodeVersionNotSupported(b []byte) (*VersionNotSupported, error) {
	v := &VersionNotSupported{}
	if err := v.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return v, nil
}

// DecodeFromBytes decodes a given byte sequence as a VersionNotSupported.
func (v *VersionNotSupported) DecodeFromBytes(b []byte) error {
	var err error
	v.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}
	if len(v.Header.Payload) < 2 {
		return nil
	}

	ie, err := ies.DecodeMultiIEs(v.Header.Payload)
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

// Len returns the actual length of Data.
func (v *VersionNotSupported) Len() int {
	l := v.Header.Len() - len(v.Header.Payload)

	for _, ie := range v.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.Len()
	}
	return l
}

// SetLength sets the length in Length field.
func (v *VersionNotSupported) SetLength() {
	v.Length = uint16(v.Len() - 8)
}

// MessageTypeName returns the name of protocol.
func (v *VersionNotSupported) MessageTypeName() string {
	return "Version Not Supported"
}

// TEID returns the TEID in human-readable string.
func (v *VersionNotSupported) TEID() uint32 {
	return v.Header.TEID
}
