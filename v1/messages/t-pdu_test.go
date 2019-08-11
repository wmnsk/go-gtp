// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/v1/messages"
	"github.com/wmnsk/go-gtp/v1/testutils"
)

func TestTPDU(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured:  messages.NewTPDU(0xdeadbeef, []byte{0xde, 0xad, 0xbe, 0xef}),
			Marshald: []byte{
				0x30, 0xff, 0x00, 0x04, 0xde, 0xad, 0xbe, 0xef,
				0xde, 0xad, 0xbe, 0xef,
			},
		}, {
			Description: "With-Sequence",
			Structured:  messages.NewTPDUWithSequence(0xdeadbeef, 0x0001, []byte{0xde, 0xad, 0xbe, 0xef}),
			Marshald: []byte{
				0x32, 0xff, 0x00, 0x08, 0xde, 0xad, 0xbe, 0xef,
				0x00, 0x01, 0x00, 0x00, 0xde, 0xad, 0xbe, 0xef,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Marshalable, error) {
		v, err := messages.ParseTPDU(b)
		if err != nil {
			return nil, err
		}
		return v, nil
	})
}
