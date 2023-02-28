package cpu

// BusMapMock is used to test proc funcs
type BusMapMock struct {
	Data map[uint16]byte
}

// NewBusMapMock returns an empty BusMapMock ready to use
func NewBusMapMock() *BusMapMock {
	return &BusMapMock{
		Data: make(map[uint16]byte),
	}
}

func (b *BusMapMock) BusRead(address uint16) byte {
	return b.Data[address]
}

func (b *BusMapMock) BusRead16(address uint16) uint16 {
	low := b.Data[address]
	high := b.Data[address + 1]
	return uint16(high<<8 | low)
}

func (b *BusMapMock) BusWrite(address uint16, value byte) {
	b.Data[address] = value
	return
}

func (b *BusMapMock) BusWrite16(address uint16, value uint16) {
	low := byte(value & 0xFF)
	high := byte(value>>8 & 0xFF)
	b.Data[address] = low
	b.Data[address+1] = high
	return
}
