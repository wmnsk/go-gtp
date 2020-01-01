// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

/*
Package messages provides encoding/decoding feature of GTPv2 protocol.
*/
package messages

import (
	"github.com/pkg/errors"
)

// Message Type definitions.
const (
	_ uint8 = iota
	MsgTypeEchoRequest
	MsgTypeEchoResponse
	MsgTypeVersionNotSupportedIndication
	MsgTypeDirectTransferRequest
	MsgTypeDirectTransferResponse
	MsgTypeNotificationRequest
	MsgTypeNotificationResponse
	_
	_
	_
	_
	_
	_
	_
	_
	_ // 8-16: Reserved for S101 interface
	MsgTypeRIMInformationTransfer
	_
	_
	_
	_
	_
	_
	_ // 18-24: Reserved for S121 interface
	MsgTypeSRVCCPsToCsRequest
	MsgTypeSRVCCPsToCsResponse
	MsgTypeSRVCCPsToCsCompleteNotification
	MsgTypeSRVCCPsToCsCompleteAcknowledge
	MsgTypeSRVCCPsToCsCancelNotification
	MsgTypeSRVCCPsToCsCancelAcknowledge
	MsgTypeSRVCCCsToPsRequest
	MsgTypeCreateSessionRequest
	MsgTypeCreateSessionResponse
	MsgTypeModifyBearerRequest
	MsgTypeModifyBearerResponse
	MsgTypeDeleteSessionRequest
	MsgTypeDeleteSessionResponse
	MsgTypeChangeNotificationRequest
	MsgTypeChangeNotificationResponse
	MsgTypeRemoteUEReportNotification
	MsgTypeRemoteUEReportAcknowledge
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
	_
	_
	_
	_ // 42-63: Reserved for S4/S11, S5/S8 interfaces
	MsgTypeModifyBearerCommand
	MsgTypeModifyBearerFailureIndication
	MsgTypeDeleteBearerCommand
	MsgTypeDeleteBearerFailureIndication
	MsgTypeBearerResourceCommand
	MsgTypeBearerResourceFailureIndication
	MsgTypeDownlinkDataNotificationFailureIndication
	MsgTypeTraceSessionActivation
	MsgTypeTraceSessionDeactivation
	MsgTypeStopPagingIndication
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
	_
	_
	_ // 74-94: Reserved for GTPv2 non-specific interfaces
	MsgTypeCreateBearerRequest
	MsgTypeCreateBearerResponse
	MsgTypeUpdateBearerRequest
	MsgTypeUpdateBearerResponse
	MsgTypeDeleteBearerRequest
	MsgTypeDeleteBearerResponse
	MsgTypeDeletePDNConnectionSetRequest
	MsgTypeDeletePDNConnectionSetResponse
	MsgTypePGWDownlinkTriggeringNotification
	MsgTypePGWDownlinkTriggeringAcknowledge
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
	_
	_
	_
	_
	_ // 105-127: Reserved for S5, S4/S11 interfaces
	MsgTypeIdentificationRequest
	MsgTypeIdentificationResponse
	MsgTypeContextRequest
	MsgTypeContextResponse
	MsgTypeContextAcknowledge
	MsgTypeForwardRelocationRequest
	MsgTypeForwardRelocationResponse
	MsgTypeForwardRelocationCompleteNotification
	MsgTypeForwardRelocationCompleteAcknowledge
	MsgTypeForwardAccessContextNotification
	MsgTypeForwardAccessContextAcknowledge
	MsgTypeRelocationCancelRequest
	MsgTypeRelocationCancelResponse
	MsgTypeConfigurationTransferTunnel
	_
	_
	_
	_
	_
	_
	_ // 142-148: Reserved for S3/S10/S16 interfaces
	MsgTypeDetachNotification
	MsgTypeDetachAcknowledge
	MsgTypeCSPagingIndication
	MsgTypeRANInformationRelay
	MsgTypeAlertMMENotification
	MsgTypeAlertMMEAcknowledge
	MsgTypeUEActivityNotification
	MsgTypeUEActivityAcknowledge
	MsgTypeISRStatusIndication
	MsgTypeUERegistrationQueryRequest
	MsgTypeUERegistrationQueryResponse
	MsgTypeCreateForwardingTunnelRequest
	MsgTypeCreateForwardingTunnelResponse
	MsgTypeSuspendNotification
	MsgTypeSuspendAcknowledge
	MsgTypeResumeNotification
	MsgTypeResumeAcknowledge
	MsgTypeCreateIndirectDataForwardingTunnelRequest
	MsgTypeCreateIndirectDataForwardingTunnelResponse
	MsgTypeDeleteIndirectDataForwardingTunnelRequest
	MsgTypeDeleteIndirectDataForwardingTunnelResponse
	MsgTypeReleaseAccessBearersRequest
	MsgTypeReleaseAccessBearersResponse
	_
	_
	_
	_ // 172-175: Reserved for S4/S11 interfaces
	MsgTypeDownlinkDataNotification
	MsgTypeDownlinkDataNotificationAcknowledge
	_
	MsgTypePGWRestartNotification
	MsgTypePGWRestartNotificationAcknowledge
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
	_ // 181-199: Reserved for S4 interface
	MsgTypeUpdatePDNConnectionSetRequest
	MsgTypeUpdatePDNConnectionSetResponse
	_
	_
	_
	_
	_
	_
	_
	_
	_ // 202-210: Reserved for S5/S8 interfaces
	MsgTypeModifyAccessBearersRequest
	MsgTypeModifyAccessBearersResponse
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
	_ // 213-230: Reserved for S11 interface
	MsgTypeMBMSSessionStartRequest
	MsgTypeMBMSSessionStartResponse
	MsgTypeMBMSSessionUpdateRequest
	MsgTypeMBMSSessionUpdateResponse
	MsgTypeMBMSSessionStopRequest
	MsgTypeMBMSSessionStopResponse
	_
	_
	_ // 237-239: Reserved for Sm/Sn interface
	MsgTypeSRVCCCsToPsResponse
	MsgTypeSRVCCCsToPsCompleteNotification
	MsgTypeSRVCCCsToPsCompleteAcknowledge
	MsgTypeSRVCCCsToPsCancelNotification
	MsgTypeSRVCCCsToPsCancelAcknowledge
	_
	_
	_ // 245-247: Reserved for Sv interface
	_
	_
	_
	_
	_
	_
	_
	_ // 248-255: Reserved for others
)

