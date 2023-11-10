package common

func GetBitMask(regValue byte, bitNumber int, bit bool) byte {
	if bit {
		return regValue | 1<<bitNumber
	}
	return regValue & ^(1 << bitNumber)
}

func GetBitRegister(regValue byte, bitNumber int) bool {
	if (regValue & (1 << bitNumber)) == 0x0 {
		return false
	}
	return true
}

func SetBitRegister(regValue byte, bitNumber int, bitValue bool) byte {
	if bitValue {
		return regValue | 1<<bitNumber
	}

	return regValue & ^(1 << bitNumber)
}

func GetHighAndLowBytes(value uint16) (high, low byte) {
	high, low = byte(value>>8), byte(value&0xFF)
	return
}
