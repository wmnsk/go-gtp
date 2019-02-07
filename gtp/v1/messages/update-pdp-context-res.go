// Copyright 2019 go-gtp authors. All rights reserveu.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/gtp/v1/ies"
)

// UpdatePDPContextResponse is a UpdatePDPContextResponse Header and its IEs above.
type UpdatePDPContextResponse struct {
	*Header
	Cause                         *ies.IE
	Recovery                      *ies.IE
	TEIDDataI                     *ies.IE
	TEIDCPlane                    *ies.IE
	ChargingID                    *ies.IE
	PCO                           *ies.IE
	GGSNAddressForCPlane          *ies.IE
	GGSNAddressForUserTraffic     *ies.IE
	AltGGSNAddressForCPlane       *ies.IE
	AltGGSNAddressForUserTraffic  *ies.IE
	QoSProfile                    *ies.IE
	ChargingGatewayAddress        *ies.IE
	AltChargingGatewayAddress     *ies.IE
	CommonFlags                   *ies.IE
	APNRestriction                *ies.IE
	BearerControlMode             *ies.IE
	MSInfoChangeReportingAction   *ies.IE
	EvolvedARPI                   *ies.IE
	CSGInformationReportingAction *ies.IE
	APNAMBR                       *ies.IE
	PrivateExtension              *ies.IE
	AdditionalIEs                 []*ies.IE
}

// NewUpdatePDPContextResponse creates a new GTPv1 UpdatePDPContextResponse.
func NewUpdatePDPContextResponse(teid uint32, seq uint16, ie ...*ies.IE) *UpdatePDPContextResponse {
	u := &UpdatePDPContextResponse{
		Header: NewHeader(0x32, MsgTypeUpdatePDPContextResponse, teid, seq, nil),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			u.Cause = i
		case ies.Recovery:
			u.Recovery = i
		case ies.TEIDDataI:
			u.TEIDDataI = i
		case ies.TEIDCPlane:
			u.TEIDCPlane = i
		case ies.ChargingID:
			u.ChargingID = i
		case ies.ProtocolConfigurationOptions:
			u.PCO = i
		case ies.GSNAddress:
			if u.GGSNAddressForCPlane == nil {
				u.GGSNAddressForCPlane = i
			} else if u.GGSNAddressForUserTraffic == nil {
				u.GGSNAddressForUserTraffic = i
			} else if u.AltGGSNAddressForCPlane == nil {
				u.AltGGSNAddressForCPlane = i
			} else if u.AltGGSNAddressForUserTraffic == nil {
				u.AltGGSNAddressForUserTraffic = i
			}
		case ies.QoSProfile:
			u.QoSProfile = i
		case ies.ChargingGatewayAddress:
			if u.ChargingGatewayAddress == nil {
				u.ChargingGatewayAddress = i
			} else if u.AltChargingGatewayAddress == nil {
				u.AltChargingGatewayAddress = i
			}
		case ies.CommonFlags:
			u.CommonFlags = i
		case ies.APNRestriction:
			u.APNRestriction = i
		case ies.BearerControlMode:
			u.BearerControlMode = i
		case ies.MSInfoChangeReportingAction:
			u.MSInfoChangeReportingAction = i
		case ies.EvolvedAllocationRetentionPriorityI:
			u.EvolvedARPI = i
		case ies.CSGInformationReportingAction:
			u.CSGInformationReportingAction = i
		case ies.AggregateMaximumBitRate:
			u.APNAMBR = i
		case ies.PrivateExtension:
			u.PrivateExtension = i
		default:
			u.AdditionalIEs = append(u.AdditionalIEs, i)
		}
	}

	u.SetLength()
	return u
}

