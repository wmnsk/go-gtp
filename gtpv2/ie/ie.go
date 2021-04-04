// Copyright 2019-2021 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

/*
Package ie provides encoding/decoding feature of GTPv2 Information Elements.
*/
package ie

import (
	"encoding/binary"
	"fmt"
	"io"
)

//

// IE definitions.
const (
	IMSI                                                 uint8 = 1
	Cause                                                uint8 = 2
	Recovery                                             uint8 = 3
	STNSR                                                uint8 = 51
	AccessPointName                                      uint8 = 71
	AggregateMaximumBitRate                              uint8 = 72
	EPSBearerID                                          uint8 = 73
	IPAddress                                            uint8 = 74
	MobileEquipmentIdentity                              uint8 = 75
	MSISDN                                               uint8 = 76
	Indication                                           uint8 = 77
	ProtocolConfigurationOptions                         uint8 = 78
	PDNAddressAllocation                                 uint8 = 79
	BearerQoS                                            uint8 = 80
	FlowQoS                                              uint8 = 81
	RATType                                              uint8 = 82
	ServingNetwork                                       uint8 = 83
	BearerTFT                                            uint8 = 84
	TrafficAggregateDescription                          uint8 = 85
	UserLocationInformation                              uint8 = 86
	FullyQualifiedTEID                                   uint8 = 87
	TMSI                                                 uint8 = 88
	GlobalCNID                                           uint8 = 89
	S103PDNDataForwardingInfo                            uint8 = 90
	S1UDataForwarding                                    uint8 = 91
	DelayValue                                           uint8 = 92
	BearerContext                                        uint8 = 93
	ChargingID                                           uint8 = 94
	ChargingCharacteristics                              uint8 = 95
	TraceInformation                                     uint8 = 96
	BearerFlags                                          uint8 = 97
	PDNType                                              uint8 = 99
	ProcedureTransactionID                               uint8 = 100
	MMContextGSMKeyAndTriplets                           uint8 = 103
	MMContextUMTSKeyUsedCipherAndQuintuplets             uint8 = 104
	MMContextGSMKeyUsedCipherAndQuintuplets              uint8 = 105
	MMContextUMTSKeyAndQuintuplets                       uint8 = 106
	MMContextEPSSecurityContextQuadrupletsAndQuintuplets uint8 = 107
	MMContextUMTSKeyQuadrupletsAndQuintuplets            uint8 = 108
	PDNConnection                                        uint8 = 109
	PDUNumbers                                           uint8 = 110
	PacketTMSI                                           uint8 = 111
	PTMSISignature                                       uint8 = 112
	HopCounter                                           uint8 = 113
	UETimeZone                                           uint8 = 114
	TraceReference                                       uint8 = 115
	CompleteRequestMessage                               uint8 = 116
	GUTI                                                 uint8 = 117
	FContainer                                           uint8 = 118
	FCause                                               uint8 = 119
	PLMNID                                               uint8 = 121
	TargetIdentification                                 uint8 = 122
	PacketFlowID                                         uint8 = 123
	RABContext                                           uint8 = 124
	SourceRNCPDCPContextInfo                             uint8 = 125
	PortNumber                                           uint8 = 126
	APNRestriction                                       uint8 = 127
	SelectionMode                                        uint8 = 128
	SourceIdentification                                 uint8 = 129
	Reserved                                             uint8 = 130
	ChangeReportingAction                                uint8 = 131
	FullyQualifiedCSID                                   uint8 = 132
	ChannelNeeded                                        uint8 = 133
	EMLPPPriority                                        uint8 = 134
	NodeType                                             uint8 = 135
	FullyQualifiedDomainName                             uint8 = 136
	TI                                                   uint8 = 137
	MBMSSessionDuration                                  uint8 = 138
	MBMSServiceArea                                      uint8 = 139
	MBMSSessionIdentifier                                uint8 = 140
	MBMSFlowIdentifier                                   uint8 = 141
	MBMSIPMulticastDistribution                          uint8 = 142
	MBMSDistributionAcknowledge                          uint8 = 143
	RFSPIndex                                            uint8 = 144
	UserCSGInformation                                   uint8 = 145
	CSGInformationReportingAction                        uint8 = 146
	CSGID                                                uint8 = 147
	CSGMembershipIndication                              uint8 = 148
	ServiceIndicator                                     uint8 = 149
	DetachType                                           uint8 = 150
	LocalDistinguishedName                               uint8 = 151
	NodeFeatures                                         uint8 = 152
	MBMSTimeToDataTransfer                               uint8 = 153
	Throttling                                           uint8 = 154
	AllocationRetensionPriority                          uint8 = 155
	EPCTimer                                             uint8 = 156
	SignallingPriorityIndication                         uint8 = 157
	TMGI                                                 uint8 = 158
	AdditionalMMContextForSRVCC                          uint8 = 159
	AdditionalFlagsForSRVCC                              uint8 = 160
	MDTConfiguration                                     uint8 = 162
	AdditionalProtocolConfigurationOptions               uint8 = 163
	AbsoluteTimeofMBMSDataTransfer                       uint8 = 164
	HeNBInformationReporting                             uint8 = 165
	IPv4ConfigurationParameters                          uint8 = 166
	ChangeToReportFlags                                  uint8 = 167
	ActionIndication                                     uint8 = 168
	TWANIdentifier                                       uint8 = 169
	ULITimestamp                                         uint8 = 170
	MBMSFlags                                            uint8 = 171
	RANNASCause                                          uint8 = 172
	CNOperatorSelectionEntity                            uint8 = 173
	TrustedWLANModeIndication                            uint8 = 174
	NodeNumber                                           uint8 = 175
	NodeIdentifier                                       uint8 = 176
	PresenceReportingAreaAction                          uint8 = 177
	PresenceReportingAreaInformation                     uint8 = 178
	TWANIdentifierTimestamp                              uint8 = 179
	OverloadControlInformation                           uint8 = 180
	LoadControlInformation                               uint8 = 181
	Metric                                               uint8 = 182
	SequenceNumber                                       uint8 = 183
	APNAndRelativeCapacity                               uint8 = 184
	WLANOffloadabilityIndication                         uint8 = 185
	PagingAndServiceInformation                          uint8 = 186
	IntegerNumber                                        uint8 = 187
	MillisecondTimeStamp                                 uint8 = 188
	MonitoringEventInformation                           uint8 = 189
	ECGIList                                             uint8 = 190
	RemoteUEContext                                      uint8 = 191
	RemoteUserID                                         uint8 = 192
	RemoteUEIPinformation                                uint8 = 193
	CIoTOptimizationsSupportIndication                   uint8 = 194
	SCEFPDNConnection                                    uint8 = 195
	HeaderCompressionConfiguration                       uint8 = 196
	ExtendedProtocolConfigurationOptions                 uint8 = 197
	ServingPLMNRateControl                               uint8 = 198
	Counter                                              uint8 = 199
	MappedUEUsageType                                    uint8 = 200
	SecondaryRATUsageDataReport                          uint8 = 201
	UPFunctionSelectionIndicationFlags                   uint8 = 202
	MaximumPacketLossRate                                uint8 = 203
	APNRateControlStatus                                 uint8 = 204
	ExtendedTraceInformation                             uint8 = 205
	MonitoringEventExtensionInformation                  uint8 = 206
	AdditionalRRMPolicyIndex                             uint8 = 207
	V2XContext                                           uint8 = 208
	PC5QoSParameters                                     uint8 = 209
	ServicesAuthorized                                   uint8 = 210
	BitRate                                              uint8 = 211
	PC5QoSFlow                                           uint8 = 212
	SGiPtPTunnelAddress                                  uint8 = 213
	SpecialIETypeForIETypeExtension                      uint8 = 254
	PrivateExtension                                     uint8 = 255
)

