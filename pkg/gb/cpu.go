package gb

type Registers struct {
	// Value
	af register
	bc register
	de register
	hl register
	pc uint16 // program ptr
	sp uint16 // stack ptr
}

const (
	// F value (flag) :
	ZF = 0x80 // Zero 0x80 -> Set if last operation produced 0
	N  = 0x40 // Operation 0x40 -> Set if last operation addition/subtraction
	H  = 0x20 // Half-carry 0x20 -> Set if last operation result [lower half] overflowed 15
	CY = 0x10 // Carry 0x10 -> Set if last operation produced over flow
)

type CPU struct {
	registers        Registers
	halt             bool
	interruptsEnable bool
	cycles           int
}

func CPUinit() *CPU {
	var cpu CPU
	cpu.registers.pc = 0x100
	//cpu.registers.pc = 0x26BE
	cpu.registers.sp = 0xFFFE
	/*
		cpu.registers.pc = 0x000
		cpu.registers.af.setHiLo(0x01B0)
		cpu.registers.bc.setHiLo(0x0013)
		cpu.registers.de.setHiLo(0x00D8)
		cpu.registers.hl.setHiLo(0x014D)
	*/
	cpu.registers.af.setHiLo(0x01B0)
	cpu.registers.bc.setHiLo(0x0013)
	cpu.registers.de.setHiLo(0x00D8)
	cpu.registers.hl.setHiLo(0x014D)
	cpu.interruptsEnable = false
	//log.Printf("AF-> %d\n", cpu.registers.hl.getHiLo())
	return &cpu
}

func (cpu *CPU) isHalt() bool {
	return cpu.halt
}

func (cpu *CPU) setHalt(status bool) {
	cpu.halt = status
}

func (cpu *CPU) isInterruptsEnable() bool {
	return cpu.interruptsEnable
}

func (cpu *CPU) setInterruptsEnable(status bool) {
	cpu.interruptsEnable = status
}

func (cpu *CPU) isZF() bool {
	return ((cpu.registers.af.getLo() & ZF) == ZF)
}

func (cpu *CPU) isN() bool {
	return ((cpu.registers.af.getLo() & N) == N)
}

func (cpu *CPU) isH() bool {
	return ((cpu.registers.af.getLo() & H) == H)
}

func (cpu *CPU) isCY() bool {
	return ((cpu.registers.af.getLo() & CY) == CY)
}

func (cpu *CPU) setFlag(val bool, flag uint8) {
	fReg := cpu.registers.af.getLo()
	if val {
		fReg = fReg | flag
	} else {
		fReg = fReg &^ flag
	}
	cpu.registers.af.setLo(fReg)
}

func (cpu *CPU) setPC(addr uint16) {
	cpu.registers.pc = addr
}

func (cpu *CPU) getPC() uint16 {
	return cpu.registers.pc
}

func (cpu *CPU) setSP(addr uint16) {
	cpu.registers.sp = addr
}

func (cpu *CPU) getSP() uint16 {
	return cpu.registers.sp
}

func (gb *Gameboy) push(val uint16) {
	gb.cpu.registers.sp -= 2
	gb.memory.write16bit(gb.cpu.registers.sp, val)
}

func (gb *Gameboy) pop() uint16 {
	addr := gb.memory.read16bit(gb.cpu.registers.sp)
	gb.cpu.registers.sp += 2
	return addr
}

func (gb *Gameboy) cpuCall(addr uint16) {
	addrSp := gb.cpu.getPC()
	gb.push(addrSp)
	gb.cpu.setPC(addr)
}
