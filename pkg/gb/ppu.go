package gb

import (
	"github.com/Miz0ka/GoBoyAdventure/pkg/byteUtils"
)

const (
	ScreenWidth  = 160
	ScreenHeight = 144
)

type PPU struct {
	mem   *memory
	cycle int
	mode  int
	line  int

	screen [ScreenWidth][ScreenHeight][3]int

	scanlineCounter int
}

func ppuInit(mem *memory) *PPU {
	var ppu PPU
	ppu.screen = [ScreenWidth][ScreenHeight][3]int{}
	ppu.mem = mem
	return &ppu
}

func (gb *Gameboy) updateGraphics(cycles int) {
	//gb.setLCDStatus()

	if !gb.isLCDEnabled() {
		return
	}

	gb.ppu.scanlineCounter -= cycles

	if gb.ppu.scanlineCounter <= 0 {

		// time to move onto next scanline
		gb.memory.IOPort[0x44]++
		currentline := gb.memory.read8bit(0xFF44)

		gb.ppu.scanlineCounter += 456

		// we have entered vertical blank period
		if currentline == 144 {
			gb.requestInterrupts(0)
		} else if currentline > 153 { // if gone past scanline 153 reset to 0
			gb.memory.IOPort[0x44] = 0
		} else if currentline < 144 { // draw the current scanline
			gb.drawScanLine()
		}
	}
}

func (gb *Gameboy) isLCDEnabled() bool {
	return byteUtils.TestBit(gb.memory.read8bit(0xFF40), 7)
}

func (gb *Gameboy) setLCDStatus() {
	status := gb.memory.read8bit(0xFF41)
	if !gb.isLCDEnabled() {
		// set the mode to 1 during lcd disabled and reset scanline
		gb.ppu.scanlineCounter = 456
		gb.memory.write8bit(0xFF44, 0)
		status &= 252
		status = gb.cpu.ulaSet(0, status)
		gb.memory.write8bit(0xFF41, status)
		return
	}

	currentline := gb.memory.read8bit(0xFF44)
	currentmode := status & 0x3

	var mode uint8 = 0
	reqInt := false

	// in vblank so set mode to 1
	if currentline >= 144 {
		mode = 1
		status = gb.cpu.ulaSet(0, status)
		status = gb.cpu.ulaRes(1, status)
		reqInt = byteUtils.TestBit(status, 4)
	} else {
		mode2bounds := 456 - 80
		mode3bounds := mode2bounds - 172

		// mode 2
		if gb.ppu.scanlineCounter >= mode2bounds {
			mode = 2
			status = gb.cpu.ulaSet(1, status)
			status = gb.cpu.ulaRes(0, status)
			reqInt = byteUtils.TestBit(status, 5)
		} else if gb.ppu.scanlineCounter >= mode3bounds {
			// mode 3
			mode = 3
			status = gb.cpu.ulaSet(1, status)
			status = gb.cpu.ulaSet(0, status)
		} else {
			// mode 0
			mode = 0
			status = gb.cpu.ulaRes(1, status)
			status = gb.cpu.ulaRes(0, status)
			reqInt = byteUtils.TestBit(status, 3)
		}
	}

	// just entered a new mode so request interupt
	if reqInt && (mode != currentmode) {
		gb.requestInterrupts(1)
	}

	// check the conincidence flag
	if currentline == gb.memory.read8bit(0xFF45) {
		status = gb.cpu.ulaSet(2, status)
		if byteUtils.TestBit(status, 6) {
			gb.requestInterrupts(1)
		}
	} else {
		status = gb.cpu.ulaRes(2, status)
	}
	gb.memory.write8bit(0xFF41, status)
}

func (gb *Gameboy) drawScanLine() {
	control := gb.memory.read8bit(0xFF40)
	if byteUtils.TestBit(control, 0) {
		gb.renderTiles(control)
	}
	if byteUtils.TestBit(control, 1) {
		gb.renderSprites(control)
	}
}

