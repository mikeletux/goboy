package bus

import "github.com/mikeletux/goboy/pkg/log"

type io struct {
	logger      log.Logger
	ioRegisters [IORegistersEnd - IORegistersStart + 1]byte
}

func NewIO(logger log.Logger) *io {
	return &io{
		logger: logger,
	}
}

func (i *io) IORead(address uint16) byte {
	if address == 0xFF44 {
		return 0x90 // Hardcoded value for Gameboy doctor
	}

	return i.ioRegisters[address-IORegistersStart]
}

func (i *io) IOWrite(address uint16, data byte) {
	i.ioRegisters[address-IORegistersStart] = data
}
