// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes StopPagingIndication into bytes.
//
// DEPRECATED: use StopPagingIndication.Marshal instead.
func (s *StopPagingIndication) Serialize() ([]byte, error) {
	log.Println("StopPagingIndication.Serialize is deprecated. use StopPagingIndication.Marshal instead")
	return s.Marshal()
}

// SerializeTo serializes StopPagingIndication into bytes given as b.
//
// DEPRECATED: use StopPagingIndication.MarshalTo instead.
func (s *StopPagingIndication) SerializeTo(b []byte) error {
	log.Println("StopPagingIndication.SerializeTo is deprecated. use StopPagingIndication.MarshalTo instead")
	return s.MarshalTo(b)
}

// DecodeStopPagingIndication decodes bytes as StopPagingIndication.
//
// DEPRECATED: use ParseStopPagingIndication instead.
func DecodeStopPagingIndication(b []byte) (*StopPagingIndication, error) {
	log.Println("DecodeStopPagingIndication is deprecated. use ParseStopPagingIndication instead")
	return ParseStopPagingIndication(b)
}

// DecodeFromBytes decodes bytes as StopPagingIndication.
//
// DEPRECATED: use StopPagingIndication.UnmarshalBinary instead.
func (s *StopPagingIndication) DecodeFromBytes(b []byte) error {
	log.Println("StopPagingIndication.DecodeFromBytes is deprecated. use StopPagingIndication.UnmarshalBinary instead")
	return s.UnmarshalBinary(b)
}

// Len returns the actual length of StopPagingIndication.
//
// DEPRECATED: use StopPagingIndication.MarshalLen instead.
func (s *StopPagingIndication) Len() int {
	log.Println("StopPagingIndication.Len is deprecated. use StopPagingIndication.MarshalLen instead")
	return s.MarshalLen()
}
