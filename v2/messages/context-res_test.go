// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages_test

import (
	"testing"

	v2 "github.com/wmnsk/go-gtp/v2"

	"github.com/wmnsk/go-gtp/v2/messages"
	"github.com/wmnsk/go-gtp/v2/testutils"

	"github.com/wmnsk/go-gtp/v2/ies"
)

func TestContextResponse(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: messages.NewContextResponse(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ies.NewCause(v2.CauseRequestAccepted, 0, 0, 0, nil),
				ies.NewIMSI("123451234567890"),
				// ies.NewMMContext(),  XXX- implement!
				ies.NewFullyQualifiedTEID(v2.IFTypeS10MMEGTPC, 0xffffffff, "1.1.1.1", ""),
			),
			Serialized: []byte{
				// Header
				0x48, 0x83, 0x00, 0x27, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
				// IMSI
				0x01, 0x00, 0x08, 0x00, 0x21, 0x43, 0x15, 0x32, 0x54, 0x76, 0x98, 0xf0,
				// F-TEID
				0x57, 0x00, 0x09, 0x00, 0x8c, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x01,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializeable, error) {
		v, err := messages.DecodeContextResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
