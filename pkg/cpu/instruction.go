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
	0x03: {Type: inInc, AddressingMode: amR, RegisterType1: rtBC, Mnemonic: "INC BC", execFunc: incExecFunc},                         // INC BC
	0x04: {Type: inInc, AddressingMode: amR, RegisterType1: rtB, Mnemonic: "INC B", execFunc: incExecFunc},                           // INC B
	0x05: {Type: inDec, AddressingMode: amR, RegisterType1: rtB, Mnemonic: "DEC B", execFunc: decExecFunc},                           // DEC B
	0x06: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtB, Mnemonic: "LD B,d8", execFunc: ldExecFunc},                        // LD B,d8
	0x07: {Type: inRlca, AddressingMode: amImp, Mnemonic: "RLCA", execFunc: rlcaExecFunc},                                            // RLCA
	0x08: {Type: inLd, AddressingMode: amA16nR, RegisterType2: rtSP, Mnemonic: "LD (a16),SP", execFunc: ldExecFunc},                  // LD (a16),SP
	0x09: {Type: inAdd, AddressingMode: amRnR, RegisterType1: rtHL, RegisterType2: rtBC, Mnemonic: "ADD HL,BC", execFunc: addExecFunc}, // ADD HL,BC
	0x0A: {Type: inLd, AddressingMode: amRnMR, RegisterType1: rtA, RegisterType2: rtBC, Mnemonic: "LD A,(BC)", execFunc: ldExecFunc}, // LD A,(BC)
	0x0B: {Type: inDec, AddressingMode: amR, RegisterType1: rtBC, Mnemonic: "DEC BC", execFunc: decExecFunc},                         // DEC BC
	0x0C: {Type: inInc, AddressingMode: amR, RegisterType1: rtC, Mnemonic: "INC C", execFunc: incExecFunc},                           // INC C
	0x0D: {Type: inDec, AddressingMode: amR, RegisterType1: rtC, Mnemonic: "DEC C", execFunc: decExecFunc},                           // DEC C
	0x0E: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtC, Mnemonic: "LD C, d8", execFunc: ldExecFunc},                       // LD C, d8
	0x0F: {Type: inRrca, AddressingMode: amImp, Mnemonic: "RRCA", execFunc: rrcaExecFunc},                                            // RRCA
	// 0x1
	0x10: {Type: inStop, Mnemonic: "STOP 0", execFunc: stopExecFunc},                                                                 // STOP 0
	0x11: {Type: inLd, AddressingMode: amRnD16, RegisterType1: rtDE, Mnemonic: "LD DE,d16", execFunc: ldExecFunc},                    // LD DE,d16
	0x12: {Type: inLd, AddressingMode: amMRnR, RegisterType1: rtDE, RegisterType2: rtA, Mnemonic: "LD (DE),A", execFunc: ldExecFunc}, // LD (DE),A
	0x13: {Type: inInc, AddressingMode: amR, RegisterType1: rtDE, Mnemonic: "INC DE", execFunc: incExecFunc},                         // INC DE
	0x14: {Type: inInc, AddressingMode: amR, RegisterType1: rtD, Mnemonic: "INC D", execFunc: incExecFunc},                           // INC D
	0x15: {Type: inDec, AddressingMode: amR, RegisterType1: rtD, Mnemonic: "DEC D", execFunc: decExecFunc},                           // DEC D
	0x16: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtD, Mnemonic: "LD D,d8", execFunc: ldExecFunc},                        // LD D,d8
	0x17: {Type: inRla, AddressingMode: amImp, Mnemonic: "RLA", execFunc: rlaExecFunc},                                               // RLA
	0x18: {Type: inJr, AddressingMode: amD8, Mnemonic: "JR r8", Condition: ctNone, execFunc: jrExecFunc},                             // JR r8
	0x19: {Type: inAdd, AddressingMode: amRnR, RegisterType1: rtHL, RegisterType2: rtDE, Mnemonic: "ADD HL,DE", execFunc: addExecFunc}, // ADD HL,DE
	0x1A: {Type: inLd, AddressingMode: amRnMR, RegisterType1: rtA, RegisterType2: rtDE, Mnemonic: "LD A,(DE)", execFunc: ldExecFunc}, // LD A,(DE)
	0x1B: {Type: inDec, AddressingMode: amR, RegisterType1: rtDE, Mnemonic: "DEC DE", execFunc: decExecFunc},                         // DEC DE
	0x1C: {Type: inInc, AddressingMode: amR, RegisterType1: rtE, Mnemonic: "INC E", execFunc: incExecFunc},                           // INC E
	0x1D: {Type: inDec, AddressingMode: amR, RegisterType1: rtE, Mnemonic: "DEC E", execFunc: decExecFunc},                           // DEC E
	0x1E: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtE, Mnemonic: "LD E,d8", execFunc: ldExecFunc},                        // LD E,d8
	0x1F: {Type: inRra, AddressingMode: amImp, Mnemonic: "RRA", execFunc: rraExecFunc},                                               // RRA
	// 0x2
	0x20: {Type: inJr, AddressingMode: amD8, Mnemonic: "JR NZ,r8", Condition: ctNZ, execFunc: jrExecFunc},                              // JR NZ,r8
	0x21: {Type: inLd, AddressingMode: amRnD16, RegisterType1: rtHL, Mnemonic: "LD HL,d16", execFunc: ldExecFunc},                      // LD HL,d16
	0x22: {Type: inLd, AddressingMode: amHLInR, RegisterType1: rtHL, RegisterType2: rtA, Mnemonic: "LD (HL+),A", execFunc: ldExecFunc}, // LD (HL+),A
	0x23: {Type: inInc, AddressingMode: amR, RegisterType1: rtHL, Mnemonic: "INC HL", execFunc: incExecFunc},                           // INC HL
	0x24: {Type: inInc, AddressingMode: amR, RegisterType1: rtH, Mnemonic: "INC H", execFunc: incExecFunc},                             // INC H
	0x25: {Type: inDec, AddressingMode: amR, RegisterType1: rtH, Mnemonic: "DEC H", execFunc: decExecFunc},                             // DEC H
	0x26: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtH, Mnemonic: "LD H,d8", execFunc: ldExecFunc},                          // LD H,d8
	0x27: {Type: inDaa, Mnemonic: "DAA", execFunc: daaExecFunc},                                                                        // DAA
	0x28: {Type: inJr, AddressingMode: amD8, Mnemonic: "JR Z,r8", Condition: ctZ, execFunc: jrExecFunc},                                // JR Z,r8
	0x29: {Type: inAdd, AddressingMode: amRnR, RegisterType1: rtHL, RegisterType2: rtHL, Mnemonic: "ADD HL,HL", execFunc: addExecFunc}, // ADD HL,HL
	0x2A: {Type: inLd, AddressingMode: amRnHLI, RegisterType1: rtA, RegisterType2: rtHL, Mnemonic: "LD A,(HL+)", execFunc: ldExecFunc}, // LD A,(HL+)
	0x2B: {Type: inDec, AddressingMode: amR, RegisterType1: rtHL, Mnemonic: "DEC HL", execFunc: decExecFunc},                           // DEC HL
	0x2C: {Type: inInc, AddressingMode: amR, RegisterType1: rtL, Mnemonic: "INC L", execFunc: incExecFunc},                             // INC L
	0x2D: {Type: inDec, AddressingMode: amR, RegisterType1: rtL, Mnemonic: "DEC L", execFunc: decExecFunc},                             // DEC L
	0x2E: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtL, Mnemonic: "LD L,d8", execFunc: ldExecFunc},                          // LD L,d8
	0x2F: {Type: inCpl, Mnemonic: "CPL", execFunc: cplExecFunc},                                                                        // CPL
	// 0x3
	0x30: {Type: inJr, AddressingMode: amD8, Mnemonic: "JR NC,r8", Condition: ctNC, execFunc: jrExecFunc},                              // JR NC,r8
	0x31: {Type: inLd, AddressingMode: amRnD16, RegisterType1: rtSP, Mnemonic: "LD SP,d16", execFunc: ldExecFunc},                      // LD SP,d16
	0x32: {Type: inLd, AddressingMode: amHLDnR, RegisterType1: rtHL, RegisterType2: rtA, Mnemonic: "LD (HL-),A", execFunc: ldExecFunc}, // LD (HL-),A
	0x33: {Type: inInc, AddressingMode: amR, RegisterType1: rtSP, Mnemonic: "INC SP", execFunc: incExecFunc},                           // INC SP
	0x34: {Type: inInc, AddressingMode: amMR, RegisterType1: rtHL, Mnemonic: "INC (HL)", execFunc: incExecFunc},                        // INC (HL)
	0x35: {Type: inDec, AddressingMode: amMR, RegisterType1: rtHL, Mnemonic: "DEC (HL)", execFunc: decExecFunc},                        // DEC (HL)
	0x36: {Type: inLd, AddressingMode: amMRnD8, RegisterType1: rtHL, Mnemonic: "LD (HL),d8", execFunc: ldExecFunc},                     // LD (HL),d8
	0x37: {Type: inScf, Mnemonic: "SCF", execFunc: scfExecFunc},                                                                        // SCF
	0x38: {Type: inJr, AddressingMode: amD8, Mnemonic: "JR C,r8", Condition: ctC, execFunc: jrExecFunc},                                // JR C,r8
	0x39: {Type: inAdd, AddressingMode: amRnR, RegisterType1: rtHL, RegisterType2: rtSP, Mnemonic: "ADD HL,SP", execFunc: addExecFunc}, // ADD HL,SP
	0x3A: {Type: inLd, AddressingMode: amRnHLD, RegisterType1: rtA, RegisterType2: rtHL, Mnemonic: "LD A,(HL-)", execFunc: ldExecFunc}, // LD A,(HL-)
	0x3B: {Type: inDec, AddressingMode: amR, RegisterType1: rtSP, Mnemonic: "DEC SP", execFunc: decExecFunc},                           // DEC SP
	0x3C: {Type: inInc, AddressingMode: amR, RegisterType1: rtA, Mnemonic: "INC A", execFunc: incExecFunc},                             // INC A
	0x3D: {Type: inDec, AddressingMode: amR, RegisterType1: rtA, Mnemonic: "DEC A", execFunc: decExecFunc},                             // DEC A
	0x3E: {Type: inLd, AddressingMode: amRnD8, RegisterType1: rtA, Mnemonic: "LD A,d8", execFunc: ldExecFunc},                          // LD A,d8
	0x3F: {Type:inCcf, Mnemonic: "CCF", execFunc: ccfExecFunc},                                                                         // CCF
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
	0x76: {Type: inHalt, Mnemonic: "HALT", execFunc: haltExecFunc},                                                                   // HALT
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
	0x80: {Type: inAdd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtB, Mnemonic: "ADD A,B", execFunc: addExecFunc}, // ADD A,B
	0x81: {Type: inAdd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtC, Mnemonic: "ADD A,C", execFunc: addExecFunc}, // ADD A,C
	0x82: {Type: inAdd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtD, Mnemonic: "ADD A,D", execFunc: addExecFunc}, // ADD A,D
	0x83: {Type: inAdd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtE, Mnemonic: "ADD A,E", execFunc: addExecFunc}, // ADD A,E
	0x84: {Type: inAdd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtH, Mnemonic: "ADD A,H", execFunc: addExecFunc}, // ADD A,H
	0x85: {Type: inAdd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtL, Mnemonic: "ADD A,L", execFunc: addExecFunc}, // ADD A,L
	0x86: {Type: inAdd, AddressingMode: amRnMR, RegisterType1: rtA, RegisterType2: rtHL, Mnemonic: "ADD A,(HL)", execFunc: addExecFunc}, // ADD A,(HL)
	0x87: {Type: inAdd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtA, Mnemonic: "ADD A,A", execFunc: addExecFunc}, // ADD A,A
	0x88: {Type: inAdc, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtB, Mnemonic: "ADC A,B", execFunc: adcExecFunc}, // ADC A,B
	0x89: {Type: inAdc, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtC, Mnemonic: "ADC A,C", execFunc: adcExecFunc}, // ADC A,C
	0x8A: {Type: inAdc, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtD, Mnemonic: "ADC A,D", execFunc: adcExecFunc}, // ADC A,D
	0x8B: {Type: inAdc, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtE, Mnemonic: "ADC A,E", execFunc: adcExecFunc}, // ADC A,E
	0x8C: {Type: inAdc, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtH, Mnemonic: "ADC A,H", execFunc: adcExecFunc}, // ADC A,H
	0x8D: {Type: inAdc, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtL, Mnemonic: "ADC A,L", execFunc: adcExecFunc}, // ADC A,L
	0x8E: {Type: inAdc, AddressingMode: amRnMR, RegisterType1: rtA, RegisterType2: rtHL, Mnemonic: "ADC A,(HL)", execFunc: adcExecFunc}, // ADC A,(HL)
	0x8F: {Type: inAdc, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtA, Mnemonic: "ADC A,A", execFunc: adcExecFunc}, // ADC A,A
	// 0x9
	0x90: {Type: inSub, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtB, Mnemonic: "SUB B", execFunc: subExecFunc}, // SUB B
	0x91: {Type: inSub, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtC, Mnemonic: "SUB C", execFunc: subExecFunc}, // SUB C
	0x92: {Type: inSub, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtD, Mnemonic: "SUB D", execFunc: subExecFunc}, // SUB D
	0x93: {Type: inSub, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtE, Mnemonic: "SUB E", execFunc: subExecFunc}, // SUB E
	0x94: {Type: inSub, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtH, Mnemonic: "SUB H", execFunc: subExecFunc}, // SUB H
	0x95: {Type: inSub, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtL, Mnemonic: "SUB L", execFunc: subExecFunc}, // SUB L
	0x96: {Type: inSub, AddressingMode: amRnMR, RegisterType1: rtA, RegisterType2: rtHL, Mnemonic: "SUB (HL)", execFunc: subExecFunc}, // SUB (HL)
	0x97: {Type: inSub, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtA, Mnemonic: "SUB A", execFunc: subExecFunc}, // SUB A
	0x98: {Type: inSbc, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtB, Mnemonic: "SBC A,B", execFunc: sbcExecFunc}, // ADC A,B
	0x99: {Type: inSbc, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtC, Mnemonic: "SBC A,C", execFunc: sbcExecFunc}, // ADC A,C
	0x9A: {Type: inSbc, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtD, Mnemonic: "SBC A,D", execFunc: sbcExecFunc}, // ADC A,D
	0x9B: {Type: inSbc, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtE, Mnemonic: "SBC A,E", execFunc: sbcExecFunc}, // ADC A,E
	0x9C: {Type: inSbc, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtH, Mnemonic: "SBC A,H", execFunc: sbcExecFunc}, // ADC A,H
	0x9D: {Type: inSbc, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtL, Mnemonic: "SBC A,L", execFunc: sbcExecFunc}, // ADC A,L
	0x9E: {Type: inSbc, AddressingMode: amRnMR, RegisterType1: rtA, RegisterType2: rtHL, Mnemonic: "SBC A,(HL)", execFunc: sbcExecFunc}, // ADC A,(HL)
	0x9F: {Type: inSbc, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtA, Mnemonic: "SBC A,A", execFunc: sbcExecFunc}, // ADC A,A
	// 0xA
	0xA0: {Type: inAnd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtB, Mnemonic: "AND B", execFunc: andExecFunc}, // AND B
	0xA1: {Type: inAnd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtC, Mnemonic: "AND C", execFunc: andExecFunc}, // AND C
	0xA2: {Type: inAnd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtD, Mnemonic: "AND D", execFunc: andExecFunc}, // AND D
	0xA3: {Type: inAnd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtE, Mnemonic: "AND E", execFunc: andExecFunc}, // AND E
	0xA4: {Type: inAnd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtH, Mnemonic: "AND H", execFunc: andExecFunc}, // AND H
	0xA5: {Type: inAnd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtL, Mnemonic: "AND L", execFunc: andExecFunc}, // AND L
	0xA6: {Type: inAnd, AddressingMode: amRnMR, RegisterType1: rtA, RegisterType2: rtHL, Mnemonic: "AND (HL)", execFunc: andExecFunc}, // AND (HL)
	0xA7: {Type: inAnd, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtA, Mnemonic: "AND A", execFunc: andExecFunc}, // AND A
	0xA8: {Type: inXor, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtB, Mnemonic: "XOR B", execFunc: xorExecFunc}, // XOR B
	0xA9: {Type: inXor, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtC, Mnemonic: "XOR C", execFunc: xorExecFunc}, // XOR C
	0xAA: {Type: inXor, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtD, Mnemonic: "XOR D", execFunc: xorExecFunc}, // XOR D
	0xAB: {Type: inXor, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtE, Mnemonic: "XOR E", execFunc: xorExecFunc}, // XOR E
	0xAC: {Type: inXor, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtH, Mnemonic: "XOR H", execFunc: xorExecFunc}, // XOR H
	0xAD: {Type: inXor, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtL, Mnemonic: "XOR L", execFunc: xorExecFunc}, // XOR L
	0xAE: {Type: inXor, AddressingMode: amRnMR, RegisterType1: rtA, RegisterType2: rtHL, Mnemonic: "XOR (HL)", execFunc: xorExecFunc}, // XOR (HL)
	0xAF: {Type: inXor, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtA, Mnemonic: "XOR A", execFunc: xorExecFunc}, // XOR A
	// 0xB
	0xB0: {Type: inOr, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtB, Mnemonic: "OR B", execFunc: orExecFunc}, // OR B
	0xB1: {Type: inOr, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtC, Mnemonic: "OR C", execFunc: orExecFunc}, // OR C
	0xB2: {Type: inOr, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtD, Mnemonic: "OR D", execFunc: orExecFunc}, // OR D
	0xB3: {Type: inOr, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtE, Mnemonic: "OR E", execFunc: orExecFunc}, // OR E
	0xB4: {Type: inOr, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtH, Mnemonic: "OR H", execFunc: orExecFunc}, // OR H
	0xB5: {Type: inOr, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtL, Mnemonic: "OR L", execFunc: orExecFunc}, // OR L
	0xB6: {Type: inOr, AddressingMode: amRnMR, RegisterType1: rtA, RegisterType2: rtHL, Mnemonic: "OR (HL)", execFunc: orExecFunc}, // OR (HL)
	0xB7: {Type: inOr, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtA, Mnemonic: "OR A", execFunc: orExecFunc}, // OR A
	0xB8: {Type: inCp, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtB, Mnemonic: "CP B", execFunc: cpExecFunc}, // CP B
	0xB9: {Type: inCp, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtC, Mnemonic: "CP C", execFunc: cpExecFunc}, // CP C
	0xBA: {Type: inCp, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtD, Mnemonic: "CP D", execFunc: cpExecFunc}, // CP D
	0xBB: {Type: inCp, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtE, Mnemonic: "CP E", execFunc: cpExecFunc}, // CP E
	0xBC: {Type: inCp, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtH, Mnemonic: "CP H", execFunc: cpExecFunc}, // CP H
	0xBD: {Type: inCp, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtL, Mnemonic: "CP L", execFunc: cpExecFunc}, // CP L
	0xBE: {Type: inCp, AddressingMode: amRnMR, RegisterType1: rtA, RegisterType2: rtHL, Mnemonic: "CP (HL)", execFunc: cpExecFunc}, // CP (HL)
	0xBF: {Type: inCp, AddressingMode: amRnR, RegisterType1: rtA, RegisterType2: rtA, Mnemonic: "CP A", execFunc: cpExecFunc}, // CP A
	// 0xC
	0xC0: {Type: inRet, Mnemonic: "RET NZ", Condition: ctNZ, execFunc: retExecFunc},                               // RET NZ
	0xC1: {Type: inPop, AddressingMode: amImp, RegisterType1: rtBC, Mnemonic: "POP BC", execFunc: popExecFunc},    // POP BC
	0xC2: {Type: inJp, AddressingMode: amD16, Mnemonic: "JP NZ,a16", Condition: ctNZ, execFunc: jpExecFunc},       // JP NZ,a16
	0xC3: {Type: inJp, AddressingMode: amD16, Mnemonic: "JP a16", execFunc: jpExecFunc},                           // JP a16
	0xC4: {Type: inCall, AddressingMode: amD16, Mnemonic: "CALL NZ,a16", Condition: ctNZ, execFunc: callExecFunc}, // CALL NZ,a16
	0xC5: {Type: inPush, AddressingMode: amImp, RegisterType1: rtBC, Mnemonic: "PUSH BC", execFunc: pushExecFunc}, // PUSH BC
	0xC6: {Type: inAdd, AddressingMode: amRnD8, RegisterType1: rtA, Mnemonic: "ADD A,d8", execFunc: addExecFunc}, // ADD A,d8
	0xC7: {Type: inRst, AddressingMode: amImp, Mnemonic: "RST 00H", Parameter: 0x0, execFunc: rstExecFunc},        // RST 00H
	0xC8: {Type: inRet, Mnemonic: "RET Z", Condition: ctZ, execFunc: retExecFunc},                                 // RET Z
	0xC9: {Type: inRet, Mnemonic: "RET", Condition: ctNone, execFunc: retExecFunc},                                // RET
	0xCA: {Type: inJp, AddressingMode: amD16, Mnemonic: "JP Z,a16", Condition: ctZ, execFunc: jpExecFunc},         // JP Z,a16
	0xCB: {Type: inCb, AddressingMode: amD8, Mnemonic: "PREFIX CB", execFunc: cbExecFunc},                         // PREFIX CB
	0xCC: {Type: inCall, AddressingMode: amD16, Mnemonic: "CALL Z,a16", Condition: ctZ, execFunc: callExecFunc},   // CALL Z,a16
	0xCD: {Type: inCall, AddressingMode: amD16, Mnemonic: "CALL a16", Condition: ctNone, execFunc: callExecFunc},  // CALL a16
	0xCE: {Type: inAdc, AddressingMode: amRnD8, RegisterType1: rtA, Mnemonic: "ADC A,d8", execFunc: adcExecFunc}, // ADC A,d8
	0xCF: {Type: inRst, AddressingMode: amImp, Mnemonic: "RST 08H", Parameter: 0x08, execFunc: rstExecFunc},       // RST 08H
	// 0xD
	0xD0: {Type: inRet, Mnemonic: "RET NC", Condition: ctNC, execFunc: retExecFunc},                               // RET NC
	0xD1: {Type: inPop, AddressingMode: amImp, RegisterType1: rtDE, Mnemonic: "POP DE", execFunc: popExecFunc},    // POP DE
	0xD2: {Type: inJp, AddressingMode: amD16, Mnemonic: "JP NC,a16", Condition: ctNC, execFunc: jpExecFunc},       // JP NC,a16
	0xD4: {Type: inCall, AddressingMode: amD16, Mnemonic: "CALL NC,a16", Condition: ctNC, execFunc: callExecFunc}, // CALL NC,a16
	0xD5: {Type: inPush, AddressingMode: amImp, RegisterType1: rtDE, Mnemonic: "PUSH DE", execFunc: pushExecFunc}, // PUSH DE
	0xD6: {Type: inSub, AddressingMode: amD8, Mnemonic: "SUB d8", execFunc: subExecFunc},                          // SUB d8
	0xD7: {Type: inRst, AddressingMode: amImp, Mnemonic: "RST 10H", Parameter: 0x10, execFunc: rstExecFunc},       // RST 10H
	0xD8: {Type: inRet, Mnemonic: "RET C", Condition: ctC, execFunc: retExecFunc},                                 // RET C
	0xD9: {Type: inReti, Mnemonic: "RETI", execFunc: retiExecFunc},                                                // RETI
	0xDA: {Type: inJp, AddressingMode: amD16, Mnemonic: "JP C,a16", Condition: ctC, execFunc: jpExecFunc},         // JP C,a16
	0xDC: {Type: inCall, AddressingMode: amD16, Mnemonic: "CALL C,a16", Condition: ctC, execFunc: callExecFunc},   // CALL C,a16
	0xDE: {Type: inSbc, AddressingMode: amRnD8, RegisterType1: rtA, Mnemonic: "SBC A,d8", execFunc: sbcExecFunc}, // SBC A,d8
	0xDF: {Type: inRst, AddressingMode: amImp, Mnemonic: "RST 18H", Parameter: 0x18, execFunc: rstExecFunc},       // RST 18H
	// 0xE
	0xE0: {Type: inLdh, AddressingMode: amA8nR, RegisterType2: rtA, Mnemonic: "LDH (a8),A", execFunc: ldhExecFunc},                 // LDH (a8),A
	0xE1: {Type: inPop, AddressingMode: amImp, RegisterType1: rtHL, Mnemonic: "POP HL", execFunc: popExecFunc},                     // POP HL
	0xE2: {Type: inLd, AddressingMode: amMRnR, RegisterType1: rtC, RegisterType2: rtA, Mnemonic: "LD (C),A", execFunc: ldExecFunc}, // LD (C),A
	0xE5: {Type: inPush, AddressingMode: amImp, RegisterType1: rtHL, Mnemonic: "PUSH HL", execFunc: pushExecFunc},                  // PUSH HL
	0xE6: {Type: inAnd, AddressingMode: amRnD8, RegisterType1: rtA, Mnemonic: "AND d8", execFunc: andExecFunc},                     // AND d8
	0xE7: {Type: inRst, AddressingMode: amImp, Mnemonic: "RST 20H", Parameter: 0x20, execFunc: rstExecFunc},                        // RST 20H
	0xE8: {Type: inAdd, AddressingMode: amRnD8, RegisterType1: rtSP, Mnemonic: "ADD SP,r8", execFunc: addExecFunc},                 // ADD SP,r8
	0xE9: {Type: inJp, AddressingMode: amMR, RegisterType1: rtHL, Mnemonic: "JP (HL)", Condition: ctNone, execFunc: jpExecFunc},    // JP (HL) CHECK VIDEO 9 DUE TO POSSIBLE BUG
	0xEA: {Type: inLd, AddressingMode: amA16nR, RegisterType2: rtA, Mnemonic: "LD (a16),A", execFunc: ldExecFunc},                  // LD (a16),A
	0xEE: {Type: inXor, AddressingMode: amRnD8, RegisterType1: rtA, Mnemonic: "XOR d8", execFunc: xorExecFunc},                     // XOR d8
	0xEF: {Type: inRst, AddressingMode: amImp, Mnemonic: "RST 28H", Parameter: 0x28, execFunc: rstExecFunc},                        // RST 28H
	// 0xF
	0xF0: {Type: inLdh, AddressingMode: amRnA8, RegisterType1: rtA, Mnemonic: "LDH A,(a8)", execFunc: ldhExecFunc},                 // LDH A,(a8)
	0xF1: {Type: inPop, AddressingMode: amImp, RegisterType1: rtAF, Mnemonic: "POP AF", execFunc: popExecFunc},                     // POP AF
	0xF2: {Type: inLd, AddressingMode: amRnMR, RegisterType1: rtA, RegisterType2: rtC, Mnemonic: "LD A,(C)", execFunc: ldExecFunc}, // LD A,(C)
	0xF3: {Type: inDi, Mnemonic: "DI", execFunc: diExecFunc},                                                                       // DI
	0xF5: {Type: inPush, AddressingMode: amImp, RegisterType1: rtAF, Mnemonic: "PUSH AF", execFunc: pushExecFunc},                  // PUSH AF
	0xF6: {Type: inOr, AddressingMode: amRnD8, RegisterType1: rtA, Mnemonic: "OR d8", execFunc: orExecFunc},                        // OR d8
	0xF7: {Type: inRst, AddressingMode: amImp, Mnemonic: "RST 30H", Parameter: 0x30, execFunc: rstExecFunc},                        // RST 30H
	0xF8: {Type: inLd, AddressingMode: amHLnSPR, RegisterType1: rtHL, RegisterType2: rtSP, Mnemonic: "LD HL,SP+r8", execFunc: ldExecFunc}, // LD HL,SP+r8
	0xF9: {Type: inLd, AddressingMode: amRnR, RegisterType1: rtSP, RegisterType2: rtHL, Mnemonic: "LD SP,HL", execFunc: ldExecFunc}, // LD SP,HL
	0xFA: {Type: inLd, AddressingMode: amRnA16, RegisterType1: rtA, Mnemonic: "LD A,(a16)", execFunc: ldExecFunc},                   // LD A,(a16)
	0xFB: {Type: inEi, Mnemonic: "EI", execFunc: eiExecFunc},                                                                        // EI
	0xFE: {Type: inCp, AddressingMode: amRnD8, RegisterType1: rtA, Mnemonic: "CP d8", execFunc: cpExecFunc},                         // CP d8
	0xFF: {Type: inRst, AddressingMode: amImp, Mnemonic: "RST 38H", Parameter: 0x38, execFunc: rstExecFunc},                         // RST 38H
}
