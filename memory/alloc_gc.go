//go:build !tinygo && !wasm && !wasi && !tinygo.wasm

package memory

import (
	"unsafe"
)

//go:linkname mallocgc runtime.mallocgc
func mallocgc(size uintptr, typ unsafe.Pointer, needzero bool) unsafe.Pointer

func GCAlloc(size uintptr) unsafe.Pointer {
	return mallocgc(size, nil, false)
	//b := make([]byte, size)
	//return unsafe.Pointer(&b[0])
}

func GCAllocZeroed(size uintptr) unsafe.Pointer {
	return mallocgc(size, nil, true)
}
