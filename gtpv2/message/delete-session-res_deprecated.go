// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes DeleteSessionResponse into bytes.
//
// Deprecated: use DeleteSessionResponse.Marshal instead.
func (d *DeleteSessionResponse) Serialize() ([]byte, error) {
	log.Println("DeleteSessionResponse.Serialize is deprecated. use DeleteSessionResponse.Marshal instead")
	return d.Marshal()
}

// SerializeTo serializes DeleteSessionResponse into bytes given as b.
//
// Deprecated: use DeleteSessionResponse.MarshalTo instead.
func (d *DeleteSessionResponse) SerializeTo(b []byte) error {
	log.Println("DeleteSessionResponse.SerializeTo is deprecated. use DeleteSessionResponse.MarshalTo instead")
	return d.MarshalTo(b)
}

// DecodeDeleteSessionResponse decodes bytes as DeleteSessionResponse.
//
// Deprecated: use ParseDeleteSessionResponse instead.
func DecodeDeleteSessionResponse(b []byte) (*DeleteSessionResponse, error) {
	log.Println("DecodeDeleteSessionResponse is deprecated. use ParseDeleteSessionResponse instead")
	return ParseDeleteSessionResponse(b)
}

// DecodeFromBytes decodes bytes as DeleteSessionResponse.
//
// Deprecated: use DeleteSessionResponse.UnmarshalBinary instead.
func (d *DeleteSessionResponse) DecodeFromBytes(b []byte) error {
	log.Println("DeleteSessionResponse.DecodeFromBytes is deprecated. use DeleteSessionResponse.UnmarshalBinary instead")
	return d.UnmarshalBinary(b)
}

// Len returns the actual length of DeleteSessionResponse.
//
// Deprecated: use DeleteSessionResponse.MarshalLen instead.
func (d *DeleteSessionResponse) Len() int {
	log.Println("DeleteSessionResponse.Len is deprecated. use DeleteSessionResponse.MarshalLen instead")
	return d.MarshalLen()
}
