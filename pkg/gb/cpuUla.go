package gb

import (
	byteUtils "github.com/Miz0ka/GoBoyAdventure/pkg/byteUtils"
)

// ============= ULA 8-Bit ===============

func (cpu *CPU) ulaAdd(reg1 uint8, reg2 uint8, ADC bool) uint8 {
	carry := byteUtils.BoolToByte(ADC)
	total := uint16(reg1) + uint16(reg2) + uint16(carry)

	cpu.setFlag(total == 0, ZF)
	cpu.setFlag(false, N)
	cpu.setFlag((reg1&0xF)+(reg2&0xF)+carry > 0xF, H)
	cpu.setFlag(total > 0xFF, CY)

	return uint8(total)
}

func (cpu *CPU) ulaSub(reg1 uint8, reg2 uint8, SBC bool) uint8 {
	carry := byteUtils.BoolToByte(SBC)
	total := int16(reg1) - int16(reg2) - int16(carry)

	cpu.setFlag(total == 0, ZF)
	cpu.setFlag(true, N)
	cpu.setFlag(int16(reg1&0xF)-int16(reg2&0xF)-int16(carry) < 0, H)
	cpu.setFlag(total < 0, CY)

	return uint8(total)
}

func (cpu *CPU) ulaAnd(reg1 uint8, reg2 uint8) uint8 {
	total := int16(reg1) & int16(reg2)

	cpu.setFlag(total == 0, ZF)
	cpu.setFlag(false, N)
	cpu.setFlag(true, H)
	cpu.setFlag(false, CY)

	return uint8(total)
}

func (cpu *CPU) ulaOr(reg1 uint8, reg2 uint8) uint8 {
	total := int16(reg1) | int16(reg2)

	cpu.setFlag(total == 0, ZF)
	cpu.setFlag(false, N)
	cpu.setFlag(false, H)
	cpu.setFlag(false, CY)

	return uint8(total)
}

func (cpu *CPU) ulaXor(reg1 uint8, reg2 uint8) uint8 {
	total := int16(reg1) ^ int16(reg2)

	cpu.setFlag(total == 0, ZF)
	cpu.setFlag(false, N)
	cpu.setFlag(false, H)
	cpu.setFlag(false, CY)

	return uint8(total)
}

func (cpu *CPU) ulaCP(reg1, reg2 uint8) {
	cpu.ulaSub(reg1, reg2, false)
}

func (cpu *CPU) ulaInc(reg uint8) uint8 {
	total := uint16(reg) + 1

	cpu.setFlag(total == 0, ZF)
	cpu.setFlag(false, N)
	cpu.setFlag((reg&0xF)+1 > 0xF, H)

	return uint8(total)
}

func (cpu *CPU) ulaDec(reg uint8) uint8 {
	total := int16(reg) - 1

	cpu.setFlag(total == 0, ZF)
	cpu.setFlag(true, N)
	cpu.setFlag((reg&0xF)+1 < 0, H)

	return uint8(total)
}

// ============= ULA 16-Bit ===============

func (cpu *CPU) ulaAdd16(reg1 uint16, reg2 uint16) uint16 {
	total := uint32(reg1) + uint32(reg2)

	cpu.setFlag(total == 0, ZF)
	cpu.setFlag(false, N)
	cpu.setFlag(uint16(reg1&0xFFF)+uint16(reg2&0xFFF) > 0xFFF, H)
	cpu.setFlag(total > 0xFFFF, CY)

	return uint16(total)
}

func (cpu *CPU) ulaAdd16Signed(reg1 uint16, reg2 int8) uint16 {
	total := int32(reg1) + int32(reg2)

	tmpflag := reg1 ^ uint16(reg2) ^ uint16(total)
	cpu.setFlag(false, ZF)
	cpu.setFlag(false, N)
	cpu.setFlag((tmpflag&0x10) == 0x10, H)
	cpu.setFlag((tmpflag&0x100) == 0x100, CY)

	return uint16(total)
}

func (cpu *CPU) ulaInc16(reg uint16) uint16 {
	total := uint32(reg) + 1
	return uint16(total)
}

func (cpu *CPU) ulaDec16(reg uint16) uint16 {
	total := int32(reg) - 1
	return uint16(total)
}

// ============= Miscellaneous ===============

