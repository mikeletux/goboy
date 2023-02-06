package bus

import (
	"fmt"
	"github.com/mikeletux/goboy/pkg/cart"
	"github.com/mikeletux/goboy/pkg/log"
	"github.com/mikeletux/goboy/pkg/misc"
	"strconv"
)

type DataBusInterface interface {
	BusRead(address uint16) byte
	BusWrite(address uint16, value byte)
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
	misc.NoImplemented(fmt.Sprintf("this bus address is not implemented yet to read. Addr -> %s",
		strconv.FormatInt(int64(address), 16)), -5)

	return 0
}

// BusWrite writes a value into bus given a bus memory area
func (b *Bus) BusWrite(address uint16, value byte) {
	if address <= RomBank01NNEnd {
		b.Cartridge.CartWrite(address, value)
		return
	}

	// NO IMPLEMENTED
	misc.NoImplemented(fmt.Sprintf("this bus address is not implemented yet to write. Addr -> %s",
		strconv.FormatInt(int64(address), 16)), -5)
}
