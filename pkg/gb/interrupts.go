package gb

import (
	"github.com/Miz0ka/GoBoyAdventure/pkg/byteUtils"
)

const (
	IF         = 0xFF0F
	IE         = 0xFFFF
	INTVBLANK  = 0x40
	INTLCDSTAT = 0x48
	INTTIMER   = 0x50
	INTJOYPAD  = 0x60
)

// IF/IE Structure
// Bit 0 : V-Blank
// Bit 1 : LCD Stat
// Bit 2 : Timer
// Bit 3 : Serial
// Bit 4 : Joypad

func (gb *Gameboy) requestInterrupts(bit uint8) {
	regval := gb.memory.read8bit(IF)
	regval = gb.cpu.ulaSet(bit, regval)
	gb.memory.write8bit(IF, regval)
}

func (gb *Gameboy) checkInterrupts() {
	if gb.cpu.isInterruptsEnable() {
		regIF := gb.memory.read8bit(IF) // IF : Interrupt Flag
		regIE := gb.memory.read8bit(IE) // IE : Interrupt Enable
		//fmt.Printf("IF> %02X IE> %02X == %02X \n", regIF, regIE, regIF&regIE)
		regIF = regIF & regIE
		var i uint8
		if regIF > 0 {
			for i = 0; i < 5; i++ {
				if byteUtils.TestBit(regIF, int(i)) {
					gb.doInterrupts(i)
					return
				}
			}
		}
	}
}

func (gb *Gameboy) doInterrupts(bit uint8) {
	gb.cpu.interruptsEnable = false
	flags := gb.memory.read8bit(IF)
	flags = gb.cpu.ulaRes(bit, flags)
	gb.memory.write8bit(IF, flags)
	gb.push(gb.cpu.getPC())
	switch bit {
	case 0:
		gb.cpu.setPC(INTVBLANK)
		break
	case 1:
		gb.cpu.setPC(INTLCDSTAT)
		break
	case 2:
		gb.cpu.setPC(INTTIMER)
		break
	// 3 -> Serial Transfert
	case 4:
		gb.cpu.setPC(INTJOYPAD)
		break
	}
}
