package bus

// MapMock is used to test proc functions
type MapMock struct {
	Data map[uint16]byte
}

// NewMapMock returns an empty MapMock ready to use
func NewMapMock() *MapMock {
	return &MapMock{
		Data: make(map[uint16]byte),
	}
}

func (b *MapMock) BusRead(address uint16) byte {
	return b.Data[address]
}

func (b *MapMock) BusRead16(address uint16) uint16 {
	low := b.Data[address]
	high := b.Data[address + 1]
	return uint16(high<<8 | low)
}

func (b *MapMock) BusWrite(address uint16, value byte) {
	b.Data[address] = value
	return
}

func (b *MapMock) BusWrite16(address uint16, value uint16) {
	low := byte(value & 0xFF)
	high := byte(value>>8 & 0xFF)
	b.Data[address] = low
	b.Data[address+1] = high
	return
}
