// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

/*
Package ie provides encoding/decoding feature of GTPv1 Information Elements.
*/
package ie

import (
	"encoding/binary"
	"fmt"
)

// TV IE definitions.
const (
	Cause                        uint8 = 1
	IMSI                         uint8 = 2
	RouteingAreaIdentity         uint8 = 3
	TemporaryLogicalLinkIdentity uint8 = 4
	PacketTMSI                   uint8 = 5
	ReorderingRequired           uint8 = 8
	AuthenticationTriplet        uint8 = 9
	MAPCause                     uint8 = 11
	PTMSISignature               uint8 = 12
	MSValidated                  uint8 = 13
	Recovery                     uint8 = 14
	SelectionMode                uint8 = 15
	TEIDDataI                    uint8 = 16
	TEIDCPlane                   uint8 = 17
	TEIDDataII                   uint8 = 18
	TeardownInd                  uint8 = 19
	NSAPI                        uint8 = 20
	RANAPCause                   uint8 = 21
	RABContext                   uint8 = 22
	RadioPrioritySMS             uint8 = 23
	RadioPriority                uint8 = 24
	PacketFlowID                 uint8 = 25
	ChargingCharacteristics      uint8 = 26
	TraceReference               uint8 = 27
	TraceType                    uint8 = 28
	MSNotReachableReason         uint8 = 29
	ChargingID                   uint8 = 127
)

// TLV IE definitions.
const (
	EndUserAddress                        uint8 = 128
	MMContext                             uint8 = 129
	PDPContext                            uint8 = 130
	AccessPointName                       uint8 = 131
	ProtocolConfigurationOptions          uint8 = 132
	GSNAddress                            uint8 = 133
	MSISDN                                uint8 = 134
	QoSProfile                            uint8 = 135
	AuthenticationQuintuplet              uint8 = 136
	TrafficFlowTemplate                   uint8 = 137
	TargetIdentification                  uint8 = 138
	UTRANTransparentContainer             uint8 = 139
	RABSetupInformation                   uint8 = 140
	ExtensionHeaderTypeList               uint8 = 141
	TriggerID                             uint8 = 142
	OMCIdentity                           uint8 = 143
	RANTransparentContainer               uint8 = 144
	PDPContextPrioritization              uint8 = 145
	AdditionalRABSetupInformation         uint8 = 146
	SGSNNumber                            uint8 = 147
	CommonFlags                           uint8 = 148
	APNRestriction                        uint8 = 149
	RadioPriorityLCS                      uint8 = 150
	RATType                               uint8 = 151
	UserLocationInformation               uint8 = 152
	MSTimeZone                            uint8 = 153
	IMEISV                                uint8 = 154
	CAMELChargingInformationContainer     uint8 = 155
	MBMSUEContext                         uint8 = 156
	TemporaryMobileGroupIdentity          uint8 = 157
	RIMRoutingAddress                     uint8 = 158
	MBMSProtocolConfigurationOptions      uint8 = 159
	MBMSServiceArea                       uint8 = 160
	SourceRNCPDCPContextInfo              uint8 = 161
	AdditionalTraceInfo                   uint8 = 162
	HopCounter                            uint8 = 163
	SelectedPLMNID                        uint8 = 164
	MBMSSessionIdentifier                 uint8 = 165
	MBMS2G3GIndicator                     uint8 = 166
	EnhancedNSAPI                         uint8 = 167
	MBMSSessionDuration                   uint8 = 168
	AdditionalMBMSTraceInfo               uint8 = 169
	MBMSSessionRepetitionNumber           uint8 = 170
	MBMSTimeToDataTransfer                uint8 = 171
	BSSContainer                          uint8 = 173
	CellIdentification                    uint8 = 174
	PDUNumbers                            uint8 = 175
	BSSGPCause                            uint8 = 176
	RequiredMBMSBearerCapabilities        uint8 = 177
	RIMRoutingAddressDiscriminator        uint8 = 178
	ListOfSetupPFCs                       uint8 = 179
	PSHandoverXIDParameters               uint8 = 180
	MSInfoChangeReportingAction           uint8 = 181
	DirectTunnelFlags                     uint8 = 182
	CorrelationID                         uint8 = 183
	BearerControlMode                     uint8 = 184
	MBMSFlowIdentifier                    uint8 = 185
	MBMSIPMulticastDistribution           uint8 = 186
	MBMSDistributionAcknowledgement       uint8 = 187
	ReliableInterRATHandoverInfo          uint8 = 188
	RFSPIndex                             uint8 = 189
	FullyQualifiedDomainName              uint8 = 190
	EvolvedAllocationRetentionPriorityI   uint8 = 191
	EvolvedAllocationRetentionPriorityII  uint8 = 192
	ExtendedCommonFlags                   uint8 = 193
	UserCSGInformation                    uint8 = 194
	CSGInformationReportingAction         uint8 = 195
	CSGID                                 uint8 = 196
	CSGMembershipIndication               uint8 = 197
	AggregateMaximumBitRate               uint8 = 198
	UENetworkCapability                   uint8 = 199
	UEAMBR                                uint8 = 200
	APNAMBRWithNSAPI                      uint8 = 201
	GGSNBackOffTime                       uint8 = 202
	SignallingPriorityIndication          uint8 = 203
	SignallingPriorityIndicationWithNSAPI uint8 = 204
	HigherBitratesThan16MbpsFlag          uint8 = 205
	AdditionalMMContextForSRVCC           uint8 = 207
	AdditionalFlagsForSRVCC               uint8 = 208
	STNSR                                 uint8 = 209
	CMSISDN                               uint8 = 210
	ExtendedRANAPCause                    uint8 = 211
	ENodeBID                              uint8 = 212
	SelectionModeWithNSAPI                uint8 = 213
	ULITimestamp                          uint8 = 214
	LHNIDWithNSAPI                        uint8 = 215
	CNOperatorSelectionEntity             uint8 = 216
	UEUsageType                           uint8 = 217
	ExtendedCommonFlagsII                 uint8 = 218
	NodeIdentifier                        uint8 = 219
	CIoTOptimizationsSupportIndication    uint8 = 220
	SCEFPDNConnection                     uint8 = 221
	IOVUpdatesCounter                     uint8 = 222
	MappedUEUsageType                     uint8 = 223
	UPFunctionSelectionIndicationFlags    uint8 = 224
	SpecialIETypeForIETypeExtension       uint8 = 238
	ChargingGatewayAddress                uint8 = 251
	PrivateExtension                      uint8 = 255
)

