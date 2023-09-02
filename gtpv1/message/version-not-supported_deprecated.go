// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes VersionNotSupported into bytes.
//
// Deprecated: use VersionNotSupported.Marshal instead.
func (v *VersionNotSupported) Serialize() ([]byte, error) {
	log.Println("VersionNotSupported.Serialize is deprecated. use VersionNotSupported.Marshal instead")
	return v.Marshal()
}

// SerializeTo serializes VersionNotSupported into bytes given as b.
//
// Deprecated: use VersionNotSupported.MarshalTo instead.
func (v *VersionNotSupported) SerializeTo(b []byte) error {
	log.Println("VersionNotSupported.SerializeTo is deprecated. use VersionNotSupported.MarshalTo instead")
	return v.MarshalTo(b)
}

// DecodeVersionNotSupported decodes bytes as VersionNotSupported.
//
// Deprecated: use ParseVersionNotSupported instead.
func DecodeVersionNotSupported(b []byte) (*VersionNotSupported, error) {
	log.Println("DecodeVersionNotSupported is deprecated. use ParseVersionNotSupported instead")
	return ParseVersionNotSupported(b)
}

// DecodeFromBytes decodes bytes as VersionNotSupported.
//
// Deprecated: use VersionNotSupported.UnmarshalBinary instead.
func (v *VersionNotSupported) DecodeFromBytes(b []byte) error {
	log.Println("VersionNotSupported.DecodeFromBytes is deprecated. use VersionNotSupported.UnmarshalBinary instead")
	return v.UnmarshalBinary(b)
}

// Len returns the actual length of VersionNotSupported.
//
// Deprecated: use VersionNotSupported.MarshalLen instead.
func (v *VersionNotSupported) Len() int {
	log.Println("VersionNotSupported.Len is deprecated. use VersionNotSupported.MarshalLen instead")
	return v.MarshalLen()
}
