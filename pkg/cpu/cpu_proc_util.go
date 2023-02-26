package cpu

func checkCondition(c *CPU) bool { // This should be a method
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

func (c *CPU) gotoAddr(address uint16, pushPC bool) {
	if checkCondition(c) {
		if pushPC {
			c.stackPush16(c.registers.PC)
			c.emulateCpuCycles(2)
		}
		c.registers.PC = address
		c.emulateCpuCycles(1)
	}
}

func is16BitRegister(register int) bool {
	return register >= rtAF
}
