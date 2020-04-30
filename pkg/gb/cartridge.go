package gb

import (
	"io/ioutil"
	"log"
)

type Cartridge struct {
	url   string
	title string
	data  []byte
}

func CartrigdeInit(url string) *Cartridge {
	/*file, err := os.Open(url)
	if err != nil {
		log.Fatal(err)
	}*/
	//cart.data = make([]byte, 0x10000)
	var cart Cartridge
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

	_ = cart.data[0x0147] // TODO : Type of Cartridge
	_ = cart.data[0x0148] // TODO : ROM size
	_ = cart.data[0x0149] // TODO : RAM size

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
