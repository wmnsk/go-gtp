// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewCommonFlags creates a new CommonFlags IE.
//
// Note: each flag should be set in 1 or 0.
func NewCommonFlags(dualAddr, upgradeQoS, nrsn, noQoS, mbmsCount, ranReady, mbmsService, prohibitComp int) *IE {
	return New(
		CommonFlags,
		[]byte{uint8(
			dualAddr<<7 | upgradeQoS<<6 | nrsn<<5 | noQoS<<4 | mbmsCount<<3 | ranReady<<2 | mbmsService<<1 | prohibitComp,
		)},
	)
}

// CommonFlags returns CommonFlags value if type matches.
func (i *IE) CommonFlags() (uint8, error) {
	if i.Type != CommonFlags {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustCommonFlags returns CommonFlags in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustCommonFlags() uint8 {
	v, _ := i.CommonFlags()
	return v
}

// IsDualAddressBearer checks if DualAddressBearer flag exists in CommonFlags.
func (i *IE) IsDualAddressBearer() bool {
	return ((i.MustCommonFlags() >> 7) & 0x01) != 0
}

// IsUpgradeQoSSupported checks if UpgradeQoSSupported flag exists in CommonFlags.
func (i *IE) IsUpgradeQoSSupported() bool {
	return ((i.MustCommonFlags() >> 6) & 0x01) != 0
}

// IsNRSN checks if NRSN flag exists in CommonFlags.
func (i *IE) IsNRSN() bool {
	return ((i.MustCommonFlags() >> 5) & 0x01) != 0
}

// IsNoQoSNegotiation checks if NoQoSNegotiation flag exists in CommonFlags.
func (i *IE) IsNoQoSNegotiation() bool {
	return ((i.MustCommonFlags() >> 4) & 0x01) != 0
}

// IsMBMSCountingInformation checks if MBMSCountingInformation flag exists in CommonFlags.
func (i *IE) IsMBMSCountingInformation() bool {
	return ((i.MustCommonFlags() >> 3) & 0x01) != 0
}

// IsRANProceduresReady checks if RANProceduresReady flag exists in CommonFlags.
func (i *IE) IsRANProceduresReady() bool {
	return ((i.MustCommonFlags() >> 2) & 0x01) != 0
}

// IsMBMSServiceType checks if MBMSServiceType flag exists in CommonFlags.
func (i *IE) IsMBMSServiceType() bool {
	return ((i.MustCommonFlags() >> 1) & 0x01) != 0
}

// IsProhibitPayloadCompression checks if ProhibitPayloadCompression flag exists in CommonFlags.
func (i *IE) IsProhibitPayloadCompression() bool {
	return (i.MustCommonFlags() & 0x01) != 0
}
