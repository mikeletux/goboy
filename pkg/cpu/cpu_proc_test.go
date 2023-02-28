package cpu

import (
	"github.com/mikeletux/goboy/pkg/log"
	"testing"
)

const(
	jpAddress = 0xFEA // Why this address? Because it means ugly is spanish :)
)

// TestJpExecFunc tests JP
func TestJpExecFunc(t *testing.T) {
	cpu := Init(&BusMapMock{}, &log.NilLogger{})
	tests := []struct{
		testName string
		addressToJump uint16 // This is the mock address to jump if condition is satisfied
		condition int // This is the mock condition from JP instruction
		Z,C bool // These are the CPU registers
		shouldJump bool // This indicates if the JP instruction should execute or not
	}{
		{
			testName: "1 - Test JP a16 and JP (HL)",
			addressToJump: jpAddress,
			condition: ctNone,
			shouldJump: true,
		},
		{
			testName: "2 - Test JP NZ,a16 with Z=0",
			addressToJump: jpAddress,
			condition: ctNZ,
			shouldJump: true,
		},
		{
			testName: "3 - Test JP NZ,a16 with Z=1",
			addressToJump: jpAddress,
			condition: ctNZ,
			Z: true,
			shouldJump: false,
		},
		{
			testName: "4 - Test JP NC,a16 with C=0",
			addressToJump: jpAddress,
			condition: ctNC,
			shouldJump: true,
		},
		{
			testName: "5 - Test JP NC,a16 with C=1",
			addressToJump: jpAddress,
			condition: ctNC,
			C: true,
			shouldJump: false,
		},
		{
			testName: "6 - Test JP Z,a16 with Z=0",
			addressToJump: jpAddress,
			condition: ctZ,
			shouldJump: false,
		},
		{
			testName: "7 - Test JP Z,a16 with Z=1",
			addressToJump: jpAddress,
			condition: ctZ,
			Z: true,
			shouldJump: true,
		},
		{
			testName: "8 - Test JP C,a16 with C=0",
			addressToJump: jpAddress,
			condition: ctC,
			shouldJump: false,
		},
		{
			testName: "9 - Test JP C,a16 with C=1",
			addressToJump: jpAddress,
			condition: ctC,
			C: true,
			shouldJump: true,
		},
	}

	for _, test := range tests {
		cpu.CurrentInstruction = Instruction{Condition: test.condition}
		cpu.FetchedData = test.addressToJump
		cpu.registers.PC = 0x0
		cpu.registers.SetFZ(test.Z)
		cpu.registers.SetFC(test.C)

		jpExecFunc(cpu)

		if test.shouldJump && cpu.registers.PC != jpAddress {
			t.Errorf("[%s] The program counter should have jumped to %X and it is %X",
				test.testName, jpAddress, cpu.registers.PC)
		}

		if !test.shouldJump && cpu.registers.PC == jpAddress {
			t.Errorf("[%s] The program counter shouldn't have jumped to %X and it is %X",
				test.testName, jpAddress, cpu.registers.PC)
		}
	}
}

// TestCallExecFunc tests CALL
func TestCallExecFunc(t *testing.T) {}

// TestRstExecFunc tests RST
func TestRstExecFunc(t *testing.T) {}

// TestRetExecFunc tests RET
func TestRetExecFunc(t *testing.T) {}

// TestRetiExecFunc tests RETI
func TestRetiExecFunc(t *testing.T) {}

// TestJrExecFunc tests JR
func TestJrExecFunc(t *testing.T) {}

// TestPopExecFunc tests POP
func TestPopExecFunc(t *testing.T) {}

// TestPushExecFunc tests PUSH
func TestPushExecFunc(t *testing.T) {}

// TestDiExecFunc tests DI
func TestDiExecFunc(t *testing.T) {}

// TestLdExecFunc tests LD
func TestLdExecFunc(t *testing.T) {}

// TestLdhExecFunc tests LDH
func TestLdhExecFunc(t *testing.T) {}

// TestIncExecFunc tests INC
func TestIncExecFunc(t *testing.T) {}

// TestDecExecFunc tests DEC
func TestDecExecFunc(t *testing.T) {}

// TestAddExecFunc tests ADD
func TestAddExecFunc(t *testing.T) {}

// TestAdcExecFunc tests ADC
func TestAdcExecFunc(t *testing.T) {}

// TestSubExecFunc tests SUB
func TestSubExecFunc(t *testing.T) {}

// TestSbcExecFunc tests SBC
func TestSbcExecFunc(t *testing.T) {}

// TestAndExecFunc tests AND
func TestAndExecFunc(t *testing.T) {}

// TestXorExecFunc tests XOR
func TestXorExecFunc(t *testing.T) {}

// TestOrExecFunc tests OR
func TestOrExecFunc(t *testing.T) {}

// TestCpExecFunc tests CP
func TestCpExecFunc(t *testing.T) {}




