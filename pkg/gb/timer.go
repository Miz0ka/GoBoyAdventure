package gb

const (
	DIVIDER = 0xFF04
	TIMA    = 0xFF05
	TMA     = 0xFF06
	TMC     = 0xFF07
)

type timer struct {
}

func timerInit() *timer {
	var tmr timer
	return &tmr
}

func (gb *Gameboy) isClockEnabled() bool {
	val := gb.memory.read8bit(TMC)
	return val&0x2 != 0
}

func (gb *Gameboy) getClockFreq() uint8 {
	return gb.memory.read8bit(TMC) & 0x3
}

func (gb *Gameboy) setClockFreq() {
	// CPU 4194304
	freq := gb.getClockFreq()
	switch freq {
	case 0x00:
		gb.timerCounter = 1024
		break // freq 4096
	case 0x01:
		gb.timerCounter = 16
		break // freq 262144
	case 0x02:
		gb.timerCounter = 64
		break // freq 65536
	case 0x03:
		gb.timerCounter = 256
		break // freq 16382
	}
}

func (gb *Gameboy) updateTimers(cycles int) {
	gb.doDividerRegister(cycles)

	// the clock must be enabled to update the clock
	if gb.isClockEnabled() {
		gb.timerCounter -= cycles

		// enough cpu clock cycles have happened to update the timer
		if gb.timerCounter <= 0 {
			// reset m_TimerTracer to the correct value
			gb.setClockFreq()

			// timer about to overflow
			if gb.memory.read8bit(TIMA) == 255 {
				val := gb.memory.read8bit(TMA)
				gb.memory.write8bit(TIMA, val)
				gb.requestInterrupts(2)
			} else {
				val := gb.memory.read8bit(TIMA) + 1
				gb.memory.write8bit(TIMA, val)
			}
		}
	}
}

func (gb *Gameboy) doDividerRegister(cycles int) {
	gb.dividerRegister += cycles
	if gb.dividerRegister >= 0xFF {
		gb.dividerRegister = 0
		gb.memory.IOPort[0x04]++
	}
}
