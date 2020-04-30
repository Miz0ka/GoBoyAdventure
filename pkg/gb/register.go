package gb

type register struct {
	value uint16
}

func (reg *register) getHi() uint8 {
	return uint8(reg.value >> 8)
}

func (reg *register) setHi(val uint8) {
	reg.value = reg.value ^ (uint16(reg.getHi()) << 8)
	reg.value = reg.value + (uint16(val) << 8)
}

func (reg *register) getLo() uint8 {
	return uint8(reg.value)
}

func (reg *register) setLo(val uint8) {
	reg.value = reg.value ^ (uint16(reg.getLo()))
	reg.value = reg.value | (uint16(val))
}

func (reg *register) getHiLo() uint16 {
	return reg.value
}

func (reg *register) setHiLo(val uint16) {
	reg.value = val
}
