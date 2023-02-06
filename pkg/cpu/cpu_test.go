package cpu

import (
	"github.com/mikeletux/goboy/pkg/log"
	"reflect"
	"testing"
)

const (
	ATestData byte = 0xDE
	FTestData byte = 0xF0

	BTestData byte = 0x10
	CTestData byte = 0xE4

	DTestData byte = 0xA9
	ETestData byte = 0x70

	HTestData byte = 0x15
	LTestData byte = 0x5B

	SPTestData uint16 = 0x05A0
	PCTestData uint16 = 0xF50D

	AFTestData uint16 = 0xDEF0
	BCTestData uint16 = 0x10E4
	DETestData uint16 = 0xA970
	HLTestData uint16 = 0x155B
)

func TestGetSetRegisters(t *testing.T) {
	// Init CPU
	cpu := Init(&BusMock{}, log.NewBuiltinStdoutLogger())

	// Set registers with some data for testing Getters
	initCPURegistersWithTestData(cpu)

	// Test GetAF()
	got := cpu.registers.GetAF()
	if got != AFTestData {
		t.Errorf("got: %d, expected %d", got, AFTestData)
	}

	// Test GetBC()
	got = cpu.registers.GetBC()
	if got != BCTestData {
		t.Errorf("got: %d, expected %d", got, BCTestData)
	}

	// Test GetDE()
	got = cpu.registers.GetDE()
	if got != DETestData {
		t.Errorf("got: %d, expected %d", got, DETestData)
	}

	// Test GetHL()
	got = cpu.registers.GetHL()
	if got != HLTestData {
		t.Errorf("got: %d, expected %d", got, HLTestData)
	}

	// Set all registers to O
	cpu.registers.A = 0
	cpu.registers.F = 0
	cpu.registers.B = 0
	cpu.registers.C = 0
	cpu.registers.D = 0
	cpu.registers.E = 0
	cpu.registers.H = 0
	cpu.registers.L = 0

	// Test SetAF
	cpu.registers.SetAF(AFTestData)
	if cpu.registers.A != ATestData || cpu.registers.F != FTestData {
		t.Errorf("got: A:%d F:%d, expected A:%d F:%d", cpu.registers.A, cpu.registers.F, ATestData, FTestData)
	}

	// Test SetBC
	cpu.registers.SetBC(BCTestData)
	if cpu.registers.B != BTestData || cpu.registers.C != CTestData {
		t.Errorf("got: B:%d C:%d, expected B:%d C:%d", cpu.registers.B, cpu.registers.C, BTestData, CTestData)
	}

	// Test SetDE
	cpu.registers.SetDE(DETestData)
	if cpu.registers.D != DTestData || cpu.registers.E != ETestData {
		t.Errorf("got: D:%d E:%d, expected D:%d E:%d", cpu.registers.D, cpu.registers.E, DTestData, ETestData)
	}

	// Test SetHL
	cpu.registers.SetHL(HLTestData)
	if cpu.registers.H != HTestData || cpu.registers.L != LTestData {
		t.Errorf("got: H:%d L:%d, expected H:%d L:%d", cpu.registers.H, cpu.registers.L, HTestData, LTestData)
	}
}

func TestFetchDataFromRegisters(t *testing.T) {
	// Init CPU
	cpu := Init(&BusMock{}, log.NewBuiltinStdoutLogger())

	// Set registers with some data for testing Getters
	initCPURegistersWithTestData(cpu)

	testCases := []struct {
		registerType  int
		expectedValue uint16
		hasError      bool
	}{
		{
			registerType:  rtA,
			expectedValue: uint16(ATestData),
			hasError:      false,
		},
		{
			registerType:  rtF,
			expectedValue: uint16(FTestData),
			hasError:      false,
		},
		{
			registerType:  rtB,
			expectedValue: uint16(BTestData),
			hasError:      false,
		},
		{
			registerType:  rtC,
			expectedValue: uint16(CTestData),
			hasError:      false,
		},
		{
			registerType:  rtD,
			expectedValue: uint16(DTestData),
			hasError:      false,
		},
		{
			registerType:  rtE,
			expectedValue: uint16(ETestData),
			hasError:      false,
		},
		{
			registerType:  rtH,
			expectedValue: uint16(HTestData),
			hasError:      false,
		},
		{
			registerType:  rtL,
			expectedValue: uint16(LTestData),
			hasError:      false,
		},
		{
			registerType: 999, // This register doesn't exist
			hasError:     true,
		},
	}

	for _, testCase := range testCases {
		got, err := cpu.registers.FetchDataFromRegisters(testCase.registerType)
		if err == nil {
			if testCase.hasError {
				t.Errorf("the test should've returned an error but it didn't")
			}
		} else {
			if !testCase.hasError {
				t.Errorf("the test should've not returned an error but it did")
			}
		}

		if got != testCase.expectedValue {
			t.Errorf("got %d expected %d for register type %d", got, testCase.expectedValue, testCase.registerType)
		}
	}

}

