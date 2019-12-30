// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/v2/ies"
)

// DeleteBearerResponse is a DeleteBearerResponse Header and its IEs above.
type DeleteBearerResponse struct {
	*Header
	Cause                              *ies.IE
	LinkedEBI                          *ies.IE
	BearerContexts                     *ies.IE
	Recovery                           *ies.IE
	MMEFQCSID                          *ies.IE
	SGWFQCSID                          *ies.IE
	EPDGFQCSID                         *ies.IE
	TWANFQCSID                         *ies.IE
	PCO                                *ies.IE
	UETimeZone                         *ies.IE
	ULI                                *ies.IE
	ULITimestamp                       *ies.IE
	TWANIdentifier                     *ies.IE
	TWANIdentifierTimestamp            *ies.IE
	MMEOverloadControlInformation      *ies.IE
	SGWOverloadControlInformation      *ies.IE
	MMESGSNIdentifier                  *ies.IE
	TWANePDGOverloadControlInformation *ies.IE
	WLANLocationInformation            *ies.IE
	WLANLocationTimestamp              *ies.IE
	UELocalIPAddress                   *ies.IE
	UEUDPPort                          *ies.IE
	NBIFOMContainer                    *ies.IE
	UETCPPort                          *ies.IE
	SecondaryRATUsageDataReport        *ies.IE
	PrivateExtension                   *ies.IE
	AdditionalIEs                      []*ies.IE
}

// NewDeleteBearerResponse creates a new DeleteBearerResponse.
func NewDeleteBearerResponse(teid, seq uint32, ie ...*ies.IE) *DeleteBearerResponse {
	d := &DeleteBearerResponse{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeDeleteBearerResponse, teid, seq, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			d.Cause = i
		case ies.EPSBearerID:
			d.LinkedEBI = i
		case ies.BearerContext:
			d.BearerContexts = i
		case ies.Recovery:
			d.Recovery = i
		case ies.FullyQualifiedCSID:
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
		case ies.ProtocolConfigurationOptions:
			d.PCO = i
		case ies.UETimeZone:
			d.UETimeZone = i
		case ies.UserLocationInformation:
			d.ULI = i
		case ies.ULITimestamp:
			d.ULITimestamp = i
		case ies.TWANIdentifier:
			switch i.Instance() {
			case 0:
				d.TWANIdentifier = i
			case 1:
				d.WLANLocationInformation = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ies.TWANIdentifierTimestamp:
			switch i.Instance() {
			case 0:
				d.TWANIdentifierTimestamp = i
			case 1:
				d.WLANLocationTimestamp = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ies.OverloadControlInformation:
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
		case ies.IPAddress:
			switch i.Instance() {
			case 0:
				d.MMESGSNIdentifier = i
			case 1:
				d.UELocalIPAddress = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ies.PortNumber:
			switch i.Instance() {
			case 0:
				d.UEUDPPort = i
			case 1:
				d.UETCPPort = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ies.FContainer:
			d.NBIFOMContainer = i
		case ies.SecondaryRATUsageDataReport:
			d.SecondaryRATUsageDataReport = i
		case ies.PrivateExtension:
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
	if ie := d.SecondaryRATUsageDataReport; ie != nil {
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
		case ies.EPSBearerID:
			d.LinkedEBI = i
		case ies.BearerContext:
			d.BearerContexts = i
		case ies.Recovery:
			d.Recovery = i
		case ies.FullyQualifiedCSID:
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
		case ies.ProtocolConfigurationOptions:
			d.PCO = i
		case ies.UETimeZone:
			d.UETimeZone = i
		case ies.UserLocationInformation:
			d.ULI = i
		case ies.ULITimestamp:
			d.ULITimestamp = i
		case ies.TWANIdentifier:
			switch i.Instance() {
			case 0:
				d.TWANIdentifier = i
			case 1:
				d.WLANLocationInformation = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ies.TWANIdentifierTimestamp:
			switch i.Instance() {
			case 0:
				d.TWANIdentifierTimestamp = i
			case 1:
				d.WLANLocationTimestamp = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ies.OverloadControlInformation:
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
		case ies.IPAddress:
			switch i.Instance() {
			case 0:
				d.MMESGSNIdentifier = i
			case 1:
				d.UELocalIPAddress = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ies.PortNumber:
			switch i.Instance() {
			case 0:
				d.UEUDPPort = i
			case 1:
				d.UETCPPort = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ies.FContainer:
			d.NBIFOMContainer = i
		case ies.SecondaryRATUsageDataReport:
			d.SecondaryRATUsageDataReport = i
		case ies.PrivateExtension:
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
	if ie := d.BearerContexts; ie != nil {
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
	if ie := d.SecondaryRATUsageDataReport; ie != nil {
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
