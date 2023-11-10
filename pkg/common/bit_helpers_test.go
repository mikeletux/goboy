package common

import "testing"

func TestSetBitRegister(t *testing.T) {
	var initialValue byte = 0b01010101
	changedValue := SetBitRegister(initialValue, 1, true)

	if changedValue != 0b01010111 {
		t.Errorf("Expected %08b got %08b", 0b01010111, changedValue)
	}

	changedValue = SetBitRegister(initialValue, 6, false)
	if changedValue != 0b00010101 {
		t.Errorf("Expected %08b got %08b", 0b00010101, changedValue)
	}
}
