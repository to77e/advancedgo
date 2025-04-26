package hw01

import "unsafe"

func ToLittleEndian(number uint32) uint32 {
	return (number&0x000000FF)<<24 | (number&0x0000FF00)<<8 | (number&0x00FF0000)>>8 | (number&0xFF000000)>>24
}

func ToLittleEndianGeneric[T uint64 | uint32 | uint16](number T) (result T) {
	size := int(unsafe.Sizeof(number))

	for i := 0; i < size; i++ {
		mask := T(0xFF) << (i * 8)
		xByte := byte((number & mask) >> (i * 8))
		shift := ((size - 1) - i) * 8
		result |= T(xByte) << shift
	}

	return
}
