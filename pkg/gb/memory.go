package gb

import (
	"fmt"
	"log"
)

type memory struct {
	bios          [0x100]byte
	bootstrapDone bool
	ROM           [0x8000]byte
	VRAM          [0x2000]byte
	RAM           [0x4000]byte
	EchoRAM       [0x1E00]byte
	OAM           [0x100]byte //Tile
	IOPort        [0x100]byte
	//HRAM          [0x100]byte
	cartridge *Cartridge
	gb        *Gameboy
}

func memoryInit(cartridgeUrl string, gb *Gameboy) *memory {
	var mem memory
	mem.bios = [0x100]byte{ // Nintendo Logo
		0x31, 0xFE, 0xFF, 0xAF, 0x21, 0xFF, 0x9F, 0x32, 0xCB, 0x7C, 0x20, 0xFB, 0x21, 0x26, 0xFF, 0x0E,
		0x11, 0x3E, 0x80, 0x32, 0xE2, 0x0C, 0x3E, 0xF3, 0xE2, 0x32, 0x3E, 0x77, 0x77, 0x3E, 0xFC, 0xE0,
		0x47, 0x11, 0x04, 0x01, 0x21, 0x10, 0x80, 0x1A, 0xCD, 0x95, 0x00, 0xCD, 0x96, 0x00, 0x13, 0x7B,
		0xFE, 0x34, 0x20, 0xF3, 0x11, 0xD8, 0x00, 0x06, 0x08, 0x1A, 0x13, 0x22, 0x23, 0x05, 0x20, 0xF9,
		0x3E, 0x19, 0xEA, 0x10, 0x99, 0x21, 0x2F, 0x99, 0x0E, 0x0C, 0x3D, 0x28, 0x08, 0x32, 0x0D, 0x20,
		0xF9, 0x2E, 0x0F, 0x18, 0xF3, 0x67, 0x3E, 0x64, 0x57, 0xE0, 0x42, 0x3E, 0x91, 0xE0, 0x40, 0x04,
		0x1E, 0x02, 0x0E, 0x0C, 0xF0, 0x44, 0xFE, 0x90, 0x20, 0xFA, 0x0D, 0x20, 0xF7, 0x1D, 0x20, 0xF2,
		0x0E, 0x13, 0x24, 0x7C, 0x1E, 0x83, 0xFE, 0x62, 0x28, 0x06, 0x1E, 0xC1, 0xFE, 0x64, 0x20, 0x06,
		0x7B, 0xE2, 0x0C, 0x3E, 0x87, 0xF2, 0xF0, 0x42, 0x90, 0xE0, 0x42, 0x15, 0x20, 0xD2, 0x05, 0x20,
		0x4F, 0x16, 0x20, 0x18, 0xCB, 0x4F, 0x06, 0x04, 0xC5, 0xCB, 0x11, 0x17, 0xC1, 0xCB, 0x11, 0x17,
		0x05, 0x20, 0xF5, 0x22, 0x23, 0x22, 0x23, 0xC9, 0xCE, 0xED, 0x66, 0x66, 0xCC, 0x0D, 0x00, 0x0B,
		0x03, 0x73, 0x00, 0x83, 0x00, 0x0C, 0x00, 0x0D, 0x00, 0x08, 0x11, 0x1F, 0x88, 0x89, 0x00, 0x0E,
		0xDC, 0xCC, 0x6E, 0xE6, 0xDD, 0xDD, 0xD9, 0x99, 0xBB, 0xBB, 0x67, 0x63, 0x6E, 0x0E, 0xEC, 0xCC,
		0xDD, 0xDC, 0x99, 0x9F, 0xBB, 0xB9, 0x33, 0x3E, 0x3c, 0x42, 0xB9, 0xA5, 0xB9, 0xA5, 0x42, 0x4C,
		0x21, 0x04, 0x01, 0x11, 0xA8, 0x00, 0x1A, 0x13, 0xBE, 0x20, 0xFE, 0x23, 0x7D, 0xFE, 0x34, 0x20,
		0xF5, 0x06, 0x19, 0x78, 0x86, 0x23, 0x05, 0x20, 0xFB, 0x86, 0x20, 0xFE, 0x3E, 0x01, 0xE0, 0x50,
	}
	mem.bootstrapDone = true
	mem.IOPort[0x00] = 0xFF // TIMA
	mem.IOPort[0x05] = 0x00 // TIMA
	mem.IOPort[0x06] = 0x00 // TMA
	mem.IOPort[0x07] = 0x00 // TAC
	mem.IOPort[0x10] = 0x80 // NR10
	mem.IOPort[0x11] = 0xBF // NR11
	mem.IOPort[0x12] = 0xF3 // NR12
	mem.IOPort[0x14] = 0xBF // NR14
	mem.IOPort[0x16] = 0x3F // NR21
	mem.IOPort[0x17] = 0x00 // NR22
	mem.IOPort[0x19] = 0xBF // NR24
	mem.IOPort[0x1A] = 0x7F // NR30
	mem.IOPort[0x1B] = 0xFF // NR31
	mem.IOPort[0x1C] = 0x9F // NR32
	mem.IOPort[0x1E] = 0xBF // NR33
	mem.IOPort[0x20] = 0xFF // NR41
	mem.IOPort[0x21] = 0x00 // NR42
	mem.IOPort[0x22] = 0x00 // NR43
	mem.IOPort[0x23] = 0xBF // NR30
	mem.IOPort[0x24] = 0x77 // NR50
	mem.IOPort[0x25] = 0xF3 // NR51
	if true {               // GB
		mem.IOPort[0x26] = 0xF1 // NR52
	} else { // TODO :SGB
		mem.IOPort[0x26] = 0xF0 // NR52
	}
	mem.IOPort[0x40] = 0x91 // LCDC
	mem.IOPort[0x42] = 0x00 // SCY
	mem.IOPort[0x43] = 0x00 // SCX
	mem.IOPort[0x45] = 0x00 // LYC
	mem.IOPort[0x47] = 0xFC // BGP
	mem.IOPort[0x48] = 0xFF // OBP0
	mem.IOPort[0x49] = 0xFF // OBP1
	mem.IOPort[0x4A] = 0x00 // WY
	mem.IOPort[0x4B] = 0x00 // WX
	mem.IOPort[0xFF] = 0x00 // IE

	mem.cartridge = CartrigdeInit(cartridgeUrl)
	mem.gb = gb

	return &mem
}

