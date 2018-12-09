package machine

type Instruction uint32

func (i Instruction) Opcode() int {
	return int(i & 0x3F)
}

func (i Instruction) ABC() (a, b, c int) {
	a = int(i >> 6 & 0xFF)
	c = int(i >> 14 & 0x1FF)
	b = int(i >> 23 & 0x1FF)
	return a, b, c
}

func (i Instruction) ABx() (a, bx int) {
	a = int(i >> 6 & 0xFF)
	bx = int(i >> 14)
	return a, bx
}

func (i Instruction) ASbx() (a, sbx int) {
	a, bx := i.ABx()
	return a, bx - (1 << 18 - 1)
}

func (i Instruction) Ax() int {
	return int(i >> 6)
}
