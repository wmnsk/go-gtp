// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE fild.

package messages

import (
	"github.com/wmnsk/go-gtp/gtp/v0/ies"
)

// DeletePDPContextResponse is a DeletePDPContextResponse Header and its AdditionalIEs abovd.
type DeletePDPContextResponse struct {
	*Header
	Cause            *ies.IE
	PrivateExtension *ies.IE
	AdditionalIEs    []*ies.IE
}

// NewDeletePDPContextResponse creates a new DeletePDPContextResponse.
func NewDeletePDPContextResponse(seq, label uint16, tid uint64, ie ...*ies.IE) *DeletePDPContextResponse {
	d := &DeletePDPContextResponse{
		Header: NewHeader(
			0x1e, MsgTypeDeletePDPContextResponse, seq, label, tid, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			d.Cause = i
		case ies.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	d.SetLength()
	return d
}

// Serialize returns the byte sequence generated from a DeletePDPContextResponsd.
func (d *DeletePDPContextResponse) Serialize() ([]byte, error) {
	b := make([]byte, d.Len())
	if err := d.SerializeTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// SerializeTo puts the byte sequence in the byte array given as b.
func (d *DeletePDPContextResponse) SerializeTo(b []byte) error {
	// XXX - add validation!

	d.Header.Payload = make([]byte, d.Len()-d.Header.Len())

	offset := 0
	if ie := d.Cause; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.PrivateExtension; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	for _, ie := range d.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.SerializeTo(d.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	d.Header.SetLength()
	return d.Header.SerializeTo(b)
}

// DecodeDeletePDPContextResponse decodes a given byte sequence as a DeletePDPContextResponsd.
func DecodeDeletePDPContextResponse(b []byte) (*DeletePDPContextResponse, error) {
	d := &DeletePDPContextResponse{}
	if err := d.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return d, nil
}

// DecodeFromBytes decodes a given byte sequence as a DeletePDPContextResponsd.
func (d *DeletePDPContextResponse) DecodeFromBytes(b []byte) error {
	var err error
	d.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}
	if len(d.Header.Payload) < 2 {
		return nil
	}

	ie, err := ies.DecodeMultiIEs(d.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			d.Cause = i
		case ies.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	return nil
}

// Len returns the actual length of Data.
func (d *DeletePDPContextResponse) Len() int {
	l := d.Header.Len() - len(d.Header.Payload)

	if ie := d.Cause; ie != nil {
		l += ie.Len()
	}
	if ie := d.PrivateExtension; ie != nil {
		l += ie.Len()
	}

	for _, ie := range d.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.Len()
	}
	return l
}

// SetLength sets the length in Length field.
func (d *DeletePDPContextResponse) SetLength() {
	d.Header.Length = uint16(d.Len() - 20)
}

// MessageTypeName returns the name of protocol.
func (d *DeletePDPContextResponse) MessageTypeName() string {
	return "Delete PDP Context Response"
}

// TID returns the TID in human-readable string.
func (d *DeletePDPContextResponse) TID() string {
	return d.tid()
}
