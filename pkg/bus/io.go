package bus

import (
	"github.com/mikeletux/goboy/pkg/log"
)

// Register addresses
const (
	serialTransferDataAddr    uint16 = 0xFF01
	serialTransferControlAddr uint16 = 0xFF02

	divRegisterAddr  uint16 = 0xFF04
	timaRegisterAddr uint16 = 0xFF05
	tmaRegisterAddr  uint16 = 0xFF06
	tacRegisterAddr  uint16 = 0xFF07

	interruptFlagRegisterAddr uint16 = 0xFF0F
	oamDmaRegisterAddr        uint16 = 0xFF46
)

const (
	initialDivRegisterValue uint16 = 0xABCC
)

type timer struct {
	divReg  uint16 // FF04
	timaReg byte   // FF05
	tmaReg  byte   // FF06
	tacReg  byte   // FF07
}

type serial struct {
	serialTransferData    byte // FF01
	serialTransferControl byte // FF02
}

type io struct {
	serial *serial
	timer  *timer
	ifReg  byte // Interrupt Flag FF0F
	ly     byte // HACK - REMOVE IT LATER
	dma    *Dma
	logger log.Logger
}

func NewIO(logger log.Logger, dma *Dma) *io {
	io := &io{
		logger: logger,
		serial: &serial{},
		timer: &timer{
			divReg: initialDivRegisterValue,
		},
		dma: dma,
	}

	return io
}

func (i *io) IORead(address uint16) byte {
	if address == 0xFF44 {
		i.ly++
		return i.ly // HACK
	}

	switch address { // This switch is for special cases (Like 16bit Timer DIV register)
	case serialTransferDataAddr:
		return i.serial.serialTransferData
	case serialTransferControlAddr:
		return i.serial.serialTransferControl
	case divRegisterAddr:
		return byte(i.timer.divReg >> 8)
	case timaRegisterAddr:
		return i.timer.timaReg
	case tmaRegisterAddr:
		return i.timer.tmaReg
	case tacRegisterAddr:
		return i.timer.tacReg
	case interruptFlagRegisterAddr:
		return i.ifReg
	default:
		return 0x0
	}
}

func (i *io) IOWrite(address uint16, data byte) {
	switch address { // This switch is for special cases (Like 16bit Timer DIV register)
	case serialTransferDataAddr:
		i.serial.serialTransferData = data
	case serialTransferControlAddr:
		i.serial.serialTransferControl = data
	case divRegisterAddr:
		i.timer.divReg = 0
	case timaRegisterAddr:
		i.timer.timaReg = data
	case tmaRegisterAddr:
		i.timer.tmaReg = data
	case tacRegisterAddr:
		i.timer.tacReg = data
	case interruptFlagRegisterAddr:
		i.ifReg = data
	case oamDmaRegisterAddr:
		i.dma.start(data)
		i.logger.Debugf("DMA STARTED\n")
	}
}
