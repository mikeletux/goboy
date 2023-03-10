package cart

import (
	"encoding/binary"
	"fmt"
	"github.com/mikeletux/goboy/pkg/log"
	"os"
)

type CartridgeHeader struct {
	// EntryPoint instruction for the boot ROM.
	EntryPoint [4]byte
	// NintendoLogo is the bitmap image that is displayed when the Game Boy is powered on.
	NintendoLogo [48]byte
	// Title contains ASCII representation of the upper case name of the game.
	Title [16]byte
	// ManufacturerCode in older cartridges was part of the title. The purpose of this is unknown.
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

func (c *CartridgeHeader) GetReadableTitle() string {
	return string(c.Title[:])
}

func (c *CartridgeHeader) SGBFlag() bool {
	if c.SgbFlag == 0x3 {
		return true
	}
	return false
}

func (c *CartridgeHeader) GetReadableCartridgeType() string {
	return CartridgeType[c.CartridgeType]
}

// GetReadableRomSize returns how much ROM is present in the cartridge in KiB
func (c *CartridgeHeader) GetReadableRomSize() int {
	return 32 * (1 << RomSize[c.RomSize])
}

func (c *CartridgeHeader) GetReadableRamSize() string {
	return RamSize[c.RamSize]
}

func (c *CartridgeHeader) GetReadableLicenseeCode() string {
	if c.OldLicenseeCode == 0x33 {
		return NewLicenseeCodePublishers[binary.BigEndian.Uint16(c.NewLicenseeCode[:])]
	}
	return OldLicenseeCode[c.OldLicenseeCode]
}

type CartridgeInterface interface {
	CartRead(address uint16) byte
	CartWrite(address uint16, value byte)
}

// Cartridge implements all the logic regarding GB cartridges
type Cartridge struct {
	CartridgeHeader *CartridgeHeader
	rawData         []byte
	logger          log.Logger
}

// NewCartridge returns a pointer to Cartridge given a Rom path
func NewCartridge(romPath string, logger log.Logger) (*Cartridge, error) {
	romData, err := os.ReadFile(romPath)
	if err != nil {
		return nil, fmt.Errorf("error while loading ROM cartridges - %v", err)
	}

	// Do checksum to ensure cartridge integrity
	var checksum uint8
	for address := TitleAddrStart; address <= MaskRomVersionNumberAddr; address++ {
		checksum = checksum - romData[address] - 1
	}

	if checksum != romData[HeaderChecksumAddr] {
		return nil, fmt.Errorf("calculated header checksum doesn't correspond with cartridge checksum")
	}

	return &Cartridge{
		CartridgeHeader: parseCartridgeHeader(romData),
		rawData:         romData,
		logger:          logger,
	}, nil
}

// CartRead returns a given byte from the cartridge given a memory address
func (c *Cartridge) CartRead(address uint16) byte {
	// For now just ROM only type supported
	return c.rawData[address]
}

// CartWrite write a value in the address specified
func (c *Cartridge) CartWrite(address uint16, value byte) {
	// For now just ROM only type supported
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

func (c *Cartridge) LogCartridgeHeaderInfo() {
	c.logger.Debugf("Cartridge information:\nTitle:%s\nLicensee:%s\nCartridge type:%s\nRom size:%d KiB\nRam size:%s KiB\nSGB flag:%t\n",
		c.CartridgeHeader.GetReadableTitle(),
		c.CartridgeHeader.GetReadableLicenseeCode(),
		c.CartridgeHeader.GetReadableCartridgeType(),
		c.CartridgeHeader.GetReadableRomSize(),
		c.CartridgeHeader.GetReadableRamSize(),
		c.CartridgeHeader.SGBFlag(),
	)
}
