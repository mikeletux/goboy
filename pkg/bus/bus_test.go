package bus

import (
	"github.com/mikeletux/goboy/pkg/log"
	"testing"
)

func TestIOTimer(t *testing.T) {
	bus := NewBus(nil, &log.NilLogger{}) // Not going to read into cartridge mappings

	testCases := []struct {
		locationToWrite   uint16
		valueToWrite      byte
		expectedReadValue byte
	}{
		{
			locationToWrite:   divRegisterAddr, // FF04
			valueToWrite:      0xFF,
			expectedReadValue: 0x0,
		},
		{
			locationToWrite:   timaRegisterAddr, // FF05
			valueToWrite:      0xEA,
			expectedReadValue: 0xEA,
		},
		{
			locationToWrite:   tmaRegisterAddr, // FF06
			valueToWrite:      0x1E,
			expectedReadValue: 0x1E,
		},
		{
			locationToWrite:   tacRegisterAddr, // FF07
			valueToWrite:      0b00000111,
			expectedReadValue: 0b00000111,
		},
	}

	// Write test values
	for _, v := range testCases {
		bus.BusWrite(v.locationToWrite, v.valueToWrite)
	}

	// Check that reads are correct
	for _, v := range testCases {
		readValue := bus.BusRead(v.locationToWrite)
		if readValue != v.expectedReadValue {
			t.Errorf("expected %X got %X", v.expectedReadValue, readValue)
		}
	}
}
