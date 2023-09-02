// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

func FuzzParse(f *testing.F) {
	for _, c := range cases {
		f.Add(c.serialized)
	}

	testIE := func(t *testing.T, v *ie.IE) {
		// Generated with
		// grep -ho -e "Must.*()" * | sort | uniq | sed -e 's/^/v./'
		v.MustAccessMode()
		v.MustAccessPointName()
		v.MustAggregateMaximumBitRateDown()
		v.MustAggregateMaximumBitRateUp()
		v.MustAPNRestriction()
		v.MustBearerFlags()
		v.MustCause()
		v.MustCauseFlags()
		v.MustChargingCharacteristics()
		v.MustChargingID()
		v.MustCMI()
		v.MustCNID()
		v.MustCSGID()
		v.MustCSIDs()
		v.MustDaylightSaving()
		v.MustDelayValue()
		v.MustDetachType()
		v.MustEBIs()
		v.MustEnterpriseID()
		v.MustEPCTimer()
		v.MustEPSBearerID()
		v.MustFullyQualifiedDomainName()
		v.MustGBRForDownlink()
		v.MustGBRForUplink()
		v.MustGREKey()
		v.MustHopCounter()
		v.MustHSGWAddress()
		v.MustIMSI()
		v.MustIntegerNumber()
		v.MustInterfaceType()
		v.MustIP()
		v.MustIPAddress()
		v.MustIPv4()
		v.MustIPv6()
		v.MustLocalDistinguishedName()
		v.MustMBMSFlags()
		v.MustMBRForDownlink()
		v.MustMBRForUplink()
		v.MustMCC()
		v.MustMMECode()
		v.MustMMEGroupID()
		v.MustMNC()
		v.MustMobileEquipmentIdentity()
		v.MustMSISDN()
		v.MustMTMSI()
		v.MustNodeFeatures()
		v.MustNodeID()
		v.MustNodeIDType()
		v.MustNodeType()
		v.MustOffendingIE()
		v.MustPacketTMSI()
		v.MustPagingPolicyIndication()
		v.MustPDNType()
		v.MustPLMNID()
		v.MustPortNumber()
		v.MustPrivateExtension()
		v.MustProcedureTransactionID()
		v.MustProtocolConfigurationOptions()
		v.MustPTMSISignature()
		v.MustRATType()
		v.MustRecovery()
		v.MustRFSPIndex()
		v.MustSelectionMode()
		v.MustServiceIndicator()
		v.MustServingNetwork()
		v.MustSGWAddress()
		v.MustTEID()
		v.MustTimestamp()
		v.MustTimeZone()
		v.MustTMSI()
		v.MustTraceID()
	}

	f.Fuzz(func(t *testing.T, data []byte) {
		if v, err := ie.Parse(data); err == nil && v == nil {
			t.Errorf("nil without error")
		} else if err == nil {
			testIE(t, v)
			for _, cie := range v.ChildIEs {
				testIE(t, cie)
			}
		}
	})
}
