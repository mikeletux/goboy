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

func ldExecFunc(c *CPU) {
	if c.DestinationIsMemory {
		// We need to write in memory
		if c.CurrentInstruction.RegisterType2 >= rtAF { // This means we need to write twice in memory.
			c.bus.BusWrite16(c.MemoryDestination, c.FetchedData)
		} else {
			c.bus.BusWrite(c.MemoryDestination, byte(c.FetchedData))
		}
		return
	}

	if c.CurrentInstruction.AddressingMode == amHLnSPR {
		c.registers.SetFZ(false)
		c.registers.SetFN(false)
		reg2Value, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType2)
		if err != nil {
			c.logger.Fatalf("error when executing LD HL SP(r) operation: %s", err)
		}

		c.registers.SetFH((reg2Value&0xF)+(c.FetchedData&0xF) >= 0x10)    // If lower 4 bits of result overflow, set H.
		c.registers.SetFC((reg2Value&0xFF)+(c.FetchedData&0xFF) >= 0x100) // If upper 4 bits of result overflow, set C.

		c.registers.SetDataToRegisters(c.CurrentInstruction.RegisterType1, reg2Value+c.FetchedData)
	}

	c.registers.SetDataToRegisters(c.CurrentInstruction.RegisterType1, c.FetchedData) // Normal case.
}

func ldhExecFunc(c *CPU) {
	if c.CurrentInstruction.RegisterType1 == rtA {
		c.registers.A = c.bus.BusRead(0xFF00 | c.FetchedData)
	} else {
		c.bus.BusWrite(0xFF00|c.FetchedData, c.registers.A)
	}

	c.emulateCpuCycles(1)
}
