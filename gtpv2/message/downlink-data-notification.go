// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// DownlinkDataNotification is a DownlinkDataNotification Header and its IEs above.
type DownlinkDataNotification struct {
	*Header
	Cause                         *ie.IE
	EPSBearerID                   *ie.IE
	AllocationRetensionPriority   *ie.IE
	IMSI                          *ie.IE
	SenderFTEIDC                  *ie.IE
	IndicationFlags               *ie.IE
	SGWNodeLoadControlInformation *ie.IE
	SGWOverloadControlInformation *ie.IE
	PagingAndServiceInformation   *ie.IE
	DLDataPacketsSize             *ie.IE
	PrivateExtension              *ie.IE
	AdditionalIEs                 []*ie.IE
}

// NewDownlinkDataNotification creates a new DownlinkDataNotification.
func NewDownlinkDataNotification(teid, seq uint32, ies ...*ie.IE) *DownlinkDataNotification {
	d := &DownlinkDataNotification{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeDownlinkDataNotification, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			d.Cause = i
		case ie.EPSBearerID:
			d.EPSBearerID = i
		case ie.AllocationRetensionPriority:
			d.AllocationRetensionPriority = i
		case ie.IMSI:
			d.IMSI = i
		case ie.FullyQualifiedTEID:
			d.SenderFTEIDC = i
		case ie.Indication:
			d.IndicationFlags = i
		case ie.LoadControlInformation:
			d.SGWNodeLoadControlInformation = i
		case ie.OverloadControlInformation:
			d.SGWOverloadControlInformation = i
		case ie.PagingAndServiceInformation:
			d.PagingAndServiceInformation = i
		case ie.IntegerNumber:
			d.DLDataPacketsSize = i
		case ie.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	d.SetLength()
	return d
}

// Marshal serializes DownlinkDataNotification into bytes.
func (d *DownlinkDataNotification) Marshal() ([]byte, error) {
	b := make([]byte, d.MarshalLen())
	if err := d.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes DownlinkDataNotification into bytes.
func (d *DownlinkDataNotification) MarshalTo(b []byte) error {
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
	if ie := d.EPSBearerID; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.AllocationRetensionPriority; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.IMSI; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.SenderFTEIDC; ie != nil {
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
	if ie := d.SGWNodeLoadControlInformation; ie != nil {
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
	if ie := d.PagingAndServiceInformation; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.DLDataPacketsSize; ie != nil {
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

// ParseDownlinkDataNotification decodes given bytes as DownlinkDataNotification.
func ParseDownlinkDataNotification(b []byte) (*DownlinkDataNotification, error) {
	d := &DownlinkDataNotification{}
	if err := d.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return d, nil
}

// UnmarshalBinary decodes given bytes as DownlinkDataNotification.
func (d *DownlinkDataNotification) UnmarshalBinary(b []byte) error {
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
		case ie.Cause:
			d.Cause = i
		case ie.EPSBearerID:
			d.EPSBearerID = i
		case ie.AllocationRetensionPriority:
			d.AllocationRetensionPriority = i
		case ie.IMSI:
			d.IMSI = i
		case ie.FullyQualifiedTEID:
			d.SenderFTEIDC = i
		case ie.Indication:
			d.IndicationFlags = i
		case ie.LoadControlInformation:
			d.SGWNodeLoadControlInformation = i
		case ie.OverloadControlInformation:
			d.SGWOverloadControlInformation = i
		case ie.PagingAndServiceInformation:
			d.PagingAndServiceInformation = i
		case ie.IntegerNumber:
			d.DLDataPacketsSize = i
		case ie.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (d *DownlinkDataNotification) MarshalLen() int {
	l := d.Header.MarshalLen() - len(d.Header.Payload)

	if ie := d.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.EPSBearerID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.AllocationRetensionPriority; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.IMSI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.SenderFTEIDC; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.IndicationFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.SGWNodeLoadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.SGWOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.PagingAndServiceInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.DLDataPacketsSize; ie != nil {
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
func (d *DownlinkDataNotification) SetLength() {
	d.Header.Length = uint16(d.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (d *DownlinkDataNotification) MessageTypeName() string {
	return "Downlink Data Notification"
}

// TEID returns the TEID in uint32.
func (d *DownlinkDataNotification) TEID() uint32 {
	return d.Header.teid()
}
