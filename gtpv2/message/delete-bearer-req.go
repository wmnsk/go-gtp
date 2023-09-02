// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// DeleteBearerRequest is a DeleteBearerRequest Header and its IEs above.
type DeleteBearerRequest struct {
	*Header
	LinkedEBI                     *ie.IE
	EBIs                          []*ie.IE
	FailedBearerContext           *ie.IE
	PTI                           *ie.IE
	PCO                           *ie.IE
	PGWFQCSID                     *ie.IE
	SGWFQCSID                     *ie.IE
	Cause                         *ie.IE
	IndicationFlags               *ie.IE
	PGWNodeLoadControlInformation *ie.IE
	PGWAPNLoadControlInformation  *ie.IE
	SGWNodeLoadControlInformation *ie.IE
	PGWOverloadControlInformation *ie.IE
	SGWOverloadControlInformation *ie.IE
	NBIFOMContainer               *ie.IE
	APNRateControlStatus          *ie.IE
	EPCO                          *ie.IE
	PrivateExtension              *ie.IE
	AdditionalIEs                 []*ie.IE
}

// NewDeleteBearerRequest creates a new DeleteBearerRequest.
func NewDeleteBearerRequest(teid, seq uint32, ies ...*ie.IE) *DeleteBearerRequest {
	d := &DeleteBearerRequest{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeDeleteBearerRequest, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.EPSBearerID:
			switch i.Instance() {
			case 0:
				d.LinkedEBI = i
			case 1:
				d.EBIs = append(d.EBIs, i)
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.BearerContext:
			d.FailedBearerContext = i
		case ie.ProcedureTransactionID:
			d.PTI = i
		case ie.ProtocolConfigurationOptions:
			d.PCO = i
		case ie.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				d.PGWFQCSID = i
			case 1:
				d.SGWFQCSID = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.Cause:
			d.Cause = i
		case ie.Indication:
			d.IndicationFlags = i
		case ie.LoadControlInformation:
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
		case ie.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				d.PGWOverloadControlInformation = i
			case 1:
				d.SGWOverloadControlInformation = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.FContainer:
			d.NBIFOMContainer = i
		case ie.APNRateControlStatus:
			d.APNRateControlStatus = i
		case ie.ExtendedProtocolConfigurationOptions:
			d.EPCO = i
		case ie.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	d.SetLength()
	return d
}

// Marshal serializes DeleteBearerRequest into bytes.
func (d *DeleteBearerRequest) Marshal() ([]byte, error) {
	b := make([]byte, d.MarshalLen())
	if err := d.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes DeleteBearerRequest into bytes.
func (d *DeleteBearerRequest) MarshalTo(b []byte) error {
	if d.Header.Payload != nil {
		d.Header.Payload = nil
	}
	d.Header.Payload = make([]byte, d.MarshalLen()-d.Header.MarshalLen())

	offset := 0

	if ie := d.LinkedEBI; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	for _, ie := range d.EBIs {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.FailedBearerContext; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.PTI; ie != nil {
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
	if ie := d.PGWFQCSID; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.SGWFQCSID; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.Cause; ie != nil {
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
	if ie := d.PGWNodeLoadControlInformation; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.PGWAPNLoadControlInformation; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.SGWNodeLoadControlInformation; ie != nil {
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
	if ie := d.NBIFOMContainer; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.APNRateControlStatus; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.EPCO; ie != nil {
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

// ParseDeleteBearerRequest decodes given bytes as DeleteBearerRequest.
func ParseDeleteBearerRequest(b []byte) (*DeleteBearerRequest, error) {
	d := &DeleteBearerRequest{}
	if err := d.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return d, nil
}

// UnmarshalBinary decodes given bytes as DeleteBearerRequest.
func (d *DeleteBearerRequest) UnmarshalBinary(b []byte) error {
	var err error
	d.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(d.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ie.ParseMultiIEs(d.Header.Payload)
	if err != nil {
		return err
	}
	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.EPSBearerID:
			switch i.Instance() {
			case 0:
				d.LinkedEBI = i
			case 1:
				d.EBIs = append(d.EBIs, i)
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.BearerContext:
			d.FailedBearerContext = i
		case ie.ProcedureTransactionID:
			d.PTI = i
		case ie.ProtocolConfigurationOptions:
			d.PCO = i
		case ie.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				d.PGWFQCSID = i
			case 1:
				d.SGWFQCSID = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.Cause:
			d.Cause = i
		case ie.Indication:
			d.IndicationFlags = i
		case ie.LoadControlInformation:
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
		case ie.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				d.PGWOverloadControlInformation = i
			case 1:
				d.SGWOverloadControlInformation = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.FContainer:
			d.NBIFOMContainer = i
		case ie.APNRateControlStatus:
			d.APNRateControlStatus = i
		case ie.ExtendedProtocolConfigurationOptions:
			d.EPCO = i
		case ie.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (d *DeleteBearerRequest) MarshalLen() int {
	l := d.Header.MarshalLen() - len(d.Header.Payload)

	if ie := d.LinkedEBI; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range d.EBIs {
		l += ie.MarshalLen()
	}
	if ie := d.FailedBearerContext; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.PTI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.PCO; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.PGWFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.SGWFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.IndicationFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.PGWNodeLoadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.PGWAPNLoadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.SGWNodeLoadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.PGWOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.SGWOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.NBIFOMContainer; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.APNRateControlStatus; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.EPCO; ie != nil {
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
func (d *DeleteBearerRequest) SetLength() {
	d.Header.Length = uint16(d.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (d *DeleteBearerRequest) MessageTypeName() string {
	return "Delete Bearer Request"
}

// TEID returns the TEID in uint32.
func (d *DeleteBearerRequest) TEID() uint32 {
	return d.Header.teid()
}
