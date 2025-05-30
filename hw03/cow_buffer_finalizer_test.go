package hw03

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCOWBufferFinalizer(t *testing.T) {
	refs := 1
	buffer := &COWBuffer{
		data: []byte("test"),
		refs: &refs,
	}

	runtime.SetFinalizer(buffer, (*COWBuffer).Close)
	buffer = nil

	runtime.GC()
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, refs, 0)
}
