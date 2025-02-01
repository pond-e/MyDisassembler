package main

import (
	"strconv"
)

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
	// fmt.Printf("modrm.modByte: %x\n", modrm.modByte)
	// fmt.Printf("operand: %x\n", operand)
	var addrBaseReg string
	var addNum byte = 0
	if modrm.rex.rexB {
		addNum = 8
	}
	if modrm.modByte == 3 {
		if operand == OpXm128 {
			addrBaseReg = "xmm" + strconv.Itoa(int(modrm.rmByte+addNum))
		} else {
			addNum := 0
			if modrm.rex.rexB {
				addNum = 8
			}
			addrBaseReg = Operand2Register(operand)[modrm.rmByte+byte(addNum)]
		}
	} else {
		if operand == OpXm128 {
			addrBaseReg = "xmm" + strconv.Itoa(int(modrm.rmByte+addNum))
		} else {
			// addrBaseReg = REGISTERS64[modrm.rmByte+addNum]
			addrBaseReg = Operand2Register(operand)[modrm.rmByte+byte(addNum)]
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
