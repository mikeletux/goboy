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
	0x18: {Type: inJr, AddressingMode: amD8, Mnemonic: "JR r8", Condition: ctNone, execFunc: jrExecFunc},                             // JR r8
	0x1A: {Type: inLd, AddressingMode: amRnMR, RegisterType1: rtA, RegisterType2: rtDE, Mnemonic: "LD A,(DE)", execFunc: ldExecFunc}, // LD A,(DE)
	0x1E: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtE, Mnemonic: "LD E,d8", execFunc: ldExecFunc},                        // LD E,d8
	// 0x2
	0x20: {Type: inJr, AddressingMode: amD8, Mnemonic: "JR NZ,r8", Condition: ctNZ, execFunc: jrExecFunc},                              // JR NZ,r8
	0x21: {Type: inLd, AddressingMode: amRnD16, RegisterType1: rtHL, Mnemonic: "LD HL,d16", execFunc: ldExecFunc},                      // LD HL,d16
	0x22: {Type: inLd, AddressingMode: amHLInR, RegisterType1: rtHL, RegisterType2: rtA, Mnemonic: "LD (HL+),A", execFunc: ldExecFunc}, // LD (HL+),A
	0x26: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtH, Mnemonic: "LD H,d8", execFunc: ldExecFunc},                          // LD H,d8
	0x28: {Type: inJr, AddressingMode: amD8, Mnemonic: "JR Z,r8", Condition: ctZ, execFunc: jrExecFunc},                                // JR Z,r8
	0x2A: {Type: inLd, AddressingMode: amRnHLI, RegisterType1: rtA, RegisterType2: rtHL, Mnemonic: "LD A,(HL+)", execFunc: ldExecFunc}, // LD A,(HL+)
	0x2E: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtL, Mnemonic: "LD L,d8", execFunc: ldExecFunc},                          // LD L,d8
	// 0x3
	0x30: {Type: inJr, AddressingMode: amD8, Mnemonic: "JR NC,r8", Condition: ctNC, execFunc: jrExecFunc},                              // JR NC,r8
	0x31: {Type: inLd, AddressingMode: amRnD16, RegisterType1: rtSP, Mnemonic: "LD SP,d16", execFunc: ldExecFunc},                      // LD SP,d16
	0x32: {Type: inLd, AddressingMode: amHLDnR, RegisterType1: rtHL, RegisterType2: rtA, Mnemonic: "LD (HL-),A", execFunc: ldExecFunc}, // LD (HL-),A
	0x36: {Type: inLd, AddressingMode: amMRnD8, RegisterType1: rtHL, Mnemonic: "LD (HL),d8", execFunc: ldExecFunc},                     // LD (HL),d8
	0x38: {Type: inJr, AddressingMode: amD8, Mnemonic: "JR C,r8", Condition: ctC, execFunc: jrExecFunc},                                // JR C,r8
	0x3A: {Type: inLd, AddressingMode: amRnHLD, RegisterType1: rtA, RegisterType2: rtHL, Mnemonic: "LD A,(HL-)", execFunc: ldExecFunc}, // LD A,(HL-)
	0x3E: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtA, Mnemonic: "LD A,d8", execFunc: ldExecFunc},                          // LD A,d8
	// 0x4
	0x40: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtB, RegisterType2: rtB, Mnemonic: "LD B,B", execFunc: ldExecFunc},      //LD B,B
	0x41: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtB, RegisterType2: rtC, Mnemonic: "LD B,C", execFunc: ldExecFunc},      // LD B,C
	0x42: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtB, RegisterType2: rtD, Mnemonic: "LD B,D", execFunc: ldExecFunc},      // LD B,D
	0x43: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtB, RegisterType2: rtE, Mnemonic: "LD B,E", execFunc: ldExecFunc},      // LD B,E
	0x44: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtB, RegisterType2: rtH, Mnemonic: "LD B,H", execFunc: ldExecFunc},      // LD B,H
	0x45: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtB, RegisterType2: rtL, Mnemonic: "LD B,L", execFunc: ldExecFunc},      // LD B,L
	0x46: {Type: inLd, AddressingMode: amRnMR, RegisterType1: rtB, RegisterType2: rtHL, Mnemonic: "LD B,(HL)", execFunc: ldExecFunc}, // LD B,(HL)
	0x47: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtB, RegisterType2: rtA, Mnemonic: "LD B,A", execFunc: ldExecFunc},      // LD B,A
	0x48: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtC, RegisterType2: rtB, Mnemonic: "LD C,B", execFunc: ldExecFunc},      // LD C,B
	0x49: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtC, RegisterType2: rtC, Mnemonic: "LD C,C", execFunc: ldExecFunc},      // LD C,C
	0x4A: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtC, RegisterType2: rtD, Mnemonic: "LD C,D", execFunc: ldExecFunc},      // LD C,D
	0x4B: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtC, RegisterType2: rtE, Mnemonic: "LD C,E", execFunc: ldExecFunc},      // LD C,E
	0x4C: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtC, RegisterType2: rtH, Mnemonic: "LD C,H", execFunc: ldExecFunc},      // LD C,H
	0x4D: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtC, RegisterType2: rtL, Mnemonic: "LD C,L", execFunc: ldExecFunc},      // LD C,L
	0x4E: {Type: inLd, AddressingMode: amRnMR, RegisterType1: rtC, RegisterType2: rtHL, Mnemonic: "LD C,(HL)", execFunc: ldExecFunc}, // LD C,(HL)
	0x4F: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtC, RegisterType2: rtA, Mnemonic: "LD C,A", execFunc: ldExecFunc},      // LD C,A
	// 0x5
	0x50: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtD, RegisterType2: rtB, Mnemonic: "LD D,B", execFunc: ldExecFunc},      // LD D,B
	0x51: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtD, RegisterType2: rtC, Mnemonic: "LD D,C", execFunc: ldExecFunc},      // LD D,C
	0x52: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtD, RegisterType2: rtD, Mnemonic: "LD D,D", execFunc: ldExecFunc},      // LD D,D
	0x53: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtD, RegisterType2: rtE, Mnemonic: "LD D,E", execFunc: ldExecFunc},      // LD D,E
	0x54: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtD, RegisterType2: rtH, Mnemonic: "LD D,H", execFunc: ldExecFunc},      // LD D,H
	0x55: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtD, RegisterType2: rtL, Mnemonic: "LD D,L", execFunc: ldExecFunc},      // LD D,L
	0x56: {Type: inLd, AddressingMode: amRnMR, RegisterType1: rtD, RegisterType2: rtHL, Mnemonic: "LD D,(HL)", execFunc: ldExecFunc}, // LD D,(HL)
	0x57: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtD, RegisterType2: rtA, Mnemonic: "LD D,A", execFunc: ldExecFunc},      // LD D,A
	0x58: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtE, RegisterType2: rtB, Mnemonic: "LD E,B", execFunc: ldExecFunc},      // LD E,B
	0x59: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtE, RegisterType2: rtC, Mnemonic: "LD E,C", execFunc: ldExecFunc},      // LD E,C
	0x5A: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtE, RegisterType2: rtD, Mnemonic: "LD E,D", execFunc: ldExecFunc},      // LD E,D
	0x5B: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtE, RegisterType2: rtE, Mnemonic: "LD E,E", execFunc: ldExecFunc},      // LD E,E
	0x5C: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtE, RegisterType2: rtH, Mnemonic: "LD E,H", execFunc: ldExecFunc},      // LD E,H
	0x5D: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtE, RegisterType2: rtL, Mnemonic: "LD E,L", execFunc: ldExecFunc},      // LD E,L
	0x5E: {Type: inLd, AddressingMode: amRnMR, RegisterType1: rtE, RegisterType2: rtHL, Mnemonic: "LD E,(HL)", execFunc: ldExecFunc}, // LD E,(HL)
	0x5F: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtE, RegisterType2: rtA, Mnemonic: "LD E,A", execFunc: ldExecFunc},      // LD E,A
	// 0x6
	0x60: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtH, RegisterType2: rtB, Mnemonic: "LD H,B", execFunc: ldExecFunc},      // LD H,B
	0x61: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtH, RegisterType2: rtC, Mnemonic: "LD H,C", execFunc: ldExecFunc},      // LD H,C
	0x62: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtH, RegisterType2: rtD, Mnemonic: "LD H,D", execFunc: ldExecFunc},      // LD H,D
	0x63: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtH, RegisterType2: rtE, Mnemonic: "LD H,E", execFunc: ldExecFunc},      // LD H,E
	0x64: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtH, RegisterType2: rtH, Mnemonic: "LD H,H", execFunc: ldExecFunc},      // LD H,H
	0x65: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtH, RegisterType2: rtL, Mnemonic: "LD H,L", execFunc: ldExecFunc},      // LD H,L
	0x66: {Type: inLd, AddressingMode: amRnMR, RegisterType1: rtH, RegisterType2: rtHL, Mnemonic: "LD H,(HL)", execFunc: ldExecFunc}, // LD H,(HL)
	0x67: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtH, RegisterType2: rtA, Mnemonic: "LD H,A", execFunc: ldExecFunc},      // LD H,A
	0x68: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtL, RegisterType2: rtB, Mnemonic: "LD L,B", execFunc: ldExecFunc},      // LD L,B
	0x69: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtL, RegisterType2: rtC, Mnemonic: "LD L,C", execFunc: ldExecFunc},      // LD L,C
	0x6A: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtL, RegisterType2: rtD, Mnemonic: "LD L,D", execFunc: ldExecFunc},      // LD L,D
	0x6B: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtL, RegisterType2: rtE, Mnemonic: "LD L,E", execFunc: ldExecFunc},      // LD L,E
	0x6C: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtL, RegisterType2: rtH, Mnemonic: "LD L,H", execFunc: ldExecFunc},      // LD L,H
	0x6D: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtL, RegisterType2: rtL, Mnemonic: "LD L,L", execFunc: ldExecFunc},      // LD L,L
	0x6E: {Type: inLd, AddressingMode: amRnMR, RegisterType1: rtL, RegisterType2: rtHL, Mnemonic: "LD L,(HL)", execFunc: ldExecFunc}, // LD L,(HL)
	0x6F: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtL, RegisterType2: rtA, Mnemonic: "LD L,A", execFunc: ldExecFunc},      // LD L,A
	// 0x7
	0x70: {Type: inLd, AddressingMode: amMRnR, RegisterType1: rtHL, RegisterType2: rtB, Mnemonic: "LD (HL),B", execFunc: ldExecFunc}, // LD (HL),B
	0x71: {Type: inLd, AddressingMode: amMRnR, RegisterType1: rtHL, RegisterType2: rtC, Mnemonic: "LD (HL),C", execFunc: ldExecFunc}, // LD (HL),C
	0x72: {Type: inLd, AddressingMode: amMRnR, RegisterType1: rtHL, RegisterType2: rtD, Mnemonic: "LD (HL),D", execFunc: ldExecFunc}, // LD (HL),D
	0x73: {Type: inLd, AddressingMode: amMRnR, RegisterType1: rtHL, RegisterType2: rtE, Mnemonic: "LD (HL),E", execFunc: ldExecFunc}, // LD (HL),E
	0x74: {Type: inLd, AddressingMode: amMRnR, RegisterType1: rtHL, RegisterType2: rtH, Mnemonic: "LD (HL),H", execFunc: ldExecFunc}, // LD (HL),H
	0x75: {Type: inLd, AddressingMode: amMRnR, RegisterType1: rtHL, RegisterType2: rtL, Mnemonic: "LD (HL),L", execFunc: ldExecFunc}, // LD (HL),L
	//	0x66:
	0x77: {Type: inLd, AddressingMode: amMRnR, RegisterType1: rtHL, RegisterType2: rtA, Mnemonic: "LD (HL),A", execFunc: ldExecFunc}, // LD (HL),A
	0x78: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtB, Mnemonic: "LD A,B", execFunc: ldExecFunc},      // LD A,B
	0x79: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtC, Mnemonic: "LD A,C", execFunc: ldExecFunc},      // LD A,C
	0x7A: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtD, Mnemonic: "LD A,D", execFunc: ldExecFunc},      // LD A,D
	0x7B: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtE, Mnemonic: "LD A,E", execFunc: ldExecFunc},      // LD A,E
	0x7C: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtH, Mnemonic: "LD A,H", execFunc: ldExecFunc},      // LD A,H
	0x7D: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtL, Mnemonic: "LD A,L", execFunc: ldExecFunc},      // LD A,L
	0x7E: {Type: inLd, AddressingMode: amRnMR, RegisterType1: rtA, RegisterType2: rtHL, Mnemonic: "LD A,(HL)", execFunc: ldExecFunc}, // LD A,(HL)
	0x7F: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtA, Mnemonic: "LD A,A", execFunc: ldExecFunc},      // LD A,A
	// 0x8
	// 0x9
	// 0xA
	0xAF: {Type: inXor, AddressingMode: amR, RegisterType1: rtA, Mnemonic: "XOR A", execFunc: xorExecFunc}, // XOR A
	// 0xB
	// 0xC
	0xC0: {Type: inRet, Mnemonic: "RET NZ", Condition: ctNZ, execFunc: retExecFunc},                               // RET NZ
	0xC1: {Type: inPop, AddressingMode: amImp, RegisterType1: rtBC, Mnemonic: "POP BC", execFunc: popExecFunc},    // POP BC
	0xC2: {Type: inJp, AddressingMode: amD16, Mnemonic: "JP NZ,a16", Condition: ctNZ, execFunc: jpExecFunc},       // JP NZ,a16
	0xC3: {Type: inJp, AddressingMode: amD16, Mnemonic: "JP a16", execFunc: jpExecFunc},                           // JP a16
	0xC4: {Type: inCall, AddressingMode: amD16, Mnemonic: "CALL NZ,a16", Condition: ctNZ, execFunc: callExecFunc}, // CALL NZ,a16
	0xC5: {Type: inPush, AddressingMode: amImp, RegisterType1: rtBC, Mnemonic: "PUSH BC", execFunc: pushExecFunc}, // PUSH BC
	0xC7: {Type: inRst, AddressingMode: amImp, Mnemonic: "RST 00H", Parameter: 0x0, execFunc: rstExecFunc},        // RST 00H
	0xC8: {Type: inRet, Mnemonic: "RET Z", Condition: ctZ, execFunc: retExecFunc},                                 // RET Z
	0xC9: {Type: inRet, Mnemonic: "RET", Condition: ctNone, execFunc: retExecFunc},                                // RET
	0xCA: {Type: inJp, AddressingMode: amD16, Mnemonic: "JP Z,a16", Condition: ctZ, execFunc: jpExecFunc},         // JP Z,a16
	0xCC: {Type: inCall, AddressingMode: amD16, Mnemonic: "CALL Z,a16", Condition: ctZ, execFunc: callExecFunc},   // CALL Z,a16
	0xCD: {Type: inCall, AddressingMode: amD16, Mnemonic: "CALL a16", Condition: ctNone, execFunc: callExecFunc},  // CALL a16
	0xCF: {Type: inRst, AddressingMode: amImp, Mnemonic: "RST 08H", Parameter: 0x08, execFunc: rstExecFunc},       // RST 08H
	// 0xD
	0xD0: {Type: inRet, Mnemonic: "RET NC", Condition: ctNC, execFunc: retExecFunc},                               // RET NC
	0xD1: {Type: inPop, AddressingMode: amImp, RegisterType1: rtDE, Mnemonic: "POP DE", execFunc: popExecFunc},    // POP DE
	0xD2: {Type: inJp, AddressingMode: amD16, Mnemonic: "JP NC,a16", Condition: ctNC, execFunc: jpExecFunc},       // JP NC,a16
	0xD4: {Type: inCall, AddressingMode: amD16, Mnemonic: "CALL NC,a16", Condition: ctNC, execFunc: callExecFunc}, // CALL NC,a16
	0xD5: {Type: inPush, AddressingMode: amImp, RegisterType1: rtDE, Mnemonic: "PUSH DE", execFunc: pushExecFunc}, // PUSH DE
	0xD7: {Type: inRst, AddressingMode: amImp, Mnemonic: "RST 10H", Parameter: 0x10, execFunc: rstExecFunc},       // RST 10H
	0xD8: {Type: inRet, Mnemonic: "RET C", Condition: ctC, execFunc: retExecFunc},                                 // RET C
	0xD9: {Type: inReti, Mnemonic: "RETI", execFunc: retiExecFunc},                                                // RETI
	0xDA: {Type: inJp, AddressingMode: amD16, Mnemonic: "JP C,a16", Condition: ctC, execFunc: jpExecFunc},         // JP C,a16
	0xDC: {Type: inCall, AddressingMode: amD16, Mnemonic: "CALL C,a16", Condition: ctC, execFunc: callExecFunc},   // CALL C,a16
	0xDF: {Type: inRst, AddressingMode: amImp, Mnemonic: "RST 18H", Parameter: 0x18, execFunc: rstExecFunc},       // RST 18H
	// 0xE
	0xE0: {Type: inLdh, AddressingMode: amA8nR, RegisterType2: rtA, Mnemonic: "LDH (a8),A", execFunc: ldhExecFunc},                 // LDH (a8),A
	0xE1: {Type: inPop, AddressingMode: amImp, RegisterType1: rtHL, Mnemonic: "POP HL", execFunc: popExecFunc},                     // POP HL
	0xE2: {Type: inLd, AddressingMode: amMRnR, RegisterType1: rtC, RegisterType2: rtA, Mnemonic: "LD (C),A", execFunc: ldExecFunc}, // LD (C),A
	0xE5: {Type: inPush, AddressingMode: amImp, RegisterType1: rtHL, Mnemonic: "PUSH HL", execFunc: pushExecFunc},                  // PUSH HL
	0xE7: {Type: inRst, AddressingMode: amImp, Mnemonic: "RST 20H", Parameter: 0x20, execFunc: rstExecFunc},                        // RST 20H
	0xE9: {Type: inJp, AddressingMode: amMR, RegisterType1: rtHL, Mnemonic: "JP (HL)", Condition: ctNone, execFunc: jpExecFunc},    // JP (HL)
	0xEA: {Type: inLd, AddressingMode: amA16nR, RegisterType2: rtA, Mnemonic: "LD (a16),A", execFunc: ldExecFunc},                  // LD (a16),A
	0xEF: {Type: inRst, AddressingMode: amImp, Mnemonic: "RST 28H", Parameter: 0x28, execFunc: rstExecFunc},                        // RST 28H
	// 0xF
	0xF0: {Type: inLdh, AddressingMode: amRnA8, RegisterType1: rtA, Mnemonic: "LDH A,(a8)", execFunc: ldhExecFunc},                 // LDH A,(a8)
	0xF1: {Type: inPop, AddressingMode: amImp, RegisterType1: rtAF, Mnemonic: "POP AF", execFunc: popExecFunc},                     // POP AF
	0xF2: {Type: inLd, AddressingMode: amRnMR, RegisterType1: rtA, RegisterType2: rtC, Mnemonic: "LD A,(C)", execFunc: ldExecFunc}, // LD A,(C)
	0xF3: {Type: inDi, Mnemonic: "DI", execFunc: diExecFunc},                                                                       // DI
	0xF5: {Type: inPush, AddressingMode: amImp, RegisterType1: rtAF, Mnemonic: "PUSH AF", execFunc: pushExecFunc},                  // PUSH AF
	0xF7: {Type: inRst, AddressingMode: amImp, Mnemonic: "RST 30H", Parameter: 0x30, execFunc: rstExecFunc},                        // RST 30H
	// 0xF8: {Type: inLd, AddressingMode: amHLnSPR, Mnemonic: "LD HL,SP+r8", execFunc: ldExecFunc},                                     // LD HL,SP+r8
	0xF9: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtSP, RegisterType2: rtHL, Mnemonic: "LD SP,HL", execFunc: ldExecFunc}, // LD SP,HL
	0xFA: {Type: inLd, AddressingMode: amRnA16, RegisterType1: rtA, Mnemonic: "LD A,(a16)", execFunc: ldExecFunc},                   // LD A,(a16)
	0xFF: {Type: inRst, AddressingMode: amImp, Mnemonic: "RST 38H", Parameter: 0x38, execFunc: rstExecFunc},                         // RST 38H
}