func (gb *Gameboy) renderTiles(lcdControl uint8) {
	var tileData uint16 = 0
	var backgroundMemory uint16 = 0
	var unsig bool = true

	// where to draw the visual area and the window
	scrollY := gb.memory.read8bit(0xFF42)
	scrollX := gb.memory.read8bit(0xFF43)
	windowY := gb.memory.read8bit(0xFF4A)
	windowX := gb.memory.read8bit(0xFF4B) - 7

	var usingWindow bool = false

	// is the window enabled?
	if byteUtils.TestBit(lcdControl, 5) {
		// is the current scanline we're drawing
		// within the windows Y pos?,
		if windowY <= gb.memory.read8bit(0xFF44) {
			usingWindow = true
		}
	}

	// which tile data are we using?
	if byteUtils.TestBit(lcdControl, 4) {
		tileData = 0x8000
	} else {
		// IMPORTANT: This memory region uses signed
		// bytes as tile identifiers
		tileData = 0x8800
		unsig = false
	}

	// which background mem?
	if false == usingWindow {
		if byteUtils.TestBit(lcdControl, 3) {
			backgroundMemory = 0x9C00
		} else {
			backgroundMemory = 0x9800
		}
	} else {
		// which window memory?
		if byteUtils.TestBit(lcdControl, 6) {
			backgroundMemory = 0x9C00
		} else {
			backgroundMemory = 0x9800
		}
	}

	var yPos byte = 0

	// yPos is used to calculate which of 32 vertical tiles the
	// current scanline is drawing
	if !usingWindow {
		yPos = scrollY + gb.memory.read8bit(0xFF44)
	} else {
		yPos = gb.memory.read8bit(0xFF44) - windowY
	}

	// which of the 8 vertical pixels of the current
	// tile is the scanline on?
	var tileRow uint16 = uint16(yPos/8) * 32

	// time to start drawing the 160 horizontal pixels
	// for this scanline
	for pixel := uint8(0); pixel < 160; pixel++ {
		var xPos uint8 = pixel + scrollX

		// translate the current x pos to window space if necessary
		if usingWindow {
			if pixel >= windowX {
				xPos = pixel - windowX
			}
		}

		// which of the 32 horizontal tiles does this xPos fall within?
		var tileCol uint16 = uint16(xPos / 8)
		var tileNum int16

		// get the tile identity number. Remember it can be signed
		// or unsigned
		var tileAddrss uint16 = backgroundMemory + tileRow + tileCol

		// deduce where this tile identifier is in memory. Remember i
		// shown this algorithm earlier
		var tileLocation uint16 = tileData

		if unsig {
			tileNum = int16(uint16(gb.memory.read8bit(tileAddrss)))
			tileLocation += uint16(tileNum * 16)
		} else {
			tileNum = int16(int8(gb.memory.read8bit(tileAddrss)))
			tileLocation += uint16((tileNum + 128) * 16)
		}

		// find the correct vertical line we're on of the
		// tile to get the tile data
		//from in memory
		var line uint8 = (yPos % 8) * 2 // each vertical line takes up two bytes of memory
		var data1 uint8 = gb.memory.read8bit(tileLocation + uint16(line))
		var data2 uint8 = gb.memory.read8bit(tileLocation + uint16(line) + 1)

		//fmt.Printf("d1: %04X d2: %04X \n", tileLocation+uint16(line), tileLocation+uint16(line)+1)

		// pixel 0 in the tile is it 7 of data 1 and data2.
		// Pixel 1 is bit 6 etc..
		var colourBit byte = byte(int8((xPos%8)-7) * -1)

		// combine data 2 and data 1 to get the colour id for this pixel
		// in the tile
		//var colourNum uint8 = byteUtils.Val(data2, colourBit)
		colourNum := (byteUtils.Val(data2, colourBit) << 1) | byteUtils.Val(data1, colourBit)
		//colourNum <<= 1
		//colourNum |= byteUtils.Val(data1, byte(colourBit))

		// now we have the colour id get the actual
		// colour from palette 0xFF47
		col := gb.getColour(colourNum, 0xFF47)
		var red int = 0
		var green int = 0
		var blue int = 0

		// setup the RGB values
		switch col {
		case WHITE:
			red = 255
			green = 255
			blue = 255
			break
		case LIGHT_GRAY:
			red = 0xCC
			green = 0xCC
			blue = 0xCC
			break
		case DARK_GRAY:
			red = 0x77
			green = 0x77
			blue = 0x77
			break
		}

		var finaly uint8 = gb.memory.read8bit(0xFF44)

		// safety check to make sure what im about
		// to set is int the 160x144 bounds
		if (finaly < 0) || (finaly > 143) || (pixel < 0) || (pixel > 159) {
			continue
		}

		gb.ppu.screen[pixel][finaly][0] = red
		gb.ppu.screen[pixel][finaly][1] = green
		gb.ppu.screen[pixel][finaly][2] = blue
	}
}

