// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv2/message"
	"github.com/wmnsk/go-gtp/gtpv2/testutils"
)

func TestHeader(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: message.NewHeader(
				message.NewHeaderFlags(2, 0, 1),
				32,         // Message type
				0xffffffff, // TEID
				0xdadada,   // Sequence Number
				[]byte{ // Payload: IMSI IE
					0x01, 0x00, 0x08, 0x00, 0x21, 0x43, 0x15, 0x32, 0x54, 0x76, 0x98, 0xf0,
				},
			),
			Serialized: []byte{
				0x48, 0x20, 0x00, 0x14, 0xff, 0xff, 0xff, 0xff,
				0xda, 0xda, 0xda, 0x00,
				0x01, 0x00, 0x08, 0x00, 0x21, 0x43, 0x15, 0x32, 0x54, 0x76, 0x98, 0xf0,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseHeader(b)
		if err != nil {
			return nil, err
		}
		return v, nil
	})
}
