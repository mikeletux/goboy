package ppu

import (
	"github.com/mikeletux/goboy/pkg/bus"
	"github.com/mikeletux/goboy/pkg/log"
)

type PPU struct {
	bus    bus.DataBusInterface
	lcd    *Lcd
	logger log.Logger
}

func InitPPU(bus bus.DataBusInterface, lcd *Lcd, logger log.Logger) *PPU {
	return &PPU{
		bus:    bus,
		lcd:    lcd,
		logger: logger,
	}
}
