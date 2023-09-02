// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"testing"
)

func TestNewUserLocationInformationStruct(t *testing.T) {
	t.Run("Test UserLocationInformation unmarshal", func(t *testing.T) {
		uli := NewUserLocationInformationStruct(
			NewCGI("123", "45", 0x1111, 0x2222),
			NewSAI("123", "45", 0x1111, 0x3333),
			NewRAI("123", "45", 0x1111, 0x4444),
			NewTAI("123", "45", 0x5555),
			NewECGI("123", "45", 0x66666666),
			NewLAI("123", "45", 0x1111),
			NewMENBI("123", "45", 0x11111111),
			NewEMENBI("123", "45", 0x22222222),
		)
		if uli == nil {
			t.Fatalf("Error in NewUserLocationInformationStruct")
		}

		uliFields, err := uli.UserLocationInformation()
		if err != nil {
			t.Fatalf("Error in unmarshal: %v", err)
		}

		// CGI
		if uliFields.CGI.LAC != 0x1111 {
			t.Errorf("wrong uliFields.CGI.LAC, got: 0x%x", uliFields.CGI.LAC)
		}
		if uliFields.CGI.CI != 0x2222 {
			t.Errorf("wrong uliFields.CGI.CI, got: 0x%x", uliFields.CGI.CI)
		}
		if uliFields.CGI.PLMN.MCC != "123" {
			t.Errorf("wrong uliFields.CGI.PLMN.MCC, got: %v", uliFields.CGI.PLMN.MCC)
		}
		if uliFields.CGI.PLMN.MNC != "45" {
			t.Errorf("wrong uliFields.CGI.PLMN.MNC, got: %v", uliFields.CGI.PLMN.MNC)
		}

		// SAI
		if uliFields.SAI.LAC != 0x1111 {
			t.Errorf("wrong uliFields.SAI.LAC, got: 0x%x", uliFields.SAI.LAC)
		}
		if uliFields.SAI.SAC != 0x3333 {
			t.Errorf("wrong uliFields.SAI.SAC, got: 0x%x", uliFields.SAI.SAC)
		}
		if uliFields.SAI.PLMN.MCC != "123" {
			t.Errorf("wrong uliFields.SAI.PLMN.MCC, got: %v", uliFields.SAI.PLMN.MCC)
		}
		if uliFields.SAI.PLMN.MNC != "45" {
			t.Errorf("wrong uliFields.SAI.PLMN.MNC, got: %v", uliFields.SAI.PLMN.MNC)
		}

		// RAI
		if uliFields.RAI.LAC != 0x1111 {
			t.Errorf("wrong uliFields.RAI.LAC, got: 0x%x", uliFields.RAI.LAC)
		}
		if uliFields.RAI.RAC != 0x4444 {
			t.Errorf("wrong uliFields.RAI.RAC, got: 0x%x", uliFields.RAI.RAC)
		}
		if uliFields.RAI.PLMN.MCC != "123" {
			t.Errorf("wrong uliFields.RAI.PLMN.MCC, got: %v", uliFields.RAI.PLMN.MCC)
		}
		if uliFields.RAI.PLMN.MNC != "45" {
			t.Errorf("wrong uliFields.RAI.PLMN.MNC, got: %v", uliFields.RAI.PLMN.MNC)
		}

		// TAI
		if uliFields.TAI.TAC != 0x5555 {
			t.Errorf("wrong uliFields.TAI.TAC, got: 0x%x", uliFields.TAI.TAC)
		}
		if uliFields.TAI.PLMN.MCC != "123" {
			t.Errorf("wrong uliFields.TAI.PLMN.MCC, got: %v", uliFields.TAI.PLMN.MCC)
		}
		if uliFields.TAI.PLMN.MNC != "45" {
			t.Errorf("wrong uliFields.TAI.PLMN.MNC, got: %v", uliFields.TAI.PLMN.MNC)
		}

		// ECGI
		if uliFields.ECGI.ECI != 0x6666666 {
			t.Errorf("wrong uliFields.ECGI.ECI, got: 0x%x", uliFields.ECGI.ECI)
		}
		if uliFields.ECGI.PLMN.MCC != "123" {
			t.Errorf("wrong uliFields.ECGI.PLMN.MCC, got: %v", uliFields.ECGI.PLMN.MCC)
		}
		if uliFields.ECGI.PLMN.MNC != "45" {
			t.Errorf("wrong uliFields.ECGI.PLMN.MNC, got: %v", uliFields.ECGI.PLMN.MNC)
		}

		// LAI
		if uliFields.LAI.LAC != 0x1111 {
			t.Errorf("wrong uliFields.LAI.LAC, got: 0x%x", uliFields.LAI.LAC)
		}
		if uliFields.LAI.PLMN.MCC != "123" {
			t.Errorf("wrong uliFields.LAI.PLMN.MCC, got: %v", uliFields.LAI.PLMN.MCC)
		}
		if uliFields.LAI.PLMN.MNC != "45" {
			t.Errorf("wrong uliFields.LAI.PLMN.MNC, got: %v", uliFields.LAI.PLMN.MNC)
		}

		// MENBI
		if uliFields.MENBI.MENBI != 0x111111 {
			t.Errorf("wrong uliFields.MENBI.MENBI, got: 0x%x", uliFields.MENBI.MENBI)
		}
		if uliFields.MENBI.PLMN.MCC != "123" {
			t.Errorf("wrong uliFields.MENBI.PLMN.MCC, got: %v", uliFields.MENBI.PLMN.MCC)
		}
		if uliFields.MENBI.PLMN.MNC != "45" {
			t.Errorf("wrong uliFields.MENBI.PLMN.MNC, got: %v", uliFields.MENBI.PLMN.MNC)
		}

		// EMENBI
		if uliFields.EMENBI.EMENBI != 0x222222 {
			t.Errorf("wrong uliFields.EMENBI.EMENBI, got: 0x%x", uliFields.EMENBI.EMENBI)
		}
		if uliFields.EMENBI.PLMN.MCC != "123" {
			t.Errorf("wrong uliFields.EMENBI.PLMN.MCC, got: %v", uliFields.EMENBI.PLMN.MCC)
		}
		if uliFields.EMENBI.PLMN.MNC != "45" {
			t.Errorf("wrong uliFields.EMENBI.PLMN.MNC, got: %v", uliFields.EMENBI.PLMN.MNC)
		}
	})
}
