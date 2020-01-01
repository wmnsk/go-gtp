// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes ErrorIndication into bytes.
//
// DEPRECATED: use ErrorIndication.Marshal instead.
func (e *ErrorIndication) Serialize() ([]byte, error) {
	log.Println("ErrorIndication.Serialize is deprecated. use ErrorIndication.Marshal instead")
	return e.Marshal()
}

// SerializeTo serializes ErrorIndication into bytes given as b.
//
// DEPRECATED: use ErrorIndication.MarshalTo instead.
func (e *ErrorIndication) SerializeTo(b []byte) error {
	log.Println("ErrorIndication.SerializeTo is deprecated. use ErrorIndication.MarshalTo instead")
	return e.MarshalTo(b)
}

// DecodeErrorIndication decodes bytes as ErrorIndication.
//
// DEPRECATED: use ParseErrorIndication instead.
func DecodeErrorIndication(b []byte) (*ErrorIndication, error) {
	log.Println("DecodeErrorIndication is deprecated. use ParseErrorIndication instead")
	return ParseErrorIndication(b)
}

// DecodeFromBytes decodes bytes as ErrorIndication.
//
// DEPRECATED: use ErrorIndication.UnmarshalBinary instead.
func (e *ErrorIndication) DecodeFromBytes(b []byte) error {
	log.Println("ErrorIndication.DecodeFromBytes is deprecated. use ErrorIndication.UnmarshalBinary instead")
	return e.UnmarshalBinary(b)
}

// Len returns the actual length of ErrorIndication.
//
// DEPRECATED: use ErrorIndication.MarshalLen instead.
func (e *ErrorIndication) Len() int {
	log.Println("ErrorIndication.Len is deprecated. use ErrorIndication.MarshalLen instead")
	return e.MarshalLen()
}
