package main

import (
	"flag"
	"github.com/mikeletux/goboy/pkg/bus"
	"github.com/mikeletux/goboy/pkg/cart"
	"github.com/mikeletux/goboy/pkg/cpu"
	"github.com/mikeletux/goboy/pkg/log"
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

	cartridge.LogCartridgeHeaderInfo()

	// Build memory bus
	memoryBus := bus.NewBus(cartridge, logger)

	// Build CPU
	gbCpu := cpu.Init(memoryBus, logger)

	for {
		gbCpu.Step()
	}
}
