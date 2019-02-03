// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

/*
Package messages provides encoding/decoding feature of GTPv1 protocol.
*/
package messages

// Message Type definitions.
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
	MsgTypeSendRoutingInfoRequest
	MsgTypeSendRoutingInfoResponse
	MsgTypeFailureReportRequest
	MsgTypeFailureReportResponse
	MsgTypeNoteMSPresentRequest
	MsgTypeNoteMSPresentResponse
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
	MsgTypeDataRecordTransferRequest  uint8 = 240
	MsgTypeDataRecordTransferResponse uint8 = 241
	MsgTypeTPDU                       uint8 = 255
)

// Message is an interface that defines Message messages.
type Message interface {
	SerializeTo([]byte) error
	DecodeFromBytes(b []byte) error
	Len() int
	Version() int
	MessageType() uint8
	MessageTypeName() string
	TEID() uint32
	SetTEID(uint32)
	Sequence() uint16
	SetSequenceNumber(uint16)
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
	case MsgTypeCreatePDPContextRequest:
		g = &CreatePDPContextRequest{}
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
	case MsgTypeCreatePDPContextResponse:
		g = &CreatePDPContextRes{}
	case MsgTypeUpdatePDPContextRequest:
		g = &UpdatePDPContextReq{}
	case MsgTypeUpdatePDPContextResponse:
		g = &UpdatePDPContextRes{}
	case MsgTypeDeletePDPContextRequest:
		g = &DeletePDPContextReq{}
	case MsgTypeDeletePDPContextResponse:
		g = &DeletePDPContextRes{}
	case MsgTypeCreateAaPDPContextRequest:
		g = &CreateAaPDPContextReq{}
	case MsgTypeCreateAaPDPContextResponse:
		g = &CreateAaPDPContextRes{}
	case MsgTypeDeleteAaPDPContextRequest:
		g = &DeleteAaPDPContextReq{}
	case MsgTypeDeleteAaPDPContextResponse:
		g = &DeleteAaPDPContextRes{}
	*/
	case MsgTypeErrorIndication:
		g = &ErrorIndication{}
	/* XXX - Implement!
	case MsgTypePduNotificationRequest:
		g = &PduNotificationReq{}
	case MsgTypePduNotificationResponse:
		g = &PduNotificationRes{}
	case MsgTypePduNotificationRejectRequest:
		g = &PduNotificationRejectReq{}
	case MsgTypePduNotificationRejectResponse:
		g = &PduNotificationRejectRes{}
	case MsgTypeSendRoutingInfoRequest:
		g = &SendRoutingInfoReq{}
	case MsgTypeSendRoutingInfoResponse:
		g = &SendRoutingInfoRes{}
	case MsgTypeFailureReportRequest:
		g = &FailureReportReq{}
	case MsgTypeFailureReportResponse:
		g = &FailureReportRes{}
	case MsgTypeNoteMsPresentRequest:
		g = &NoteMsPresentReq{}
	case MsgTypeNoteMsPresentResponse:
		g = &NoteMsPresentRes{}
	case MsgTypeIdentificationRequest:
		g = &IdentificationReq{}
	case MsgTypeIdentificationResponse:
		g = &IdentificationRes{}
	case MsgTypeSgsnContextRequest:
		g = &SgsnContextReq{}
	case MsgTypeSgsnContextResponse:
		g = &SgsnContextRes{}
	case MsgTypeSgsnContextAcknowledge:
		g = &SgsnContextAck{}
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
		return nil, err
	}
	return g, nil
}
