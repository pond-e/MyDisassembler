package main

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
