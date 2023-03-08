package cpu

const (
	vblankInterruptFlag  byte = 0x1
	lcdStatInterruptFlag byte = 0x2
	timerInterruptFlag   byte = 0x4
	serialInterruptFlag  byte = 0x8
	joypadInterruptFlag  byte = 0x16
)

const (
	vblankInterruptAddr  uint16 = 0x40
	lcdStatInterruptAddr uint16 = 0x48
	timerInterruptAddr   uint16 = 0x50
	serialInterruptAddr  uint16 = 0x58
	joypadInterruptAddr  uint16 = 0x60
)

func (c *CPU) pushPCToStack(address uint16) {
	c.stackPush16(c.registers.PC)
	c.registers.PC = address
}

func (c *CPU) interruptCheck(addressToJump uint16, interruptType byte) bool {
	ieRegister := c.bus.BusRead(interruptEnableAddr)
	ifRegister := c.bus.BusRead(interruptFlagIOAddr)

	if ifRegister&interruptType == interruptType &&
		ieRegister&interruptType == interruptType {

		c.pushPCToStack(addressToJump)
		c.bus.BusWrite(interruptFlagIOAddr, ifRegister & ^interruptType)
		c.Halted = false

		return true
	}

	return false
}

func (c *CPU) handleInterruptions() {
	if c.interruptCheck(vblankInterruptAddr, vblankInterruptFlag) {
	} else if c.interruptCheck(lcdStatInterruptAddr, lcdStatInterruptFlag) {
	} else if c.interruptCheck(timerInterruptAddr, timerInterruptFlag) {
	} else if c.interruptCheck(serialInterruptAddr, serialInterruptFlag) {
	} else if c.interruptCheck(joypadInterruptAddr, joypadInterruptFlag) {
	}
}

func (c *CPU) requestInterrupt(interruptType byte) {
	interrupts := c.bus.BusRead(interruptFlagIOAddr)
	c.bus.BusWrite(interruptFlagIOAddr, interrupts|interruptType)
}