// IE is a GTPv1 Information Element.
type IE struct {
	Type    uint8
	Length  uint16
	Payload []byte
}

// New creates new IE.
func New(t uint8, p []byte) *IE {
	i := &IE{Type: t, Payload: p}
	i.SetLength()

	return i
}

// Marshal returns the byte sequence generated from an IE instance.
func (i *IE) Marshal() ([]byte, error) {
	b := make([]byte, i.MarshalLen())
	if err := i.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (i *IE) MarshalTo(b []byte) error {
	if len(b) < i.MarshalLen() {
		return ErrTooShortToMarshal
	}

	var offset = 1
	b[0] = i.Type

	if i.Type == ExtensionHeaderTypeList {
		b[1] = uint8(i.Length)
		offset++
	} else if !i.IsTV() {
		binary.BigEndian.PutUint16(b[1:3], i.Length)
		offset += 2
	}
	copy(b[offset:i.MarshalLen()], i.Payload)
	return nil
}

// Parse decodes given byte sequence as a GTPv1 Information Element.
func Parse(b []byte) (*IE, error) {
	i := &IE{}
	if err := i.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return i, nil
}

// UnmarshalBinary sets the values retrieved from byte sequence in GTPv1 IE.
func (i *IE) UnmarshalBinary(b []byte) error {
	if len(b) < 2 {
		return ErrTooShortToParse
	}

	i.Type = b[0]
	if i.Type == ExtensionHeaderTypeList {
		return decodeExtensionHeaderTypeList(i, b)
	}

	if i.IsTV() {
		return decodeTVFromBytes(i, b)
	}

	return decodeTLVFromBytes(i, b)
}

func decodeTVFromBytes(i *IE, b []byte) error {
	l := len(b)
	if l < 2 {
		return ErrTooShortToParse
	}
	if i.MarshalLen() > l {
		return ErrInvalidLength
	}

	i.Length = 0
	i.Payload = b[1:i.MarshalLen()]

	return nil
}

func decodeTLVFromBytes(i *IE, b []byte) error {
	l := len(b)
	if l < 3 {
		return ErrTooShortToParse
	}

	i.Length = binary.BigEndian.Uint16(b[1:3])
	if int(i.Length)+3 > l {
		return ErrInvalidLength
	}

	i.Payload = b[3 : 3+int(i.Length)]
	return nil
}

func decodeExtensionHeaderTypeList(i *IE, b []byte) error {
	l := len(b)
	if l < 2 {
		return ErrTooShortToParse
	}
	if i.MarshalLen() > l {
		return ErrInvalidLength
	}

	i.Length = uint16(b[1])
	n := 2 + int(i.Length)
	if n > l {
		return ErrInvalidLength
	}

	i.Payload = b[2:n]

	return nil
}

var tvLengthMap = map[int]int{
	0:   0,  // Reserved
	1:   1,  // Cause
	2:   8,  // IMSI
	3:   6,  // RouteingAreaIdentity
	4:   4,  // TLLI
	5:   4,  // P-TMSI
	8:   1,  // Reorder Required
	9:   28, // Authentication Triplet
	11:  1,  // MAP Cause
	12:  3,  // P-TMSI Signature
	13:  1,  // MS Validated
	14:  1,  // Recovery
	15:  1,  // Selection Mode
	16:  4,  // TEID Data 1
	17:  4,  // TEID C-Plane
	18:  4,  // TEID Data 2
	19:  1,  // Teardown Indication
	20:  1,  // NSAPI
	21:  1,  // RANAP Cause
	22:  9,  // RAB Context
	23:  1,  // Radio Priority SMS
	24:  1,  // Radio Priority
	25:  2,  // Packet Flow ID
	26:  2,  // Charging Characteristics
	27:  2,  // Trace Preference
	28:  2,  // Trace Type
	29:  1,  // MS Not Reachable Reason
	127: 4,  // Charging ID
}

// IsTV checks if a IE is TV format. If false, it indicates the IE has Length inside.
func (i *IE) IsTV() bool {
	return int(i.Type) < 0x80
}

// MarshalLen returns the serial length of IE.
func (i *IE) MarshalLen() int {
	if l, ok := tvLengthMap[int(i.Type)]; ok {
		return l + 1
	}

	if i.Type < 128 {
		return 1 + len(i.Payload)
	}

	if i.Type == ExtensionHeaderTypeList {
		return 2 + len(i.Payload)
	}

	return 3 + len(i.Payload)
}

// SetLength sets the length in Length field.
func (i *IE) SetLength() {
	if _, ok := tvLengthMap[int(i.Type)]; ok {
		i.Length = 0
		return
	}

	i.Length = uint16(len(i.Payload))
}

// Name returns the name of IE in string.
func (i *IE) Name() string {
	if n, ok := ieTypeNameMap[i.Type]; ok {
		return n
	}
	return "Undefined"
}

// String returns the GTPv1 IE values in human readable format.
func (i *IE) String() string {
	if i == nil {
		return "nil"
	}
	return fmt.Sprintf("{%s: {Type: %d, Length: %d, Payload: %#v}}",
		i.Name(),
		i.Type,
		i.Length,
		i.Payload,
	)
}

// ParseMultiIEs decodes multiple (unspecified number of) IEs to []*IE at a time.
func ParseMultiIEs(b []byte) ([]*IE, error) {
	var ies []*IE
	for {
		if len(b) == 0 {
			break
		}

		i, err := Parse(b)
		if err != nil {
			return nil, err
		}

		ies = append(ies, i)
		b = b[i.MarshalLen():]
		continue
	}
	return ies, nil
}

func newUint8ValIE(t, v uint8) *IE {
	return New(t, []byte{v})
}

// left for future use.
// func newUint16ValIE(t uint8, v uint16) *IE {
// 	i := New(t, make([]byte, 2))
// 	binary.BigEndian.PutUint16(i.Payload, v)
// 	return i
// }

func newUint32ValIE(t uint8, v uint32) *IE {
	i := New(t, make([]byte, 4))
	binary.BigEndian.PutUint32(i.Payload, v)
	return i
}

var ieTypeNameMap = map[uint8]string{
	128: "EndUserAddress",
	129: "MMContext",
	130: "PDPContext",
	131: "AccessPointName",
	132: "ProtocolConfigurationOptions",
	133: "GSNAddress",
	134: "MSISDN",
	135: "QoSProfile",
	136: "AuthenticationQuintuplet",
	137: "TrafficFlowTemplate",
	138: "TargetIdentification",
	139: "UTRANTransparentContainer",
	140: "RABSetupInformation",
	141: "ExtensionHeaderTypeList",
	142: "TriggerID",
	143: "OMCIdentity",
	144: "RANTransparentContainer",
	145: "PDPContextPrioritization",
	146: "AdditionalRABSetupInformation",
	147: "SGSNNumber",
	148: "CommonFlags",
	149: "APNRestriction",
	150: "RadioPriorityLCS",
	151: "RATType",
	152: "UserLocationInformation",
	153: "MSTimeZone",
	154: "IMEISV",
	155: "CAMELChargingInformationContainer",
	156: "MBMSUEContext",
	157: "TemporaryMobileGroupIdentity",
	158: "RIMRoutingAddress",
	159: "MBMSProtocolConfigurationOptions",
	160: "MBMSServiceArea",
	161: "SourceRNCPDCPContextInfo",
	162: "AdditionalTraceInfo",
	163: "HopCounter",
	164: "SelectedPLMNID",
	165: "MBMSSessionIdentifier",
	166: "MBMS2G3GIndicator",
	167: "EnhancedNSAPI",
	168: "MBMSSessionDuration",
	169: "AdditionalMBMSTraceInfo",
	170: "MBMSSessionRepetitionNumber",
	171: "MBMSTimeToDataTransfer",
	173: "BSSContainer",
	174: "CellIdentification",
	175: "PDUNumbers",
	176: "BSSGPCause",
	177: "RequiredMBMSBearerCapabilities",
	178: "RIMRoutingAddressDiscriminator",
	179: "ListOfSetupPFCs",
	180: "PSHandoverXIDParameters",
	181: "MSInfoChangeReportingAction",
	182: "DirectTunnelFlags",
	183: "CorrelationID",
	184: "BearerControlMode",
	185: "MBMSFlowIdentifier",
	186: "MBMSIPMulticastDistribution",
	187: "MBMSDistributionAcknowledgement",
	188: "ReliableInterRATHandoverInfo",
	189: "RFSPIndex",
	190: "FullyQualifiedDomainName",
	191: "EvolvedAllocationRetentionPriorityI",
	192: "EvolvedAllocationRetentionPriorityII",
	193: "ExtendedCommonFlags",
	194: "UserCSGInformation",
	195: "CSGInformationReportingAction",
	196: "CSGID",
	197: "CSGMembershipIndication",
	198: "AggregateMaximumBitRate",
	199: "UENetworkCapability",
	200: "UEAMBR",
	201: "APNAMBRWithNSAPI",
	202: "GGSNBackOffTime",
	203: "SignallingPriorityIndication",
	204: "SignallingPriorityIndicationWithNSAPI",
	205: "HigherBitratesThan16MbpsFlag",
	207: "AdditionalMMContextForSRVCC",
	208: "AdditionalFlagsForSRVCC",
	209: "STNSR",
	210: "CMSISDN",
	211: "ExtendedRANAPCause",
	212: "ENodeBID",
	213: "SelectionModeWithNSAPI",
	214: "ULITimestamp",
	215: "LHNIDWithNSAPI",
	216: "CNOperatorSelectionEntity",
	217: "UEUsageType",
	218: "ExtendedCommonFlagsII",
	219: "NodeIdentifier",
	220: "CIoTOptimizationsSupportIndication",
	221: "SCEFPDNConnection",
	222: "IOVUpdatesCounter",
	223: "MappedUEUsageType",
	224: "UPFunctionSelectionIndicationFlags",
	238: "SpecialIETypeForIETypeExtension",
	251: "ChargingGatewayAddress",
	255: "PrivateExtension",
}
