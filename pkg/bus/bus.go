package bus

import (
	"github.com/mikeletux/goboy/pkg/cart"
	"github.com/mikeletux/goboy/pkg/log"
)

type DataBusInterface interface {
	BusRead(address uint16) byte
	BusWrite(address uint16, value byte)
	BusRead16(address uint16) uint16
	BusWrite16(address uint16, value uint16)
}

// Bus represents the whole Gameboy bus
type Bus struct {
	// Cartridge represents the Gameboy game cartridge. It has functions to read and write its memory.
	Cartridge cart.CartridgeInterface
	logger    log.Logger
}

// NewBus initializes a bus given a type that implements the cart.CartridgeInterface interface.
func NewBus(cartridge cart.CartridgeInterface, logger log.Logger) *Bus {
	return &Bus{
		Cartridge: cartridge,
		logger:    logger,
	}
}

// BusRead given an address it returns the value from that bus memory area.
func (b *Bus) BusRead(address uint16) byte {
	if address <= RomBank01NNEnd {
		return b.Cartridge.CartRead(address)
	}

	// NO IMPLEMENTED
	b.logger.Fatalf("bus address 0x%X yet not implemented to read", address)

	return 0
}

// BusWrite writes a byte into bus given a bus memory area.
func (b *Bus) BusWrite(address uint16, value byte) {
	if address <= RomBank01NNEnd {
		b.Cartridge.CartWrite(address, value)
		return
	}

	// NO IMPLEMENTED
	b.logger.Fatalf("bus address 0x%X yet not implemented to write", address)
}

// BusRead16 given an address it returns the 16 bit value from that area.
func (b *Bus) BusRead16(address uint16) uint16 {
	low := b.BusRead(address)
	high := b.BusRead(address + 1)
	return uint16(high<<8 | low)
}

// BusWrite16 writes 16 bit value into bus given a bus memory area.
func (b *Bus) BusWrite16(address uint16, value uint16) {
	b.BusWrite(address, byte(value&0xFF))        // Low
	b.BusWrite(address+1, byte((value>>8)&0xFF)) // High
}
