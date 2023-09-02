// Copyright 2019-2023 go-gtp authors. All rights reserveu.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv1/ie"
)

// UpdatePDPContextRequest is a UpdatePDPContextRequest Header and its IEs above.
type UpdatePDPContextRequest struct {
	*Header
	IMSI                                 *ie.IE
	RAI                                  *ie.IE
	Recovery                             *ie.IE
	TEIDDataI                            *ie.IE
	TEIDCPlane                           *ie.IE
	NSAPI                                *ie.IE
	TraceReference                       *ie.IE
	TraceType                            *ie.IE
	PCO                                  *ie.IE
	SGSNAddressForCPlane                 *ie.IE
	SGSNAddressForUserTraffic            *ie.IE
	AlternativeSGSNAddressForCPlane      *ie.IE
	AlternativeSGSNAddressForUserTraffic *ie.IE
	QoSProfile                           *ie.IE
	TFT                                  *ie.IE
	TriggerID                            *ie.IE
	OMCIdentity                          *ie.IE
	CommonFlags                          *ie.IE
	RATType                              *ie.IE
	ULI                                  *ie.IE
	MSTimeZone                           *ie.IE
	AdditionalTraceInfo                  *ie.IE
	DirectTunnelFlags                    *ie.IE
	EvolvedARPI                          *ie.IE
	ExtendedCommonFlags                  *ie.IE
	UCI                                  *ie.IE
	APNAMBR                              *ie.IE
	SignallingPriorityIndication         *ie.IE
	CNOperatorSelectionEntity            *ie.IE
	IMEI                                 *ie.IE
	PrivateExtension                     *ie.IE
	AdditionalIEs                        []*ie.IE
}

// NewUpdatePDPContextRequest creates a new GTPv1 UpdatePDPContextRequest.
func NewUpdatePDPContextRequest(teid uint32, seq uint16, ies ...*ie.IE) *UpdatePDPContextRequest {
	u := &UpdatePDPContextRequest{
		Header: NewHeader(0x32, MsgTypeUpdatePDPContextRequest, teid, seq, nil),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.IMSI:
			u.IMSI = i
		case ie.RouteingAreaIdentity:
			u.RAI = i
		case ie.Recovery:
			u.Recovery = i
		case ie.TEIDDataI:
			u.TEIDDataI = i
		case ie.TEIDCPlane:
			u.TEIDCPlane = i
		case ie.NSAPI:
			u.NSAPI = i
		case ie.TraceReference:
			u.TraceReference = i
		case ie.TraceType:
			u.TraceType = i
		case ie.ProtocolConfigurationOptions:
			u.PCO = i
		case ie.GSNAddress:
			if u.SGSNAddressForCPlane == nil {
				u.SGSNAddressForCPlane = i
			} else if u.SGSNAddressForUserTraffic == nil {
				u.SGSNAddressForUserTraffic = i
			} else if u.AlternativeSGSNAddressForCPlane == nil {
				u.AlternativeSGSNAddressForCPlane = i
			} else if u.AlternativeSGSNAddressForUserTraffic == nil {
				u.AlternativeSGSNAddressForUserTraffic = i
			}
		case ie.QoSProfile:
			u.QoSProfile = i
		case ie.TrafficFlowTemplate:
			u.TFT = i
		case ie.TriggerID:
			u.TriggerID = i
		case ie.OMCIdentity:
			u.OMCIdentity = i
		case ie.CommonFlags:
			u.CommonFlags = i
		case ie.RATType:
			u.RATType = i
		case ie.UserLocationInformation:
			u.ULI = i
		case ie.MSTimeZone:
			u.MSTimeZone = i
		case ie.AdditionalTraceInfo:
			u.AdditionalTraceInfo = i
		case ie.DirectTunnelFlags:
			u.DirectTunnelFlags = i
		case ie.EvolvedAllocationRetentionPriorityI:
			u.EvolvedARPI = i
		case ie.ExtendedCommonFlags:
			u.ExtendedCommonFlags = i
		case ie.UserCSGInformation:
			u.UCI = i
		case ie.AggregateMaximumBitRate:
			u.APNAMBR = i
		case ie.SignallingPriorityIndication:
			u.SignallingPriorityIndication = i
		case ie.CNOperatorSelectionEntity:
			u.CNOperatorSelectionEntity = i
		case ie.IMEISV:
			u.IMEI = i
		case ie.PrivateExtension:
			u.PrivateExtension = i
		default:
			u.AdditionalIEs = append(u.AdditionalIEs, i)
		}
	}

	u.SetLength()
	return u
}

