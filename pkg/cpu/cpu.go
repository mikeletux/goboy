package cpu

import (
	"fmt"
	"github.com/mikeletux/goboy/pkg/bus"
	"github.com/mikeletux/goboy/pkg/log"
)

// Registers model all CPU registers from Gameboy
type Registers struct {
	A  byte   // Accumulator
	F  byte   // Flags - Bit 7 (z) Zero flag | Bit 6 (n) Subtraction flag (BCD) | Bit 5 (h) Half Carry flag (BCD) | Bit 4 (c) Carry Flag
	B  byte   // BC
	C  byte   // BC
	D  byte   // DE
	E  byte   // DE
	H  byte   // HL
	L  byte   // HL
	SP uint16 // Stack Pointer
	PC uint16 // Program Counter
}

// GetAF returns a 16Bit value from CPU registers A and F
func (r *Registers) GetAF() uint16 {
	return uint16(r.A)<<8 | uint16(r.F)
}

// GetBC returns a 16Bit value from CPU registers B and C
func (r *Registers) GetBC() uint16 {
	return uint16(r.B)<<8 | uint16(r.C)
}

// GetDE returns a 16Bit value from CPU registers D and E
func (r *Registers) GetDE() uint16 {
	return uint16(r.D)<<8 | uint16(r.E)
}

// GetHL returns a 16Bit value from CPU registers H and L
func (r *Registers) GetHL() uint16 {
	return uint16(r.H)<<8 | uint16(r.L)
}

// SetAF sets a 16Bit value between registers A and F
func (r *Registers) SetAF(value uint16) {
	r.A, r.F = getHighAndLowBytes(value)
}

// SetBC sets a 16Bit value between registers B and C
func (r *Registers) SetBC(value uint16) {
	r.B, r.C = getHighAndLowBytes(value)
}

// SetDE sets a 16Bit value between registers D and E
func (r *Registers) SetDE(value uint16) {
	r.D, r.E = getHighAndLowBytes(value)
}

// SetHL sets a 16Bit value between registers H and L
func (r *Registers) SetHL(value uint16) {
	r.H, r.L = getHighAndLowBytes(value)
}

// SetFZ sets Zero flag (Z) from F to bit value. True means 1 whilst false means 0.
func (r *Registers) SetFZ(bit bool) {
	r.F = getBitMask(r.F, 7, bit)
}

// GetFZ returns if the Z bit is set or not.
func (r *Registers) GetFZ() bool {
	return getBitRegister(r.F, 7)
}

// SetFN sets Subtraction flag (N) from F to bit value. True means 1 whilst false means 0.
func (r *Registers) SetFN(bit bool) {
	r.F = getBitMask(r.F, 6, bit)
}

// GetFN returns if the N bit is set or not.
func (r *Registers) GetFN() bool {
	return getBitRegister(r.F, 6)
}

// SetFH sets Half Carry flag (H) from F to bit value. True means 1 whilst false means 0.
func (r *Registers) SetFH(bit bool) {
	r.F = getBitMask(r.F, 5, bit)
}

// GetFH returns if the H bit is set or not.
func (r *Registers) GetFH() bool {
	return getBitRegister(r.F, 5)
}

// SetFC sets Carry flag (C) from F to bit value. True means 1 whilst false means 0.
func (r *Registers) SetFC(bit bool) {
	r.F = getBitMask(r.F, 4, bit)
}

// GetFC returns if the C bit is set or not.
func (r *Registers) GetFC() bool {
	return getBitRegister(r.F, 4)
}

func getBitMask(regValue byte, bitNumber int, bit bool) byte {
	if bit {
		return regValue | 1<<bitNumber
	}
	return regValue & ^(1 << bitNumber)
}

func getBitRegister(regValue byte, bitNumber int) bool {
	if (regValue & (1 << bitNumber)) == 0x0 {
		return false
	}
	return true
}

