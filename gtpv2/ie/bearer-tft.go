// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// TFT Operation Code definitions.
const (
	TFTOpIgnoreThisIE                       uint8 = 0
	TFTOpCreateNewTFT                       uint8 = 1
	TFTOpDeleteExistingTFT                  uint8 = 2
	TFTOpAddPacketFiltersToExistingTFT      uint8 = 3
	TFTOpReplacePacketFiltersInExistingTFT  uint8 = 4
	TFTOpDeletePacketFiltersFromExistingTFT uint8 = 5
	TFTOpNoTFTOperation                     uint8 = 6
)

// NewBearerTFT creates a new BearerTFT IE.
//
// Custom constructors for each operation code are available, which does not require
// unnecessary parameters.
func NewBearerTFT(op uint8, filters []*TFTPacketFilter, ids []uint8, params []*TFTParameter) *IE {
	v := NewTrafficFlowTemplate(op, filters, ids, params)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(BearerTFT, 0x00, b)
}

// NewBearerTFTCreateNewTFT creates a new BearerTFT IE with opcode=CreateNewTFT.
func NewBearerTFTCreateNewTFT(filters []*TFTPacketFilter, params []*TFTParameter) *IE {
	return NewBearerTFT(TFTOpCreateNewTFT, filters, nil, params)
}

// NewBearerTFTAddPacketFilters creates a new BearerTFT IE with opcode=AddPacketFiltersToExistingTFT.
func NewBearerTFTAddPacketFilters(filters []*TFTPacketFilter, params []*TFTParameter) *IE {
	return NewBearerTFT(TFTOpAddPacketFiltersToExistingTFT, filters, nil, params)
}

// NewBearerTFTReplacePacketFilters creates a new BearerTFT IE with opcode=ReplacePacketFiltersInExistingTFT.
func NewBearerTFTReplacePacketFilters(filters []*TFTPacketFilter, params []*TFTParameter) *IE {
	return NewBearerTFT(TFTOpReplacePacketFiltersInExistingTFT, filters, nil, params)
}

// NewBearerTFTDeletePacketFilters creates a new BearerTFT IE with opcode=DeletePacketFiltersFromExistingTFT.
func NewBearerTFTDeletePacketFilters(ids []uint8, params ...*TFTParameter) *IE {
	return NewBearerTFT(TFTOpDeletePacketFiltersFromExistingTFT, nil, ids, params)
}

// NewBearerTFTDeleteExistingTFT creates a new BearerTFT IE with opcode=DeleteExistingTFT.
func NewBearerTFTDeleteExistingTFT(params ...*TFTParameter) *IE {
	return NewBearerTFT(TFTOpDeleteExistingTFT, nil, nil, params)
}

// NewBearerTFTNoTFTOperation creates a new BearerTFT IE with opcode=NoTFTOperation.
func NewBearerTFTNoTFTOperation(params ...*TFTParameter) *IE {
	return NewBearerTFT(TFTOpNoTFTOperation, nil, nil, params)
}

// BearerTFT returns TrafficFlowTemplate struct if the type of IE matches.
func (i *IE) BearerTFT() (*TrafficFlowTemplate, error) {
	return i.TrafficFlowTemplate()
}
