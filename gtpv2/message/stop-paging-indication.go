// Copyright 2019-2023 go-gtp authors. All rights reservev.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "github.com/wmnsk/go-gtp/gtpv2/ie"

// StopPagingIndication is a StopPagingIndication Header and its IEs above.
type StopPagingIndication struct {
	*Header
	IMSI             *ie.IE
	PrivateExtension *ie.IE
	AdditionalIEs    []*ie.IE
}

// NewStopPagingIndication creates a new StopPagingIndication.
func NewStopPagingIndication(teid, seq uint32, ies ...*ie.IE) *StopPagingIndication {
	s := &StopPagingIndication{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeStopPagingIndication, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.IMSI:
			s.IMSI = i
		case ie.PrivateExtension:
			s.PrivateExtension = i
		default:
			s.AdditionalIEs = append(s.AdditionalIEs, i)
		}
	}

	s.SetLength()
	return s
}

// Marshal serializes StopPagingIndication into bytes.
func (s *StopPagingIndication) Marshal() ([]byte, error) {
	b := make([]byte, s.MarshalLen())
	if err := s.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes StopPagingIndication into bytes.
func (s *StopPagingIndication) MarshalTo(b []byte) error {
	if s.Header.Payload != nil {
		s.Header.Payload = nil
	}
	s.Header.Payload = make([]byte, s.MarshalLen()-s.Header.MarshalLen())

	offset := 0
	if ie := s.IMSI; ie != nil {
		if err := ie.MarshalTo(s.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := s.PrivateExtension; ie != nil {
		if err := ie.MarshalTo(s.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	for _, ie := range s.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.MarshalTo(s.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	s.Header.SetLength()
	return s.Header.MarshalTo(b)
}

// ParseStopPagingIndication decodes given bytes as StopPagingIndication.
func ParseStopPagingIndication(b []byte) (*StopPagingIndication, error) {
	s := &StopPagingIndication{}
	if err := s.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return s, nil
}

// UnmarshalBinary decodes given bytes as StopPagingIndication.
func (s *StopPagingIndication) UnmarshalBinary(b []byte) error {
	var err error
	s.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(s.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ie.ParseMultiIEs(s.Header.Payload)
	if err != nil {
		return err
	}
	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.IMSI:
			s.IMSI = i
		case ie.PrivateExtension:
			s.PrivateExtension = i
		default:
			s.AdditionalIEs = append(s.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (s *StopPagingIndication) MarshalLen() int {
	l := s.Header.MarshalLen() - len(s.Header.Payload)
	if ie := s.IMSI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := s.PrivateExtension; ie != nil {
		l += ie.MarshalLen()
	}

	for _, ie := range s.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.MarshalLen()
	}
	return l
}

// SetLength sets the length in Length field.
func (s *StopPagingIndication) SetLength() {
	s.Header.Length = uint16(s.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (s *StopPagingIndication) MessageTypeName() string {
	return "Stop Paging Indication"
}

// TEID returns the TEID in uint32.
func (s *StopPagingIndication) TEID() uint32 {
	return s.Header.teid()
}
