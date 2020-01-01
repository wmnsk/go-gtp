// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes DeleteSessionRequest into bytes.
//
// DEPRECATED: use DeleteSessionRequest.Marshal instead.
func (d *DeleteSessionRequest) Serialize() ([]byte, error) {
	log.Println("DeleteSessionRequest.Serialize is deprecated. use DeleteSessionRequest.Marshal instead")
	return d.Marshal()
}

// SerializeTo serializes DeleteSessionRequest into bytes given as b.
//
// DEPRECATED: use DeleteSessionRequest.MarshalTo instead.
func (d *DeleteSessionRequest) SerializeTo(b []byte) error {
	log.Println("DeleteSessionRequest.SerializeTo is deprecated. use DeleteSessionRequest.MarshalTo instead")
	return d.MarshalTo(b)
}

// DecodeDeleteSessionRequest decodes bytes as DeleteSessionRequest.
//
// DEPRECATED: use ParseDeleteSessionRequest instead.
func DecodeDeleteSessionRequest(b []byte) (*DeleteSessionRequest, error) {
	log.Println("DecodeDeleteSessionRequest is deprecated. use ParseDeleteSessionRequest instead")
	return ParseDeleteSessionRequest(b)
}

// DecodeFromBytes decodes bytes as DeleteSessionRequest.
//
// DEPRECATED: use DeleteSessionRequest.UnmarshalBinary instead.
func (d *DeleteSessionRequest) DecodeFromBytes(b []byte) error {
	log.Println("DeleteSessionRequest.DecodeFromBytes is deprecated. use DeleteSessionRequest.UnmarshalBinary instead")
	return d.UnmarshalBinary(b)
}

// Len returns the actual length of DeleteSessionRequest.
//
// DEPRECATED: use DeleteSessionRequest.MarshalLen instead.
func (d *DeleteSessionRequest) Len() int {
	log.Println("DeleteSessionRequest.Len is deprecated. use DeleteSessionRequest.MarshalLen instead")
	return d.MarshalLen()
}
