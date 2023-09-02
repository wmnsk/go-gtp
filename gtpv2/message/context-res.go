// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// ContextResponse is a ContextResponse Header and its IEs above.
type ContextResponse struct {
	*Header
	Cause                               *ie.IE
	IMSI                                *ie.IE
	UEMMContext                         *ie.IE
	UEPDNConnections                    []*ie.IE
	SenderFTEID                         *ie.IE
	SGWS11S4FTEID                       *ie.IE
	SGWNodeName                         *ie.IE
	IndicationFlags                     *ie.IE
	TraceInformation                    *ie.IE
	S101IPAddress                       *ie.IE
	S102IPAddress                       *ie.IE
	SubscribedRFSPIndex                 *ie.IE
	RFSPIndexInUse                      *ie.IE
	UETimeZone                          *ie.IE
	MMESGSNLDN                          *ie.IE
	MDTConfiguration                    *ie.IE
	SGSNNodeName                        *ie.IE
	MMENodeName                         *ie.IE
	UCI                                 *ie.IE
	MonitoringEventInformation          *ie.IE
	MonitoringEventExtensionInformation *ie.IE
	UEUsageType                         *ie.IE
	SCEFPDNConnection                   []*ie.IE
	RATType                             *ie.IE
	ServingPLMNRateControl              *ie.IE
	MOExceptionDataCounter              *ie.IE
	RemainingRunningServiceGapTimer     *ie.IE
	ExtendedTraceInformation            *ie.IE
	SubscribedAdditionalRRMPolicyIndex  *ie.IE
	AdditionalRRMPolicyIndexInUse       *ie.IE
	PrivateExtension                    *ie.IE
	AdditionalIEs                       []*ie.IE
}

