// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/v2/ies"
	"github.com/wmnsk/go-gtp/v2/messages"
	"github.com/wmnsk/go-gtp/v2/testutils"
)

func TestReleaseAccessBearersRequest(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal/NoIE",
			Structured:  messages.NewReleaseAccessBearersRequest(testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq),
			Serialized: []byte{
				// Header
				0x48, 0xaa, 0x00, 0x08, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
			},
		}, {
			Description: "Normal/WithIndication",
			Structured: messages.NewReleaseAccessBearersRequest(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ies.NewIndicationFromOctets(0xa1, 0x08, 0x15, 0x10, 0x88, 0x81, 0x40),
			),
			Serialized: []byte{
				// Header
				0x48, 0xaa, 0x00, 0x13, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// Indication
				0x4d, 0x00, 0x07, 0x00, 0xa1, 0x08, 0x15, 0x10, 0x88, 0x81, 0x40,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := messages.ParseReleaseAccessBearersRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
