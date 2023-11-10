package ui

import "github.com/veandco/go-sdl2/sdl"

type GameboyWindow struct {
	sdlWindow   *sdl.Window
	sdlRenderer *sdl.Renderer
	sdlTexture  *sdl.Texture
	sdlScreen   *sdl.Surface
}

func InitGameBoyWindow() (*GameboyWindow, error) {
	sdlWindow, sdlRenderer, err := sdl.CreateWindowAndRenderer(gameBoyScreenWidth, gameBoyScreenHeight, 0)
	if err != nil {
		return nil, err
	}

	return &GameboyWindow{
		sdlWindow:   sdlWindow,
		sdlRenderer: sdlRenderer,
	}, nil
}
