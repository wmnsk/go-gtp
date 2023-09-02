// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes VersionNotSupportedIndication into bytes.
//
// Deprecated: use VersionNotSupportedIndication.Marshal instead.
func (v *VersionNotSupportedIndication) Serialize() ([]byte, error) {
	log.Println("VersionNotSupportedIndication.Serialize is deprecated. use VersionNotSupportedIndication.Marshal instead")
	return v.Marshal()
}

// SerializeTo serializes VersionNotSupportedIndication into bytes given as b.
//
// Deprecated: use VersionNotSupportedIndication.MarshalTo instead.
func (v *VersionNotSupportedIndication) SerializeTo(b []byte) error {
	log.Println("VersionNotSupportedIndication.SerializeTo is deprecated. use VersionNotSupportedIndication.MarshalTo instead")
	return v.MarshalTo(b)
}

// DecodeVersionNotSupportedIndication decodes bytes as VersionNotSupportedIndication.
//
// Deprecated: use ParseVersionNotSupportedIndication instead.
func DecodeVersionNotSupportedIndication(b []byte) (*VersionNotSupportedIndication, error) {
	log.Println("DecodeVersionNotSupportedIndication is deprecated. use ParseVersionNotSupportedIndication instead")
	return ParseVersionNotSupportedIndication(b)
}

// DecodeFromBytes decodes bytes as VersionNotSupportedIndication.
//
// Deprecated: use VersionNotSupportedIndication.UnmarshalBinary instead.
func (v *VersionNotSupportedIndication) DecodeFromBytes(b []byte) error {
	log.Println("VersionNotSupportedIndication.DecodeFromBytes is deprecated. use VersionNotSupportedIndication.UnmarshalBinary instead")
	return v.UnmarshalBinary(b)
}

// Len returns the actual length of VersionNotSupportedIndication.
//
// Deprecated: use VersionNotSupportedIndication.MarshalLen instead.
func (v *VersionNotSupportedIndication) Len() int {
	log.Println("VersionNotSupportedIndication.Len is deprecated. use VersionNotSupportedIndication.MarshalLen instead")
	return v.MarshalLen()
}
