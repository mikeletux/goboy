package main

import (
	"flag"
	"github.com/mikeletux/goboy/pkg/bus"
	"github.com/mikeletux/goboy/pkg/cart"
	"github.com/mikeletux/goboy/pkg/cpu"
	"github.com/mikeletux/goboy/pkg/lcd"
	"github.com/mikeletux/goboy/pkg/log"
	"github.com/veandco/go-sdl2/sdl"
	"sync"
)

func main() {
	romCartridgePath := flag.String("romPath", "", "Path to the Gameboy rom to load")
	flag.Parse()

	// Build stdout logger
	logger, err := log.NewBuiltinStdoutLogger(true, "/home/mikeletux/development/goboy/log/goboy.log")
	if err != nil {
		panic(err)
	}

	// Build cartridge
	cartridge, err := cart.NewCartridge(*romCartridgePath, logger)
	if err != nil {
		logger.Fatal(err)
	}

	// Build memory bus
	memoryBus := bus.NewBus(cartridge, logger)

	// Build CPU
	gbCpu := cpu.Init(memoryBus, logger)

	die := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(die <-chan bool) {
		defer wg.Done()
		for {
			select {
			case <-die:
				return
			default:
				gbCpu.Step()
			}
		}
	}(die)

	// Build UI
	gbScreen := lcd.NewGameboyScreen()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				die <- true
				wg.Wait()
				gbScreen.DestroyWindow()
				return
			}
		}
	}
}
