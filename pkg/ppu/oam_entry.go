package ppu

const (
	dmgPaletteAttrFlagBitPos = 4
	xFlipAttrFlagBitPos      = 5
	yFlipAttrFlagBitPos      = 6
	priorityAttrFlagBitPos   = 7
)

type OamEntry struct {
	y              byte
	x              byte
	tileIndex      byte
	attributesFlag byte
}

func (o *OamEntry) GetDMGPalette() bool {
	return o.getBitAttributesFlag(dmgPaletteAttrFlagBitPos)
}

func (o *OamEntry) SetDMGPalette(bit bool) {
	o.setBitAttributesFlag(dmgPaletteAttrFlagBitPos, bit)
}

func (o *OamEntry) GetXFlip() bool {
	return o.getBitAttributesFlag(xFlipAttrFlagBitPos)
}

func (o *OamEntry) SetXFlip(bit bool) {
	o.setBitAttributesFlag(xFlipAttrFlagBitPos, bit)
}

func (o *OamEntry) GetYFlip() bool {
	return o.getBitAttributesFlag(yFlipAttrFlagBitPos)
}

func (o *OamEntry) SetYFlip(bit bool) {
	o.setBitAttributesFlag(yFlipAttrFlagBitPos, bit)
}

func (o *OamEntry) GetPriority() bool {
	return o.getBitAttributesFlag(priorityAttrFlagBitPos)
}

func (o *OamEntry) SetPriority(bit bool) {
	o.setBitAttributesFlag(priorityAttrFlagBitPos, bit)
}

func (o *OamEntry) getBitAttributesFlag(position int) bool {
	return (o.attributesFlag>>position)&1 == 1
}

func (o *OamEntry) setBitAttributesFlag(position int, bit bool) {
	if bit {
		o.attributesFlag |= 1 << position
	} else {
		o.attributesFlag &^= 1 << position
	}
}
