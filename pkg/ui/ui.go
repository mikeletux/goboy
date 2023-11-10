package ui

import (
	"github.com/mikeletux/goboy/pkg/bus"
	"github.com/mikeletux/goboy/pkg/log"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	gameBoyScreenWidth  = 800
	gameBoyScreenHeight = 600
	scale               = 4
)

type GameboyScreen struct {
	window      *GameboyWindow
	debugWindow *GameboyDebugWindow
}

func NewGameboyScreen(logger log.Logger, bus bus.DataBusInterface) *GameboyScreen {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}

	/*	gameBoyWindow, err := InitGameBoyWindow()
		if err != nil {
			panic(err) // Handle better
		}*/

	gameBoyDebugWindow, err := InitGameBoyDebugWindow(logger, bus)
	if err != nil {
		panic(err) // Handle better
	}

	// Set one window aside the other one
	//	x, y := gameBoyWindow.sdlWindow.GetPosition()
	//	gameBoyDebugWindow.sdlWindow.SetPosition(x+gameBoyScreenWidth+10, y)

	return &GameboyScreen{
		//		window:      gameBoyWindow,
		debugWindow: gameBoyDebugWindow,
	}
}

func (g *GameboyScreen) UpdateUI() {
	g.debugWindow.updateWindow()
}

func (g *GameboyScreen) DestroyWindow() {
	g.window.sdlWindow.Destroy()
	sdl.Quit()
}

func displayTile(surface *sdl.Surface, bus bus.DataBusInterface, startLocation uint16, tileNum uint16, x, y int) {
	for tileY := uint16(0); tileY < 16; tileY += 2 {
		b1 := bus.BusRead(startLocation + tileNum*16 + tileY)
		b2 := bus.BusRead(startLocation + tileNum*16 + tileY + 1)

		for bit := 7; bit >= 0; bit-- {
			// hi := !!(b1 & (1 << bit)) << 1
			var hi byte
			if b1&(1<<bit) > 0 {
				hi = 1 << 1
			} else {
				hi = 0
			}
			// lo := !!(b2 & (1 << bit))

			var lo byte
			if b2&(1<<bit) > 0 {
				lo = 1
			} else {
				lo = 0
			}
			color := hi | lo

			rc := &sdl.Rect{
				X: int32(x + ((7 - bit) * scale)),
				Y: int32(y + (int(tileY) / 2 * scale)),
				W: scale,
				H: scale,
			}

			surface.FillRect(rc, getPixelColor(color))
		}
	}
}

func getPixelColor(code byte) uint32 {
	/*
		0xFFFFFFFF -> White
		0xFFAAAAAA -> Light grey
		0xFF5555555 -> Dark grey
		0xFF000000 -> Black
	*/
	colors := []uint32{0xFFFFFFFF, 0xFFAAAAAA, 0xFF555555, 0xFF000000}
	return colors[code]
}
