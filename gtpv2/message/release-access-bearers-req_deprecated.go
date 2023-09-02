// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes ReleaseAccessBearersRequest into bytes.
//
// Deprecated: use ReleaseAccessBearersRequest.Marshal instead.
func (r *ReleaseAccessBearersRequest) Serialize() ([]byte, error) {
	log.Println("ReleaseAccessBearersRequest.Serialize is deprecated. use ReleaseAccessBearersRequest.Marshal instead")
	return r.Marshal()
}

// SerializeTo serializes ReleaseAccessBearersRequest into bytes given as b.
//
// Deprecated: use ReleaseAccessBearersRequest.MarshalTo instead.
func (r *ReleaseAccessBearersRequest) SerializeTo(b []byte) error {
	log.Println("ReleaseAccessBearersRequest.SerializeTo is deprecated. use ReleaseAccessBearersRequest.MarshalTo instead")
	return r.MarshalTo(b)
}

// DecodeReleaseAccessBearersRequest decodes bytes as ReleaseAccessBearersRequest.
//
// Deprecated: use ParseReleaseAccessBearersRequest instead.
func DecodeReleaseAccessBearersRequest(b []byte) (*ReleaseAccessBearersRequest, error) {
	log.Println("DecodeReleaseAccessBearersRequest is deprecated. use ParseReleaseAccessBearersRequest instead")
	return ParseReleaseAccessBearersRequest(b)
}

// DecodeFromBytes decodes bytes as ReleaseAccessBearersRequest.
//
// Deprecated: use ReleaseAccessBearersRequest.UnmarshalBinary instead.
func (r *ReleaseAccessBearersRequest) DecodeFromBytes(b []byte) error {
	log.Println("ReleaseAccessBearersRequest.DecodeFromBytes is deprecated. use ReleaseAccessBearersRequest.UnmarshalBinary instead")
	return r.UnmarshalBinary(b)
}

// Len returns the actual length of ReleaseAccessBearersRequest.
//
// Deprecated: use ReleaseAccessBearersRequest.MarshalLen instead.
func (r *ReleaseAccessBearersRequest) Len() int {
	log.Println("ReleaseAccessBearersRequest.Len is deprecated. use ReleaseAccessBearersRequest.MarshalLen instead")
	return r.MarshalLen()
}
