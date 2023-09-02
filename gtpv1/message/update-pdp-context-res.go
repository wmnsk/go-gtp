// Copyright 2019-2023 go-gtp authors. All rights reserveu.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv1/ie"
)

// UpdatePDPContextResponse is a UpdatePDPContextResponse Header and its IEs above.
type UpdatePDPContextResponse struct {
	*Header
	Cause                         *ie.IE
	Recovery                      *ie.IE
	TEIDDataI                     *ie.IE
	TEIDCPlane                    *ie.IE
	ChargingID                    *ie.IE
	PCO                           *ie.IE
	GGSNAddressForCPlane          *ie.IE
	GGSNAddressForUserTraffic     *ie.IE
	AltGGSNAddressForCPlane       *ie.IE
	AltGGSNAddressForUserTraffic  *ie.IE
	QoSProfile                    *ie.IE
	ChargingGatewayAddress        *ie.IE
	AltChargingGatewayAddress     *ie.IE
	CommonFlags                   *ie.IE
	APNRestriction                *ie.IE
	BearerControlMode             *ie.IE
	MSInfoChangeReportingAction   *ie.IE
	EvolvedARPI                   *ie.IE
	CSGInformationReportingAction *ie.IE
	APNAMBR                       *ie.IE
	PrivateExtension              *ie.IE
	AdditionalIEs                 []*ie.IE
}

// NewUpdatePDPContextResponse creates a new GTPv1 UpdatePDPContextResponse.
func NewUpdatePDPContextResponse(teid uint32, seq uint16, ies ...*ie.IE) *UpdatePDPContextResponse {
	u := &UpdatePDPContextResponse{
		Header: NewHeader(0x32, MsgTypeUpdatePDPContextResponse, teid, seq, nil),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			u.Cause = i
		case ie.Recovery:
			u.Recovery = i
		case ie.TEIDDataI:
			u.TEIDDataI = i
		case ie.TEIDCPlane:
			u.TEIDCPlane = i
		case ie.ChargingID:
			u.ChargingID = i
		case ie.ProtocolConfigurationOptions:
			u.PCO = i
		case ie.GSNAddress:
			if u.GGSNAddressForCPlane == nil {
				u.GGSNAddressForCPlane = i
			} else if u.GGSNAddressForUserTraffic == nil {
				u.GGSNAddressForUserTraffic = i
			} else if u.AltGGSNAddressForCPlane == nil {
				u.AltGGSNAddressForCPlane = i
			} else if u.AltGGSNAddressForUserTraffic == nil {
				u.AltGGSNAddressForUserTraffic = i
			}
		case ie.QoSProfile:
			u.QoSProfile = i
		case ie.ChargingGatewayAddress:
			if u.ChargingGatewayAddress == nil {
				u.ChargingGatewayAddress = i
			} else if u.AltChargingGatewayAddress == nil {
				u.AltChargingGatewayAddress = i
			}
		case ie.CommonFlags:
			u.CommonFlags = i
		case ie.APNRestriction:
			u.APNRestriction = i
		case ie.BearerControlMode:
			u.BearerControlMode = i
		case ie.MSInfoChangeReportingAction:
			u.MSInfoChangeReportingAction = i
		case ie.EvolvedAllocationRetentionPriorityI:
			u.EvolvedARPI = i
		case ie.CSGInformationReportingAction:
			u.CSGInformationReportingAction = i
		case ie.AggregateMaximumBitRate:
			u.APNAMBR = i
		case ie.PrivateExtension:
			u.PrivateExtension = i
		default:
			u.AdditionalIEs = append(u.AdditionalIEs, i)
		}
	}

	u.SetLength()
	return u
}

