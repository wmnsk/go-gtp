// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

/*
Package message provides encoding/decoding feature of GTPv2 protocol.
*/
package message

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/wmnsk/go-gtp/gtpv2/ie"
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

// Message is an interface that defines GTPv2 message.
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

	if len(b) < 2 {
		return nil, io.ErrUnexpectedEOF
	}

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
	case MsgTypeUpdateBearerRequest:
		m = &UpdateBearerRequest{}
	case MsgTypeUpdateBearerResponse:
		m = &UpdateBearerResponse{}
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
	case MsgTypeDeletePDNConnectionSetRequest:
		m = &DeletePDNConnectionSetRequest{}
	case MsgTypeDeletePDNConnectionSetResponse:
		m = &DeletePDNConnectionSetResponse{}
	case MsgTypeUpdatePDNConnectionSetRequest:
		m = &UpdatePDNConnectionSetRequest{}
	case MsgTypeUpdatePDNConnectionSetResponse:
		m = &UpdatePDNConnectionSetResponse{}
	case MsgTypePGWRestartNotification:
		m = &PGWRestartNotification{}
	case MsgTypePGWRestartNotificationAcknowledge:
		m = &PGWRestartNotificationAcknowledge{}
	case MsgTypeDetachNotification:
		m = &DetachNotification{}
	case MsgTypeDetachAcknowledge:
		m = &DetachAcknowledge{}
	case MsgTypeResumeAcknowledge:
		m = &ResumeAcknowledge{}
	case MsgTypeResumeNotification:
		m = &ResumeNotification{}
	case MsgTypeSuspendAcknowledge:
		m = &SuspendAcknowledge{}
	case MsgTypeSuspendNotification:
		m = &SuspendNotification{}
	case MsgTypeChangeNotificationRequest:
		m = &ChangeNotificationRequest{}
	case MsgTypeChangeNotificationResponse:
		m = &ChangeNotificationResponse{}
	case MsgTypeDownlinkDataNotification:
		m = &DownlinkDataNotification{}
	case MsgTypeDownlinkDataNotificationAcknowledge:
		m = &DownlinkDataNotificationAcknowledge{}
	case MsgTypeDownlinkDataNotificationFailureIndication:
		m = &DownlinkDataNotificationFailureIndication{}
	default:
		m = &Generic{}
	}

	if err := m.UnmarshalBinary(b); err != nil {
		return nil, fmt.Errorf("failed to decode GTPv2 Message: %w", err)
	}
	return m, nil
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
	if i == nil {
		return fmt.Sprintf("{%s: %v}", name, i)
	}

	if i.IsGrouped() {
		vals := make([]string, len(i.ChildIEs))
		for i, val := range i.ChildIEs {
			vals[i] = fmt.Sprint(val)
		}
		return fmt.Sprintf("{%s: [%v]}", name, strings.Join(vals, ", "))
	}

	return fmt.Sprintf("{%s: %v}", name, i)
}
