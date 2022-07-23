//go:build !tinygo && !wasm && !wasi && !tinygo.wasm && cgo

package memory

import (
	"github.com/pidato/unsafe/memory/rpmalloc"
)

func Init() {}

// Alloc calls Alloc on the system allocator
//export alloc
func Alloc(size uintptr) Pointer {
	return Pointer(rpmalloc.Malloc(size))
}

//export allocCap
func AllocCap(size uintptr) (Pointer, uintptr) {
	p, c := rpmalloc.MallocCap(size)
	return Pointer(p), c
}

func AllocZeroed(size uintptr) Pointer {
	return Pointer(rpmalloc.MallocZeroed(size))
}

func AllocZeroedCap(size uintptr) (Pointer, uintptr) {
	p, c := rpmalloc.MallocZeroedCap(size)
	return Pointer(p), c
}

// Alloc calls Alloc on the system allocator
func Calloc(num, size uintptr) Pointer {
	return Pointer(rpmalloc.Calloc(num, size))
}

func CallocCap(num, size uintptr) (Pointer, uintptr) {
	p, c := rpmalloc.CallocCap(num, size)
	return Pointer(p), c
}

// Realloc calls Realloc on the system allocator
func Realloc(p Pointer, size uintptr) Pointer {
	return Pointer(rpmalloc.Realloc(uintptr(p), size))
}

//export reallocCap
func ReallocCap(p Pointer, size uintptr) (Pointer, uintptr) {
	newPtr, c := rpmalloc.ReallocCap(uintptr(p), size)
	return Pointer(newPtr), c
}

// Free calls Free on the system allocator
func Free(p Pointer) {
	rpmalloc.Free(uintptr(p))
}

func Sizeof(ptr Pointer) uintptr {
	return rpmalloc.UsableSize(uintptr(ptr))
}
