// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"
	"time"

	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
	"github.com/wmnsk/go-gtp/gtpv2/testutils"
)

func TestModifyBearerCommand(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: message.NewModifyBearerCommand(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewAggregateMaximumBitRate(0x11111111, 0x22222222),
				ie.NewBearerContext(ie.NewDelayValue(500*time.Millisecond), ie.NewDelayValue(100*time.Millisecond)),
			),
			Serialized: []byte{
				// Header
				0x48, 0x40, 0x00, 0x22, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// APN-AMBR
				0x48, 0x00, 0x08, 0x00, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22,
				// BearerContext
				0x5d, 0x00, 0x0a, 0x00, 0x5c, 0x00, 0x01, 0x00, 0x0a, 0x5c, 0x00, 0x01, 0x00, 0x02,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseModifyBearerCommand(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