// Serialize returns the byte sequence generated from a UpdatePDPContextResponse.
func (u *UpdatePDPContextResponse) Serialize() ([]byte, error) {
	b := make([]byte, u.Len())
	if err := u.SerializeTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// SerializeTo puts the byte sequence in the byte array given as b.
func (u *UpdatePDPContextResponse) SerializeTo(b []byte) error {
	if len(b) < u.Len() {
		return ErrTooShortToSerialize
	}
	u.Header.Payload = make([]byte, u.Len()-u.Header.Len())

	offset := 0
	if ie := u.Cause; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.Recovery; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.TEIDDataI; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.TEIDCPlane; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.ChargingID; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.PCO; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.GGSNAddressForCPlane; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.GGSNAddressForUserTraffic; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.AltGGSNAddressForCPlane; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.AltGGSNAddressForUserTraffic; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.QoSProfile; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.ChargingGatewayAddress; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.AltChargingGatewayAddress; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.CommonFlags; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.APNRestriction; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.BearerControlMode; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.MSInfoChangeReportingAction; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.EvolvedARPI; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.CSGInformationReportingAction; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.APNAMBR; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.PrivateExtension; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	for _, ie := range u.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.SerializeTo(u.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	u.Header.SetLength()
	return u.Header.SerializeTo(b)
}

// DecodeUpdatePDPContextResponse decodes a given byte sequence as a UpdatePDPContextResponse.
func DecodeUpdatePDPContextResponse(b []byte) (*UpdatePDPContextResponse, error) {
	u := &UpdatePDPContextResponse{}
	if err := u.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return u, nil
}

// DecodeFromBytes decodes a given byte sequence as a UpdatePDPContextResponse.
func (u *UpdatePDPContextResponse) DecodeFromBytes(b []byte) error {
	var err error
	u.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}
	if len(u.Header.Payload) < 2 {
		return nil
	}

	ie, err := ies.DecodeMultiIEs(u.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			u.Cause = i
		case ies.Recovery:
			u.Recovery = i
		case ies.TEIDDataI:
			u.TEIDDataI = i
		case ies.TEIDCPlane:
			u.TEIDCPlane = i
		case ies.ChargingID:
			u.ChargingID = i
		case ies.ProtocolConfigurationOptions:
			u.PCO = i
		case ies.GSNAddress:
			if u.GGSNAddressForCPlane == nil {
				u.GGSNAddressForCPlane = i
			} else if u.GGSNAddressForUserTraffic == nil {
				u.GGSNAddressForUserTraffic = i
			} else if u.AltGGSNAddressForCPlane == nil {
				u.AltGGSNAddressForCPlane = i
			} else if u.AltGGSNAddressForUserTraffic == nil {
				u.AltGGSNAddressForUserTraffic = i
			}
		case ies.QoSProfile:
			u.QoSProfile = i
		case ies.ChargingGatewayAddress:
			if u.ChargingGatewayAddress == nil {
				u.ChargingGatewayAddress = i
			} else if u.AltChargingGatewayAddress == nil {
				u.AltChargingGatewayAddress = i
			}
		case ies.CommonFlags:
			u.CommonFlags = i
		case ies.APNRestriction:
			u.APNRestriction = i
		case ies.BearerControlMode:
			u.BearerControlMode = i
		case ies.MSInfoChangeReportingAction:
			u.MSInfoChangeReportingAction = i
		case ies.EvolvedAllocationRetentionPriorityI:
			u.EvolvedARPI = i
		case ies.CSGInformationReportingAction:
			u.CSGInformationReportingAction = i
		case ies.AggregateMaximumBitRate:
			u.APNAMBR = i
		case ies.PrivateExtension:
			u.PrivateExtension = i
		default:
			u.AdditionalIEs = append(u.AdditionalIEs, i)
		}
	}
	return nil
}

// Len returns the actual length of Data.
func (u *UpdatePDPContextResponse) Len() int {
	l := u.Header.Len() - len(u.Header.Payload)

	if ie := u.Cause; ie != nil {
		l += ie.Len()
	}
	if ie := u.Recovery; ie != nil {
		l += ie.Len()
	}
	if ie := u.TEIDDataI; ie != nil {
		l += ie.Len()
	}
	if ie := u.TEIDCPlane; ie != nil {
		l += ie.Len()
	}
	if ie := u.ChargingID; ie != nil {
		l += ie.Len()
	}
	if ie := u.PCO; ie != nil {
		l += ie.Len()
	}
	if ie := u.GGSNAddressForCPlane; ie != nil {
		l += ie.Len()
	}
	if ie := u.GGSNAddressForUserTraffic; ie != nil {
		l += ie.Len()
	}
	if ie := u.AltGGSNAddressForCPlane; ie != nil {
		l += ie.Len()
	}
	if ie := u.AltGGSNAddressForUserTraffic; ie != nil {
		l += ie.Len()
	}
	if ie := u.QoSProfile; ie != nil {
		l += ie.Len()
	}
	if ie := u.ChargingGatewayAddress; ie != nil {
		l += ie.Len()
	}
	if ie := u.AltChargingGatewayAddress; ie != nil {
		l += ie.Len()
	}
	if ie := u.CommonFlags; ie != nil {
		l += ie.Len()
	}
	if ie := u.APNRestriction; ie != nil {
		l += ie.Len()
	}
	if ie := u.BearerControlMode; ie != nil {
		l += ie.Len()
	}
	if ie := u.MSInfoChangeReportingAction; ie != nil {
		l += ie.Len()
	}
	if ie := u.EvolvedARPI; ie != nil {
		l += ie.Len()
	}
	if ie := u.CSGInformationReportingAction; ie != nil {
		l += ie.Len()
	}
	if ie := u.APNAMBR; ie != nil {
		l += ie.Len()
	}
	if ie := u.PrivateExtension; ie != nil {
		l += ie.Len()
	}

	for _, ie := range u.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.Len()
	}
	return l
}

// SetLength sets the length in Length field.
func (u *UpdatePDPContextResponse) SetLength() {
	u.Length = uint16(u.Len() - 8)
}

// MessageTypeName returns the name of protocol.
func (u *UpdatePDPContextResponse) MessageTypeName() string {
	return "Update PDP Context Response"
}

// TEID returns the TEID in human-readable string.
func (u *UpdatePDPContextResponse) TEID() uint32 {
	return u.Header.TEID
}
