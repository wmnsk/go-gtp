// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv2"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
	"github.com/wmnsk/go-gtp/gtpv2/testutils"
)

func TestResumeAcknowledge(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Just Cause",
			Structured: message.NewResumeAcknowledge(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
			),
			Serialized: []byte{
				// Header
				0x48, 165, 0x00, 14, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
			},
		},
		{
			Description: "With Private Extension",
			Structured: message.NewResumeAcknowledge(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
				ie.NewPrivateExtension(234, []byte{2, 3, 4, 5, 6}),
			),
			Serialized: []byte{
				// Header
				0x48, 165, 0x00, 25, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
				// Private Extension
				0xff, 0x00, 0x07, 0x00, 0x00, 234, 2, 3, 4, 5, 6,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseResumeAcknowledge(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
