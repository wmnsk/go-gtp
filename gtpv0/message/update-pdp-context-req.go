// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv0/ie"
)

// UpdatePDPContextRequest is a UpdatePDPContextRequest Header and its IEs above.
type UpdatePDPContextRequest struct {
	*Header
	RAI                       *ie.IE
	QoSProfile                *ie.IE
	Recovery                  *ie.IE
	FlowLabelDataI            *ie.IE
	FlowLabelSignalling       *ie.IE
	EndUserAddress            *ie.IE
	SGSNAddressForSignalling  *ie.IE
	SGSNAddressForUserTraffic *ie.IE
	PrivateExtension          *ie.IE
	AdditionalIEs             []*ie.IE
}

// NewUpdatePDPContextRequest creates a new UpdatePDPContextRequest.
func NewUpdatePDPContextRequest(seq, label uint16, tid uint64, ies ...*ie.IE) *UpdatePDPContextRequest {
	u := &UpdatePDPContextRequest{
		Header: NewHeader(
			0x1e, MsgTypeUpdatePDPContextRequest, seq, label, tid, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.RouteingAreaIdentity:
			u.RAI = i
		case ie.QualityOfServiceProfile:
			u.QoSProfile = i
		case ie.Recovery:
			u.Recovery = i
		case ie.FlowLabelDataI:
			u.FlowLabelDataI = i
		case ie.FlowLabelSignalling:
			u.FlowLabelSignalling = i
		case ie.EndUserAddress:
			u.EndUserAddress = i
		case ie.GSNAddress:
			if u.SGSNAddressForSignalling == nil {
				u.SGSNAddressForSignalling = i
			} else {
				u.SGSNAddressForUserTraffic = i
			}
		case ie.PrivateExtension:
			u.PrivateExtension = i
		default:
			u.AdditionalIEs = append(u.AdditionalIEs, i)
		}
	}

	u.SetLength()
	return u
}

// Marshal returns the byte sequence generated from a UpdatePDPContextRequest.
func (u *UpdatePDPContextRequest) Marshal() ([]byte, error) {
	b := make([]byte, u.MarshalLen())
	if err := u.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (u *UpdatePDPContextRequest) MarshalTo(b []byte) error {
	if u.Header.Payload != nil {
		u.Header.Payload = nil
	}
	u.Header.Payload = make([]byte, u.MarshalLen()-u.Header.MarshalLen())

	offset := 0
	if ie := u.RAI; ie != nil {
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
	if ie := u.EndUserAddress; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.SGSNAddressForSignalling; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.SGSNAddressForUserTraffic; ie != nil {
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

// ParseUpdatePDPContextRequest parses a given byte sequence as a UpdatePDPContextRequest.
func ParseUpdatePDPContextRequest(b []byte) (*UpdatePDPContextRequest, error) {
	u := &UpdatePDPContextRequest{}
	if err := u.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return u, nil
}

// UnmarshalBinary parses a given byte sequence as a UpdatePDPContextRequest.
func (u *UpdatePDPContextRequest) UnmarshalBinary(b []byte) error {
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
		case ie.RouteingAreaIdentity:
			u.RAI = i
		case ie.QualityOfServiceProfile:
			u.QoSProfile = i
		case ie.Recovery:
			u.Recovery = i
		case ie.FlowLabelDataI:
			u.FlowLabelDataI = i
		case ie.FlowLabelSignalling:
			u.FlowLabelSignalling = i
		case ie.EndUserAddress:
			u.EndUserAddress = i
		case ie.GSNAddress:
			if u.SGSNAddressForSignalling == nil {
				u.SGSNAddressForSignalling = i
			} else {
				u.SGSNAddressForUserTraffic = i
			}
		case ie.PrivateExtension:
			u.PrivateExtension = i
		default:
			u.AdditionalIEs = append(u.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length of Data.
func (u *UpdatePDPContextRequest) MarshalLen() int {
	l := u.Header.MarshalLen() - len(u.Header.Payload)

	if ie := u.RAI; ie != nil {
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
	if ie := u.EndUserAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.SGSNAddressForSignalling; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.SGSNAddressForUserTraffic; ie != nil {
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
func (u *UpdatePDPContextRequest) SetLength() {
	u.Header.Length = uint16(u.MarshalLen() - 20)
}

// MessageTypeName returns the name of protocol.
func (u *UpdatePDPContextRequest) MessageTypeName() string {
	return "Update PDP Context Request"
}

// TID returns the TID in human-readable string.
func (u *UpdatePDPContextRequest) TID() string {
	return u.tid()
}
