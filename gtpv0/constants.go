// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtpv0

// Cause definitions.
const (
	CauseRequestIMSI              uint8 = 0
	CauseRequestIMEI              uint8 = 1
	CauseRequestIMSIandIMEI       uint8 = 2
	CauseNoIdentityNeeded         uint8 = 3
	CauseRequestAccepted          uint8 = 128
	CauseNonExistent              uint8 = 192
	CauseInvalidMessageFormat     uint8 = 193
	CauseIMSINotKnown             uint8 = 194
	CauseMSIsGPRSDetached         uint8 = 195
	CauseMSIsNotGPRSResponding    uint8 = 196
	CauseMSRefuses                uint8 = 197
	CauseVersionNotSupported      uint8 = 198
	CauseNoResourcesAvailable     uint8 = 199
	CauseServiceNotSupported      uint8 = 200
	CauseMandatoryIEIncorrect     uint8 = 201
	CauseMandatoryIEMissing       uint8 = 202
	CauseOptionalIEIncorrect      uint8 = 203
	CauseSystemFailure            uint8 = 204
	CauseRoamingRestriction       uint8 = 205
	CausePTMSISignatureMismatch   uint8 = 206
	CauseGPRSConnectionSuspended  uint8 = 207
	CauseAuthenticationFailure    uint8 = 208
	CauseUserAuthenticationFailed uint8 = 209
)

// PDP Type Organization definitions.
const (
	PDPTypeETSI uint8 = iota | 0xf0
	PDPTypeIETF
)

// SelectionMode definitions.
const (
	SelectionModeMSorNetworkProvidedAPNSubscribedVerified uint8 = iota | 0xf0
	SelectionModeMSProvidedAPNSubscriptionNotVerified
	SelectionModeNetworkProvidedAPNSubscriptionNotVerified
)
