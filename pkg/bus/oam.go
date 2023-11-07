package bus

import "github.com/mikeletux/goboy/pkg/log"

const oamSize = OamEnd - OamStart + 1

type Oam struct {
	logger                log.Logger
	objectAttributeMemory [oamSize]byte
}

func NewOam(logger log.Logger) *Oam {
	return &Oam{
		logger: logger,
	}
}

func (o *Oam) readOam(address uint16) byte {
	address -= OamStart
	if address >= oamSize {
		o.logger.Fatalf("Invalid OAM read address 0x%X", address)
	}

	return o.objectAttributeMemory[address]
}

func (o *Oam) writeOam(address uint16, value byte) {
	address -= OamStart
	if address >= oamSize {
		o.logger.Fatalf("Invalid OAM write address 0x%X", address)
	}

	o.objectAttributeMemory[address] = value
}
