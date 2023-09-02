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

func TestDownlinkDataNotificationAcknowledge(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: message.NewDownlinkDataNotificationAcknowledge(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
				ie.NewDelayValue(500*time.Millisecond),
				ie.NewRecovery(0xff),
				ie.NewThrottling(20*time.Hour, 80),
				ie.NewIMSI("123451234567890"),
				ie.NewEPCTimer(20*time.Hour),
				ie.NewIntegerNumber(2020),
			),
			Serialized: []byte{
				// Header
				0x48, 0xb1, 0x00, 0x35, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
				// DataNotificationDelay
				0x5c, 0x00, 0x01, 0x00, 0x0a,
				// Recovery
				0x03, 0x00, 0x01, 0x00, 0xff,
				// DLLowPriorityTrafficThrottling
				0x9a, 0x00, 0x02, 0x00, 0x82, 0x50,
				// IMSI
				0x01, 0x00, 0x08, 0x00, 0x21, 0x43, 0x15, 0x32, 0x54, 0x76, 0x98, 0xf0,
				// DLBufferingDuration
				0x9c, 0x00, 0x01, 0x00, 0x82,
				// DLBufferingSuggestedPacketCount
				0xbb, 0x00, 0x02, 0x00, 0x07, 0xe4,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseDownlinkDataNotificationAcknowledge(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
