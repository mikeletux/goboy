package cpu

func nopExecFunc(c *CPU) {
	return
}

func jpExecFunc(c *CPU) {
	c.gotoAddr(c.FetchedData, false)
}

func callExecFunc(c *CPU) {
	c.gotoAddr(c.FetchedData, true)
}

func rstExecFunc(c *CPU) {
	c.gotoAddr(uint16(c.CurrentInstruction.Parameter)&0xFF, true)
}

func retExecFunc(c *CPU) {
	if c.CurrentInstruction.Condition != ctNone {
		c.emulateCpuCycles(1)
	}

	if c.checkCondition() {
		low := c.stackPop()
		c.emulateCpuCycles(1)
		high := c.stackPop()
		c.emulateCpuCycles(1)

		c.registers.PC = uint16(high)<<8 | uint16(low)
		c.emulateCpuCycles(1)
	}
}

func retiExecFunc(c *CPU) {
	c.EnableMasterInterruptions = true
	retExecFunc(c)
}

func jrExecFunc(c *CPU) {
	rel := int8(c.FetchedData & 0xFF) // This byte must be signed
	addr := c.registers.PC + uint16(rel)
	c.gotoAddr(addr, false)
}

func popExecFunc(c *CPU) {
	low := uint16(c.stackPop()) // Read the least significant byte
	c.emulateCpuCycles(1)
	high := uint16(c.stackPop()) // Read the most significant byte
	c.emulateCpuCycles(1)
	c.registers.SetDataToRegisters(c.CurrentInstruction.RegisterType1, high<<8|low)

	if c.CurrentInstruction.RegisterType1 == rtAF {
		c.registers.SetDataToRegisters(c.CurrentInstruction.RegisterType1, (high<<8|low)&0xFFF0)
	}
}

func pushExecFunc(c *CPU) {
	value, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType1)
	if err != nil {
		c.logger.Fatal(err)
	}

	c.stackPush(byte(value>>8) & 0xFF) // Push the most significant byte
	c.emulateCpuCycles(1)
	c.stackPush(byte(value) & 0xFF) // Push the least significant byte
	c.emulateCpuCycles(1)

	c.emulateCpuCycles(1)
}

func diExecFunc(c *CPU) {
	c.EnableMasterInterruptions = false
}

func eiExecFunc(c *CPU) {
	c.EnablingIme = true
}

func ldExecFunc(c *CPU) {
	if c.DestinationIsMemory {
		// We need to write in memory
		if is16BitRegister(c.CurrentInstruction.RegisterType2) { // This means we need to write twice in memory.
			c.emulateCpuCycles(1)
			c.bus.BusWrite16(c.MemoryDestination, c.FetchedData)
		} else {
			c.bus.BusWrite(c.MemoryDestination, byte(c.FetchedData))
		}

		c.emulateCpuCycles(1)
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

		c.registers.SetDataToRegisters(c.CurrentInstruction.RegisterType1, reg2Value+uint16(int8(c.FetchedData)))
		return
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

func incExecFunc(c *CPU) {
	value, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType1)
	if err != nil {
		c.logger.Fatal(err)
	}

	value++ // Increment is done here
	if is16BitRegister(c.CurrentInstruction.RegisterType1) {
		c.emulateCpuCycles(1)
	}

	if c.CurrentInstruction.RegisterType1 == rtHL && c.DestinationIsMemory {
		value = c.FetchedData + 1
		value &= 0xFF
		c.bus.BusWrite(c.registers.GetHL(), byte(value))
	} else {
		err = c.registers.SetDataToRegisters(c.CurrentInstruction.RegisterType1, value)
		value &= 0xFF
		if err != nil {
			c.logger.Fatal(err)
		}
	}

	if (c.CurrentOperationCode & 0x03) == 0x03 { // 0xX3 INC instruction doesn't change flags
		return
	}

	c.registers.SetFZ(value == 0)
	c.registers.SetFN(false)
	c.registers.SetFH(value&0x0F == 0)
}

func decExecFunc(c *CPU) {
	value, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType1)
	if err != nil {
		c.logger.Fatal(err)
	}

	value-- // Decrement is done here
	if is16BitRegister(c.CurrentInstruction.RegisterType1) {
		c.emulateCpuCycles(1)
	}

	if c.CurrentInstruction.RegisterType1 == rtHL && c.DestinationIsMemory {
		value = c.FetchedData - 1
		c.bus.BusWrite(c.registers.GetHL(), byte(value))
	} else {
		err = c.registers.SetDataToRegisters(c.CurrentInstruction.RegisterType1, value)
		if err != nil {
			c.logger.Fatal(err)
		}
	}

	if (c.CurrentOperationCode & 0x0B) == 0x0B { // 0xXB DEC instruction doesn't change flags
		return
	}

	c.registers.SetFZ(value == 0)
	c.registers.SetFN(true)
	c.registers.SetFH(value&0x0F == 0x0F)
}

