// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/v1/ies"
)

// CreatePDPContextRequest is a CreatePDPContextRequest Header and its IEs above.
type CreatePDPContextRequest struct {
	*Header
	IMSI                               *ies.IE
	RAI                                *ies.IE
	Recovery                           *ies.IE
	SelectionMode                      *ies.IE
	TEIDDataI                          *ies.IE
	TEIDCPlane                         *ies.IE
	NSAPI                              *ies.IE
	LinkedNSAPI                        *ies.IE
	ChargingCharacteristics            *ies.IE
	TraceReference                     *ies.IE
	TraceType                          *ies.IE
	EndUserAddress                     *ies.IE
	APN                                *ies.IE
	PCO                                *ies.IE
	SGSNAddressForSignalling           *ies.IE
	SGSNAddressForUserTraffic          *ies.IE
	MSISDN                             *ies.IE
	QoSProfile                         *ies.IE
	TFT                                *ies.IE
	TriggerID                          *ies.IE
	OMCIdentity                        *ies.IE
	CommonFlags                        *ies.IE
	APNRestriction                     *ies.IE
	RATType                            *ies.IE
	UserLocationInformation            *ies.IE
	MSTimeZone                         *ies.IE
	IMEI                               *ies.IE
	CAMELChargingInformationContainer  *ies.IE
	AdditionalTraceInfo                *ies.IE
	CorrelationID                      *ies.IE
	EvolvedARPI                        *ies.IE
	ExtendedCommonFlags                *ies.IE
	UCI                                *ies.IE
	APNAMBR                            *ies.IE
	SignallingPriorityIndication       *ies.IE
	CNOperatorSelectionEntity          *ies.IE
	MappedUEUsageType                  *ies.IE
	UPFunctionSelectionIndicationFlags *ies.IE
	PrivateExtension                   *ies.IE
	AdditionalIEs                      []*ies.IE
}

