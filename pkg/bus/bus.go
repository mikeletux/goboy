package bus

import (
	"github.com/mikeletux/goboy/pkg/cart"
	"github.com/mikeletux/goboy/pkg/log"
)

type DataBusInterface interface {
	BusRead(address uint16) byte
	BusWrite(address uint16, value byte)
	BusWrite16(address uint16, value uint16)
}

// Bus represents the whole Gameboy bus
type Bus struct {
	// Cartridge represents the Gameboy game cartridge. It has functions to read and write its memory.
	Cartridge cart.CartridgeInterface
	logger    log.Logger
}

// NewBus initializes a bus given a type that implements the cart.CartridgeInterface interface
func NewBus(cartridge cart.CartridgeInterface, logger log.Logger) *Bus {
	return &Bus{
		Cartridge: cartridge,
		logger:    logger,
	}
}

// BusRead given an address it returns the value from that bus memory area
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

// BusWrite16 writes two bytes into bus given a bus memory area
func (b *Bus) BusWrite16(address uint16, value uint16) {}
