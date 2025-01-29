package main

type Operand int

const (
	OpOne Operand = iota
	OpImm8
	OpImm16
	OpImm32
	OpImm64
	OpReg8
	OpReg16
	OpReg32
	OpReg64
	OpRm8
	OpRm16
	OpRm32
	OpRm64
	OpReg
	OpSreg
	OpXmm
	OpYmm
	OpXm128
	OpYm256
	OpM
	OpM32fp
	OpM64fp
	OpAl
	OpAx
	OpEax
	OpRax
	OpMoffs8
	OpMoffs16
	OpMoffs32
	OpMoffs64
	OpCl
	OpDx
	OpSt0
	OpSti
)

func IsAReg(operand Operand) bool {
	return operand == OpAl || operand == OpAx ||
		operand == OpEax || operand == OpRax
}

func IsRM(operand Operand) bool {
	return operand == OpRm8 || operand == OpRm16 ||
		operand == OpEax || operand == OpRax
}

func IsM(operand Operand) bool {
	return operand == OpM || operand == OpM32fp || operand == OpM64fp
}

func IsREG(operand Operand) bool {
	return operand == OpReg8 || operand == OpReg16 ||
		operand == OpReg32 || operand == OpReg64 ||
		operand == OpXmm || operand == OpYmm
}

func IsIMM(operand Operand) bool {
	return operand == OpImm8 || operand == OpImm16 ||
		operand == OpImm32 || operand == OpImm64
}

func Is8Bit(operand Operand) bool {
	return operand == OpRm8 || operand == OpReg8
}

func Is16Bit(operand Operand) bool {
	return operand == OpRm16 || operand == OpReg16
}

func Is32Bit(operand Operand) bool {
	return operand == OpRm32 || operand == OpReg32
}

func Is64Bit(operand Operand) bool {
	return operand == OpRm64 || operand == OpReg64
}

type OpEnc int

const (
	OpEncI OpEnc = iota
	OpEncD
	OpEncM
	OpEncO
	OpEncNP
	OpEncMC
	OpEncMI
	OpEncM1
	OpEncMR
	OpEncRM
	OpEncRMI
	OpEncMRI
	OpEncMRC
	OpEncOI
	OpEncFD
	OpEncTD
	OpEncS
	OpEncA
	OpEncB
	OpEncC
)

var TWO_BYTES_OPCODE_PREFIX = [4]int{0x0F, 0xD8, 0xD9, 0xDC}

func IsJCCInstruction(mnemonic Mnemonic) bool {
	return mnemonic == MnJO || mnemonic == MnJNO ||
		mnemonic == MnJNAE || mnemonic == MnJNB ||
		mnemonic == MnJZ || mnemonic == MnJNZ ||
		mnemonic == MnJNA || mnemonic == MnJNBE ||
		mnemonic == MnJS || mnemonic == MnJNS ||
		mnemonic == MnJP || mnemonic == MnJPO ||
		mnemonic == MnJNGE || mnemonic == MnJNL ||
		mnemonic == MnJNG || mnemonic == MnJNLE
}

func IsLOOPInstruction(mnemonic Mnemonic) bool {
	return mnemonic == MnLOOP || mnemonic == MnLOOPE || mnemonic == MnLOOPNE
}

func IsControlFlowInstruction(mnemonic Mnemonic) bool {
	return mnemonic == MnCALL || mnemonic == MnJMP ||
		IsJCCInstruction(mnemonic) || IsLOOPInstruction(mnemonic)
}

var REGISTERS8 = [16]string{
	"al", "cl", "dl", "bl", "spl", "bpl", "sil", "dil",
	"r8b", "r9b", "r10b", "r11b", "r12b", "r13b", "r14b", "r15b",
}
var REGISTERS16 = [16]string{
	"ax", "cx", "dx", "bx", "sp", "bp", "si", "di",
	"r8w", "r9w", "r10w", "r11w", "r12w", "r13w", "r14w", "r15w",
}
var REGISTERS32 = [16]string{
	"eax", "ecx", "edx", "ebx", "esp", "ebp", "esi", "edi",
	"r8d", "r9d", "r10d", "r11d", "r12d", "r13d", "r14d", "r15d",
}
var REGISTERS64 = [16]string{
	"rax", "rcx", "rdx", "rbx", "rsp", "rbp", "rsi", "rdi",
	"r8", "r9", "r10", "r11", "r12", "r13", "r14", "r15",
}

