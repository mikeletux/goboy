package cart

import (
	"fmt"
	"os"
)

type CartridgeHeader struct {
	// EntryPoint instruction for the boot ROM.
	EntryPoint [4]byte
	// NintendoLogo is the bitmap image that is displayed when the Game Boy is powered on.
	NintendoLogo [30]byte
	// Title contains ASCII representation of the upper case name of the game.
	Title [10]byte
	// ManufacturerCode in older cartridges was part of the title. The purpose of this si unknown.
	ManufacturerCode [4]byte
	// CgbFlag in older cartridges was part of the title. It decides whether enable color mode.
	CgbFlag              byte
	NewLicenseeCode      [2]byte
	SgbFlag              byte
	CartridgeType        byte
	RomSize              byte
	RamSize              byte
	DestinationCode      byte
	OldLicenseeCode      byte
	MaskRomVersionNumber byte
	HeaderCheckSum       byte
	GlobalChecksum       [2]byte
}

// Cartridge implements all the logic regarding GB cartridges
type Cartridge struct {
	rawData []byte
}

// NewCartridge returns a pointer to Cartridge given a Rom path
func NewCartridge(romPath string) (*Cartridge, error) {
	romData, err := os.ReadFile(romPath)
	if err != nil {
		return nil, fmt.Errorf("error while loading ROM cartridges - %v", err)
	}

	// Perform checks that the ROM is ok

	return &Cartridge{
		rawData: romData,
	}, nil
}
