package ppu

import (
	"testing"
)

func TestOamEntryAttributesFlagDMGPalette(t *testing.T) {
	oamEntry := OamEntry{
		attributesFlag: 0x0,
	}

	oamEntry.SetDMGPalette(true)
	if oamEntry.attributesFlag != 0b00010000 {
		t.Error("oamEntry.SetDMGPalette didn't set to 1 the needed bit")
	}
	if !oamEntry.GetDMGPalette() {
		t.Error("oamEntry.GetDMGPalette didn't get the needed bit")
	}

	oamEntry.attributesFlag = 0xFF
	oamEntry.SetDMGPalette(false)
	if oamEntry.attributesFlag != 0b11101111 {
		t.Error("oamEntry.SetDMGPalette didn't set to 0 the needed bit")
	}
	if oamEntry.GetDMGPalette() {
		t.Error("oamEntry.GetDMGPalette didn't get the needed bit")
	}
}

func TestOamEntryAttributesFlagXFlip(t *testing.T) {
	oamEntry := OamEntry{
		attributesFlag: 0x0,
	}

	oamEntry.SetXFlip(true)
	if oamEntry.attributesFlag != 0b00100000 {
		t.Error("oamEntry.SetXFlip didn't set to 1 the needed bit")
	}
	if !oamEntry.GetXFlip() {
		t.Error("oamEntry.GetXFlip didn't get the needed bit")
	}

	oamEntry.attributesFlag = 0xFF
	oamEntry.SetXFlip(false)
	if oamEntry.attributesFlag != 0b11011111 {
		t.Error("oamEntry.SetXFlip didn't set to 0 the needed bit")
	}
	if oamEntry.GetXFlip() {
		t.Error("oamEntry.GetXFlip didn't get the needed bit")
	}
}

func TestOamEntryAttributesFlagYFlip(t *testing.T) {
	oamEntry := OamEntry{
		attributesFlag: 0x0,
	}

	oamEntry.SetYFlip(true)
	if oamEntry.attributesFlag != 0b01000000 {
		t.Error("oamEntry.SetYFlip didn't set to 1 the needed bit")
	}
	if !oamEntry.GetYFlip() {
		t.Error("oamEntry.GetYFlip didn't get the needed bit")
	}

	oamEntry.attributesFlag = 0xFF
	oamEntry.SetYFlip(false)
	if oamEntry.attributesFlag != 0b10111111 {
		t.Error("oamEntry.SetYFlip didn't set to 0 the needed bit")
	}
	if oamEntry.GetYFlip() {
		t.Error("oamEntry.GetYFlip didn't get the needed bit")
	}
}

func TestOamEntryAttributesFlagPriority(t *testing.T) {
	oamEntry := OamEntry{
		attributesFlag: 0x0,
	}

	oamEntry.SetPriority(true)
	if oamEntry.attributesFlag != 0b10000000 {
		t.Error("oamEntry.SetPriority didn't set to 1 the needed bit")
	}
	if !oamEntry.GetPriority() {
		t.Error("oamEntry.GetPriority didn't get the needed bit")
	}

	oamEntry.attributesFlag = 0xFF
	oamEntry.SetPriority(false)
	if oamEntry.attributesFlag != 0b01111111 {
		t.Error("oamEntry.SetPriority didn't set to 0 the needed bit")
	}
	if oamEntry.GetPriority() {
		t.Error("oamEntry.GetPriority didn't get the needed bit")
	}
}