func Operand2Register(operand Operand) []string {
	if Is8Bit(operand) {
		return REGISTERS8[:]
	} else if Is16Bit(operand) {
		return REGISTERS16[:]
	} else if Is32Bit(operand) {
		return REGISTERS32[:]
	} else if Is64Bit(operand) {
		return REGISTERS64[:]
	}

	return nil
}

func HasModrm(opEnc OpEnc) bool {
	switch opEnc {
	case OpEncI:
		return false
	case OpEncA:
		return true
	case OpEncB:
		return true
	case OpEncD:
		return false
	case OpEncM:
		return true
	case OpEncO:
		return false
	case OpEncNP:
		return false
	case OpEncMI:
		return true
	case OpEncM1:
		return true
	case OpEncMC:
		return true
	case OpEncMR:
		return true
	case OpEncRM:
		return true
	case OpEncRMI:
		return true
	case OpEncMRI:
		return true
	case OpEncMRC:
		return true
	case OpEncOI:
		return false
	case OpEncFD:
		return false
	case OpEncTD:
		return false
	case OpEncS:
		return false
	default:
		return false
	}
}

type Prefix int

const (
	PrefixNONE Prefix = iota
	PrefixP66         // change the default operand size
	PrefixREXW        // use R8-R15 registers
	PrefixREX         // use 64 bit registers
)

type Mnemonic int

const (
	MnSETNE Mnemonic = iota
	MnFXCH
	MnFADD
	MnENDBR64
	MnENDBR32
	MnSHLD
	MnSHRD
	MnMOV
	MnCMOVE
	MnMOVSX
	MnMOVZX
	MnMOVAPS
	MnSCASQ
	MnLODSQ
	MnSTOSQ
	MnLEA
	MnADD
	MnADC
	MnSUB
	MnSBB
	MnMUL
	MnIML
	MnDIV
	MnIDIV
	MnINC
	MnDEC
	MnAND
	MnOR
	MnXOR
	MnNOT
	MnNEG
	MnCMP
	MnTEST
	MnSAL
	MnSHL
	MnSAR
	MnSHR
	MnRCL
	MnRCR
	MnROL
	MnROR
	MnJMP
	MnLOOP
	MnLOOPE
	MnLOOPNE
	MnJZ
	MnJNZ
	MnJP
	MnJO
	MnJNO
	MnJS
	MnJECXZ
	MnJNB
	MnJNBE
	MnJNG
	MnJNGE
	MnJNL
	MnJNLE
	MnJNS
	MnJNAE
	MnJNA
	MnJPO
	MnCALL
	MnRET
	MnPUSH
	MnPOP
	MnMOVSB
	MnMOVSW
	MnMOVSD
	MnMOVSQ
	MnCLD
	MnSTD
	MnLODSB
	MnLODSW
	MnLODSD
	MnSTOSB
	MnSTOSW
	MnSTOSD
	MnSCASB
	MnSCASW
	MnSCASD
	MnCMPSB
	MnCMPSW
	MnCMPSD
	MnCMPSQ
	MnIN
	MnOUT
	MnINSB
	MnINSW
	MnINSD
	MnOUTSB
	MnOUTSW
	MnOUTSD
	MnCBW
	MnCWD
	MnCWDE
	MnCDQ
	MnCDQE
	MnCQO
	MnINT21
	MNENTER
	MnLEAVE
	MnNOP
	MnUD2
	MnCPUID
	MnXCHG
	MnSTC
	MnCLC
	MnBSWAP
)

var SCALE_FACTOR = [4]int{1, 2, 4, 8}
