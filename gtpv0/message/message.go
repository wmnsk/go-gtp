// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

/*
Package message provides encoding/decoding feature of GTPv0 protocol.
*/
package message

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/wmnsk/go-gtp/gtpv0/ie"
)

// MessageType definitions.
const (
	_ uint8 = iota
	MsgTypeEchoRequest
	MsgTypeEchoResponse
	MsgTypeVersionNotSupported
	MsgTypeNodeAliveRequest
	MsgTypeNodeAliveResponse
	MsgTypeRedirectionRequest
	MsgTypeRedirectionResponse
	_
	_
	_
	_
	_
	_
	_
	_
	MsgTypeCreatePDPContextRequest // 16
	MsgTypeCreatePDPContextResponse
	MsgTypeUpdatePDPContextRequest
	MsgTypeUpdatePDPContextResponse
	MsgTypeDeletePDPContextRequest
	MsgTypeDeletePDPContextResponse
	MsgTypeCreateAAPDPContextRequest
	MsgTypeCreateAAPDPContextResponse
	MsgTypeDeleteAAPDPContextRequest
	MsgTypeDeleteAAPDPContextResponse
	MsgTypeErrorIndication
	MsgTypePDUNotificationRequest
	MsgTypePDUNotificationResponse
	MsgTypePDUNotificationRejectRequest
	MsgTypePDUNotificationRejectResponse
	_
	MsgTypeSendRouteingInformationforGPRSRequest // 32
	MsgTypeSendRouteingInformationforGPRSResponse
	MsgTypeFailureReportRequest
	MsgTypeFailureReportResponse
	MsgTypeNoteMSGPRSPresentRequest
	MsgTypeNoteMSGPRSPresentResponse
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	MsgTypeIdentificationRequest // 48
	MsgTypeIdentificationResponse
	MsgTypeSGSNContextRequest
	MsgTypeSGSNContextResponse
	MsgTypeSGSNContextAcknowledge
	MsgTypeDataRecordTransferRequest  = 240
	MsgTypeDataRecordTransferResponse = 241
	MsgTypeTPDU                       = 255
)

// Message is an interface that defines Message message.
type Message interface {
	MarshalTo([]byte) error
	UnmarshalBinary(b []byte) error
	MarshalLen() int
	String() string
	Version() int
	MessageType() uint8
	MessageTypeName() string
	TID() string

	// deprecated
	SerializeTo([]byte) error
	DecodeFromBytes(b []byte) error
}

