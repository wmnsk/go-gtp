// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages_test

import (
	"testing"
	"time"

	v2 "github.com/wmnsk/go-gtp/v2"
	"github.com/wmnsk/go-gtp/v2/ies"
	"github.com/wmnsk/go-gtp/v2/messages"
	"github.com/wmnsk/go-gtp/v2/testutils"
)

func TestDeleteBearerFailureIndication(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: messages.NewDeleteBearerFailureIndication(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ies.NewCause(v2.CauseRequestAccepted, 0, 0, 0, nil),
				ies.NewBearerContext(ies.NewDelayValue(500*time.Millisecond), ies.NewDelayValue(100*time.Millisecond)),
			),
			Serialized: []byte{
				// Header
				0x48, 0x43, 0x00, 0x1c, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
				// BearerContexts
				0x5d, 0x00, 0x0a, 0x00, 0x5c, 0x00, 0x01, 0x00, 0x0a, 0x5c, 0x00, 0x01, 0x00, 0x02,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := messages.ParseDeleteBearerFailureIndication(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
