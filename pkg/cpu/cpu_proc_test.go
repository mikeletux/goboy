package cpu

import (
	"github.com/mikeletux/goboy/pkg/bus"
	"github.com/mikeletux/goboy/pkg/log"
	"testing"
)

const(
	// Constants regarding TestJpExecFunc
	jpAddress uint16 = 0xFEA // Why this address? Because it means ugly is spanish :)

	// Constants regarding TestCallExecFunc
	callDummyAddress         uint16 = 0x2AE
	callInitPCAddress            uint16 = 0x150
	callStackPointerInitPosition uint16 = 0xDFFF

	rstInitPCAddress uint16 = 0x150
)

// TestJpExecFunc tests JP
func TestJpExecFunc(t *testing.T) {
	cpu := Init(nil, &log.NilLogger{}) // Bus can be Nil since it's not being used for JP

	tests := []struct{
		testName string
		addressToJump uint16 // This is the mock address to jump if condition is satisfied
		condition int // This is the mock condition from JP instruction
		Z,C bool // These are the CPU registers
		shouldJump bool // This indicates if the JP instruction should jump or not
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
func TestCallExecFunc(t *testing.T) {
	cpu := Init(bus.NewMapMock(), &log.NilLogger{})

	tests := []struct{
		testName string
		addressToCall uint16 // This is the mock address to call if condition is satisfied
		condition int // This is the mock condition from CALL instruction
		Z,C bool // These are the CPU registers
		shouldJump bool // This indicates if the CALL instruction should execute or not
	}{
		{
			testName: "1 - Test CALL NZ,a16 with Z=0",
			addressToCall: callDummyAddress,
			condition: ctNZ,
			shouldJump: true,
		},
		{
			testName: "2 - Test CALL NZ,a16 with Z=1",
			addressToCall: callDummyAddress,
			condition: ctNZ,
			Z: true,
			shouldJump: false,
		},
		{
			testName: "3 - Test CALL NC,a16 with C=0",
			addressToCall: callDummyAddress,
			condition: ctNC,
			shouldJump: true,
		},
		{
			testName: "4 - Test CALL NC,a16 with C=1",
			addressToCall: callDummyAddress,
			condition: ctNC,
			C: true,
			shouldJump: false,
		},
		{
			testName: "5 - Test CALL Z,a16 with Z=0",
			addressToCall: callDummyAddress,
			condition: ctZ,
			shouldJump: false,
		},
		{
			testName: "6 - Test CALL Z,a16 with Z=1",
			addressToCall: callDummyAddress,
			condition: ctZ,
			Z: true,
			shouldJump: true,
		},
		{
			testName: "7 - Test CALL C,a16 with C=0",
			addressToCall: callDummyAddress,
			condition: ctC,
			shouldJump: false,
		},
		{
			testName: "8 - Test CALL C,a16 with C=1",
			addressToCall: callDummyAddress,
			condition: ctC,
			C: true,
			shouldJump: true,
		},
		{
			testName: "9 - Test CALL a16 without condition",
			addressToCall: callDummyAddress,
			condition: ctNone,
			shouldJump: true,
		},
	}

	for _, test := range tests {
		cpu.CurrentInstruction = Instruction{Condition: test.condition}
		cpu.FetchedData = test.addressToCall
		cpu.registers.PC = callInitPCAddress
		cpu.registers.SP = callStackPointerInitPosition
		cpu.registers.SetFZ(test.Z)
		cpu.registers.SetFC(test.C)

		callExecFunc(cpu)

		if test.shouldJump  {
			if cpu.registers.PC != callDummyAddress { // Check that PC has jump to where it should
				t.Errorf("[%s] The program counter should have jumped to %X and it is %X",
					test.testName, callDummyAddress, cpu.registers.PC)
			}

			if cpu.registers.SP != callStackPointerInitPosition- 2 { // Check that SP has decreased by 2 bytes
				t.Errorf("[%s] SP should be %d and it is %d",
					test.testName, callStackPointerInitPosition- 2, cpu.registers.SP)
			}

			low := cpu.bus.BusRead(cpu.registers.SP)
			high := cpu.bus.BusRead(cpu.registers.SP + 1)

			if callInitPCAddress != uint16(high)<<8 | uint16(low) { // Check that written PC in the stack is the one prior CALL
				t.Errorf("[%s] The PC addr recovered from stack %X does not match with the init one %X",
					test.testName, uint16(high)<<8 | uint16(low), callInitPCAddress)
			}
		}

		if !test.shouldJump {
			if cpu.registers.PC != callInitPCAddress { // Check that PC has not moved
				t.Errorf("[%s] The program counter shouldn't have from %X and it is %X",
					test.testName, callInitPCAddress, cpu.registers.PC)
			}

			if cpu.registers.SP != callStackPointerInitPosition { // Check that SP is the same prior the call
				t.Errorf("[%s] SP shouldn't have moved from %X and it is %X",
					test.testName, callStackPointerInitPosition, cpu.registers.SP)
			}

		}
	}

}

// TestRstExecFunc tests RST
func TestRstExecFunc(t *testing.T) {
	cpu := Init(bus.NewMapMock(), &log.NilLogger{})

	tests := []struct{
		testName string
		addressToJump byte

	}{
		{
			testName: "1 - Test RST 00H",
			addressToJump: 0x00,

		},
		{
			testName: "2 - Test RST 08H",
			addressToJump: 0x08,

		},
		{
			testName: "3 - Test RST 10H",
			addressToJump: 0x10,

		},
		{
			testName: "4 - Test RST 18H",
			addressToJump: 0x18,

		},
		{
			testName: "5 - Test RST 20H",
			addressToJump: 0x20,

		},
		{
			testName: "6 - Test RST 28H",
			addressToJump: 0x28,

		},
		{
			testName: "7 - RST 30H",
			addressToJump: 0x30,

		},
		{
			testName: "8 - Test RST 38H",
			addressToJump: 0x38,

		},
	}

	for _, test := range tests {
		cpu.CurrentInstruction = Instruction{Parameter: test.addressToJump}
		cpu.registers.PC = rstInitPCAddress

		rstExecFunc(cpu)

		if byte(cpu.registers.PC & 0xFF) != test.addressToJump {
			t.Errorf("[%s] PC should have jumped to %X and it is %X",
				test.testName, test.addressToJump, cpu.registers.PC)
		}

		low := cpu.bus.BusRead(cpu.registers.SP)
		high := cpu.bus.BusRead(cpu.registers.SP + 1)

		if rstInitPCAddress != uint16(high)<<8 | uint16(low) { // Check that written PC in the stack is the one prior CALL
			t.Errorf("[%s] PC address recovered from stack %X does not match with the init one %X",
				test.testName, uint16(high)<<8 | uint16(low), rstInitPCAddress)
		}

	}
}

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




