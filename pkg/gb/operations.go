package gb

var OpcodeCycles = []int{
	1, 3, 2, 2, 1, 1, 2, 1, 5, 2, 2, 2, 1, 1, 2, 1, // 0
	0, 3, 2, 2, 1, 1, 2, 1, 3, 2, 2, 2, 1, 1, 2, 1, // 1
	2, 3, 2, 2, 1, 1, 2, 1, 2, 2, 2, 2, 1, 1, 2, 1, // 2
	2, 3, 2, 2, 3, 3, 3, 1, 2, 2, 2, 2, 1, 1, 2, 1, // 3
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 4
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 5
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 6
	2, 2, 2, 2, 2, 2, 0, 2, 1, 1, 1, 1, 1, 1, 2, 1, // 7
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 8
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 9
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // a
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // b
	2, 3, 3, 4, 3, 4, 2, 4, 2, 4, 3, 0, 3, 6, 2, 4, // c
	2, 3, 3, 0, 3, 4, 2, 4, 2, 4, 3, 0, 3, 0, 2, 4, // d
	3, 3, 2, 0, 0, 4, 2, 4, 4, 1, 4, 0, 0, 0, 2, 4, // e
	3, 3, 2, 1, 0, 4, 2, 4, 3, 2, 4, 1, 0, 0, 2, 4, // f
} //0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f

func (gb *gameboy) executeNextOpcode() {
	opCode := gb.readOPArg8bit()
	gb.cpu.cycles = OpcodeCBCycles[opCode] * 4
	gb.opCode[opCode]()
}

func (gb *gameboy) readOPArg8bit() uint8 {
	val := gb.memory.read8bit(gb.cpu.registers.pc)
	gb.cpu.registers.pc++
	return val
}

func (gb *gameboy) readOPArg16bit() uint16 {
	val := gb.memory.read16bit(gb.cpu.registers.pc)
	gb.cpu.registers.pc++
	gb.cpu.registers.pc++
	return val
}

