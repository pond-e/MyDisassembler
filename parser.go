package main

import (
	"fmt"
	"log"
	"strconv"

	"golang.org/x/exp/slices"
)

func byte2int(val []byte) int {
	str := string(val)

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return i
}

func decodeOffset(val byte) int {
	offset := int(val)

	if offset <= 0xFF {
		if offset > 0x7F {
			offset -= 0x100
		}
	} else if offset <= 0xFFFF {
		if offset > 0x7FFF {
			offset -= 0x10000
		}
	} else if offset <= 0xFFFFFFFF {
		if offset > 0x7FFFFFFF {
			offset -= 0x100000000
		}
	}

	return offset
}

func decodeOffsetArray(val []byte) int {
	offset := byte2int(val)

	if offset <= 0xFF {
		if offset > 0x7F {
			offset -= 0x100
		}
	} else if offset <= 0xFFFF {
		if offset > 0x7FFF {
			offset -= 0x10000
		}
	} else if offset <= 0xFFFFFFFF {
		if offset > 0x7FFFFFFF {
			offset -= 0x100000000
		}
	}

	return offset
}

type State struct {
	objectSource []byte
	addr2symbol  map[uint64]int

	hasInstructionPrefix     bool
	hasSegmentOverridePrefix bool
	hasREX                   bool
	hasSIB                   bool
	hasDisp8                 bool
	hasDisp32                bool

	curAddr                     uint64
	disassembledInstructionSize uint64
	prefixOffset                uint64

	instructionPrefixByte byte
	opcodeByte            byte
	modrmByte             byte
	sibByte               byte

	prefixInstructionStr     string
	prefixSegmentOverrideStr string

	disassembledInstruction []string
	disassembledOperands    []string

	mnemonic Mnemonic
	prefix   Prefix
	rex      REX
	modrm    ModRM
	sib      SIB

	opEnc    OpEnc
	remOps   []string
	operands []Operand

	disp8  string
	disp32 string
}

type REX struct {
	rexB bool
	rexX bool
	rexR bool
	rexW bool
}

func NewREX(rexByte byte) *REX {
	rexB := (rexByte & 0x1) == 0x1
	rexX := (rexByte & 0x2) == 0x2
	rexR := (rexByte & 0x4) == 0x4
	rexW := (rexByte & 0x8) == 0x8
	return &REX{
		rexB: rexB,
		rexX: rexX,
		rexR: rexR,
		rexW: rexW,
	}
}

func (state *State) ParseREX() {
	if (state.objectSource[state.curAddr] >> 4) == 4 { // 上位4ビットが0100
		state.hasREX = true
		state.rex = *NewREX(state.objectSource[state.curAddr])
		state.disassembledInstructionSize++
		state.curAddr++

		if state.rex.rexW {
			state.prefix = PrefixREXW
		} else {
			state.prefix = PrefixREX
		}
	}
}

type ModRM struct {
	rex     REX
	modByte byte
	regByte byte
	rmByte  byte

	hasDisp8  bool
	hasDisp32 bool
	hasSib    bool
}

func NewModRM(modrmByte byte, rex REX) *ModRM {
	rmByte := modrmByte & 0x7
	regByte := (modrmByte >> 3) & 0x7
	modByte := (modrmByte >> 6) & 0x3
	hasSib := false
	var hasDisp8, hasDisp32 bool

	if modByte < 3 && rmByte == 4 {
		hasSib = true
	}
	switch modByte {
	case 0:
		hasDisp8 = false
		hasDisp32 = false
	case 1:
		hasDisp8 = true
		hasDisp32 = false
	case 2:
		hasDisp8 = false
		hasDisp32 = true
	case 3:
		hasDisp8 = false
		hasDisp32 = false
	}

	if modByte == 0 && rmByte == 5 {
		hasDisp8 = false
		hasDisp32 = true
	}

	return &ModRM{
		rex:       rex,
		modByte:   modByte,
		regByte:   regByte,
		rmByte:    rmByte,
		hasDisp8:  hasDisp8,
		hasDisp32: hasDisp32,
		hasSib:    hasSib,
	}
}

func (modrm *ModRM) GetReg(operand Operand) string {
	var addNum byte = 0
	if modrm.rex.rexB {
		addNum = 8
	}

	if operand == OpXmm {
		return "xmm" + strconv.Itoa(int(modrm.regByte+addNum))
	} else {
		return Operand2Register(operand)[modrm.regByte+addNum]
	}
}