func addExecFunc(c *CPU) {
	var value uint32
	regValue, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType1)
	if err != nil {
		c.logger.Fatal(err)
	}

	value = uint32(regValue + c.FetchedData)

	if is16BitRegister(c.CurrentInstruction.RegisterType1) {
		c.emulateCpuCycles(1)
	}

	if c.CurrentInstruction.RegisterType1 == rtSP {
		r := int8(c.FetchedData & 0xFF) // r is a signed 8 bit integer
		value = uint32(regValue + uint16(r))
	}

	var flagZ, flagH, flagC bool
	if !is16BitRegister(c.CurrentInstruction.RegisterType1) { // for 8 bit instructions
		flagZ = value&0xFF == 0
		flagH = (regValue&0xF)+(c.FetchedData&0xF) >= 0x10
		flagC = (regValue&0xFF)+(c.FetchedData&0xFF) >= 0x100
	} else { // For 16 bit instructions
		if c.CurrentInstruction.RegisterType1 != rtSP { // If not special case
			flagZ = c.registers.GetFZ()
			flagH = (regValue&0xFFF)+(c.FetchedData&0xFFF) >= 0x1000
			flagC = uint32(regValue)+uint32(c.FetchedData) >= 0x10000
		} else { // If special case SP
			flagZ = false
			flagH = (regValue&0xF)+(c.FetchedData&0xF) >= 0x10
			flagC = (regValue&0xFF)+(c.FetchedData&0xFF) >= 0x100
		}
	}

	c.registers.SetDataToRegisters(c.CurrentInstruction.RegisterType1, uint16(value&0xFFFF))
	c.registers.SetFZ(flagZ)
	c.registers.SetFN(false)
	c.registers.SetFH(flagH)
	c.registers.SetFC(flagC)
}

func adcExecFunc(c *CPU) {
	regAValue, err := c.registers.FetchDataFromRegisters(rtA)
	if err != nil {
		c.logger.Fatal(err)
	}

	var result, carry uint16
	if c.registers.GetFC() {
		carry = 1
	}

	result = regAValue + c.FetchedData + carry
	c.registers.A = byte(result & 0xFF)

	c.registers.SetFZ(result&0xFF == 0)
	c.registers.SetFN(false)
	c.registers.SetFH(regAValue&0xF+c.FetchedData&0xF+carry >= 0x10)
	c.registers.SetFC(result >= 0x100)
}

func subExecFunc(c *CPU) {
	regAValue, err := c.registers.FetchDataFromRegisters(rtA)
	if err != nil {
		c.logger.Fatal(err)
	}

	c.registers.A = byte(regAValue&0xFF) - byte(c.FetchedData&0xFF)

	c.registers.SetFZ(regAValue&0xFF-c.FetchedData&0xFF == 0)
	c.registers.SetFN(true)
	c.registers.SetFH(int(regAValue&0xF)-int(c.FetchedData&0xF) < 0)
	c.registers.SetFC(int(regAValue&0xFF)-int(c.FetchedData&0xFF) < 0)
}

func sbcExecFunc(c *CPU) {
	regAValue, err := c.registers.FetchDataFromRegisters(rtA)
	if err != nil {
		c.logger.Fatal(err)
	}

	var carry byte
	if c.registers.GetFC() {
		carry = 1
	}

	c.registers.A = byte(regAValue&0xFF) - (byte(c.FetchedData&0xFF) + carry)

	c.registers.SetFZ(byte(regAValue&0xFF)-(byte(c.FetchedData&0xFF)+carry) == 0)
	c.registers.SetFN(true)
	c.registers.SetFH(int(regAValue&0xF)-int(c.FetchedData&0xF)-int(carry) < 0)
	c.registers.SetFC(int(regAValue&0xFF)-int(c.FetchedData&0xFF)-int(carry) < 0)
}

