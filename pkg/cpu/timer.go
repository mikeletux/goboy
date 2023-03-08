package cpu

const (
	divRegisterAddr  uint16 = 0xFF04
	timaRegisterAddr uint16 = 0xFF05
	tmaRegisterAddr  uint16 = 0xFF06
	tacRegisterAddr  uint16 = 0xFF07
)

func (c *CPU) timerTick() {
	previousDiv := c.bus.GetTimerDiv()
	currentDiv := c.bus.IncrementTimerDiv()

	timerUpdate := false

	tac := c.getTac()
	switch tac & 0b11 {
	case 0b00:
		timerUpdate = (previousDiv&(1<<9) == 1<<9) && !(currentDiv&(1<<9) == 1<<9)
		break
	case 0b01:
		timerUpdate = (previousDiv&(1<<3) == 1<<3) && !(currentDiv&(1<<3) == 1<<3)
		break
	case 0b10:
		timerUpdate = (previousDiv&(1<<5) == 1<<5) && !(currentDiv&(1<<5) == 1<<5)
		break
	case 0b11:
		timerUpdate = (previousDiv&(1<<7) == 1<<7) && !(currentDiv&(1<<7) == 1<<7)
		break
	}

	if timerUpdate && tac&(1<<2) == 1<<2 {
		tima := c.incrementTima()
		if tima == 0xFF {
			c.setTima(c.getTma())

			c.requestInterrupt(timerInterruptFlag)
		}
	}
}

func (c *CPU) getTima() byte      { return c.bus.BusRead(timaRegisterAddr) }
func (c *CPU) setTima(value byte) { c.bus.BusWrite(timaRegisterAddr, value) }
func (c *CPU) incrementTima() byte {
	tima := c.getTima()
	c.bus.BusWrite(timaRegisterAddr, tima+1)
	return tima + 1
}

func (c *CPU) getTma() byte { return c.bus.BusRead(tmaRegisterAddr) }
func (c *CPU) getTac() byte { return c.bus.BusRead(tacRegisterAddr) }
