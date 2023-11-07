package ppu

import (
	"github.com/mikeletux/goboy/pkg/bus"
	"github.com/mikeletux/goboy/pkg/log"
)

type PPU struct {
	bus    bus.DataBusInterface
	logger log.Logger
}

func Init(bus bus.DataBusInterface, logger log.Logger) *PPU {
	return &PPU{
		bus:    bus,
		logger: logger,
	}
}
