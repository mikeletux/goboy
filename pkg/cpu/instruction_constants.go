package cpu

// Constant about instruction types
const (
	inNone = 0
	inNop  = 1
	inLd   = 2
	inInc  = 3
	inDec  = 4
	inRlca = 5
	inAdd  = 6
	inRrca = 7
	inStop = 8
	inRla  = 9
	inJr   = 10
	inRra  = 11
	inDaa  = 12
	inCpl  = 13
	inScf  = 14
	inCcf  = 15
	inHalt = 16
	inAdc  = 17
	inSub  = 18
	inSbc  = 19
	inAnd  = 20
	inXor  = 21
	inOr   = 22
	inCp   = 23
	inPop  = 24
	inJp   = 25
	inPush = 26
	inRet  = 27
	inCb   = 28
	inCall = 29
	inReti = 30
	inLdh  = 31
	inJphl = 32
	inDi   = 33
	inEi   = 34
	inRst  = 35
	inErr  = 36

	// CB instructions
	inRlc  = 37
	inRrc  = 38
	inRl   = 39
	inRr   = 40
	inSla  = 41
	inSra  = 42
	inSwap = 43
	inSrl  = 44
	inBit  = 45
	inRes  = 46
	inSet  = 47
)

// Constants about addressing mode (/!\using n as separator for them) ...TBD docs!
const (
	amRnD16  = 0
	amRnR    = 1
	amMRnR   = 2
	amR      = 3
	amRnD8   = 4
	amRnMR   = 5
	amRnHLI  = 6
	amRnHLD  = 7
	amHLInR  = 8
	amHLRnR  = 9
	amRnA8   = 10
	amA8nR   = 11
	amHLnSPR = 12
	amD16    = 13
	amD8     = 14
	amImp    = 15 // Implied means there's nothing else to read after the instruction
	amD16nR  = 16
	amMRnD8  = 17
	amMR     = 18
	amA16nR  = 19
	amRnA16  = 20
)

// Constants about register types
const (
	rtNone = 0  // No register
	rtA    = 1  // Register A
	rtF    = 2  // Register F
	rtB    = 3  // Register B
	rtC    = 4  // Register C
	rtD    = 5  // Register D
	rtE    = 6  // Register E
	rtH    = 7  // Register H
	rtL    = 8  // Register L
	rtAF   = 9  // Register AF
	rtBC   = 10 // Register BC
	rtDE   = 11 // Register DE
	rtHL   = 12 // Register HL
	rtSP   = 13 // Stack Pointer
	rtPC   = 14 // Program Counter
)

// Constants about conditions
const (
	ctNone = 0 // No condition
	ctNZ   = 1 // No Z flag set
	ctZ    = 2 // Z flag set
	ctNC   = 3 // No C flag set
	ctC    = 4 // C flag set
)
