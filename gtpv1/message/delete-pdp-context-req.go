// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv1/ie"
)

// DeletePDPContextRequest is a DeletePDPContextRequest Header and its IEs above.
type DeletePDPContextRequest struct {
	*Header
	Cause               *ie.IE
	TeardownInd         *ie.IE
	NSAPI               *ie.IE
	PCO                 *ie.IE
	ULI                 *ie.IE
	MSTimeZone          *ie.IE
	ExtendedCommonFlags *ie.IE
	ULITimestamp        *ie.IE
	PrivateExtension    *ie.IE
	AdditionalIEs       []*ie.IE
}

// NewDeletePDPContextRequest creates a new GTPv1 DeletePDPContextRequest.
func NewDeletePDPContextRequest(teid uint32, seq uint16, ies ...*ie.IE) *DeletePDPContextRequest {
	d := &DeletePDPContextRequest{
		Header: NewHeader(0x32, MsgTypeDeletePDPContextRequest, teid, seq, nil),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			d.Cause = i
		case ie.TeardownInd:
			d.TeardownInd = i
		case ie.NSAPI:
			d.NSAPI = i
		case ie.ProtocolConfigurationOptions:
			d.PCO = i
		case ie.UserLocationInformation:
			d.ULI = i
		case ie.MSTimeZone:
			d.MSTimeZone = i
		case ie.ExtendedCommonFlags:
			d.ExtendedCommonFlags = i
		case ie.ULITimestamp:
			d.ULITimestamp = i
		case ie.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	d.SetLength()
	return d
}

// Marshal returns the byte sequence generated from a DeletePDPContextRequest.
func (d *DeletePDPContextRequest) Marshal() ([]byte, error) {
	b := make([]byte, d.MarshalLen())
	if err := d.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (d *DeletePDPContextRequest) MarshalTo(b []byte) error {
	if len(b) < d.MarshalLen() {
		return ErrTooShortToMarshal
	}
	d.Header.Payload = make([]byte, d.MarshalLen()-d.Header.MarshalLen())

	offset := 0
	if ie := d.Cause; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.TeardownInd; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.NSAPI; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.PCO; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.ULI; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.MSTimeZone; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.ExtendedCommonFlags; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.ULITimestamp; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.PrivateExtension; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	for _, ie := range d.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.MarshalTo(d.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	d.Header.SetLength()
	return d.Header.MarshalTo(b)
}

// ParseDeletePDPContextRequest decodes a given byte sequence as a DeletePDPContextRequest.
func ParseDeletePDPContextRequest(b []byte) (*DeletePDPContextRequest, error) {
	d := &DeletePDPContextRequest{}
	if err := d.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return d, nil
}

// UnmarshalBinary decodes a given byte sequence as a DeletePDPContextRequest.
func (d *DeletePDPContextRequest) UnmarshalBinary(b []byte) error {
	var err error
	d.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(d.Header.Payload) < 2 {
		return nil
	}

	ies, err := ie.ParseMultiIEs(d.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			d.Cause = i
		case ie.TeardownInd:
			d.TeardownInd = i
		case ie.NSAPI:
			d.NSAPI = i
		case ie.ProtocolConfigurationOptions:
			d.PCO = i
		case ie.UserLocationInformation:
			d.ULI = i
		case ie.MSTimeZone:
			d.MSTimeZone = i
		case ie.ExtendedCommonFlags:
			d.ExtendedCommonFlags = i
		case ie.ULITimestamp:
			d.ULITimestamp = i
		case ie.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}
	return nil
}

// MarshalLen returns the serial length of Data.
func (d *DeletePDPContextRequest) MarshalLen() int {
	l := d.Header.MarshalLen() - len(d.Header.Payload)

	if ie := d.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.TeardownInd; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.NSAPI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.PCO; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.ULI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.MSTimeZone; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.ExtendedCommonFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.ULITimestamp; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.PrivateExtension; ie != nil {
		l += ie.MarshalLen()
	}

	for _, ie := range d.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.MarshalLen()
	}
	return l
}

// SetLength sets the length in Length field.
func (d *DeletePDPContextRequest) SetLength() {
	d.Length = uint16(d.MarshalLen() - 8)
}

// MessageTypeName returns the name of protocol.
func (d *DeletePDPContextRequest) MessageTypeName() string {
	return "Delete PDP Context Request"
}

// TEID returns the TEID in human-readable string.
func (d *DeletePDPContextRequest) TEID() uint32 {
	return d.Header.TEID
}
