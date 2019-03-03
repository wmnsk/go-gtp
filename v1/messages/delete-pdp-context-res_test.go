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

func TestDeletePDPContextResponse(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: messages.NewDeletePDPContextResponse(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ies.NewCause(v1.ResCauseRequestAccepted),
			),
			Serialized: []byte{
				// Header
				0x32, 0x15, 0x00, 0x06, 0x11, 0x22, 0x33, 0x44,
				0x00, 0x01, 0x00, 0x00,
				// Cause
				0x01, 0x80,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializeable, error) {
		v, err := messages.DecodeDeletePDPContextResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
