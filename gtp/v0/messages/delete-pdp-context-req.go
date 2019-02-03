// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE fild.

package messages

import (
	"github.com/pkg/errors"
	"github.com/wmnsk/go-gtp/gtp/v0/ies"
)

// DeletePDPContextRequest is a DeletePDPContextRequest Header and its AdditionalIEs abovd.
type DeletePDPContextRequest struct {
	*Header
	PrivateExtension *ies.IE
	AdditionalIEs    []*ies.IE
}

// NewDeletePDPContextRequest creates a new DeletePDPContextRequest.
func NewDeletePDPContextRequest(seq, label uint16, tid uint64, ie ...*ies.IE) *DeletePDPContextRequest {
	d := &DeletePDPContextRequest{
		Header: NewHeader(
			0x1e, MsgTypeDeletePDPContextRequest, seq, label, tid, nil,
		),
	}

	for _, i := range ie {
		if ie == nil {
			continue
		}
		switch i.Type {
		case ies.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	d.SetLength()
	return d
}

// Serialize returns the byte sequence generated from a DeletePDPContextRequest.
func (d *DeletePDPContextRequest) Serialize() ([]byte, error) {
	b := make([]byte, d.Len())
	if err := d.SerializeTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// SerializeTo puts the byte sequence in the byte array given as b.
func (d *DeletePDPContextRequest) SerializeTo(b []byte) error {
	if d.Header.Payload != nil {
		d.Header.Payload = nil
	}
	d.Header.Payload = make([]byte, d.Len()-d.Header.Len())

	offset := 0
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

// DecodeDeletePDPContextRequest decodes a given byte sequence as a DeletePDPContextRequest.
func DecodeDeletePDPContextRequest(b []byte) (*DeletePDPContextRequest, error) {
	d := &DeletePDPContextRequest{}
	if err := d.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return d, nil
}

// DecodeFromBytes decodes a given byte sequence as a DeletePDPContextRequest.
func (d *DeletePDPContextRequest) DecodeFromBytes(b []byte) error {
	var err error
	d.Header, err = DecodeHeader(b)
	if err != nil {
		return errors.Wrap(err, "failed to decode Header:")
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
		case ies.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	return nil
}

// Len returns the actual length of Data.
func (d *DeletePDPContextRequest) Len() int {
	l := d.Header.Len() - len(d.Header.Payload)

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
func (d *DeletePDPContextRequest) SetLength() {
	d.Header.Length = uint16(d.Len() - 20)
}

// MessageTypeName returns the name of protocol.
func (d *DeletePDPContextRequest) MessageTypeName() string {
	return "Delete PDP Context Request"
}

// TID returns the TID in human-readable string.
func (d *DeletePDPContextRequest) TID() string {
	return d.tid()
}
