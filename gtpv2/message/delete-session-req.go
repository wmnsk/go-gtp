// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// DeleteSessionRequest is a DeleteSessionRequest Header and its IEs above.
type DeleteSessionRequest struct {
	*Header
	Cause                             *ie.IE
	LinkedEBI                         *ie.IE
	ULI                               *ie.IE
	IndicationFlags                   *ie.IE
	PCO                               *ie.IE
	OriginatingNode                   *ie.IE
	SenderFTEIDC                      *ie.IE
	UETimeZone                        *ie.IE
	ULITimestamp                      *ie.IE
	RANNASReleaseCause                *ie.IE
	TWANIdentifier                    *ie.IE
	TWANIdentifierTimestamp           *ie.IE
	MMESGSNOverloadControlInformation *ie.IE
	SGWOverloadControlInformaion      *ie.IE
	TWANePDGOverloadControlInformaion *ie.IE
	WLANLocationInformation           *ie.IE
	WLANLocationTimeStamp             *ie.IE
	UELocalIPAddress                  *ie.IE
	UEUDPPort                         *ie.IE
	EPCO                              *ie.IE
	UETCPPort                         *ie.IE
	SecondaryRATUsageDataReport       *ie.IE
	PrivateExtension                  *ie.IE
	AdditionalIEs                     []*ie.IE
}

// NewDeleteSessionRequest creates a new DeleteSessionRequest.
func NewDeleteSessionRequest(teid, seq uint32, ies ...*ie.IE) *DeleteSessionRequest {
	d := &DeleteSessionRequest{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeDeleteSessionRequest, teid, seq, nil,
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
		case ie.UserLocationInformation:
			d.ULI = i
		case ie.Indication:
			d.IndicationFlags = i
		case ie.ProtocolConfigurationOptions:
			d.PCO = i
		case ie.NodeType:
			d.OriginatingNode = i
		case ie.FullyQualifiedTEID:
			d.SenderFTEIDC = i
		case ie.UETimeZone:
			d.UETimeZone = i
		case ie.ULITimestamp:
			d.ULITimestamp = i
		case ie.RANNASCause:
			d.RANNASReleaseCause = i
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
				d.WLANLocationTimeStamp = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				d.MMESGSNOverloadControlInformation = i
			case 1:
				d.SGWOverloadControlInformaion = i
			case 2:
				d.TWANePDGOverloadControlInformaion = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.IPAddress:
			d.UELocalIPAddress = i
		case ie.PortNumber:
			switch i.Instance() {
			case 0:
				d.UEUDPPort = i
			case 1:
				d.UETCPPort = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.ExtendedProtocolConfigurationOptions:
			d.EPCO = i
		case ie.SecondaryRATUsageDataReport:
			d.SecondaryRATUsageDataReport = i
		case ie.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	d.SetLength()
	return d
}

// Marshal serializes DeleteSessionRequest into bytes.
func (d *DeleteSessionRequest) Marshal() ([]byte, error) {
	b := make([]byte, d.MarshalLen())
	if err := d.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes DeleteSessionRequest into bytes.
func (d *DeleteSessionRequest) MarshalTo(b []byte) error {
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
	if ie := d.ULI; ie != nil {
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
	if ie := d.PCO; ie != nil {
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
	if ie := d.SenderFTEIDC; ie != nil {
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
	if ie := d.ULITimestamp; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.RANNASReleaseCause; ie != nil {
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
	if ie := d.MMESGSNOverloadControlInformation; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.SGWOverloadControlInformaion; ie != nil {
		if err := ie.MarshalTo(d.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := d.TWANePDGOverloadControlInformaion; ie != nil {
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
	if ie := d.WLANLocationTimeStamp; ie != nil {
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
	if ie := d.EPCO; ie != nil {
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

// ParseDeleteSessionRequest decodes given bytes as DeleteSessionRequest.
func ParseDeleteSessionRequest(b []byte) (*DeleteSessionRequest, error) {
	d := &DeleteSessionRequest{}
	if err := d.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return d, nil
}

// UnmarshalBinary decodes given bytes as DeleteSessionRequest.
func (d *DeleteSessionRequest) UnmarshalBinary(b []byte) error {
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
		case ie.UserLocationInformation:
			d.ULI = i
		case ie.Indication:
			d.IndicationFlags = i
		case ie.ProtocolConfigurationOptions:
			d.PCO = i
		case ie.NodeType:
			d.OriginatingNode = i
		case ie.FullyQualifiedTEID:
			d.SenderFTEIDC = i
		case ie.UETimeZone:
			d.UETimeZone = i
		case ie.ULITimestamp:
			d.ULITimestamp = i
		case ie.RANNASCause:
			d.RANNASReleaseCause = i
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
				d.WLANLocationTimeStamp = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				d.MMESGSNOverloadControlInformation = i
			case 1:
				d.SGWOverloadControlInformaion = i
			case 2:
				d.TWANePDGOverloadControlInformaion = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.IPAddress:
			d.UELocalIPAddress = i
		case ie.PortNumber:
			switch i.Instance() {
			case 0:
				d.UEUDPPort = i
			case 1:
				d.UETCPPort = i
			default:
				d.AdditionalIEs = append(d.AdditionalIEs, i)
			}
		case ie.ExtendedProtocolConfigurationOptions:
			d.EPCO = i
		case ie.SecondaryRATUsageDataReport:
			d.SecondaryRATUsageDataReport = i
		case ie.PrivateExtension:
			d.PrivateExtension = i
		default:
			d.AdditionalIEs = append(d.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (d *DeleteSessionRequest) MarshalLen() int {
	l := d.Header.MarshalLen() - len(d.Header.Payload)

	if ie := d.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.LinkedEBI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.ULI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.IndicationFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.PCO; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.OriginatingNode; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.SenderFTEIDC; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.UETimeZone; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.ULITimestamp; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.RANNASReleaseCause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.TWANIdentifier; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.TWANIdentifierTimestamp; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.MMESGSNOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.SGWOverloadControlInformaion; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.TWANePDGOverloadControlInformaion; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.WLANLocationInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.WLANLocationTimeStamp; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.UELocalIPAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.UEUDPPort; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := d.EPCO; ie != nil {
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
func (d *DeleteSessionRequest) SetLength() {
	d.Header.Length = uint16(d.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (d *DeleteSessionRequest) MessageTypeName() string {
	return "Delete Session Request"
}

// TEID returns the TEID in uint32.
func (d *DeleteSessionRequest) TEID() uint32 {
	return d.Header.teid()
}
