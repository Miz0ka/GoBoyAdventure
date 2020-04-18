package gb

var OpcodeCBCycles = []int{
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 0
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 1
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 2
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 3
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2, // 4
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2, // 5
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2, // 6
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2, // 7
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 8
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 9
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // A
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // B
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // C
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // D
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // E
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // F
} //0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f

func (gb *gameboy) genereOPCBDic() [0x100]func() {
	dicOP := [0x100]func(){
		//--------------------------
		//	    Miscellaneous
		//--------------------------

		// SWAP n
		0x37: func() { // SWAP A
			val := gb.cpu.registers.af.getHi()
			val = gb.cpu.ulaSwap(val)
			gb.cpu.registers.af.setHi(val)
		},
		0x30: func() { // SWAP B
			val := gb.cpu.registers.bc.getHi()
			val = gb.cpu.ulaSwap(val)
			gb.cpu.registers.bc.setHi(val)
		},
		0x31: func() { // SWAP C
			val := gb.cpu.registers.bc.getLo()
			val = gb.cpu.ulaSwap(val)
			gb.cpu.registers.bc.setLo(val)
		},
		0x32: func() { // SWAP D
			val := gb.cpu.registers.de.getHi()
			val = gb.cpu.ulaSwap(val)
			gb.cpu.registers.de.setHi(val)
		},
		0x33: func() { // SWAP E
			val := gb.cpu.registers.de.getLo()
			val = gb.cpu.ulaSwap(val)
			gb.cpu.registers.de.setLo(val)
		},
		0x34: func() { // SWAP H
			val := gb.cpu.registers.hl.getHi()
			val = gb.cpu.ulaSwap(val)
			gb.cpu.registers.hl.setHi(val)
		},
		0x35: func() { // SWAP L
			val := gb.cpu.registers.hl.getLo()
			val = gb.cpu.ulaSwap(val)
			gb.cpu.registers.hl.setLo(val)
		},
		0x36: func() { // SWAP (HL)
			addr := gb.cpu.registers.hl.getHiLo()
			val := gb.memory.read8bit(addr)
			val = gb.cpu.ulaSwap(val)
			gb.memory.write8bit(addr, val)
		},

		//--------------------------
		//	   Rotates & Shifts
		//--------------------------

		// RLC r
		0x07: func() { // RLC A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaRLC(val)
			gb.cpu.registers.af.setHi(total)
		},
		0x00: func() { // RLC B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaRLC(val)
			gb.cpu.registers.bc.setHi(total)
		},
		0x01: func() { // RLC C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaRLC(val)
			gb.cpu.registers.bc.setLo(total)
		},
		0x02: func() { // RLC D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaRLC(val)
			gb.cpu.registers.de.setHi(total)
		},
		0x03: func() { // RLC E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaRLC(val)
			gb.cpu.registers.de.setLo(total)
		},
		0x04: func() { // RLC H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaRLC(val)
			gb.cpu.registers.hl.setHi(total)
		},
		0x05: func() { // RLC L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaRLC(val)
			gb.cpu.registers.hl.setLo(total)
		},
		0x06: func() { // RLC (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaRLC(val)
			gb.memory.write8bit(addr, total)
		},

		// RL r
		0x17: func() { // RL A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaRL(val)
			gb.cpu.registers.af.setHi(total)
		},
		0x10: func() { // RL B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaRL(val)
			gb.cpu.registers.bc.setHi(total)
		},
		0x11: func() { // RL C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaRL(val)
			gb.cpu.registers.bc.setLo(total)
		},
		0x12: func() { // RL D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaRL(val)
			gb.cpu.registers.de.setHi(total)
		},
		0x13: func() { // RL E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaRL(val)
			gb.cpu.registers.de.setLo(total)
		},
		0x14: func() { // RL H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaRL(val)
			gb.cpu.registers.hl.setHi(total)
		},
		0x15: func() { // RL L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaRL(val)
			gb.cpu.registers.hl.setLo(total)
		},
		0x16: func() { // RL (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaRL(val)
			gb.memory.write8bit(addr, total)
		},

		// RRC r
		0x0F: func() { // RRC A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaRRC(val)
			gb.cpu.registers.af.setHi(total)
		},
		0x08: func() { // RRC B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaRRC(val)
			gb.cpu.registers.bc.setHi(total)
		},
		0x09: func() { // RRC C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaRRC(val)
			gb.cpu.registers.bc.setLo(total)
		},
		0x0A: func() { // RRC D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaRRC(val)
			gb.cpu.registers.de.setHi(total)
		},
		0x0B: func() { // RRC E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaRRC(val)
			gb.cpu.registers.de.setLo(total)
		},
		0x0C: func() { // RRC H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaRRC(val)
			gb.cpu.registers.hl.setHi(total)
		},
		0x0D: func() { // RRC L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaRRC(val)
			gb.cpu.registers.hl.setLo(total)
		},
		0x0E: func() { // RRC (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaRRC(val)
			gb.memory.write8bit(addr, total)
		},

		// RR r
		0x1F: func() { // RR A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaRR(val)
			gb.cpu.registers.af.setHi(total)
		},
		0x18: func() { // RR B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaRR(val)
			gb.cpu.registers.bc.setHi(total)
		},
		0x19: func() { // RR C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaRR(val)
			gb.cpu.registers.bc.setLo(total)
		},
		0x1A: func() { // RR D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaRR(val)
			gb.cpu.registers.de.setHi(total)
		},
		0x1B: func() { // RR E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaRR(val)
			gb.cpu.registers.de.setLo(total)
		},
		0x1C: func() { // RR H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaRR(val)
			gb.cpu.registers.hl.setHi(total)
		},
		0x1D: func() { // RR L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaRR(val)
			gb.cpu.registers.hl.setLo(total)
		},
		0x1E: func() { // RR (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaRR(val)
			gb.memory.write8bit(addr, total)
		},

		// SLA r
		0x27: func() { // SLA A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaSLA(val)
			gb.cpu.registers.af.setHi(total)
		},
		0x20: func() { // SLA B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaSLA(val)
			gb.cpu.registers.bc.setHi(total)
		},
		0x21: func() { // SLA C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaSLA(val)
			gb.cpu.registers.bc.setLo(total)
		},
		0x22: func() { // SLA D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaSLA(val)
			gb.cpu.registers.de.setHi(total)
		},
		0x23: func() { // SLA E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaSLA(val)
			gb.cpu.registers.de.setLo(total)
		},
		0x24: func() { // SLA H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaSLA(val)
			gb.cpu.registers.hl.setHi(total)
		},
		0x25: func() { // SLA L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaSLA(val)
			gb.cpu.registers.hl.setLo(total)
		},
		0x26: func() { // SLA (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaSLA(val)
			gb.memory.write8bit(addr, total)
		},

		// SRA r
		0x2F: func() { // SRA A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaSRA(val)
			gb.cpu.registers.af.setHi(total)
		},
		0x28: func() { // SRA B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaSRA(val)
			gb.cpu.registers.bc.setHi(total)
		},
		0x29: func() { // SRA C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaSRA(val)
			gb.cpu.registers.bc.setLo(total)
		},
		0x2A: func() { // SRA D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaSRA(val)
			gb.cpu.registers.de.setHi(total)
		},
		0x2B: func() { // SRA E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaSRA(val)
			gb.cpu.registers.de.setLo(total)
		},
		0x2C: func() { // SRA H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaSRA(val)
			gb.cpu.registers.hl.setHi(total)
		},
		0x2D: func() { // SRA L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaSRA(val)
			gb.cpu.registers.hl.setLo(total)
		},
		0x2E: func() { // SRA (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaSRA(val)
			gb.memory.write8bit(addr, total)
		},

		// SRL n
		0x3F: func() { // SRL A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaSRL(val)
			gb.cpu.registers.af.setHi(total)
		},
		0x38: func() { // SRL B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaSRL(val)
			gb.cpu.registers.bc.setHi(total)
		},
		0x39: func() { // SRL C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaSRL(val)
			gb.cpu.registers.bc.setLo(total)
		},
		0x3A: func() { // SRL D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaSRL(val)
			gb.cpu.registers.de.setHi(total)
		},
		0x3B: func() { // SRL E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaSRL(val)
			gb.cpu.registers.de.setLo(total)
		},
		0x3C: func() { // SRL H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaSRL(val)
			gb.cpu.registers.hl.setHi(total)
		},
		0x3D: func() { // SRL L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaSRL(val)
			gb.cpu.registers.hl.setLo(total)
		},
		0x3E: func() { // SRL (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaSRL(val)
			gb.memory.write8bit(addr, total)
		},

		// BIT 0, r
		0x47: func() { // BIT 0 A
			val := gb.cpu.registers.af.getHi()
			gb.cpu.ulaBit(0, val)
		},
		0x40: func() { // BIT 0 B
			val := gb.cpu.registers.bc.getHi()
			gb.cpu.ulaBit(0, val)
		},
		0x41: func() { // BIT 0 C
			val := gb.cpu.registers.bc.getLo()
			gb.cpu.ulaBit(0, val)
		},
		0x42: func() { // BIT 0 D
			val := gb.cpu.registers.de.getHi()
			gb.cpu.ulaBit(0, val)
		},
		0x43: func() { // BIT 0 E
			val := gb.cpu.registers.de.getLo()
			gb.cpu.ulaBit(0, val)
		},
		0x44: func() { // BIT 0 H
			val := gb.cpu.registers.hl.getHi()
			gb.cpu.ulaBit(0, val)
		},
		0x45: func() { // BIT 0 L
			val := gb.cpu.registers.hl.getLo()
			gb.cpu.ulaBit(0, val)
		},
		0x46: func() { // BIT 0 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.ulaBit(0, val)
		},

		// BIT 1 r
		0x4F: func() { // BIT 1 A
			val := gb.cpu.registers.af.getHi()
			gb.cpu.ulaBit(1, val)
		},
		0x48: func() { // BIT 1 B
			val := gb.cpu.registers.bc.getHi()
			gb.cpu.ulaBit(1, val)
		},
		0x49: func() { // BIT 1 C
			val := gb.cpu.registers.bc.getLo()
			gb.cpu.ulaBit(1, val)
		},
		0x4A: func() { // BIT 1 D
			val := gb.cpu.registers.de.getHi()
			gb.cpu.ulaBit(1, val)
		},
		0x4B: func() { // BIT 1 E
			val := gb.cpu.registers.de.getLo()
			gb.cpu.ulaBit(1, val)
		},
		0x4C: func() { // BIT 1 H
			val := gb.cpu.registers.hl.getHi()
			gb.cpu.ulaBit(1, val)
		},
		0x4D: func() { // BIT 1 L
			val := gb.cpu.registers.hl.getLo()
			gb.cpu.ulaBit(1, val)
		},
		0x4E: func() { // BIT 1 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.ulaBit(1, val)
		},

		// BIT 2, r
		0x57: func() { // BIT 2 A
			val := gb.cpu.registers.af.getHi()
			gb.cpu.ulaBit(2, val)
		},
		0x50: func() { // BIT 2 B
			val := gb.cpu.registers.bc.getHi()
			gb.cpu.ulaBit(2, val)
		},
		0x51: func() { // BIT 2 C
			val := gb.cpu.registers.bc.getLo()
			gb.cpu.ulaBit(2, val)
		},
		0x52: func() { // BIT 2 D
			val := gb.cpu.registers.de.getHi()
			gb.cpu.ulaBit(2, val)
		},
		0x53: func() { // BIT 2 E
			val := gb.cpu.registers.de.getLo()
			gb.cpu.ulaBit(2, val)
		},
		0x54: func() { // BIT 2 H
			val := gb.cpu.registers.hl.getHi()
			gb.cpu.ulaBit(2, val)
		},
		0x55: func() { // BIT 2 L
			val := gb.cpu.registers.hl.getLo()
			gb.cpu.ulaBit(2, val)
		},
		0x56: func() { // BIT 2 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.ulaBit(2, val)
		},

		// BIT 3 r
		0x5F: func() { // BIT 3 A
			val := gb.cpu.registers.af.getHi()
			gb.cpu.ulaBit(3, val)
		},
		0x58: func() { // BIT 3 B
			val := gb.cpu.registers.bc.getHi()
			gb.cpu.ulaBit(3, val)
		},
		0x59: func() { // BIT 3 C
			val := gb.cpu.registers.bc.getLo()
			gb.cpu.ulaBit(3, val)
		},
		0x5A: func() { // BIT 3 D
			val := gb.cpu.registers.de.getHi()
			gb.cpu.ulaBit(3, val)
		},
		0x5B: func() { // BIT 3 E
			val := gb.cpu.registers.de.getLo()
			gb.cpu.ulaBit(3, val)
		},
		0x5C: func() { // BIT 3 H
			val := gb.cpu.registers.hl.getHi()
			gb.cpu.ulaBit(3, val)
		},
		0x5D: func() { // BIT 3 L
			val := gb.cpu.registers.hl.getLo()
			gb.cpu.ulaBit(3, val)
		},
		0x5E: func() { // BIT 3 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.ulaBit(3, val)
		},

		// BIT 4, r
		0x67: func() { // BIT 4 A
			val := gb.cpu.registers.af.getHi()
			gb.cpu.ulaBit(4, val)
		},
		0x60: func() { // BIT 4 B
			val := gb.cpu.registers.bc.getHi()
			gb.cpu.ulaBit(4, val)
		},
		0x61: func() { // BIT 4 C
			val := gb.cpu.registers.bc.getLo()
			gb.cpu.ulaBit(4, val)
		},
		0x62: func() { // BIT 4 D
			val := gb.cpu.registers.de.getHi()
			gb.cpu.ulaBit(4, val)
		},
		0x63: func() { // BIT 4 E
			val := gb.cpu.registers.de.getLo()
			gb.cpu.ulaBit(4, val)
		},
		0x64: func() { // BIT 4 H
			val := gb.cpu.registers.hl.getHi()
			gb.cpu.ulaBit(4, val)
		},
		0x65: func() { // BIT 4 L
			val := gb.cpu.registers.hl.getLo()
			gb.cpu.ulaBit(4, val)
		},
		0x66: func() { // BIT 4 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.ulaBit(4, val)
		},

		// BIT 5 r
		0x6F: func() { // BIT 5 A
			val := gb.cpu.registers.af.getHi()
			gb.cpu.ulaBit(5, val)
		},
		0x68: func() { // BIT 5 B
			val := gb.cpu.registers.bc.getHi()
			gb.cpu.ulaBit(5, val)
		},
		0x69: func() { // BIT 5 C
			val := gb.cpu.registers.bc.getLo()
			gb.cpu.ulaBit(5, val)
		},
		0x6A: func() { // BIT 5 D
			val := gb.cpu.registers.de.getHi()
			gb.cpu.ulaBit(5, val)
		},
		0x6B: func() { // BIT 5 E
			val := gb.cpu.registers.de.getLo()
			gb.cpu.ulaBit(5, val)
		},
		0x6C: func() { // BIT 5 H
			val := gb.cpu.registers.hl.getHi()
			gb.cpu.ulaBit(5, val)
		},
		0x6D: func() { // BIT 5 L
			val := gb.cpu.registers.hl.getLo()
			gb.cpu.ulaBit(5, val)
		},
		0x6E: func() { // BIT 5 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.ulaBit(5, val)
		},

		// BIT 6, r
		0x77: func() { // BIT 6 A
			val := gb.cpu.registers.af.getHi()
			gb.cpu.ulaBit(6, val)
		},
		0x70: func() { // BIT 6 B
			val := gb.cpu.registers.bc.getHi()
			gb.cpu.ulaBit(6, val)
		},
		0x71: func() { // BIT 6 C
			val := gb.cpu.registers.bc.getLo()
			gb.cpu.ulaBit(6, val)
		},
		0x72: func() { // BIT 6 D
			val := gb.cpu.registers.de.getHi()
			gb.cpu.ulaBit(6, val)
		},
		0x73: func() { // BIT 6 E
			val := gb.cpu.registers.de.getLo()
			gb.cpu.ulaBit(6, val)
		},
		0x74: func() { // BIT 6 H
			val := gb.cpu.registers.hl.getHi()
			gb.cpu.ulaBit(6, val)
		},
		0x75: func() { // BIT 6 L
			val := gb.cpu.registers.hl.getLo()
			gb.cpu.ulaBit(6, val)
		},
		0x76: func() { // BIT 6 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.ulaBit(6, val)
		},

		// BIT 7 r
		0x7F: func() { // BIT 7 A
			val := gb.cpu.registers.af.getHi()
			gb.cpu.ulaBit(7, val)
		},
		0x78: func() { // BIT 7 B
			val := gb.cpu.registers.bc.getHi()
			gb.cpu.ulaBit(7, val)
		},
		0x79: func() { // BIT 7 C
			val := gb.cpu.registers.bc.getLo()
			gb.cpu.ulaBit(7, val)
		},
		0x7A: func() { // BIT 7 D
			val := gb.cpu.registers.de.getHi()
			gb.cpu.ulaBit(7, val)
		},
		0x7B: func() { // BIT 7 E
			val := gb.cpu.registers.de.getLo()
			gb.cpu.ulaBit(7, val)
		},
		0x7C: func() { // BIT 7 H
			val := gb.cpu.registers.hl.getHi()
			gb.cpu.ulaBit(7, val)
		},
		0x7D: func() { // BIT 7 L
			val := gb.cpu.registers.hl.getLo()
			gb.cpu.ulaBit(7, val)
		},
		0x7E: func() { // BIT 7 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			gb.cpu.ulaBit(7, val)
		},

		// RES 0 r
		0x87: func() { // RES 0 A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaRes(0, val)
			gb.cpu.registers.af.setHi(total)
		},
		0x80: func() { // RES 0 B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaRes(0, val)
			gb.cpu.registers.bc.setHi(total)
		},
		0x81: func() { // RES 0 C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaRes(0, val)
			gb.cpu.registers.bc.setLo(total)
		},
		0x82: func() { // RES 0 D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaRes(0, val)
			gb.cpu.registers.de.setHi(total)
		},
		0x83: func() { // RES 0 E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaRes(0, val)
			gb.cpu.registers.de.setLo(total)
		},
		0x84: func() { // RES 0 H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaRes(0, val)
			gb.cpu.registers.hl.setHi(total)
		},
		0x85: func() { // RES 0 L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaRes(0, val)
			gb.cpu.registers.hl.setLo(total)
		},
		0x86: func() { // RES 0 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaRes(0, val)
			gb.memory.write8bit(addr, total)
		},

		// RES 1 r
		0x8F: func() { // RES 1 A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaRes(1, val)
			gb.cpu.registers.af.setHi(total)
		},
		0x88: func() { // RES 1 B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaRes(1, val)
			gb.cpu.registers.bc.setHi(total)
		},
		0x89: func() { // RES 1 C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaRes(1, val)
			gb.cpu.registers.bc.setLo(total)
		},
		0x8A: func() { // RES 1 D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaRes(1, val)
			gb.cpu.registers.de.setHi(total)
		},
		0x8B: func() { // RES 1 E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaRes(1, val)
			gb.cpu.registers.de.setLo(total)
		},
		0x8C: func() { // RES 1 H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaRes(1, val)
			gb.cpu.registers.hl.setHi(total)
		},
		0x8D: func() { // RES 1 L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaRes(1, val)
			gb.cpu.registers.hl.setLo(total)
		},
		0x8E: func() { // RES 1 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaRes(1, val)
			gb.memory.write8bit(addr, total)
		},

		// RES 2 r
		0x97: func() { // RES 2 A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaRes(2, val)
			gb.cpu.registers.af.setHi(total)
		},
		0x90: func() { // RES 2 B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaRes(2, val)
			gb.cpu.registers.bc.setHi(total)
		},
		0x91: func() { // RES 2 C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaRes(2, val)
			gb.cpu.registers.bc.setLo(total)
		},
		0x92: func() { // RES 2 D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaRes(2, val)
			gb.cpu.registers.de.setHi(total)
		},
		0x93: func() { // RES 2 E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaRes(2, val)
			gb.cpu.registers.de.setLo(total)
		},
		0x94: func() { // RES 2 H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaRes(2, val)
			gb.cpu.registers.hl.setHi(total)
		},
		0x95: func() { // RES 2 L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaRes(2, val)
			gb.cpu.registers.hl.setLo(total)
		},
		0x96: func() { // RES 2 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaRes(2, val)
			gb.memory.write8bit(addr, total)
		},

		// RES 3 r
		0x9F: func() { // RES 3 A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaRes(3, val)
			gb.cpu.registers.af.setHi(total)
		},
		0x98: func() { // RES 3 B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaRes(3, val)
			gb.cpu.registers.bc.setHi(total)
		},
		0x99: func() { // RES 3 C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaRes(3, val)
			gb.cpu.registers.bc.setLo(total)
		},
		0x9A: func() { // RES 3 D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaRes(3, val)
			gb.cpu.registers.de.setHi(total)
		},
		0x9B: func() { // RES 3 E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaRes(3, val)
			gb.cpu.registers.de.setLo(total)
		},
		0x9C: func() { // RES 3 H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaRes(3, val)
			gb.cpu.registers.hl.setHi(total)
		},
		0x9D: func() { // RES 3 L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaRes(3, val)
			gb.cpu.registers.hl.setLo(total)
		},
		0x9E: func() { // RES 3 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaRes(3, val)
			gb.memory.write8bit(addr, total)
		},

		// RES 4 r
		0xA7: func() { // RES 4 A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaRes(4, val)
			gb.cpu.registers.af.setHi(total)
		},
		0xA0: func() { // RES 4 B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaRes(4, val)
			gb.cpu.registers.bc.setHi(total)
		},
		0xA1: func() { // RES 4 C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaRes(4, val)
			gb.cpu.registers.bc.setLo(total)
		},
		0xA2: func() { // RES 4 D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaRes(4, val)
			gb.cpu.registers.de.setHi(total)
		},
		0xA3: func() { // RES 4 E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaRes(4, val)
			gb.cpu.registers.de.setLo(total)
		},
		0xA4: func() { // RES 4 H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaRes(4, val)
			gb.cpu.registers.hl.setHi(total)
		},
		0xA5: func() { // RES 4 L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaRes(4, val)
			gb.cpu.registers.hl.setLo(total)
		},
		0xA6: func() { // RES 4 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaRes(4, val)
			gb.memory.write8bit(addr, total)
		},

		// RES 5 r
		0xAF: func() { // RES 5 A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaRes(5, val)
			gb.cpu.registers.af.setHi(total)
		},
		0xA8: func() { // RES 5 B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaRes(5, val)
			gb.cpu.registers.bc.setHi(total)
		},
		0xA9: func() { // RES 5 C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaRes(5, val)
			gb.cpu.registers.bc.setLo(total)
		},
		0xAA: func() { // RES 5 D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaRes(5, val)
			gb.cpu.registers.de.setHi(total)
		},
		0xAB: func() { // RES 5 E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaRes(5, val)
			gb.cpu.registers.de.setLo(total)
		},
		0xAC: func() { // RES 5 H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaRes(5, val)
			gb.cpu.registers.hl.setHi(total)
		},
		0xAD: func() { // RES 5 L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaRes(5, val)
			gb.cpu.registers.hl.setLo(total)
		},
		0xAE: func() { // RES 5 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaRes(5, val)
			gb.memory.write8bit(addr, total)
		},
		// RES 6 r
		0xB7: func() { // RES 6 A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaRes(6, val)
			gb.cpu.registers.af.setHi(total)
		},
		0xB0: func() { // RES 6 B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaRes(6, val)
			gb.cpu.registers.bc.setHi(total)
		},
		0xB1: func() { // RES 6 C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaRes(6, val)
			gb.cpu.registers.bc.setLo(total)
		},
		0xB2: func() { // RES 6 D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaRes(6, val)
			gb.cpu.registers.de.setHi(total)
		},
		0xB3: func() { // RES 6 E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaRes(6, val)
			gb.cpu.registers.de.setLo(total)
		},
		0xB4: func() { // RES 6 H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaRes(6, val)
			gb.cpu.registers.hl.setHi(total)
		},
		0xB5: func() { // RES 6 L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaRes(6, val)
			gb.cpu.registers.hl.setLo(total)
		},
		0xB6: func() { // RES 6 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaRes(6, val)
			gb.memory.write8bit(addr, total)
		},

		// RES 7 r
		0xBF: func() { // RES 7 A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaRes(7, val)
			gb.cpu.registers.af.setHi(total)
		},
		0xB8: func() { // RES 7 B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaRes(7, val)
			gb.cpu.registers.bc.setHi(total)
		},
		0xB9: func() { // RES 7 C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaRes(7, val)
			gb.cpu.registers.bc.setLo(total)
		},
		0xBA: func() { // RES 7 D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaRes(7, val)
			gb.cpu.registers.de.setHi(total)
		},
		0xBB: func() { // RES 7 E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaRes(7, val)
			gb.cpu.registers.de.setLo(total)
		},
		0xBC: func() { // RES 7 H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaRes(7, val)
			gb.cpu.registers.hl.setHi(total)
		},
		0xBD: func() { // RES 7 L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaRes(7, val)
			gb.cpu.registers.hl.setLo(total)
		},
		0xBE: func() { // RES 7 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaRes(7, val)
			gb.memory.write8bit(addr, total)
		},

		// SET 0 r
		0xC7: func() { // SET 0 A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaSet(0, val)
			gb.cpu.registers.af.setHi(total)
		},
		0xC0: func() { // SET 0 B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaSet(0, val)
			gb.cpu.registers.bc.setHi(total)
		},
		0xC1: func() { // SET 0 C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaSet(0, val)
			gb.cpu.registers.bc.setLo(total)
		},
		0xC2: func() { // SET 0 D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaSet(0, val)
			gb.cpu.registers.de.setHi(total)
		},
		0xC3: func() { // SET 0 E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaSet(0, val)
			gb.cpu.registers.de.setLo(total)
		},
		0xC4: func() { // SET 0 H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaSet(0, val)
			gb.cpu.registers.hl.setHi(total)
		},
		0xC5: func() { // SET 0 L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaSet(0, val)
			gb.cpu.registers.hl.setLo(total)
		},
		0xC6: func() { // SET 0 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaSet(0, val)
			gb.memory.write8bit(addr, total)
		},

		// SET 1 r
		0xCF: func() { // SET 1 A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaSet(1, val)
			gb.cpu.registers.af.setHi(total)
		},
		0xC8: func() { // SET 1 B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaSet(1, val)
			gb.cpu.registers.bc.setHi(total)
		},
		0xC9: func() { // SET 1 C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaSet(1, val)
			gb.cpu.registers.bc.setLo(total)
		},
		0xCA: func() { // SET 1 D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaSet(1, val)
			gb.cpu.registers.de.setHi(total)
		},
		0xCB: func() { // SET 1 E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaSet(1, val)
			gb.cpu.registers.de.setLo(total)
		},
		0xCC: func() { // SET 1 H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaSet(1, val)
			gb.cpu.registers.hl.setHi(total)
		},
		0xCD: func() { // SET 1 L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaSet(1, val)
			gb.cpu.registers.hl.setLo(total)
		},
		0xCE: func() { // SET 1 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaSet(1, val)
			gb.memory.write8bit(addr, total)
		},

		// SET 2 r
		0xD7: func() { // SET 2 A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaSet(2, val)
			gb.cpu.registers.af.setHi(total)
		},
		0xD0: func() { // SET 2 B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaSet(2, val)
			gb.cpu.registers.bc.setHi(total)
		},
		0xD1: func() { // SET 2 C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaSet(2, val)
			gb.cpu.registers.bc.setLo(total)
		},
		0xD2: func() { // SET 2 D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaSet(2, val)
			gb.cpu.registers.de.setHi(total)
		},
		0xD3: func() { // SET 2 E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaSet(2, val)
			gb.cpu.registers.de.setLo(total)
		},
		0xD4: func() { // SET 2 H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaSet(2, val)
			gb.cpu.registers.hl.setHi(total)
		},
		0xD5: func() { // SET 2 L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaSet(2, val)
			gb.cpu.registers.hl.setLo(total)
		},
		0xD6: func() { // SET 2 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaSet(2, val)
			gb.memory.write8bit(addr, total)
		},

		// SET 3 r
		0xDF: func() { // SET 3 A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaSet(3, val)
			gb.cpu.registers.af.setHi(total)
		},
		0xD8: func() { // SET 3 B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaSet(3, val)
			gb.cpu.registers.bc.setHi(total)
		},
		0xD9: func() { // SET 3 C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaSet(3, val)
			gb.cpu.registers.bc.setLo(total)
		},
		0xDA: func() { // SET 3 D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaSet(3, val)
			gb.cpu.registers.de.setHi(total)
		},
		0xDB: func() { // SET 3 E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaSet(3, val)
			gb.cpu.registers.de.setLo(total)
		},
		0xDC: func() { // SET 3 H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaSet(3, val)
			gb.cpu.registers.hl.setHi(total)
		},
		0xDD: func() { // SET 3 L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaSet(3, val)
			gb.cpu.registers.hl.setLo(total)
		},
		0xDE: func() { // SET 3 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaSet(3, val)
			gb.memory.write8bit(addr, total)
		},

		// SET 4 r
		0xE7: func() { // SET 4 A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaSet(4, val)
			gb.cpu.registers.af.setHi(total)
		},
		0xE0: func() { // SET 4 B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaSet(4, val)
			gb.cpu.registers.bc.setHi(total)
		},
		0xE1: func() { // SET 4 C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaSet(4, val)
			gb.cpu.registers.bc.setLo(total)
		},
		0xE2: func() { // SET 4 D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaSet(4, val)
			gb.cpu.registers.de.setHi(total)
		},
		0xE3: func() { // SET 4 E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaSet(4, val)
			gb.cpu.registers.de.setLo(total)
		},
		0xE4: func() { // SET 4 H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaSet(4, val)
			gb.cpu.registers.hl.setHi(total)
		},
		0xE5: func() { // SET 4 L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaSet(4, val)
			gb.cpu.registers.hl.setLo(total)
		},
		0xE6: func() { // SET 4 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaSet(4, val)
			gb.memory.write8bit(addr, total)
		},

		// SET 5 r
		0xEF: func() { // SET 5 A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaSet(5, val)
			gb.cpu.registers.af.setHi(total)
		},
		0xE8: func() { // SET 5 B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaSet(5, val)
			gb.cpu.registers.bc.setHi(total)
		},
		0xE9: func() { // SET 5 C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaSet(5, val)
			gb.cpu.registers.bc.setLo(total)
		},
		0xEA: func() { // SET 5 D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaSet(5, val)
			gb.cpu.registers.de.setHi(total)
		},
		0xEB: func() { // SET 5 E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaSet(5, val)
			gb.cpu.registers.de.setLo(total)
		},
		0xEC: func() { // SET 5 H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaSet(5, val)
			gb.cpu.registers.hl.setHi(total)
		},
		0xED: func() { // SET 5 L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaSet(5, val)
			gb.cpu.registers.hl.setLo(total)
		},
		0xEE: func() { // SET 5 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaSet(5, val)
			gb.memory.write8bit(addr, total)
		},
		// SET 6 r
		0xF7: func() { // SET 6 A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaSet(6, val)
			gb.cpu.registers.af.setHi(total)
		},
		0xF0: func() { // SET 6 B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaSet(6, val)
			gb.cpu.registers.bc.setHi(total)
		},
		0xF1: func() { // SET 6 C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaSet(6, val)
			gb.cpu.registers.bc.setLo(total)
		},
		0xF2: func() { // SET 6 D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaSet(6, val)
			gb.cpu.registers.de.setHi(total)
		},
		0xF3: func() { // SET 6 E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaSet(6, val)
			gb.cpu.registers.de.setLo(total)
		},
		0xF4: func() { // SET 6 H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaSet(6, val)
			gb.cpu.registers.hl.setHi(total)
		},
		0xF5: func() { // SET 6 L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaSet(6, val)
			gb.cpu.registers.hl.setLo(total)
		},
		0xF6: func() { // SET 6 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaSet(6, val)
			gb.memory.write8bit(addr, total)
		},

		// SET 7 r
		0xFF: func() { // SET 7 A
			val := gb.cpu.registers.af.getHi()
			total := gb.cpu.ulaSet(7, val)
			gb.cpu.registers.af.setHi(total)
		},
		0xF8: func() { // SET 7 B
			val := gb.cpu.registers.bc.getHi()
			total := gb.cpu.ulaSet(7, val)
			gb.cpu.registers.bc.setHi(total)
		},
		0xF9: func() { // SET 7 C
			val := gb.cpu.registers.bc.getLo()
			total := gb.cpu.ulaSet(7, val)
			gb.cpu.registers.bc.setLo(total)
		},
		0xFA: func() { // SET 7 D
			val := gb.cpu.registers.de.getHi()
			total := gb.cpu.ulaSet(7, val)
			gb.cpu.registers.de.setHi(total)
		},
		0xFB: func() { // SET 7 E
			val := gb.cpu.registers.de.getLo()
			total := gb.cpu.ulaSet(7, val)
			gb.cpu.registers.de.setLo(total)
		},
		0xFC: func() { // SET 7 H
			val := gb.cpu.registers.hl.getHi()
			total := gb.cpu.ulaSet(7, val)
			gb.cpu.registers.hl.setHi(total)
		},
		0xFD: func() { // SET 7 L
			val := gb.cpu.registers.hl.getLo()
			total := gb.cpu.ulaSet(7, val)
			gb.cpu.registers.hl.setLo(total)
		},
		0xFE: func() { // SET 7 (HL)
			addr := gb.cpu.registers.bc.getHiLo()
			val := gb.memory.read8bit(addr)
			total := gb.cpu.ulaSet(7, val)
			gb.memory.write8bit(addr, total)
		},
	}
	return dicOP
}
