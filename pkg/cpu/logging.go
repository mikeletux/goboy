package cpu

func (c *CPU) logRegisterValues(instructionPC uint16) {
	c.logger.Debugf("[PC:%X]:%X(%s) - [A:%X] [BC:%X] [DE:%X] [HL:%X] [Z:%d N:%d H:%d C:%d] [SP:%X]",
		instructionPC, c.CurrentOperationCode, c.CurrentInstruction.Mnemonic, c.registers.A, c.registers.GetBC(),
		c.registers.GetDE(), c.registers.GetHL(),
		fromBoolToInt(c.registers.GetFZ()),fromBoolToInt(c.registers.GetFN()),
		fromBoolToInt(c.registers.GetFH()),fromBoolToInt(c.registers.GetFC()),
		c.registers.SP)
}

// fromBoolToInt is used to print more readable values for flags in logging
func fromBoolToInt(value bool) int {
	if value {
		return 1
	}
	return 0
}
