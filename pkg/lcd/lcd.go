package lcd

import "github.com/veandco/go-sdl2/sdl"

type Screen interface {
}

type GameboyScreen struct {
	window *sdl.Window
}

func NewGameboyScreen() *GameboyScreen {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	window.UpdateSurface()

	return &GameboyScreen{
		window: window,
	}
}

func (g *GameboyScreen) DestroyWindow() {
	g.window.Destroy()
	sdl.Quit()
}