// Marshal returns the byte sequence generated from a Message instance.
// Better to use MarshalXxx instead if you know the name of message to be Serialized.
func Marshal(g Message) ([]byte, error) {
	b := make([]byte, g.MarshalLen())
	if err := g.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// Parse parses the given bytes as a Message.
func Parse(b []byte) (Message, error) {
	if len(b) < 2 {
		return nil, ErrTooShortToParse
	}

	var g Message

	switch b[1] {
	case MsgTypeEchoRequest:
		g = &EchoRequest{}
	case MsgTypeEchoResponse:
		g = &EchoResponse{}
	/* XXX - Implement!
	case MsgTypeVersionNotSupported:
		g = &VerNotSupported{}
	case MsgTypeNodeAliveRequest:
		g = &NodeAliveReq{}
	case MsgTypeNodeAliveResponse:
		g = &NodeAliveRes{}
	case MsgTypeRedirectionRequest:
		g = &RedirectionReq{}
	case MsgTypeRedirectionResponse:
		g = &RedirectionRes{}
	*/
	case MsgTypeCreatePDPContextRequest:
		g = &CreatePDPContextRequest{}
	case MsgTypeCreatePDPContextResponse:
		g = &CreatePDPContextResponse{}
	case MsgTypeUpdatePDPContextRequest:
		g = &UpdatePDPContextRequest{}
	case MsgTypeUpdatePDPContextResponse:
		g = &UpdatePDPContextResponse{}
	case MsgTypeDeletePDPContextRequest:
		g = &DeletePDPContextRequest{}
	case MsgTypeDeletePDPContextResponse:
		g = &DeletePDPContextResponse{}
	/* XXX - Implement!
	case MsgTypeCreateAAPDPContextRequest:
		g = &CreateAAPDPContextReq{}
	case MsgTypeCreateAAPDPContextResponse:
		g = &CreateAAPDPContextRes{}
	case MsgTypeDeleteAAPDPContextRequest:
		g = &DeleteAAPDPContextReq{}
	case MsgTypeDeleteAAPDPContextResponse:
		g = &DeleteAAPDPContextRes{}
	case MsgTypeErrorIndication:
		g = &ErrorInd{}
	case MsgTypePDUNotificationRequest:
		g = &PDUNotificationReq{}
	case MsgTypePDUNotificationResponse:
		g = &PDUNotificationRes{}
	case MsgTypePDUNotificationRejectRequest:
		g = &PDUNotificationRejectReq{}
	case MsgTypePDUNotificationRejectResponse:
		g = &PDUNotificationRejectRes{}
	case MsgTypeSendRouteingInformationforGPRSRequest:
		g = &SendRouteingInformationforGPRSReq{}
	case MsgTypeSendRouteingInformationforGPRSResponse:
		g = &SendRouteingInformationforGPRSRes{}
	case MsgTypeFailureReportRequest:
		g = &FailureReportReq{}
	case MsgTypeFailureReportResponse:
		g = &FailureReportRes{}
	case MsgTypeNoteMSGPRSPresentRequest:
		g = &NoteMSGPRSPresentReq{}
	case MsgTypeNoteMSGPRSPresentResponse:
		g = &NoteMSGPRSPresentRes{}
	case MsgTypeIdentificationRequest:
		g = &IdentificationReq{}
	case MsgTypeIdentificationResponse:
		g = &IdentificationRes{}
	case MsgTypeSGSNContextRequest:
		g = &SGSNContextReq{}
	case MsgTypeSGSNContextResponse:
		g = &SGSNContextRes{}
	case MsgTypeSGSNContextAcknowledge:
		g = &SGSNContextAck{}
	case MsgTypeDataRecordTransferRequest:
		g = &DataRecordTransferReq{}
	case MsgTypeDataRecordTransferResponse:
		g = &DataRecordTransferRes{}
	*/
	case MsgTypeTPDU:
		g = &TPDU{}
	default:
		g = &Generic{}
	}

	if err := g.UnmarshalBinary(b); err != nil {
		return nil, fmt.Errorf("failed to Parse Message: %w", err)
	}
	return g, nil
}

// Decapsulate decapsulates given bytes and returns Payload in []byte.
func Decapsulate(b []byte) ([]byte, error) {
	header, err := ParseHeader(b)
	if err != nil {
		return nil, err
	}

	if header.Type != MsgTypeTPDU {
		return nil, nil
	}
	return header.Payload, nil
}

// Prettify returns a Message in prettified representation in string.
//
// Note that this relies much on reflect package, and thus the frequent use of
// this function may have a serious impact on the performance of your software.
func Prettify(m Message) string {
	name := m.MessageTypeName()
	header := strings.TrimSuffix(fmt.Sprint(m), "}")

	v := reflect.Indirect(reflect.ValueOf(m))
	n := v.NumField() - 1
	fields := make([]*field, n)
	for i := 1; i < n+1; i++ { // Skip *Header
		fields[i-1] = &field{name: v.Type().Field(i).Name, maybeIE: v.Field(i).Interface()}
	}

	return fmt.Sprintf("{%s: %s, IEs: [%v]}", name, header, strings.Join(prettifyFields(fields), ", "))
}

type field struct {
	name    string
	maybeIE interface{}
}

func prettifyFields(fields []*field) []string {
	ret := []string{}
	for _, field := range fields {
		if field.maybeIE == nil {
			ret = append(ret, prettifyIE(field.name, nil))
			continue
		}

		// TODO: do this recursively?
		v, ok := field.maybeIE.(*ie.IE)
		if !ok {
			// only for AdditionalIEs field
			if ies, ok := field.maybeIE.([]*ie.IE); ok {
				vals := make([]string, len(ies))
				for i, val := range ies {
					vals[i] = fmt.Sprint(val)
				}
				ret = append(ret, fmt.Sprintf("{%s: [%v]}", field.name, strings.Join(vals, ", ")))
			}
			continue
		}

		ret = append(ret, prettifyIE(field.name, v))
	}

	return ret
}

func prettifyIE(name string, i *ie.IE) string {
	return fmt.Sprintf("{%s: %v}", name, i)
}
