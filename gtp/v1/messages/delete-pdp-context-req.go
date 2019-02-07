// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/gtp/v1/ies"
)

// DeletePDPContextRequest is a DeletePDPContextRequest Header and its IEs above.
type DeletePDPContextRequest struct {
	*Header
	Cause               *ies.IE
	TeardownInd         *ies.IE
	NSAPI               *ies.IE
	PCO                 *ies.IE
	ULI                 *ies.IE
	MSTimeZone          *ies.IE
	ExtendedCommonFlags *ies.IE
	ULITimestamp        *ies.IE
	PrivateExtension    *ies.IE
	AdditionalIEs       []*ies.IE
}

// NewDeletePDPContextRequest creates a new GTPv1 DeletePDPContextRequest.
func NewDeletePDPContextRequest(teid uint32, seq uint16, ie ...*ies.IE) *DeletePDPContextRequest {
	d := &DeletePDPContextRequest{
		Header: NewHeader(0x32, MsgTypeDeletePDPContextRequest, teid, seq, nil),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			d.Cause = i
		case ies.TeardownInd:
			d.TeardownInd = i
		case ies.NSAPI:
			d.NSAPI = i
		case ies.ProtocolConfigurationOptions:
			d.PCO = i
		case ies.UserLocationInformation:
			d.ULI = i
		case ies.MSTimeZone:
			d.MSTimeZone = i
		case ies.ExtendedCommonFlags:
			d.ExtendedCommonFlags = i
		case ies.ULITimestamp:
			d.ULITimestamp = i
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
	if len(b) < d.Len() {
		return ErrTooShortToSerialize
	}
	d.Header.Payload = make([]byte, d.Len()-d.Header.Len())

	offset := 0
	if ie := d.Cause; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.TeardownInd; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.NSAPI; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.PCO; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.ULI; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.MSTimeZone; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.ExtendedCommonFlags; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.ULITimestamp; ie != nil {
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
		case ies.TeardownInd:
			d.TeardownInd = i
		case ies.NSAPI:
			d.NSAPI = i
		case ies.ProtocolConfigurationOptions:
			d.PCO = i
		case ies.UserLocationInformation:
			d.ULI = i
		case ies.MSTimeZone:
			d.MSTimeZone = i
		case ies.ExtendedCommonFlags:
			d.ExtendedCommonFlags = i
		case ies.ULITimestamp:
			d.ULITimestamp = i
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

	if ie := d.Cause; ie != nil {
		l += ie.Len()
	}
	if ie := d.TeardownInd; ie != nil {
		l += ie.Len()
	}
	if ie := d.NSAPI; ie != nil {
		l += ie.Len()
	}
	if ie := d.PCO; ie != nil {
		l += ie.Len()
	}
	if ie := d.ULI; ie != nil {
		l += ie.Len()
	}
	if ie := d.MSTimeZone; ie != nil {
		l += ie.Len()
	}
	if ie := d.ExtendedCommonFlags; ie != nil {
		l += ie.Len()
	}
	if ie := d.ULITimestamp; ie != nil {
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
func (d *DeletePDPContextRequest) SetLength() {
	d.Length = uint16(d.Len() - 8)
}

// MessageTypeName returns the name of protocol.
func (d *DeletePDPContextRequest) MessageTypeName() string {
	return "Delete PDP Context Request"
}

// TEID returns the TEID in human-readable string.
func (d *DeletePDPContextRequest) TEID() uint32 {
	return d.Header.TEID
}
