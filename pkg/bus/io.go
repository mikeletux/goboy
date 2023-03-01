package bus

import "github.com/mikeletux/goboy/pkg/log"

type io struct {
	logger log.Logger
	serialData [2]byte
}

func NewIO(logger log.Logger) *io {
	return &io{
		logger: logger,
	}
}

func (i *io) IORead(address uint16) byte {
	if address == 0xFF01 {
		return i.serialData[0]
	}

	if address == 0xFF02 {
		return i.serialData[1]
	}

	i.logger.Debug("unsupported io read")
	return 0
}

func (i *io) IOWrite(address uint16, data byte) {
	if address == 0xFF01 {
		i.serialData[0] = data
	}

	if address == 0xFF02 {
		i.serialData[1] = data
	}

	i.logger.Debug("unsupported io write")
}
