// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv0/message"
	"github.com/wmnsk/go-gtp/gtpv0/testutils"
)

func TestHeader(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "normal",
			Structured: message.NewHeader(
				message.HeaderFlags(
					0, // version
					1, // Protocol Type
					0, // N-PDU?
				), //Flags
				0x10, // Message type
				testutils.TestFlow.Seq, testutils.TestFlow.Label, testutils.TestFlow.TID,
				[]byte{ // Payload
					0xde, 0xad, 0xbe, 0xef,
				},
			),
			Serialized: []byte{
				// Flags
				0x1e,
				// MessageType
				0x10,
				// SequenceNumber
				0x00, 0x04, 0x00, 0x01,
				// FlowLabel
				0x00, 0x00,
				// SndcpNumber
				0xff, 0xff, 0xff, 0xff,
				// TID
				0x21, 0x43, 0x65, 0x87, 0x09, 0x21, 0x43, 0x55,
				// dummy Payload
				0xde, 0xad, 0xbe, 0xef,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseHeader(b)
		if err != nil {
			return nil, err
		}

		return v, nil
	})
}
