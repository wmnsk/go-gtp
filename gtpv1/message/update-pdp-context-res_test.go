// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv1"
	"github.com/wmnsk/go-gtp/gtpv1/ie"
	"github.com/wmnsk/go-gtp/gtpv1/message"
	"github.com/wmnsk/go-gtp/gtpv1/testutils"
)

func TestUpdatePDPContextResponse(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: message.NewUpdatePDPContextResponse(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewCause(gtpv1.ResCauseRequestAccepted),
				ie.NewRecovery(0),
				ie.NewTEIDDataI(0xdeadbeef),
				ie.NewTEIDCPlane(0xdeadbeef),
				ie.NewGSNAddress("1.1.1.1"),
				ie.NewGSNAddress("2.2.2.2"),
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
		v, err := message.ParseUpdatePDPContextResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
