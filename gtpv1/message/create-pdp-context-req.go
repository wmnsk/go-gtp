// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv1/ie"
)

// CreatePDPContextRequest is a CreatePDPContextRequest Header and its IEs above.
type CreatePDPContextRequest struct {
	*Header
	IMSI                               *ie.IE
	RAI                                *ie.IE
	Recovery                           *ie.IE
	SelectionMode                      *ie.IE
	TEIDDataI                          *ie.IE
	TEIDCPlane                         *ie.IE
	NSAPI                              *ie.IE
	LinkedNSAPI                        *ie.IE
	ChargingCharacteristics            *ie.IE
	TraceReference                     *ie.IE
	TraceType                          *ie.IE
	EndUserAddress                     *ie.IE
	APN                                *ie.IE
	PCO                                *ie.IE
	SGSNAddressForSignalling           *ie.IE
	SGSNAddressForUserTraffic          *ie.IE
	MSISDN                             *ie.IE
	QoSProfile                         *ie.IE
	TFT                                *ie.IE
	TriggerID                          *ie.IE
	OMCIdentity                        *ie.IE
	CommonFlags                        *ie.IE
	APNRestriction                     *ie.IE
	RATType                            *ie.IE
	UserLocationInformation            *ie.IE
	MSTimeZone                         *ie.IE
	IMEI                               *ie.IE
	CAMELChargingInformationContainer  *ie.IE
	AdditionalTraceInfo                *ie.IE
	CorrelationID                      *ie.IE
	EvolvedARPI                        *ie.IE
	ExtendedCommonFlags                *ie.IE
	UCI                                *ie.IE
	APNAMBR                            *ie.IE
	SignallingPriorityIndication       *ie.IE
	CNOperatorSelectionEntity          *ie.IE
	MappedUEUsageType                  *ie.IE
	UPFunctionSelectionIndicationFlags *ie.IE
	PrivateExtension                   *ie.IE
	AdditionalIEs                      []*ie.IE
}

