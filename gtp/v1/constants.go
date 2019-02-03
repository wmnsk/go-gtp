// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v1

// SelectionMode definitions.
const (
	SelectionModeMSorNetworkProvidedAPNSubscribedVerified uint8 = iota | 0xf0
	SelectionModeMSProvidedAPNSubscriptionNotVerified
	SelectionModeNetworkProvidedAPNSubscriptionNotVerified
)

// PDP Type Organization definitions.
const (
	PDPTypeETSI uint8 = iota | 0xf0
	PDPTypeIETF
)

// Protocol ID definitions.
// For more identifiers, see RFC 3232.
const (
	ProtoIDLCP  uint16 = 0xc021
	ProtoIDPAP  uint16 = 0xc023
	ProtoIDCHAP uint16 = 0xc223
	ProtoIDIPCP uint16 = 0x8021
)

// Container ID definitions.
const (
	_ uint16 = iota
	ContIDPCSCFIPv6AddressRequest
	ContIDIMCNSubsystemSignalingFlag
	ContIDDNSServerIPv6AddressRequest
	ContIDNotSupported
	ContIDMSSupportofNetworkRequestedBearerControlIndicator
	_
	ContIDDSMIPv6HomeAgentAddressRequest
	ContIDDSMIPv6HomeNetworkPrefixRequest
	ContIDDSMIPv6IPv4HomeAgentAddressRequest
	ContIDIPaddressAllocationViaNASSignalling
	ContIDIPv4addressAllocationViaDHCPv4
	ContIDPCSCFIPv4AddressRequest
	ContIDDNSServerIPv4AddressRequest
	ContIDMSISDNRequest
	ContIDIFOMSupportRequest
	ContIDIPv4LinkMTURequest
	ContIDMSSupportOfLocalAddressInTFTIndicator
	ContIDPCSCFReselectionSupport
	ContIDNBIFOMRequestIndicator
	ContIDNBIFOMMode
	ContIDNonIPLinkMTURequest
	ContIDAPNRateControlSupportIndicator
	ContID3GPPPSDataOffUEstatus
	ContIDReliableDataServiceRequestIndicator
	ContIDAdditionalAPNRateControlForExceptionDataSupportIndicator
	ContIDPDUSessionID
	_
	_
	_
	_
	_
	ContIDEthernetFramePayloadMTURequest
	ContIDUnstructuredLinkMTURequest
	ContID5GSMCauseValue
)

// Configuration Protocol definitions.
const (
	ConfigProtocolPPPWithIP uint8 = 0
)

// RATType definitions.
const (
	_ uint8 = iota
	RatTypeUTRAN
	RatTypeGERAN
	RatTypeWLAN
	RatTypeGAN
	RatTypeHSPAEvolution
	RatTypeEUTRAN
)

// UserLocationInformation GeographicLocationType definitions.
const (
	LocTypeCGI uint8 = iota
	LocTypeSAI
	LocTypeRAI
)
