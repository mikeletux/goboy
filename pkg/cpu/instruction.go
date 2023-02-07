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
	0x00: {Type: inNop, Mnemonic: "NOP", execFunc: nopExecFunc},                                         // 0x00 NOP
	0x05: {Type: inDec, AddressingMode: amR, RegisterType1: rtB, Mnemonic: "DEC B", execFunc: nil},      // 0x05 DEC B
	0x0E: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtC, Mnemonic: "LD C, d8", execFunc: nil}, // 0x0E LD C, d8
	// 0x1
	// 0x2
	// 0x3
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
