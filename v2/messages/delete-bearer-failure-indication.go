// Copyright 2019-2020 go-gtp authors. All rights reservev.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "github.com/wmnsk/go-gtp/v2/ies"

// DeleteBearerFailureIndication is a DeleteBearerFailureIndication Header and its IEs above.
type DeleteBearerFailureIndication struct {
	*Header
	Cause                         *ies.IE
	BearerContexts                *ies.IE
	Recovery                      *ies.IE
	IndicationFlags               *ies.IE
	PGWOverloadControlInformation *ies.IE
	SGWOverloadControlInformation *ies.IE
	PrivateExtension              *ies.IE
	AdditionalIEs                 []*ies.IE
}

// NewDeleteBearerFailureIndication creates a new DeleteBearerFailureIndication.
func NewDeleteBearerFailureIndication(teid, seq uint32, ie ...*ies.IE) *DeleteBearerFailureIndication {
	d := &DeleteBearerFailureIndication{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeDeleteBearerFailureIndication, teid, seq, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			d.Cause = i
		case ies.BearerContext:
			d.BearerContexts = i
		case ies.Recovery:
			d.Recovery = i
		case ies.Indication:
			d.IndicationFlags = i
		case ies.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				d.PGWOverloadControlInformation = i
			case 1:
				d.SGWOverloadControlInformation = i
			}
		case ies.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	d.SetLength()
	return d
}

// Marshal serializes DeleteBearerFailureIndication into bytes.
func (d *DeleteBearerFailureIndication) Marshal() ([]byte, error) {
	b := make([]byte, d.MarshalLen())
	if err := d.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes DeleteBearerFailureIndication into bytes.
func (d *DeleteBearerFailureIndication) MarshalTo(b []byte) error {
	if d.Header.Payload != nil {
		d.Header.Payload = nil
	}
	d.Header.Payload = make([]byte, d.MarshalLen()-d.Header.MarshalLen())

	offset := 0
	if ie := d.Cause; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.BearerContexts; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.Recovery; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.IndicationFlags; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.PGWOverloadControlInformation; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.SGWOverloadControlInformation; ie != nil {
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

// ParseDeleteBearerFailureIndication decodes given bytes as DeleteBearerFailureIndication.
func ParseDeleteBearerFailureIndication(b []byte) (*DeleteBearerFailureIndication, error) {
	d := &DeleteBearerFailureIndication{}
	if err := d.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return d, nil
}

// UnmarshalBinary decodes given bytes as DeleteBearerFailureIndication.
func (d *DeleteBearerFailureIndication) UnmarshalBinary(b []byte) error {
	var err error
	d.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(d.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ies.ParseMultiIEs(d.Header.Payload)
	if err != nil {
		return err
	}
	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			d.Cause = i
		case ies.BearerContext:
			d.BearerContexts = i
		case ies.Recovery:
			d.Recovery = i
		case ies.Indication:
			d.IndicationFlags = i
		case ies.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				d.PGWOverloadControlInformation = i
			case 1:
				d.SGWOverloadControlInformation = i
			}
		case ies.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (d *DeleteBearerFailureIndication) MarshalLen() int {
	l := d.Header.MarshalLen() - len(d.Header.Payload)
	if ie := d.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.BearerContexts; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.Recovery; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.IndicationFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.PGWOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.SGWOverloadControlInformation; ie != nil {
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
func (d *DeleteBearerFailureIndication) SetLength() {
	d.Header.Length = uint16(d.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (d *DeleteBearerFailureIndication) MessageTypeName() string {
	return "Delete Bearer Failure Indication"
}

// TEID returns the TEID in uint32.
func (d *DeleteBearerFailureIndication) TEID() uint32 {
	return d.Header.teid()
}