func andExecFunc(c *CPU) {
	c.registers.A &= byte(c.FetchedData & 0xFF)

	c.registers.SetFZ(c.registers.A == 0)
	c.registers.SetFN(false)
	c.registers.SetFH(true)
	c.registers.SetFC(false)
}

func xorExecFunc(c *CPU) {
	c.registers.A ^= byte(c.FetchedData & 0xFF)

	c.registers.SetFZ(c.registers.A == 0)
	c.registers.SetFN(false)
	c.registers.SetFH(false)
	c.registers.SetFC(false)
}
func orExecFunc(c *CPU) {
	c.registers.A |= byte(c.FetchedData & 0xFF)

	c.registers.SetFZ(c.registers.A == 0)
	c.registers.SetFN(false)
	c.registers.SetFH(false)
	c.registers.SetFC(false)
}
func cpExecFunc(c *CPU) {
	c.registers.SetFZ(c.registers.A-byte(c.FetchedData) == 0)
	c.registers.SetFN(true)
	c.registers.SetFH(int(c.registers.A&0xF)-int(c.FetchedData&0xF) < 0)
	c.registers.SetFC(int(c.registers.A&0xFF)-int(c.FetchedData&0xFF) < 0)
}

func cbExecFunc(c *CPU) {
	cbOperation := byte(c.FetchedData & 0xFF)
	registerType := decodePrefixCBRegister(cbOperation)
	bit := (cbOperation >> 3) & 0b111
	bitOperation := (cbOperation >> 6) & 0b11
	registerValue := c.fetchRegisterPrefixCB(registerType)

	c.emulateCpuCycles(1)

	if registerType == rtHL {
		c.emulateCpuCycles(2)
	}

	switch bitOperation {
	case 1: // BIT
		c.registers.SetFZ((registerValue & (1 << bit)) != (1 << bit))
		c.registers.SetFN(false)
		c.registers.SetFH(true)
		return
	case 2: // RES
		registerValue &^= 1 << bit
		c.setRegisterPrefixCB(registerType, registerValue)
		return
	case 3: // SET
		registerValue |= 1 << bit
		c.setRegisterPrefixCB(registerType, registerValue)
		return
	}

	switch bit {
	case 0: // RLC
		setC := false
		result := (registerValue << 1) & 0xFF

		if registerValue&(1<<7) != 0 {
			result |= 1
			setC = true
		}

		c.registers.SetFZ(result == 0)
		c.registers.SetFN(false)
		c.registers.SetFH(false)
		c.registers.SetFC(setC)
		c.setRegisterPrefixCB(registerType, result)

		return

	case 1: // RRC
		old := registerValue
		registerValue >>= 1
		registerValue |= old << 7

		c.setRegisterPrefixCB(registerType, registerValue)
		c.registers.SetFZ(registerValue == 0)
		c.registers.SetFN(false)
		c.registers.SetFH(false)
		c.registers.SetFC(old&1 == 1)

		return
	case 2: // RL
		old := registerValue
		registerValue <<= 1
		if c.registers.GetFC() {
			registerValue |= 0x1
		}

		c.setRegisterPrefixCB(registerType, registerValue)
		c.registers.SetFZ(registerValue == 0)
		c.registers.SetFN(false)
		c.registers.SetFH(false)
		c.registers.SetFC(old&0x80 == 0x80)

		return
	case 3: // RR
		old := registerValue
		registerValue >>= 1
		if c.registers.GetFC() {
			registerValue |= 1 << 7
		}

		c.setRegisterPrefixCB(registerType, registerValue)
		c.registers.SetFZ(registerValue == 0)
		c.registers.SetFN(false)
		c.registers.SetFH(false)
		c.registers.SetFC(old&0x1 == 0x1)

		return
	case 4: // SLA
		old := registerValue
		registerValue <<= 1

		c.setRegisterPrefixCB(registerType, registerValue)
		c.registers.SetFZ(registerValue == 0)
		c.registers.SetFN(false)
		c.registers.SetFH(false)
		c.registers.SetFC(old&0x80 == 0x80)

		return
	case 5: // SRA (arithmetic shift to the right)
		u := int8(registerValue) >> 1
		c.setRegisterPrefixCB(registerType, byte(u))
		c.registers.SetFZ(u == 0)
		c.registers.SetFN(false)
		c.registers.SetFH(false)
		c.registers.SetFC(registerValue&0x1 == 0x1)

		return
	case 6: // SWAP (swap high nibble with low nibble)
		registerValue = (registerValue&0xF0)>>4 | (registerValue&0xF)<<4
		c.setRegisterPrefixCB(registerType, registerValue)
		c.registers.SetFZ(registerValue == 0)
		c.registers.SetFN(false)
		c.registers.SetFH(false)
		c.registers.SetFC(false)

		return
	case 7: // SRL (logical shift to the right)
		old := registerValue
		registerValue >>= 1

		c.setRegisterPrefixCB(registerType, registerValue)
		c.registers.SetFZ(registerValue == 0)
		c.registers.SetFN(false)
		c.registers.SetFH(false)
		c.registers.SetFC(old&0x1 == 0x1)

		return
	}

	c.logger.Fatalf("Invalid CB %X", cbOperation)
}

