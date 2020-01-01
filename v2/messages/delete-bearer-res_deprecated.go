// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes DeleteBearerResponse into bytes.
//
// DEPRECATED: use DeleteBearerResponse.Marshal instead.
func (d *DeleteBearerResponse) Serialize() ([]byte, error) {
	log.Println("DeleteBearerResponse.Serialize is deprecated. use DeleteBearerResponse.Marshal instead")
	return d.Marshal()
}

// SerializeTo serializes DeleteBearerResponse into bytes given as b.
//
// DEPRECATED: use DeleteBearerResponse.MarshalTo instead.
func (d *DeleteBearerResponse) SerializeTo(b []byte) error {
	log.Println("DeleteBearerResponse.SerializeTo is deprecated. use DeleteBearerResponse.MarshalTo instead")
	return d.MarshalTo(b)
}

// DecodeDeleteBearerResponse decodes bytes as DeleteBearerResponse.
//
// DEPRECATED: use ParseDeleteBearerResponse instead.
func DecodeDeleteBearerResponse(b []byte) (*DeleteBearerResponse, error) {
	log.Println("DecodeDeleteBearerResponse is deprecated. use ParseDeleteBearerResponse instead")
	return ParseDeleteBearerResponse(b)
}

// DecodeFromBytes decodes bytes as DeleteBearerResponse.
//
// DEPRECATED: use DeleteBearerResponse.UnmarshalBinary instead.
func (d *DeleteBearerResponse) DecodeFromBytes(b []byte) error {
	log.Println("DeleteBearerResponse.DecodeFromBytes is deprecated. use DeleteBearerResponse.UnmarshalBinary instead")
	return d.UnmarshalBinary(b)
}

// Len returns the actual length of DeleteBearerResponse.
//
// DEPRECATED: use DeleteBearerResponse.MarshalLen instead.
func (d *DeleteBearerResponse) Len() int {
	log.Println("DeleteBearerResponse.Len is deprecated. use DeleteBearerResponse.MarshalLen instead")
	return d.MarshalLen()
}