// NewCreatePDPContextRequest creates a new GTPv1 CreatePDPContextRequest.
func NewCreatePDPContextRequest(teid uint32, seq uint16, ies ...*ie.IE) *CreatePDPContextRequest {
	c := &CreatePDPContextRequest{
		Header: NewHeader(0x32, MsgTypeCreatePDPContextRequest, teid, seq, nil),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.IMSI:
			c.IMSI = i
		case ie.RouteingAreaIdentity:
			c.RAI = i
		case ie.Recovery:
			c.Recovery = i
		case ie.SelectionMode:
			c.SelectionMode = i
		case ie.TEIDDataI:
			c.TEIDDataI = i
		case ie.TEIDCPlane:
			c.TEIDCPlane = i
		case ie.NSAPI:
			if c.NSAPI == nil {
				c.NSAPI = i
			} else if c.LinkedNSAPI == nil {
				c.LinkedNSAPI = i
			} else {
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.ChargingCharacteristics:
			c.ChargingCharacteristics = i
		case ie.TraceReference:
			c.TraceReference = i
		case ie.TraceType:
			c.TraceType = i
		case ie.EndUserAddress:
			c.EndUserAddress = i
		case ie.AccessPointName:
			c.APN = i
		case ie.ProtocolConfigurationOptions:
			c.PCO = i
		case ie.GSNAddress:
			if c.SGSNAddressForSignalling == nil {
				c.SGSNAddressForSignalling = i
			} else if c.SGSNAddressForUserTraffic == nil {
				c.SGSNAddressForUserTraffic = i
			} else {
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.MSISDN:
			c.MSISDN = i
		case ie.QoSProfile:
			c.QoSProfile = i
		case ie.TrafficFlowTemplate:
			c.TFT = i
		case ie.TriggerID:
			c.TriggerID = i
		case ie.OMCIdentity:
			c.OMCIdentity = i
		case ie.CommonFlags:
			c.CommonFlags = i
		case ie.APNRestriction:
			c.APNRestriction = i
		case ie.RATType:
			c.RATType = i
		case ie.UserLocationInformation:
			c.UserLocationInformation = i
		case ie.MSTimeZone:
			c.MSTimeZone = i
		case ie.IMEISV:
			c.IMEI = i
		case ie.CAMELChargingInformationContainer:
			c.CAMELChargingInformationContainer = i
		case ie.AdditionalTraceInfo:
			c.AdditionalTraceInfo = i
		case ie.CorrelationID:
			c.CorrelationID = i
		case ie.EvolvedAllocationRetentionPriorityI:
			c.EvolvedARPI = i
		case ie.ExtendedCommonFlags:
			c.ExtendedCommonFlags = i
		case ie.UserCSGInformation:
			c.UCI = i
		case ie.AggregateMaximumBitRate:
			c.APNAMBR = i
		case ie.SignallingPriorityIndication:
			c.SignallingPriorityIndication = i
		case ie.CNOperatorSelectionEntity:
			c.CNOperatorSelectionEntity = i
		case ie.MappedUEUsageType:
			c.MappedUEUsageType = i
		case ie.UPFunctionSelectionIndicationFlags:
			c.UPFunctionSelectionIndicationFlags = i
		case ie.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	c.SetLength()
	return c
}

// Marshal returns the byte sequence generated from a CreatePDPContextRequest.
func (c *CreatePDPContextRequest) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (c *CreatePDPContextRequest) MarshalTo(b []byte) error {
	if len(b) < c.MarshalLen() {
		return ErrTooShortToMarshal
	}
	c.Header.Payload = make([]byte, c.MarshalLen()-c.Header.MarshalLen())

	offset := 0
	if ie := c.IMSI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.RAI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.Recovery; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SelectionMode; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.TEIDDataI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.TEIDCPlane; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.NSAPI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.LinkedNSAPI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.ChargingCharacteristics; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.TraceReference; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.TraceType; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.EndUserAddress; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.APN; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PCO; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SGSNAddressForSignalling; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SGSNAddressForUserTraffic; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MSISDN; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.QoSProfile; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.TFT; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.TriggerID; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.OMCIdentity; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.CommonFlags; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.APNRestriction; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.RATType; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.UserLocationInformation; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MSTimeZone; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.IMEI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.CAMELChargingInformationContainer; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.AdditionalTraceInfo; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.CorrelationID; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.EvolvedARPI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.ExtendedCommonFlags; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.UCI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.APNAMBR; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SignallingPriorityIndication; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.CNOperatorSelectionEntity; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MappedUEUsageType; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.UPFunctionSelectionIndicationFlags; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PrivateExtension; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	for _, ie := range c.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.MarshalTo(c.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	c.Header.SetLength()
	return c.Header.MarshalTo(b)
}

// ParseCreatePDPContextRequest decodes a given byte sequence as a CreatePDPContextRequest.
func ParseCreatePDPContextRequest(b []byte) (*CreatePDPContextRequest, error) {
	c := &CreatePDPContextRequest{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary decodes a given byte sequence as a CreatePDPContextRequest.
func (c *CreatePDPContextRequest) UnmarshalBinary(b []byte) error {
	var err error
	c.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(c.Header.Payload) < 2 {
		return nil
	}

	ies, err := ie.ParseMultiIEs(c.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.IMSI:
			c.IMSI = i
		case ie.RouteingAreaIdentity:
			c.RAI = i
		case ie.Recovery:
			c.Recovery = i
		case ie.SelectionMode:
			c.SelectionMode = i
		case ie.TEIDDataI:
			c.TEIDDataI = i
		case ie.TEIDCPlane:
			c.TEIDCPlane = i
		case ie.NSAPI:
			if c.NSAPI == nil {
				c.NSAPI = i
			} else if c.LinkedNSAPI == nil {
				c.LinkedNSAPI = i
			} else {
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.ChargingCharacteristics:
			c.ChargingCharacteristics = i
		case ie.TraceReference:
			c.TraceReference = i
		case ie.TraceType:
			c.TraceType = i
		case ie.EndUserAddress:
			c.EndUserAddress = i
		case ie.AccessPointName:
			c.APN = i
		case ie.ProtocolConfigurationOptions:
			c.PCO = i
		case ie.GSNAddress:
			if c.SGSNAddressForSignalling == nil {
				c.SGSNAddressForSignalling = i
			} else if c.SGSNAddressForUserTraffic == nil {
				c.SGSNAddressForUserTraffic = i
			} else {
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.MSISDN:
			c.MSISDN = i
		case ie.QoSProfile:
			c.QoSProfile = i
		case ie.TrafficFlowTemplate:
			c.TFT = i
		case ie.TriggerID:
			c.TriggerID = i
		case ie.OMCIdentity:
			c.OMCIdentity = i
		case ie.CommonFlags:
			c.CommonFlags = i
		case ie.APNRestriction:
			c.APNRestriction = i
		case ie.RATType:
			c.RATType = i
		case ie.UserLocationInformation:
			c.UserLocationInformation = i
		case ie.MSTimeZone:
			c.MSTimeZone = i
		case ie.IMEISV:
			c.IMEI = i
		case ie.CAMELChargingInformationContainer:
			c.CAMELChargingInformationContainer = i
		case ie.AdditionalTraceInfo:
			c.AdditionalTraceInfo = i
		case ie.CorrelationID:
			c.CorrelationID = i
		case ie.EvolvedAllocationRetentionPriorityI:
			c.EvolvedARPI = i
		case ie.ExtendedCommonFlags:
			c.ExtendedCommonFlags = i
		case ie.UserCSGInformation:
			c.UCI = i
		case ie.AggregateMaximumBitRate:
			c.APNAMBR = i
		case ie.SignallingPriorityIndication:
			c.SignallingPriorityIndication = i
		case ie.CNOperatorSelectionEntity:
			c.CNOperatorSelectionEntity = i
		case ie.MappedUEUsageType:
			c.MappedUEUsageType = i
		case ie.UPFunctionSelectionIndicationFlags:
			c.UPFunctionSelectionIndicationFlags = i
		case ie.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}
	return nil
}

// MarshalLen returns the serial length of Data.
func (c *CreatePDPContextRequest) MarshalLen() int {
	l := c.Header.MarshalLen() - len(c.Header.Payload)

	if ie := c.IMSI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.RAI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.Recovery; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SelectionMode; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TEIDDataI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TEIDCPlane; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.NSAPI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.LinkedNSAPI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ChargingCharacteristics; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TraceReference; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TraceType; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.EndUserAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.APN; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PCO; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGSNAddressForSignalling; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGSNAddressForUserTraffic; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MSISDN; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.QoSProfile; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TFT; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TriggerID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.OMCIdentity; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.CommonFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.APNRestriction; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.RATType; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UserLocationInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MSTimeZone; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.IMEI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.CAMELChargingInformationContainer; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.AdditionalTraceInfo; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.CorrelationID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.EvolvedARPI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ExtendedCommonFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UCI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.APNAMBR; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SignallingPriorityIndication; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.CNOperatorSelectionEntity; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MappedUEUsageType; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UPFunctionSelectionIndicationFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PrivateExtension; ie != nil {
		l += ie.MarshalLen()
	}

	for _, ie := range c.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.MarshalLen()
	}
	return l
}

// SetLength sets the length in Length field.
func (c *CreatePDPContextRequest) SetLength() {
	c.Length = uint16(c.MarshalLen() - 8)
}

// MessageTypeName returns the name of protocol.
func (c *CreatePDPContextRequest) MessageTypeName() string {
	return "Create PDP Context Request"
}

// TEID returns the TEID in human-readable string.
func (c *CreatePDPContextRequest) TEID() uint32 {
	return c.Header.TEID
}
