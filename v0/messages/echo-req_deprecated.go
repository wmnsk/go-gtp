// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes EchoRequest into bytes.
//
// DEPRECATED: use EchoRequest.Marshal instead.
func (e *EchoRequest) Serialize() ([]byte, error) {
	log.Println("EchoRequest.Serialize is deprecated. use EchoRequest.Marshal instead")
	return e.Marshal()
}

// SerializeTo serializes EchoRequest into bytes given as b.
//
// DEPRECATED: use EchoRequest.MarshalTo instead.
func (e *EchoRequest) SerializeTo(b []byte) error {
	log.Println("EchoRequest.SerializeTo is deprecated. use EchoRequest.MarshalTo instead")
	return e.MarshalTo(b)
}

// DecodeEchoRequest decodes bytes as EchoRequest.
//
// DEPRECATED: use ParseEchoRequest instead.
func DecodeEchoRequest(b []byte) (*EchoRequest, error) {
	log.Println("DecodeEchoRequest is deprecated. use ParseEchoRequest instead")
	return ParseEchoRequest(b)
}

// DecodeFromBytes decodes bytes as EchoRequest.
//
// DEPRECATED: use EchoRequest.UnmarshalBinary instead.
func (e *EchoRequest) DecodeFromBytes(b []byte) error {
	log.Println("EchoRequest.DecodeFromBytes is deprecated. use EchoRequest.UnmarshalBinary instead")
	return e.UnmarshalBinary(b)
}

// Len returns the actual length of EchoRequest.
//
// DEPRECATED: use EchoRequest.MarshalLen instead.
func (e *EchoRequest) Len() int {
	log.Println("EchoRequest.Len is deprecated. use EchoRequest.MarshalLen instead")
	return e.MarshalLen()
}
