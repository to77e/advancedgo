package hw10

import (
	"unsafe"
)

func Defragment(memory []byte, pointers []unsafe.Pointer) {
	if len(memory) == 0 || len(pointers) == 0 {
		return
	}

	base := uintptr(unsafe.Pointer(&memory[0]))
	write := uintptr(0)

	for read := uintptr(0); read < uintptr(len(memory)); read++ {
		var usedHere bool
		for i, p := range pointers {
			if uintptr(p)-base == read {
				usedHere = true
				if read != write {
					memory[write] = memory[read]
				}
				pointers[i] = unsafe.Pointer(&memory[write])
				break
			}
		}
		if usedHere {
			write++
		}
	}

	for i := write; i < uintptr(len(memory)); i++ {
		memory[i] = 0
	}
}
