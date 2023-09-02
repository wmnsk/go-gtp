// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes Message into bytes.
//
// Deprecated: use Marshal instead.
func Serialize(m Message) ([]byte, error) {
	log.Println("Serialize is deprecated. use Marshal instead")
	return Marshal(m)
}

// Decode decodes bytes as Message.
//
// Deprecated: use Parse instead.
func Decode(b []byte) (Message, error) {
	log.Println("Decode is deprecated. use Parse instead")
	return Parse(b)
}