// Marshal returns the byte sequence generated from a UpdatePDPContextRequest.
func (u *UpdatePDPContextRequest) Marshal() ([]byte, error) {
	b := make([]byte, u.MarshalLen())
	if err := u.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (u *UpdatePDPContextRequest) MarshalTo(b []byte) error {
	if len(b) < u.MarshalLen() {
		return ErrTooShortToMarshal
	}
	u.Header.Payload = make([]byte, u.MarshalLen()-u.Header.MarshalLen())

	offset := 0
	if ie := u.IMSI; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.RAI; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.Recovery; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.TEIDDataI; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.TEIDCPlane; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.NSAPI; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.TraceReference; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.TraceType; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.PCO; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.SGSNAddressForCPlane; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.SGSNAddressForUserTraffic; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.AlternativeSGSNAddressForCPlane; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.AlternativeSGSNAddressForUserTraffic; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.QoSProfile; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.TFT; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.TriggerID; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.OMCIdentity; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.CommonFlags; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.RATType; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.ULI; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.MSTimeZone; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.AdditionalTraceInfo; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.DirectTunnelFlags; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.EvolvedARPI; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.ExtendedCommonFlags; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.UCI; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.APNAMBR; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.SignallingPriorityIndication; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.CNOperatorSelectionEntity; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.IMEI; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := u.PrivateExtension; ie != nil {
		if err := ie.MarshalTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	for _, ie := range u.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.MarshalTo(u.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	u.Header.SetLength()
	return u.Header.MarshalTo(b)
}

// ParseUpdatePDPContextRequest decodes a given byte sequence as a UpdatePDPContextRequest.
func ParseUpdatePDPContextRequest(b []byte) (*UpdatePDPContextRequest, error) {
	u := &UpdatePDPContextRequest{}
	if err := u.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return u, nil
}

// UnmarshalBinary decodes a given byte sequence as a UpdatePDPContextRequest.
func (u *UpdatePDPContextRequest) UnmarshalBinary(b []byte) error {
	var err error
	u.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(u.Header.Payload) < 2 {
		return nil
	}

	ies, err := ie.ParseMultiIEs(u.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.IMSI:
			u.IMSI = i
		case ie.RouteingAreaIdentity:
			u.RAI = i
		case ie.Recovery:
			u.Recovery = i
		case ie.TEIDDataI:
			u.TEIDDataI = i
		case ie.TEIDCPlane:
			u.TEIDCPlane = i
		case ie.NSAPI:
			u.NSAPI = i
		case ie.TraceReference:
			u.TraceReference = i
		case ie.TraceType:
			u.TraceType = i
		case ie.ProtocolConfigurationOptions:
			u.PCO = i
		case ie.GSNAddress:
			if u.SGSNAddressForCPlane == nil {
				u.SGSNAddressForCPlane = i
			} else if u.SGSNAddressForUserTraffic == nil {
				u.SGSNAddressForUserTraffic = i
			} else if u.AlternativeSGSNAddressForCPlane == nil {
				u.AlternativeSGSNAddressForCPlane = i
			} else if u.AlternativeSGSNAddressForUserTraffic == nil {
				u.AlternativeSGSNAddressForUserTraffic = i
			}
		case ie.QoSProfile:
			u.QoSProfile = i
		case ie.TrafficFlowTemplate:
			u.TFT = i
		case ie.TriggerID:
			u.TriggerID = i
		case ie.OMCIdentity:
			u.OMCIdentity = i
		case ie.CommonFlags:
			u.CommonFlags = i
		case ie.RATType:
			u.RATType = i
		case ie.UserLocationInformation:
			u.ULI = i
		case ie.MSTimeZone:
			u.MSTimeZone = i
		case ie.AdditionalTraceInfo:
			u.AdditionalTraceInfo = i
		case ie.DirectTunnelFlags:
			u.DirectTunnelFlags = i
		case ie.EvolvedAllocationRetentionPriorityI:
			u.EvolvedARPI = i
		case ie.ExtendedCommonFlags:
			u.ExtendedCommonFlags = i
		case ie.UserCSGInformation:
			u.UCI = i
		case ie.AggregateMaximumBitRate:
			u.APNAMBR = i
		case ie.SignallingPriorityIndication:
			u.SignallingPriorityIndication = i
		case ie.CNOperatorSelectionEntity:
			u.CNOperatorSelectionEntity = i
		case ie.IMEISV:
			u.IMEI = i
		case ie.PrivateExtension:
			u.PrivateExtension = i
		default:
			u.AdditionalIEs = append(u.AdditionalIEs, i)
		}
	}
	return nil
}

// MarshalLen returns the serial length of Data.
func (u *UpdatePDPContextRequest) MarshalLen() int {
	l := u.Header.MarshalLen() - len(u.Header.Payload)

	if ie := u.IMSI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.RAI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.Recovery; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.TEIDDataI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.TEIDCPlane; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.NSAPI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.TraceReference; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.TraceType; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.PCO; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.SGSNAddressForCPlane; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.SGSNAddressForUserTraffic; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.AlternativeSGSNAddressForCPlane; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.AlternativeSGSNAddressForUserTraffic; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.QoSProfile; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.TFT; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.TriggerID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.OMCIdentity; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.CommonFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.RATType; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.ULI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.MSTimeZone; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.AdditionalTraceInfo; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.DirectTunnelFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.EvolvedARPI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.ExtendedCommonFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.UCI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.APNAMBR; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.SignallingPriorityIndication; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.CNOperatorSelectionEntity; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.IMEI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := u.PrivateExtension; ie != nil {
		l += ie.MarshalLen()
	}

	for _, ie := range u.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.MarshalLen()
	}
	return l
}

// SetLength sets the length in Length field.
func (u *UpdatePDPContextRequest) SetLength() {
	u.Length = uint16(u.MarshalLen() - 8)
}

// MessageTypeName returns the name of protocol.
func (u *UpdatePDPContextRequest) MessageTypeName() string {
	return "Update PDP Context Request"
}

// TEID returns the TEID in human-readable string.
func (u *UpdatePDPContextRequest) TEID() uint32 {
	return u.Header.TEID
}
