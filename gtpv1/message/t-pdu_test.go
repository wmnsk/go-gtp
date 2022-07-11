// Copyright 2019-2022 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv1/message"
	"github.com/wmnsk/go-gtp/gtpv1/testutils"
)

func TestTPDU(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured:  message.NewTPDU(0xdeadbeef, []byte{0xde, 0xad, 0xbe, 0xef}),
			Serialized: []byte{
				0x30, 0xff, 0x00, 0x04, 0xde, 0xad, 0xbe, 0xef,
				0xde, 0xad, 0xbe, 0xef,
			},
		}, {
			Description: "With-Sequence",
			Structured:  message.NewTPDUWithSequence(0xdeadbeef, 0x0001, []byte{0xde, 0xad, 0xbe, 0xef}),
			Serialized: []byte{
				0x32, 0xff, 0x00, 0x08, 0xde, 0xad, 0xbe, 0xef,
				0x00, 0x01, 0x00, 0x00, 0xde, 0xad, 0xbe, 0xef,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseTPDU(b)
		if err != nil {
			return nil, err
		}
		return v, nil
	})
}
