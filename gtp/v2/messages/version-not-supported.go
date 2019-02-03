// Copyright 2019 go-gtp authors. All rights reservev.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/gtp/v2/ies"
)

// VersionNotSupportedIndication is a VersionNotSupportedIndication Header and its IEs above.
type VersionNotSupportedIndication struct {
	*Header
	AdditionalIEs []*ies.IE
}

// NewVersionNotSupportedIndication creates a new VersionNotSupportedIndication.
func NewVersionNotSupportedIndication(teid, seq uint32, ie ...*ies.IE) *VersionNotSupportedIndication {
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

// Serialize serializes VersionNotSupportedIndication into bytes.
func (v *VersionNotSupportedIndication) Serialize() ([]byte, error) {
	b := make([]byte, v.Len())
	if err := v.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes VersionNotSupportedIndication into bytes.
func (v *VersionNotSupportedIndication) SerializeTo(b []byte) error {
	if v.Header.Payload != nil {
		v.Header.Payload = nil
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

// DecodeVersionNotSupportedIndication decodes given bytes as VersionNotSupportedIndication.
func DecodeVersionNotSupportedIndication(b []byte) (*VersionNotSupportedIndication, error) {
	v := &VersionNotSupportedIndication{}
	if err := v.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return v, nil
}

// DecodeFromBytes decodes given bytes as VersionNotSupportedIndication.
func (v *VersionNotSupportedIndication) DecodeFromBytes(b []byte) error {
	var err error
	v.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}
	if len(v.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ies.DecodeMultiIEs(v.Header.Payload)
	if err != nil {
		if err == ErrTooShortToDecode {
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

// Len returns the actual length in int.
func (v *VersionNotSupportedIndication) Len() int {
	l := v.Header.Len() - len(v.Header.Payload)
	for _, ie := range v.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.Len()
	}
	return l
}

// SetLength sets the length in Length fielv.
func (v *VersionNotSupportedIndication) SetLength() {
	v.Header.Length = uint16(v.Len() - 4)
}

// MessageTypeName returns the name of protocol.
func (v *VersionNotSupportedIndication) MessageTypeName() string {
	return "Version Not Supported Indication"
}

// TEID returns the TEID in uint32.
func (v *VersionNotSupportedIndication) TEID() uint32 {
	return v.Header.teid()
}
