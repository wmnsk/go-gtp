// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// DeleteBearerResponse is a DeleteBearerResponse Header and its IEs above.
type DeleteBearerResponse struct {
	*Header
	Cause                              *ie.IE
	LinkedEBI                          *ie.IE
	BearerContexts                     []*ie.IE
	Recovery                           *ie.IE
	MMEFQCSID                          *ie.IE
	SGWFQCSID                          *ie.IE
	EPDGFQCSID                         *ie.IE
	TWANFQCSID                         *ie.IE
	PCO                                *ie.IE
	UETimeZone                         *ie.IE
	ULI                                *ie.IE
	ULITimestamp                       *ie.IE
	TWANIdentifier                     *ie.IE
	TWANIdentifierTimestamp            *ie.IE
	MMEOverloadControlInformation      *ie.IE
	SGWOverloadControlInformation      *ie.IE
	MMESGSNIdentifier                  *ie.IE
	TWANePDGOverloadControlInformation *ie.IE
	WLANLocationInformation            *ie.IE
	WLANLocationTimestamp              *ie.IE
	UELocalIPAddress                   *ie.IE
	UEUDPPort                          *ie.IE
	NBIFOMContainer                    *ie.IE
	UETCPPort                          *ie.IE
	SecondaryRATUsageDataReport        []*ie.IE
	PrivateExtension                   *ie.IE
	AdditionalIEs                      []*ie.IE
}

