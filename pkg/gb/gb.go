package gb

type Gameboy struct {
	cpu    *CPU
	memory *memory
	ppu    *PPU

	opCode   [0x100]func()
	opCodeCB [0x100]func()

	timerCounter    int
	dividerRegister int

	directionKeys byte
	buttonkeys    byte

	Debug bool
}

const (
	MAXCYCLES = 69905 // 4194304 / 60 ( cpu cycles per sec / frame rate)
)

func (gb *Gameboy) Init(cartridgeUrl string) {
	gb.cpu = CPUinit()
	gb.memory = memoryInit(cartridgeUrl, gb)
	gb.ppu = ppuInit(gb.memory)

	gb.setClockFreq()
	gb.opCode = gb.genereOPDic()
	gb.opCodeCB = gb.genereOPCBDic()
	gb.directionKeys = 0xF
	gb.buttonkeys = 0xF
}

func (gb *Gameboy) Update() {
	var cycleTotal int = 0
	var cycleCurrent int = 0
	for cycleTotal < MAXCYCLES {
		cycleCurrent = gb.cpu.cycles
		cycleTotal += cycleCurrent

		//fmt.Printf(" debut Op %02X [%d] ", gb.cpu.getPC(), cycleCurrent)
		gb.executeNextOpcode()
		//fmt.Printf("LCD : %02X \n", gb.memory.read8bit(0xFF40))
		gb.updateTimers(cycleCurrent)
		gb.checkInterrupts()
		gb.updateGraphics(cycleCurrent)
		//time.Sleep(time.Second / 1000)
	}
}

func (gb *Gameboy) GetScreen() [ScreenWidth][ScreenHeight][3]int {
	return gb.ppu.screen
}

///////////// DEBUG FUNCT ////////////////////

func (gb *Gameboy) GetVRAM() [0x2000]byte {
	return gb.memory.VRAM
}

func (gb *Gameboy) GetCPUSTAT() [10]uint16 {
	var stat [10]uint16
	stat[0] = gb.cpu.getPC()
	stat[1] = gb.cpu.getSP()
	stat[2] = gb.cpu.registers.af.getHiLo()
	stat[3] = gb.cpu.registers.bc.value
	stat[4] = gb.cpu.registers.de.value
	stat[5] = gb.cpu.registers.hl.value
	stat[6] = uint16(gb.ppu.scanlineCounter)
	stat[7] = uint16(gb.buttonkeys)
	stat[8] = uint16(gb.directionKeys)

	return stat
}

func (gb *Gameboy) GetGameTitle() string {
	return gb.memory.cartridge.title
}

// halt = 1 cycle GB

// TODO
// INIT
// Graphics part
// MEMORY BANK
// Interrucp
// Input
// Sound
