package gtpv2_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv2"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
	"github.com/wmnsk/go-gtp/gtpv2/testutils"
)

var nbLoop = 1
var debug = false

func BenchmarkEchoRequestNormal(b *testing.B) {

	b.Log("BenchmarkEchoRequestNormal")
	msg := message.NewEchoRequest(0, ie.NewRecovery(0x80), ie.NewNodeFeatures(0x01))
	i := 0

	for i < nbLoop {
		/*payload*/
		_, err := msg.Marshal()
		if err != nil {
			b.Errorf("could not Marshal EchoRequest : %s", err)
		}
		i++
	}

}

func BenchmarkEchoRequestReflect(b *testing.B) {
	b.Log("BenchmarkEchoRequestReflect")
	msg1 := message.NewEchoRequest(0, ie.NewRecovery(0x80), ie.NewNodeFeatures(0x01))
	i := 0
	for i < nbLoop {

		Totlen1 := msg1.MarshalLen1(debug)
		Payloadtmp := make([]byte, Totlen1)
		err := msg1.MarshalTo1(Payloadtmp, int(Totlen1), debug)
		if err != nil {
			b.Errorf("could not Marshal EchoRequest : %s", err)
		}

		i++

	}

}

/////////////////////////////////////////////////////////////////////////////

func BenchmarkCreateSessionRequestNormal(b *testing.B) {

	b.Log("BenchmarkCreateSessionRequestNormal")

	i := 0

	msg := message.NewCreateSessionRequest(
		testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
		ie.NewIMSI("123451234567890"),
		ie.NewMSISDN("123450123456789"),
		ie.NewAccessPointName("some.apn.example"),
		ie.NewFullyQualifiedTEID(gtpv2.IFTypeS11MMEGTPC, 0xffffffff, "1.1.1.1", ""),
		ie.NewFullyQualifiedTEID(gtpv2.IFTypeS5S8PGWGTPC, 0xffffffff, "1.1.1.2", "").WithInstance(1),
		ie.NewPDNType(gtpv2.PDNTypeIPv4),
		ie.NewAggregateMaximumBitRate(0x11111111, 0x22222222),
		ie.NewIndicationFromOctets(0xa1, 0x08, 0x15, 0x10, 0x88, 0x81, 0x40),
		ie.NewBearerContext(
			ie.NewEPSBearerID(0x05),
			ie.NewBearerQoS(1, 2, 1, 0xff, 0x1111111111, 0x2222222222, 0x1111111111, 0x2222222222),
		),
		ie.NewBearerContext(
			ie.NewEPSBearerID(0x06),
			ie.NewBearerQoS(1, 2, 1, 0xff, 0x1111111111, 0x2222222222, 0x1111111111, 0x2222222222),
		),
		ie.NewMobileEquipmentIdentity("123450123456789"),
		ie.NewServingNetwork("123", "45"),
		ie.NewPDNAddressAllocation("2.2.2.2"),
		ie.NewAPNRestriction(gtpv2.APNRestrictionPublic1),
		ie.NewUserLocationInformationStruct(
			nil, nil, nil, ie.NewTAI("123", "45", 0x0001),
			ie.NewECGI("123", "45", 0x00000101), nil, nil, nil,
		),
		ie.NewRATType(gtpv2.RATTypeEUTRAN),
		ie.NewSelectionMode(gtpv2.SelectionModeMSorNetworkProvidedAPNSubscribedVerified),
	)

	for i < nbLoop {
		/*payload*/
		_, err := msg.Marshal()
		if err != nil {
			b.Errorf("could not Marshal CreateSessionRequest : %s", err)
		}
		i++
	}

}

func BenchmarkCreateSessionRequestReflect(b *testing.B) {
	b.Log("BenchmarkCreateSessionRequestReflect")
	i := 0

	msg1 := message.NewCreateSessionRequest(
		testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
		ie.NewIMSI("123451234567890"),
		ie.NewMSISDN("123450123456789"),
		ie.NewAccessPointName("some.apn.example"),
		ie.NewFullyQualifiedTEID(gtpv2.IFTypeS11MMEGTPC, 0xffffffff, "1.1.1.1", ""),
		ie.NewFullyQualifiedTEID(gtpv2.IFTypeS5S8PGWGTPC, 0xffffffff, "1.1.1.2", "").WithInstance(1),
		ie.NewPDNType(gtpv2.PDNTypeIPv4),
		ie.NewAggregateMaximumBitRate(0x11111111, 0x22222222),
		ie.NewIndicationFromOctets(0xa1, 0x08, 0x15, 0x10, 0x88, 0x81, 0x40),
		ie.NewBearerContext(
			ie.NewEPSBearerID(0x05),
			ie.NewBearerQoS(1, 2, 1, 0xff, 0x1111111111, 0x2222222222, 0x1111111111, 0x2222222222),
		),
		ie.NewBearerContext(
			ie.NewEPSBearerID(0x06),
			ie.NewBearerQoS(1, 2, 1, 0xff, 0x1111111111, 0x2222222222, 0x1111111111, 0x2222222222),
		),
		ie.NewMobileEquipmentIdentity("123450123456789"),
		ie.NewServingNetwork("123", "45"),
		ie.NewPDNAddressAllocation("2.2.2.2"),
		ie.NewAPNRestriction(gtpv2.APNRestrictionPublic1),
		ie.NewUserLocationInformationStruct(
			nil, nil, nil, ie.NewTAI("123", "45", 0x0001),
			ie.NewECGI("123", "45", 0x00000101), nil, nil, nil,
		),
		ie.NewRATType(gtpv2.RATTypeEUTRAN),
		ie.NewSelectionMode(gtpv2.SelectionModeMSorNetworkProvidedAPNSubscribedVerified),
	)
	for i < nbLoop {

		Totlen1 := msg1.MarshalLen1(debug)
		Payloadtmp := make([]byte, Totlen1)
		err := msg1.MarshalTo1(Payloadtmp, int(Totlen1), debug)
		if err != nil {
			b.Errorf("could not Marshal CreateSessionRequest : %s", err)
		}

		i++

	}

}
