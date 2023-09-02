// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes DeleteBearerRequest into bytes.
//
// Deprecated: use DeleteBearerRequest.Marshal instead.
func (d *DeleteBearerRequest) Serialize() ([]byte, error) {
	log.Println("DeleteBearerRequest.Serialize is deprecated. use DeleteBearerRequest.Marshal instead")
	return d.Marshal()
}

// SerializeTo serializes DeleteBearerRequest into bytes given as b.
//
// Deprecated: use DeleteBearerRequest.MarshalTo instead.
func (d *DeleteBearerRequest) SerializeTo(b []byte) error {
	log.Println("DeleteBearerRequest.SerializeTo is deprecated. use DeleteBearerRequest.MarshalTo instead")
	return d.MarshalTo(b)
}

// DecodeDeleteBearerRequest decodes bytes as DeleteBearerRequest.
//
// Deprecated: use ParseDeleteBearerRequest instead.
func DecodeDeleteBearerRequest(b []byte) (*DeleteBearerRequest, error) {
	log.Println("DecodeDeleteBearerRequest is deprecated. use ParseDeleteBearerRequest instead")
	return ParseDeleteBearerRequest(b)
}

// DecodeFromBytes decodes bytes as DeleteBearerRequest.
//
// Deprecated: use DeleteBearerRequest.UnmarshalBinary instead.
func (d *DeleteBearerRequest) DecodeFromBytes(b []byte) error {
	log.Println("DeleteBearerRequest.DecodeFromBytes is deprecated. use DeleteBearerRequest.UnmarshalBinary instead")
	return d.UnmarshalBinary(b)
}

// Len returns the actual length of DeleteBearerRequest.
//
// Deprecated: use DeleteBearerRequest.MarshalLen instead.
func (d *DeleteBearerRequest) Len() int {
	log.Println("DeleteBearerRequest.Len is deprecated. use DeleteBearerRequest.MarshalLen instead")
	return d.MarshalLen()
}