func (gb *gameboy) genereOPDic() [0x100]func() {
	dicOP := [0x100]func(){
		//--------------------------
		//	 8-Bit Loads
		//--------------------------

		// LD r, n
		0x06: func() { // LD B, d8
			val := gb.readOPArg8bit()
			gb.cpu.registers.bc.setHi(val)
		},
		0x0E: func() { // LD C, d8
			val := gb.readOPArg8bit()
			gb.cpu.registers.bc.setLo(val)
		},
		0x16: func() { // LD D, d8
			val := gb.readOPArg8bit()
			gb.cpu.registers.de.setHi(val)
		},
		0x1E: func() { // LD E, d8
			val := gb.readOPArg8bit()
			gb.cpu.registers.de.setLo(val)
		},
		0x26: func() { // LD H, d8
			val := gb.readOPArg8bit()
			gb.cpu.registers.hl.setHi(val)
		},
		0x2E: func() { // LD L, d8
			val := gb.readOPArg8bit()
			gb.cpu.registers.hl.setLo(val)
		},

		// LD r1 r2
		0x7F: func() { // LD A, A
			val := gb.cpu.registers.af.getHi()
			gb.cpu.registers.bc.setHi(val)
		},
		0x78: func() { // LD A, B
			val := gb.cpu.registers.bc.getHi()
			gb.cpu.registers.af.setHi(val)
		},
		0x79: func() { // LD A, C
			val := gb.cpu.registers.bc.getLo()
			gb.cpu.registers.af.setHi(val)
		},
		0x7A: func() { // LD A, D
			val := gb.cpu.registers.de.getHi()
			gb.cpu.registers.af.setHi(val)
		},
		0x7B: func() { // LD A, E
			val := gb.cpu.registers.de.getLo()
			gb.cpu.registers.af.setHi(val)
		},
		0x7C: func() { // LD A, H
			val := gb.cpu.registers.hl.getHi()
			gb.cpu.registers.af.setHi(val)
		},
		0x7D: func() { // LD A, L
			val := gb.cpu.registers.hl.getLo()
			gb.cpu.registers.af.setHi(val)
		},
		0x7E: func() { // LD A, (HL)
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.registers.af.setHi(val)
		},
		0x40: func() { // LD B, B
			val := gb.cpu.registers.bc.getHi()
			gb.cpu.registers.bc.setHi(val)
		},
		0x41: func() { // LD B, C
			val := gb.cpu.registers.bc.getLo()
			gb.cpu.registers.bc.setHi(val)
		},
		0x42: func() { // LD B, D
			val := gb.cpu.registers.de.getHi()
			gb.cpu.registers.bc.setHi(val)
		},
		0x43: func() { // LD B, E
			val := gb.cpu.registers.de.getLo()
			gb.cpu.registers.bc.setHi(val)
		},
		0x44: func() { // LD B, H
			val := gb.cpu.registers.hl.getHi()
			gb.cpu.registers.bc.setHi(val)
		},
		0x45: func() { // LD B, L
			val := gb.cpu.registers.hl.getLo()
			gb.cpu.registers.bc.setHi(val)
		},
		0x46: func() { // LD B, (HL)
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.registers.bc.setHi(val)
		},
		0x47: func() { // LD B, A
			val := gb.cpu.registers.af.getHi()
			gb.cpu.registers.bc.setHi(val)
		},
		0x48: func() { // LD C, B
			val := gb.cpu.registers.bc.getHi()
			gb.cpu.registers.bc.setLo(val)
		},
		0x49: func() { // LD C, C
			val := gb.cpu.registers.bc.getLo()
			gb.cpu.registers.bc.setLo(val)
		},
		0x4A: func() { // LD C, D
			val := gb.cpu.registers.de.getHi()
			gb.cpu.registers.bc.setLo(val)
		},
		0x4B: func() { // LD C, E
			val := gb.cpu.registers.de.getLo()
			gb.cpu.registers.bc.setLo(val)
		},
		0x4C: func() { // LD C, H
			val := gb.cpu.registers.hl.getHi()
			gb.cpu.registers.bc.setLo(val)
		},
		0x4D: func() { // LD C, L
			val := gb.cpu.registers.hl.getLo()
			gb.cpu.registers.bc.setLo(val)
		},
		0x4E: func() { // LD C, (HL)
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.registers.bc.setLo(val)
		},
		0x4F: func() { // LD C, A
			val := gb.cpu.registers.af.getHi()
			gb.cpu.registers.bc.setLo(val)
		},
		0x50: func() { // LD D, B
			val := gb.cpu.registers.bc.getHi()
			gb.cpu.registers.de.setHi(val)
		},
		0x51: func() { // LD D, C
			val := gb.cpu.registers.bc.getLo()
			gb.cpu.registers.de.setHi(val)
		},
		0x52: func() { // LD D, D
			val := gb.cpu.registers.de.getHi()
			gb.cpu.registers.de.setHi(val)
		},
		0x53: func() { // LD D, E
			val := gb.cpu.registers.de.getLo()
			gb.cpu.registers.de.setHi(val)
		},
		0x54: func() { // LD D, H
			val := gb.cpu.registers.hl.getHi()
			gb.cpu.registers.de.setHi(val)
		},
		0x55: func() { // LD D, L
			val := gb.cpu.registers.hl.getLo()
			gb.cpu.registers.de.setHi(val)
		},
		0x56: func() { // LD D, (HL)
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.registers.de.setHi(val)
		},
		0x57: func() { // LD D, A
			val := gb.cpu.registers.af.getHi()
			gb.cpu.registers.de.setHi(val)
		},
		0x58: func() { // LD E, B
			val := gb.cpu.registers.bc.getHi()
			gb.cpu.registers.de.setLo(val)
		},
		0x59: func() { // LD E, C
			val := gb.cpu.registers.bc.getLo()
			gb.cpu.registers.de.setLo(val)
		},
		0x5A: func() { // LD E, D
			val := gb.cpu.registers.de.getHi()
			gb.cpu.registers.de.setLo(val)
		},
		0x5B: func() { // LD E, E
			val := gb.cpu.registers.de.getLo()
			gb.cpu.registers.de.setLo(val)
		},
		0x5C: func() { // LD E, H
			val := gb.cpu.registers.hl.getHi()
			gb.cpu.registers.de.setLo(val)
		},
		0x5D: func() { // LD E, L
			val := gb.cpu.registers.hl.getLo()
			gb.cpu.registers.de.setLo(val)
		},
		0x5E: func() { // LD E, (HL)
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.registers.de.setLo(val)
		},
		0x5F: func() { // LD E, A
			val := gb.cpu.registers.af.getHi()
			gb.cpu.registers.de.setLo(val)
		},
		0x60: func() { // LD H, B
			val := gb.cpu.registers.bc.getHi()
			gb.cpu.registers.hl.setHi(val)
		},
		0x61: func() { // LD H, C
			val := gb.cpu.registers.bc.getLo()
			gb.cpu.registers.hl.setHi(val)
		},
		0x62: func() { // LD H, D
			val := gb.cpu.registers.de.getHi()
			gb.cpu.registers.hl.setHi(val)
		},
		0x63: func() { // LD H, E
			val := gb.cpu.registers.de.getLo()
			gb.cpu.registers.hl.setHi(val)
		},
		0x64: func() { // LD H, H
			val := gb.cpu.registers.hl.getHi()
			gb.cpu.registers.hl.setHi(val)
		},
		0x65: func() { // LD H, L
			val := gb.cpu.registers.hl.getLo()
			gb.cpu.registers.hl.setHi(val)
		},
		0x66: func() { // LD H, (HL)
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.registers.hl.setHi(val)
		},
		0x67: func() { // LD H, A
			val := gb.cpu.registers.af.getHi()
			gb.cpu.registers.hl.setHi(val)
		},
		0x68: func() { // LD L, B
			val := gb.cpu.registers.bc.getHi()
			gb.cpu.registers.hl.setLo(val)
		},
		0x69: func() { // LD L, C
			val := gb.cpu.registers.bc.getLo()
			gb.cpu.registers.hl.setLo(val)
		},
		0x6A: func() { // LD L, D
			val := gb.cpu.registers.de.getHi()
			gb.cpu.registers.hl.setLo(val)
		},
		0x6B: func() { // LD L, E
			val := gb.cpu.registers.de.getLo()
			gb.cpu.registers.hl.setLo(val)
		},
		0x6C: func() { // LD L, H
			val := gb.cpu.registers.hl.getHi()
			gb.cpu.registers.hl.setLo(val)
		},
		0x6D: func() { // LD L, L
			val := gb.cpu.registers.hl.getLo()
			gb.cpu.registers.hl.setLo(val)
		},
		0x6E: func() { // LD L, (HL)
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.registers.hl.setLo(val)
		},
		0x6F: func() { // LD L, A
			val := gb.cpu.registers.af.getHi()
			gb.cpu.registers.hl.setLo(val)
		},
		0x70: func() { // LD (HL), B
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.cpu.registers.bc.getHi()
			gb.memory.write8bit(addr, val)
		},
		0x71: func() { // LD (HL), C
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.cpu.registers.bc.getLo()
			gb.memory.write8bit(addr, val)
		},
		0x72: func() { // LD (HL), D
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.cpu.registers.de.getHi()
			gb.memory.write8bit(addr, val)
		},
		0x73: func() { // LD (HL), E
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.cpu.registers.de.getLo()
			gb.memory.write8bit(addr, val)
		},
		0x74: func() { // LD (HL), H
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.cpu.registers.hl.getHi()
			gb.memory.write8bit(addr, val)
		},
		0x75: func() { // LD (HL), L
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.cpu.registers.hl.getLo()
			gb.memory.write8bit(addr, val)
		},
		0x36: func() { /// LD (HL), d8
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.readOPArg8bit()
			gb.memory.write8bit(addr, val)
		},

		// LD A,n
		0x0A: func() { // LD A, (BC)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.registers.af.setHi(val)
		},
		0x1A: func() { // LD A, (DE)
			addr := gb.cpu.registers.de.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.registers.af.setHi(val)
		},
		0xFA: func() { /// LD A, (nn)
			addr := gb.readOPArg16bit()
			val := gb.memory.read8bit(addr)
			gb.cpu.registers.af.setHi(val)
		},
		0x3E: func() { /// LD A, d8
			val := gb.readOPArg8bit()
			gb.cpu.registers.af.setHi(val)
		},

		//LD n, A

		0x02: func() { // LD (BC),A
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.cpu.registers.af.getHi()
			gb.memory.write8bit(addr, val)
		},
		0x12: func() { // LD (DE),A
			addr := gb.cpu.registers.de.getHiLo()
			val := gb.cpu.registers.af.getHi()
			gb.memory.write8bit(addr, val)
		},
		0x77: func() { // LD (HL),A
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.cpu.registers.af.getHi()
			gb.memory.write8bit(addr, val)
		},
		0xEA: func() { // LD (nn),A
			addr := gb.readOPArg16bit()
			val := gb.cpu.registers.af.getHi()
			gb.memory.write8bit(addr, val)
		},

		// LD A,(C)
		0xF2: func() { /// LD A, (C)
			addr := uint16(0xFF00) + uint16(gb.cpu.registers.bc.getLo())
			val := gb.memory.read8bit(addr)
			gb.cpu.registers.af.setHi(val)
		},

		// LD (C), A
		0xE2: func() { /// LD (C), A
			addr := uint16(0xFF00) + uint16(gb.cpu.registers.bc.getLo())
			val := gb.cpu.registers.af.getHi()
			gb.memory.write8bit(addr, val)
		},

		// LDD A, (HL)
		0x3A: func() { /// LD A, (HL-)
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.registers.af.setHi(val)
			gb.cpu.registers.hl.setHiLo(addr - 1)
		},

		// LDD (HL), A
		0x32: func() { // LD (HL-), A
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.cpu.registers.af.getHi()
			gb.memory.write8bit(addr, val)
			gb.cpu.registers.hl.setHiLo(addr - 1)
		},

		// LDI A,(HL)
		0x2A: func() { // LD A, (HL+)
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.registers.af.setHi(val)
			gb.cpu.registers.hl.setHiLo(addr + 1)
		},

		// LDI (HL), A
		0x22: func() { // LD (HL+), A
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.cpu.registers.af.getHi()
			gb.memory.write8bit(addr, val)
			gb.cpu.registers.hl.setHiLo(addr + 1)
		},

		// LDH (n), A
		0xE0: func() { // LDH (n),A
			addr := uint16(gb.readOPArg8bit()) + uint16(0xFF00)
			val := gb.cpu.registers.af.getHi()
			gb.memory.write8bit(addr, val)
		},

		// LDH A, (n)
		0xF0: func() { // LDH (n),A
			addr := uint16(gb.readOPArg8bit()) + uint16(0xFF00)
			val := gb.memory.read8bit(addr)
			gb.cpu.registers.af.setHi(val)
		},

		//--------------------------
		//	 16-Bit Loads
		//--------------------------

		// LD r, nn
		0x01: func() { // LD BC X
			val := gb.readOPArg16bit()
			gb.cpu.registers.bc.setHiLo(val)
		},
		0x11: func() { // LD DE X
			val := gb.readOPArg16bit()
			gb.cpu.registers.de.setHiLo(val)
		},
		0x21: func() { // LD HL X
			val := gb.readOPArg16bit()
			gb.cpu.registers.hl.setHiLo(val)
		},
		0x31: func() { // LD sp X
			val := gb.readOPArg16bit()
			gb.cpu.setSP(val)
		},

		// LD SP, HL
		0xF9: func() { // LD SP, HL
			val := gb.cpu.registers.hl.getHiLo()
			gb.cpu.setSP(val)
		},

		// LDHL SP, n
		0xF8: func() { // LDHL SP, n
			nVal := gb.readOPArg8bit()
			spVal := gb.cpu.getSP()
			val := gb.cpu.ulaAdd16Signed(spVal, int8(nVal))
			gb.cpu.registers.hl.setHiLo(val)
		},

		// LD (nn), SP
		0x08: func() { // LD (a16) SP
			addr := gb.readOPArg16bit()
			gb.memory.write16bit(addr, gb.cpu.getSP())
		},

		// PUSH rr
		0xF5: func() { // PUSH AF
			val := gb.cpu.registers.af.getHiLo()
			gb.push(val)
		},
		0xC5: func() { // PUSH BC
			val := gb.cpu.registers.bc.getHiLo()
			gb.push(val)
		},
		0xD5: func() { // PUSH DE
			val := gb.cpu.registers.de.getHiLo()
			gb.push(val)
		},
		0xE5: func() { // PUSH HL
			val := gb.cpu.registers.hl.getHiLo()
			gb.push(val)
		},

		// POP rr
		0xF1: func() { // POP AF
			val := gb.pop()
			gb.cpu.registers.af.setHiLo(val)
		},
		0xC1: func() { // POP BC
			val := gb.pop()
			gb.cpu.registers.bc.setHiLo(val)
		},
		0xD1: func() { // POP DE
			val := gb.pop()
			gb.cpu.registers.de.setHiLo(val)
		},
		0xE1: func() { // POP HL
			val := gb.pop()
			gb.cpu.registers.hl.setHiLo(val)
		},

		//--------------------------
		//	 8-Bit ALU
		//--------------------------

		// ADD A, n
		0x87: func() { // ADD A, A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaAdd(val, val, false)
			gb.cpu.registers.af.setHi(total)
		},
		0x80: func() { // ADD A, B
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaAdd(val1, val2, false)
			gb.cpu.registers.af.setHi(total)
		},
		0x81: func() { // ADD A, C
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaAdd(val1, val2, false)
			gb.cpu.registers.af.setHi(total)
		},
		0x82: func() { // ADD A, D
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaAdd(val1, val2, false)
			gb.cpu.registers.af.setHi(total)
		},
		0x83: func() { // ADD A, E
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaAdd(val1, val2, false)
			gb.cpu.registers.af.setHi(total)
		},
		0x84: func() { // ADD A, H
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaAdd(val1, val2, false)
			gb.cpu.registers.af.setHi(total)
		},
		0x85: func() { // ADD A, L
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaAdd(val1, val2, false)
			gb.cpu.registers.af.setHi(total)
		},
		0x86: func() { // ADD A, (HL)
			val1 := gb.cpu.registers.af.getHi()
			addr := gb.cpu.registers.hl.getHiLo()
			val2 := gb.memory.read8bit(addr)
			total := gb.cpu.ulaAdd(val1, val2, false)
			gb.cpu.registers.af.setHi(total)
		},
		0xC6: func() { // ADD A, n
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.readOPArg8bit()
			total := gb.cpu.ulaAdd(val1, val2, false)
			gb.cpu.registers.af.setHi(total)
		},

		// ADC A, n
		0x8F: func() { // ADC A, A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaAdd(val, val, true)
			gb.cpu.registers.af.setHi(total)
		},
		0x88: func() { // ADC A, B
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaAdd(val1, val2, true)
			gb.cpu.registers.af.setHi(total)
		},
		0x89: func() { // ADC A, C
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaAdd(val1, val2, true)
			gb.cpu.registers.af.setHi(total)
		},
		0x8A: func() { // ADC A, D
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaAdd(val1, val2, true)
			gb.cpu.registers.af.setHi(total)
		},
		0x8B: func() { // ADC A, E
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaAdd(val1, val2, true)
			gb.cpu.registers.af.setHi(total)
		},
		0x8C: func() { // ADC A, H
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaAdd(val1, val2, true)
			gb.cpu.registers.af.setHi(total)
		},
		0x8D: func() { // ADC A, L
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaAdd(val1, val2, true)
			gb.cpu.registers.af.setHi(total)
		},
		0x8E: func() { // ADC A, (HL)
			val1 := gb.cpu.registers.af.getHi()
			addr := gb.cpu.registers.hl.getHiLo()
			val2 := gb.memory.read8bit(addr)
			total := gb.cpu.ulaAdd(val1, val2, true)
			gb.cpu.registers.af.setHi(total)
		},
		0xCE: func() { // ADC A, n
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.readOPArg8bit()
			total := gb.cpu.ulaAdd(val1, val2, true)
			gb.cpu.registers.af.setHi(total)
		},

		// SUB A, n
		0x97: func() { // SUB A, A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaSub(val, val, false)
			gb.cpu.registers.af.setHi(total)
		},
		0x90: func() { // SUB A, B
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaSub(val1, val2, false)
			gb.cpu.registers.af.setHi(total)
		},
		0x91: func() { // SUB A, C
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaSub(val1, val2, false)
			gb.cpu.registers.af.setHi(total)
		},
		0x92: func() { // SUB A, D
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaSub(val1, val2, false)
			gb.cpu.registers.af.setHi(total)
		},
		0x93: func() { // SUB A, E
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaSub(val1, val2, false)
			gb.cpu.registers.af.setHi(total)
		},
		0x94: func() { // SUB A, H
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaSub(val1, val2, false)
			gb.cpu.registers.af.setHi(total)
		},
		0x95: func() { // SUB A, L
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaSub(val1, val2, false)
			gb.cpu.registers.af.setHi(total)
		},
		0x96: func() { // SUB A, (HL)
			val1 := gb.cpu.registers.af.getHi()
			addr := gb.cpu.registers.hl.getHiLo()
			val2 := gb.memory.read8bit(addr)
			total := gb.cpu.ulaSub(val1, val2, false)
			gb.cpu.registers.af.setHi(total)
		},
		0xD6: func() { // SUB A, n
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.readOPArg8bit()
			total := gb.cpu.ulaSub(val1, val2, false)
			gb.cpu.registers.af.setHi(total)
		},

		// SBC A, n
		0x9F: func() { // SBC A, A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaSub(val, val, true)
			gb.cpu.registers.af.setHi(total)
		},
		0x98: func() { // SBC A, B
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaSub(val1, val2, true)
			gb.cpu.registers.af.setHi(total)
		},
		0x99: func() { // SBC A, C
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaSub(val1, val2, true)
			gb.cpu.registers.af.setHi(total)
		},
		0x9A: func() { // SBC A, D
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaSub(val1, val2, true)
			gb.cpu.registers.af.setHi(total)
		},
		0x9B: func() { // SBC A, E
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaSub(val1, val2, true)
			gb.cpu.registers.af.setHi(total)
		},
		0x9C: func() { // SBC A, H
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaSub(val1, val2, true)
			gb.cpu.registers.af.setHi(total)
		},
		0x9D: func() { // SBC A, L
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaSub(val1, val2, true)
			gb.cpu.registers.af.setHi(total)
		},
		0x9E: func() { // SBC A, (HL)
			val1 := gb.cpu.registers.af.getHi()
			addr := gb.cpu.registers.hl.getHiLo()
			val2 := gb.memory.read8bit(addr)
			total := gb.cpu.ulaSub(val1, val2, true)
			gb.cpu.registers.af.setHi(total)
		},
		/*0xCE : func () { // SBC A, n
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.readOPArg8bit()
			total := gb.cpu.ulaSub(val1, val2, true)
			gb.cpu.registers.af.setHi(total)
		},*/

		// AND A, n
		0xA7: func() { // AND A, A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaAnd(val, val)
			gb.cpu.registers.af.setHi(total)
		},
		0xA0: func() { // AND A, B
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaAnd(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xA1: func() { // AND A, C
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaAnd(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xA2: func() { // AND A, D
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaAnd(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xA3: func() { // AND A, E
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaAnd(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xA4: func() { // AND A, H
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaAnd(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xA5: func() { // AND A, L
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaAnd(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xA6: func() { // AND A, (HL)
			val1 := gb.cpu.registers.af.getHi()
			addr := gb.cpu.registers.hl.getHiLo()
			val2 := gb.memory.read8bit(addr)
			total := gb.cpu.ulaAnd(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xE6: func() { // AND A, n
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.readOPArg8bit()
			total := gb.cpu.ulaAnd(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},

		// OR A, n
		0xB7: func() { // OR A, A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaOr(val, val)
			gb.cpu.registers.af.setHi(total)
		},
		0xB0: func() { // OR A, B
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaOr(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xB1: func() { // OR A, C
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaOr(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xB2: func() { // OR A, D
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaOr(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xB3: func() { // OR A, E
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaOr(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xB4: func() { // OR A, H
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaOr(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xB5: func() { // OR A, L
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaOr(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xB6: func() { // OR A, (HL)
			val1 := gb.cpu.registers.af.getHi()
			addr := gb.cpu.registers.hl.getHiLo()
			val2 := gb.memory.read8bit(addr)
			total := gb.cpu.ulaOr(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xF6: func() { // OR A, n
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.readOPArg8bit()
			total := gb.cpu.ulaOr(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},

		// XOR A, n
		0xAF: func() { // XOR A, A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaXor(val, val)
			gb.cpu.registers.af.setHi(total)
		},
		0xA8: func() { // XOR A, B
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaXor(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xA9: func() { // XOR A, C
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaXor(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xAA: func() { // XOR A, D
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaXor(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xAB: func() { // XOR A, E
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaXor(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xAC: func() { // XOR A, H
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaXor(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xAD: func() { // XOR A, L
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaXor(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xAE: func() { // XOR A, (HL)
			val1 := gb.cpu.registers.af.getHi()
			addr := gb.cpu.registers.hl.getHiLo()
			val2 := gb.memory.read8bit(addr)
			total := gb.cpu.ulaXor(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},
		0xEE: func() { // XOR A, n
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.readOPArg8bit()
			total := gb.cpu.ulaXor(val1, val2)
			gb.cpu.registers.af.setHi(total)
		},

		// CP A, n
		0xBF: func() { // CP A, A
			val := gb.cpu.registers.af.getHi()
			gb.cpu.ulaCP(val, val)
		},
		0xB8: func() { // CP A, B
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.bc.getHi()
			gb.cpu.ulaCP(val1, val2)
		},
		0xB9: func() { // CP A, C
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.bc.getLo()
			gb.cpu.ulaCP(val1, val2)
		},
		0xBA: func() { // CP A, D
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.de.getHi()
			gb.cpu.ulaCP(val1, val2)
		},
		0xBB: func() { // CP A, E
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.de.getLo()
			gb.cpu.ulaCP(val1, val2)
		},
		0xBC: func() { // CP A, H
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.hl.getHi()
			gb.cpu.ulaCP(val1, val2)
		},
		0xBD: func() { // CP A, L
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.cpu.registers.hl.getLo()
			gb.cpu.ulaCP(val1, val2)
		},
		0xBE: func() { // CP A, (HL)
			val1 := gb.cpu.registers.af.getHi()
			addr := gb.cpu.registers.hl.getHiLo()
			val2 := gb.memory.read8bit(addr)
			gb.cpu.ulaCP(val1, val2)
		},
		0xFE: func() { // CP A, n
			val1 := gb.cpu.registers.af.getHi()
			val2 := gb.readOPArg8bit()
			gb.cpu.ulaCP(val1, val2)
		},

		// INC n
		0x3C: func() { // INC A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaInc(val)
			gb.cpu.registers.af.setHi(total)
		},
		0x04: func() { // INC B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaInc(val)
			gb.cpu.registers.bc.setHi(total)
		},
		0x0C: func() { // INC C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaInc(val)
			gb.cpu.registers.bc.setLo(total)
		},
		0x14: func() { // INC D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaInc(val)
			gb.cpu.registers.de.setHi(total)
		},
		0x1C: func() { // INC E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaInc(val)
			gb.cpu.registers.de.setLo(total)
		},
		0x24: func() { // INC H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaInc(val)
			gb.cpu.registers.hl.setHi(total)
		},
		0x2C: func() { // INC L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaInc(val)
			gb.cpu.registers.hl.setLo(total)
		},
		0x34: func() { // INC (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaInc(val)
			gb.memory.write8bit(addr, total)
		},

		// DEC r
		0x3D: func() { // DEC A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaDec(val)
			gb.cpu.registers.af.setHi(total)
		},
		0x05: func() { // DEC B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaDec(val)
			gb.cpu.registers.bc.setHi(total)
		},
		0x0D: func() { // DEC C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaDec(val)
			gb.cpu.registers.bc.setLo(total)
		},
		0x15: func() { // DEC D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaDec(val)
			gb.cpu.registers.de.setHi(total)
		},
		0x1D: func() { // DEC E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaDec(val)
			gb.cpu.registers.de.setLo(total)
		},
		0x25: func() { // DEC H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaDec(val)
			gb.cpu.registers.hl.setHi(total)
		},
		0x2D: func() { // DEC L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaDec(val)
			gb.cpu.registers.hl.setLo(total)
		},
		0x35: func() { // DEC (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaDec(val)
			gb.memory.write8bit(addr, total)
		},

		//--------------------------
		//	 16-Bit ALU
		//--------------------------

		// ADD HL, rr
		0x09: func() { // ADD HL, BC
			val1 := gb.cpu.registers.hl.getHiLo()
			val2 := gb.cpu.registers.bc.getHiLo()
			total := gb.cpu.ulaAdd16(val1, val2)
			gb.cpu.registers.hl.setHiLo(total)
		},
		0x19: func() { // ADD HL, DE
			val1 := gb.cpu.registers.hl.getHiLo()
			val2 := gb.cpu.registers.de.getHiLo()
			total := gb.cpu.ulaAdd16(val1, val2)
			gb.cpu.registers.hl.setHiLo(total)
		},
		0x29: func() { // ADD HL, HL
			val := gb.cpu.registers.hl.getHiLo()
			total := gb.cpu.ulaAdd16(val, val)
			gb.cpu.registers.hl.setHiLo(total)
		},
		0x39: func() { // ADD HL, SP
			val1 := gb.cpu.registers.hl.getHiLo()
			val2 := gb.cpu.getSP()
			total := gb.cpu.ulaAdd16(val1, val2)
			gb.cpu.registers.af.setHiLo(total)
		},

		//ADD SP, n
		0xE8: func() { // ADD SP, n
			val1 := gb.cpu.getSP()
			val2 := gb.readOPArg8bit()
			total := gb.cpu.ulaAdd16Signed(val1, int8(val2))
			gb.cpu.setSP(total)
		},

		// INC rr
		0x03: func() { // INC BC
			gb.cpu.registers.bc.value++
		},
		0x13: func() { // INC DE
			gb.cpu.registers.de.value++
		},
		0x23: func() { // INC HL
			gb.cpu.registers.hl.value++
		},
		0x33: func() { // INC SP
			gb.cpu.registers.sp++
		},

		// DEC rr
		0x0B: func() { // DEC BC
			gb.cpu.registers.bc.value--
		},
		0x1B: func() { // DEC DE
			gb.cpu.registers.de.value--
		},
		0x2B: func() { // DEC HL
			gb.cpu.registers.hl.value--
		},
		0x3B: func() { // DEC SP
			gb.cpu.registers.sp--
		},

		//--------------------------
		//	    Miscellaneous
		//--------------------------

		// DAA
		0x27: func() { // DAA
			val := gb.cpu.registers.af.getHi()
			val = gb.cpu.ulaDAA(val)
			gb.cpu.registers.af.setHi(val)
		},

		// CPL
		0x2F: func() { // CPL
			val := gb.cpu.registers.af.getHi()
			val = gb.cpu.ulaCPL(val)
			gb.cpu.registers.af.setHi(val)
		},

		// CCF
		0x3F: func() { // CCF
			gb.cpu.ulaCCF()
		},

		// CCF
		0x37: func() { // SCF
			gb.cpu.ulaSCF()
		},

		// NOP
		0x00: func() {
			// NOP
		},

		// HALT
		0x76: func() { // HALT
			gb.cpu.setHalt(true)
		},

		// STOP
		0x10: func() { // STOP 0
			gb.cpu.setHalt(true)
			gb.readOPArg8bit()
		},

		// DI
		0xF3: func() { // DI
			gb.cpu.setInterruptsEnable(false)
		},

		// EL
		0xFB: func() { // EL
			gb.cpu.setInterruptsEnable(true)
		},

		//--------------------------
		//	   Rotates & Shifts
		//--------------------------

		// RLC
		0x07: func() { // RLCA
			val := gb.cpu.registers.af.getHi()
			val = gb.cpu.ulaRLC(val)
			gb.cpu.registers.af.setHi(val)
		},

		// RL
		0x17: func() { // RLA
			val := gb.cpu.registers.af.getHi()
			val = gb.cpu.ulaRL(val)
			gb.cpu.registers.af.setHi(val)
		},

		// RLC
		0x0F: func() { // RRCA
			val := gb.cpu.registers.af.getHi()
			val = gb.cpu.ulaRRC(val)
			gb.cpu.registers.af.setHi(val)
		},

		// RL
		0x1F: func() { // RRA
			val := gb.cpu.registers.af.getHi()
			val = gb.cpu.ulaRR(val)
			gb.cpu.registers.af.setHi(val)
		},

		//--------------------------
		//	   		Jumps
		//--------------------------

		// JP nn
		0xC3: func() { // JP nn
			val := gb.readOPArg16bit()
			gb.cpu.setPC(val)
		},

		// JC cc nn
		0xC2: func() { // JC NZ nn
			val := gb.readOPArg16bit()
			if !gb.cpu.isZF() {
				gb.cpu.setPC(val)
				gb.cpu.cycles += 4
			}
		},
		0xCA: func() { // JC Z nn
			val := gb.readOPArg16bit()
			if gb.cpu.isZF() {
				gb.cpu.setPC(val)
				gb.cpu.cycles += 4
			}
		},
		0xD2: func() { // JC NC nn
			val := gb.readOPArg16bit()
			if !gb.cpu.isCY() {
				gb.cpu.setPC(val)
				gb.cpu.cycles += 4
			}
		},
		0xDA: func() { // JC C nn
			val := gb.readOPArg16bit()
			if gb.cpu.isCY() {
				gb.cpu.setPC(val)
				gb.cpu.cycles += 4
			}
		},

		// JP (HL)
		0xE9: func() { // JP (HL)
			addr := gb.cpu.registers.hl.getHiLo()
			gb.cpu.setPC(addr)
		},

		// JR n
		0x18: func() { // JR n
			addr := gb.cpu.getPC()
			val := int8(gb.readOPArg8bit())
			addr = uint16(int32(addr) + int32(val))
			gb.cpu.setPC(addr)
		},

		// JR cc n
		0x20: func() { // JR NZ n
			addr := gb.cpu.getPC()
			val := int8(gb.readOPArg8bit())
			if !gb.cpu.isZF() {
				addr = uint16(int32(addr) + int32(val))
				gb.cpu.setPC(addr)
				gb.cpu.cycles += 4
			}
		},
		0x28: func() { // JR Z n
			addr := gb.cpu.getPC()
			val := int8(gb.readOPArg8bit())
			if gb.cpu.isZF() {
				addr = uint16(int32(addr) + int32(val))
				gb.cpu.setPC(addr)
				gb.cpu.cycles += 4
			}
		},
		0x30: func() { // JR NC n
			addr := gb.cpu.getPC()
			val := int8(gb.readOPArg8bit())
			if !gb.cpu.isCY() {
				addr = uint16(int32(addr) + int32(val))
				gb.cpu.setPC(addr)
				gb.cpu.cycles += 4
			}
		},
		0x38: func() { // JR C n
			addr := gb.cpu.getPC()
			val := int8(gb.readOPArg8bit())
			if gb.cpu.isCY() {
				addr = uint16(int32(addr) + int32(val))
				gb.cpu.setPC(addr)
				gb.cpu.cycles += 4
			}
		},
		// WARNING
		// CALL n
		0xCD: func() { // CALL n
			addr := gb.readOPArg16bit()
			gb.cpuCall(addr)
		},

		// CALL cc nn
		0xC4: func() { // CALL NZ nn
			addr := gb.readOPArg16bit()
			if !gb.cpu.isZF() {
				gb.cpuCall(addr)
				gb.cpu.cycles += 12
			}
		},
		0xCC: func() { // CALL Z nn
			addr := gb.readOPArg16bit()
			if gb.cpu.isZF() {
				gb.cpuCall(addr)
				gb.cpu.cycles += 12
			}
		},
		0xD4: func() { // CALL NC nn
			addr := gb.readOPArg16bit()
			if !gb.cpu.isCY() {
				gb.cpuCall(addr)
				gb.cpu.cycles += 12
			}
		},
		0xDC: func() { // CALL C nn
			addr := gb.readOPArg16bit()
			if gb.cpu.isCY() {
				gb.cpuCall(addr)
				gb.cpu.cycles += 12
			}
		},

		// RST n
		0xC7: func() { // RST 0x0000
			addr := gb.cpu.getPC()
			gb.push(addr)
			gb.cpu.setPC(0x0000)
		},
		0xCF: func() { // RST 0x08
			addr := gb.cpu.getPC()
			gb.push(addr)
			gb.cpu.setPC(0x0008)
		},
		0xD7: func() { // RST 0x10
			addr := gb.cpu.getPC()
			gb.push(addr)
			gb.cpu.setPC(0x0010)
		},
		0xDF: func() { // RST 0x18
			addr := gb.cpu.getPC()
			gb.push(addr)
			gb.cpu.setPC(0x0018)
		},
		0xE7: func() { // RST 0x20
			addr := gb.cpu.getPC()
			gb.push(addr)
			gb.cpu.setPC(0x0020)
		},
		0xEF: func() { // RST 0x28
			addr := gb.cpu.getPC()
			gb.push(addr)
			gb.cpu.setPC(0x0028)
		},
		0xF7: func() { // RST 0x30
			addr := gb.cpu.getPC()
			gb.push(addr)
			gb.cpu.setPC(0x0030)
		},
		0xFF: func() { // RST 0x38
			addr := gb.cpu.getPC()
			gb.push(addr)
			gb.cpu.setPC(0x0038)
		},

		// RET
		0xC9: func() { // RET
			addr := gb.pop()
			gb.cpu.setPC(addr)
		},

		// RET cc
		0xC0: func() { // RET NZ
			if !gb.cpu.isZF() {
				addr := gb.pop()
				gb.cpu.setPC(addr)
				gb.cpu.cycles += 12
			}
		},
		0xC8: func() { // RET NZ
			if gb.cpu.isZF() {
				addr := gb.pop()
				gb.cpu.setPC(addr)
				gb.cpu.cycles += 12
			}
		},
		0xD0: func() { // RET NZ
			if !gb.cpu.isCY() {
				addr := gb.pop()
				gb.cpu.setPC(addr)
				gb.cpu.cycles += 12
			}
		},
		0xD8: func() { // RET NZ
			if gb.cpu.isCY() {
				addr := gb.pop()
				gb.cpu.setPC(addr)
				gb.cpu.cycles += 12
			}
		},

		// RET I
		0xD9: func() { // RET I
			addr := gb.pop()
			gb.cpu.setPC(addr)
			gb.cpu.setInterruptsEnable(true)
		},

		// CB
		0xCB: func() { // CB
			opCodeCB := gb.readOPArg8bit()
			gb.cpu.cycles = OpcodeCBCycles[opCodeCB] * 4
			gb.opCodeCB[opCodeCB]()
		},
	}
	return dicOP
}