// NewDeleteBearerResponse creates a new DeleteBearerResponse.
func NewDeleteBearerResponse(teid, seq uint32, ies ...*ie.IE) *DeleteBearerResponse {
	d := &DeleteBearerResponse{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeDeleteBearerResponse, teid, seq, nil,
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
			d.LinkedEBI = i
		case ie.BearerContext:
			d.BearerContexts = append(d.BearerContexts, i)
		case ie.Recovery:
			d.Recovery = i
		case ie.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				d.MMEFQCSID = i
			case 1:
				d.SGWFQCSID = i
			case 2:
				d.EPDGFQCSID = i
			case 3:
				d.TWANFQCSID = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.ProtocolConfigurationOptions:
			d.PCO = i
		case ie.UETimeZone:
			d.UETimeZone = i
		case ie.UserLocationInformation:
			d.ULI = i
		case ie.ULITimestamp:
			d.ULITimestamp = i
		case ie.TWANIdentifier:
			switch i.Instance() {
			case 0:
				d.TWANIdentifier = i
			case 1:
				d.WLANLocationInformation = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.TWANIdentifierTimestamp:
			switch i.Instance() {
			case 0:
				d.TWANIdentifierTimestamp = i
			case 1:
				d.WLANLocationTimestamp = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				d.MMEOverloadControlInformation = i
			case 1:
				d.SGWOverloadControlInformation = i
			case 2:
				d.TWANePDGOverloadControlInformation = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.IPAddress:
			switch i.Instance() {
			case 0:
				d.MMESGSNIdentifier = i
			case 1:
				d.UELocalIPAddress = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.PortNumber:
			switch i.Instance() {
			case 0:
				d.UEUDPPort = i
			case 1:
				d.UETCPPort = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.FContainer:
			d.NBIFOMContainer = i
		case ie.SecondaryRATUsageDataReport:
			d.SecondaryRATUsageDataReport = append(d.SecondaryRATUsageDataReport, i)
		case ie.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	d.SetLength()
	return d
}

// Marshal serializes DeleteBearerResponse into bytes.
func (d *DeleteBearerResponse) Marshal() ([]byte, error) {
	b := make([]byte, d.MarshalLen())
	if err := d.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes DeleteBearerResponse into bytes.
func (d *DeleteBearerResponse) MarshalTo(b []byte) error {
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
	if ie := d.LinkedEBI; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	for _, ie := range d.BearerContexts {
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
	if ie := d.MMEFQCSID; ie != nil {
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
	if ie := d.EPDGFQCSID; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.TWANFQCSID; ie != nil {
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
	if ie := d.UETimeZone; ie != nil {
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
	if ie := d.ULITimestamp; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.TWANIdentifier; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.TWANIdentifierTimestamp; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.MMEOverloadControlInformation; ie != nil {
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
	if ie := d.MMESGSNIdentifier; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.TWANePDGOverloadControlInformation; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.WLANLocationInformation; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.WLANLocationTimestamp; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.UELocalIPAddress; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.UEUDPPort; ie != nil {
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
	if ie := d.UETCPPort; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	for _, ie := range d.SecondaryRATUsageDataReport {
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

// ParseDeleteBearerResponse decodes given bytes as DeleteBearerResponse.
func ParseDeleteBearerResponse(b []byte) (*DeleteBearerResponse, error) {
	d := &DeleteBearerResponse{}
	if err := d.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return d, nil
}

// UnmarshalBinary decodes given bytes as DeleteBearerResponse.
func (d *DeleteBearerResponse) UnmarshalBinary(b []byte) error {
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
			d.LinkedEBI = i
		case ie.BearerContext:
			d.BearerContexts = append(d.BearerContexts, i)
		case ie.Recovery:
			d.Recovery = i
		case ie.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				d.MMEFQCSID = i
			case 1:
				d.SGWFQCSID = i
			case 2:
				d.EPDGFQCSID = i
			case 3:
				d.TWANFQCSID = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.ProtocolConfigurationOptions:
			d.PCO = i
		case ie.UETimeZone:
			d.UETimeZone = i
		case ie.UserLocationInformation:
			d.ULI = i
		case ie.ULITimestamp:
			d.ULITimestamp = i
		case ie.TWANIdentifier:
			switch i.Instance() {
			case 0:
				d.TWANIdentifier = i
			case 1:
				d.WLANLocationInformation = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.TWANIdentifierTimestamp:
			switch i.Instance() {
			case 0:
				d.TWANIdentifierTimestamp = i
			case 1:
				d.WLANLocationTimestamp = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				d.MMEOverloadControlInformation = i
			case 1:
				d.SGWOverloadControlInformation = i
			case 2:
				d.TWANePDGOverloadControlInformation = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.IPAddress:
			switch i.Instance() {
			case 0:
				d.MMESGSNIdentifier = i
			case 1:
				d.UELocalIPAddress = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.PortNumber:
			switch i.Instance() {
			case 0:
				d.UEUDPPort = i
			case 1:
				d.UETCPPort = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.FContainer:
			d.NBIFOMContainer = i
		case ie.SecondaryRATUsageDataReport:
			d.SecondaryRATUsageDataReport = append(d.SecondaryRATUsageDataReport, i)
		case ie.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (d *DeleteBearerResponse) MarshalLen() int {
	l := d.Header.MarshalLen() - len(d.Header.Payload)

	if ie := d.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.LinkedEBI; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range d.BearerContexts {
		l += ie.MarshalLen()
	}
	if ie := d.Recovery; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.MMEFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.SGWFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.EPDGFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.TWANFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.PCO; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.UETimeZone; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.ULI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.ULITimestamp; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.TWANIdentifier; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.TWANIdentifierTimestamp; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.MMEOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.SGWOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.MMESGSNIdentifier; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.TWANePDGOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.WLANLocationInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.WLANLocationTimestamp; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.UELocalIPAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.UEUDPPort; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.NBIFOMContainer; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.UETCPPort; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range d.SecondaryRATUsageDataReport {
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
func (d *DeleteBearerResponse) SetLength() {
	d.Header.Length = uint16(d.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (d *DeleteBearerResponse) MessageTypeName() string {
	return "Delete Bearer Response"
}

// TEID returns the TEID in uint32.
func (d *DeleteBearerResponse) TEID() uint32 {
	return d.Header.teid()
}
