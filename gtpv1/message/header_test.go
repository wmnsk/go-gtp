// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv1/message"
	"github.com/wmnsk/go-gtp/gtpv1/testutils"
)

func TestHeader(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "NoOptionals",
			Structured: message.NewHeader(
				message.NewHeaderFlags(
					1, // version
					1, // Protocol Type
					0, // Next Extension Header?
					0, // Sequence Number?
					0, // N-PDU Number?
				), // Flags
				message.MsgTypeTPDU, // Message type
				0xdeadbeef,          // TEID
				0x00,                // Sequence Number
				[]byte{ // Payload
					0xde, 0xad, 0xbe, 0xef,
				},
			),
			Serialized: []byte{
				0x30, 0xff, 0x00, 0x04, 0xde, 0xad, 0xbe, 0xef,
				0xde, 0xad, 0xbe, 0xef,
			},
		}, {
			Description: "WithSequenceNumber",
			Structured: message.NewHeader(
				message.NewHeaderFlags(
					1, // version
					1, // Protocol Type
					0, // Next Extension Header?
					1, // Sequence Number?
					0, // N-PDU Number?
				), // Flags
				message.MsgTypeEchoRequest, // Message type
				0xdeadbeef,                 // TEID
				0xcafe,                     // Sequence Number
				[]byte{ // Payload
					0xde, 0xad, 0xbe, 0xef,
				},
			),
			Serialized: []byte{
				0x32, 0x01, 0x00, 0x08, 0xde, 0xad, 0xbe, 0xef,
				0xca, 0xfe, 0x00, 0x00, 0xde, 0xad, 0xbe, 0xef,
			},
		}, {
			Description: "WithNPDUNumber",
			Structured: message.NewHeader(
				message.NewHeaderFlags(
					1, // version
					1, // Protocol Type
					0, // Next Extension Header?
					0, // Sequence Number?
					0, // N-PDU Number?: set to zero at first, set by With... method
				), // Flags
				message.MsgTypeEchoRequest, // Message type
				0xdeadbeef,                 // TEID
				0x00,                       // Sequence Number
				[]byte{ // Payload
					0xde, 0xad, 0xbe, 0xef,
				},
			).WithNPDUNumber(0xff),
			Serialized: []byte{
				0x31, 0x01, 0x00, 0x08, 0xde, 0xad, 0xbe, 0xef,
				0x00, 0x00, 0xff, 0x00, 0xde, 0xad, 0xbe, 0xef,
			},
		}, {
			Description: "WithExtensionHeaders",
			Structured: message.NewHeader(
				message.NewHeaderFlags(
					1, // version
					1, // Protocol Type
					0, // Next Extension Header?: set to zero at first, set by With... method
					0, // Sequence Number?
					0, // N-PDU Number?
				), // Flags
				message.MsgTypeEchoRequest, // Message type
				0xdeadbeef,                 // TEID
				0x00,                       // Sequence Number
				[]byte{ // Payload
					0xde, 0xad, 0xbe, 0xef,
				},
			).WithExtensionHeaders(
				message.NewExtensionHeader(
					message.ExtHeaderTypeUDPPort,
					[]byte{0x22, 0xb8},
					message.ExtHeaderTypeUDPPort,
				),
				message.NewExtensionHeader(
					message.ExtHeaderTypeUDPPort,
					[]byte{0x22, 0xb8},
					message.ExtHeaderTypeNoMoreExtensionHeaders,
				),
			),
			Serialized: []byte{
				0x34, 0x01, 0x00, 0x10, 0xde, 0xad, 0xbe, 0xef,
				0x00, 0x00, 0x00, 0x40, 0x01, 0x22, 0xb8, 0x40,
				0x01, 0x22, 0xb8, 0x00, 0xde, 0xad, 0xbe, 0xef,
			},
		}, {
			Description: "WithSequenceAndNPDUNumber",
			Structured: message.NewHeader(
				message.NewHeaderFlags(
					1, // version
					1, // Protocol Type
					0, // Next Extension Header?
					0, // Sequence Number?
					0, // N-PDU Number?: set to zero at first, set by With... method
				), // Flags
				message.MsgTypeEchoRequest, // Message type
				0xdeadbeef,                 // TEID
				0x00,                       // Sequence Number
				[]byte{ // Payload
					0xde, 0xad, 0xbe, 0xef,
				},
			).WithSequenceNumber(0xcafe).WithNPDUNumber(0xff),
			Serialized: []byte{
				0x33, 0x01, 0x00, 0x08, 0xde, 0xad, 0xbe, 0xef,
				0xca, 0xfe, 0xff, 0x00, 0xde, 0xad, 0xbe, 0xef,
			},
		}, {
			Description: "WithSequenceAndExtensionHeaders",
			Structured: message.NewHeader(
				message.NewHeaderFlags(
					1, // version
					1, // Protocol Type
					0, // Next Extension Header?: set to zero at first, set by With... method
					0, // Sequence Number?: set to zero at first, set by With... method
					0, // N-PDU Number?
				), // Flags
				message.MsgTypeEchoRequest, // Message type
				0xdeadbeef,                 // TEID
				0x00,                       // Sequence Number
				[]byte{ // Payload
					0xde, 0xad, 0xbe, 0xef,
				},
			).WithSequenceNumber(0xcafe).WithExtensionHeaders(
				message.NewExtensionHeader(
					message.ExtHeaderTypeUDPPort,
					[]byte{0x022, 0xb8},
					message.ExtHeaderTypeUDPPort,
				),
				message.NewExtensionHeader(
					message.ExtHeaderTypeUDPPort,
					[]byte{0x022, 0xb8},
					message.ExtHeaderTypeNoMoreExtensionHeaders,
				),
			),
			Serialized: []byte{
				0x36, 0x01, 0x00, 0x10, 0xde, 0xad, 0xbe, 0xef,
				0xca, 0xfe, 0x00, 0x40, 0x01, 0x22, 0xb8, 0x40,
				0x01, 0x22, 0xb8, 0x00, 0xde, 0xad, 0xbe, 0xef,
			},
		}, {
			Description: "WithSequenceAndNPDUAndExtensionHeaders",
			Structured: message.NewHeader(
				message.NewHeaderFlags(
					1, // version
					1, // Protocol Type
					0, // Next Extension Header?: set to zero at first, set by With... method
					0, // Sequence Number?: set to zero at first, set by With... method
					0, // N-PDU Number?: set to zero at first, set by With... method
				), // Flags
				message.MsgTypeEchoRequest, // Message type
				0xdeadbeef,                 // TEID
				0x00,                       // Sequence Number
				[]byte{ // Payload
					0xde, 0xad, 0xbe, 0xef,
				},
			).WithNPDUNumber(0xff).WithSequenceNumber(0xcafe).WithExtensionHeaders(
				message.NewExtensionHeader(
					message.ExtHeaderTypeUDPPort,
					[]byte{0x022, 0xb8},
					message.ExtHeaderTypeUDPPort,
				),
				message.NewExtensionHeader(
					message.ExtHeaderTypeUDPPort,
					[]byte{0x022, 0xb8},
					message.ExtHeaderTypeNoMoreExtensionHeaders,
				),
			),
			Serialized: []byte{
				0x37, 0x01, 0x00, 0x10, 0xde, 0xad, 0xbe, 0xef,
				0xca, 0xfe, 0xff, 0x40, 0x01, 0x22, 0xb8, 0x40,
				0x01, 0x22, 0xb8, 0x00, 0xde, 0xad, 0xbe, 0xef,
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

func TestHeaderErrorDetection(t *testing.T) {
	cases := []struct {
		Description string
		Serialized  []byte
		Error       error
	}{
		{
			Description: "ExtFlagSetButMissingExt",
			Serialized: []byte{
				0b00110110, // Version (3 bits), PT, (*), E, S, PN
				0x03,       // Message Type
				0x00, 0x04, // Length (2 octets)
				0x00, 0x00, 0x00, 0x00, // TEID (4 octets)
				0x00, 0x06, // Seqence Number (2 octets)
				0xff,       // N-PDU Number (to be ignored)
				0b11000001, // Next Extension Header Type
				// Extension Header would go here, but is missing
			},
			Error: message.ErrTooShortToParse, // Expect too short to parse error, as missing ext header has length 0
		},
		{
			Description: "IncorrectLength",
			Serialized: []byte{
				0b00110010, // Version (3 bits), PT, (*), E, S, PN
				0x03,       // Message Type
				0x30, 0x33, // Length (2 octets)
				0x00, 0x00, 0x00, 0x00, // TEID (4 octets)
				0x00, 0x06, // Seqence Number (2 octets)
				0xff,       // N-PDU Number (to be ignored)
				0b00000000, // Next Extension Header Type
			},
			Error: message.ErrTooShortToParse, // Expect too short to parse error, as missing ext header has length 0
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			_, err := message.ParseHeader(c.Serialized)
			if err != c.Error {
				t.Fatalf("got '%v' want '%v'", err, c.Error)
			}
		})
	}
}
