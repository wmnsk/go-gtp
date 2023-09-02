// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"testing"
)

func TestMCCAndMNCFromULI(t *testing.T) {
	tests := []struct {
		name    string
		i       *IE
		wantMCC string
		wantMNC string
		wantErr bool
	}{
		{
			name: "Test getting MCC and MNC from ULI from CGI",
			i: NewUserLocationInformationStruct(
				NewCGI("123", "45", 0x1111, 0x2222), nil, nil, nil, nil, nil, nil, nil,
			),
			wantMCC: "123",
			wantMNC: "45",
		},
		{
			name: "Test getting MCC and MNC from ULI from SAI",
			i: NewUserLocationInformationStruct(
				nil, NewSAI("234", "56", 0x1111, 0x3333), nil, nil, nil, nil, nil, nil,
			),
			wantMCC: "234",
			wantMNC: "56",
		},
		{
			name: "Test getting MCC and MNC from ULI from RAI",
			i: NewUserLocationInformationStruct(
				nil, nil, NewRAI("456", "78", 0x1111, 0x4444), nil, nil, nil, nil, nil,
			),
			wantMCC: "456",
			wantMNC: "78",
		},
		{
			name: "Test getting MCC and MNC from ULI from TAI",
			i: NewUserLocationInformationStruct(
				nil, nil, nil, NewTAI("567", "89", 0x5555), nil, nil, nil, nil,
			),
			wantMCC: "567",
			wantMNC: "89",
		},
		{
			name: "Test getting MCC and MNC from ULI from ECGI",
			i: NewUserLocationInformationStruct(
				nil, nil, nil, nil, NewECGI("678", "90", 0x66666666), nil, nil, nil,
			),
			wantMCC: "678",
			wantMNC: "90",
		},
		{
			name: "Test getting MCC and MNC from ULI from LAI",
			i: NewUserLocationInformationStruct(
				nil, nil, nil, nil, nil, NewLAI("789", "01", 0x1111), nil, nil,
			),
			wantMCC: "789",
			wantMNC: "01",
		},
		{
			name: "Test getting MCC and MNC from ULI from MENBI",
			i: NewUserLocationInformationStruct(
				nil, nil, nil, nil, nil, nil, NewMENBI("890", "12", 0x11111111), nil,
			),
			wantMCC: "890",
			wantMNC: "12",
		},
		{
			name: "Test getting MCC and MNC from ULI from EMENBI",
			i: NewUserLocationInformationStruct(
				nil, nil, nil, nil, nil, nil, nil, NewEMENBI("321", "54", 0x22222222),
			),
			wantMCC: "321",
			wantMNC: "54",
		},
		{
			name: "Test getting empty strings when uli does not contain information elemtents",
			i: NewUserLocationInformationStruct(
				nil, nil, nil, nil, nil, nil, nil, nil,
			),
			wantMCC: "",
			wantMNC: "",
			wantErr: true,
		},
		{
			name: "Test getting MCC and MNC from ULI from CGI when all elements are given",
			i: NewUserLocationInformationStruct(
				NewCGI("123", "45", 0x1111, 0x2222),
				NewSAI("234", "56", 0x1111, 0x3333),
				NewRAI("456", "78", 0x1111, 0x4444),
				NewTAI("567", "89", 0x5555),
				NewECGI("678", "90", 0x66666666),
				NewLAI("789", "01", 0x1111),
				NewMENBI("890", "12", 0x11111111),
				NewEMENBI("321", "54", 0x22222222),
			),
			wantMCC: "123",
			wantMNC: "45",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.MCC()
			if (err != nil) != tt.wantErr {
				t.Errorf("IE.MCC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantMCC {
				t.Errorf("IE.MCC() = %v, want %v", got, tt.wantMCC)
			}
			got, err = tt.i.MNC()
			if (err != nil) != tt.wantErr {
				t.Errorf("IE.MNC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantMNC {
				t.Errorf("IE.MNC() = %v, want %v", got, tt.wantMNC)
			}
		})
	}
}
