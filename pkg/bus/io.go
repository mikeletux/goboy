package bus

import "github.com/mikeletux/goboy/pkg/log"

const (
	divRegisterAddr  uint16 = 0xFF04
	timaRegisterAddr uint16 = 0xFF05
	tmaRegisterAddr  uint16 = 0xFF06
	tacRegisterAddr  uint16 = 0xFF07

	initialDivRegisterValue uint16 = 0xABCC
)

type timer struct {
	divReg  uint16 // FF04
	timaReg byte   // FF05
	tmaReg  byte   // FF06
	tacReg  byte   // FF07
}

type io struct {
	logger log.Logger
	// ioRegisters [IORegistersEnd - IORegistersStart + 1]byte

	timer *timer
}

func NewIO(logger log.Logger) *io {
	io := &io{
		logger: logger,
		timer: &timer{
			divReg: initialDivRegisterValue,
		},
	}

	return io
}

func (i *io) IORead(address uint16) byte {
	if address == 0xFF44 {
		return 0x90 // Hardcoded value for Gameboy doctor
	}

	switch address { // This switch is for special cases (Like 16bit Timer DIV register)
	case divRegisterAddr:
		return byte(i.timer.divReg >> 8)
	case timaRegisterAddr:
		return i.timer.timaReg
	case tmaRegisterAddr:
		return i.timer.tmaReg
	case tacRegisterAddr:
		return i.timer.tacReg
	default:
		return 0x0
	}

	// return i.ioRegisters[address-IORegistersStart]
}

func (i *io) IOWrite(address uint16, data byte) {
	switch address { // This switch is for special cases (Like 16bit Timer DIV register)
	case divRegisterAddr:
		i.timer.divReg = 0
	case timaRegisterAddr:
		i.timer.timaReg = data
	case tmaRegisterAddr:
		i.timer.tmaReg = data
	case tacRegisterAddr:
		i.timer.tacReg = data
	}

	// i.ioRegisters[address-IORegistersStart] = data
}
