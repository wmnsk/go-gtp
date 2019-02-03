// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/gtp/v2/ies"
)

// DeleteSessionResponse is a DeleteSessionResponse Header and its IEs above.
type DeleteSessionResponse struct {
	*Header
	Cause                         *ies.IE
	Recovery                      *ies.IE
	PCO                           *ies.IE
	IndicationFlags               *ies.IE
	PGWNodeLoadControlInformation *ies.IE
	PGWAPNLoadControlInformation  *ies.IE
	SGWNodeLoadControlInformation *ies.IE
	PGWOverloadControlInformation *ies.IE
	SGWOverloadControlInformation *ies.IE
	EPCO                          *ies.IE
	APNRateControlStatus          *ies.IE
	PrivateExtension              *ies.IE
	AdditionalIEs                 []*ies.IE
}

// NewDeleteSessionResponse creates a new DeleteSessionResponse.
func NewDeleteSessionResponse(teid, seq uint32, ie ...*ies.IE) *DeleteSessionResponse {
	d := &DeleteSessionResponse{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeDeleteSessionResponse, teid, seq, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			d.Cause = i
		case ies.Recovery:
			d.Recovery = i
		case ies.ProtocolConfigurationOptions:
			d.PCO = i
		case ies.Indication:
			d.IndicationFlags = i
		case ies.LoadControlInformation:
			switch i.Instance() {
			case 1:
				d.PGWNodeLoadControlInformation = i
			case 2:
				d.PGWAPNLoadControlInformation = i
			case 3:
				d.SGWNodeLoadControlInformation = i
			}
		case ies.OverloadControlInformation:
			switch i.Instance() {
			case 1:
				d.PGWOverloadControlInformation = i
			case 2:
				d.SGWOverloadControlInformation = i
			}
		case ies.ExtendedProtocolConfigurationOptions:
			d.EPCO = i
		case ies.APNRateControlStatus:
			d.APNRateControlStatus = i
		case ies.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	d.SetLength()
	return d
}

// Serialize serializes DeleteSessionResponse into bytes.
func (d *DeleteSessionResponse) Serialize() ([]byte, error) {
	b := make([]byte, d.Len())
	if err := d.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes DeleteSessionResponse into bytes.
func (d *DeleteSessionResponse) SerializeTo(b []byte) error {
	if d.Header.Payload != nil {
		d.Header.Payload = nil
	}
	d.Header.Payload = make([]byte, d.Len()-d.Header.Len())

	offset := 0
	if ie := d.Cause; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.Recovery; ie != nil {
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
	if ie := d.IndicationFlags; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.PGWNodeLoadControlInformation; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.PGWAPNLoadControlInformation; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.SGWNodeLoadControlInformation; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.PGWOverloadControlInformation; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.SGWOverloadControlInformation; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.EPCO; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.APNRateControlStatus; ie != nil {
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

// DecodeDeleteSessionResponse decodes given bytes as DeleteSessionResponse.
func DecodeDeleteSessionResponse(b []byte) (*DeleteSessionResponse, error) {
	d := &DeleteSessionResponse{}
	if err := d.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return d, nil
}

// DecodeFromBytes decodes given bytes as DeleteSessionResponse.
func (d *DeleteSessionResponse) DecodeFromBytes(b []byte) error {
	var err error
	d.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}
	if len(d.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ies.DecodeMultiIEs(d.Header.Payload)
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
		case ies.Recovery:
			d.Recovery = i
		case ies.ProtocolConfigurationOptions:
			d.PCO = i
		case ies.Indication:
			d.IndicationFlags = i
		case ies.LoadControlInformation:
			switch i.Instance() {
			case 1:
				d.PGWNodeLoadControlInformation = i
			case 2:
				d.PGWAPNLoadControlInformation = i
			case 3:
				d.SGWNodeLoadControlInformation = i
			}
		case ies.OverloadControlInformation:
			switch i.Instance() {
			case 1:
				d.PGWOverloadControlInformation = i
			case 2:
				d.SGWOverloadControlInformation = i
			}
		case ies.ExtendedProtocolConfigurationOptions:
			d.EPCO = i
		case ies.APNRateControlStatus:
			d.APNRateControlStatus = i
		case ies.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	return nil
}

// Len returns the actual length in int.
func (d *DeleteSessionResponse) Len() int {
	l := d.Header.Len() - len(d.Header.Payload)

	if ie := d.Cause; ie != nil {
		l += ie.Len()
	}
	if ie := d.Recovery; ie != nil {
		l += ie.Len()
	}
	if ie := d.PCO; ie != nil {
		l += ie.Len()
	}
	if ie := d.IndicationFlags; ie != nil {
		l += ie.Len()
	}
	if ie := d.PGWNodeLoadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := d.PGWAPNLoadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := d.SGWNodeLoadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := d.PGWOverloadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := d.SGWOverloadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := d.EPCO; ie != nil {
		l += ie.Len()
	}
	if ie := d.APNRateControlStatus; ie != nil {
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
func (d *DeleteSessionResponse) SetLength() {
	d.Header.Length = uint16(d.Len() - 4)
}

// MessageTypeName returns the name of protocol.
func (d *DeleteSessionResponse) MessageTypeName() string {
	return "Delete Session Response"
}

// TEID returns the TEID in uint32.
func (d *DeleteSessionResponse) TEID() uint32 {
	return d.Header.teid()
}
