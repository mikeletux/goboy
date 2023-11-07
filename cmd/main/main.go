package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/mikeletux/goboy/pkg/bus"
	"github.com/mikeletux/goboy/pkg/cart"
	"github.com/mikeletux/goboy/pkg/config"
	"github.com/mikeletux/goboy/pkg/cpu"
	"github.com/mikeletux/goboy/pkg/lcd"
	"github.com/mikeletux/goboy/pkg/log"
	"os"
	"sync"
	"time"
)

func main() {
	configValues := configureEmulator()

	// Build stdout logger
	logger, err := log.NewBuiltinStdoutLogger(configValues.LogStdoutEnable, configValues.LogFileEnable, configValues.LogFilePath) // Add truncate!!! TO-DO
	if err != nil {
		panic(err)
	}

	// Build cartridge
	cartridge, err := cart.NewCartridge(configValues.RomPath, logger)
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
	gbScreen := lcd.NewGameboyScreen(logger, memoryBus)

	for {
		time.Sleep(1000 * time.Microsecond)
		gbScreen.UpdateUI()
	}
}

func configureEmulator() *config.Config {
	configFilePath := flag.String("configFilePath", "", "Path to the GoBoy config path")
	flag.Parse()

	configParser, err := config.NewConfigParser(*configFilePath)
	if err != nil {
		var confNotFoundErr *config.ConfigurationFileNotFound
		if errors.As(err, &confNotFoundErr) {
			fmt.Printf("%s", confNotFoundErr) // probably in the future show a window?
			os.Exit(-1)
		} else {
			panic(err)
		}
	}

	configValues, err := configParser.Parse()
	if err != nil {
		var parseInfoErr *config.ParsingError
		var missingConfigValuesErr *config.MissingConfigValuesError
		var romNotFoundErr *config.RomNotFoundError
		var dirWriteErr *config.LogWriteError

		if errors.As(err, &parseInfoErr) {
			fmt.Printf("%s", parseInfoErr) // probably in the future show a window?
			os.Exit(-1)
		} else if errors.As(err, &missingConfigValuesErr) {
			fmt.Printf("%s", missingConfigValuesErr) // probably in the future show a window?
			os.Exit(-1)
		} else if errors.As(err, &romNotFoundErr) {
			fmt.Printf("%s", romNotFoundErr) // probably in the future show a window?
			os.Exit(-1)
		} else if errors.As(err, &dirWriteErr) {
			fmt.Printf("%s", dirWriteErr)
			os.Exit(-1)
		} else {
			panic(err)
		}
	}

	return configValues
}