func (modrm *ModRM) GetAddrMode(operand Operand, disp8 string, disp32 string) string {
	var addrBaseReg string
	var addNum byte = 0
	if modrm.rex.rexB {
		addNum = 8
	}
	if modrm.modByte == 3 {
		if operand == OpXm128 {
			addrBaseReg = "xmm" + strconv.Itoa(int(modrm.rmByte+addNum))
		} else {
			addrBaseReg = "xmm" + strconv.Itoa(int(modrm.rmByte+addNum))
		}
	} else {
		if operand == OpXm128 {
			addrBaseReg = "xmm" + strconv.Itoa(int(modrm.rmByte+addNum))
		} else {
			addrBaseReg = REGISTERS64[modrm.rmByte+addNum]
		}
	}

	if modrm.modByte < 3 && modrm.rmByte == 4 {
		addrBaseReg = "SIB"
	}

	var addressingMode string
	switch modrm.modByte {
	case 0:
		addressingMode = "[" + addrBaseReg + "]"
	case 1:
		addressingMode = "[" + addrBaseReg + disp8 + "]"
	case 2:
		addressingMode = "[" + addrBaseReg + disp32 + "]"
	case 3:
		addressingMode = addrBaseReg
	}

	if modrm.modByte == 0 && modrm.rmByte == 5 {
		addressingMode = "[rip" + disp32 + "]"
	}

	return addressingMode
}

func (state *State) ParseModRM() {
	if HasModrm(state.opEnc) {
		if state.modrmByte < 0 {
			log.Fatal("Expected ModRM byte but there aren't any bytes left.")
		}
		state.disassembledInstructionSize++
		state.curAddr++
		state.modrm = *NewModRM(state.modrmByte, state.rex)
	}
}

type SIB struct {
	scaleByte byte
	indexByte byte
	baseByte  byte
	modByte   byte
	rex       REX

	address     string
	addrBaseReg string
	indexReg    string
	scale       int
	hasDisp8    bool
	hasDisp32   bool
}

func NewSIB(sibByte byte, modByte byte, rex REX) *SIB {
	scaleByte := (sibByte >> 6) & 0x3
	indexByte := (sibByte >> 3) & 0x7
	baseByte := sibByte & 0x7

	hasDisp8 := baseByte == 5 && modByte == 1
	hasDisp32 := baseByte == 5 && modByte != 1

	return &SIB{
		scaleByte:   scaleByte,
		indexByte:   indexByte,
		baseByte:    baseByte,
		rex:         rex,
		address:     "",
		addrBaseReg: "",
		indexReg:    "",
		scale:       0,
		hasDisp8:    hasDisp8,
		hasDisp32:   hasDisp32,
	}
}

func (sib *SIB) GetAddr(operand Operand, disp8 string, disp32 string) string {
	offset := ""

	if sib.baseByte == 5 {
		switch sib.modByte {
		case 0:
			if len(disp32) > 2 {
				if disp32[1] == '+' {
					sib.addrBaseReg = disp32[3:]
				} else {
					sib.addrBaseReg = "-" + disp32[3:]
				}
			}
		case 1:
			if sib.rex.rexB {
				sib.addrBaseReg = "r13"
			} else {
				sib.addrBaseReg = "rbp"
			}
			offset = disp8
		case 2:
			if sib.rex.rexB {
				sib.addrBaseReg = "r13"
			} else {
				sib.addrBaseReg = "rbp"
			}
			offset = disp32
		}
	} else {
		var addNum byte = 0
		if sib.rex.rexB {
			addNum = 8
		}
		sib.addrBaseReg = REGISTERS64[sib.baseByte+addNum]
		if sib.modByte == 1 {
			offset = disp8
		} else if sib.modByte == 2 {
			offset = disp32
		}
	}

	if sib.modByte == 0 && sib.baseByte == 5 && sib.indexByte == 4 {
		sib.address = sib.addrBaseReg
	} else if sib.indexByte == 4 && (!sib.rex.rexX) {
		sib.address = "[" + sib.addrBaseReg + offset + "]"
	} else {
		var addNum byte = 0
		if sib.rex.rexX {
			addNum = 8
		}
		sib.indexReg = REGISTERS64[sib.indexByte+addNum]
		sib.scale = SCALE_FACTOR[sib.scaleByte]
		sib.address = "[" + sib.addrBaseReg + " + " + sib.indexReg + " * " + strconv.Itoa(sib.scale) + offset + "]"
	}
	return sib.address
}

func (state *State) ParseSIB() {
	if HasModrm(state.opEnc) && state.modrm.hasSib {
		if state.curAddr < uint64(len(state.objectSource)) {
			state.sibByte = state.objectSource[state.curAddr]
		}
		if state.sibByte < 0 {
			log.Fatal("Expected SIB byte but there aren't any bytes left.")
		}
	}
	state.sib = *NewSIB(state.sibByte, state.modrm.modByte, state.rex)
	state.disassembledInstructionSize++
	state.curAddr++
}

