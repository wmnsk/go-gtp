// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/v2/ies"
)

// DeleteBearerRequest is a DeleteBearerRequest Header and its IEs above.
type DeleteBearerRequest struct {
	*Header
	LinkedEBI                     *ies.IE
	EBI                           *ies.IE
	FailedBearerContext           *ies.IE
	PTI                           *ies.IE
	PCO                           *ies.IE
	PGWFQCSID                     *ies.IE
	SGWFQCSID                     *ies.IE
	Cause                         *ies.IE
	IndicationFlags               *ies.IE
	PGWNodeLoadControlInformation *ies.IE
	PGWAPNLoadControlInformation  *ies.IE
	SGWNodeLoadControlInformation *ies.IE
	PGWOverloadControlInformation *ies.IE
	SGWOverloadControlInformation *ies.IE
	NBIFOMContainer               *ies.IE
	APNRateControlStatus          *ies.IE
	EPCO                          *ies.IE
	PrivateExtension              *ies.IE
	AdditionalIEs                 []*ies.IE
}

// NewDeleteBearerRequest creates a new DeleteBearerRequest.
func NewDeleteBearerRequest(teid, seq uint32, ie ...*ies.IE) *DeleteBearerRequest {
	d := &DeleteBearerRequest{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeDeleteBearerRequest, teid, seq, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.EPSBearerID:
			switch i.Instance() {
			case 0:
				d.LinkedEBI = i
			case 1:
				d.EBI = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ies.BearerContext:
			d.FailedBearerContext = i
		case ies.ProcedureTransactionID:
			d.PTI = i
		case ies.ProtocolConfigurationOptions:
			d.PCO = i
		case ies.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				d.PGWFQCSID = i
			case 1:
				d.SGWFQCSID = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ies.Cause:
			d.Cause = i
		case ies.Indication:
			d.IndicationFlags = i
		case ies.LoadControlInformation:
			switch i.Instance() {
			case 0:
				d.PGWNodeLoadControlInformation = i
			case 1:
				d.PGWAPNLoadControlInformation = i
			case 2:
				d.SGWNodeLoadControlInformation = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ies.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				d.PGWOverloadControlInformation = i
			case 1:
				d.SGWOverloadControlInformation = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ies.FContainer:
			d.NBIFOMContainer = i
		case ies.APNRateControlStatus:
			d.APNRateControlStatus = i
		case ies.ExtendedProtocolConfigurationOptions:
			d.EPCO = i
		case ies.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	d.SetLength()
	return d
}

// Serialize serializes DeleteBearerRequest into bytes.
func (d *DeleteBearerRequest) Serialize() ([]byte, error) {
	b := make([]byte, d.Len())
	if err := d.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes DeleteBearerRequest into bytes.
func (d *DeleteBearerRequest) SerializeTo(b []byte) error {
	if d.Header.Payload != nil {
		d.Header.Payload = nil
	}
	d.Header.Payload = make([]byte, d.Len()-d.Header.Len())

	offset := 0

	if ie := d.LinkedEBI; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.EBI; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.FailedBearerContext; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.PTI; ie != nil {
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
	if ie := d.PGWFQCSID; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.SGWFQCSID; ie != nil {
		if err := ie.SerializeTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := d.Cause; ie != nil {
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
	if ie := d.NBIFOMContainer; ie != nil {
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
	if ie := d.EPCO; ie != nil {
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

// DecodeDeleteBearerRequest decodes given bytes as DeleteBearerRequest.
func DecodeDeleteBearerRequest(b []byte) (*DeleteBearerRequest, error) {
	d := &DeleteBearerRequest{}
	if err := d.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return d, nil
}

// DecodeFromBytes decodes given bytes as DeleteBearerRequest.
func (d *DeleteBearerRequest) DecodeFromBytes(b []byte) error {
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
		case ies.EPSBearerID:
			switch i.Instance() {
			case 0:
				d.LinkedEBI = i
			case 1:
				d.EBI = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ies.BearerContext:
			d.FailedBearerContext = i
		case ies.ProcedureTransactionID:
			d.PTI = i
		case ies.ProtocolConfigurationOptions:
			d.PCO = i
		case ies.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				d.PGWFQCSID = i
			case 1:
				d.SGWFQCSID = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ies.Cause:
			d.Cause = i
		case ies.Indication:
			d.IndicationFlags = i
		case ies.LoadControlInformation:
			switch i.Instance() {
			case 0:
				d.PGWNodeLoadControlInformation = i
			case 1:
				d.PGWAPNLoadControlInformation = i
			case 2:
				d.SGWNodeLoadControlInformation = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ies.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				d.PGWOverloadControlInformation = i
			case 1:
				d.SGWOverloadControlInformation = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ies.FContainer:
			d.NBIFOMContainer = i
		case ies.APNRateControlStatus:
			d.APNRateControlStatus = i
		case ies.ExtendedProtocolConfigurationOptions:
			d.EPCO = i
		case ies.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	return nil
}

// Len returns the actual length in int.
func (d *DeleteBearerRequest) Len() int {
	l := d.Header.Len() - len(d.Header.Payload)

	if ie := d.LinkedEBI; ie != nil {
		l += ie.Len()
	}
	if ie := d.EBI; ie != nil {
		l += ie.Len()
	}
	if ie := d.FailedBearerContext; ie != nil {
		l += ie.Len()
	}
	if ie := d.PTI; ie != nil {
		l += ie.Len()
	}
	if ie := d.PCO; ie != nil {
		l += ie.Len()
	}
	if ie := d.PGWFQCSID; ie != nil {
		l += ie.Len()
	}
	if ie := d.SGWFQCSID; ie != nil {
		l += ie.Len()
	}
	if ie := d.Cause; ie != nil {
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
	if ie := d.NBIFOMContainer; ie != nil {
		l += ie.Len()
	}
	if ie := d.APNRateControlStatus; ie != nil {
		l += ie.Len()
	}
	if ie := d.EPCO; ie != nil {
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
func (d *DeleteBearerRequest) SetLength() {
	d.Header.Length = uint16(d.Len() - 4)
}

// MessageTypeName returns the name of protocol.
func (d *DeleteBearerRequest) MessageTypeName() string {
	return "Delete Bearer Request"
}

// TEID returns the TEID in uint32.
func (d *DeleteBearerRequest) TEID() uint32 {
	return d.Header.teid()
}