// IE is a GTPv2 Information Element.
type IE struct {
	Type     uint8
	Length   uint16
	instance uint8
	Payload  []byte
	ChildIEs []*IE
}

// New creates new IE.
func New(itype, ins uint8, data []byte) *IE {
	ie := &IE{
		Type:     itype,
		instance: ins & 0x0f,
		Payload:  data,
	}
	ie.SetLength()

	return ie
}

// SetInstance sets the instance.
func (i *IE) SetInstance(ins uint8) {
	i.instance = ins & 0x0f
}

// WithInstance sets the instance and returns IE.
func (i *IE) WithInstance(ins uint8) *IE {
	i.instance = ins & 0x0f
	return i
}

// Instance returns instance value in uint8
func (i *IE) Instance() uint8 {
	return i.instance & 0x0f
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
	l := len(b)
	if l < 4 {
		return io.ErrUnexpectedEOF
	}

	b[0] = i.Type
	binary.BigEndian.PutUint16(b[1:3], i.Length)
	b[3] = i.instance
	if i.IsGrouped() {
		offset := 4
		for _, ie := range i.ChildIEs {
			if l < offset+ie.MarshalLen() {
				break
			}

			if err := ie.MarshalTo(b[offset : offset+ie.MarshalLen()]); err != nil {
				return err
			}
			offset += ie.MarshalLen()
		}
		return nil
	}

	if l < i.MarshalLen() {
		return io.ErrUnexpectedEOF
	}

	copy(b[4:i.MarshalLen()], i.Payload)
	return nil
}

// Parse decodes given byte sequence as a GTPv2 Information Element.
func Parse(b []byte) (*IE, error) {
	ie := &IE{}
	if err := ie.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return ie, nil
}

