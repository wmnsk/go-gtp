// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"strconv"
)

// TODO: add setters
// TODO: consider using struct?

// NewIndication creates a new Indication IE.
// Note that each parameters should be 0 if false and 1 if true. Otherwise,
// the value won't be set as expected in the bitwise operations.
func NewIndication(
	daf, dtf, hi, dfi, oi, isrsi, israi, sgwci,
	sqci, uimsi, cfsi, crsi, ps, pt, si, msv,
	retLoc, pbic, srni, s6af, s4af, mbmdt, israu, ccrsi,
	cprai, arrl, ppoff, ppon, ppsi, csfbi, clii, cpsr,
	nsi, uasi, dtci, bdwi, psci, pcri, aosi, aopi,
	roaai, epcosi, cpopci, pmtmsi, s11tf, pnsi, unaccsi, wpmsi,
	fgsnn26, reprefi, fgsiwk, eevrsi, ltemui, ltempi, enbcrsi, tspcmi,
	csrmfi, mtedtn, mtedta, n5gnmi, fgcnrs, fgcnri, fsrhoi, ethpdn,
	sp98, sp97, sp96, sp95, sp94, sp93, sp92, emci uint8,
) *IE {
	i := New(Indication, 0x00, make([]byte, 9))
	i.Payload[0] |= sgwci
	i.Payload[0] |= (israi << 1)
	i.Payload[0] |= (isrsi << 2)
	i.Payload[0] |= (oi << 3)
	i.Payload[0] |= (dfi << 4)
	i.Payload[0] |= (hi << 5)
	i.Payload[0] |= (dtf << 6)
	i.Payload[0] |= (daf << 7)

	i.Payload[1] |= msv
	i.Payload[1] |= (si << 1)
	i.Payload[1] |= (pt << 2)
	i.Payload[1] |= (ps << 3)
	i.Payload[1] |= (crsi << 4)
	i.Payload[1] |= (cfsi << 5)
	i.Payload[1] |= (uimsi << 6)
	i.Payload[1] |= (sqci << 7)

	i.Payload[2] |= ccrsi
	i.Payload[2] |= (israu << 1)
	i.Payload[2] |= (mbmdt << 2)
	i.Payload[2] |= (s4af << 3)
	i.Payload[2] |= (s6af << 4)
	i.Payload[2] |= (srni << 5)
	i.Payload[2] |= (pbic << 6)
	i.Payload[2] |= (retLoc << 7)

	i.Payload[3] |= cpsr
	i.Payload[3] |= (clii << 1)
	i.Payload[3] |= (csfbi << 2)
	i.Payload[3] |= (ppsi << 3)
	i.Payload[3] |= (ppon << 4)
	i.Payload[3] |= (ppoff << 5)
	i.Payload[3] |= (arrl << 6)
	i.Payload[3] |= (cprai << 7)

	i.Payload[4] |= aopi
	i.Payload[4] |= (aosi << 1)
	i.Payload[4] |= (pcri << 2)
	i.Payload[4] |= (psci << 3)
	i.Payload[4] |= (bdwi << 4)
	i.Payload[4] |= (dtci << 5)
	i.Payload[4] |= (uasi << 6)
	i.Payload[4] |= (nsi << 7)

	i.Payload[5] |= wpmsi
	i.Payload[5] |= (unaccsi << 1)
	i.Payload[5] |= (pnsi << 2)
	i.Payload[5] |= (s11tf << 3)
	i.Payload[5] |= (pmtmsi << 4)
	i.Payload[5] |= (cpopci << 5)
	i.Payload[5] |= (epcosi << 6)
	i.Payload[5] |= (roaai << 7)

	i.Payload[6] |= tspcmi
	i.Payload[6] |= (enbcrsi << 1)
	i.Payload[6] |= (ltempi << 2)
	i.Payload[6] |= (ltemui << 3)
	i.Payload[6] |= (eevrsi << 4)
	i.Payload[6] |= (fgsiwk << 5)
	i.Payload[6] |= (reprefi << 6)
	i.Payload[6] |= (fgsnn26 << 7)

	i.Payload[7] |= ethpdn
	i.Payload[7] |= (fsrhoi << 1)
	i.Payload[7] |= (fgcnri << 2)
	i.Payload[7] |= (fgcnrs << 3)
	i.Payload[7] |= (n5gnmi << 4)
	i.Payload[7] |= (mtedta << 5)
	i.Payload[7] |= (mtedtn << 6)
	i.Payload[7] |= (csrmfi << 7)

	i.Payload[8] |= emci
	i.Payload[8] |= (sp92 << 1)
	i.Payload[8] |= (sp93 << 2)
	i.Payload[8] |= (sp94 << 3)
	i.Payload[8] |= (sp95 << 4)
	i.Payload[8] |= (sp96 << 5)
	i.Payload[8] |= (sp97 << 6)
	i.Payload[8] |= (sp98 << 7)
	return i
}

