// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv0/ie"
	"github.com/wmnsk/go-gtp/gtpv0/message"
	"github.com/wmnsk/go-gtp/gtpv0/testutils"
)

func TestGeneric(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "echo-req",
			Structured: message.NewGeneric(
				message.MsgTypeEchoRequest, testutils.TestFlow.Seq, testutils.TestFlow.Label, testutils.TestFlow.TID,
			),
			Serialized: []byte{
				// Header
				0x1e, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
				0xff, 0xff, 0xff, 0xff, 0x21, 0x43, 0x65, 0x87,
				0x09, 0x21, 0x43, 0x55,
			},
		}, {
			Description: "echo-res",
			Structured: message.NewGeneric(
				message.MsgTypeEchoResponse, testutils.TestFlow.Seq, testutils.TestFlow.Label, testutils.TestFlow.TID,
				ie.NewRecovery(0x80),
			),
			Serialized: []byte{
				// Hewader
				0x1e, 0x02, 0x00, 0x02, 0x00, 0x01, 0x00, 0x00,
				0xff, 0xff, 0xff, 0xff, 0x21, 0x43, 0x65, 0x87,
				0x09, 0x21, 0x43, 0x55,
				// Recovery
				0x0e, 0x80,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseGeneric(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
