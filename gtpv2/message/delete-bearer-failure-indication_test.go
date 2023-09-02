// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"
	"time"

	"github.com/wmnsk/go-gtp/gtpv2"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
	"github.com/wmnsk/go-gtp/gtpv2/testutils"
)

func TestDeleteBearerFailureIndication(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: message.NewDeleteBearerFailureIndication(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
				ie.NewBearerContext(ie.NewEPSBearerID(5), ie.NewDelayValue(500*time.Millisecond), ie.NewDelayValue(100*time.Millisecond)),
				ie.NewBearerContext(ie.NewEPSBearerID(6), ie.NewDelayValue(500*time.Millisecond), ie.NewDelayValue(100*time.Millisecond)),
			),
			Serialized: []byte{
				// Header
				0x48, 0x43, 0x00, 0x34, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
				// BearerContexts 1
				0x5d, 0x00, 0x0f, 0x00,
				0x49, 0x00, 0x01, 0x00, 0x05,
				0x5c, 0x00, 0x01, 0x00, 0x0a,
				0x5c, 0x00, 0x01, 0x00, 0x02,
				// BearerContexts 2
				0x5d, 0x00, 0x0f, 0x00,
				0x49, 0x00, 0x01, 0x00, 0x06,
				0x5c, 0x00, 0x01, 0x00, 0x0a,
				0x5c, 0x00, 0x01, 0x00, 0x02,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseDeleteBearerFailureIndication(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
