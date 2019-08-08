// Copyright 2019 go-gtp authors. All rights reservev.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "github.com/wmnsk/go-gtp/v2/ies"

// StopPagingIndication is a StopPagingIndication Header and its IEs above.
type StopPagingIndication struct {
	*Header
	IMSI             *ies.IE
	PrivateExtension *ies.IE
	AdditionalIEs    []*ies.IE
}

// NewStopPagingIndication creates a new StopPagingIndication.
func NewStopPagingIndication(teid, seq uint32, ie ...*ies.IE) *StopPagingIndication {
	s := &StopPagingIndication{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeStopPagingIndication, teid, seq, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.IMSI:
			s.IMSI = i
		case ies.PrivateExtension:
			s.PrivateExtension = i
		default:
			s.AdditionalIEs = append(s.AdditionalIEs, i)
		}
	}

	s.SetLength()
	return s
}

// Serialize serializes StopPagingIndication into bytes.
func (s *StopPagingIndication) Serialize() ([]byte, error) {
	b := make([]byte, s.Len())
	if err := s.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes StopPagingIndication into bytes.
func (s *StopPagingIndication) SerializeTo(b []byte) error {
	if s.Header.Payload != nil {
		s.Header.Payload = nil
	}
	s.Header.Payload = make([]byte, s.Len()-s.Header.Len())

	offset := 0
	if ie := s.IMSI; ie != nil {
		if err := ie.SerializeTo(s.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := s.PrivateExtension; ie != nil {
		if err := ie.SerializeTo(s.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	for _, ie := range s.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.SerializeTo(s.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	s.Header.SetLength()
	return s.Header.SerializeTo(b)
}

// DecodeStopPagingIndication decodes given bytes as StopPagingIndication.
func DecodeStopPagingIndication(b []byte) (*StopPagingIndication, error) {
	s := &StopPagingIndication{}
	if err := s.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return s, nil
}

// DecodeFromBytes decodes given bytes as StopPagingIndication.
func (s *StopPagingIndication) DecodeFromBytes(b []byte) error {
	var err error
	s.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}
	if len(s.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ies.DecodeMultiIEs(s.Header.Payload)
	if err != nil {
		return err
	}
	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.IMSI:
			s.IMSI = i
		case ies.PrivateExtension:
			s.PrivateExtension = i
		default:
			s.AdditionalIEs = append(s.AdditionalIEs, i)
		}
	}

	return nil
}

// Len returns the actual length in int.
func (s *StopPagingIndication) Len() int {
	l := s.Header.Len() - len(s.Header.Payload)
	if ie := s.IMSI; ie != nil {
		l += ie.Len()
	}
	if ie := s.PrivateExtension; ie != nil {
		l += ie.Len()
	}

	for _, ie := range s.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.Len()
	}
	return l
}

// SetLength sets the length in Length field.
func (s *StopPagingIndication) SetLength() {
	s.Header.Length = uint16(s.Len() - 4)
}

// MessageTypeName returns the name of protocol.
func (s *StopPagingIndication) MessageTypeName() string {
	return "Stop Paging Indication"
}

// TEID returns the TEID in uint32.
func (s *StopPagingIndication) TEID() uint32 {
	return s.Header.teid()
}