// NewIndicationFromBitSequence creates a new Indication IE from string-formatted
// sequence of bits. The sequence should look the same as Wireshark appearance.
// The input is to be like "101000010000100000010101000100001000100010000001010000001010000000000001".
func NewIndicationFromBitSequence(bitSequence string) *IE {
	if len(bitSequence) != 72 {
		return nil
	}

	ie := New(Indication, 0x00, make([]byte, 9))
	for i, r := range bitSequence {
		bit, err := strconv.Atoi(string(r))
		if err != nil {
			return nil
		}
		if bit > 1 {
			bit = 1
		}

		// index:
		//   0 =< i < 8 : set bits in first octet
		//   8 =< i < 16: set bits in second octet ...
		// bit shift:
		//   0 =< i < 8 : shift i to left
		//   8 =< i < 16: shift i - 8 to left ...
		ie.Payload[i/8] |= uint8(bit << uint8(8*(i/8+1)-i-1))
	}

	return ie
}

// NewIndicationFromOctets creates a new IndicationFromOctets IE from the set of octets.
func NewIndicationFromOctets(octs ...uint8) *IE {
	ie := New(Indication, 0x00, make([]byte, 0))
	for i, o := range octs {
		if i >= 9 {
			break
		}
		ie.Payload = append(ie.Payload, o)
	}
	ie.SetLength()
	return ie
}

// Indication returns Indication flags in []byte if the type of IE matches.
func (i *IE) Indication() ([]byte, error) {
	if i.Type != Indication {
		return nil, &InvalidTypeError{i.Type}
	}

	return i.Payload, nil
}

// HasSGWCI reports whether an IE has SGWCI bit.
func (i *IE) HasSGWCI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 1 {
		return false
	}

	return has1stBit(v[0])
}

// HasISRAI reports whether an IE has ISRAI bit.
func (i *IE) HasISRAI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 1 {
		return false
	}

	return has2ndBit(v[0])
}

// HasISRSI reports whether an IE has ISRSI bit.
func (i *IE) HasISRSI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 1 {
		return false
	}

	return has3rdBit(v[0])
}

// HasOI reports whether an IE has OI bit.
func (i *IE) HasOI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 1 {
		return false
	}

	return has4thBit(v[0])
}

// HasDFI reports whether an IE has DFI bit.
func (i *IE) HasDFI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 1 {
		return false
	}

	return has5thBit(v[0])
}

// HasHI reports whether an IE has HI bit.
func (i *IE) HasHI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 1 {
		return false
	}

	return has6thBit(v[0])
}

// HasDTF reports whether an IE has DTF bit.
func (i *IE) HasDTF() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 1 {
		return false
	}

	return has7thBit(v[0])
}

// HasDAF reports whether an IE has DAF bit.
func (i *IE) HasDAF() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 1 {
		return false
	}

	return has8thBit(v[0])
}

// HasMSV reports whether an IE has MSV bit.
func (i *IE) HasMSV() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 2 {
		return false
	}

	return has1stBit(v[1])
}

// HasSI reports whether an IE has SI bit.
func (i *IE) HasSI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 2 {
		return false
	}

	return has2ndBit(v[1])
}

// HasPT reports whether an IE has PT bit.
func (i *IE) HasPT() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 2 {
		return false
	}

	return has3rdBit(v[1])
}

// HasPS reports whether an IE has PS bit.
func (i *IE) HasPS() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 2 {
		return false
	}

	return has4thBit(v[1])
}

// HasCRSI reports whether an IE has CRSI bit.
func (i *IE) HasCRSI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 2 {
		return false
	}

	return has5thBit(v[1])
}

// HasCFSI reports whether an IE has CFSI bit.
func (i *IE) HasCFSI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 2 {
		return false
	}

	return has6thBit(v[1])
}

