package cpu

// Constant about instruction types
const (
	inNone = iota
	inNop
	inLd
	inInc
	inDec
	inRlca
	inAdd
	inRrca
	inStop
	inRla
	inJr
	inRra
	inDaa
	inCpl
	inScf
	inCcf
	inHalt
	inAdc
	inSub
	inSbc
	inAnd
	inXor
	inOr
	inCp
	inPop
	inJp
	inPush
	inRet
	inCb
	inCall
	inReti
	inLdh
	inJphl
	inDi
	inEi
	inRst
	inErr

	// CB instructions
	inRlc
	inRrc
	inRl
	inRr
	inSla
	inSra
	inSwap
	inSrl
	inBit
	inRes
	inSet
)

// Constants about addressing mode (/!\using n as separator for them) ...TBD docs!
const (
	amImp = iota // Implied means there's nothing else to read after the instruction
	amRnD16
	amRnR
	amMRnR
	amR
	amRnD8
	amRnMR
	amRnHLI
	amRnHLD
	amHLInR
	amHLRnR
	amRnA8
	amA8nR
	amHLnSPR
	amD16
	amD8
	amD16nR
	amMRnD8
	amMR
	amA16nR
	amRnA16
)

// Constants about register types
const (
	rtNone = iota // No register
	rtA           // Register A
	rtF           // Register F
	rtB           // Register B
	rtC           // Register C
	rtD           // Register D
	rtE           // Register E
	rtH           // Register H
	rtL           // Register L
	rtAF          // Register AF
	rtBC          // Register BC
	rtDE          // Register DE
	rtHL          // Register HL
	rtSP          // Stack Pointer
	rtPC          // Program Counter
)

// Constants about conditions
const (
	ctNone = iota // No condition
	ctNZ          // No Z flag set
	ctZ           // Z flag set
	ctNC          // No C flag set
	ctC           // C flag set
)
