package cpu

import "fmt"

func (c *CPU) logRegisterValues(instructionPC uint16) {
	c.logger.Debugf("[PC:%X]:%X(%s) - [A:%X] [BC:%X] [DE:%X] [HL:%X] [Z:%d N:%d H:%d C:%d] [SP:%X]",
		instructionPC, c.CurrentOperationCode, c.CurrentInstruction.Mnemonic, c.registers.A, c.registers.GetBC(),
		c.registers.GetDE(), c.registers.GetHL(),
		fromBoolToInt(c.registers.GetFZ()), fromBoolToInt(c.registers.GetFN()),
		fromBoolToInt(c.registers.GetFH()), fromBoolToInt(c.registers.GetFC()),
		c.registers.SP)
}

// fromBoolToInt is used to print more readable values for flags in logging
func fromBoolToInt(value bool) int {
	if value {
		return 1
	}
	return 0
}

func (c *CPU) logRegistersGameboyDoctor(instructionPC uint16) {
	pc0 := c.bus.BusRead(instructionPC)
	pc1 := c.bus.BusRead(instructionPC + 1)
	pc2 := c.bus.BusRead(instructionPC + 2)
	pc3 := c.bus.BusRead(instructionPC + 3)

	c.logger.Debugf("A:%s F:%s B:%s C:%s D:%s E:%s H:%s L:%s SP:%s PC:%s PCMEM:%s,%s,%s,%s",
		printAllNibbles(false, uint16(c.registers.A)),
		printAllNibbles(false, uint16(c.registers.F)),
		printAllNibbles(false, uint16(c.registers.B)),
		printAllNibbles(false, uint16(c.registers.C)),
		printAllNibbles(false, uint16(c.registers.D)),
		printAllNibbles(false, uint16(c.registers.E)),
		printAllNibbles(false, uint16(c.registers.H)),
		printAllNibbles(false, uint16(c.registers.L)),
		printAllNibbles(true, c.registers.SP),
		printAllNibbles(true, instructionPC),
		printAllNibbles(false, uint16(pc0)),
		printAllNibbles(false, uint16(pc1)),
		printAllNibbles(false, uint16(pc2)),
		printAllNibbles(false, uint16(pc3)),
	)
}

func printAllNibbles(is16Bit bool, value uint16) string {
	if is16Bit {
		return fmt.Sprintf("%X%X%X%X", value>>12&0xF, value>>8&0xF, value>>4&0xF, value&0xF)
	}

	return fmt.Sprintf("%X%X", value>>4&0xF, value&0xF)
}
