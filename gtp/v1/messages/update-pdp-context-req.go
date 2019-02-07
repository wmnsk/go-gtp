// Copyright 2019 go-gtp authors. All rights reserveu.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/gtp/v1/ies"
)

// UpdatePDPContextRequest is a UpdatePDPContextRequest Header and its IEs above.
type UpdatePDPContextRequest struct {
	*Header
	IMSI                                 *ies.IE
	RAI                                  *ies.IE
	Recovery                             *ies.IE
	TEIDDataI                            *ies.IE
	TEIDCPlane                           *ies.IE
	NSAPI                                *ies.IE
	TraceReference                       *ies.IE
	TraceType                            *ies.IE
	PCO                                  *ies.IE
	SGSNAddressForCPlane                 *ies.IE
	SGSNAddressForUserTraffic            *ies.IE
	AlternativeSGSNAddressForCPlane      *ies.IE
	AlternativeSGSNAddressForUserTraffic *ies.IE
	QoSProfile                           *ies.IE
	TFT                                  *ies.IE
	TriggerID                            *ies.IE
	OMCIdentity                          *ies.IE
	CommonFlags                          *ies.IE
	RATType                              *ies.IE
	ULI                                  *ies.IE
	MSTimeZone                           *ies.IE
	AdditionalTraceInfo                  *ies.IE
	DirectTunnelFlags                    *ies.IE
	EvolvedARPI                          *ies.IE
	ExtendedCommonFlags                  *ies.IE
	UCI                                  *ies.IE
	APNAMBR                              *ies.IE
	SignallingPriorityIndication         *ies.IE
	CNOperatorSelectionEntity            *ies.IE
	IMEI                                 *ies.IE
	PrivateExtension                     *ies.IE
	AdditionalIEs                        []*ies.IE
}

// NewUpdatePDPContextRequest creates a new GTPv1 UpdatePDPContextRequest.
func NewUpdatePDPContextRequest(teid uint32, seq uint16, ie ...*ies.IE) *UpdatePDPContextRequest {
	u := &UpdatePDPContextRequest{
		Header: NewHeader(0x32, MsgTypeUpdatePDPContextRequest, teid, seq, nil),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.IMSI:
			u.IMSI = i
		case ies.RouteingAreaIdentity:
			u.RAI = i
		case ies.Recovery:
			u.Recovery = i
		case ies.TEIDDataI:
			u.TEIDDataI = i
		case ies.TEIDCPlane:
			u.TEIDCPlane = i
		case ies.NSAPI:
			u.NSAPI = i
		case ies.TraceReference:
			u.TraceReference = i
		case ies.TraceType:
			u.TraceType = i
		case ies.ProtocolConfigurationOptions:
			u.PCO = i
		case ies.GSNAddress:
			if u.SGSNAddressForCPlane == nil {
				u.SGSNAddressForCPlane = i
			} else if u.SGSNAddressForUserTraffic == nil {
				u.SGSNAddressForUserTraffic = i
			} else if u.AlternativeSGSNAddressForCPlane == nil {
				u.AlternativeSGSNAddressForCPlane = i
			} else if u.AlternativeSGSNAddressForUserTraffic == nil {
				u.AlternativeSGSNAddressForUserTraffic = i
			}
		case ies.QoSProfile:
			u.QoSProfile = i
		case ies.TrafficFlowTemplate:
			u.TFT = i
		case ies.TriggerID:
			u.TriggerID = i
		case ies.OMCIdentity:
			u.OMCIdentity = i
		case ies.CommonFlags:
			u.CommonFlags = i
		case ies.RATType:
			u.RATType = i
		case ies.UserLocationInformation:
			u.ULI = i
		case ies.MSTimeZone:
			u.MSTimeZone = i
		case ies.AdditionalTraceInfo:
			u.AdditionalTraceInfo = i
		case ies.DirectTunnelFlags:
			u.DirectTunnelFlags = i
		case ies.EvolvedAllocationRetentionPriorityI:
			u.EvolvedARPI = i
		case ies.ExtendedCommonFlags:
			u.ExtendedCommonFlags = i
		case ies.UserCSGInformation:
			u.UCI = i
		case ies.AggregateMaximumBitRate:
			u.APNAMBR = i
		case ies.SignallingPriorityIndication:
			u.SignallingPriorityIndication = i
		case ies.CNOperatorSelectionEntity:
			u.CNOperatorSelectionEntity = i
		case ies.IMEISV:
			u.IMEI = i
		case ies.PrivateExtension:
			u.PrivateExtension = i
		default:
			u.AdditionalIEs = append(u.AdditionalIEs, i)
		}
	}

	u.SetLength()
	return u
}

