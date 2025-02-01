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

func NewState(objectSource []byte, curAddr uint64) *State {
	return &State{
		objectSource: objectSource,
		curAddr:      curAddr,

		prefix: PrefixNONE,
	}
}

func (state *State) ParsePrefixInstrucsions() {
	// Prefix group 1
	if state.objectSource[state.curAddr] == 0xF0 ||
		state.objectSource[state.curAddr] == 0xF2 ||
		state.objectSource[state.curAddr] == 0xF3 {
		state.hasInstructionPrefix = true
		state.instructionPrefixByte = state.objectSource[state.curAddr]
		state.prefixOffset = 1
		state.disassembledInstructionSize++
		state.curAddr++
	}
}

func (state *State) ParseSegmentOverrridePrefix() {
	// Prefix group 2 segment override
	if segment, ok := SEGMENT_OVERRIDE[state.objectSource[state.curAddr]]; ok {
		state.hasSegmentOverridePrefix = true
		state.prefixSegmentOverrideStr = segment
		state.disassembledInstructionSize++
		state.curAddr++
	}
}

func (state *State) ParseBranch() {
	if state.objectSource[state.curAddr] == 0x66 {
		// TODO: add something
		state.disassembledInstructionSize++
		state.curAddr++
		// fmt.Printf("mnemonic nun: %x\n", state.objectSource[state.curAddr])
	}
}

func (state *State) ParseOperandSizeOverridePrefix() {
	// Prefix group 3
	if state.objectSource[state.curAddr] == 0x66 {
		state.prefix = PrefixP66
		state.disassembledInstructionSize++
		state.curAddr++
		// fmt.Printf("mnemonic nun: %x\n", state.objectSource[state.curAddr])
	}
}

func (state *State) ParseAddressSizeOverridePrefix() {
	// Prefix group 4
	if state.objectSource[state.curAddr] == 0x67 {
		// TODO: set prefix ?
		// state.opEnc = OpEncRM
		state.disassembledInstructionSize++
		state.curAddr++
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
	state.opcodeByte = state.objectSource[state.curAddr]
	state.disassembledInstructionSize++
	state.curAddr++

	pottentialOpCodeByte := (state.opcodeByte << 8) + state.objectSource[state.curAddr]

	_, okOpLookUp := OP_LOOKUP[PrefixOpcode{Prefix: state.prefix, Opcode: int(state.opcodeByte)}]
	// _, okOpLookUpRexw := OP_LOOKUP[PrefixOpcode{Prefix: PrefixREXW, Opcode: int(state.opcodeByte)}]
	_, okOpLookUpRex := OP_LOOKUP[PrefixOpcode{Prefix: PrefixREX, Opcode: int(state.opcodeByte)}]
	_, okOpLookUpNone := OP_LOOKUP[PrefixOpcode{Prefix: PrefixNONE, Opcode: int(state.opcodeByte)}]
	if slices.Contains(TWO_BYTES_OPCODE_PREFIX[:], int(state.opcodeByte)) || state.prefix == PrefixREXW && okOpLookUpRex || state.prefix == PrefixREX && okOpLookUpNone {
		state.opcodeByte = pottentialOpCodeByte
		state.disassembledInstructionSize++
		state.curAddr++
	}

	// (prefix, opcode) -> (reg, mnemonic)
	reg2mnem := make(map[int]Mnemonic)
	if okOpLookUp {
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
		log.Fatal("Unknown combination of the prefix and the opcodeByte: (" + PrefixToString(state.prefix) + ")")
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

	state.disassembledInstruction = append(state.disassembledInstruction, MnemonicToString(state.mnemonic))

	eleOperandLookUp, okOperandLookUp := OPERAND_LOOKUP[PrefixMnemonicInt{prefix: state.prefix, mnemonic: state.mnemonic, num: int(state.opcodeByte)}]
	if okOperandLookUp {
		state.opEnc = eleOperandLookUp.openc
		state.remOps = eleOperandLookUp.vecString
		// fmt.Println(eleOperandLookUp.vecOperand)
		state.operands = eleOperandLookUp.vecOperand
	} else {
		log.Fatal("Unknown combination of prefix, mnemonic and opcodeByte: (" + PrefixToString(state.prefix) + ", " + MnemonicToString(state.mnemonic) + ", )")
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

	// TODO: check EndBr
	// parse prefix
	state.ParsePrefixInstrucsions()
	state.ParseSegmentOverrridePrefix()
	state.ParseOperandSizeOverridePrefix()
	state.ParseAddressSizeOverridePrefix()
	// parsePrefix

	state.ParseREX()
	state.ParseOpecode()
	state.ParseModRM()
	state.ParseAddressOffset()

	// parse operand
	var imm []byte
	for _, operand := range state.operands {
		// fmt.Printf("parse operand loop\n")
		// fmt.Printf("operand: %x\n", operand)
		decodedTranslatedValue := ""

		if IsAReg(operand) || operand == OpCl || operand == OpDx {
			// fmt.Println("hoge1")
			decodedTranslatedValue = OperandToString(operand)
		} else if operand == OpSti {
			// fmt.Println("hoge2")
			decodedTranslatedValue = "st(" + state.remOps[0] + ")"
		} else if IsRM(operand) || IsREG(operand) || IsM(operand) {
			if HasModrm(state.opEnc) {
				if IsRM(operand) || IsM(operand) {
					// fmt.Println("hoge3")
					decodedTranslatedValue = state.modrm.GetAddrMode(operand, state.disp8, state.disp32)
					// fmt.Println(decodedTranslatedValue)
				} else {
					// fmt.Println("hoge4")
					decodedTranslatedValue = state.modrm.GetReg(operand)
					// fmt.Println(decodedTranslatedValue)
				}
			} else {
				// fmt.Println("bbb")
				var regIdx int
				if state.hasREX && state.rex.rexB {
					regIdx, _ = strconv.Atoi(state.remOps[0])
					regIdx += 8
				} else {
					regIdx, _ = strconv.Atoi(state.remOps[0])
				}

				if Is8Bit(operand) {
					// fmt.Println("hoge5")
					decodedTranslatedValue = REGISTERS8[regIdx]
				} else if Is16Bit(operand) {
					// fmt.Println("hoge6")
					decodedTranslatedValue = REGISTERS16[regIdx]
				} else if Is32Bit(operand) {
					// fmt.Println("hoge7")
					decodedTranslatedValue = REGISTERS32[regIdx]
				} else if operand == OpXm128 {
					// fmt.Println("hoge8")
					decodedTranslatedValue = "xmm" + state.remOps[0]
				}
			}

			if (IsRM(operand) || IsM(operand)) && HasModrm(state.opEnc) && state.modrm.hasSib {
				// fmt.Println("hoge9")
				decodedTranslatedValue = state.sib.GetAddr(operand, state.disp8, state.disp32)
			}
			if state.hasSegmentOverridePrefix {
				// fmt.Println("hoge11")
				decodedTranslatedValue = state.prefixSegmentOverrideStr + ":" + decodedTranslatedValue
			}

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
			// fmt.Println("hoge10")
			decodedTranslatedValue = tmpStr
		}

		state.disassembledOperands = append(state.disassembledOperands, decodedTranslatedValue)
	}

	// Append the mnemonic to the disassembled instruction
	instruction := MnemonicToString(state.mnemonic)
	state.disassembledInstruction = append(state.disassembledInstruction, instruction)
}
