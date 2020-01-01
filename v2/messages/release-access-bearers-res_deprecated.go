// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes ReleaseAccessBearersResponse into bytes.
//
// DEPRECATED: use ReleaseAccessBearersResponse.Marshal instead.
func (r *ReleaseAccessBearersResponse) Serialize() ([]byte, error) {
	log.Println("ReleaseAccessBearersResponse.Serialize is deprecated. use ReleaseAccessBearersResponse.Marshal instead")
	return r.Marshal()
}

// SerializeTo serializes ReleaseAccessBearersResponse into bytes given as b.
//
// DEPRECATED: use ReleaseAccessBearersResponse.MarshalTo instead.
func (r *ReleaseAccessBearersResponse) SerializeTo(b []byte) error {
	log.Println("ReleaseAccessBearersResponse.SerializeTo is deprecated. use ReleaseAccessBearersResponse.MarshalTo instead")
	return r.MarshalTo(b)
}

// DecodeReleaseAccessBearersResponse decodes bytes as ReleaseAccessBearersResponse.
//
// DEPRECATED: use ParseReleaseAccessBearersResponse instead.
func DecodeReleaseAccessBearersResponse(b []byte) (*ReleaseAccessBearersResponse, error) {
	log.Println("DecodeReleaseAccessBearersResponse is deprecated. use ParseReleaseAccessBearersResponse instead")
	return ParseReleaseAccessBearersResponse(b)
}

// DecodeFromBytes decodes bytes as ReleaseAccessBearersResponse.
//
// DEPRECATED: use ReleaseAccessBearersResponse.UnmarshalBinary instead.
func (r *ReleaseAccessBearersResponse) DecodeFromBytes(b []byte) error {
	log.Println("ReleaseAccessBearersResponse.DecodeFromBytes is deprecated. use ReleaseAccessBearersResponse.UnmarshalBinary instead")
	return r.UnmarshalBinary(b)
}

// Len returns the actual length of ReleaseAccessBearersResponse.
//
// DEPRECATED: use ReleaseAccessBearersResponse.MarshalLen instead.
func (r *ReleaseAccessBearersResponse) Len() int {
	log.Println("ReleaseAccessBearersResponse.Len is deprecated. use ReleaseAccessBearersResponse.MarshalLen instead")
	return r.MarshalLen()
}
