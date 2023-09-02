// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtpv2

// Fixes for the constants with wrong names (original ones are kept for compatibility).
const (
	ContIDMSSupportOfNetworkRequestedBearerControlIndicator uint16 = 5  // ContIDMSSupportofNetworkRequestedBearerControlIndicator
	ContIDIPAddressAllocationViaNASSignalling               uint16 = 10 // ContIDIPaddressAllocationViaNASSignalling
	ContIDIPv4AddressAllocationViaDHCPv4                    uint16 = 11 // ContIDIPv4addressAllocationViaDHCPv4
	ContID3GPPPSDataOffUEStatus                             uint16 = 23 // ContID3GPPPSDataOffUEstatus

	SelectionModeMSOrNetworkProvidedAPNSubscribedVerified uint8 = 0 // SelectionModeMSorNetworkProvidedAPNSubscribedVerified
)

// Registered UDP ports
const (
	GTPCPort = ":2123"
	GTPUPort = ":2152"
)

// InterfaceType definitions.
const (
	IFTypeS1UeNodeBGTPU uint8 = iota
	IFTypeS1USGWGTPU
	IFTypeS12RNCGTPU
	IFTypeS12SGWGTPU
	IFTypeS5S8SGWGTPU
	IFTypeS5S8PGWGTPU
	IFTypeS5S8SGWGTPC
	IFTypeS5S8PGWGTPC
	IFTypeS5S8SGWPMIPv6
	IFTypeS5S8PGWPMIPv6
	IFTypeS11MMEGTPC
	IFTypeS11S4SGWGTPC
	IFTypeS10MMEGTPC
	IFTypeS3MMEGTPC
	IFTypeS3SGSNGTPC
	IFTypeS4SGSNGTPU
	IFTypeS4SGWGTPU
	IFTypeS4SGSNGTPC
	IFTypeS16SGSNGTPC
	IFTypeeNodeBGTPUForDL
	IFTypeeNodeBGTPUForUL
	IFTypeRNCGTPUForData
	IFTypeSGSNGTPUForData
	IFTypeSGWUPFGTPUForDL
	IFTypeSmMBMSGWGTPC
	IFTypeSnMBMSGWGTPC
	IFTypeSmMMEGTPC
	IFTypeSnSGSNGTPC
	IFTypeSGWGTPUForUL
	IFTypeSnSGSNGTPU
	IFTypeS2bePDGGTPC
	IFTypeS2bUePDGGTPU
	IFTypeS2bPGWGTPC
	IFTypeS2bUPGWGTPU
	IFTypeS2aTWANGTPU
	IFTypeS2aTWANGTPC
	IFTypeS2aPGWGTPC
	IFTypeS2aPGWGTPU
	IFTypeS11MMEGTPU
	IFTypeS11SGWGTPU
)

// APN Restriction definitions.
const (
	APNRestrictionNoExistingContextsorRestriction uint8 = iota
	APNRestrictionPublic1
	APNRestrictionPublic2
	APNRestrictionPrivate1
	APNRestrictionPrivate2
)

