package hw03

import (
	"runtime"
	"unsafe"
)

type COWBuffer struct {
	data []byte
	refs *int
}

func NewCOWBuffer(data []byte) COWBuffer {
	refs := 1
	buf := COWBuffer{
		data: data,
		refs: &refs,
	}
	runtime.SetFinalizer(&buf, (*COWBuffer).Close)
	return buf
}

func (b *COWBuffer) Clone() COWBuffer {
	*b.refs++
	clone := COWBuffer{
		data: b.data,
		refs: b.refs,
	}
	runtime.SetFinalizer(&clone, (*COWBuffer).Close)
	return clone
}

func (b *COWBuffer) Close() {
	*b.refs--
	runtime.SetFinalizer(b, nil)
}

func (b *COWBuffer) Update(index int, value byte) bool {
	if index < 0 || index >= len(b.data) {
		return false
	}

	if *b.refs > 1 {
		copiedData := make([]byte, len(b.data))
		copy(copiedData, b.data)
		newBuff := NewCOWBuffer(copiedData)

		*b.refs--
		b.data = newBuff.data
		b.refs = newBuff.refs
	}

	b.data[index] = value
	return true
}

func (b *COWBuffer) String() string {
	return unsafe.String(unsafe.SliceData(b.data), len(b.data))
}
