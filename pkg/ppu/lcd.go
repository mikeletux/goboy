package ppu

import (
	"github.com/mikeletux/goboy/pkg/bus"
	"github.com/mikeletux/goboy/pkg/common"
	"github.com/mikeletux/goboy/pkg/log"
	"github.com/mikeletux/goboy/pkg/ui"
)

const (
	modeHBlank = 0
	modeVBlank = 1
	modeOam    = 2
	modeXfer   = 3
)

type Lcd struct {
	bus    bus.DataBusInterface
	ui     *ui.UI
	logger log.Logger

	bgColors  [4]uint32
	sp1Colors [4]uint32
	sp2Colors [4]uint32
}

func InitLcd(bus bus.DataBusInterface, ui *ui.UI, logger log.Logger) *Lcd {
	return &Lcd{
		bus: bus,
		ui:  ui, // MAKE UI AT SOME POINT INTERFACE
	}
}

/*
   Bit 7: LCD & PPU enable: 0 = Off; 1 = On
   Bit 6: Window tile map area: 0 = 9800–9BFF; 1 = 9C00–9FFF
   Bit 5: Window enable: 0 = Off; 1 = On
   Bit 4: BG & Window tile data area: 0 = 8800–97FF; 1 = 8000–8FFF
   Bit 3: BG tile map area: 0 = 9800–9BFF; 1 = 9C00–9FFF
   Bit 2: OBJ size: 0 = 8×8; 1 = 8×16
   Bit 1: OBJ enable: 0 = Off; 1 = On
   Bit 0: BG & Window enable / priority [Different meaning in CGB Mode]: 0 = Off; 1 = On
*/

func (l *Lcd) lcdBgwEnable() bool {
	return common.GetBitRegister(l.bus.BusRead(bus.LcdControlRegisterAddr), 0)
}

func (l *Lcd) lcdObjEnable() bool {
	return common.GetBitRegister(l.bus.BusRead(bus.LcdControlRegisterAddr), 1)
}

func (l *Lcd) lcdObjSize() int {
	if common.GetBitRegister(l.bus.BusRead(bus.LcdControlRegisterAddr), 2) {
		return 16
	}
	return 8
}

func (l *Lcd) lcdBgMapArea() uint16 {
	if common.GetBitRegister(l.bus.BusRead(bus.LcdControlRegisterAddr), 3) {
		return 0x9C00
	}

	return 0x9800
}

func (l *Lcd) lcdBgWindowDataArea() uint16 {
	if common.GetBitRegister(l.bus.BusRead(bus.LcdControlRegisterAddr), 4) {
		return 0x8000
	}

	return 0x8800
}

func (l *Lcd) lcdWindowEnable() bool {
	return common.GetBitRegister(l.bus.BusRead(bus.LcdControlRegisterAddr), 5)
}

func (l *Lcd) lcdWindowMapArea() uint16 {
	if common.GetBitRegister(l.bus.BusRead(bus.LcdControlRegisterAddr), 6) {
		return 0x9C00
	}

	return 0x9800
}

func (l *Lcd) lcdEnable() bool {
	return common.GetBitRegister(l.bus.BusRead(bus.LcdControlRegisterAddr), 7)
}

/*
   Bit 7: No use
   Bit 6: LYC int select (Read/Write): If set, selects the LYC == LY condition for the STAT interrupt.
   Bit 5: Mode 2 int select (Read/Write): If set, selects the Mode 2 condition for the STAT interrupt.
   Bit 4: Mode 1 int select (Read/Write): If set, selects the Mode 1 condition for the STAT interrupt.
   Bit 3: Mode 0 int select (Read/Write): If set, selects the Mode 0 condition for the STAT interrupt.
   Bit 2: LYC == LY (Read-only): Set when LY contains the same value as LYC; it is constantly updated.
   Bit 0-1: PPU mode (Read-only): Indicates the PPU’s current status.
*/