// NewCreatePDPContextRequest creates a new GTPv1 CreatePDPContextRequest.
func NewCreatePDPContextRequest(teid uint32, seq uint16, ie ...*ies.IE) *CreatePDPContextRequest {
	c := &CreatePDPContextRequest{
		Header: NewHeader(0x32, MsgTypeCreatePDPContextRequest, teid, seq, nil),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.IMSI:
			c.IMSI = i
		case ies.RouteingAreaIdentity:
			c.RAI = i
		case ies.Recovery:
			c.Recovery = i
		case ies.SelectionMode:
			c.SelectionMode = i
		case ies.TEIDDataI:
			c.TEIDDataI = i
		case ies.TEIDCPlane:
			c.TEIDCPlane = i
		case ies.NSAPI:
			if c.NSAPI == nil {
				c.NSAPI = i
			} else if c.LinkedNSAPI == nil {
				c.LinkedNSAPI = i
			} else {
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.ChargingCharacteristics:
			c.ChargingCharacteristics = i
		case ies.TraceReference:
			c.TraceReference = i
		case ies.TraceType:
			c.TraceType = i
		case ies.EndUserAddress:
			c.EndUserAddress = i
		case ies.AccessPointName:
			c.APN = i
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.GSNAddress:
			if c.SGSNAddressForSignalling == nil {
				c.SGSNAddressForSignalling = i
			} else if c.SGSNAddressForUserTraffic == nil {
				c.SGSNAddressForUserTraffic = i
			} else {
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.MSISDN:
			c.MSISDN = i
		case ies.QoSProfile:
			c.QoSProfile = i
		case ies.TrafficFlowTemplate:
			c.TFT = i
		case ies.TriggerID:
			c.TriggerID = i
		case ies.OMCIdentity:
			c.OMCIdentity = i
		case ies.CommonFlags:
			c.CommonFlags = i
		case ies.APNRestriction:
			c.APNRestriction = i
		case ies.RATType:
			c.RATType = i
		case ies.UserLocationInformation:
			c.UserLocationInformation = i
		case ies.MSTimeZone:
			c.MSTimeZone = i
		case ies.IMEISV:
			c.IMEI = i
		case ies.CAMELChargingInformationContainer:
			c.CAMELChargingInformationContainer = i
		case ies.AdditionalTraceInfo:
			c.AdditionalTraceInfo = i
		case ies.CorrelationID:
			c.CorrelationID = i
		case ies.EvolvedAllocationRetentionPriorityI:
			c.EvolvedARPI = i
		case ies.ExtendedCommonFlags:
			c.ExtendedCommonFlags = i
		case ies.UserCSGInformation:
			c.UCI = i
		case ies.AggregateMaximumBitRate:
			c.APNAMBR = i
		case ies.SignallingPriorityIndication:
			c.SignallingPriorityIndication = i
		case ies.CNOperatorSelectionEntity:
			c.CNOperatorSelectionEntity = i
		case ies.MappedUEUsageType:
			c.MappedUEUsageType = i
		case ies.UPFunctionSelectionIndicationFlags:
			c.UPFunctionSelectionIndicationFlags = i
		case ies.PrivateExtension:
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

	ie, err := ies.ParseMultiIEs(c.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.IMSI:
			c.IMSI = i
		case ies.RouteingAreaIdentity:
			c.RAI = i
		case ies.Recovery:
			c.Recovery = i
		case ies.SelectionMode:
			c.SelectionMode = i
		case ies.TEIDDataI:
			c.TEIDDataI = i
		case ies.TEIDCPlane:
			c.TEIDCPlane = i
		case ies.NSAPI:
			if c.NSAPI == nil {
				c.NSAPI = i
			} else if c.LinkedNSAPI == nil {
				c.LinkedNSAPI = i
			} else {
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.ChargingCharacteristics:
			c.ChargingCharacteristics = i
		case ies.TraceReference:
			c.TraceReference = i
		case ies.TraceType:
			c.TraceType = i
		case ies.EndUserAddress:
			c.EndUserAddress = i
		case ies.AccessPointName:
			c.APN = i
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.GSNAddress:
			if c.SGSNAddressForSignalling == nil {
				c.SGSNAddressForSignalling = i
			} else if c.SGSNAddressForUserTraffic == nil {
				c.SGSNAddressForUserTraffic = i
			} else {
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.MSISDN:
			c.MSISDN = i
		case ies.QoSProfile:
			c.QoSProfile = i
		case ies.TrafficFlowTemplate:
			c.TFT = i
		case ies.TriggerID:
			c.TriggerID = i
		case ies.OMCIdentity:
			c.OMCIdentity = i
		case ies.CommonFlags:
			c.CommonFlags = i
		case ies.APNRestriction:
			c.APNRestriction = i
		case ies.RATType:
			c.RATType = i
		case ies.UserLocationInformation:
			c.UserLocationInformation = i
		case ies.MSTimeZone:
			c.MSTimeZone = i
		case ies.IMEISV:
			c.IMEI = i
		case ies.CAMELChargingInformationContainer:
			c.CAMELChargingInformationContainer = i
		case ies.AdditionalTraceInfo:
			c.AdditionalTraceInfo = i
		case ies.CorrelationID:
			c.CorrelationID = i
		case ies.EvolvedAllocationRetentionPriorityI:
			c.EvolvedARPI = i
		case ies.ExtendedCommonFlags:
			c.ExtendedCommonFlags = i
		case ies.UserCSGInformation:
			c.UCI = i
		case ies.AggregateMaximumBitRate:
			c.APNAMBR = i
		case ies.SignallingPriorityIndication:
			c.SignallingPriorityIndication = i
		case ies.CNOperatorSelectionEntity:
			c.CNOperatorSelectionEntity = i
		case ies.MappedUEUsageType:
			c.MappedUEUsageType = i
		case ies.UPFunctionSelectionIndicationFlags:
			c.UPFunctionSelectionIndicationFlags = i
		case ies.PrivateExtension:
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