func initCPURegistersWithTestData(cpu *CPU) {
	cpu.registers.A = ATestData
	cpu.registers.F = FTestData
	cpu.registers.B = BTestData
	cpu.registers.C = CTestData
	cpu.registers.D = DTestData
	cpu.registers.E = ETestData
	cpu.registers.H = HTestData
	cpu.registers.L = LTestData
}

func TestGetPCAndIncrement(t *testing.T) {
	// Init CPU
	cpu := Init(&BusMock{}, log.NewBuiltinStdoutLogger())
	cpu.registers.PC = PCTestData

	for i := uint16(0); i < 4; i++ {
		got := cpu.registers.GetPCAndIncrement()
		if got != PCTestData+i {
			t.Errorf("got: %d expected %d in the PC", got, PCTestData+i)
		}
	}
}

// TestCPU_StepInstruction just checks that the right instruction is loaded in the CPU
func TestCPU_StepInstruction(t *testing.T) {
	busMock := &BusMock{
		Data: []byte{
			0x00,             // NOP
			0xC3, 0x05, 0x00, // JP 0x0004
			0xAF,       // XOR A
			0x0E, 0x10, // LD C, 0x10
			0x05, // DEC B
		}, // 5 instructions
	}

	testCases := []struct {
		FetchedData          uint16
		CheckFetchData       bool // This indicates that the test needs to check FetchedData
		CurrentOperationCode byte
		CurrentInstruction   Instruction
	}{
		{
			CheckFetchData:       false,
			CurrentOperationCode: 0x00, // NOP
			CurrentInstruction: Instruction{
				Type:           inNop,
				AddressingMode: amImp,
			},
		},
		{
			FetchedData:          0x0005,
			CheckFetchData:       true,
			CurrentOperationCode: 0xC3, // JP a16
			CurrentInstruction: Instruction{
				Type:           inJp,
				AddressingMode: amD16,
			},
		},
		{
			CheckFetchData:       false,
			CurrentOperationCode: 0xAF, // XOR A
			CurrentInstruction: Instruction{
				Type:           inXor,
				AddressingMode: amR,
				RegisterType1:  rtA,
			},
		},
		{
			FetchedData:          0x10,
			CheckFetchData:       true,
			CurrentOperationCode: 0x0E, // LD C, 0x10
			CurrentInstruction: Instruction{
				Type:           inLd,
				AddressingMode: amRnD8,
				RegisterType1:  rtC,
			},
		},
		{
			CheckFetchData:       false,
			CurrentOperationCode: 0x05, // DEC B
			CurrentInstruction: Instruction{
				Type:           inDec,
				AddressingMode: amR,
				RegisterType1:  rtB,
			},
		},
	}

	// Init CPU
	cpu := Init(busMock, log.NewBuiltinStdoutLogger())
	for _, testCase := range testCases {
		_ = cpu.Step()

		// Do assertions
		if cpu.CurrentOperationCode != testCase.CurrentOperationCode {
			t.Errorf("got %d expected %d for current operation code",
				cpu.CurrentOperationCode, testCase.CurrentOperationCode)
		}

		if testCase.CheckFetchData {
			if cpu.FetchedData != testCase.FetchedData {
				t.Errorf("got %d expected %d for fetched data", cpu.FetchedData, testCase.FetchedData)
			}
		}

		if !reflect.DeepEqual(cpu.CurrentInstruction, testCase.CurrentInstruction) {
			t.Errorf("got %v expected %v for CPU current instruction", cpu.CurrentInstruction, testCase.CurrentInstruction)
		}
	}
}
