// Copyright 2019-2021 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"testing"
)

func TestUserLocationInformationWithCGI(t *testing.T) {
	t.Run("Test UserLocationInformation with CGI", func(t *testing.T) {
		ie := NewUserLocationInformationWithCGI("123", "45", 111, 222)

		cgi := ie.MustCGI()
		if cgi != 222 {
			t.Errorf("wrong cgi, got %v", cgi)
		}

		lac := ie.MustLAC()
		if lac != 111 {
			t.Errorf("wrong lac, got %v", lac)
		}

		mcc := ie.MustMCC()
		if mcc != "123" {
			t.Errorf("wrong mcc, got %v", mcc)
		}

		mnc := ie.MustMNC()
		if mnc != "45" {
			t.Errorf("wrong mnc, got %v", mnc)
		}
	})
}

func TestUserLocationInformationWithRAI(t *testing.T) {
	t.Run("Test UserLocationInformation with RAI", func(t *testing.T) {
		ie := NewUserLocationInformationWithRAI("123", "45", 111, 222)

		rac := ie.MustRAC()
		if rac != 222 {
			t.Errorf("wrong rac, got %v", rac)
		}

		lac := ie.MustLAC()
		if lac != 111 {
			t.Errorf("wrong lac, got %v", lac)
		}

		mcc := ie.MustMCC()
		if mcc != "123" {
			t.Errorf("wrong mcc, got %v", mcc)
		}

		mnc := ie.MustMNC()
		if mnc != "45" {
			t.Errorf("wrong mnc, got %v", mnc)
		}
	})
}

func TestUserLocationInformationWithSAI(t *testing.T) {
	t.Run("Test UserLocationInformation with SAI", func(t *testing.T) {
		ie := NewUserLocationInformationWithSAI("123", "45", 111, 222)

		sac := ie.MustSAC()
		if sac != 222 {
			t.Errorf("wrong sac, got %v", sac)
		}

		lac := ie.MustLAC()
		if lac != 111 {
			t.Errorf("wrong lac, got %v", lac)
		}

		mcc := ie.MustMCC()
		if mcc != "123" {
			t.Errorf("wrong mcc, got %v", mcc)
		}

		mnc := ie.MustMNC()
		if mnc != "45" {
			t.Errorf("wrong mnc, got %v", mnc)
		}
	})
}
