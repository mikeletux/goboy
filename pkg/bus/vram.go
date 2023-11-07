package bus

import "github.com/mikeletux/goboy/pkg/log"

const vramSize = VramEnd - VramStart + 1

type VRam struct {
	logger   log.Logger
	VideoRam [vramSize]byte // length 0x2000
}

func NewVRam(logger log.Logger) *VRam {
	return &VRam{
		logger: logger,
	}
}

func (v *VRam) readVRam(address uint16) byte {
	address -= VramStart
	if address > vramSize {
		v.logger.Fatalf("Invalid Video RAM read address 0x%X", address)
	}

	return v.VideoRam[address]
}

func (v *VRam) writeVRam(address uint16, value byte) {
	address -= VramStart
	if address > vramSize {
		v.logger.Fatalf("Invalid Video RAM write address 0x%X", address)
	}

	v.VideoRam[address] = value
}
