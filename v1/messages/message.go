// Copyright 2019-2020 go-gtp authors. All rights reserved.
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
	MarshalTo([]byte) error
	UnmarshalBinary(b []byte) error
	MarshalLen() int
	Version() int
	MessageType() uint8
	MessageTypeName() string
	TEID() uint32
	SetTEID(uint32)
	Sequence() uint16
	SetSequenceNumber(uint16)

	// deprecated
	SerializeTo([]byte) error
	DecodeFromBytes(b []byte) error
}

// Marshal returns the byte sequence generated from a Message instance.
// Better to use MarshalXxx instead if you know the name of message to be serialized.
func Marshal(g Message) ([]byte, error) {
	b := make([]byte, g.MarshalLen())
	if err := g.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// Parse decodes the given bytes as Message.
func Parse(b []byte) (Message, error) {
	var m Message

	switch b[1] {
	case MsgTypeEchoRequest:
		m = &EchoRequest{}
	case MsgTypeEchoResponse:
		m = &EchoResponse{}
	case MsgTypeCreatePDPContextRequest:
		m = &CreatePDPContextRequest{}
	case MsgTypeCreatePDPContextResponse:
		m = &CreatePDPContextResponse{}
	case MsgTypeUpdatePDPContextRequest:
		m = &UpdatePDPContextRequest{}
	case MsgTypeUpdatePDPContextResponse:
		m = &UpdatePDPContextResponse{}
	case MsgTypeDeletePDPContextRequest:
		m = &DeletePDPContextRequest{}
	case MsgTypeVersionNotSupported:
		m = &VersionNotSupported{}
	case MsgTypeDeletePDPContextResponse:
		m = &DeletePDPContextResponse{}
	/* XXX - Implement!
	case MsgTypeNodeAliveRequest:
		m = &NodeAliveReq{}
	case MsgTypeNodeAliveResponse:
		m = &NodeAliveRes{}
	case MsgTypeRedirectionRequest:
		m = &RedirectionReq{}
	case MsgTypeRedirectionResponse:
		m = &RedirectionRes{}
	case MsgTypeCreateAaPDPContextRequest:
		m = &CreateAaPDPContextReq{}
	case MsgTypeCreateAaPDPContextResponse:
		m = &CreateAaPDPContextRes{}
	case MsgTypeDeleteAaPDPContextRequest:
		m = &DeleteAaPDPContextReq{}
	case MsgTypeDeleteAaPDPContextResponse:
		m = &DeleteAaPDPContextRes{}
	*/
	case MsgTypeErrorIndication:
		m = &ErrorIndication{}
	/* XXX - Implement!
	case MsgTypePduNotificationRequest:
		m = &PduNotificationReq{}
	case MsgTypePduNotificationResponse:
		m = &PduNotificationRes{}
	case MsgTypePduNotificationRejectRequest:
		m = &PduNotificationRejectReq{}
	case MsgTypePduNotificationRejectResponse:
		m = &PduNotificationRejectRes{}
	case MsgTypeSendRoutingInfoRequest:
		m = &SendRoutingInfoReq{}
	case MsgTypeSendRoutingInfoResponse:
		m = &SendRoutingInfoRes{}
	case MsgTypeFailureReportRequest:
		m = &FailureReportReq{}
	case MsgTypeFailureReportResponse:
		m = &FailureReportRes{}
	case MsgTypeNoteMsPresentRequest:
		m = &NoteMsPresentReq{}
	case MsgTypeNoteMsPresentResponse:
		m = &NoteMsPresentRes{}
	case MsgTypeIdentificationRequest:
		m = &IdentificationReq{}
	case MsgTypeIdentificationResponse:
		m = &IdentificationRes{}
	case MsgTypeSgsnContextRequest:
		m = &SgsnContextReq{}
	case MsgTypeSgsnContextResponse:
		m = &SgsnContextRes{}
	case MsgTypeSgsnContextAcknowledge:
		m = &SgsnContextAck{}
	case MsgTypeDataRecordTransferRequest:
		m = &DataRecordTransferReq{}
	case MsgTypeDataRecordTransferResponse:
		m = &DataRecordTransferRes{}
	*/
	case MsgTypeTPDU:
		m = &TPDU{}
	default:
		m = &Generic{}
	}

	if err := m.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return m, nil
}
