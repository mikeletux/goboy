package cpu

var dbgMsg [1024]byte
var msgSize = 0

func (c *CPU) dbgUpdate() {
	if c.bus.BusRead(serialTransferControlIOAddr) == 0x81 {
		v := c.bus.BusRead(serialTransferDataIOAddr)
		dbgMsg[msgSize] = v
		msgSize++

		c.bus.BusWrite(serialTransferControlIOAddr, 0)
	}
}

func (c *CPU) dbgPrint() {
	if dbgMsg[0] != 0 {
		c.logger.Debugf("DBG: %s", string(dbgMsg[:]))
	}
}
