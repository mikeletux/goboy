package cpu

import (
	"github.com/mikeletux/goboy/pkg/bus"
	"github.com/mikeletux/goboy/pkg/log"
)

const (
	initARegisterValue  byte   = 0x01
	initFRegisterValue  byte   = 0xB0
	initBRegisterValue  byte   = 0x00
	initCRegisterValue  byte   = 0x13
	initDRegisterValue  byte   = 0x00
	initERegisterValue  byte   = 0xD8
	initHRegisterValue  byte   = 0x01
	initLRegisterValue  byte   = 0x4D
	initSPRegisterValue uint16 = 0xFFFE
	initPCRegisterValue uint16 = 0x0100
)

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

	EnableMasterInterruptions bool
	EnablingIme               bool
	InterruptionsRegister     byte

	Halted   bool
	Stepping bool

	tick uint64
}

func Init(bus bus.DataBusInterface, logger log.Logger) *CPU {
	return &CPU{
		registers: &Registers{
			A:  initARegisterValue,
			F:  initFRegisterValue,
			B:  initBRegisterValue,
			C:  initCRegisterValue,
			D:  initDRegisterValue,
			E:  initERegisterValue,
			H:  initHRegisterValue,
			L:  initLRegisterValue,
			SP: initSPRegisterValue,
			PC: initPCRegisterValue,
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

		c.logRegistersGameboyDoctor(instructionPC)
		//c.logRegisterValues(instructionPC) // used for debugging purposes

		c.dbgUpdate() // Useful for debugging roms
		c.dbgPrint()

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
		c.tick++

	} else {
		// CPU is halted at this point
		c.emulateCpuCycles(1)

		if c.InterruptionsRegister != 0x0 {
			c.Halted = false
		}
	}

	if c.EnableMasterInterruptions {
		// handle interrupt
		c.EnablingIme = false
	}

	if c.EnablingIme {
		c.EnableMasterInterruptions = true
	}

	return true // Check this
}

func (c *CPU) emulateCpuCycles(numCycles int) { // TO BE IMPLEMENTED
	return
}
