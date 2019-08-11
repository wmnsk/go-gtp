// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages_test

import (
	"testing"

	v0 "github.com/wmnsk/go-gtp/v0"
	"github.com/wmnsk/go-gtp/v0/ies"
	"github.com/wmnsk/go-gtp/v0/messages"
	"github.com/wmnsk/go-gtp/v0/testutils"
)

func TestCreatePDPContextRequest(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "mandatory-only",
			Structured: messages.NewCreatePDPContextRequest(
				testutils.TestFlow.Seq, testutils.TestFlow.Label, testutils.TestFlow.TID,
				ies.NewQualityOfServiceProfile(1, 1, 1, 1, 1),
				ies.NewSelectionMode(v0.SelectionModeMSorNetworkProvidedAPNSubscribedVerified),
				ies.NewFlowLabelDataI(11),
				ies.NewFlowLabelSignalling(22),
				ies.NewEndUserAddress("1.1.1.1"),
				ies.NewAccessPointName("some.apn.example"),
				ies.NewGSNAddress("2.2.2.2"),
				ies.NewGSNAddress("3.3.3.3"),
				ies.NewMSISDN("819012345678"),
			),
			Serialized: []byte{
				// Header
				0x1e, 0x10, 0x00, 0x41,
				// SequenceNumber
				0x00, 0x01, 0x00, 0x00,
				// Sndpd
				0xff, 0xff, 0xff, 0xff,
				// TID
				0x21, 0x43, 0x65, 0x87, 0x09, 0x21, 0x43, 0x55,
				// QualityOfServiceProfile
				0x06, 0x09, 0x11, 0x01,
				// SelectionMode
				0x0f, 0xf0,
				// FlowLabelDataI
				0x10, 0x00, 0x0b,
				// FlowLabelSignalling
				0x11, 0x00, 0x16,
				// EndUserAddress
				0x80, 0x00, 0x06, 0xf1, 0x21, 0x01, 0x01, 0x01, 0x01,
				// AccessPointName
				0x83, 0x00, 0x11, 0x04, 0x73, 0x6f, 0x6d, 0x65,
				0x03, 0x61, 0x70, 0x6e, 0x07, 0x65, 0x78, 0x61,
				0x6d, 0x70, 0x6c, 0x65,
				// SGSNAddressForSignalling
				0x85, 0x00, 0x04, 0x02, 0x02, 0x02, 0x02,
				// SGSNAddressForUserData
				0x85, 0x00, 0x04, 0x03, 0x03, 0x03, 0x03,
				// MSISDN
				0x86, 0x00, 0x07, 0x91, 0x18, 0x09, 0x21, 0x43,
				0x65, 0x87,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := messages.ParseCreatePDPContextRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
