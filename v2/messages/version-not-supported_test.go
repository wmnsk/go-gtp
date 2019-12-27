// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages_test

import (
	"testing"

	"github.com/ErvinsK/go-gtp/v2/messages"
	"github.com/ErvinsK/go-gtp/v2/testutils"
)

func TestVersionNotSupportedIndication(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal/CauseOnly",
			Structured:  messages.NewVersionNotSupportedIndication(testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq),
			Serialized: []byte{
				// Header
				0x48, 0x03, 0x00, 0x08, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := messages.ParseVersionNotSupportedIndication(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