// Cause definitions.
const (
	_                                                                                   uint8 = 0
	_                                                                                   uint8 = 1
	CauseLocalDetach                                                                    uint8 = 2
	CauseCompleteDetach                                                                 uint8 = 3
	CauseRATChangedFrom3GPPToNon3GPP                                                    uint8 = 4
	CauseISRDeactivation                                                                uint8 = 5
	CauseErrorIndicationReceivedFromRNCeNodeBS4SGSNMME                                  uint8 = 6
	CauseIMSIDetachOnly                                                                 uint8 = 7
	CauseReactivationRequested                                                          uint8 = 8
	CausePDNReconnectionToThisAPNDisallowed                                             uint8 = 9
	CauseAccessChangedFromNon3GPPTo3GPP                                                 uint8 = 10
	CausePDNConnectionInactivityTimerExpires                                            uint8 = 11
	CausePGWNotResponding                                                               uint8 = 12
	CauseNetworkFailure                                                                 uint8 = 13
	CauseQoSParameterMismatch                                                           uint8 = 14
	_                                                                                   uint8 = 15
	CauseRequestAccepted                                                                uint8 = 16
	CauseRequestAcceptedPartially                                                       uint8 = 17
	CauseNewPDNTypeDueToNetworkPreference                                               uint8 = 18
	CauseNewPDNTypeDueToSingleAddressBearerOnly                                         uint8 = 19
	CauseContextNotFound                                                                uint8 = 64
	CauseInvalidMessageFormat                                                           uint8 = 65
	CauseVersionNotSupportedByNextPeer                                                  uint8 = 66
	CauseInvalidLength                                                                  uint8 = 67
	CauseServiceNotSupported                                                            uint8 = 68
	CauseMandatoryIEIncorrect                                                           uint8 = 69
	CauseMandatoryIEMissing                                                             uint8 = 70
	_                                                                                   uint8 = 71
	CauseSystemFailure                                                                  uint8 = 72
	CauseNoResourcesAvailable                                                           uint8 = 73
	CauseSemanticErrorInTheTFTOperation                                                 uint8 = 74
	CauseSyntacticErrorInTheTFTOperation                                                uint8 = 75
	CauseSemanticErrorsInPacketFilters                                                  uint8 = 76
	CauseSyntacticErrorsInPacketFilters                                                 uint8 = 77
	CauseMissingOrUnknownAPN                                                            uint8 = 78
	_                                                                                   uint8 = 79
	CauseGREKeyNotFound                                                                 uint8 = 80
	CauseRelocationFailure                                                              uint8 = 81
	CauseDeniedInRAT                                                                    uint8 = 82
	CausePreferredPDNTypeNotSupported                                                   uint8 = 83
	CauseAllDynamicAddressesAreOccupied                                                 uint8 = 84
	CauseUEContextWithoutTFTAlreadyActivated                                            uint8 = 85
	CauseProtocolTypeNotSupported                                                       uint8 = 86
	CauseUENotResponding                                                                uint8 = 87
	CauseUERefuses                                                                      uint8 = 88
	CauseServiceDenied                                                                  uint8 = 89
	CauseUnableToPageUE                                                                 uint8 = 90
	CauseNoMemoryAvailable                                                              uint8 = 91
	CauseUserAuthenticationFailed                                                       uint8 = 92
	CauseAPNAccessDeniedNoSubscription                                                  uint8 = 93
	CauseRequestRejectedReasonNotSpecified                                              uint8 = 94
	CausePTMSISignatureMismatch                                                         uint8 = 95
	CauseIMSIIMEINotKnown                                                               uint8 = 96
	CauseSemanticErrorInTheTADOperation                                                 uint8 = 97
	CauseSyntacticErrorInTheTADOperation                                                uint8 = 98
	_                                                                                   uint8 = 99
	CauseRemotePeerNotResponding                                                        uint8 = 100
	CauseCollisionWithNetworkInitiatedRequest                                           uint8 = 101
	CauseUnableToPageUEDueToSuspension                                                  uint8 = 102
	CauseConditionalIEMissing                                                           uint8 = 103
	CauseAPNRestrictionTypeIncompatibleWithCurrentlyActivePDNConnection                 uint8 = 104
	CauseInvalidOverallLengthOfTheTriggeredResponseMessageAndAPiggybackedInitialMessage uint8 = 105
	CauseDataForwardingNotSupported                                                     uint8 = 106
	CauseInvalidReplyFromRemotePeer                                                     uint8 = 107
	CauseFallbackToGTPv1                                                                uint8 = 108
	CauseInvalidPeer                                                                    uint8 = 109
	CauseTemporarilyRejectedDueToHandoverTAURAUProcedureInProgress                      uint8 = 110
	CauseModificationsNotLimitedToS1UBearers                                            uint8 = 111
	CauseRequestRejectedForAPMIPv6Reason                                                uint8 = 112
	CauseAPNCongestion                                                                  uint8 = 113
	CauseBearerHandlingNotSupported                                                     uint8 = 114
	CauseUEAlreadyReattached                                                            uint8 = 115
	CauseMultiplePDNConnectionsForAGivenAPNNotAllowed                                   uint8 = 116
	CauseTargetAccessRestrictedForTheSubscriber                                         uint8 = 117
	_                                                                                   uint8 = 118
	CauseMMESGSNRefusesDueToVPLMNPolicy                                                 uint8 = 119
	CauseGTPCEntityCongestion                                                           uint8 = 120
	CauseLateOverlappingRequest                                                         uint8 = 121
	CauseTimedOutRequest                                                                uint8 = 122
	CauseUEIsTemporarilyNotReachableDueToPowerSaving                                    uint8 = 123
	CauseRelocationFailureDueToNASMessageRedirection                                    uint8 = 124
	CauseUENotAuthorisedByOCSOrExternalAAAServer                                        uint8 = 125
	CauseMultipleAccessesToAPDNConnectionNotAllowed                                     uint8 = 126
	CauseRequestRejectedDueToUECapability                                               uint8 = 127
	CauseS1UPathFailure                                                                 uint8 = 128
)

// CSG Membership Indication definitions.
const (
	CMINonCSG uint8 = iota
	CMICSG
)

// Detach Type definitions.
const (
	_ uint8 = iota
	DetachTypePS
	DetachTypeCombinedPSCS
)

// Node-ID Type definitions.
const (
	NodeIDIPv4 uint8 = iota
	NodeIDIPv6
	NodeIDOther
)

// Node Type definitions.
const (
	NodeTypeSGSN uint8 = iota
	NodeTypeMME
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

// PDN Type definitions.
const (
	_ uint8 = iota
	PDNTypeIPv4
	PDNTypeIPv6
	PDNTypeIPv4v6
	PDNTypeNonIP
)

// Protocol Type definitions.
const (
	_ uint8 = iota
	ProtoTypeS1APCause
	ProtoTypeEMMCause
	ProtoTypeESMCause
	ProtoTypeDiameterCause
	ProtoTypeIKEv2Cause
)

// Cause Type definitions.
const (
	CauseTypeRadioNetworkLayer uint8 = iota
	CauseTypeTransportLayer
	CauseTypeNAS
	CauseTypeProtocol
	CauseTypeMiscellaneous
)

// RAT Type definitions.
const (
	_ uint8 = iota
	RATTypeUTRAN
	RATTypeGERAN
	RATTypeWLAN
	RATTypeGAN
	RATTypeHSPAEvolution
	RATTypeEUTRAN
	RATTypeVirtual
	RATTypeEUTRANNBIoT
	RATTypeLTEM
	RATTypeNR
)

// SelectionMode definitions.
const (
	SelectionModeMSorNetworkProvidedAPNSubscribedVerified uint8 = iota
	SelectionModeMSProvidedAPNSubscriptionNotVerified
	SelectionModeNetworkProvidedAPNSubscriptionNotVerified
)

// Service Indicator definitions.
const (
	_ uint8 = iota
	ServiceIndCSCall
	ServiceIndSMS
)

// Access Mode definitions.
const (
	AccessModeClosed uint8 = iota
	AccessModeHybrid
)

// Daylight Saving Time definitions.
const (
	DaylightSavingNoAdjustment uint8 = iota
	DaylightSavingPlusOneHour
	DaylightSavingPlusTwoHours
)