// HasUIMSI reports whether an IE has UIMSI bit.
func (i *IE) HasUIMSI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 2 {
		return false
	}

	return has7thBit(v[1])
}

// HasSQCI reports whether an IE has SQCI bit.
func (i *IE) HasSQCI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 2 {
		return false
	}

	return has8thBit(v[1])
}

// HasCCRSI reports whether an IE has CCRSI bit.
func (i *IE) HasCCRSI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 3 {
		return false
	}

	return has1stBit(v[2])
}

// HasISRAU reports whether an IE has ISRAU bit.
func (i *IE) HasISRAU() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 3 {
		return false
	}

	return has2ndBit(v[2])
}

// HasMBMDT reports whether an IE has MBMDT bit.
func (i *IE) HasMBMDT() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 3 {
		return false
	}

	return has3rdBit(v[2])
}

// HasS4AF reports whether an IE has S4AF bit.
func (i *IE) HasS4AF() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 3 {
		return false
	}

	return has4thBit(v[2])
}

// HasS6AF reports whether an IE has S6AF bit.
func (i *IE) HasS6AF() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 3 {
		return false
	}

	return has5thBit(v[2])
}

// HasSRNI reports whether an IE has SRNI bit.
func (i *IE) HasSRNI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 3 {
		return false
	}

	return has6thBit(v[2])
}

// HasPBIC reports whether an IE has PBIC bit.
func (i *IE) HasPBIC() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 3 {
		return false
	}

	return has7thBit(v[2])
}

// HasRETLOC reports whether an IE has RETLOC bit.
func (i *IE) HasRETLOC() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 3 {
		return false
	}

	return has8thBit(v[2])
}

// HasCPSR reports whether an IE has CPSR bit.
func (i *IE) HasCPSR() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 4 {
		return false
	}

	return has1stBit(v[3])
}

// HasCLII reports whether an IE has CLII bit.
func (i *IE) HasCLII() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 4 {
		return false
	}

	return has2ndBit(v[3])
}

// HasCSFBI reports whether an IE has CSFBI bit.
func (i *IE) HasCSFBI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 4 {
		return false
	}

	return has3rdBit(v[3])
}

// HasPPSI reports whether an IE has PPSI bit.
func (i *IE) HasPPSI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 4 {
		return false
	}

	return has4thBit(v[3])
}

// HasPPON reports whether an IE has PPON bit.
func (i *IE) HasPPON() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 4 {
		return false
	}

	return has5thBit(v[3])
}

// HasPPOF reports whether an IE has PPOF bit.
func (i *IE) HasPPOF() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 4 {
		return false
	}

	return has6thBit(v[3])
}

// HasARRL reports whether an IE has ARRL bit.
func (i *IE) HasARRL() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 4 {
		return false
	}

	return has7thBit(v[3])
}

// HasCPRAI reports whether an IE has CPRAI bit.
func (i *IE) HasCPRAI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 4 {
		return false
	}

	return has8thBit(v[3])
}

// HasAOPI reports whether an IE has AOPI bit.
func (i *IE) HasAOPI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 5 {
		return false
	}

	return has1stBit(v[4])
}

// HasAOSI reports whether an IE has AOSI bit.
func (i *IE) HasAOSI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 5 {
		return false
	}

	return has2ndBit(v[4])
}

// HasPCRI reports whether an IE has PCRI bit.
func (i *IE) HasPCRI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 5 {
		return false
	}

	return has3rdBit(v[4])
}

// HasPSCI reports whether an IE has PSCI bit.
func (i *IE) HasPSCI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 5 {
		return false
	}

	return has4thBit(v[4])
}

// HasBDWI reports whether an IE has BDWI bit.
func (i *IE) HasBDWI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 5 {
		return false
	}

	return has5thBit(v[4])
}

// HasDTCI reports whether an IE has DTCI bit.
func (i *IE) HasDTCI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 5 {
		return false
	}

	return has6thBit(v[4])
}

// HasUACI reports whether an IE has UACI bit.
func (i *IE) HasUACI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 5 {
		return false
	}

	return has7thBit(v[4])
}

// HasNSI reports whether an IE has NSI bit.
func (i *IE) HasNSI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 5 {
		return false
	}

	return has8thBit(v[4])
}

// HasWPMSI reports whether an IE has WPMSI bit.
func (i *IE) HasWPMSI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 6 {
		return false
	}

	return has1stBit(v[5])
}

