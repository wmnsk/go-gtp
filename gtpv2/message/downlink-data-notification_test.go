// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv2"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
	"github.com/wmnsk/go-gtp/gtpv2/testutils"
)

func TestDownlinkDataNotification(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: message.NewDownlinkDataNotification(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
				ie.NewEPSBearerID(0x05),
				ie.NewAllocationRetensionPriority(1, 2, 1),
				ie.NewIMSI("123451234567890"),
				ie.NewFullyQualifiedTEID(gtpv2.IFTypeS4SGSNGTPU, 0xffffffff, "1.1.1.1", ""),
				ie.NewIndicationFromOctets(0xa1, 0x08),
				ie.NewPagingAndServiceInformation(5, 0x01, 0xff),
				ie.NewIntegerNumber(2020),
			),
			Serialized: []byte{
				// Header
				0x48, 0xb0, 0x00, 0x44, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
				// EBI
				0x49, 0x00, 0x01, 0x00, 0x05,
				// ARP
				0x9b, 0x00, 0x01, 0x00, 0x49,
				// IMSI
				0x01, 0x00, 0x08, 0x00, 0x21, 0x43, 0x15, 0x32, 0x54, 0x76, 0x98, 0xf0,
				// F-TEID
				0x57, 0x00, 0x09, 0x00, 0x8f, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x01,
				// Indication
				0x4d, 0x00, 0x02, 0x00, 0xa1, 0x08,
				// PagingAndServiceInformation
				0xba, 0x00, 0x03, 0x00, 0x05, 0x01, 0x7f,
				// DLDataPacketSize
				0xbb, 0x00, 0x02, 0x00, 0x07, 0xe4,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseDownlinkDataNotification(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
