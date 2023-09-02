// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes DeletePDPContextResponse into bytes.
//
// Deprecated: use DeletePDPContextResponse.Marshal instead.
func (d *DeletePDPContextResponse) Serialize() ([]byte, error) {
	log.Println("DeletePDPContextResponse.Serialize is deprecated. use DeletePDPContextResponse.Marshal instead")
	return d.Marshal()
}

// SerializeTo serializes DeletePDPContextResponse into bytes given as b.
//
// Deprecated: use DeletePDPContextResponse.MarshalTo instead.
func (d *DeletePDPContextResponse) SerializeTo(b []byte) error {
	log.Println("DeletePDPContextResponse.SerializeTo is deprecated. use DeletePDPContextResponse.MarshalTo instead")
	return d.MarshalTo(b)
}

// DecodeDeletePDPContextResponse decodes bytes as DeletePDPContextResponse.
//
// Deprecated: use ParseDeletePDPContextResponse instead.
func DecodeDeletePDPContextResponse(b []byte) (*DeletePDPContextResponse, error) {
	log.Println("DecodeDeletePDPContextResponse is deprecated. use ParseDeletePDPContextResponse instead")
	return ParseDeletePDPContextResponse(b)
}

// DecodeFromBytes decodes bytes as DeletePDPContextResponse.
//
// Deprecated: use DeletePDPContextResponse.UnmarshalBinary instead.
func (d *DeletePDPContextResponse) DecodeFromBytes(b []byte) error {
	log.Println("DeletePDPContextResponse.DecodeFromBytes is deprecated. use DeletePDPContextResponse.UnmarshalBinary instead")
	return d.UnmarshalBinary(b)
}

// Len returns the actual length of DeletePDPContextResponse.
//
// Deprecated: use DeletePDPContextResponse.MarshalLen instead.
func (d *DeletePDPContextResponse) Len() int {
	log.Println("DeletePDPContextResponse.Len is deprecated. use DeletePDPContextResponse.MarshalLen instead")
	return d.MarshalLen()
}