func (state *State) ParseOpecode() {
	opcodeByte := state.objectSource[state.curAddr]
	state.disassembledInstructionSize++
	state.curAddr++

	pottentialOpCodeByte := (opcodeByte << 8) + state.objectSource[state.curAddr]

	_, okOPLookUpTwoBytes := OP_LOOKUP[PrefixOpcode{Prefix: state.prefix, Opcode: int(state.opcodeByte)}] // TODO: 修正
	// _, okOpLookUpRexw := OP_LOOKUP[PrefixOpcode{Prefix: PrefixREXW, Opcode: int(state.opcodeByte)}]
	_, okOpLookUpRex := OP_LOOKUP[PrefixOpcode{Prefix: PrefixREX, Opcode: int(state.opcodeByte)}]
	_, okOpLookUpNone := OP_LOOKUP[PrefixOpcode{Prefix: PrefixNONE, Opcode: int(state.opcodeByte)}]
	if slices.Contains(TWO_BYTES_OPCODE_PREFIX[:], int(state.opcodeByte)) && okOPLookUpTwoBytes || state.prefix == PrefixREXW && okOpLookUpRex || state.prefix == PrefixREX && okOpLookUpNone {
		state.opcodeByte = pottentialOpCodeByte
		state.disassembledInstructionSize++
		state.curAddr++
	}

	// (prefix, opcode) -> (reg, mnemonic)
	reg2mnem := make(map[int]Mnemonic)
	if okOPLookUpTwoBytes {
		reg2mnemTmp := OP_LOOKUP[PrefixOpcode{Prefix: state.prefix, Opcode: int(state.opcodeByte)}]
		reg2mnem[reg2mnemTmp.Reg] = reg2mnemTmp.Operator
	} else if state.prefix == PrefixREXW && okOpLookUpRex {
		reg2mnemTmp := OP_LOOKUP[PrefixOpcode{Prefix: PrefixREX, Opcode: int(state.opcodeByte)}]
		reg2mnem[reg2mnemTmp.Reg] = reg2mnemTmp.Operator
		state.prefix = PrefixREX
	} else if state.prefix == PrefixREX && okOpLookUpNone {
		reg2mnemTmp := OP_LOOKUP[PrefixOpcode{Prefix: PrefixNONE, Opcode: int(state.opcodeByte)}]
		reg2mnem[reg2mnemTmp.Reg] = reg2mnemTmp.Operator
		state.prefix = PrefixNONE
	} else {
		fmt.Printf("%x\n", state.opcodeByte)
		log.Fatal("Unknown combination of the prefix and the opcodeByte: (" + strconv.Itoa(int(state.prefix)) + ")")
	}

	// We sometimes need reg of modrm to determine the opecode
	// e.g. 83 /4 -> AND
	//      83 /1 -> OR
	if state.curAddr < uint64(len(state.objectSource)) {
		state.modrmByte = state.objectSource[state.curAddr]
	}

	if state.modrmByte >= 0 {
		reg := (state.modrmByte >> 3) & 0x7
		if mnem, ok := reg2mnem[int(reg)]; ok {
			state.mnemonic = mnem
		} else if defaultMnem, ok := reg2mnem[-1]; ok {
			state.mnemonic = defaultMnem
		} else {
			log.Fatal("Unable to determine mnemonic for opcode")
		}
	} else if defaultMnem, ok := reg2mnem[-1]; ok {
		state.mnemonic = defaultMnem
	} else {
		log.Fatal("Unable to determine mnemonic for opcode")
	}

	if state.hasInstructionPrefix {
		if state.instructionPrefixByte == 0xF0 {
			state.disassembledInstruction = append(state.disassembledInstruction, "lock")
		} else if state.instructionPrefixByte == 0xF2 {
			if IsControlFlowInstruction(state.mnemonic) {
				state.disassembledInstruction = append(state.disassembledInstruction, "bnd")
			} else {
				state.disassembledInstruction = append(state.disassembledInstruction, "repne")
			}
		} else if state.instructionPrefixByte == 0xF3 {
			state.disassembledInstruction = append(state.disassembledInstruction, "rep")
		} else if state.instructionPrefixByte == 0x3E {
			state.disassembledInstruction = append(state.disassembledInstruction, "notrack")
		}
	}

	state.disassembledInstruction = append(state.disassembledInstruction, strconv.Itoa(int(state.mnemonic))) // TODO: append mnemonic string

	eleOperandLookUp, okOperandLookUp := OPERAND_LOOKUP[PrefixMnemonicInt{prefix: state.prefix, mnemonic: state.mnemonic, num: int(state.opcodeByte)}]
	if okOperandLookUp {
		state.opEnc = eleOperandLookUp.openc
		state.remOps = eleOperandLookUp.vecString
		state.operands = eleOperandLookUp.vecOperand
	} else {
		fmt.Printf("%x\n", state.opcodeByte)
		log.Fatal("Unknown combination of prefix, mnemonic and opcodeByte: (" + strconv.Itoa(int(state.prefix)) + ", " + strconv.Itoa(int(state.mnemonic)) + ", )")
	}
}

