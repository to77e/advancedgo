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

	pointerPos := make(map[uintptr]int)
	for i, p := range pointers {
		pointerPos[uintptr(p)-base] = i
	}

	for read := uintptr(0); read < uintptr(len(memory)); read++ {
		if i, exists := pointerPos[read]; exists {
			if read != write {
				memory[write] = memory[read]
			}
			pointers[i] = unsafe.Pointer(&memory[write])
			write++
		}
	}

	for i := write; i < uintptr(len(memory)); i++ {
		memory[i] = 0
	}
}
