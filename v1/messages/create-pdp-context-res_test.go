// Copyright 2019 go-gtp authors. All rights reserved.
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

func TestCreatePDPContextResponse(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: messages.NewCreatePDPContextResponse(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ies.NewCause(v1.ResCauseRequestAccepted),
				ies.NewReorderingRequired(false),
				ies.NewRecovery(0),
				ies.NewTEIDDataI(0xdeadbeef),
				ies.NewTEIDCPlane(0xdeadbeef),
				ies.NewEndUserAddress("10.10.10.10"),
				ies.NewGSNAddress("1.1.1.1"),
				ies.NewGSNAddress("2.2.2.2"),
			),
			Serialized: []byte{
				// Header
				0x32, 0x11, 0x00, 0x2b, 0x11, 0x22, 0x33, 0x44,
				0x00, 0x01, 0x00, 0x00,
				// Cause
				0x01, 0x80,
				// ReorderingRequired
				0x08, 0xfe,
				// Recovery
				0x0e, 0x00,
				// TEID-U
				0x10, 0xde, 0xad, 0xbe, 0xef,
				// TEID-C
				0x11, 0xde, 0xad, 0xbe, 0xef,
				// End User Address
				0x80, 0x00, 0x06, 0xf1, 0x21, 0x0a, 0x0a, 0x0a, 0x0a,
				// GSN Address
				0x85, 0x00, 0x04, 0x01, 0x01, 0x01, 0x01,
				// GSN Address
				0x85, 0x00, 0x04, 0x02, 0x02, 0x02, 0x02,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := messages.ParseCreatePDPContextResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