func (mem *memory) read8bit(addr uint16) uint8 {
	var val uint8

	if !mem.bootstrapDone && addr < 0x100 {
		val = mem.bios[addr]
		if (addr) == 0xFF {
			mem.bootstrapDone = true
		}
	} else if addr < 0x8000 {
		//val = mem.ROM[addr]
		val = mem.cartridge.data[addr]
	} else if (addr >= 0x8000) && (addr < 0xA000) {
		val = mem.VRAM[addr-0x8000]
	} else if (addr >= 0xA000) && (addr < 0xE000) {
		val = mem.RAM[addr-0xA000]
	} else if (addr >= 0xE000) && (addr < 0xFE00) {
		val = mem.EchoRAM[addr-0xE000]
	} else if (addr >= 0xFE00) && (addr < 0xFEA0) {
		val = mem.OAM[addr-0xFE00]
	} else if addr == 0xFF00 {
		val = mem.gb.readInput(mem.IOPort[addr-0xFF00])
		fmt.Printf("IO:%02X\n", val)
	} else if (addr >= 0xFF00) && (addr < 0xFF80) {
		val = mem.IOPort[addr-0xFF00]
	} else if (addr >= 0xFF80) && (addr <= 0xFFFF) {
		val = mem.IOPort[addr-0xFF00]
	} else {
		val = 0
		log.Printf("Reading Address Invalid : %02X \n", addr)
	}
	return val
}

func (mem *memory) read16bit(addr uint16) uint16 {
	val := uint16(mem.read8bit(addr))
	val = val | uint16(mem.read8bit(addr+1))<<8
	return val
}

func (mem *memory) write8bit(addr uint16, val uint8) {
	if addr < 0x8000 {
		mem.cartridge.data[addr] = val
	} else if (addr >= 0x8000) && (addr < 0xA000) {
		mem.VRAM[addr-0x8000] = val
	} else if (addr >= 0xA000) && (addr < 0xE000) {
		mem.RAM[addr-0xA000] = val
	} else if (addr >= 0xE000) && (addr < 0xFE00) {
		mem.EchoRAM[addr-0xE000] = val
		mem.write8bit(addr-0x2000, val)
	} else if (addr >= 0xFE00) && (addr < 0xFEA0) {
		mem.OAM[addr-0xFE00] = val
	} else if addr == TMC {
		currentfreq := mem.gb.getClockFreq()
		mem.IOPort[TMC-0xFF00] = val
		newfreq := mem.gb.getClockFreq()

		if currentfreq != newfreq {
			mem.gb.setClockFreq()
		}
	} else if addr == DIVIDER {
		mem.IOPort[addr-0xFF00] = 0
	} else if addr == 0xFF44 {
		mem.IOPort[addr-0xFF00] = 0
	} else if addr == 0xFF46 {
		mem.doDMATransfer(val)
	} else if (addr >= 0xFF00) && (addr < 0xFF80) {
		mem.IOPort[addr-0xFF00] = val
	} else if (addr >= 0xFF80) && (addr <= 0xFFFF) { // oupsi
		mem.IOPort[addr-0xFF00] = val
	} else {
		log.Printf("Writing Address Invalid : %02X \n", addr)
	}
}

func (mem *memory) write16bit(addr uint16, val uint16) {
	mem.write8bit(addr, uint8(val))
	mem.write8bit(addr+1, uint8(val>>8))
}

func (mem *memory) doDMATransfer(data uint8) {
	var addr uint16 = uint16(data) << 8
	for i := 0; i < 0xA0; i++ {
		mem.write8bit(0xFE00+uint16(i), mem.read8bit(addr+uint16(i)))
	}
}
