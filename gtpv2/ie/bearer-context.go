// Copyright 2019-2024 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewBearerContext creates a new BearerContext IE.
func NewBearerContext(ies ...*IE) *IE {
	return NewGroupedIE(BearerContext, ies...)
}

// NewBearerContextWithinCreateBearerRequest creates a new BearerContext used within CreateBearerRequest.
func NewBearerContextWithinCreateBearerRequest(ebi, tft, qos, chargeID, flags, pco, epco, mplr *IE, fTEIDs ...*IE) *IE {
	n := 8 + len(fTEIDs)
	ies := make([]*IE, n)

	ies[0] = ebi
	ies[1] = tft
	ies[2] = qos
	ies[3] = chargeID
	ies[4] = flags
	ies[5] = pco
	ies[6] = epco
	ies[7] = mplr
	copy(ies[8:], fTEIDs)

	return NewBearerContext(ies...)
}

// NewBearerContextWithinCreateBearerResponse creates a new BearerContext used within CreateBearerResponse.
func NewBearerContextWithinCreateBearerResponse(ebi, cause, pco, rannasCause, epco *IE, fTEIDs ...*IE) *IE {
	n := 5 + len(fTEIDs)
	ies := make([]*IE, n)

	ies[0] = ebi
	ies[1] = cause
	ies[2] = pco
	ies[3] = rannasCause
	ies[4] = epco
	copy(ies[5:], fTEIDs)

	return NewBearerContext(ies...)
}

// NewBearerContextWithinDeleteBearerRequest creates a new BearerContext used within DeleteBearerRequest.
func NewBearerContextWithinDeleteBearerRequest(ebi, cause *IE) *IE {
	return NewBearerContext(ebi, cause)
}

// NewBearerContextWithinDeleteBearerResponse creates a new BearerContext used within DeleteBearerResponse.
func NewBearerContextWithinDeleteBearerResponse(ebi, cause, pco, rannasCause, epco *IE) *IE {
	return NewBearerContext(ebi, cause, pco, rannasCause, epco)
}

// NewBearerContextWithinModifyBearerCommand creates a new BearerContext used within ModifyBearerCommand.
func NewBearerContextWithinModifyBearerCommand(ebi, qos *IE) *IE {
	return NewBearerContext(ebi, qos)
}

// NewBearerContextWithinUpdateBearerRequest creates a new BearerContext used within UpdateBearerRequest.
func NewBearerContextWithinUpdateBearerRequest(ebi, tft, qos, flags, pco, apco, epco, mplr *IE) *IE {
	return NewBearerContext(ebi, tft, qos, flags, pco, apco, epco, mplr)
}

// NewBearerContextWithinUpdateBearerResponse creates a new BearerContext used within UpdateBearerResponse.
func NewBearerContextWithinUpdateBearerResponse(ebi, cause, pco, rannasCause, epco *IE, fTEIDs ...*IE) *IE {
	n := 5 + len(fTEIDs)
	ies := make([]*IE, n)

	ies[0] = ebi
	ies[1] = cause
	ies[2] = pco
	ies[3] = rannasCause
	ies[4] = epco
	copy(ies[5:], fTEIDs)

	return NewBearerContext(ies...)
}

// NewBearerContextWithinDeleteBearerCommand creates a new BearerContext used within DeleteBearerCommand.
func NewBearerContextWithinDeleteBearerCommand(ebi, flags, rannasCause *IE) *IE {
	return NewBearerContext(ebi, flags, rannasCause)
}

// NewBearerContextWithinDeleteBearerFailureIndication creates a new BearerContext used within DeleteBearerFailureIndication.
func NewBearerContextWithinDeleteBearerFailureIndication(ebi, cause *IE) *IE {
	return NewBearerContext(ebi, cause)
}

// NewBearerContextWithinCreateIndirectDataForwardingTunnelRequest creates a new BearerContext used within CreateIndirectDataForwardingTunnelRequest.
func NewBearerContextWithinCreateIndirectDataForwardingTunnelRequest(ebi *IE, fTEIDs ...*IE) *IE {
	n := 1 + len(fTEIDs)
	ies := make([]*IE, n)

	ies[0] = ebi
	copy(ies[1:], fTEIDs)

	return NewBearerContext(ies...)
}

// NewBearerContextWithinCreateIndirectDataForwardingTunnelResponse creates a new BearerContext used within CreateIndirectDataForwardingTunnelResponse.
func NewBearerContextWithinCreateIndirectDataForwardingTunnelResponse(ebi, cause *IE, fTEIDs ...*IE) *IE {
	n := 2 + len(fTEIDs)
	ies := make([]*IE, n)

	ies[0] = ebi
	ies[1] = cause
	copy(ies[2:], fTEIDs)

	return NewBearerContext(ies...)
}

// NewBearerContextWithinForwardRelocationRequest creates a new BearerContext used within  ForwardRelocationRequest.
func NewBearerContextWithinForwardRelocationRequest(ebi, tft, qos, container, ti, flags *IE, fTEIDs ...*IE) *IE {
	n := 6 + len(fTEIDs)
	ies := make([]*IE, n)

	ies[0] = ebi
	ies[1] = tft
	ies[2] = qos
	ies[3] = container
	ies[4] = ti
	ies[5] = flags
	copy(ies[6:], fTEIDs)

	return NewBearerContext(ies...)
}

// NewBearerContextWithinContextResponse creates a new BearerContext used within  ContextResponse.
func NewBearerContextWithinContextResponse(ebi, tft, qos, container, ti *IE, fTEIDs ...*IE) *IE {
	n := 5 + len(fTEIDs)
	ies := make([]*IE, n)

	ies[0] = ebi
	ies[1] = tft
	ies[2] = qos
	ies[3] = container
	ies[4] = ti
	copy(ies[5:], fTEIDs)

	return NewBearerContext(ies...)
}

// NewBearerContextWithinContextAcknowledge creates a new BearerContext used within ContextAcknowledge.
func NewBearerContextWithinContextAcknowledge(ebi, fwdFTEID *IE) *IE {
	n := 2
	ies := make([]*IE, n)

	ies[0] = ebi
	ies[1] = fwdFTEID

	return NewBearerContext(ies...)
}

// BearerContext returns the []*IE inside BearerContext IE.
func (i *IE) BearerContext() ([]*IE, error) {
	if i.Type != BearerContext {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return nil, io.ErrUnexpectedEOF
	}

	ies, err := ParseMultiIEs(i.Payload)
	if err != nil {
		return nil, err
	}

	return ies, nil
}