// FetchDataFromRegisters returns the register value given its register type constant
func (r *Registers) FetchDataFromRegisters(registerType int) (uint16, error) {
	switch registerType {
	case rtA:
		return uint16(r.A), nil
	case rtF:
		return uint16(r.F), nil
	case rtB:
		return uint16(r.B), nil
	case rtC:
		return uint16(r.C), nil
	case rtD:
		return uint16(r.D), nil
	case rtE:
		return uint16(r.E), nil
	case rtH:
		return uint16(r.H), nil
	case rtL:
		return uint16(r.L), nil
	case rtAF:
		return r.GetAF(), nil
	case rtBC:
		return r.GetBC(), nil
	case rtDE:
		return r.GetDE(), nil
	case rtHL:
		return r.GetHL(), nil
	case rtSP:
		return r.SP, nil
	case rtPC:
		return r.PC, nil
	}

	return 0, fmt.Errorf("the processor register provided doesn't exist")
}

// SetDataToRegisters sets data to a register given its register constant.
func (r *Registers) SetDataToRegisters(registerType int, data uint16) error {
	switch registerType {
	case rtA:
		r.A = byte(0x00FF & data) // I think doing & 0x00FF is not needed but it makes clearer which part of the two bytes we want
	case rtF:
		r.F = byte(0x00FF & data)
	case rtB:
		r.B = byte(0x00FF & data)
	case rtC:
		r.C = byte(0x00FF & data)
	case rtD:
		r.D = byte(0x00FF & data)
	case rtE:
		r.E = byte(0x00FF & data)
	case rtH:
		r.H = byte(0x00FF & data)
	case rtL:
		r.L = byte(0x00FF & data)
	case rtAF:
		r.SetAF(data)
	case rtBC:
		r.SetBC(data)
	case rtDE:
		r.SetDE(data)
	case rtHL:
		r.SetHL(data)
	case rtSP:
		r.SP = data
	case rtPC:
		r.PC = data
	default:
		return fmt.Errorf("the processor register provided doesn't exist")
	}

	return nil
}

// GetPCAndIncrement returns the PC and increments it by 1
func (r *Registers) GetPCAndIncrement() (pc uint16) {
	pc = r.PC
	r.PC++
	return
}

func getHighAndLowBytes(value uint16) (high, low byte) {
	high, low = byte(value>>8), byte(value&0xFF)
	return
}

type CPU struct {
	registers *Registers
	bus       bus.DataBusInterface
	logger    log.Logger

	// Current fetch
	FetchedData               uint16
	MemoryDestination         uint16
	DestinationIsMemory       bool
	CurrentOperationCode      byte
	CurrentInstruction        Instruction
	EnableMasterInterruptions bool

	Halted   bool
	Stepping bool
}

func Init(bus bus.DataBusInterface, logger log.Logger) *CPU {
	return &CPU{
		registers: &Registers{
			PC: cpuEntryPoint, // The Gameboy entry point is 0x100
		},
		bus:    bus,
		logger: logger,
	}
}

func (c *CPU) Step() bool {
	if !c.Halted {
		// Fetch instruction
		instructionPC := c.registers.PC // used for debugging purposes
		c.CurrentOperationCode = c.bus.BusRead(c.registers.GetPCAndIncrement())
		instruction, ok := instructionsMap[c.CurrentOperationCode]
		if !ok {
			c.logger.Fatalf("instruction with code %X doesn't exist", c.CurrentOperationCode)
		}
		c.CurrentInstruction = instruction
		c.logRegisterValues(instructionPC) // used for debugging purposes

		// Fetch data
		err := c.fetchData()
		if err != nil {
			c.logger.Fatal(err)
		}

		// Execute instruction
		execFunc := instruction.execFunc

		if execFunc == nil { // If nil means that the instruction has not been implemented yet.
			c.logger.Fatalf("instruction %s has not been implemented yet", instruction.Mnemonic)
		}

		execFunc(c)

	}
	return true // Check this
}

func (c *CPU) emulateCpuCycles(numCycles int) { // TO BE IMPLEMENTED
	return
}
