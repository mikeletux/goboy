package ppu

import (
	"github.com/mikeletux/goboy/pkg/bus"
	"github.com/mikeletux/goboy/pkg/log"
	"github.com/mikeletux/goboy/pkg/ui"
)

type Lcd struct {
	bus    bus.DataBusInterface
	ui     *ui.UI
	logger log.Logger
}

func InitLcd(bus bus.DataBusInterface, ui *ui.UI, logger log.Logger) *Lcd {
	return &Lcd{
		bus: bus,
		ui:  ui, // MAKE UI AT SOME POINT INTERFACE
	}
}