// HasUNACCSI reports whether an IE has UNACCSI bit.
func (i *IE) HasUNACCSI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 6 {
		return false
	}

	return has2ndBit(v[5])
}

// HasPNSI reports whether an IE has PNSI bit.
func (i *IE) HasPNSI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 6 {
		return false
	}

	return has3rdBit(v[5])
}

// HasS11TF reports whether an IE has S11TF bit.
func (i *IE) HasS11TF() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 6 {
		return false
	}

	return has4thBit(v[5])
}

// HasPMTMSI reports whether an IE has PMTMSI bit.
func (i *IE) HasPMTMSI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 6 {
		return false
	}

	return has5thBit(v[5])
}

// HasCPOPCI reports whether an IE has CPOPCI bit.
func (i *IE) HasCPOPCI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 6 {
		return false
	}

	return has6thBit(v[5])
}

// HasEPCOSI reports whether an IE has EPCOSI bit.
func (i *IE) HasEPCOSI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 6 {
		return false
	}

	return has7thBit(v[5])
}

// HasROAAI reports whether an IE has ROAAI bit.
func (i *IE) HasROAAI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 6 {
		return false
	}

	return has8thBit(v[5])
}

// HasTSPCMI reports whether an IE has TSPCMI bit.
func (i *IE) HasTSPCMI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 7 {
		return false
	}

	return has1stBit(v[6])
}

// HasENBCRSI reports whether an IE has ENBCRSI bit.
func (i *IE) HasENBCRSI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 7 {
		return false
	}

	return has2ndBit(v[6])
}

// HasLTEMPI reports whether an IE has LTEMPI bit.
func (i *IE) HasLTEMPI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 7 {
		return false
	}

	return has3rdBit(v[6])
}

// HasLTEMUI reports whether an IE has LTEMUI bit.
func (i *IE) HasLTEMUI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 7 {
		return false
	}

	return has4thBit(v[6])
}

// HasEEVRSI reports whether an IE has EEVRSI bit.
func (i *IE) HasEEVRSI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 7 {
		return false
	}

	return has5thBit(v[6])
}

// Has5GSIWK reports whether an IE has 5GSIWK bit.
func (i *IE) Has5GSIWK() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 7 {
		return false
	}

	return has6thBit(v[6])
}

// HasREPREFI reports whether an IE has REPREFI bit.
func (i *IE) HasREPREFI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 7 {
		return false
	}

	return has7thBit(v[6])
}

// Has5GSNN26 reports whether an IE has 5GSNN26 bit.
func (i *IE) Has5GSNN26() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 7 {
		return false
	}

	return has8thBit(v[6])
}

// HasETHPDN reports whether an IE has ETHPDN bit.
func (i *IE) HasETHPDN() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 8 {
		return false
	}

	return has1stBit(v[7])
}

// Has5SRHOI reports whether an IE has 5SRHOI bit.
func (i *IE) Has5SRHOI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 8 {
		return false
	}

	return has2ndBit(v[7])
}

// Has5GCNRI reports whether an IE has 5GCNRI bit.
func (i *IE) Has5GCNRI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 8 {
		return false
	}

	return has3rdBit(v[7])
}

// Has5GCNRS reports whether an IE has 5GCNRS bit.
func (i *IE) Has5GCNRS() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 8 {
		return false
	}

	return has4thBit(v[7])
}

// HasN5GNMI reports whether an IE has N5GNMI bit.
func (i *IE) HasN5GNMI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 8 {
		return false
	}

	return has5thBit(v[7])
}

// HasMTEDTA reports whether an IE has MTEDTA bit.
func (i *IE) HasMTEDTA() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 8 {
		return false
	}

	return has6thBit(v[7])
}

// HasMTEDTN reports whether an IE has MTEDTN bit.
func (i *IE) HasMTEDTN() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 8 {
		return false
	}

	return has7thBit(v[7])
}

// HasCSRMFI reports whether an IE has CSRMFI bit.
func (i *IE) HasCSRMFI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 8 {
		return false
	}

	return has8thBit(v[7])
}

// HasEMCI reports whether an IE has EMCI bit.
func (i *IE) HasEMCI() bool {
	v, err := i.Indication()
	if err != nil {
		return false
	}
	if len(v) < 9 {
		return false
	}

	return has1stBit(v[8])
}
