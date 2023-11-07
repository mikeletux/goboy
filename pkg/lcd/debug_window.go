package lcd

import (
	"github.com/mikeletux/goboy/pkg/bus"
	"github.com/mikeletux/goboy/pkg/log"
	"github.com/veandco/go-sdl2/sdl"
	"unsafe"
)

type GameboyDebugWindow struct {
	sdlWindow   *sdl.Window
	sdlRenderer *sdl.Renderer
	sdlTexture  *sdl.Texture
	sdlScreen   *sdl.Surface

	logger log.Logger
	bus    bus.DataBusInterface
}

func InitGameBoyDebugWindow(logger log.Logger, bus bus.DataBusInterface) (*GameboyDebugWindow, error) {
	sdlWindow, sdlRenderer, err := sdl.CreateWindowAndRenderer(16*8*scale, 32*8*scale, 0)
	if err != nil {
		return nil, err
	}

	sdlTexture, err := sdlRenderer.CreateTexture(
		sdl.PIXELFORMAT_ABGR8888,
		sdl.TEXTUREACCESS_STREAMING,
		(16*8*scale)+(16*scale),
		(32*8*scale)+(64*scale),
	)

	if err != nil {
		return nil, err
	}

	sdlSurface, err := sdl.CreateRGBSurface(
		0,
		(16*8*scale)+(16*scale),
		(32*8*scale)+(64*scale),
		32,
		0x00FF0000,
		0x0000FF00,
		0x000000FF,
		0xFF000000,
	)

	if err != nil {
		return nil, err
	}

	return &GameboyDebugWindow{
		sdlWindow:   sdlWindow,
		sdlRenderer: sdlRenderer,
		sdlTexture:  sdlTexture,
		sdlScreen:   sdlSurface,
		logger:      logger,
		bus:         bus,
	}, nil
}

func (g *GameboyDebugWindow) updateWindow() {
	xDraw := 0
	yDraw := 0
	tileNum := 0

	var rc sdl.Rect
	rc.X = 0
	rc.Y = 0
	rc.W = g.sdlScreen.W
	rc.H = g.sdlScreen.H
	g.sdlScreen.FillRect(&rc, 0xFF111111) // HANDLE ERROR HERE

	var addr uint16 = 0x8000

	//384 tiles, 24 x 16
	for y := 0; y < 24; y++ {
		for x := 0; x < 16; x++ {
			displayTile(g.sdlScreen, g.bus, addr, uint16(tileNum), xDraw+(x*scale), yDraw+(y*scale))
			xDraw += 8 * scale
			tileNum++
		}

		yDraw += 8 * scale
		xDraw = 0
	}

	pixels := g.sdlScreen.Pixels()
	g.sdlTexture.Update(nil, unsafe.Pointer(&pixels[0]), int(g.sdlScreen.Pitch))
	g.sdlRenderer.Clear()
	g.sdlRenderer.Copy(g.sdlTexture, nil, nil)
	g.sdlRenderer.Present()
}
