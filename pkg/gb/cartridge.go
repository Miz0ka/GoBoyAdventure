package gb

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Miz0ka/GoBoyAdventure/pkg/byteUtils"
)

type Cartridge struct {
	url   string
	title string
	mbc   uint8

	data           []byte
	currentRomBank int
	ram            [0x8000]byte
	currentRamBank int
	enableRam      bool
	romBanking     bool
}

func CartrigdeInit(url string) *Cartridge {
	/*file, err := os.Open(url)
	if err != nil {
		log.Fatal(err)
	}*/
	//cart.data = make([]byte, 0x10000)
	var cart Cartridge
	fmt.Printf("ROM: %s\n", url)
	cart.url = url
	var err error
	cart.data, err = ioutil.ReadFile(url)
	if err != nil {
		log.Fatal(err)
	}
	cart.title = string(cart.data[0x0134:0x0143])
	if false { //CGB
		_ = cart.data[0x0143]
		// CGB flag :
		// 0x80 : Game supports CGB functions, but works on old gameboys also
		// 0xC0 : Game works on CGB only (physically the same as 80h)
	}
	if false { //SGB
		_ = cart.data[0x0146] // 0x00 -> No SGB function | 0x03 -> SGB function supported
		_ = cart.data[0x014B] // Need to be at 0x014B
	}

	if cart.data[0x0147] == 0x00 {
		log.Println("Japanese Cartbridge")
	} else if cart.data[0x0147] == 0x01 {
		log.Println("Non-Japanese Cartbridge")
	} else {
		log.Println("Unknow Cartbridge")
	}

	_ = cart.data[0x0148] // TODO : ROM size
	_ = cart.data[0x0149] // TODO : RAM size

	log.Printf("%d \n", cart.data[0x0147])

	switch cart.data[0x0147] { // Type of Cartridge
	case 0:
		log.Println("TYPE: ROM ONLY")
		cart.mbc = 0
		break
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		log.Println("TYPE: MBC1")
		cart.mbc = 1
		break
	case 4:
		fallthrough
	case 5:
		fallthrough
	case 6:
		log.Println("TYPE: MBC2")
		cart.mbc = 2
		break
	default:
		log.Println("This type of cartridge isn't manage by this emulator")
		break
	}

	cart.currentRomBank = 1
	cart.currentRamBank = 0

	log.Printf("Check Sum result : header -> 0x%02X body -> 0x%04X", cart.headerCheckSum(), cart.bodyCheckSum())

	return &cart
}

func (cart *Cartridge) headerCheckSum() uint8 {
	var x int = 0
	for i := 0x0134; i < 0x014D; i++ {
		x = x - int(uint8(cart.data[i])) - 1
	}
	return uint8(x)
}

func (cart *Cartridge) bodyCheckSum() uint16 {
	var x int = 0
	for i := 0x0000; i < len(cart.data); i++ {
		if i != 0x014E && i != 0x014F {
			x += int(uint8(cart.data[i]))
		}
	}
	return uint16(x)
}

func (cart *Cartridge) read8bit(addr uint16) uint8 {

	if (addr >= 0x4000) && (addr <= 0x7FFF) {
		addr = uint16(addr-0x4000) + uint16(cart.currentRomBank*0x4000)
		return uint8(cart.data[addr])
	} else if (addr >= 0xA000) && (addr <= 0xBFFF) {
		addr = uint16(addr-0xA000) + uint16(cart.currentRamBank*0x2000)
		return uint8(cart.ram[addr])
	}

	return uint8(cart.data[addr])
}

func (cart *Cartridge) write8bit(addr uint16, val uint8) {
	if addr < 0x8000 {
		cart.handleBanking(addr, val)
	} else if (addr >= 0xA000) && (addr <= 0xC000) {
		if cart.enableRam {
			addr = uint16(addr-0xA000) + uint16(cart.currentRamBank*0x2000)
			cart.ram[addr] = val
		}
	}
}

func (cart *Cartridge) handleBanking(addr uint16, val uint8) {
	if addr < 0x2000 {
		if cart.mbc > 0 {
			cart.doRAMBankEnable(addr, val)
		}
	} else if (addr >= 0x200) && (addr < 0x4000) {
		if cart.mbc > 0 {
			cart.doChangeLoROMBank(val)
		}
	} else if (addr >= 0x4000) && (addr < 0x6000) {
		if cart.mbc == 1 {
			if cart.romBanking {
				cart.doChangeHiROMBank(val)
			} else {
				cart.doChangeRAMBankChange(val)
			}
		}
	} else if (addr >= 0x6000) && (addr < 0x8000) {
		if cart.mbc == 1 {
			cart.doChangeROMRAMMode(val)
		}
	}
}

func (cart *Cartridge) doRAMBankEnable(addr uint16, val uint8) {
	if cart.mbc == 2 {
		if byteUtils.TestBit(uint8(addr), 4) {
			return
		}
	}
	if val&0xF == 0xA {
		cart.enableRam = true
	} else if val&0xF == 0x0 {
		cart.enableRam = false
	}
}

func (cart *Cartridge) doChangeLoROMBank(val uint8) {
	if cart.mbc == 2 {
		cart.currentRomBank = int(val & 0xF)
		if cart.currentRomBank == 0 {
			cart.currentRomBank++
		}
	} else {
		cart.currentRomBank &= 0xE0
		cart.currentRomBank |= int(val & 0x1F)
		if cart.currentRomBank == 0 {
			cart.currentRomBank++
		}
	}
}

func (cart *Cartridge) doChangeHiROMBank(val uint8) {
	cart.currentRomBank &= 0x1F
	cart.currentRomBank |= int(val & 0xE0)
	if cart.currentRomBank == 0 {
		cart.currentRomBank++
	}
}

func (cart *Cartridge) doChangeRAMBankChange(val uint8) {
	cart.currentRamBank = int(val & 0x03)
}

func (cart *Cartridge) doChangeROMRAMMode(val uint8) {
	data := val & 0x1
	cart.romBanking = (data == 0)
	if cart.romBanking {
		cart.currentRamBank = 0
	}
}
