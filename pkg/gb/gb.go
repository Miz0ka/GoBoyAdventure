package gb

type gameboy struct {
	cpu    *CPU
	memory *Memory

	opCode   [0x100]func()
	opCodeCB [0x100]func()
}

func (gb *gameboy) Init() {
	//Les restes
	gb.opCode = gb.genereOPDic()
	gb.opCodeCB = gb.genereOPCBDic()
}

func (gb *gameboy) Update() {
	for {
		gb.executeNextOpcode()
		//UpdateTimers(cycles)
		//UpdateGraphics(cycles)
		//DoInterupts()
	}
	//RenderScreen()
}
