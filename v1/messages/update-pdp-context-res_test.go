// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages_test

import (
	"testing"

	v1 "github.com/wmnsk/go-gtp/v1"
	"github.com/wmnsk/go-gtp/v1/ies"
	"github.com/wmnsk/go-gtp/v1/messages"
	"github.com/wmnsk/go-gtp/v1/testutils"
)

func TestUpdatePDPContextResponse(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: messages.NewUpdatePDPContextResponse(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ies.NewCause(v1.ResCauseRequestAccepted),
				ies.NewRecovery(0),
				ies.NewTEIDDataI(0xdeadbeef),
				ies.NewTEIDCPlane(0xdeadbeef),
				ies.NewGSNAddress("1.1.1.1"),
				ies.NewGSNAddress("2.2.2.2"),
			),
			Serialized: []byte{
				// Header
				0x32, 0x13, 0x00, 0x20, 0x11, 0x22, 0x33, 0x44,
				0x00, 0x01, 0x00, 0x00,
				// Cause
				0x01, 0x80,
				// Recovery
				0x0e, 0x00,
				// TEID-U
				0x10, 0xde, 0xad, 0xbe, 0xef,
				// TEID-C
				0x11, 0xde, 0xad, 0xbe, 0xef,
				// GSN Address
				0x85, 0x00, 0x04, 0x01, 0x01, 0x01, 0x01,
				// GSN Address
				0x85, 0x00, 0x04, 0x02, 0x02, 0x02, 0x02,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := messages.ParseUpdatePDPContextResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
