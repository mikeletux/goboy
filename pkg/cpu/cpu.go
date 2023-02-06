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

// SetFN sets Subtraction flag (N) from F to bit value. True means 1 whilst false means 0.
func (r *Registers) SetFN(bit bool) {
	r.F = getBitMask(r.F, 6, bit)
}

// SetFH sets Half Carry flag (H) from F to bit value. True means 1 whilst false means 0.
func (r *Registers) SetFH(bit bool) {
	r.F = getBitMask(r.F, 5, bit)
}

// SetFC sets Carry flag (C) from F to bit value. True means 1 whilst false means 0.
func (r *Registers) SetFC(bit bool) {
	r.F = getBitMask(r.F, 4, bit)
}

func getBitMask(regValue byte, bitNumber int, bit bool) byte {
	if bit {
		return regValue | 1<<bitNumber
	}
	return regValue & ^(1 << bitNumber)
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
	FetchedData          uint16
	MemoryDestination    uint16
	DestinationIsMemory  bool
	CurrentOperationCode byte
	CurrentInstruction   Instruction

	Halted   bool
	Stepping bool
}

func Init(bus bus.DataBusInterface, logger log.Logger) *CPU {
	return &CPU{
		registers: &Registers{},
		bus:       bus,
		logger:    log.NewBuiltinStdoutLogger(),
	}
}

func (c *CPU) Step() bool {
	if !c.Halted {
		// Fetch instruction
		c.CurrentOperationCode = c.bus.BusRead(c.registers.GetPCAndIncrement())
		instruction, ok := instructionsMap[c.CurrentOperationCode]
		if !ok {
			c.logger.Fatalf("instruction %d doesn't exist", c.CurrentOperationCode)
		}
		c.CurrentInstruction = instruction
		c.logger.Debugf("instruction to be execute: %s", c.CurrentInstruction.Mnemonic)

		// Fetch data
		err := c.fetchData()
		if err != nil {
			c.logger.Fatal(err)
		}

		// Execute instruction
		execFunc, exist := execInstructionMap[instruction.Type]
		if !exist { // Check if instruction has an entry on execInstructionMap.
			c.logger.Fatalf("instruction %s doesn't exist in execInstructionMap", instruction.Mnemonic)
		}

		if execFunc == nil { // If nil means that the instruction has not been implemented yet.
			c.logger.Fatalf("instruction %s has not been implemented yet", instruction.Mnemonic)
		}

		execFunc(c)
		c.logRegisterValues()

	}
	return true // Check this
}

func (c *CPU) fetchData() error {
	c.MemoryDestination = 0
	c.DestinationIsMemory = false

	switch c.CurrentInstruction.AddressingMode {
	case amImp:
		return nil

	case amR:
		fetchedData, err := c.registers.FetchDataFromRegisters(c.CurrentInstruction.RegisterType1)
		if err != nil {
			return err
		}
		c.FetchedData = fetchedData
		return nil

	case amRnD8:
		c.FetchedData = uint16(c.bus.BusRead(c.registers.GetPCAndIncrement()))
		c.emulateCpuCycles(1)
		return nil

	case amD16:
		var low = c.bus.BusRead(c.registers.GetPCAndIncrement())
		c.emulateCpuCycles(1)

		var high = c.bus.BusRead(c.registers.GetPCAndIncrement())
		c.emulateCpuCycles(1)

		c.FetchedData = uint16(low) | uint16(high)<<8
		return nil

	// To be done still
	default:
		return fmt.Errorf("addressing mode %d doesn't exist", c.CurrentInstruction.AddressingMode)
	}

	return nil // This return should not be reached ever
}

func (c *CPU) emulateCpuCycles(numCycles int) { // TO BE IMPLEMENTED
	return
}

func (c *CPU) logRegisterValues() {
	c.logger.Debugf("AF:%X BC:%X DE:%X HL%X",
		c.registers.GetAF(), c.registers.GetBC(), c.registers.GetDE(), c.registers.GetHL())
}
