// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes DeletePDPContextRequest into bytes.
//
// Deprecated: use DeletePDPContextRequest.Marshal instead.
func (d *DeletePDPContextRequest) Serialize() ([]byte, error) {
	log.Println("DeletePDPContextRequest.Serialize is deprecated. use DeletePDPContextRequest.Marshal instead")
	return d.Marshal()
}

// SerializeTo serializes DeletePDPContextRequest into bytes given as b.
//
// Deprecated: use DeletePDPContextRequest.MarshalTo instead.
func (d *DeletePDPContextRequest) SerializeTo(b []byte) error {
	log.Println("DeletePDPContextRequest.SerializeTo is deprecated. use DeletePDPContextRequest.MarshalTo instead")
	return d.MarshalTo(b)
}

// DecodeDeletePDPContextRequest decodes bytes as DeletePDPContextRequest.
//
// Deprecated: use ParseDeletePDPContextRequest instead.
func DecodeDeletePDPContextRequest(b []byte) (*DeletePDPContextRequest, error) {
	log.Println("DecodeDeletePDPContextRequest is deprecated. use ParseDeletePDPContextRequest instead")
	return ParseDeletePDPContextRequest(b)
}

// DecodeFromBytes decodes bytes as DeletePDPContextRequest.
//
// Deprecated: use DeletePDPContextRequest.UnmarshalBinary instead.
func (d *DeletePDPContextRequest) DecodeFromBytes(b []byte) error {
	log.Println("DeletePDPContextRequest.DecodeFromBytes is deprecated. use DeletePDPContextRequest.UnmarshalBinary instead")
	return d.UnmarshalBinary(b)
}

// Len returns the actual length of DeletePDPContextRequest.
//
// Deprecated: use DeletePDPContextRequest.MarshalLen instead.
func (d *DeletePDPContextRequest) Len() int {
	log.Println("DeletePDPContextRequest.Len is deprecated. use DeletePDPContextRequest.MarshalLen instead")
	return d.MarshalLen()
}
