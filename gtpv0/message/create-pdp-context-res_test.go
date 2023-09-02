// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv0"
	"github.com/wmnsk/go-gtp/gtpv0/ie"
	"github.com/wmnsk/go-gtp/gtpv0/message"
	"github.com/wmnsk/go-gtp/gtpv0/testutils"
)

func TestCreatePDPContextResponse(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "request-accepted",
			Structured: message.NewCreatePDPContextResponse(
				testutils.TestFlow.Seq, testutils.TestFlow.Label, testutils.TestFlow.TID,
				ie.NewCause(gtpv0.CauseRequestAccepted),
				ie.NewQualityOfServiceProfile(1, 1, 1, 1, 1),
				ie.NewReorderingRequired(false),
				ie.NewFlowLabelDataI(11),
				ie.NewFlowLabelSignalling(22),
				ie.NewChargingID(0xff),
				ie.NewEndUserAddress("1.1.1.1"),
				ie.NewGSNAddress("2.2.2.2"),
				ie.NewGSNAddress("3.3.3.3"),
			),
			Serialized: []byte{
				// Header
				0x1e, 0x11, 0x00, 0x2a,
				// SequenceNumber
				0x00, 0x01, 0x00, 0x00,
				// Sndpd
				0xff, 0xff, 0xff, 0xff,
				// TID
				0x21, 0x43, 0x65, 0x87, 0x09, 0x21, 0x43, 0x55,
				// Cause
				0x01, 0x80,
				// QualityOfServiceProfile
				0x06, 0x09, 0x11, 0x01,
				// ReorderingRequired
				0x08, 0xfe,
				// FlowLabelDataI
				0x10, 0x00, 0x0b,
				// FlowLabelSignalling
				0x11, 0x00, 0x16,
				// ChargingID
				0x7f, 0x00, 0x00, 0x00, 0xff,
				// EndUserAddress
				0x80, 0x00, 0x06, 0xf1, 0x21, 0x01, 0x01, 0x01, 0x01,
				// SGSNAddressForSignalling
				0x85, 0x00, 0x04, 0x02, 0x02, 0x02, 0x02,
				// SGSNAddressForUserData
				0x85, 0x00, 0x04, 0x03, 0x03, 0x03, 0x03,
			},
		}, {
			Description: "no-resources",
			Structured: message.NewCreatePDPContextResponse(
				testutils.TestFlow.Seq, testutils.TestFlow.Label, testutils.TestFlow.TID,
				ie.NewCause(gtpv0.CauseNoResourcesAvailable),
			),
			Serialized: []byte{
				// Header
				0x1e, 0x11, 0x00, 0x02,
				// SequenceNumber
				0x00, 0x01, 0x00, 0x00,
				// Sndpd
				0xff, 0xff, 0xff, 0xff,
				// TID
				0x21, 0x43, 0x65, 0x87, 0x09, 0x21, 0x43, 0x55,
				// Cause
				0x01, 0xc7,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseCreatePDPContextResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
