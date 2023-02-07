package cpu

// BusMock is used to test CPU
type BusMock struct {
	Data []byte
}

func (b *BusMock) BusRead(address uint16) byte {
	return b.Data[address]
}

func (b *BusMock) BusRead16(address uint16) uint16 {
	return 0xFFFF // TBD!!!
}

func (b *BusMock) BusWrite(address uint16, value byte) {
	if int(address) > len(b.Data) {
		return
	}

	b.Data[address] = value
	return
}

func (b *BusMock) BusWrite16(address uint16, value uint16) {
	return // TBD!!!!
}
