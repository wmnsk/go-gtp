// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages_test

import (
	"testing"

	"github.com/ErvinsK/go-gtp/v2/ies"
	"github.com/ErvinsK/go-gtp/v2/messages"
	"github.com/ErvinsK/go-gtp/v2/testutils"
)

func TestStopPagingIndication(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal/CauseOnly",
			Structured: messages.NewStopPagingIndication(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ies.NewIMSI("123451234567890"),
			),
			Serialized: []byte{
				// Header
				0x48, 0x49, 0x00, 0x14, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// IMSI
				0x01, 0x00, 0x08, 0x00, 0x21, 0x43, 0x15, 0x32, 0x54, 0x76, 0x98, 0xf0,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := messages.ParseStopPagingIndication(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
