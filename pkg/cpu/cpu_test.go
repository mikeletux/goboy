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
	cpu := Init(&BusMapMock{}, &log.NilLogger{})

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
	//cpu.registers.SetAF(AFTestData)
	cpu.registers.SetDataToRegisters(rtAF, AFTestData)
	if cpu.registers.A != ATestData || cpu.registers.F != FTestData {
		t.Errorf("got: A:%d F:%d, expected A:%d F:%d", cpu.registers.A, cpu.registers.F, ATestData, FTestData)
	}

	// Test SetBC
	//cpu.registers.SetBC(BCTestData)
	cpu.registers.SetDataToRegisters(rtBC, BCTestData)
	if cpu.registers.B != BTestData || cpu.registers.C != CTestData {
		t.Errorf("got: B:%d C:%d, expected B:%d C:%d", cpu.registers.B, cpu.registers.C, BTestData, CTestData)
	}

	// Test SetDE
	//cpu.registers.SetDE(DETestData)
	cpu.registers.SetDataToRegisters(rtDE, DETestData)
	if cpu.registers.D != DTestData || cpu.registers.E != ETestData {
		t.Errorf("got: D:%d E:%d, expected D:%d E:%d", cpu.registers.D, cpu.registers.E, DTestData, ETestData)
	}

	// Test SetHL
	// cpu.registers.SetHL(HLTestData)
	cpu.registers.SetDataToRegisters(rtHL, HLTestData)
	if cpu.registers.H != HTestData || cpu.registers.L != LTestData {
		t.Errorf("got: H:%d L:%d, expected H:%d L:%d", cpu.registers.H, cpu.registers.L, HTestData, LTestData)
	}
}

func TestSetGetFlagsFromFlagsRegister(t *testing.T) {
	// Init CPU
	cpu := Init(&BusMapMock{}, &log.NilLogger{})
	cpu.registers = &Registers{
		F: 0x0,
	}

	// Test Get and Set for Z flag
	cpu.registers.SetFZ(true) // Now register should be 0x80
	result := reflect.DeepEqual(cpu.registers.F, byte(0x80))
	if !result {
		t.Errorf("got: %X expected: %X for flag Z on F register", cpu.registers.F, 0x80)
	}

	result = reflect.DeepEqual(cpu.registers.GetFZ(), true)
	if !result {
		t.Errorf("got: %t expected: %t for flag Z on F register", cpu.registers.GetFZ(), true)
	}

	// Test Get and Set for N flag
	cpu.registers.SetFN(true) // Now register should be 0xC0
	result = reflect.DeepEqual(cpu.registers.F, byte(0xC0))
	if !result {
		t.Errorf("got: %X expected: %X for flag N on F register", cpu.registers.F, 0xC0)
	}

	result = reflect.DeepEqual(cpu.registers.GetFN(), true)
	if !result {
		t.Errorf("got: %t expected: %t for flag N on F register", cpu.registers.GetFN(), true)
	}

	// Test Get and Set for H flag
	cpu.registers.SetFH(true) // Now register should be 0xE0
	result = reflect.DeepEqual(cpu.registers.F, byte(0xE0))
	if !result {
		t.Errorf("got: %X expected: %X on F register", cpu.registers.F, 0xE0)
	}

	result = reflect.DeepEqual(cpu.registers.GetFH(), true)
	if !result {
		t.Errorf("got: %t expected: %t for flag H on F register", cpu.registers.GetFH(), true)
	}

	// Test Get and Set for C flag
	cpu.registers.SetFC(true) // Now register should be 0xF0
	result = reflect.DeepEqual(cpu.registers.F, byte(0xF0))
	if !result {
		t.Errorf("got: %X expected: %X on F register", cpu.registers.F, 0xF0)
	}

	result = reflect.DeepEqual(cpu.registers.GetFC(), true)
	if !result {
		t.Errorf("got: %t expected: %t for flag C on F register", cpu.registers.GetFC(), true)
	}

	// Test Get and Set for Z flag
	cpu.registers.SetFZ(false) // Now register should be 0x70
	result = reflect.DeepEqual(cpu.registers.F, byte(0x70))
	if !result {
		t.Errorf("got: %X expected: %X for flag Z on F register", cpu.registers.F, 0x70)
	}

	result = reflect.DeepEqual(cpu.registers.GetFZ(), false)
	if !result {
		t.Errorf("got: %t expected: %t for flag Z on F register", cpu.registers.GetFZ(), false)
	}

	// Test Get and Set for N flag
	cpu.registers.SetFN(false) // Now register should be 0x30
	result = reflect.DeepEqual(cpu.registers.F, byte(0x30))
	if !result {
		t.Errorf("got: %X expected: %X for flag N on F register", cpu.registers.F, 0x30)
	}

	result = reflect.DeepEqual(cpu.registers.GetFN(), false)
	if !result {
		t.Errorf("got: %t expected: %t for flag N on F register", cpu.registers.GetFN(), false)
	}

	// Test Get and Set for H flag
	cpu.registers.SetFH(false) // Now register should be 0x10
	result = reflect.DeepEqual(cpu.registers.F, byte(0x10))
	if !result {
		t.Errorf("got: %X expected: %X on F register", cpu.registers.F, 0x10)
	}

	result = reflect.DeepEqual(cpu.registers.GetFH(), false)
	if !result {
		t.Errorf("got: %t expected: %t for flag H on F register", cpu.registers.GetFH(), false)
	}

	// Test Get and Set for C flag
	cpu.registers.SetFC(false) // Now register should be 0x00
	result = reflect.DeepEqual(cpu.registers.F, byte(0x00))
	if !result {
		t.Errorf("got: %X expected: %X on F register", cpu.registers.F, 0x00)
	}

	result = reflect.DeepEqual(cpu.registers.GetFC(), false)
	if !result {
		t.Errorf("got: %t expected: %t for flag C on F register", cpu.registers.GetFC(), false)
	}
}

func TestFetchDataFromRegisters(t *testing.T) {
	// Init CPU
	cpu := Init(&BusMapMock{}, &log.NilLogger{})

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
	cpu.registers.SetDataToRegisters(rtA, uint16(ATestData))
	cpu.registers.SetDataToRegisters(rtF, uint16(FTestData))
	cpu.registers.SetDataToRegisters(rtB, uint16(BTestData))
	cpu.registers.SetDataToRegisters(rtC, uint16(CTestData))
	cpu.registers.SetDataToRegisters(rtD, uint16(DTestData))
	cpu.registers.SetDataToRegisters(rtE, uint16(ETestData))
	cpu.registers.SetDataToRegisters(rtH, uint16(HTestData))
	cpu.registers.SetDataToRegisters(rtL, uint16(LTestData))
}

func TestGetPCAndIncrement(t *testing.T) {
	// Init CPU
	cpu := Init(&BusMapMock{}, &log.NilLogger{})
	cpu.registers.PC = PCTestData

	for i := uint16(0); i < 4; i++ {
		got := cpu.registers.GetPCAndIncrement()
		if got != PCTestData+i {
			t.Errorf("got: %d expected %d in the PC", got, PCTestData+i)
		}
	}
}
