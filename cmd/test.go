package main

import (
	"fmt"
	"image/color"
	"strconv"
	"time"

	"github.com/Miz0ka/GoBoyAdventure/pkg/gb"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	var gbemu gb.Gameboy
	//gbemu.Init("../roms/Zoids Densetsu.GB")
	//gbemu.Init("../roms/Monster Maker.GB")
	//gbemu.Init("../roms/Doraemon.GB")
	//gbemu.Init("../roms/Tetris.GB")
	//gbemu.Init("../roms/01-special.GB")
	//gbemu.Init("../roms/02-interrupts.GB")
	//gbemu.Init("../roms/03-op sp,hl.GB") Ok
	gbemu.Init("../roms/04-op r,imm.GB") // Crash
	//gbemu.Init("../roms/05-op rp.GB") Ok
	//gbemu.Init("../roms/06-ld r,r.GB") Ok
	//gbemu.Init("../roms/07-jr,jp,call,ret,rst.GB") // Long
	//gbemu.Init("../roms/08-misc instrs.GB") // F1
	//gbemu.Init("../roms/09-op r,r.GB")
	//gbemu.Init("../roms/10-bit ops.GB") Ok
	//gbemu.Init("../roms/11-op a,(hl).GB")
	//gbemu.Init("../roms/bgbtest.GB")

	cfg := pixelgl.WindowConfig{
		Title:  gbemu.GetGameTitle(),
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	vramTxt := text.New(pixel.V(700, 500), basicAtlas)
	statText := text.New(pixel.V(850, 700), basicAtlas)

	imd := imdraw.New(nil)

	var stat [10]uint16
	stat[0] = 0x0000

	frameTime := time.Second / 100

	ticker := time.NewTicker(frameTime)
	start := time.Now()
	frames := 0
	for range ticker.C {

		if win.Closed() {
			return
		}

		win.Clear(colornames.Black)
		vramTxt.Clear()
		statText.Clear()
		imd.Clear()
		inputGB(win, &gbemu)
		gbemu.Update()

		for x, val := range gbemu.GetVRAM() {
			fmt.Fprintf(vramTxt, "%02X ", val)
			if (x+1)%16 == 0 {
				fmt.Fprintf(vramTxt, "\n")
			}
		}
		stat = gbemu.GetCPUSTAT()
		fmt.Fprintf(statText, "PC -> %04X \n", stat[0])
		//fmt.Printf("PC -> %04X ", stat[0])
		fmt.Fprintf(statText, "SP -> %04X \n", stat[1])
		//fmt.Printf("SP -> %04X ", stat[1])
		fmt.Fprintf(statText, "AF -> %04X \n", stat[2])
		//fmt.Printf("AF -> %04X ", stat[2])
		fmt.Fprintf(statText, "BC -> %04X \n", stat[3])
		//fmt.Printf("BC -> %04X ", stat[3])
		fmt.Fprintf(statText, "DE -> %04X \n", stat[4])
		//fmt.Printf("DE -> %04X ", stat[4])
		fmt.Fprintf(statText, "HL -> %04X \n", stat[5])
		//fmt.Printf("HL -> %04X ", stat[5])
		//fmt.Printf("Scanline -> %d \n", stat[6])
		fmt.Fprintf(statText, "Scanline -> %d \n", stat[6])
		fmt.Fprintf(statText, "Button -> %04b \n", stat[7])
		fmt.Fprintf(statText, "direction -> %04b \n", stat[8])

		screen := gbemu.GetScreen()

		for x, line := range screen {
			for y, pixels := range line {
				imd.Push(pixel.V(float64(x*5), float64(800-y*5)))
				imd.Push(pixel.V(float64((x+1)*5), float64(800-(y+1)*5)))
				imd.Color = pixel.ToRGBA(color.RGBA{uint8(pixels[0]), uint8(pixels[1]), uint8(pixels[2]), 255}) //(float64(pixels[0]), float64(pixels[1]), float64(pixels[2]))
				//fmt.Printf("%d, %d, %d", pixels[0], pixels[1], pixels[2])
				imd.Rectangle(0)
				/*var bck bloc
				bck.rect. = pixel.R(-50, -34, 50, -32)
				rect.Color = color
				rect.Draw(win)*/
			}
		}

		frames++
		imd.Draw(win)
		vramTxt.Draw(win, pixel.IM)
		statText.Draw(win, pixel.IM)
		win.Update()
		//time.Sleep(time.Millisecond * 1)
		since := time.Since(start)
		if since > time.Second {
			start = time.Now()
			win.SetTitle(strconv.Itoa(frames))
			frames = 0
		}
	}
}

func inputGB(win *pixelgl.Window, gbemu *gb.Gameboy) {
	if win.JustPressed(pixelgl.KeyZ) {
		gbemu.KeyPressed(gb.INPUT_A)
	}
	if win.JustReleased(pixelgl.KeyZ) {
		gbemu.KeyReleased(gb.INPUT_A)
	}

	if win.JustPressed(pixelgl.KeyX) {
		gbemu.KeyPressed(gb.INPUT_B)
	}
	if win.JustReleased(pixelgl.KeyX) {
		gbemu.KeyReleased(gb.INPUT_B)
	}

	if win.JustPressed(pixelgl.KeyA) {
		gbemu.KeyPressed(gb.INPUT_SELECT)
	}
	if win.JustReleased(pixelgl.KeyA) {
		gbemu.KeyReleased(gb.INPUT_SELECT)
	}

	if win.JustPressed(pixelgl.KeyS) {
		gbemu.KeyPressed(gb.INPUT_START)
	}
	if win.JustReleased(pixelgl.KeyS) {
		gbemu.KeyReleased(gb.INPUT_START)
	}

	if win.JustPressed(pixelgl.KeyRight) {
		gbemu.KeyPressed(gb.INPUT_RIGHT)
	}
	if win.JustReleased(pixelgl.KeyRight) {
		gbemu.KeyReleased(gb.INPUT_RIGHT)
	}

	if win.JustPressed(pixelgl.KeyLeft) {
		gbemu.KeyPressed(gb.INPUT_LEFT)
	}
	if win.JustReleased(pixelgl.KeyLeft) {
		gbemu.KeyReleased(gb.INPUT_LEFT)
	}

	if win.JustPressed(pixelgl.KeyUp) {
		gbemu.KeyPressed(gb.INPUT_UP)
	}
	if win.JustReleased(pixelgl.KeyUp) {
		gbemu.KeyReleased(gb.INPUT_UP)
	}

	if win.JustPressed(pixelgl.KeyDown) {
		gbemu.KeyPressed(gb.INPUT_DOWN)
	}
	if win.JustReleased(pixelgl.KeyDown) {
		gbemu.KeyReleased(gb.INPUT_DOWN)
	}
}
