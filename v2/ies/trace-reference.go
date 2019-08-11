// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "github.com/wmnsk/go-gtp/utils"

// NewTraceReference creates a new TraceReference IE.
func NewTraceReference(mcc, mnc string, traceID uint32) *IE {
	i := New(TraceReference, 0x00, make([]byte, 6))
	plmn, err := utils.EncodePLMN(mcc, mnc)
	if err != nil {
		return nil
	}
	copy(i.Payload[0:3], plmn)
	copy(i.Payload[3:6], utils.Uint32To24(traceID))

	return i
}

// TraceID returns TraceID in uint32 if the type of IE matches.
func (i *IE) TraceID() uint32 {
	switch i.Type {
	case TraceReference, TraceInformation:
		if len(i.Payload) < 6 {
			return 0
		}
		return utils.Uint24To32(i.Payload[3:6])
	default:
		return 0
	}
}
