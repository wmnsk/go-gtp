// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv2"

	"github.com/wmnsk/go-gtp/gtpv2/message"
	"github.com/wmnsk/go-gtp/gtpv2/testutils"

	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

func TestContextAcknowledge(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: message.NewContextAcknowledge(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
				ie.NewFullyQualifiedTEID(gtpv2.IFTypeS4SGSNGTPU, 0xffffffff, "1.1.1.1", ""),
			),
			Serialized: []byte{
				// Header
				0x48, 0x84, 0x00, 0x1b, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
				// F-TEID
				0x57, 0x00, 0x09, 0x00, 0x8f, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x01,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseContextAcknowledge(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