func (cpu *CPU) ulaSwap(reg uint8) uint8 {
	val := byteUtils.Swap(byte(reg))

	cpu.setFlag(val == 0, ZF)

	return val
}

func (cpu *CPU) ulaDAA(reg uint8) uint8 {
	val := uint16(0)
	if cpu.isH() || ((reg&0xF) > 9 && !cpu.isN()) { // Hi
		val |= 0x06
	}
	if cpu.isCY() || (reg > 0x99 && !cpu.isN()) { // Lo
		val |= 0x60
		cpu.setFlag(true, CY)
	}

	if cpu.isN() {
		val -= uint16(reg)
	} else {
		val += uint16(reg)
	}

	cpu.setFlag(val == 0, ZF)
	cpu.setFlag(false, H)
	if !cpu.isCY() {
		cpu.setFlag(val > 0xFFFF, CY)
	}

	return uint8(val)
}

func (cpu *CPU) ulaCPL(reg uint8) uint8 {
	reg = reg ^ 0xFF

	cpu.setFlag(true, N)
	cpu.setFlag(true, H)

	return reg
}

func (cpu *CPU) ulaCCF() {
	cpu.setFlag(false, N)
	cpu.setFlag(false, H)
	cpu.setFlag(!cpu.isCY(), CY)
}

func (cpu *CPU) ulaSCF() {
	cpu.setFlag(false, N)
	cpu.setFlag(false, H)
	cpu.setFlag(true, CY)
}

// ============= Rotates & Shifts ===============

func (cpu *CPU) ulaRLC(reg uint8) uint8 {
	val := reg<<1 + (reg & 0x80 >> 7)

	cpu.setFlag(val == 0, N)
	cpu.setFlag(false, N)
	cpu.setFlag(false, H)
	cpu.setFlag(reg&0x80 == 0x80, CY)

	return val
}

func (cpu *CPU) ulaRL(reg uint8) uint8 {
	val := reg<<1 + byteUtils.BoolToByte(cpu.isCY())

	cpu.setFlag(val == 0, N)
	cpu.setFlag(false, N)
	cpu.setFlag(false, H)
	cpu.setFlag(reg&0x80 == 0x80, CY)

	return val
}

func (cpu *CPU) ulaRRC(reg uint8) uint8 {
	val := reg>>1 + (reg & 0x01 << 7)

	cpu.setFlag(val == 0, N)
	cpu.setFlag(false, N)
	cpu.setFlag(false, H)
	cpu.setFlag(reg&0x01 == 0x01, CY)

	return val
}

func (cpu *CPU) ulaRR(reg uint8) uint8 {
	val := reg>>1 + (byteUtils.BoolToByte(cpu.isCY()) << 7)

	cpu.setFlag(val == 0, N)
	cpu.setFlag(false, N)
	cpu.setFlag(false, H)
	cpu.setFlag(reg&0x01 == 0x01, CY)

	return val
}

func (cpu *CPU) ulaSLA(reg uint8) uint8 {
	val := reg << 1

	cpu.setFlag(val == 0, N)
	cpu.setFlag(false, N)
	cpu.setFlag(false, H)
	cpu.setFlag(reg&0x80 == 0x80, CY)

	return val
}

func (cpu *CPU) ulaSRA(reg uint8) uint8 {
	val := reg>>1 | (reg & 0x80)

	cpu.setFlag(val == 0, N)
	cpu.setFlag(false, N)
	cpu.setFlag(false, H)
	cpu.setFlag(reg&0x01 == 0x01, CY)

	return val
}

func (cpu *CPU) ulaSRL(reg uint8) uint8 {
	val := reg >> 1

	cpu.setFlag(val == 0, N)
	cpu.setFlag(false, N)
	cpu.setFlag(false, H)
	cpu.setFlag(reg&0x01 == 0x01, CY)

	return val
}

// ============= Bit Opcodes ===============

func (cpu *CPU) ulaBit(b byte, reg uint8) {
	cpu.setFlag((reg>>b)&0x1 == 0x1, N)
	cpu.setFlag(false, N)
	cpu.setFlag(true, H)
}

func (cpu *CPU) ulaSet(b byte, reg uint8) uint8 {
	return reg | (0x1 << b)
}

func (cpu *CPU) ulaRes(b byte, reg uint8) uint8 {
	return reg ^ (reg & (0x1 << b))
}
