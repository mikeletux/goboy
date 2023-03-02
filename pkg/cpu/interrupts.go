package cpu

func (c *CPU) handleInterrupt(address uint16) {
	c.stackPush16(c.registers.PC)
	c.registers.PC = address
}


