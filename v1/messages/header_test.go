// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/v1/messages"
	"github.com/wmnsk/go-gtp/v1/testutils"
)

func TestHeader(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: messages.NewHeader(
				messages.NewHeaderFlags(
					1, // version
					1, // Protocol Type
					0, // Next Extension Header?
					1, // Sequence Number?
					0, // N-PDU Number?
				), //Flags
				0x10,       // Message type
				0xdeadbeef, // TEID
				0xcafe,     // Sequence Number
				[]byte{ // Payload
					0xde, 0xad, 0xbe, 0xef,
				},
			),
			Marshald: []byte{
				0x32, 0x10, 0x00, 0x08, 0xde, 0xad, 0xbe, 0xef,
				0xca, 0xfe, 0x00, 0x00, 0xde, 0xad, 0xbe, 0xef,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Marshalable, error) {
		v, err := messages.ParseHeader(b)
		if err != nil {
			return nil, err
		}
		return v, nil
	})
}
