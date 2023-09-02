// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

/*
Package ie provides encoding/decoding feature of GTPv0 Information Elements.
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
	QualityOfServiceProfile      uint8 = 6
	ReorderingRequired           uint8 = 8
	AuthenticationTriplet        uint8 = 9
	MAPCause                     uint8 = 11
	PTMSISignature               uint8 = 12
	MSValidated                  uint8 = 13
	Recovery                     uint8 = 14
	SelectionMode                uint8 = 15
	FlowLabelDataI               uint8 = 16
	FlowLabelSignalling          uint8 = 17
	FlowLabelDataII              uint8 = 18
	MSNotReachableReason         uint8 = 19
	ChargingID                   uint8 = 127
)

// TLV IE definitions.
const (
	EndUserAddress               uint8 = 128
	MMContext                    uint8 = 129
	PDPContext                   uint8 = 130
	AccessPointName              uint8 = 131
	ProtocolConfigurationOptions uint8 = 132
	GSNAddress                   uint8 = 133
	MSISDN                       uint8 = 134
	ChargingGatewayAddress       uint8 = 251
	PrivateExtension             uint8 = 255
)

// IE is a GTPv0 Information Element.
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
	if !i.IsTV() {
		binary.BigEndian.PutUint16(b[1:3], i.Length)
		offset += 2
	}
	copy(b[offset:i.MarshalLen()], i.Payload)
	return nil
}

// Parse Parses given byte sequence as a GTPv0 Information Element.
func Parse(b []byte) (*IE, error) {
	i := &IE{}
	if err := i.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return i, nil
}

// UnmarshalBinary sets the values retrieved from byte sequence in GTPv0 IE.
func (i *IE) UnmarshalBinary(b []byte) error {
	if len(b) < 2 {
		return ErrTooShortToParse
	}

	i.Type = b[0]
	if i.IsTV() {
		return parseTVFromBytes(i, b)
	}
	return parseTLVFromBytes(i, b)
}

func parseTVFromBytes(i *IE, b []byte) error {
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

func parseTLVFromBytes(i *IE, b []byte) error {
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

var tvLengthMap = map[uint8]int{
	0:   0,  // Reserved
	1:   1,  // Cause
	2:   8,  // IMSI
	3:   6,  // RAI
	4:   4,  // TLLI
	5:   4,  // P-TMSI
	6:   3,  // QoS
	8:   1,  // Reordering Required
	9:   28, // Authentication Triplet
	11:  1,  // MAP Cause
	12:  3,  // P-TMSI Signature
	13:  1,  // MS Validated
	14:  1,  // Recovery
	15:  1,  // Selection Mode
	16:  2,  // Flow Label Data I
	17:  2,  // Flow Label Signalling
	18:  3,  // Flow Label Data II
	19:  1,  // MS Not Reachable Reason
	127: 4,  // Charging ID
}

// IsTV checks if a IE is TV format. If false, it indicates the IE has Length inside.
func (i *IE) IsTV() bool {
	return int(i.Type) < 0x80
}

// MarshalLen returns the serial length of IE.
func (i *IE) MarshalLen() int {
	if l, ok := tvLengthMap[i.Type]; ok {
		return l + 1
	}
	if i.Type < 128 {
		return 1 + len(i.Payload)
	}
	return 3 + len(i.Payload)
}

// SetLength sets the length in Length field.
func (i *IE) SetLength() {
	if _, ok := tvLengthMap[i.Type]; ok {
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

// String returns the GTPv0 IE values in human readable format.
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

// ParseMultiIEs Parses multiple (unspecified number of) IEs to []*IE at a time.
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

func newUint16ValIE(t uint8, v uint16) *IE {
	i := New(t, make([]byte, 2))
	binary.BigEndian.PutUint16(i.Payload, v)
	return i
}

func newUint32ValIE(t uint8, v uint32) *IE {
	i := New(t, make([]byte, 4))
	binary.BigEndian.PutUint32(i.Payload, v)
	return i
}

// left for future use.
// func newStringIE(t uint8, str string) *IE {
// 	return New(t, []byte(str))
// }

var ieTypeNameMap = map[uint8]string{
	1:   "Cause",
	2:   "IMSI",
	3:   "RouteingAreaIdentity",
	4:   "TemporaryLogicalLinkIdentity",
	5:   "PacketTMSI",
	6:   "QualityOfServiceProfile",
	8:   "ReorderingRequired",
	9:   "AuthenticationTriplet",
	11:  "MAPCause",
	12:  "PTMSISignature",
	13:  "MSValidated",
	14:  "Recovery",
	15:  "SelectionMode",
	16:  "FlowLabelDataI",
	17:  "FlowLabelSignalling",
	18:  "FlowLabelDataII",
	19:  "MSNotReachableReason",
	127: "ChargingID",
	128: "EndUserAddress",
	129: "MMContext",
	130: "PDPContext",
	131: "AccessPointName",
	132: "ProtocolConfigurationOptions",
	133: "GSNAddress",
	134: "MSISDN",
	251: "ChargingGatewayAddress",
	255: "PrivateExtension",
}