func (gb *Gameboy) getColour(colourNum uint8, address uint16) color {
	var res color = WHITE
	palette := gb.memory.read8bit(address)
	//var hi uint8 = 0
	//var lo uint8 = 0

	hi := colourNum<<1 | 1
	lo := colourNum << 1

	// which bits of the colour palette does the colour id map to?
	/*switch colourNum {
	case 0:
		hi = 1
		lo = 0
		break
	case 1:
		hi = 3
		lo = 2
		break
	case 2:
		hi = 5
		lo = 4
		break
	case 3:
		hi = 7
		lo = 6
		break
	}*/

	// use the palette to get the colour
	var col color = 0
	col = color(byteUtils.Val(palette, byte(hi))) << 1
	col |= color(byteUtils.Val(palette, lo))

	// convert the game colour to emulator colour
	switch col {
	case 0:
		res = WHITE
		break
	case 1:
		res = LIGHT_GRAY
		break
	case 2:
		res = DARK_GRAY
		break
	case 3:
		res = BLACK
		break
	}

	return res
}

func (gb *Gameboy) renderSprites(lcdControl uint8) {
	var use8x16 bool = false
	if byteUtils.TestBit(lcdControl, 2) {
		use8x16 = true
	}
	var sprite uint8
	for sprite = 0; sprite < 40; sprite++ {
		// sprite occupies 4 bytes in the sprite attributes table
		var index uint16 = uint16(sprite * 4)
		var yPos uint16 = uint16(gb.memory.read8bit(0xFE00+index) - 16)
		var xPos uint16 = uint16(gb.memory.read8bit(0xFE00+index+1) - 8)
		var tileLocation uint16 = uint16(gb.memory.read8bit(0xFE00 + index + 2))
		var attributes uint8 = gb.memory.read8bit(0xFE00 + index + 3)

		var yFlip bool = byteUtils.TestBit(attributes, 6)
		var xFlip bool = byteUtils.TestBit(attributes, 5)

		var scanline uint16 = uint16(gb.memory.read8bit(0xFF44))

		var ysize uint16 = 8
		if use8x16 {
			ysize = 16
		}

		// does this sprite intercept with the scanline?
		if (scanline >= yPos) && (scanline < (yPos + ysize)) {
			var line int16 = int16(scanline - yPos)

			// read the sprite in backwards in the y axis
			if yFlip {
				line -= int16(ysize)
				line *= -1
			}

			line *= 2 // same as for tiles
			var dataAddress uint16 = uint16(int32(0x8000+(tileLocation*16)) + int32(line))
			var data1 uint8 = gb.memory.read8bit(dataAddress)
			var data2 uint8 = gb.memory.read8bit(dataAddress + 1)

			// its easier to read in from right to left as pixel 0 is
			// bit 7 in the colour data, pixel 1 is bit 6 etc...
			for tilePixel := 7; tilePixel >= 0; tilePixel-- {
				colourbit := tilePixel
				// read the sprite in backwards for the x axis
				if xFlip {
					colourbit -= 7
					colourbit *= -1
				}

				// the rest is the same as for tiles
				colourNum := byteUtils.Val(data2, byte(colourbit))
				colourNum <<= 1
				colourNum |= byteUtils.Val(data1, byte(colourbit))

				var colourAddress uint16
				if byteUtils.TestBit(attributes, 4) {
					colourAddress = 0xFF49
				} else {
					colourAddress = 0xFF48
				}
				col := gb.getColour(colourNum, colourAddress)

				// white is transparent for sprites.
				if col == WHITE {
					continue
				}
				red := 0
				green := 0
				blue := 0

				switch col {
				case WHITE:
					red = 255
					green = 255
					blue = 255
					break
				case LIGHT_GRAY:
					//red = 0xCC
					//green = 0xCC
					//blue = 0xCC
					red = 255
					green = 0
					blue = 0
					break
				case DARK_GRAY:
					//red = 0x77
					//green = 0x77
					//blue = 0x77
					red = 0
					green = 0
					blue = 255
					break
				}

				xPix := 0 - tilePixel
				xPix += 7

				pixel := xPos + uint16(xPix)

				// sanity check
				if (scanline < 0) || (scanline > 143) || (pixel < 0) || (pixel > 159) {
					continue
				}

				gb.ppu.screen[pixel][scanline][0] = red
				gb.ppu.screen[pixel][scanline][1] = green
				gb.ppu.screen[pixel][scanline][2] = blue
			}
		}
	}
}
