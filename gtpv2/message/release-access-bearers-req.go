// Copyright 2019-2023 go-gtp authors. All rights reservev.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "github.com/wmnsk/go-gtp/gtpv2/ie"

// ReleaseAccessBearersRequest is a ReleaseAccessBearersRequest Header and its IEs above.
type ReleaseAccessBearersRequest struct {
	*Header
	ListOfRABs                  []*ie.IE
	OriginatingNode             *ie.IE
	IndicationFlags             *ie.IE
	SecondaryRATUsageDataReport []*ie.IE
	PrivateExtension            *ie.IE
	AdditionalIEs               []*ie.IE
}

// NewReleaseAccessBearersRequest creates a new ReleaseAccessBearersRequest.
func NewReleaseAccessBearersRequest(teid, seq uint32, ies ...*ie.IE) *ReleaseAccessBearersRequest {
	r := &ReleaseAccessBearersRequest{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeReleaseAccessBearersRequest, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.EPSBearerID:
			r.ListOfRABs = append(r.ListOfRABs, i)
		case ie.NodeType:
			r.OriginatingNode = i
		case ie.Indication:
			r.IndicationFlags = i
		case ie.SecondaryRATUsageDataReport:
			r.SecondaryRATUsageDataReport = append(r.SecondaryRATUsageDataReport, i)
		case ie.PrivateExtension:
			r.PrivateExtension = i
		default:
			r.AdditionalIEs = append(r.AdditionalIEs, i)
		}
	}

	r.SetLength()
	return r
}

// Marshal serializes ReleaseAccessBearersRequest into bytes.
func (r *ReleaseAccessBearersRequest) Marshal() ([]byte, error) {
	b := make([]byte, r.MarshalLen())
	if err := r.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes ReleaseAccessBearersRequest into bytes.
func (r *ReleaseAccessBearersRequest) MarshalTo(b []byte) error {
	if r.Header.Payload != nil {
		r.Header.Payload = nil
	}
	r.Header.Payload = make([]byte, r.MarshalLen()-r.Header.MarshalLen())

	offset := 0
	for _, ie := range r.ListOfRABs {
		if err := ie.MarshalTo(r.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := r.OriginatingNode; ie != nil {
		if err := ie.MarshalTo(r.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := r.IndicationFlags; ie != nil {
		if err := ie.MarshalTo(r.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	for _, ie := range r.SecondaryRATUsageDataReport {
		if err := ie.MarshalTo(r.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := r.PrivateExtension; ie != nil {
		if err := ie.MarshalTo(r.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	for _, ie := range r.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.MarshalTo(r.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	r.Header.SetLength()
	return r.Header.MarshalTo(b)
}

// ParseReleaseAccessBearersRequest decodes given bytes as ReleaseAccessBearersRequest.
func ParseReleaseAccessBearersRequest(b []byte) (*ReleaseAccessBearersRequest, error) {
	r := &ReleaseAccessBearersRequest{}
	if err := r.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return r, nil
}

// UnmarshalBinary decodes given bytes as ReleaseAccessBearersRequest.
func (r *ReleaseAccessBearersRequest) UnmarshalBinary(b []byte) error {
	var err error
	r.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(r.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ie.ParseMultiIEs(r.Header.Payload)
	if err != nil {
		return err
	}
	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.EPSBearerID:
			r.ListOfRABs = append(r.ListOfRABs, i)
		case ie.NodeType:
			r.OriginatingNode = i
		case ie.Indication:
			r.IndicationFlags = i
		case ie.SecondaryRATUsageDataReport:
			r.SecondaryRATUsageDataReport = append(r.SecondaryRATUsageDataReport, i)
		case ie.PrivateExtension:
			r.PrivateExtension = i
		default:
			r.AdditionalIEs = append(r.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (r *ReleaseAccessBearersRequest) MarshalLen() int {
	l := r.Header.MarshalLen() - len(r.Header.Payload)
	for _, ie := range r.ListOfRABs {
		l += ie.MarshalLen()
	}
	if ie := r.OriginatingNode; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := r.IndicationFlags; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range r.SecondaryRATUsageDataReport {
		l += ie.MarshalLen()
	}
	if ie := r.PrivateExtension; ie != nil {
		l += ie.MarshalLen()
	}

	for _, ie := range r.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.MarshalLen()
	}
	return l
}

// SetLength sets the length in Length field.
func (r *ReleaseAccessBearersRequest) SetLength() {
	r.Header.Length = uint16(r.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (r *ReleaseAccessBearersRequest) MessageTypeName() string {
	return "Release Access Bearers Request"
}

// TEID returns the TEID in uint32.
func (r *ReleaseAccessBearersRequest) TEID() uint32 {
	return r.Header.teid()
}
