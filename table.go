package main

import "fmt"

// PrefixOpcode holds a prefix and its associated opcode.
type PrefixOpcode struct {
	Prefix Prefix
	Opcode int
}

// NewPrefixOpcode creates and returns a new PrefixOpcode.
func NewPrefixOpcode(prefix Prefix, opcode int) (*PrefixOpcode, error) {
	if opcode < 0 {
		return nil, fmt.Errorf("invalid opcode: %d", opcode)
	}
	return &PrefixOpcode{
		Prefix: prefix,
		Opcode: opcode,
	}, nil
}

// RegOperator associates a register with an operator mnemonic.
type RegOperator struct {
	Reg      int
	Operator Mnemonic
}

// NewRegOperator creates and returns a new RegOperator.
func NewRegOperator(reg int, operator Mnemonic) (*RegOperator, error) {
	// Insert any validation logic if needed
	return &RegOperator{
		Reg:      reg,
		Operator: operator,
	}, nil
}

// OP_LOOKUP provides a mapping from PrefixOpcode to RegOperator.
var OP_LOOKUP = map[PrefixOpcode]RegOperator{
	// SETNE
	{Prefix: PrefixNONE, Opcode: 0x0F95}: {Reg: -1, Operator: MnSETNE},
	{Prefix: PrefixREX, Opcode: 0x0F95}:  {Reg: -1, Operator: MnSETNE},

	// FADD
	{Prefix: PrefixNONE, Opcode: 0xD8}:   {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xDC}:   {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xD8C0}: {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xD8C1}: {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xD8C2}: {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xD8C3}: {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xD8C4}: {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xD8C5}: {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xD8C6}: {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xD8C7}: {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xDCC0}: {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xDCC1}: {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xDCC2}: {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xDCC3}: {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xDCC4}: {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xDCC5}: {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xDCC6}: {Reg: -1, Operator: MnFADD},
	{Prefix: PrefixNONE, Opcode: 0xDCC7}: {Reg: -1, Operator: MnFADD},

	// FXCH
	{Prefix: PrefixNONE, Opcode: 0xD9C8}: {Reg: -1, Operator: MnFXCH},
	{Prefix: PrefixNONE, Opcode: 0xD9C9}: {Reg: -1, Operator: MnFXCH},
	{Prefix: PrefixNONE, Opcode: 0xD9CA}: {Reg: -1, Operator: MnFXCH},
	{Prefix: PrefixNONE, Opcode: 0xD9CB}: {Reg: -1, Operator: MnFXCH},
	{Prefix: PrefixNONE, Opcode: 0xD9CC}: {Reg: -1, Operator: MnFXCH},
	{Prefix: PrefixNONE, Opcode: 0xD9CD}: {Reg: -1, Operator: MnFXCH},
	{Prefix: PrefixNONE, Opcode: 0xD9CE}: {Reg: -1, Operator: MnFXCH},
	{Prefix: PrefixNONE, Opcode: 0xD9CF}: {Reg: -1, Operator: MnFXCH},

	// CMOVE
	{Prefix: PrefixNONE, Opcode: 0x0F44}: {Reg: -1, Operator: MnCMOVE},

	// MOVAPS
	{Prefix: PrefixNONE, Opcode: 0x0F28}: {Reg: -1, Operator: MnMOVAPS},
	{Prefix: PrefixNONE, Opcode: 0x0F29}: {Reg: -1, Operator: MnMOVAPS},

	// ENDBR
	{Prefix: PrefixNONE, Opcode: 0xF30F1EFA}: {Reg: -1, Operator: MnENDBR64},
	{Prefix: PrefixNONE, Opcode: 0xF30F1EFB}: {Reg: -1, Operator: MnENDBR32},

	// IN
	{Prefix: PrefixNONE, Opcode: 0xE4}: {Reg: -1, Operator: MnIN},
	{Prefix: PrefixNONE, Opcode: 0xE5}: {Reg: -1, Operator: MnIN},
	{Prefix: PrefixNONE, Opcode: 0xEC}: {Reg: -1, Operator: MnIN},
	{Prefix: PrefixNONE, Opcode: 0xED}: {Reg: -1, Operator: MnIN},

	// OUT
	{Prefix: PrefixNONE, Opcode: 0xE6}: {Reg: -1, Operator: MnOUT},
	{Prefix: PrefixNONE, Opcode: 0xE7}: {Reg: -1, Operator: MnOUT},
	{Prefix: PrefixNONE, Opcode: 0xEE}: {Reg: -1, Operator: MnOUT},
	{Prefix: PrefixNONE, Opcode: 0xEF}: {Reg: -1, Operator: MnOUT},

	// LOOP
	{Prefix: PrefixNONE, Opcode: 0xE2}: {Reg: -1, Operator: MnLOOP},
	{Prefix: PrefixNONE, Opcode: 0xE1}: {Reg: -1, Operator: MnLOOPE},
	{Prefix: PrefixNONE, Opcode: 0xE0}: {Reg: -1, Operator: MnLOOPNE},

	{Prefix: PrefixNONE, Opcode: 0xE2}:   {Reg: -1, Operator: MnLOOP},
	{Prefix: PrefixNONE, Opcode: 0x6C}:   {Reg: -1, Operator: MnINSB},
	{Prefix: PrefixNONE, Opcode: 0x6D}:   {Reg: -1, Operator: MnINSW},
	{Prefix: PrefixNONE, Opcode: 0x6E}:   {Reg: -1, Operator: MnOUTSB},
	{Prefix: PrefixNONE, Opcode: 0x6C}:   {Reg: -1, Operator: MnINSB},
	{Prefix: PrefixNONE, Opcode: 0xAE}:   {Reg: -1, Operator: MnSCASB},
	{Prefix: PrefixNONE, Opcode: 0xAF}:   {Reg: -1, Operator: MnSCASW},
	{Prefix: PrefixREXW, Opcode: 0xAF}:   {Reg: -1, Operator: MnSCASQ},
	{Prefix: PrefixNONE, Opcode: 0xAA}:   {Reg: -1, Operator: MnSTOSB},
	{Prefix: PrefixNONE, Opcode: 0xAB}:   {Reg: -1, Operator: MnSTOSW},
	{Prefix: PrefixREXW, Opcode: 0xAB}:   {Reg: -1, Operator: MnSTOSQ},
	{Prefix: PrefixNONE, Opcode: 0xAC}:   {Reg: -1, Operator: MnLODSB},
	{Prefix: PrefixNONE, Opcode: 0xAD}:   {Reg: -1, Operator: MnLODSW},
	{Prefix: PrefixREXW, Opcode: 0xAD}:   {Reg: -1, Operator: MnLODSW},
	{Prefix: PrefixNONE, Opcode: 0xA4}:   {Reg: -1, Operator: MnMOVSB},
	{Prefix: PrefixNONE, Opcode: 0xA5}:   {Reg: -1, Operator: MnMOVSW},
	{Prefix: PrefixREXW, Opcode: 0xA5}:   {Reg: -1, Operator: MnMOVSQ},
	{Prefix: PrefixNONE, Opcode: 0xA6}:   {Reg: -1, Operator: MnCMPSB},
	{Prefix: PrefixNONE, Opcode: 0xA7}:   {Reg: -1, Operator: MnCMPSW},
	{Prefix: PrefixREXW, Opcode: 0xA7}:   {Reg: -1, Operator: MnCMPSQ},
	{Prefix: PrefixNONE, Opcode: 0xC9}:   {Reg: -1, Operator: MnLEAVE},
	{Prefix: PrefixNONE, Opcode: 0x0FA2}: {Reg: -1, Operator: MnCPUID},
	{Prefix: PrefixNONE, Opcode: 0xF8}:   {Reg: -1, Operator: MnCLC},
	{Prefix: PrefixNONE, Opcode: 0xFC}:   {Reg: -1, Operator: MnCLD},
	{Prefix: PrefixNONE, Opcode: 0xF9}:   {Reg: -1, Operator: MnSTC},

	// ADD, ADC, SUB, SBB, AND, OR, XOR, CMP
	{Prefix: PrefixNONE, Opcode: 0x04}: {Reg: -1, Operator: MnADD},
	{Prefix: PrefixNONE, Opcode: 0x14}: {Reg: -1, Operator: MnADC},
	{Prefix: PrefixNONE, Opcode: 0x2C}: {Reg: -1, Operator: MnSUB},
	{Prefix: PrefixNONE, Opcode: 0x1C}: {Reg: -1, Operator: MnSBB},

	{Prefix: PrefixNONE, Opcode: 0x24}: {Reg: -1, Operator: MnAND},
	{Prefix: PrefixNONE, Opcode: 0x0C}: {Reg: -1, Operator: MnOR},
	{Prefix: PrefixNONE, Opcode: 0x34}: {Reg: -1, Operator: MnXOR},

	{Prefix: PrefixNONE, Opcode: 0x3C}: {Reg: -1, Operator: MnCMP},
	{Prefix: PrefixNONE, Opcode: 0x3D}: {Reg: -1, Operator: MnCMP},
	{Prefix: PrefixREXW, Opcode: 0x3D}: {Reg: -1, Operator: MnCMP},
	{Prefix: PrefixNONE, Opcode: 0x38}: {Reg: -1, Operator: MnCMP},
	{Prefix: PrefixREXW, Opcode: 0x38}: {Reg: -1, Operator: MnCMP},
	{Prefix: PrefixNONE, Opcode: 0x39}: {Reg: -1, Operator: MnCMP},
	{Prefix: PrefixREXW, Opcode: 0x39}: {Reg: -1, Operator: MnCMP},
	{Prefix: PrefixNONE, Opcode: 0x3A}: {Reg: -1, Operator: MnCMP},
	{Prefix: PrefixREXW, Opcode: 0x3A}: {Reg: -1, Operator: MnCMP},
	{Prefix: PrefixNONE, Opcode: 0x3B}: {Reg: -1, Operator: MnCMP},
	{Prefix: PrefixREXW, Opcode: 0x3B}: {Reg: -1, Operator: MnCMP},

	{Prefix: PrefixNONE, Opcode: 0x05}: {Reg: -1, Operator: MnADD},
	{Prefix: PrefixREXW, Opcode: 0x05}: {Reg: -1, Operator: MnADD},
	{Prefix: PrefixNONE, Opcode: 0x15}: {Reg: -1, Operator: MnADC},
	{Prefix: PrefixREXW, Opcode: 0x15}: {Reg: -1, Operator: MnADC},
	{Prefix: PrefixNONE, Opcode: 0x2D}: {Reg: -1, Operator: MnSUB},
	{Prefix: PrefixREXW, Opcode: 0x2D}: {Reg: -1, Operator: MnSUB},
	{Prefix: PrefixNONE, Opcode: 0x1D}: {Reg: -1, Operator: MnSBB},
	{Prefix: PrefixREXW, Opcode: 0x1D}: {Reg: -1, Operator: MnSBB},

	{Prefix: PrefixNONE, Opcode: 0x25}: {Reg: -1, Operator: MnAND},
	{Prefix: PrefixREXW, Opcode: 0x25}: {Reg: -1, Operator: MnAND},
	{Prefix: PrefixP66, Opcode: 0x25}:  {Reg: -1, Operator: MnAND},
	{Prefix: PrefixNONE, Opcode: 0x0D}: {Reg: -1, Operator: MnOR},
	{Prefix: PrefixREXW, Opcode: 0x0D}: {Reg: -1, Operator: MnOR},
	{Prefix: PrefixP66, Opcode: 0x0D}:  {Reg: -1, Operator: MnOR},
	{Prefix: PrefixNONE, Opcode: 0x35}: {Reg: -1, Operator: MnXOR},
	{Prefix: PrefixREXW, Opcode: 0x35}: {Reg: -1, Operator: MnXOR},
	{Prefix: PrefixP66, Opcode: 0x35}:  {Reg: -1, Operator: MnXOR},

	{Prefix: PrefixNONE, Opcode: 0x80}: {Reg: 0, Operator: MnADD},
	{Prefix: PrefixNONE, Opcode: 0x80}: {Reg: 2, Operator: MnADC},
	{Prefix: PrefixNONE, Opcode: 0x80}: {Reg: 5, Operator: MnSUB},
	{Prefix: PrefixNONE, Opcode: 0x80}: {Reg: 3, Operator: MnSBB},
	{Prefix: PrefixNONE, Opcode: 0x80}: {Reg: 4, Operator: MnAND},
	{Prefix: PrefixNONE, Opcode: 0x80}: {Reg: 1, Operator: MnOR},
	{Prefix: PrefixNONE, Opcode: 0x80}: {Reg: 6, Operator: MnXOR},
	{Prefix: PrefixNONE, Opcode: 0x80}: {Reg: 7, Operator: MnCMP},

	{Prefix: PrefixREX, Opcode: 0x80}: {Reg: 0, Operator: MnADD},
	{Prefix: PrefixREX, Opcode: 0x80}: {Reg: 2, Operator: MnADC},
	{Prefix: PrefixREX, Opcode: 0x80}: {Reg: 5, Operator: MnSUB},
	{Prefix: PrefixREX, Opcode: 0x80}: {Reg: 3, Operator: MnSBB},
	{Prefix: PrefixREX, Opcode: 0x80}: {Reg: 4, Operator: MnAND},
	{Prefix: PrefixREX, Opcode: 0x80}: {Reg: 1, Operator: MnOR},
	{Prefix: PrefixREX, Opcode: 0x80}: {Reg: 6, Operator: MnXOR},
	{Prefix: PrefixREX, Opcode: 0x80}: {Reg: 7, Operator: MnCMP},

	{Prefix: PrefixNONE, Opcode: 0x81}: {Reg: 0, Operator: MnADD},
	{Prefix: PrefixNONE, Opcode: 0x81}: {Reg: 2, Operator: MnADC},
	{Prefix: PrefixNONE, Opcode: 0x81}: {Reg: 5, Operator: MnSUB},
	{Prefix: PrefixNONE, Opcode: 0x81}: {Reg: 3, Operator: MnSBB},
	{Prefix: PrefixNONE, Opcode: 0x81}: {Reg: 4, Operator: MnAND},
	{Prefix: PrefixNONE, Opcode: 0x81}: {Reg: 1, Operator: MnOR},
	{Prefix: PrefixNONE, Opcode: 0x81}: {Reg: 6, Operator: MnXOR},
	{Prefix: PrefixNONE, Opcode: 0x81}: {Reg: 7, Operator: MnCMP},

	{Prefix: PrefixREXW, Opcode: 0x81}: {Reg: 0, Operator: MnADD},
	{Prefix: PrefixREXW, Opcode: 0x81}: {Reg: 2, Operator: MnADC},
	{Prefix: PrefixREXW, Opcode: 0x81}: {Reg: 5, Operator: MnSUB},
	{Prefix: PrefixREXW, Opcode: 0x81}: {Reg: 3, Operator: MnSBB},
	{Prefix: PrefixREXW, Opcode: 0x81}: {Reg: 4, Operator: MnAND},
	{Prefix: PrefixREXW, Opcode: 0x81}: {Reg: 1, Operator: MnOR},
	{Prefix: PrefixREXW, Opcode: 0x81}: {Reg: 6, Operator: MnXOR},
	{Prefix: PrefixREXW, Opcode: 0x81}: {Reg: 7, Operator: MnCMP},

	{Prefix: PrefixNONE, Opcode: 0x83}: {Reg: 0, Operator: MnADD},
	{Prefix: PrefixNONE, Opcode: 0x83}: {Reg: 2, Operator: MnADC},
	{Prefix: PrefixNONE, Opcode: 0x83}: {Reg: 5, Operator: MnSUB},
	{Prefix: PrefixNONE, Opcode: 0x83}: {Reg: 3, Operator: MnSBB},
	{Prefix: PrefixNONE, Opcode: 0x83}: {Reg: 4, Operator: MnAND},
	{Prefix: PrefixNONE, Opcode: 0x83}: {Reg: 1, Operator: MnOR},
	{Prefix: PrefixNONE, Opcode: 0x83}: {Reg: 6, Operator: MnXOR},
	{Prefix: PrefixNONE, Opcode: 0x83}: {Reg: 7, Operator: MnCMP},

	{Prefix: PrefixREXW, Opcode: 0x83}: {Reg: 0, Operator: MnADD},
	{Prefix: PrefixREXW, Opcode: 0x83}: {Reg: 2, Operator: MnADC},
	{Prefix: PrefixREXW, Opcode: 0x83}: {Reg: 5, Operator: MnSUB},
	{Prefix: PrefixREXW, Opcode: 0x83}: {Reg: 3, Operator: MnSBB},
	{Prefix: PrefixREXW, Opcode: 0x83}: {Reg: 4, Operator: MnAND},
	{Prefix: PrefixREXW, Opcode: 0x83}: {Reg: 1, Operator: MnOR},
	{Prefix: PrefixREXW, Opcode: 0x83}: {Reg: 6, Operator: MnXOR},
	{Prefix: PrefixREXW, Opcode: 0x83}: {Reg: 7, Operator: MnCMP},

	{Prefix: PrefixNONE, Opcode: 0x81}: {Reg: 0, Operator: MnADD},
	{Prefix: PrefixNONE, Opcode: 0x81}: {Reg: 2, Operator: MnADC},
	{Prefix: PrefixNONE, Opcode: 0x81}: {Reg: 5, Operator: MnSUB},
	{Prefix: PrefixNONE, Opcode: 0x81}: {Reg: 3, Operator: MnSBB},
	{Prefix: PrefixNONE, Opcode: 0x81}: {Reg: 4, Operator: MnAND},
	{Prefix: PrefixNONE, Opcode: 0x81}: {Reg: 1, Operator: MnOR},
	{Prefix: PrefixNONE, Opcode: 0x81}: {Reg: 6, Operator: MnXOR},
	{Prefix: PrefixNONE, Opcode: 0x81}: {Reg: 7, Operator: MnCMP},

	{Prefix: PrefixNONE, Opcode: 0x00}: {Reg: -1, Operator: MnADD},
	{Prefix: PrefixREX, Opcode: 0x00}:  {Reg: -1, Operator: MnADD},
	{Prefix: PrefixNONE, Opcode: 0x10}: {Reg: -1, Operator: MnADC},
	{Prefix: PrefixREX, Opcode: 0x10}:  {Reg: -1, Operator: MnADC},
	{Prefix: PrefixNONE, Opcode: 0x28}: {Reg: -1, Operator: MnSUB},
	{Prefix: PrefixREX, Opcode: 0x28}:  {Reg: -1, Operator: MnSUB},
	{Prefix: PrefixNONE, Opcode: 0x18}: {Reg: -1, Operator: MnSBB},
	{Prefix: PrefixREX, Opcode: 0x18}:  {Reg: -1, Operator: MnSBB},
	{Prefix: PrefixNONE, Opcode: 0x20}: {Reg: -1, Operator: MnAND},
	{Prefix: PrefixREX, Opcode: 0x20}:  {Reg: -1, Operator: MnAND},
	{Prefix: PrefixNONE, Opcode: 0x08}: {Reg: -1, Operator: MnOR},
	{Prefix: PrefixREX, Opcode: 0x08}:  {Reg: -1, Operator: MnOR},
	{Prefix: PrefixNONE, Opcode: 0x30}: {Reg: -1, Operator: MnXOR},
	{Prefix: PrefixREX, Opcode: 0x30}:  {Reg: -1, Operator: MnXOR},

	{Prefix: PrefixNONE, Opcode: 0x01}: {Reg: -1, Operator: MnADD},
	{Prefix: PrefixREXW, Opcode: 0x01}: {Reg: -1, Operator: MnADD},
	{Prefix: PrefixNONE, Opcode: 0x11}: {Reg: -1, Operator: MnADC},
	{Prefix: PrefixREXW, Opcode: 0x11}: {Reg: -1, Operator: MnADC},
	{Prefix: PrefixNONE, Opcode: 0x29}: {Reg: -1, Operator: MnSUB},
	{Prefix: PrefixREXW, Opcode: 0x29}: {Reg: -1, Operator: MnSUB},
	{Prefix: PrefixNONE, Opcode: 0x19}: {Reg: -1, Operator: MnSBB},
	{Prefix: PrefixREXW, Opcode: 0x19}: {Reg: -1, Operator: MnSBB},

	{Prefix: PrefixP66, Opcode: 0x20}:  {Reg: -1, Operator: MnAND},
	{Prefix: PrefixREXW, Opcode: 0x20}: {Reg: -1, Operator: MnAND},
	{Prefix: PrefixP66, Opcode: 0x08}:  {Reg: -1, Operator: MnOR},
	{Prefix: PrefixREXW, Opcode: 0x08}: {Reg: -1, Operator: MnOR},
	{Prefix: PrefixP66, Opcode: 0x30}:  {Reg: -1, Operator: MnXOR},
	{Prefix: PrefixREXW, Opcode: 0x30}: {Reg: -1, Operator: MnXOR},

	{Prefix: PrefixNONE, Opcode: 0x02}: {Reg: -1, Operator: MnADD},
	{Prefix: PrefixREX, Opcode: 0x02}:  {Reg: -1, Operator: MnADD},
	{Prefix: PrefixNONE, Opcode: 0x12}: {Reg: -1, Operator: MnADC},
	{Prefix: PrefixREX, Opcode: 0x12}:  {Reg: -1, Operator: MnADC},
	{Prefix: PrefixNONE, Opcode: 0x2A}: {Reg: -1, Operator: MnSUB},
	{Prefix: PrefixREX, Opcode: 0x2A}:  {Reg: -1, Operator: MnSUB},
	{Prefix: PrefixNONE, Opcode: 0x1A}: {Reg: -1, Operator: MnSBB},
	{Prefix: PrefixREX, Opcode: 0x1A}:  {Reg: -1, Operator: MnSBB},

	{Prefix: PrefixNONE, Opcode: 0x21}: {Reg: -1, Operator: MnAND},
	{Prefix: PrefixREXW, Opcode: 0x21}: {Reg: -1, Operator: MnAND},
	{Prefix: PrefixP66, Opcode: 0x21}:  {Reg: -1, Operator: MnAND},
	{Prefix: PrefixNONE, Opcode: 0x09}: {Reg: -1, Operator: MnOR},
	{Prefix: PrefixREXW, Opcode: 0x09}: {Reg: -1, Operator: MnOR},
	{Prefix: PrefixP66, Opcode: 0x09}:  {Reg: -1, Operator: MnOR},
	{Prefix: PrefixNONE, Opcode: 0x31}: {Reg: -1, Operator: MnXOR},
	{Prefix: PrefixREXW, Opcode: 0x31}: {Reg: -1, Operator: MnXOR},
	{Prefix: PrefixP66, Opcode: 0x31}:  {Reg: -1, Operator: MnXOR},

	{Prefix: PrefixNONE, Opcode: 0x03}: {Reg: -1, Operator: MnADD},
	{Prefix: PrefixREX, Opcode: 0x03}:  {Reg: -1, Operator: MnADD},
	{Prefix: PrefixNONE, Opcode: 0x13}: {Reg: -1, Operator: MnADC},
	{Prefix: PrefixREX, Opcode: 0x13}:  {Reg: -1, Operator: MnADC},
	{Prefix: PrefixNONE, Opcode: 0x2B}: {Reg: -1, Operator: MnSUB},
	{Prefix: PrefixREX, Opcode: 0x2B}:  {Reg: -1, Operator: MnSUB},
	{Prefix: PrefixNONE, Opcode: 0x1B}: {Reg: -1, Operator: MnSBB},
	{Prefix: PrefixREX, Opcode: 0x1B}:  {Reg: -1, Operator: MnSBB},

	{Prefix: PrefixNONE, Opcode: 0x22}: {Reg: -1, Operator: MnAND},
	{Prefix: PrefixREX, Opcode: 0x22}:  {Reg: -1, Operator: MnAND},
	{Prefix: PrefixNONE, Opcode: 0x0A}: {Reg: -1, Operator: MnOR},
	{Prefix: PrefixREX, Opcode: 0x0A}:  {Reg: -1, Operator: MnOR},
	{Prefix: PrefixNONE, Opcode: 0x32}: {Reg: -1, Operator: MnXOR},
	{Prefix: PrefixREX, Opcode: 0x32}:  {Reg: -1, Operator: MnXOR},

	{Prefix: PrefixNONE, Opcode: 0x23}: {Reg: -1, Operator: MnAND},
	{Prefix: PrefixREXW, Opcode: 0x23}: {Reg: -1, Operator: MnAND},
	{Prefix: PrefixP66, Opcode: 0x23}:  {Reg: -1, Operator: MnAND},
	{Prefix: PrefixNONE, Opcode: 0x0B}: {Reg: -1, Operator: MnOR},
	{Prefix: PrefixREXW, Opcode: 0x0B}: {Reg: -1, Operator: MnOR},
	{Prefix: PrefixP66, Opcode: 0x0B}:  {Reg: -1, Operator: MnOR},
	{Prefix: PrefixNONE, Opcode: 0x33}: {Reg: -1, Operator: MnXOR},
	{Prefix: PrefixREXW, Opcode: 0x33}: {Reg: -1, Operator: MnXOR},
	{Prefix: PrefixP66, Opcode: 0x33}:  {Reg: -1, Operator: MnXOR},

	// MOV
	{Prefix: PrefixNONE, Opcode: 0x88}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREX, Opcode: 0x88}:  {Reg: -1, Operator: MnMOV},

	{Prefix: PrefixP66, Opcode: 0x89}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0x89}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0x89}: {Reg: -1, Operator: MnMOV},

	{Prefix: PrefixNONE, Opcode: 0x8A}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREX, Opcode: 0x8A}:  {Reg: -1, Operator: MnMOV},

	{Prefix: PrefixP66, Opcode: 0x8B}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0x8B}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0x8B}: {Reg: -1, Operator: MnMOV},

	{Prefix: PrefixNONE, Opcode: 0x8C}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0x8C}: {Reg: -1, Operator: MnMOV},

	{Prefix: PrefixNONE, Opcode: 0x8C}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0x8C}: {Reg: -1, Operator: MnMOV},

	{Prefix: PrefixNONE, Opcode: 0x8E}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0x8E}: {Reg: -1, Operator: MnMOV},

	{Prefix: PrefixNONE, Opcode: 0xA0}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0xA0}: {Reg: -1, Operator: MnMOV},

	{Prefix: PrefixNONE, Opcode: 0xA1}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0xA1}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixP66, Opcode: 0xA1}:  {Reg: -1, Operator: MnMOV},

	{Prefix: PrefixNONE, Opcode: 0xA2}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0xA2}: {Reg: -1, Operator: MnMOV},

	{Prefix: PrefixNONE, Opcode: 0xA3}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0xA3}: {Reg: -1, Operator: MnMOV},

	{Prefix: PrefixNONE, Opcode: 0xB0}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREX, Opcode: 0xB0}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xB1}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREX, Opcode: 0xB1}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xB2}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREX, Opcode: 0xB2}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xB3}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREX, Opcode: 0xB3}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xB4}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREX, Opcode: 0xB4}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xB5}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREX, Opcode: 0xB5}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xB6}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREX, Opcode: 0xB6}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xB7}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREX, Opcode: 0xB7}:  {Reg: -1, Operator: MnMOV},

	{Prefix: PrefixP66, Opcode: 0xB8}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xB8}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0xB8}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixP66, Opcode: 0xB9}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xB9}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0xB9}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixP66, Opcode: 0xBA}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xBA}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0xBA}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixP66, Opcode: 0xBB}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xBB}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0xBB}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixP66, Opcode: 0xBC}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xBC}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0xBC}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixP66, Opcode: 0xBD}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xBD}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0xBD}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixP66, Opcode: 0xBE}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xBE}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0xBE}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixP66, Opcode: 0xBF}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xBF}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREXW, Opcode: 0xBF}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xC6}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREX, Opcode: 0xC6}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixP66, Opcode: 0xC7}:  {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixNONE, Opcode: 0xC7}: {Reg: -1, Operator: MnMOV},
	{Prefix: PrefixREX, Opcode: 0xC7}:  {Reg: -1, Operator: MnMOV},
}

