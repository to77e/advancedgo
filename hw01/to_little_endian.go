package hw01

func ToLittleEndian(number uint32) uint32 {
	var result uint32
	result |= (number & 0x000000FF) << 24
	result |= (number & 0x0000FF00) << 8
	result |= (number & 0x00FF0000) >> 8
	result |= (number & 0xFF000000) >> 24
	return result
}

func ToLittleEndianGeneric[T uint64 | uint32 | uint16](number T) (result T) {
	var size int

	switch any(number).(type) {
	case uint16:
		size = 2
	case uint32:
		size = 4
	case uint64:
		size = 8
	}

	for i := 0; i < size; i++ {
		mask := T(0xFF) << (i * 8)
		xByte := byte((number & mask) >> (i * 8))
		shift := ((size - 1) - i) * 8
		result |= T(xByte) << shift
	}

	return
}