// Marshal returns the byte sequence generated from a UpdatePDPContextResponse.
func (u *UpdatePDPContextResponse) Marshal() ([]byte, error) {
	b := make([]byte, u.MarshalLen())
	if err := u.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (u *UpdatePDPContextResponse) MarshalTo(b []byte) error {
	if len(b) < u.MarshalLen() {
		return ErrTooShortToMarshal
	}
	u.Header.Payload = make([]byte, u.MarshalLen()-u.Header.MarshalLen())

	offset := 0
	if ie := u.Cause; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.Recovery; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.TEIDDataI; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.TEIDCPlane; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.ChargingID; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.PCO; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.GGSNAddressForCPlane; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.GGSNAddressForUserTraffic; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.AltGGSNAddressForCPlane; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.AltGGSNAddressForUserTraffic; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.QoSProfile; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.ChargingGatewayAddress; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.AltChargingGatewayAddress; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.CommonFlags; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.APNRestriction; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.BearerControlMode; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.MSInfoChangeReportingAction; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.EvolvedARPI; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.CSGInformationReportingAction; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.APNAMBR; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.PrivateExtension; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	for _, ie := range u.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.MarshalTo(u.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	u.Header.SetLength()
	return u.Header.MarshalTo(b)
}

// ParseUpdatePDPContextResponse decodes a given byte sequence as a UpdatePDPContextResponse.
func ParseUpdatePDPContextResponse(b []byte) (*UpdatePDPContextResponse, error) {
	u := &UpdatePDPContextResponse{}
	if err := u.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return u, nil
}

// UnmarshalBinary decodes a given byte sequence as a UpdatePDPContextResponse.
func (u *UpdatePDPContextResponse) UnmarshalBinary(b []byte) error {
	var err error
	u.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(u.Header.Payload) < 2 {
		return nil
	}

	ies, err := ie.ParseMultiIEs(u.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			u.Cause = i
		case ie.Recovery:
			u.Recovery = i
		case ie.TEIDDataI:
			u.TEIDDataI = i
		case ie.TEIDCPlane:
			u.TEIDCPlane = i
		case ie.ChargingID:
			u.ChargingID = i
		case ie.ProtocolConfigurationOptions:
			u.PCO = i
		case ie.GSNAddress:
			if u.GGSNAddressForCPlane == nil {
				u.GGSNAddressForCPlane = i
			} else if u.GGSNAddressForUserTraffic == nil {
				u.GGSNAddressForUserTraffic = i
			} else if u.AltGGSNAddressForCPlane == nil {
				u.AltGGSNAddressForCPlane = i
			} else if u.AltGGSNAddressForUserTraffic == nil {
				u.AltGGSNAddressForUserTraffic = i
			}
		case ie.QoSProfile:
			u.QoSProfile = i
		case ie.ChargingGatewayAddress:
			if u.ChargingGatewayAddress == nil {
				u.ChargingGatewayAddress = i
			} else if u.AltChargingGatewayAddress == nil {
				u.AltChargingGatewayAddress = i
			}
		case ie.CommonFlags:
			u.CommonFlags = i
		case ie.APNRestriction:
			u.APNRestriction = i
		case ie.BearerControlMode:
			u.BearerControlMode = i
		case ie.MSInfoChangeReportingAction:
			u.MSInfoChangeReportingAction = i
		case ie.EvolvedAllocationRetentionPriorityI:
			u.EvolvedARPI = i
		case ie.CSGInformationReportingAction:
			u.CSGInformationReportingAction = i
		case ie.AggregateMaximumBitRate:
			u.APNAMBR = i
		case ie.PrivateExtension:
			u.PrivateExtension = i
		default:
			u.AdditionalIEs = append(u.AdditionalIEs, i)
		}
	}
	return nil
}

// MarshalLen returns the serial length of Data.
func (u *UpdatePDPContextResponse) MarshalLen() int {
	l := u.Header.MarshalLen() - len(u.Header.Payload)

	if ie := u.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.Recovery; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.TEIDDataI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.TEIDCPlane; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.ChargingID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.PCO; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.GGSNAddressForCPlane; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.GGSNAddressForUserTraffic; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.AltGGSNAddressForCPlane; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.AltGGSNAddressForUserTraffic; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.QoSProfile; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.ChargingGatewayAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.AltChargingGatewayAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.CommonFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.APNRestriction; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.BearerControlMode; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.MSInfoChangeReportingAction; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.EvolvedARPI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.CSGInformationReportingAction; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.APNAMBR; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.PrivateExtension; ie != nil {
		l += ie.MarshalLen()
	}

	for _, ie := range u.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.MarshalLen()
	}
	return l
}

// SetLength sets the length in Length field.
func (u *UpdatePDPContextResponse) SetLength() {
	u.Length = uint16(u.MarshalLen() - 8)
}

// MessageTypeName returns the name of protocol.
func (u *UpdatePDPContextResponse) MessageTypeName() string {
	return "Update PDP Context Response"
}

// TEID returns the TEID in human-readable string.
func (u *UpdatePDPContextResponse) TEID() uint32 {
	return u.Header.TEID
}