type PrefixMnemonicInt struct {
	prefix   Prefix
	mnemonic Mnemonic
	num      int
}

type OpEncVecStrVecOprand struct {
	openc      OpEnc
	vecString  []string
	vecOperand []Operand
}

var OPERAND_LOOKUP = map[PrefixMnemonicInt]OpEncVecStrVecOprand{
	// SETNE
	{prefix: PrefixNONE, mnemonic: MnSETNE, num: 0x0F95}: {openc: OpEncM, vecString: []string{}, vecOperand: []Operand{OpRm8}},

	// MOV
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0x88}: {openc: OpEncMR, vecString: []string{"/r"}, vecOperand: []Operand{OpRm8, OpReg8}},
	{prefix: PrefixREX, mnemonic: MnMOV, num: 0x88}:  {openc: OpEncMR, vecString: []string{"/r"}, vecOperand: []Operand{OpRm8, OpReg8}},
	{prefix: PrefixP66, mnemonic: MnMOV, num: 0x89}:  {openc: OpEncMR, vecString: []string{"/r"}, vecOperand: []Operand{OpRm16, OpReg16}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0x89}: {openc: OpEncMR, vecString: []string{"/r"}, vecOperand: []Operand{OpRm32, OpReg32}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0x89}: {openc: OpEncMR, vecString: []string{"/r"}, vecOperand: []Operand{OpRm64, OpReg64}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0x8A}: {openc: OpEncRM, vecString: []string{"/r"}, vecOperand: []Operand{OpReg8, OpRm8}},
	{prefix: PrefixREX, mnemonic: MnMOV, num: 0x8A}:  {openc: OpEncRM, vecString: []string{"/r"}, vecOperand: []Operand{OpReg8, OpRm8}},
	{prefix: PrefixP66, mnemonic: MnMOV, num: 0x8B}:  {openc: OpEncRM, vecString: []string{"/r"}, vecOperand: []Operand{OpReg16, OpRm16}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0x8B}: {openc: OpEncRM, vecString: []string{"/r"}, vecOperand: []Operand{OpReg32, OpRm32}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0x8B}: {openc: OpEncRM, vecString: []string{"/r"}, vecOperand: []Operand{OpReg64, OpRm64}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0x8C}: {openc: OpEncMR, vecString: []string{"/r"}, vecOperand: []Operand{OpRm16, OpSreg}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0x8C}: {openc: OpEncMR, vecString: []string{"/r"}, vecOperand: []Operand{OpRm64, OpSreg}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0x8E}: {openc: OpEncRM, vecString: []string{"/r"}, vecOperand: []Operand{OpSreg, OpRm16}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0x8E}: {openc: OpEncRM, vecString: []string{"/r"}, vecOperand: []Operand{OpSreg, OpRm64}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xA0}: {openc: OpEncFD, vecString: []string{""}, vecOperand: []Operand{OpAl, OpMoffs8}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0xA0}: {openc: OpEncFD, vecString: []string{""}, vecOperand: []Operand{OpAl, OpMoffs8}},
	{prefix: PrefixP66, mnemonic: MnMOV, num: 0xA1}:  {openc: OpEncFD, vecString: []string{""}, vecOperand: []Operand{OpAx, OpMoffs16}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xA1}: {openc: OpEncFD, vecString: []string{""}, vecOperand: []Operand{OpEax, OpMoffs32}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0xA1}: {openc: OpEncFD, vecString: []string{""}, vecOperand: []Operand{OpRax, OpMoffs64}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xA2}: {openc: OpEncTD, vecString: []string{""}, vecOperand: []Operand{OpMoffs8, OpAl}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0xA2}: {openc: OpEncTD, vecString: []string{""}, vecOperand: []Operand{OpMoffs8, OpAl}},
	{prefix: PrefixP66, mnemonic: MnMOV, num: 0xA3}:  {openc: OpEncTD, vecString: []string{""}, vecOperand: []Operand{OpMoffs16, OpAx}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0xA3}: {openc: OpEncTD, vecString: []string{""}, vecOperand: []Operand{OpMoffs64, OpRax}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xB0}: {openc: OpEncOI, vecString: []string{"0"}, vecOperand: []Operand{OpReg8, OpImm8}},
	{prefix: PrefixREX, mnemonic: MnMOV, num: 0xB0}:  {openc: OpEncOI, vecString: []string{"0"}, vecOperand: []Operand{OpReg8, OpImm8}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xB1}: {openc: OpEncOI, vecString: []string{"1"}, vecOperand: []Operand{OpReg8, OpImm8}},
	{prefix: PrefixREX, mnemonic: MnMOV, num: 0xB1}:  {openc: OpEncOI, vecString: []string{"1"}, vecOperand: []Operand{OpReg8, OpImm8}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xB2}: {openc: OpEncOI, vecString: []string{"2"}, vecOperand: []Operand{OpReg8, OpImm8}},
	{prefix: PrefixREX, mnemonic: MnMOV, num: 0xB2}:  {openc: OpEncOI, vecString: []string{"2"}, vecOperand: []Operand{OpReg8, OpImm8}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xB3}: {openc: OpEncOI, vecString: []string{"3"}, vecOperand: []Operand{OpReg8, OpImm8}},
	{prefix: PrefixREX, mnemonic: MnMOV, num: 0xB3}:  {openc: OpEncOI, vecString: []string{"3"}, vecOperand: []Operand{OpReg8, OpImm8}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xB4}: {openc: OpEncOI, vecString: []string{"4"}, vecOperand: []Operand{OpReg8, OpImm8}},
	{prefix: PrefixREX, mnemonic: MnMOV, num: 0xB4}:  {openc: OpEncOI, vecString: []string{"4"}, vecOperand: []Operand{OpReg8, OpImm8}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xB5}: {openc: OpEncOI, vecString: []string{"5"}, vecOperand: []Operand{OpReg8, OpImm8}},
	{prefix: PrefixREX, mnemonic: MnMOV, num: 0xB5}:  {openc: OpEncOI, vecString: []string{"5"}, vecOperand: []Operand{OpReg8, OpImm8}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xB6}: {openc: OpEncOI, vecString: []string{"6"}, vecOperand: []Operand{OpReg8, OpImm8}},
	{prefix: PrefixREX, mnemonic: MnMOV, num: 0xB6}:  {openc: OpEncOI, vecString: []string{"6"}, vecOperand: []Operand{OpReg8, OpImm8}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xB7}: {openc: OpEncOI, vecString: []string{"7"}, vecOperand: []Operand{OpReg8, OpImm8}},
	{prefix: PrefixREX, mnemonic: MnMOV, num: 0xB7}:  {openc: OpEncOI, vecString: []string{"7"}, vecOperand: []Operand{OpReg8, OpImm8}},
	{prefix: PrefixP66, mnemonic: MnMOV, num: 0xB8}:  {openc: OpEncOI, vecString: []string{"0"}, vecOperand: []Operand{OpReg16, OpImm16}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xB8}: {openc: OpEncOI, vecString: []string{"0"}, vecOperand: []Operand{OpReg32, OpImm32}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0xB8}: {openc: OpEncOI, vecString: []string{"0"}, vecOperand: []Operand{OpReg64, OpImm64}},
	{prefix: PrefixP66, mnemonic: MnMOV, num: 0xB9}:  {openc: OpEncOI, vecString: []string{"1"}, vecOperand: []Operand{OpReg16, OpImm16}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xB9}: {openc: OpEncOI, vecString: []string{"1"}, vecOperand: []Operand{OpReg32, OpImm32}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0xB9}: {openc: OpEncOI, vecString: []string{"1"}, vecOperand: []Operand{OpReg64, OpImm64}},
	{prefix: PrefixP66, mnemonic: MnMOV, num: 0xBA}:  {openc: OpEncOI, vecString: []string{"2"}, vecOperand: []Operand{OpReg16, OpImm16}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xBA}: {openc: OpEncOI, vecString: []string{"2"}, vecOperand: []Operand{OpReg32, OpImm32}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0xBA}: {openc: OpEncOI, vecString: []string{"2"}, vecOperand: []Operand{OpReg64, OpImm64}},
	{prefix: PrefixP66, mnemonic: MnMOV, num: 0xBB}:  {openc: OpEncOI, vecString: []string{"3"}, vecOperand: []Operand{OpReg16, OpImm16}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xBB}: {openc: OpEncOI, vecString: []string{"3"}, vecOperand: []Operand{OpReg32, OpImm32}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0xBB}: {openc: OpEncOI, vecString: []string{"3"}, vecOperand: []Operand{OpReg64, OpImm64}},
	{prefix: PrefixP66, mnemonic: MnMOV, num: 0xBC}:  {openc: OpEncOI, vecString: []string{"4"}, vecOperand: []Operand{OpReg16, OpImm16}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xBC}: {openc: OpEncOI, vecString: []string{"4"}, vecOperand: []Operand{OpReg32, OpImm32}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0xBC}: {openc: OpEncOI, vecString: []string{"4"}, vecOperand: []Operand{OpReg64, OpImm64}},
	{prefix: PrefixP66, mnemonic: MnMOV, num: 0xBD}:  {openc: OpEncOI, vecString: []string{"5"}, vecOperand: []Operand{OpReg16, OpImm16}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xBD}: {openc: OpEncOI, vecString: []string{"5"}, vecOperand: []Operand{OpReg32, OpImm32}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0xBD}: {openc: OpEncOI, vecString: []string{"5"}, vecOperand: []Operand{OpReg64, OpImm64}},
	{prefix: PrefixP66, mnemonic: MnMOV, num: 0xBE}:  {openc: OpEncOI, vecString: []string{"6"}, vecOperand: []Operand{OpReg16, OpImm16}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xBE}: {openc: OpEncOI, vecString: []string{"6"}, vecOperand: []Operand{OpReg32, OpImm32}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0xBE}: {openc: OpEncOI, vecString: []string{"6"}, vecOperand: []Operand{OpReg64, OpImm64}},
	{prefix: PrefixP66, mnemonic: MnMOV, num: 0xBF}:  {openc: OpEncOI, vecString: []string{"7"}, vecOperand: []Operand{OpReg16, OpImm16}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xBF}: {openc: OpEncOI, vecString: []string{"7"}, vecOperand: []Operand{OpReg32, OpImm32}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0xBF}: {openc: OpEncOI, vecString: []string{"7"}, vecOperand: []Operand{OpReg64, OpImm64}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xC6}: {openc: OpEncMI, vecString: []string{"ib"}, vecOperand: []Operand{OpRm8, OpImm8}},
	{prefix: PrefixREX, mnemonic: MnMOV, num: 0xC6}:  {openc: OpEncMI, vecString: []string{"ib"}, vecOperand: []Operand{OpRm8, OpImm8}},
	{prefix: PrefixP66, mnemonic: MnMOV, num: 0xC7}:  {openc: OpEncMI, vecString: []string{"iw"}, vecOperand: []Operand{OpRm16, OpImm16}},
	{prefix: PrefixNONE, mnemonic: MnMOV, num: 0xC7}: {openc: OpEncMI, vecString: []string{"id"}, vecOperand: []Operand{OpRm32, OpImm32}},
	{prefix: PrefixREXW, mnemonic: MnMOV, num: 0xC7}: {openc: OpEncMI, vecString: []string{"id"}, vecOperand: []Operand{OpRm64, OpImm32}},

	// ..etc
}

var SEGMENT_OVERRIDE = map[byte]string{
	// prefix group 2 segment override
	0x2E: "CS",
	0x36: "SS",
	0x3E: "DS",
	0x26: "ES",
	0x64: "FS",
	0x65: "GS",
}
