// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/v2/ies"
	"github.com/wmnsk/go-gtp/v2/messages"
	"github.com/wmnsk/go-gtp/v2/testutils"
)

func TestGeneric(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: messages.NewGeneric(
				messages.MsgTypeEchoRequest,
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ies.NewRecovery(0x80),
			),
			Serialized: []byte{
				0x48, 0x01, 0x00, 0x0d, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00, 0x03, 0x00, 0x01, 0x00,
				0x80,
			},
		}, {
			Description: "No-TEID",
			Structured: messages.NewGeneric(
				messages.MsgTypeEchoRequest,
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ies.NewRecovery(0x80),
			),
			Serialized: []byte{
				0x48, 0x01, 0x00, 0x0d, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00, 0x03, 0x00, 0x01, 0x00,
				0x80,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := messages.ParseGeneric(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
