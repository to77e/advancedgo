package hw03

import (
	"log/slog"
	"runtime"
	"sync"
	"unsafe"
)

type COWBuffer struct {
	data []byte
	refs *int
	mu   *sync.Mutex
}

func NewCOWBuffer(data []byte) COWBuffer {
	copiedData := make([]byte, len(data))
	copy(copiedData, data)
	refs := 1

	buf := COWBuffer{
		data: data,
		refs: &refs,
		mu:   &sync.Mutex{},
	}
	runtime.SetFinalizer(&buf, (*COWBuffer).Close)
	return buf
}

func (b *COWBuffer) Clone() COWBuffer {
	if b.refs == nil {
		slog.Debug("cloned closed buffer")
		return COWBuffer{}
	}

	b.mu.Lock()
	*b.refs++
	b.mu.Unlock()

	clone := COWBuffer{
		data: b.data,
		refs: b.refs,
		mu:   b.mu,
	}
	runtime.SetFinalizer(&clone, (*COWBuffer).Close)
	return clone
}

func (b *COWBuffer) Close() {
	slog.Debug("closing buffer")

	if b.refs == nil {
		slog.Debug("closed twice")
		return
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	*b.refs--
	if *b.refs == 0 {
		b.data, b.refs, b.mu = nil, nil, nil
	}

	runtime.SetFinalizer(b, nil)
}

func (b *COWBuffer) Update(index int, value byte) bool {
	if b.refs == nil {
		slog.Debug("update on closed buffer")
		return false
	}

	if index < 0 || index >= len(b.data) {
		slog.Debug("index out of bounds")
		return false
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	if *b.refs > 1 {
		newData := make([]byte, len(b.data))
		copy(newData, b.data)

		*b.refs--
		newRefs := 1

		b.data = newData
		b.refs = &newRefs
		b.mu = &sync.Mutex{}
	}

	b.data[index] = value
	return true
}

func (b *COWBuffer) String() string {
	if b.data == nil {
		slog.Debug("string on closed buffer")
		return ""
	}
	return unsafe.String(unsafe.SliceData(b.data), len(b.data))
}
