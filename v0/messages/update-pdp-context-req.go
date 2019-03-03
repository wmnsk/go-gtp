// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/v0/ies"
)

// UpdatePDPContextRequest is a UpdatePDPContextRequest Header and its IEs above.
type UpdatePDPContextRequest struct {
	*Header
	RAI                       *ies.IE
	QoSProfile                *ies.IE
	Recovery                  *ies.IE
	FlowLabelDataI            *ies.IE
	FlowLabelSignalling       *ies.IE
	EndUserAddress            *ies.IE
	SGSNAddressForSignalling  *ies.IE
	SGSNAddressForUserTraffic *ies.IE
	PrivateExtension          *ies.IE
	AdditionalIEs             []*ies.IE
}

// NewUpdatePDPContextRequest creates a new UpdatePDPContextRequest.
func NewUpdatePDPContextRequest(seq, label uint16, tid uint64, ie ...*ies.IE) *UpdatePDPContextRequest {
	u := &UpdatePDPContextRequest{
		Header: NewHeader(
			0x1e, MsgTypeUpdatePDPContextRequest, seq, label, tid, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.RouteingAreaIdentity:
			u.RAI = i
		case ies.QualityOfServiceProfile:
			u.QoSProfile = i
		case ies.Recovery:
			u.Recovery = i
		case ies.FlowLabelDataI:
			u.FlowLabelDataI = i
		case ies.FlowLabelSignalling:
			u.FlowLabelSignalling = i
		case ies.EndUserAddress:
			u.EndUserAddress = i
		case ies.GSNAddress:
			if u.SGSNAddressForSignalling == nil {
				u.SGSNAddressForSignalling = i
			} else {
				u.SGSNAddressForUserTraffic = i
			}
		case ies.PrivateExtension:
			u.PrivateExtension = i
		default:
			u.AdditionalIEs = append(u.AdditionalIEs, i)
		}
	}

	u.SetLength()
	return u
}

// Serialize returns the byte sequence generated from a UpdatePDPContextRequest.
func (u *UpdatePDPContextRequest) Serialize() ([]byte, error) {
	b := make([]byte, u.Len())
	if err := u.SerializeTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// SerializeTo puts the byte sequence in the byte array given as b.
func (u *UpdatePDPContextRequest) SerializeTo(b []byte) error {
	if u.Header.Payload != nil {
		u.Header.Payload = nil
	}
	u.Header.Payload = make([]byte, u.Len()-u.Header.Len())

	offset := 0
	if ie := u.RAI; ie != nil {
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
	if ie := u.Recovery; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.FlowLabelDataI; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.FlowLabelSignalling; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.EndUserAddress; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.SGSNAddressForSignalling; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.SGSNAddressForUserTraffic; ie != nil {
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

// DecodeUpdatePDPContextRequest decodes a given byte sequence as a UpdatePDPContextRequest.
func DecodeUpdatePDPContextRequest(b []byte) (*UpdatePDPContextRequest, error) {
	u := &UpdatePDPContextRequest{}
	if err := u.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return u, nil
}

// DecodeFromBytes decodes a given byte sequence as a UpdatePDPContextRequest.
func (u *UpdatePDPContextRequest) DecodeFromBytes(b []byte) error {
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
		case ies.RouteingAreaIdentity:
			u.RAI = i
		case ies.QualityOfServiceProfile:
			u.QoSProfile = i
		case ies.Recovery:
			u.Recovery = i
		case ies.FlowLabelDataI:
			u.FlowLabelDataI = i
		case ies.FlowLabelSignalling:
			u.FlowLabelSignalling = i
		case ies.EndUserAddress:
			u.EndUserAddress = i
		case ies.GSNAddress:
			if u.SGSNAddressForSignalling == nil {
				u.SGSNAddressForSignalling = i
			} else {
				u.SGSNAddressForUserTraffic = i
			}
		case ies.PrivateExtension:
			u.PrivateExtension = i
		default:
			u.AdditionalIEs = append(u.AdditionalIEs, i)
		}
	}

	return nil
}

// Len returns the actual length of Data.
func (u *UpdatePDPContextRequest) Len() int {
	l := u.Header.Len() - len(u.Header.Payload)

	if ie := u.RAI; ie != nil {
		l += ie.Len()
	}
	if ie := u.QoSProfile; ie != nil {
		l += ie.Len()
	}
	if ie := u.Recovery; ie != nil {
		l += ie.Len()
	}
	if ie := u.FlowLabelDataI; ie != nil {
		l += ie.Len()
	}
	if ie := u.FlowLabelSignalling; ie != nil {
		l += ie.Len()
	}
	if ie := u.EndUserAddress; ie != nil {
		l += ie.Len()
	}
	if ie := u.SGSNAddressForSignalling; ie != nil {
		l += ie.Len()
	}
	if ie := u.SGSNAddressForUserTraffic; ie != nil {
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
func (u *UpdatePDPContextRequest) SetLength() {
	u.Header.Length = uint16(u.Len() - 20)
}

// MessageTypeName returns the name of protocol.
func (u *UpdatePDPContextRequest) MessageTypeName() string {
	return "Update PDP Context Request"
}

// TID returns the TID in human-readable string.
func (u *UpdatePDPContextRequest) TID() string {
	return u.tid()
}
