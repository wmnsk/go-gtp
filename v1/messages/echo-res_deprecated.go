// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes EchoResponse into bytes.
//
// DEPRECATED: use EchoResponse.Marshal instead.
func (e *EchoResponse) Serialize() ([]byte, error) {
	log.Println("EchoResponse.Serialize is deprecated. use EchoResponse.Marshal instead")
	return e.Marshal()
}

// SerializeTo serializes EchoResponse into bytes given as b.
//
// DEPRECATED: use EchoResponse.MarshalTo instead.
func (e *EchoResponse) SerializeTo(b []byte) error {
	log.Println("EchoResponse.SerializeTo is deprecated. use EchoResponse.MarshalTo instead")
	return e.MarshalTo(b)
}

// DecodeEchoResponse decodes bytes as EchoResponse.
//
// DEPRECATED: use ParseEchoResponse instead.
func DecodeEchoResponse(b []byte) (*EchoResponse, error) {
	log.Println("DecodeEchoResponse is deprecated. use ParseEchoResponse instead")
	return ParseEchoResponse(b)
}

// DecodeFromBytes decodes bytes as EchoResponse.
//
// DEPRECATED: use EchoResponse.UnmarshalBinary instead.
func (e *EchoResponse) DecodeFromBytes(b []byte) error {
	log.Println("EchoResponse.DecodeFromBytes is deprecated. use EchoResponse.UnmarshalBinary instead")
	return e.UnmarshalBinary(b)
}

// Len returns the actual length of EchoResponse.
//
// DEPRECATED: use EchoResponse.MarshalLen instead.
func (e *EchoResponse) Len() int {
	log.Println("EchoResponse.Len is deprecated. use EchoResponse.MarshalLen instead")
	return e.MarshalLen()
}
