// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes Message into bytes.
//
// DEPRECATED: use Marshal instead.
func Serialize(m Message) ([]byte, error) {
	log.Println("Serialize is deprecated. use Marshal instead")
	return Marshal(m)
}

// Decode decodes bytes as Message.
//
// DEPRECATED: use Parse instead.
func Decode(b []byte) (Message, error) {
	log.Println("Decode is deprecated. use Parse instead")
	return Parse(b)
}
