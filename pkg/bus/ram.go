package bus

import "github.com/mikeletux/goboy/pkg/log"

const (
	workingRamSize = WorkRam1End - WorkRam0Start + 1
	highRamSize    = HighRamEnd - HighRamStart + 1
)

type Ram struct {
	logger     log.Logger
	WorkingRam [workingRamSize]byte // length 0x2000
	HighRam    [highRamSize]byte    // length 0x7F
}

func NewRam(logger log.Logger) *Ram {
	return &Ram{
		logger: logger,
	}
}

func (r *Ram) readWorkingRam(address uint16) byte {
	address -= WorkRam0Start
	if address > workingRamSize { // Bigger than working RAM size
		r.logger.Fatalf("Invalid working RAM address 0x%X", address)
	}

	return r.WorkingRam[address]
}

func (r *Ram) writeWorkingRam(address uint16, value byte) {
	address -= WorkRam0Start
	if address > workingRamSize { // Bigger than working RAM size
		r.logger.Fatalf("Invalid working RAM address 0x%X", address)
	}

	r.WorkingRam[address] = value
}

func (r *Ram) readHighRam(address uint16) byte {
	address -= HighRamStart
	if address > highRamSize { // Bigger than working RAM size
		r.logger.Fatalf("Invalid high RAM address 0x%X", address)
	}

	return r.HighRam[address]
}

func (r *Ram) writeHighRam(address uint16, value byte) {
	address -= HighRamStart
	if address > highRamSize { // Bigger than working RAM size
		r.logger.Fatalf("Invalid high RAM address 0x%X", address)
	}

	r.HighRam[address] = value
}
