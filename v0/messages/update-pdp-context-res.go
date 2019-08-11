// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/v0/ies"
)

// UpdatePDPContextResponse is a UpdatePDPContextResponse Header and its IEs above.
type UpdatePDPContextResponse struct {
	*Header
	Cause                     *ies.IE
	QoSProfile                *ies.IE
	Recovery                  *ies.IE
	FlowLabelDataI            *ies.IE
	FlowLabelSignalling       *ies.IE
	ChargingID                *ies.IE
	EndUserAddress            *ies.IE
	GGSNAddressForSignalling  *ies.IE
	GGSNAddressForUserTraffic *ies.IE
	ChargingGatewayAddress    *ies.IE
	PrivateExtension          *ies.IE
	AdditionalIEs             []*ies.IE
}

// NewUpdatePDPContextResponse creates a new UpdatePDPContextResponse.
func NewUpdatePDPContextResponse(seq, label uint16, tid uint64, ie ...*ies.IE) *UpdatePDPContextResponse {
	u := &UpdatePDPContextResponse{
		Header: NewHeader(
			0x1e, MsgTypeUpdatePDPContextResponse, seq, label, tid, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			u.Cause = i
		case ies.QualityOfServiceProfile:
			u.QoSProfile = i
		case ies.Recovery:
			u.Recovery = i
		case ies.FlowLabelDataI:
			u.FlowLabelDataI = i
		case ies.FlowLabelSignalling:
			u.FlowLabelSignalling = i
		case ies.ChargingID:
			u.ChargingID = i
		case ies.EndUserAddress:
			u.EndUserAddress = i
		case ies.GSNAddress:
			if u.GGSNAddressForSignalling == nil {
				u.GGSNAddressForSignalling = i
			} else {
				u.GGSNAddressForUserTraffic = i
			}
		case ies.ChargingGatewayAddress:
			u.ChargingGatewayAddress = i
		case ies.PrivateExtension:
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
	if u.Header.Payload != nil {
		u.Header.Payload = nil
	}
	u.Header.Payload = make([]byte, u.MarshalLen()-u.Header.MarshalLen())

	offset := 0
	if ie := u.Cause; ie != nil {
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
	if ie := u.Recovery; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.FlowLabelDataI; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.FlowLabelSignalling; ie != nil {
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
	if ie := u.EndUserAddress; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.GGSNAddressForSignalling; ie != nil {
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
	if ie := u.ChargingGatewayAddress; ie != nil {
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

// ParseUpdatePDPContextResponse Parses a given byte sequence as a UpdatePDPContextResponse.
func ParseUpdatePDPContextResponse(b []byte) (*UpdatePDPContextResponse, error) {
	u := &UpdatePDPContextResponse{}
	if err := u.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return u, nil
}

// UnmarshalBinary Parses a given byte sequence as a UpdatePDPContextResponse.
func (u *UpdatePDPContextResponse) UnmarshalBinary(b []byte) error {
	var err error
	u.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(u.Header.Payload) < 2 {
		return nil
	}

	ie, err := ies.ParseMultiIEs(u.Header.Payload)
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
		case ies.QualityOfServiceProfile:
			u.QoSProfile = i
		case ies.Recovery:
			u.Recovery = i
		case ies.FlowLabelDataI:
			u.FlowLabelDataI = i
		case ies.FlowLabelSignalling:
			u.FlowLabelSignalling = i
		case ies.ChargingID:
			u.ChargingID = i
		case ies.EndUserAddress:
			u.EndUserAddress = i
		case ies.GSNAddress:
			if u.GGSNAddressForSignalling == nil {
				u.GGSNAddressForSignalling = i
			} else {
				u.GGSNAddressForUserTraffic = i
			}
		case ies.ChargingGatewayAddress:
			u.ChargingGatewayAddress = i
		case ies.PrivateExtension:
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
	if ie := u.QoSProfile; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.Recovery; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.FlowLabelDataI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.FlowLabelSignalling; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.ChargingID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.EndUserAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.GGSNAddressForSignalling; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.GGSNAddressForUserTraffic; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.ChargingGatewayAddress; ie != nil {
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
	u.Header.Length = uint16(u.MarshalLen() - 20)
}

// MessageTypeName returns the name of protocol.
func (u *UpdatePDPContextResponse) MessageTypeName() string {
	return "Update PDP Context Response"
}

// TID returns the TID in human-readable string.
func (u *UpdatePDPContextResponse) TID() string {
	return u.tid()
}
