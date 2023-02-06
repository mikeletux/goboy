package cpu

func nopExecFunc(c *CPU) {
	c.emulateCpuCycles(4)
	return
}

func xorExecFunc(c *CPU) {
	c.registers.A ^= c.registers.A // This is for sure wrong. Double check
	if c.registers.A == 0x0 {
		c.registers.SetFZ(true)
	}
	c.emulateCpuCycles(4)
}

func jpExecFunc(c *CPU) {
	if checkCondition(c) {
		c.registers.PC = c.FetchedData
		c.emulateCpuCycles(1)
	}
}

func checkCondition(c *CPU) bool {
	fz := c.registers.GetFZ()
	fc := c.registers.GetFC()

	switch c.CurrentInstruction.Condition {
	case ctNone:
		return true
	case ctZ:
		return fz
	case ctNZ:
		return !fz
	case ctC:
		return fc
	case ctNC:
		return !fc
	}
	return true // This never should be reached
}

func diExecFunc(c *CPU) {
	c.EnableMasterInterruptions = false
}