// Message is an interface that defines GTPv2 messages.
type Message interface {
	MarshalTo([]byte) error
	UnmarshalBinary(b []byte) error
	MarshalLen() int
	Version() int
	MessageType() uint8
	MessageTypeName() string
	TEID() uint32
	SetTEID(uint32)
	Sequence() uint32
	SetSequenceNumber(uint32)

	// deprecated
	SerializeTo([]byte) error
	DecodeFromBytes(b []byte) error
}

// Marshal returns the byte sequence generated from a Message instance.
// Better to use MarshalXxx instead if you know the name of message to be serialized.
func Marshal(m Message) ([]byte, error) {
	b := make([]byte, m.MarshalLen())
	if err := m.MarshalTo(b); err != nil {
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
	case MsgTypeVersionNotSupportedIndication:
		m = &VersionNotSupportedIndication{}
	case MsgTypeCreateSessionRequest:
		m = &CreateSessionRequest{}
	case MsgTypeCreateSessionResponse:
		m = &CreateSessionResponse{}
	case MsgTypeDeleteSessionRequest:
		m = &DeleteSessionRequest{}
	case MsgTypeDeleteSessionResponse:
		m = &DeleteSessionResponse{}
	case MsgTypeModifyBearerCommand:
		m = &ModifyBearerCommand{}
	case MsgTypeModifyBearerFailureIndication:
		m = &ModifyBearerFailureIndication{}
	case MsgTypeDeleteBearerCommand:
		m = &DeleteBearerCommand{}
	case MsgTypeDeleteBearerFailureIndication:
		m = &DeleteBearerFailureIndication{}
	case MsgTypeDeleteBearerRequest:
		m = &DeleteBearerRequest{}
	case MsgTypeCreateBearerRequest:
		m = &CreateBearerRequest{}
	case MsgTypeCreateBearerResponse:
		m = &CreateBearerResponse{}
	case MsgTypeDeleteBearerResponse:
		m = &DeleteBearerResponse{}
	case MsgTypeModifyBearerRequest:
		m = &ModifyBearerRequest{}
	case MsgTypeModifyBearerResponse:
		m = &ModifyBearerResponse{}
	case MsgTypeContextRequest:
		m = &ContextRequest{}
	case MsgTypeContextResponse:
		m = &ContextResponse{}
	case MsgTypeContextAcknowledge:
		m = &ContextAcknowledge{}
	case MsgTypeReleaseAccessBearersRequest:
		m = &ReleaseAccessBearersRequest{}
	case MsgTypeReleaseAccessBearersResponse:
		m = &ReleaseAccessBearersResponse{}
	case MsgTypeStopPagingIndication:
		m = &StopPagingIndication{}
	case MsgTypeModifyAccessBearersRequest:
		m = &ModifyAccessBearersRequest{}
	case MsgTypeModifyAccessBearersResponse:
		m = &ModifyAccessBearersResponse{}
	default:
		m = &Generic{}
	}

	if err := m.UnmarshalBinary(b); err != nil {
		return nil, errors.Wrap(err, "failed to decode GTPv2 Message")
	}
	return m, nil
}
