package cpu

// regsPrefixCBDecode is used to decode the reg when executing the Prefix CB instruction
var regsPrefixCBDecode = []int {
	rtB,  // 0b000
	rtC,  // 0b001
	rtD,  // 0b010
	rtE,  // 0b011
	rtH,  // 0b100
	rtL,  // 0b101
	rtHL, // 0b110
	rtA,  // 0b111
}

func (c *CPU)checkCondition() bool { // This should be a method
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
	if c.checkCondition() {
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

// decodePrefixCBRegister decodes the kind of register to be used by the Prefix CB instruction
// cbByte is the Prefix CB operation type.
func decodePrefixCBRegister(cbByte byte) int {
	encodedRegister := 0b111 & cbByte
	return regsPrefixCBDecode[encodedRegister]
}

func (c *CPU) fetchRegisterPrefixCB(register int) byte{
	switch register{
	case rtHL:
		return c.bus.BusRead(c.registers.GetHL())
	default:
		data, err := c.registers.FetchDataFromRegisters(register)
		if err != nil {
			c.logger.Fatal(err)
		}

		return byte(data & 0xFF)
	}
}

func (c *CPU) setRegisterPrefixCB(register int, data byte) {
	switch register{
	case rtHL:
		c.bus.BusWrite(c.registers.GetHL(), data)
	default:
		err := c.registers.SetDataToRegisters(register, uint16(data))
		if err != nil {
			c.logger.Fatal(err)
		}
	}
}
