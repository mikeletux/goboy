package cpu

var dbgMsg [1024]byte
var msgSize = 0

func (c *CPU) dbgUpdate(){
	if c.bus.BusRead(0xFF02) == 0x81 {
		v := c.bus.BusRead(0xFF01)
		dbgMsg[msgSize] = v
		msgSize++

		c.bus.BusWrite(0xFF02, 0)
	}
}

func (c *CPU) dbgPrint(){
	if dbgMsg[0] != 0 {
		c.logger.Debugf("DBG: %s", string(dbgMsg[:]))
	}
}