package ppu

const (
	dmgPaletteAttrFlagBitPos = 4
	xFlipAttrFlagBitPos      = 5
	yFlipAttrFlagBitPos      = 6
	priorityAttrFlagBitPos   = 7
)

type OamEntry struct {
	y         byte
	x         byte
	tileIndex byte
	/*
			attributesFlag
		 Bit7   BG and Window over OBJ (0=No, 1=BG and Window colors 1-3 over the OBJ)
		 Bit6   Y flip          (0=Normal, 1=Vertically mirrored)
		 Bit5   X flip          (0=Normal, 1=Horizontally mirrored)
		 Bit4   Palette number  **Non CGB Mode Only** (0=OBP0, 1=OBP1)
		 Bit3   Tile VRAM-Bank  **CGB Mode Only**     (0=Bank 0, 1=Bank 1)
		 Bit2-0 Palette number  **CGB Mode Only**     (OBP0-7)
	*/
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
