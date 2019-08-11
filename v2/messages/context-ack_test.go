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

func TestContextAcknowledge(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: messages.NewContextAcknowledge(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ies.NewCause(v2.CauseRequestAccepted, 0, 0, 0, nil),
				ies.NewFullyQualifiedTEID(v2.IFTypeS4SGSNGTPU, 0xffffffff, "1.1.1.1", ""),
			),
			Marshald: []byte{
				// Header
				0x48, 0x84, 0x00, 0x1b, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
				// F-TEID
				0x57, 0x00, 0x09, 0x00, 0x8f, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x01,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Marshalable, error) {
		v, err := messages.ParseContextAcknowledge(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