func (state *State) ParseAddressOffset() {
	if HasModrm(state.opEnc) && state.modrm.hasDisp8 ||
		HasModrm(state.opEnc) && state.modrm.hasSib && state.sib.hasDisp8 ||
		HasModrm(state.opEnc) && state.modrm.hasSib && state.modrm.modByte == 1 && state.sib.baseByte == 5 {
		disp8 := decodeOffset(state.objectSource[state.curAddr])
		state.disp8 = strconv.Itoa(disp8)

		state.hasDisp8 = true
		state.disassembledInstructionSize++
		state.curAddr++
	}

	if HasModrm(state.opEnc) && state.modrm.hasDisp32 ||
		HasModrm(state.opEnc) && state.modrm.hasSib && state.sib.hasDisp32 ||
		HasModrm(state.opEnc) && state.modrm.hasSib && state.modrm.modByte == 2 && state.sib.baseByte == 5 {
		disp32 := decodeOffsetArray(state.objectSource[state.curAddr : state.curAddr+4])
		state.disp32 = strconv.Itoa(disp32)

		state.hasDisp32 = true
		state.disassembledInstructionSize += 4
		state.curAddr += 4
	}
}

type DisassembledResult struct {
	startAddr                   uint64
	disassembledInstructionSize uint64
	mnemonic                    Mnemonic
	disassembledInstructionStr  []string
	nextOffset                  int64
}

func reverseBytes(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

// func (state *State) step(startAddr uint64) DisassembledResult {
func (state *State) step(startAddr uint64) {
	// init
	state.curAddr = startAddr

	// parse prefix
	// parsePrefixInstrucsions()
	// parsesegmentOverrridePrefix()
	// parsePrefix

	state.ParseREX()
	state.ParseOpecode()
	state.ParseModRM()
	state.ParseAddressOffset()

	// parse operand
	var imm []byte
	for _, operand := range state.operands {
		decodedTranslatedValue := ""

		if IsAReg(operand) || operand == OpCl || operand == OpDx {
			decodedTranslatedValue = strconv.Itoa(int(operand)) // TODO: fix
		} else if operand == OpSti {
			decodedTranslatedValue = "st(" + state.remOps[0] + ")"
		} else if IsRM(operand) || IsREG(operand) || IsM(operand) {
			if HasModrm(state.opEnc) {
				if IsRM(operand) || IsM(operand) {
					decodedTranslatedValue = state.modrm.GetAddrMode(operand, state.disp8, state.disp32)
				} else {
					decodedTranslatedValue = state.modrm.GetReg(operand)
				}
			} else {
				var regIdx int
				if state.hasREX && state.rex.rexB {
					regIdx, _ = strconv.Atoi(state.remOps[0])
					regIdx += 8
				} else {
					regIdx, _ = strconv.Atoi(state.remOps[0])
				}

				if Is8Bit(operand) {
					decodedTranslatedValue = REGISTERS8[regIdx]
				} else if Is16Bit(operand) {
					decodedTranslatedValue = REGISTERS16[regIdx]
				} else if Is32Bit(operand) {
					decodedTranslatedValue = REGISTERS32[regIdx]
				} else if operand == OpXm128 {
					decodedTranslatedValue = "xmm" + state.remOps[0]
				}
			}

			if (IsRM(operand) || IsM(operand)) && HasModrm(state.opEnc) && state.modrm.hasSib {
				decodedTranslatedValue = state.sib.GetAddr(operand, state.disp8, state.disp32)
			}
			// if hasSegmentOverridePrefix

		} else if IsIMM(operand) {
			immSize := 0
			if operand == OpImm64 || operand == OpYmm {
				immSize = 8
			} else if operand == OpImm32 || operand == OpXmm {
				immSize = 4
			} else if operand == OpImm16 {
				immSize = 2
			} else if operand == OpImm8 {
				immSize = 1
			}
			imm = state.objectSource[state.curAddr : state.curAddr+uint64(immSize)]
			imm = reverseBytes(imm)
			state.disassembledInstructionSize += uint64(immSize)
			state.curAddr += uint64(immSize)

			tmpStr := "0x"
			for i := 0; i < len(imm); i++ {
				tmpStr += fmt.Sprintf("%x", imm[i])
			}
			decodedTranslatedValue = tmpStr
		}

		state.disassembledOperands = append(state.disassembledOperands, decodedTranslatedValue)
	}

}
