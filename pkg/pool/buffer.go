package pool

import (
	"bytes"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

// GetBuffer returns a reset buffer from the pool.
func GetBuffer() *bytes.Buffer {
	b := bufferPool.Get().(*bytes.Buffer)
	b.Reset()
	return b
}

// PutBuffer returns a buffer to the pool.
func PutBuffer(b *bytes.Buffer) {
	bufferPool.Put(b)
}