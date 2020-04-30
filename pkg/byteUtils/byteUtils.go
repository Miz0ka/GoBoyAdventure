package byteUtils

func BoolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func Swap(octet byte) byte {
	up := octet & 0xF0
	up = up >> 4
	octet = octet << 4
	return octet | up
}

func Val(value byte, bit byte) byte {
	return (value >> bit) & 0x01
}

func TestBit(value byte, bit int) bool {
	return (value>>bit)&0x1 != 0
}