// Serialize returns the byte sequence generated from a UpdatePDPContextRequest.
func (u *UpdatePDPContextRequest) Serialize() ([]byte, error) {
	b := make([]byte, u.Len())
	if err := u.SerializeTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// SerializeTo puts the byte sequence in the byte array given as b.
func (u *UpdatePDPContextRequest) SerializeTo(b []byte) error {
	if len(b) < u.Len() {
		return ErrTooShortToSerialize
	}
	u.Header.Payload = make([]byte, u.Len()-u.Header.Len())

	offset := 0
	if ie := u.IMSI; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.RAI; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.Recovery; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.TEIDDataI; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.TEIDCPlane; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.NSAPI; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.TraceReference; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.TraceType; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.PCO; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.SGSNAddressForCPlane; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.SGSNAddressForUserTraffic; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.AlternativeSGSNAddressForCPlane; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.AlternativeSGSNAddressForUserTraffic; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.QoSProfile; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.TFT; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.TriggerID; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.OMCIdentity; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.CommonFlags; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.RATType; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.ULI; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.MSTimeZone; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.AdditionalTraceInfo; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.DirectTunnelFlags; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.EvolvedARPI; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.ExtendedCommonFlags; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.UCI; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.APNAMBR; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.SignallingPriorityIndication; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.CNOperatorSelectionEntity; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.IMEI; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := u.PrivateExtension; ie != nil {
		if err := ie.SerializeTo(u.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	for _, ie := range u.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.SerializeTo(u.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	u.Header.SetLength()
	return u.Header.SerializeTo(b)
}

// DecodeUpdatePDPContextRequest decodes a given byte sequence as a UpdatePDPContextRequest.
func DecodeUpdatePDPContextRequest(b []byte) (*UpdatePDPContextRequest, error) {
	u := &UpdatePDPContextRequest{}
	if err := u.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return u, nil
}

// DecodeFromBytes decodes a given byte sequence as a UpdatePDPContextRequest.
func (u *UpdatePDPContextRequest) DecodeFromBytes(b []byte) error {
	var err error
	u.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}
	if len(u.Header.Payload) < 2 {
		return nil
	}

	ie, err := ies.DecodeMultiIEs(u.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.IMSI:
			u.IMSI = i
		case ies.RouteingAreaIdentity:
			u.RAI = i
		case ies.Recovery:
			u.Recovery = i
		case ies.TEIDDataI:
			u.TEIDDataI = i
		case ies.TEIDCPlane:
			u.TEIDCPlane = i
		case ies.NSAPI:
			u.NSAPI = i
		case ies.TraceReference:
			u.TraceReference = i
		case ies.TraceType:
			u.TraceType = i
		case ies.ProtocolConfigurationOptions:
			u.PCO = i
		case ies.GSNAddress:
			if u.SGSNAddressForCPlane == nil {
				u.SGSNAddressForCPlane = i
			} else if u.SGSNAddressForUserTraffic == nil {
				u.SGSNAddressForUserTraffic = i
			} else if u.AlternativeSGSNAddressForCPlane == nil {
				u.AlternativeSGSNAddressForCPlane = i
			} else if u.AlternativeSGSNAddressForUserTraffic == nil {
				u.AlternativeSGSNAddressForUserTraffic = i
			}
		case ies.QoSProfile:
			u.QoSProfile = i
		case ies.TrafficFlowTemplate:
			u.TFT = i
		case ies.TriggerID:
			u.TriggerID = i
		case ies.OMCIdentity:
			u.OMCIdentity = i
		case ies.CommonFlags:
			u.CommonFlags = i
		case ies.RATType:
			u.RATType = i
		case ies.UserLocationInformation:
			u.ULI = i
		case ies.MSTimeZone:
			u.MSTimeZone = i
		case ies.AdditionalTraceInfo:
			u.AdditionalTraceInfo = i
		case ies.DirectTunnelFlags:
			u.DirectTunnelFlags = i
		case ies.EvolvedAllocationRetentionPriorityI:
			u.EvolvedARPI = i
		case ies.ExtendedCommonFlags:
			u.ExtendedCommonFlags = i
		case ies.UserCSGInformation:
			u.UCI = i
		case ies.AggregateMaximumBitRate:
			u.APNAMBR = i
		case ies.SignallingPriorityIndication:
			u.SignallingPriorityIndication = i
		case ies.CNOperatorSelectionEntity:
			u.CNOperatorSelectionEntity = i
		case ies.IMEISV:
			u.IMEI = i
		case ies.PrivateExtension:
			u.PrivateExtension = i
		default:
			u.AdditionalIEs = append(u.AdditionalIEs, i)
		}
	}
	return nil
}

// Len returns the actual length of Data.
func (u *UpdatePDPContextRequest) Len() int {
	l := u.Header.Len() - len(u.Header.Payload)

	if ie := u.IMSI; ie != nil {
		l += ie.Len()
	}
	if ie := u.RAI; ie != nil {
		l += ie.Len()
	}
	if ie := u.Recovery; ie != nil {
		l += ie.Len()
	}
	if ie := u.TEIDDataI; ie != nil {
		l += ie.Len()
	}
	if ie := u.TEIDCPlane; ie != nil {
		l += ie.Len()
	}
	if ie := u.NSAPI; ie != nil {
		l += ie.Len()
	}
	if ie := u.TraceReference; ie != nil {
		l += ie.Len()
	}
	if ie := u.TraceType; ie != nil {
		l += ie.Len()
	}
	if ie := u.PCO; ie != nil {
		l += ie.Len()
	}
	if ie := u.SGSNAddressForCPlane; ie != nil {
		l += ie.Len()
	}
	if ie := u.SGSNAddressForUserTraffic; ie != nil {
		l += ie.Len()
	}
	if ie := u.AlternativeSGSNAddressForCPlane; ie != nil {
		l += ie.Len()
	}
	if ie := u.AlternativeSGSNAddressForUserTraffic; ie != nil {
		l += ie.Len()
	}
	if ie := u.QoSProfile; ie != nil {
		l += ie.Len()
	}
	if ie := u.TFT; ie != nil {
		l += ie.Len()
	}
	if ie := u.TriggerID; ie != nil {
		l += ie.Len()
	}
	if ie := u.OMCIdentity; ie != nil {
		l += ie.Len()
	}
	if ie := u.CommonFlags; ie != nil {
		l += ie.Len()
	}
	if ie := u.RATType; ie != nil {
		l += ie.Len()
	}
	if ie := u.ULI; ie != nil {
		l += ie.Len()
	}
	if ie := u.MSTimeZone; ie != nil {
		l += ie.Len()
	}
	if ie := u.AdditionalTraceInfo; ie != nil {
		l += ie.Len()
	}
	if ie := u.DirectTunnelFlags; ie != nil {
		l += ie.Len()
	}
	if ie := u.EvolvedARPI; ie != nil {
		l += ie.Len()
	}
	if ie := u.ExtendedCommonFlags; ie != nil {
		l += ie.Len()
	}
	if ie := u.UCI; ie != nil {
		l += ie.Len()
	}
	if ie := u.APNAMBR; ie != nil {
		l += ie.Len()
	}
	if ie := u.SignallingPriorityIndication; ie != nil {
		l += ie.Len()
	}
	if ie := u.CNOperatorSelectionEntity; ie != nil {
		l += ie.Len()
	}
	if ie := u.IMEI; ie != nil {
		l += ie.Len()
	}
	if ie := u.PrivateExtension; ie != nil {
		l += ie.Len()
	}

	for _, ie := range u.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.Len()
	}
	return l
}

// SetLength sets the length in Length field.
func (u *UpdatePDPContextRequest) SetLength() {
	u.Length = uint16(u.Len() - 8)
}

// MessageTypeName returns the name of protocol.
func (u *UpdatePDPContextRequest) MessageTypeName() string {
	return "Update PDP Context Request"
}

// TEID returns the TEID in human-readable string.
func (u *UpdatePDPContextRequest) TEID() uint32 {
	return u.Header.TEID
}
