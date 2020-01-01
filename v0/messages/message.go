// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

/*
Package messages provides encoding/decoding feature of GTPv0 protocol.
*/
package messages

import (
	"github.com/pkg/errors"
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

// Message is an interface that defines Message messages.
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

// Parse Parses the given bytes as Message.
func Parse(b []byte) (Message, error) {
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
		return nil, errors.Wrap(err, "failed to Parse Message:")
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
