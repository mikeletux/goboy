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
	CgbFlag byte
	// NewLicenseeCode contains a two-character ASCII “licensee code” indicating the game’s publisher.
	NewLicenseeCode [2]byte
	// SgbFlag specifies whether the game supports SGB functions.
	SgbFlag byte
	// CartridgeType indicates what kind of hardware is present on the cartridge.
	CartridgeType byte
	// RomSize indicates how much ROM is present on the cartridge.
	RomSize byte
	// RamSize indicates how much RAM is present on the cartridge, if any.
	RamSize byte
	// DestinationCode specifies whether this version of the game is intended to be sold in Japan or elsewhere.
	DestinationCode byte
	// OldLicenseeCode is used in older (pre-SGB) cartridges to specify the game’s publisher.
	OldLicenseeCode byte
	// MaskRomVersionNumber specifies the version number of the game. It is usually 0x0.
	MaskRomVersionNumber byte
	// HeaderCheckSum contains an 8-bit checksum computed from the cartridge header bytes 0x134–0x14C.
	HeaderCheckSum byte
	// GlobalChecksum contain a 16-bit (big-endian) checksum simply computed as the sum of all the bytes of
	// the cartridge ROM (except these two checksum bytes).
	GlobalChecksum [2]byte
}

// Cartridge implements all the logic regarding GB cartridges
type Cartridge struct {
	CartridgeHeader *CartridgeHeader
	rawData         []byte
}

// NewCartridge returns a pointer to Cartridge given a Rom path
func NewCartridge(romPath string) (*Cartridge, error) {
	romData, err := os.ReadFile(romPath)
	if err != nil {
		return nil, fmt.Errorf("error while loading ROM cartridges - %v", err)
	}

	return &Cartridge{
		CartridgeHeader: parseCartridgeHeader(romData),
		rawData:         romData,
	}, nil
}

func parseCartridgeHeader(cartridgeRawData []byte) *CartridgeHeader {
	cartridgeHeader := &CartridgeHeader{}
	copy(cartridgeHeader.EntryPoint[:], cartridgeRawData[EntryPointAddrStart:EntryPointAddrEnd+1])
	copy(cartridgeHeader.NintendoLogo[:], cartridgeRawData[NintendoLogoAddrStart:NintendoLogoAddrEnd+1])
	copy(cartridgeHeader.Title[:], cartridgeRawData[TitleAddrStart:TitleAddrEnd+1])
	copy(cartridgeHeader.ManufacturerCode[:], cartridgeRawData[ManufacturerAddrStart:ManufacturerAddrEnd+1])
	cartridgeHeader.CgbFlag = cartridgeRawData[CgbFlagAddr]
	copy(cartridgeHeader.NewLicenseeCode[:], cartridgeRawData[NewLicenseeCodeAddrStart:NewLicenseeCodeAddrEnd+1])
	cartridgeHeader.SgbFlag = cartridgeRawData[SgbFlagAddr]
	cartridgeHeader.CartridgeType = cartridgeRawData[CartridgeTypeAddr]
	cartridgeHeader.RomSize = cartridgeRawData[RomSizeAddr]
	cartridgeHeader.RamSize = cartridgeRawData[RamSizeAddr]
	cartridgeHeader.DestinationCode = cartridgeRawData[DestinationCodeAddr]
	cartridgeHeader.OldLicenseeCode = cartridgeRawData[OldLicenseeCodeAddr]
	cartridgeHeader.MaskRomVersionNumber = cartridgeRawData[MaskRomVersionNumberAddr]
	cartridgeHeader.HeaderCheckSum = cartridgeRawData[HeaderChecksumAddr]
	copy(cartridgeHeader.GlobalChecksum[:], cartridgeRawData[GlobalChecksumAddrStart:GlobalChecksumAddrEnd+1])
	return cartridgeHeader
}
