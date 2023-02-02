package cpu

// Constant about instruction types
const (
	inNop  = 0
	inLd   = 1
	inInc  = 2
	inDec  = 3
	inRlca = 4
	inAdd  = 5
	inRrca = 6
	inStop = 7
	inRla  = 8
	inJr   = 9
	inRra  = 10
	inDaa  = 11
	inCpl  = 12
	inScf  = 13
	inCcf  = 14
	inHalt = 15
	inAdc  = 16
	inSub  = 17
	inSbc  = 18
	inAnd  = 19
	inXor  = 20
	inOr   = 21
	inCp   = 22
	inPop  = 23
	inJp   = 24
	inPush = 25
	inRet  = 26
	inCb   = 27
	inCall = 28
	inReti = 29
	inLdh  = 30
	inJphl = 31
	inDi   = 32
	inEi   = 33
	inRst  = 34
	inErr  = 35

	// CB instructions
	inRlc  = 36
	inRrc  = 37
	inRl   = 38
	inRr   = 39
	inSla  = 40
	inSra  = 41
	inSwap = 42
	inSrl  = 43
	inBit  = 44
	inRes  = 45
	inSet  = 46
)

// Constants about addressing mode (/!\using n as separator for them) ...TBD docs!
const (
	amRnD16  = 0 // From 16bit address to register
	amRnR    = 1 // From register to register
	amMRnR   = 2 // Register to memory location of register
	amR      = 3 // single register
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
	amImp    = 15
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
