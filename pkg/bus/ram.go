package bus

import "github.com/mikeletux/goboy/pkg/log"

type Ram struct {
	logger     *log.Logger
	WorkingRam [WorkRam1End - WorkRam0Start + 1]byte // length 0x2000
	HighRam    [HighRamEnd - HighRamStart + 1]byte   // length 0x7F
}

func NewRam(logger *log.Logger) *Ram {
	return &Ram{
		logger: logger,
	}
}

func (r *Ram) readRam(address uint16) byte {
	return 0
}

func (r *Ram) writeRam(address uint16, value byte) {

}

func (r *Ram) readHighRam(address uint16) byte {
	return 0
}

func (r *Ram) writeHighRam(address uint16, value byte) {

}
