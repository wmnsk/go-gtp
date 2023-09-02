// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
	"github.com/wmnsk/go-gtp/gtpv2/testutils"
)

func TestEchoRequest(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: message.NewEchoRequest(
				0,
				ie.NewRecovery(0x80),
				ie.NewNodeFeatures(0x01),
			),
			Serialized: []byte{
				0x40, 0x01, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00,
				// Recovery
				0x03, 0x00, 0x01, 0x00, 0x80,
				// Node Features
				0x98, 0x00, 0x01, 0x00, 0x01,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseEchoRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
