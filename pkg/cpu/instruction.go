package cpu

type Instruction struct {
	// Type of instruction to be used.
	Type int
	// Addressing mode to be used. I.e: register to register, single register, etc.
	AddressingMode int
	// First register to be used in the instruction if any.
	RegisterType1 int
	// Second register to be used in the instruction if any.
	RegisterType2 int
	// Condition relative to the instruction if any. This will be used for JP and CALL instructions.
	Condition int
	// Parameter is specially used for CB.
	Parameter byte
	// Mnemonic is the human-readable instruction.
	Mnemonic string
	// execFunc is the function that will carry out the instruction changes in the CPU.
	execFunc func(c *CPU)
}

var instructionsMap = map[byte]Instruction{
	// 0x0
	0x00: {Type: inNop, Mnemonic: "NOP", execFunc: nopExecFunc},                                                                      // NOP
	0x01: {Type: inLd, AddressingMode: amRnD16, RegisterType1: rtBC, Mnemonic: "LD BC,d16", execFunc: ldExecFunc},                    // LD BC,d16
	0x02: {Type: inLd, AddressingMode: amMRnR, RegisterType1: rtBC, RegisterType2: rtA, Mnemonic: "LD (BC),A", execFunc: ldExecFunc}, // LD (BC),A
	0x05: {Type: inDec, AddressingMode: amR, RegisterType1: rtB, Mnemonic: "DEC B", execFunc: nil},                                   // DEC B
	0x06: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtB, Mnemonic: "LD B,d8", execFunc: ldExecFunc},                        // LD B,d8
	0x08: {Type: inLd, AddressingMode: amA16nR, RegisterType2: rtSP, Mnemonic: "LD (a16),SP", execFunc: ldExecFunc},                  // LD (a16),SP
	0x0A: {Type: inLd, AddressingMode: amRnMR, RegisterType1: rtA, RegisterType2: rtBC, Mnemonic: "LD A,(BC)", execFunc: ldExecFunc}, // LD A,(BC)
	0x0E: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtC, Mnemonic: "LD C, d8", execFunc: ldExecFunc},                       // LD C, d8
	// 0x1
	0x11: {Type: inLd, AddressingMode: amRnD16, RegisterType1: rtDE, Mnemonic: "LD DE,d16", execFunc: ldExecFunc},                    // LD DE,d16
	0x12: {Type: inLd, AddressingMode: amMRnR, RegisterType1: rtDE, RegisterType2: rtA, Mnemonic: "LD (DE),A", execFunc: ldExecFunc}, // LD (DE),A
	0x16: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtD, Mnemonic: "LD D,d8", execFunc: ldExecFunc},                        // LD D,d8
	0x1A: {Type: inLd, AddressingMode: amRnMR, RegisterType1: rtA, RegisterType2: rtDE, Mnemonic: "LD A,(DE)", execFunc: ldExecFunc}, // LD A,(DE)
	0x1E: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtE, Mnemonic: "LD E,d8", execFunc: ldExecFunc},                        // LD E,d8
	// 0x2
	0x21: {Type: inLd, AddressingMode: amRnD16, RegisterType1: rtHL, Mnemonic: "LD HL,d16", execFunc: ldExecFunc},                      // LD HL,d16
	0x22: {Type: inLd, AddressingMode: amHLInR, RegisterType1: rtHL, RegisterType2: rtA, Mnemonic: "LD (HL+),A", execFunc: ldExecFunc}, // LD (HL+),A
	0x26: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtH, Mnemonic: "LD H,d8", execFunc: ldExecFunc},                          // LD H,d8
	0x2A: {Type: inLd, AddressingMode: amRnHLI, RegisterType1: rtA, RegisterType2: rtHL, Mnemonic: "LD A,(HL+)", execFunc: ldExecFunc}, // LD A,(HL+)
	0x2E: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtL, Mnemonic: "LD L,d8", execFunc: ldExecFunc},                          // LD L,d8
	// 0x3 CONTINUE HERE!!!
	// 0x4
	// 0x5
	// 0x6
	// 0x7
	// 0x8
	// 0x9
	// 0xA
	0xAF: {Type: inXor, AddressingMode: amR, RegisterType1: rtA, Mnemonic: "XOR A", execFunc: xorExecFunc}, // 0xAF XOR A
	// 0xB
	// 0xC
	0xC3: {Type: inJp, AddressingMode: amD16, Mnemonic: "JP a16", execFunc: jpExecFunc}, // 0xC3 JP a16
	// 0xD
	// 0xE
	// 0xF
	0xF3: {Type: inDi, Mnemonic: "DI", execFunc: diExecFunc}, // DI
}
