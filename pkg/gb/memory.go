package gb

import "log"

type Memory struct {
	ROM     [0x8000]byte
	VRAM    [0x2000]byte
	RAM     [0x4000]byte
	EchoRAM [0x1E00]byte
	OAM     [0x100]byte //Tile
	IOPort  [0x100]byte
	HRAM    [0x100]byte
}

func (mem *Memory) read8bit(addr uint16) uint8 {
	var val uint8
	if addr < 0x8000 {
		val = mem.ROM[addr]
	} else if (addr >= 0x8000) && (addr < 0xA000) {
		val = mem.VRAM[addr-0x8000]
	} else if (addr >= 0xA000) && (addr < 0xE000) {
		val = mem.RAM[addr-0xA000]
	} else if (addr >= 0xE000) && (addr < 0xFE00) {
		val = mem.EchoRAM[addr-0xE000]
	} else if (addr >= 0xFE00) && (addr < 0xFEA0) {
		val = mem.OAM[addr-0xFE00]
	} else if (addr >= 0xFF00) && (addr < 0xFF80) {
		val = mem.IOPort[addr-0xFF00]
	} else if (addr >= 0xFF80) && (addr < 0xFFFF) {
		val = mem.IOPort[addr-0xFF00]
	} else {
		val = 0
		log.Printf("Address Invalid : %d \n", addr)
	}
	return val
}

func (mem *Memory) read16bit(addr uint16) uint16 {
	val := uint16(mem.read8bit(addr))
	val = val | uint16(mem.read8bit(addr+1))<<8
	return val
}

func (mem *Memory) write8bit(addr uint16, val uint8) {
	if addr < 0x8000 {
		mem.ROM[addr] = val
	} else if (addr >= 0x8000) && (addr < 0xA000) {
		mem.VRAM[addr-0x8000] = val
	} else if (addr >= 0xA000) && (addr < 0xE000) {
		mem.RAM[addr-0xA000] = val
	} else if (addr >= 0xE000) && (addr < 0xFE00) {
		mem.EchoRAM[addr-0xE000] = val
		mem.write8bit(addr-0x2000, val)
	} else if (addr >= 0xFE00) && (addr < 0xFEA0) {
		mem.OAM[addr-0xFE00] = val
	} else if (addr >= 0xFF00) && (addr < 0xFF80) {
		mem.IOPort[addr-0xFF00] = val
	} else if (addr >= 0xFF80) && (addr < 0xFFFF) {
		mem.IOPort[addr-0xFF00] = val
	} else {
		log.Printf("Address Invalid : %d \n", addr)
	}
}

func (mem *Memory) write16bit(addr uint16, val uint16) {
	mem.write8bit(addr, uint8(val))
	mem.write8bit(addr+1, uint8(val>>8))
}
