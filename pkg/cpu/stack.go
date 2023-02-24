package cpu

func (c *CPU) stackPush(value byte) {
	c.registers.SP--
	c.bus.BusWrite(c.registers.SP, value)
}
func (c *CPU) stackPush16(value uint16) {
	c.stackPush(byte(value >> 8 & 0xFF))
	c.stackPush(byte(value & 0xFF))
}

func (c *CPU) stackPop() byte {
	value := c.bus.BusRead(c.registers.SP)
	c.registers.SP++
	return value
}
func (c *CPU) stackPop16() uint16 {
	low := uint16(c.stackPop())
	high := uint16(c.stackPop())
	return high<<8 | low
}
