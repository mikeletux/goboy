package cpu

// BusMock is used to test CPU
type BusMock struct {
	Data []byte
}

func (b *BusMock) BusRead(address uint16) byte {
	if int(address) > len(b.Data) {
		return 0
	}

	return b.Data[address]
}
func (b *BusMock) BusWrite(address uint16, value byte) {
	if int(address) > len(b.Data) {
		return
	}

	b.Data[address] = value
	return
}
