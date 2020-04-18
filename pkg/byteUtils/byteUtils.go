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