// UnmarshalBinary sets the values retrieved from byte sequence in GTPv2 IE.
func (i *IE) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 5 {
		return io.ErrUnexpectedEOF
	}

	i.Type = b[0]
	i.Length = binary.BigEndian.Uint16(b[1:3])
	if int(i.Length) > l-4 {
		return ErrInvalidLength
	}

	i.instance = b[3]
	i.Payload = b[4 : 4+int(i.Length)]

	if i.IsGrouped() {
		var err error
		i.ChildIEs, err = ParseMultiIEs(i.Payload)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalLen returns field length in integer.
func (i *IE) MarshalLen() int {
	if i.IsGrouped() {
		l := 4
		for _, ie := range i.ChildIEs {
			l += ie.MarshalLen()
		}
		return l
	}
	return 4 + len(i.Payload)
}

// SetLength sets the length in Length field.
func (i *IE) SetLength() {
	if i.IsGrouped() {
		l := 0
		for _, ie := range i.ChildIEs {
			l += ie.MarshalLen()
		}
		i.Length = uint16(l)
	}
	i.Length = uint16(len(i.Payload))
}

// String returns the GTPv2 IE values in human readable format.
func (i *IE) String() string {
	return fmt.Sprintf("{Type: %d, Length: %d, Instance: %#x, Payload: %#v}",
		i.Type,
		i.Length,
		i.Instance(),
		i.Payload,
	)
}

var grouped = []uint8{
	BearerContext,
	PDNConnection,
	OverloadControlInformation,
	LoadControlInformation,
	RemoteUEContext,
	SCEFPDNConnection,
	V2XContext,
	PC5QoSParameters,
}

// IsGrouped reports whether an IE is grouped type or not.
func (i *IE) IsGrouped() bool {
	for _, itype := range grouped {
		if i.Type == itype {
			return true
		}
	}
	return false
}

// Add adds variable number of IEs to a IE if the IE is grouped type and update length.
// Otherwise, this does nothing(no errors).
func (i *IE) Add(ies ...*IE) {
	if !i.IsGrouped() {
		return
	}

	i.Payload = nil
	i.ChildIEs = append(i.ChildIEs, ies...)
	for _, ie := range i.ChildIEs {
		serialized, err := ie.Marshal()
		if err != nil {
			continue
		}
		i.Payload = append(i.Payload, serialized...)
	}
	i.SetLength()
}

// Remove removes an IE looked up by type and instance.
func (i *IE) Remove(typ, instance uint8) {
	if !i.IsGrouped() {
		return
	}

	i.Payload = nil
	newChildren := make([]*IE, len(i.ChildIEs))
	idx := 0
	for _, ie := range i.ChildIEs {
		if ie.Type == typ && ie.Instance() == instance {
			newChildren = newChildren[:len(newChildren)-1]
			continue
		}
		newChildren[idx] = ie
		idx++

		serialized, err := ie.Marshal()
		if err != nil {
			continue
		}
		i.Payload = append(i.Payload, serialized...)
	}
	i.ChildIEs = newChildren
	i.SetLength()
}

// FindByType returns IE looked up by type and instance.
//
// The program may be slower when calling this method multiple times
// because this ranges over a ChildIEs each time it is called.
func (i *IE) FindByType(typ, instance uint8) (*IE, error) {
	if !i.IsGrouped() {
		return nil, ErrInvalidType
	}

	for _, ie := range i.ChildIEs {
		if ie.Type == typ && ie.Instance() == instance {
			return ie, nil
		}
	}
	return nil, ErrIENotFound
}

// ParseMultiIEs decodes multiple IEs at a time.
// This is easy and useful but slower than decoding one by one.
// When you don't know the number of IEs, this is the only way to decode them.
// See benchmarks in diameter_test.go for the detail.
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
	}
	return ies, nil
}

func newUint8ValIE(t, v uint8) *IE {
	return New(t, 0x00, []byte{v})
}

func newUint16ValIE(t uint8, v uint16) *IE {
	i := New(t, 0x00, make([]byte, 2))
	binary.BigEndian.PutUint16(i.Payload, v)
	return i
}

func newUint32ValIE(t uint8, v uint32) *IE {
	i := New(t, 0x00, make([]byte, 4))
	binary.BigEndian.PutUint32(i.Payload, v)
	return i
}

// unused for now.
// func newUint64ValIE(t uint8, v uint64) *IE {
// 	i := New(t, 0x00, make([]byte, 8))
// 	binary.BigEndian.PutUint64(i.Payload, v)
// 	return i
// }

func newStringIE(t uint8, v string) *IE {
	return New(t, 0x00, []byte(v))
}

func newGroupedIE(itype uint8, ies ...*IE) *IE {
	i := New(itype, 0x00, make([]byte, 0))
	i.ChildIEs = ies
	for _, ie := range i.ChildIEs {
		serialized, err := ie.Marshal()
		if err != nil {
			return nil
		}
		i.Payload = append(i.Payload, serialized...)
	}
	i.SetLength()

	return i
}
