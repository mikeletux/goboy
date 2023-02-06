package bus

// 0000	3FFF	16 KiB ROM bank 00	From cartridge, usually a fixed bank
// 4000	7FFF	16 KiB ROM Bank 01~NN	From cartridge, switchable bank via mapper (if any)
// 8000	9FFF	8 KiB Video RAM (VRAM)	In CGB mode, switchable bank 0/1
// A000	BFFF	8 KiB External RAM	From cartridge, switchable bank if any
// C000	CFFF	4 KiB Work RAM (WRAM)
// D000	DFFF	4 KiB Work RAM (WRAM)	In CGB mode, switchable bank 1~7
// E000	FDFF	Mirror of C000~DDFF (ECHO RAM)	Nintendo says use of this area is prohibited.
// FE00	FE9F	Sprite attribute table (OAM)
// FEA0	FEFF	Not Usable	Nintendo says use of this area is prohibited
// FF00	FF7F	I/O registers
// FF80	FFFE	High RAM (HRAM)
// FFFF	FFFF	Interrupt Enable register (IE)

const (
	RomBank00Start                uint16 = 0x0000
	RomBank00End                  uint16 = 0x3FFF
	RomBank01NNStart              uint16 = 0x4000
	RomBank01NNEnd                uint16 = 0x7FFF
	VramStart                     uint16 = 0x8000
	VramEnd                       uint16 = 0x9FFF
	ExternalRamFromCartridgeStart uint16 = 0xA000
	ExternalRamFromCartridgeEnd   uint16 = 0xBFFF
	WorkRam0Start                 uint16 = 0xC000
	WorkRam0End                   uint16 = 0xCFFF
	WorkRam1Start                 uint16 = 0xD000
	WorkRam1End                   uint16 = 0xDFFF
	EchoRamStart                  uint16 = 0xE000
	EchoRamEnd                    uint16 = 0xFDFF
	OamStart                      uint16 = 0xFE00
	OamEnd                        uint16 = 0xFE9F
	IORegistersStart              uint16 = 0xFF00
	IORegistersEnd                uint16 = 0xFF7F
	HighRamStart                  uint16 = 0xFF80
	HighRamEnd                    uint16 = 0xFFFE
	InterruptEnableRegister       uint16 = 0xFFFF
)
