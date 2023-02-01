package cpu

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
	return uint16(r.A<<8 + r.F)
}

// GetBC returns a 16Bit value from CPU registers B and C
func (r *Registers) GetBC() uint16 {
	return uint16(r.B<<8 + r.C)
}

// GetDE returns a 16Bit value from CPU registers D and E
func (r *Registers) GetDE() uint16 {
	return uint16(r.D<<8 + r.E)
}

// GetHL returns a 16Bit value from CPU registers H and L
func (r *Registers) GetHL() uint16 {
	return uint16(r.H<<8 + r.L)
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

func getHighAndLowBytes(value uint16) (high, low byte) {
	high, low = byte(value>>8), byte(value&0xFF)
	return
}

type CPU struct {
	Registers *Registers
}
