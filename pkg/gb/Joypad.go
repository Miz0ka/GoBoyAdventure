package gb

import (
	"github.com/Miz0ka/GoBoyAdventure/pkg/byteUtils"
)

const (
	INPUT_A = iota
	INPUT_B
	INPUT_SELECT
	INPUT_START
	INPUT_RIGHT
	INPUT_LEFT
	INPUT_UP
	INPUT_DOWN
)

func (gb *Gameboy) KeyReleased(key int) {
	if key >= INPUT_RIGHT {
		gb.directionKeys = gb.cpu.ulaSet(byte(key-INPUT_RIGHT), uint8(gb.directionKeys))
	} else {
		gb.buttonkeys = gb.cpu.ulaSet(byte(key), uint8(gb.buttonkeys))
	}
}

func (gb *Gameboy) KeyPressed(key int) {
	if key >= INPUT_RIGHT {
		gb.directionKeys = gb.cpu.ulaRes(byte(key-INPUT_RIGHT), uint8(gb.directionKeys))
	} else {
		gb.buttonkeys = gb.cpu.ulaRes(byte(key), uint8(gb.buttonkeys))
	}
	gb.requestInterrupts(4)
}

func (gb *Gameboy) readInput(flag uint8) byte {
	if byteUtils.TestBit(flag, 4) { //Directions
		return gb.buttonkeys | flag | 0xC0
	} else if byteUtils.TestBit(flag, 5) { // Button
		return gb.directionKeys | flag | 0xC0
	}
	return 0xCF | flag
}
