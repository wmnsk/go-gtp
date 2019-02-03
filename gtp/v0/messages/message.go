// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

/*
Package messages provides encoding/decoding feature of GTPv0 protocol.
*/
package messages

import (
	"github.com/pkg/errors"
	"github.com/wmnsk/go-gtp/gtp/v0/ies"
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
	SerializeTo([]byte) error
	DecodeFromBytes(b []byte) error
	Len() int
	String() string
	Version() int
	MessageType() uint8
	MessageTypeName() string
	TID() string
}

// Serialize returns the byte sequence generated from a Message instance.
// Better to use SerializeXxx instead if you know the name of message to be serialized.
func Serialize(g Message) ([]byte, error) {
	b := make([]byte, g.Len())
	if err := g.SerializeTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// Decode decodes the given bytes as Message.
func Decode(b []byte) (Message, error) {
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

	if err := g.DecodeFromBytes(b); err != nil {
		return nil, errors.Wrap(err, "failed to decode Message:")
	}
	return g, nil
}

// Decapsulate decapsulates given bytes and returns Payload in []byte.
func Decapsulate(b []byte) ([]byte, error) {
	header, err := DecodeHeader(b)
	if err != nil {
		return nil, err
	}

	if header.Type != MsgTypeTPDU {
		return nil, nil
	}
	return header.Payload, nil
}

func sumMultiIELen(multipleIEs ...*ies.IE) int {
	l := 0
	for _, ie := range multipleIEs {
		if ie != nil {
			l += ie.Len()
		}
	}
	return l
}

func setMultiIELength(multipleIEs ...*ies.IE) {
	for _, ie := range multipleIEs {
		if ie != nil {
			ie.SetLength()
		}
	}
}