func rlcaExecFunc(c *CPU) {
	u := c.registers.A
	var carry byte
	if u>>7 == 0x1 {
		carry = 0x1
	}

	u = (u << 1) | carry
	c.registers.A = u
	c.registers.SetFZ(false)
	c.registers.SetFN(false)
	c.registers.SetFH(false)
	c.registers.SetFC(carry == 0x1)
}

func rrcaExecFunc(c *CPU) {
	b := c.registers.A & 0x1
	c.registers.A >>= 1
	c.registers.A |= b << 7

	c.registers.SetFZ(false)
	c.registers.SetFN(false)
	c.registers.SetFH(false)
	c.registers.SetFC(b == 0x1)
}

func rlaExecFunc(c *CPU) {
	u := c.registers.A
	var carryFlag byte
	if c.registers.GetFC() {
		carryFlag = 0x1
	}

	carry := (u >> 7) & 0x1

	c.registers.A = (u << 1) | carryFlag
	c.registers.SetFZ(false)
	c.registers.SetFN(false)
	c.registers.SetFH(false)
	c.registers.SetFC(carry == 0x1)
}

func rraExecFunc(c *CPU) {
	var carry byte
	if c.registers.GetFC() {
		carry = 0x1 << 7
	}

	newCarry := c.registers.A & 0x1

	c.registers.A >>= 1
	c.registers.A |= carry
	c.registers.SetFZ(false)
	c.registers.SetFN(false)
	c.registers.SetFH(false)
	c.registers.SetFC(newCarry == 0x1)
}

func stopExecFunc(c *CPU) {
	c.logger.Fatal("Stopping Gameboy...") // Do some research about how to implement this better.
}

func daaExecFunc(c *CPU) {
	u := byte(0)
	fc := 0

	if c.registers.GetFH() || (!c.registers.GetFN() && (c.registers.A&0xF) > 9) {
		u = 6
	}

	if c.registers.GetFC() || (!c.registers.GetFN() && c.registers.A > 0x99) {
		u |= 0x60
		fc = 1
	}

	if c.registers.GetFN() {
		c.registers.A += -u
	} else {
		c.registers.A += u
	}

	c.registers.SetFZ(c.registers.A == 0)
	c.registers.SetFH(false)
	c.registers.SetFC(fc == 1)
}

func cplExecFunc(c *CPU) {
	c.registers.A = ^c.registers.A
	c.registers.SetFN(true)
	c.registers.SetFH(true)
}

func scfExecFunc(c *CPU) {
	c.registers.SetFN(false)
	c.registers.SetFH(false)
	c.registers.SetFC(true)
}

func ccfExecFunc(c *CPU) {
	c.registers.SetFN(false)
	c.registers.SetFH(false)
	c.registers.SetFC(!c.registers.GetFC()) // Flip the carry flag
}

func haltExecFunc(c *CPU) {
	c.Halted = true
}