// NewContextResponse creates a new ContextResponse.
func NewContextResponse(teid, seq uint32, ies ...*ie.IE) *ContextResponse {
	c := &ContextResponse{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeContextResponse, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			c.Cause = i
		case ie.IMSI:
			c.IMSI = i
		case ie.MMContextEPSSecurityContextQuadrupletsAndQuintuplets,
			ie.MMContextGSMKeyAndTriplets, ie.MMContextGSMKeyUsedCipherAndQuintuplets,
			ie.MMContextUMTSKeyAndQuintuplets, ie.MMContextUMTSKeyQuadrupletsAndQuintuplets,
			ie.MMContextUMTSKeyUsedCipherAndQuintuplets:
			if c.UEMMContext != nil {
				c.UEMMContext = i
			} else {
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.PDNConnection:
			c.UEPDNConnections = append(c.UEPDNConnections, i)
		case ie.FullyQualifiedTEID:
			switch i.Instance() {
			case 0:
				c.SenderFTEID = i
			case 1:
				c.SGWS11S4FTEID = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.FullyQualifiedDomainName:
			switch i.Instance() {
			case 0:
				c.SGWNodeName = i
			case 1:
				c.SGSNNodeName = i
			case 2:
				c.MMENodeName = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.Indication:
			c.IndicationFlags = i
		case ie.TraceInformation:
			c.TraceInformation = i
		case ie.IPAddress:
			switch i.Instance() {
			case 0:
				c.S101IPAddress = i
			case 1:
				c.S102IPAddress = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.RFSPIndex:
			switch i.Instance() {
			case 0:
				c.SubscribedRFSPIndex = i
			case 1:
				c.RFSPIndexInUse = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.UETimeZone:
			c.UETimeZone = i
		case ie.LocalDistinguishedName:
			c.MMESGSNLDN = i
		case ie.MDTConfiguration:
			c.MDTConfiguration = i
		case ie.UserCSGInformation:
			c.UCI = i
		case ie.MonitoringEventInformation:
			c.MonitoringEventInformation = i
		case ie.MonitoringEventExtensionInformation:
			c.MonitoringEventExtensionInformation = i
		case ie.IntegerNumber:
			switch i.Instance() {
			case 0:
				c.UEUsageType = i
			case 1:
				c.RemainingRunningServiceGapTimer = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.SCEFPDNConnection:
			c.SCEFPDNConnection = append(c.SCEFPDNConnection, i)
		case ie.RATType:
			c.RATType = i
		case ie.ServingPLMNRateControl:
			c.ServingPLMNRateControl = i
		case ie.Counter:
			c.MOExceptionDataCounter = i
		case ie.ExtendedTraceInformation:
			c.ExtendedTraceInformation = i
		case ie.AdditionalRRMPolicyIndex:
			switch i.Instance() {
			case 0:
				c.SubscribedAdditionalRRMPolicyIndex = i
			case 1:
				c.AdditionalRRMPolicyIndexInUse = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	c.SetLength()
	return c
}

// Marshal serializes ContextResponse into bytes.
func (c *ContextResponse) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes ContextResponse into bytes.
func (c *ContextResponse) MarshalTo(b []byte) error {
	if c.Header.Payload != nil {
		c.Header.Payload = nil
	}
	c.Header.Payload = make([]byte, c.MarshalLen()-c.Header.MarshalLen())

	offset := 0
	if ie := c.Cause; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.IMSI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.UEMMContext; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	for _, ie := range c.UEPDNConnections {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SenderFTEID; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SGWS11S4FTEID; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SGWNodeName; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.IndicationFlags; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.TraceInformation; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.S101IPAddress; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.S102IPAddress; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SubscribedRFSPIndex; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.RFSPIndexInUse; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.UETimeZone; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MMESGSNLDN; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MDTConfiguration; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SGSNNodeName; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MMENodeName; ie != nil {
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
	if ie := c.MonitoringEventInformation; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.UEUsageType; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	for _, ie := range c.SCEFPDNConnection {
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
	if ie := c.ServingPLMNRateControl; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MOExceptionDataCounter; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.RemainingRunningServiceGapTimer; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.ExtendedTraceInformation; ie != nil {
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

// ParseContextResponse decodes given bytes as ContextResponse.
func ParseContextResponse(b []byte) (*ContextResponse, error) {
	c := &ContextResponse{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary decodes given bytes as ContextResponse.
func (c *ContextResponse) UnmarshalBinary(b []byte) error {
	var err error
	c.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(c.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ie.ParseMultiIEs(c.Header.Payload)
	if err != nil {
		return err
	}
	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			c.Cause = i
		case ie.IMSI:
			c.IMSI = i
		case ie.MMContextEPSSecurityContextQuadrupletsAndQuintuplets,
			ie.MMContextGSMKeyAndTriplets, ie.MMContextGSMKeyUsedCipherAndQuintuplets,
			ie.MMContextUMTSKeyAndQuintuplets, ie.MMContextUMTSKeyQuadrupletsAndQuintuplets,
			ie.MMContextUMTSKeyUsedCipherAndQuintuplets:
			if c.UEMMContext != nil {
				c.UEMMContext = i
			} else {
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.PDNConnection:
			c.UEPDNConnections = append(c.UEPDNConnections, i)
		case ie.FullyQualifiedTEID:
			switch i.Instance() {
			case 0:
				c.SenderFTEID = i
			case 1:
				c.SGWS11S4FTEID = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.FullyQualifiedDomainName:
			switch i.Instance() {
			case 0:
				c.SGWNodeName = i
			case 1:
				c.SGSNNodeName = i
			case 2:
				c.MMENodeName = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.Indication:
			c.IndicationFlags = i
		case ie.TraceInformation:
			c.TraceInformation = i
		case ie.IPAddress:
			switch i.Instance() {
			case 0:
				c.S101IPAddress = i
			case 1:
				c.S102IPAddress = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.RFSPIndex:
			switch i.Instance() {
			case 0:
				c.SubscribedRFSPIndex = i
			case 1:
				c.RFSPIndexInUse = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.UETimeZone:
			c.UETimeZone = i
		case ie.LocalDistinguishedName:
			c.MMESGSNLDN = i
		case ie.MDTConfiguration:
			c.MDTConfiguration = i
		case ie.UserCSGInformation:
			c.UCI = i
		case ie.MonitoringEventInformation:
			c.MonitoringEventInformation = i
		case ie.IntegerNumber:
			switch i.Instance() {
			case 0:
				c.UEUsageType = i
			case 1:
				c.RemainingRunningServiceGapTimer = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.SCEFPDNConnection:
			c.SCEFPDNConnection = append(c.SCEFPDNConnection, i)
		case ie.RATType:
			c.RATType = i
		case ie.ServingPLMNRateControl:
			c.ServingPLMNRateControl = i
		case ie.Counter:
			c.MOExceptionDataCounter = i
		case ie.ExtendedTraceInformation:
			c.ExtendedTraceInformation = i
		case ie.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (c *ContextResponse) MarshalLen() int {
	l := c.Header.MarshalLen() - len(c.Header.Payload)

	if ie := c.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.IMSI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UEMMContext; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range c.UEPDNConnections {
		l += ie.MarshalLen()
	}
	if ie := c.SenderFTEID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGWS11S4FTEID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGWNodeName; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.IndicationFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TraceInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.S101IPAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.S102IPAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SubscribedRFSPIndex; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.RFSPIndexInUse; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UETimeZone; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MMESGSNLDN; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MDTConfiguration; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGSNNodeName; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MMENodeName; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UCI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MonitoringEventInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UEUsageType; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range c.SCEFPDNConnection {
		l += ie.MarshalLen()
	}
	if ie := c.RATType; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ServingPLMNRateControl; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MOExceptionDataCounter; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.RemainingRunningServiceGapTimer; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ExtendedTraceInformation; ie != nil {
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
func (c *ContextResponse) SetLength() {
	c.Header.Length = uint16(c.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (c *ContextResponse) MessageTypeName() string {
	return "Context Response"
}

// TEID returns the TEID in uint32.
func (c *ContextResponse) TEID() uint32 {
	return c.Header.teid()
}
