// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// DownlinkDataNotificationFailureIndication is a DownlinkDataNotificationFailureIndication Header and its IEs above.
type DownlinkDataNotificationFailureIndication struct {
	*Header
	Cause            *ie.IE
	OriginatingNode  *ie.IE
	IMSI             *ie.IE
	PrivateExtension *ie.IE
	AdditionalIEs    []*ie.IE
}

// NewDownlinkDataNotificationFailureIndication creates a new DownlinkDataNotificationFailureIndication.
func NewDownlinkDataNotificationFailureIndication(teid, seq uint32, ies ...*ie.IE) *DownlinkDataNotificationFailureIndication {
	d := &DownlinkDataNotificationFailureIndication{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeDownlinkDataNotificationFailureIndication, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			d.Cause = i
		case ie.NodeType:
			d.OriginatingNode = i
		case ie.IMSI:
			d.IMSI = i
		case ie.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	d.SetLength()
	return d
}

// Marshal serializes DownlinkDataNotificationFailureIndication into bytes.
func (d *DownlinkDataNotificationFailureIndication) Marshal() ([]byte, error) {
	b := make([]byte, d.MarshalLen())
	if err := d.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes DownlinkDataNotificationFailureIndication into bytes.
func (d *DownlinkDataNotificationFailureIndication) MarshalTo(b []byte) error {
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
	if ie := d.OriginatingNode; ie != nil {
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

// ParseDownlinkDataNotificationFailureIndication decodes given bytes as DownlinkDataNotificationFailureIndication.
func ParseDownlinkDataNotificationFailureIndication(b []byte) (*DownlinkDataNotificationFailureIndication, error) {
	d := &DownlinkDataNotificationFailureIndication{}
	if err := d.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return d, nil
}

// UnmarshalBinary decodes given bytes as DownlinkDataNotificationFailureIndication.
func (d *DownlinkDataNotificationFailureIndication) UnmarshalBinary(b []byte) error {
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
		case ie.NodeType:
			d.OriginatingNode = i
		case ie.IMSI:
			d.IMSI = i
		case ie.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (d *DownlinkDataNotificationFailureIndication) MarshalLen() int {
	l := d.Header.MarshalLen() - len(d.Header.Payload)

	if ie := d.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.OriginatingNode; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.IMSI; ie != nil {
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
func (d *DownlinkDataNotificationFailureIndication) SetLength() {
	d.Header.Length = uint16(d.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (d *DownlinkDataNotificationFailureIndication) MessageTypeName() string {
	return "Downlink Data Notification Failure Indication"
}

// TEID returns the TEID in uint32.
func (d *DownlinkDataNotificationFailureIndication) TEID() uint32 {
	return d.Header.teid()
}
