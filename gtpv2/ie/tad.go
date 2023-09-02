// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewTrafficAggregateDescription creates a new TrafficAggregateDescription IE.
//
// Custom constructors for each operation code are available, which does not require
// unnecessary parameters.
func NewTrafficAggregateDescription(op uint8, filters []*TFTPacketFilter, ids []uint8, params []*TFTParameter) *IE {
	v := NewTrafficFlowTemplate(op, filters, ids, params)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(TrafficAggregateDescription, 0x00, b)
}

// NewTrafficAggregateDescriptionCreateNewTFT creates a new TrafficAggregateDescription IE with opcode=CreateNewTFT.
func NewTrafficAggregateDescriptionCreateNewTFT(filters []*TFTPacketFilter, params []*TFTParameter) *IE {
	return NewTrafficAggregateDescription(TFTOpCreateNewTFT, filters, nil, params)
}

// NewTrafficAggregateDescriptionAddPacketFilters creates a new TrafficAggregateDescription IE with opcode=AddPacketFiltersToExistingTFT.
func NewTrafficAggregateDescriptionAddPacketFilters(filters []*TFTPacketFilter, params []*TFTParameter) *IE {
	return NewTrafficAggregateDescription(TFTOpAddPacketFiltersToExistingTFT, filters, nil, params)
}

// NewTrafficAggregateDescriptionReplacePacketFilters creates a new TrafficAggregateDescription IE with opcode=ReplacePacketFiltersInExistingTFT.
func NewTrafficAggregateDescriptionReplacePacketFilters(filters []*TFTPacketFilter, params []*TFTParameter) *IE {
	return NewTrafficAggregateDescription(TFTOpReplacePacketFiltersInExistingTFT, filters, nil, params)
}

// NewTrafficAggregateDescriptionDeletePacketFilters creates a new TrafficAggregateDescription IE with opcode=DeletePacketFiltersFromExistingTFT.
func NewTrafficAggregateDescriptionDeletePacketFilters(ids []uint8, params ...*TFTParameter) *IE {
	return NewTrafficAggregateDescription(TFTOpDeletePacketFiltersFromExistingTFT, nil, ids, params)
}

// NewTrafficAggregateDescriptionDeleteExistingTFT creates a new TrafficAggregateDescription IE with opcode=DeleteExistingTFT.
func NewTrafficAggregateDescriptionDeleteExistingTFT(params ...*TFTParameter) *IE {
	return NewTrafficAggregateDescription(TFTOpDeleteExistingTFT, nil, nil, params)
}

// NewTrafficAggregateDescriptionNoTFTOperation creates a new TrafficAggregateDescription IE with opcode=NoTFTOperation.
func NewTrafficAggregateDescriptionNoTFTOperation(params ...*TFTParameter) *IE {
	return NewTrafficAggregateDescription(TFTOpNoTFTOperation, nil, nil, params)
}

// TrafficAggregateDescription returns TrafficAggregateDescription in TrafficFlowTemplate type if the type of IE matches.
func (i *IE) TrafficAggregateDescription() (*TrafficFlowTemplate, error) {
	return i.TrafficFlowTemplate()
}
