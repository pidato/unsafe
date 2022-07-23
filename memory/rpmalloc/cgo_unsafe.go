//go:build !tinygo && (amd64 || arm64)
// +build !tinygo
// +build amd64 arm64

package rpmalloc

/*
//#cgo darwin,amd64 LDFLAGS: -Wl,-rpath,${SRCDIR}/lib/darwin_amd64 -L${SRCDIR}/lib/darwin_amd64
#cgo darwin,amd64 LDFLAGS: -ldl -lc -lm
#cgo linux,amd64 CFLAGS: -I${SRCDIR}/src
//#cgo linux,amd64 LDFLAGS: -Wl,-rpath,${SRCDIR}/lib/linux_amd64 -L${SRCDIR}/lib/linux_amd64
#cgo linux,amd64 LDFLAGS: -ldl -lc -lm
#cgo linux,amd64 CFLAGS: -D_GNU_SOURCE
#include "rpmalloc.h"
#include <rpmalloc.h>
#include <stdlib.h>
#include <inttypes.h>
#include <string.h>

typedef struct {
	size_t size;
	size_t ptr;
} malloc_t;

void do_malloc(uintptr_t arg0, uintptr_t arg1) {
	malloc_t* args = (malloc_t*)arg0;
	args->ptr = (size_t)malloc((size_t)args->size);
}
void do_free(size_t ptr, size_t arg1, uintptr_t arg2, uintptr_t arg3) {
	free((void*)ptr);
}

void do_rpmalloc_thread_statistics(uintptr_t arg0, uintptr_t arg1) {
	rpmalloc_thread_statistics((rpmalloc_thread_statistics_t*)(void*)arg0);
}

void do_rpmalloc_global_statistics(size_t arg0, size_t arg1) {
	rpmalloc_global_statistics((rpmalloc_global_statistics_t*)(void*)arg0);
}

void do_rpmalloc(size_t arg0, size_t arg1) {
	malloc_t* args = (malloc_t*)arg0;
	args->ptr = (size_t)rpmalloc((size_t)args->size);
}

void do_rpmalloc_zero(size_t arg0, size_t arg1) {
	malloc_t* args = (malloc_t*)arg0;
	args->ptr = (size_t)rpmalloc((size_t)args->size);
	memset((void*)args->ptr, (int)0, (size_t)args->size);
}

typedef struct {
	size_t size;
	size_t ptr;
	size_t cap;
} malloc_cap_t;

void do_rpmalloc_cap(size_t arg0, size_t arg1) {
	malloc_cap_t* args = (malloc_cap_t*)arg0;
	args->ptr = (size_t)rpmalloc((size_t)args->size);
	args->cap = (size_t)rpmalloc_usable_size((void*)args->ptr);
}

void do_rpmalloc_zero_cap(size_t arg0, size_t arg1) {
	malloc_cap_t* args = (malloc_cap_t*)arg0;
	args->ptr = (size_t)rpmalloc((size_t)args->size);
	args->cap = (size_t)rpmalloc_usable_size((void*)args->ptr);
	memset((void*)args->ptr, 0, args->cap);
}

typedef struct {
	size_t num;
	size_t size;
	size_t ptr;
} calloc_t;

void do_rpcalloc(size_t arg0, size_t arg1) {
	calloc_t* args = (calloc_t*)(void*)arg0;
	args->ptr = (size_t)rpcalloc(args->num, args->size);
}

typedef struct {
	size_t num;
	size_t size;
	size_t ptr;
	size_t cap;
} calloc_cap_t;

void do_rpcalloc_cap(size_t arg0, size_t arg1) {
	calloc_cap_t* args = (calloc_cap_t*)(void*)arg0;
	args->ptr = (size_t)rpcalloc(args->num, args->size);
	args->cap = (size_t)rpmalloc_usable_size((void*)args->ptr);
}

typedef struct {
	size_t ptr;
	size_t size;
	size_t newptr;
} realloc_t;

void do_rprealloc(size_t arg0, size_t arg1) {
	realloc_t* args = (realloc_t*)(void*)arg0;
	args->newptr = (size_t)rprealloc((void*)args->ptr, args->size);
}

typedef struct {
	size_t ptr;
	size_t size;
	size_t newptr;
	size_t cap;
} realloc_cap_t;

void do_rprealloc_cap(size_t arg0, size_t arg1) {
	realloc_cap_t* args = (realloc_cap_t*)(void*)arg0;
	args->newptr = (size_t)rprealloc((void*)args->ptr, args->size);
	args->cap = (size_t)rpmalloc_usable_size((void*)args->newptr);
}

void do_rpfree(size_t ptr, size_t arg2) {
	rpfree((void*)ptr);
}

typedef struct {
	size_t ptr;
	size_t size;
} usable_size_t;

void do_rpmalloc_usable_size(size_t arg0, size_t arg1) {
	usable_size_t* args = (usable_size_t*)arg0;
	args->size = (size_t)rpmalloc_usable_size((void*)args->ptr);
}

typedef struct {
	size_t ptr;
} heap_acquire_t;

void do_rpmalloc_heap_acquire(size_t arg0, size_t arg1) {
	heap_acquire_t* args = (heap_acquire_t*)(void*)arg0;
	args->ptr = (size_t)rpmalloc_heap_acquire();
}

typedef struct {
	size_t ptr;
} heap_release_t;

void do_rpmalloc_heap_release(size_t arg0, size_t arg1) {
	heap_release_t* args = (heap_release_t*)(void*)arg0;
	rpmalloc_heap_release((rpmalloc_heap_t*)(void*)args->ptr);
}

typedef struct {
	size_t heap;
	size_t size;
	size_t ptr;
} heap_alloc_t;

void do_rpmalloc_heap_alloc(size_t arg0, size_t arg1) {
	heap_alloc_t* args = (heap_alloc_t*)(void*)arg0;
	args->ptr = (size_t)rpmalloc_heap_alloc((rpmalloc_heap_t*)(void*)args->heap, args->size);
}

typedef struct {
	size_t heap;
	size_t size;
	size_t ptr;
	size_t cap;
} heap_alloc_cap_t;

void do_rpmalloc_heap_alloc_cap(size_t arg0, size_t arg1) {
	heap_alloc_cap_t* args = (heap_alloc_cap_t*)(void*)arg0;
	args->ptr = (size_t)rpmalloc_heap_alloc((rpmalloc_heap_t*)(void*)args->heap, args->size);
	args->cap = (size_t)rpmalloc_usable_size((void*)args->ptr);
}

typedef struct {
	size_t heap;
	size_t num;
	size_t size;
	size_t ptr;
} heap_calloc_t;

void do_rpmalloc_heap_calloc(size_t arg0, size_t arg1) {
	heap_calloc_t* args = (heap_calloc_t*)(void*)arg0;
	args->ptr = (size_t)rpmalloc_heap_calloc((rpmalloc_heap_t*)(void*)args->heap, args->num, args->size);
}

typedef struct {
	size_t heap;
	size_t num;
	size_t size;
	size_t ptr;
	size_t cap;
} heap_calloc_cap_t;

void do_rpmalloc_heap_calloc_cap(size_t arg0, size_t arg1) {
	heap_calloc_cap_t* args = (heap_calloc_cap_t*)(void*)arg0;
	args->ptr = (size_t)rpmalloc_heap_calloc((rpmalloc_heap_t*)(void*)args->heap, args->num, args->size);
	args->cap = (size_t)rpmalloc_usable_size((void*)args->ptr);
}

typedef struct {
	size_t heap;
	size_t ptr;
	size_t size;
	size_t newptr;
	int flags;
} heap_realloc_t;

void do_rpmalloc_heap_realloc(size_t arg0, size_t arg1) {
	heap_realloc_t* args = (heap_realloc_t*)(void*)arg0;
	args->newptr = (size_t)rpmalloc_heap_realloc((rpmalloc_heap_t*)(void*)args->heap, (void*)args->ptr, args->size, args->flags);
}

typedef struct {
	size_t heap;
	size_t ptr;
	size_t size;
	size_t newptr;
	size_t cap;
	int flags;
} heap_realloc_cap_t;

void do_rpmalloc_heap_realloc_cap(size_t arg0, size_t arg1) {
	heap_realloc_cap_t* args = (heap_realloc_cap_t*)(void*)arg0;
	args->newptr = (size_t)rpmalloc_heap_realloc((rpmalloc_heap_t*)(void*)args->heap, (void*)args->ptr, args->size, args->flags);
	args->cap = (size_t)rpmalloc_usable_size((void*)args->newptr);
}

typedef struct {
	size_t heap;
	size_t ptr;
} heap_free_t;

void do_rpmalloc_heap_free(size_t arg0, size_t arg1) {
	heap_free_t* args = (heap_free_t*)(void*)arg0;
	rpmalloc_heap_free((rpmalloc_heap_t*)(void*)args->heap, (void*)args->ptr);
}

typedef struct {
	size_t heap;
} heap_free_all_t;

void do_rpmalloc_heap_free_all(size_t arg0, size_t arg1) {
	heap_free_all_t* args = (heap_free_all_t*)(void*)arg0;
	rpmalloc_heap_free_all((rpmalloc_heap_t*)(void*)args->heap);
}
*/
import "C"
import (
	"github.com/pidato/unsafe/cgo"
	"unsafe"
)

// ReadThreadStats get thread statistics
func ReadThreadStats(stats *ThreadStats) {
	cgo.NonBlocking((*byte)(C.rpmalloc_thread_statistics), uintptr(unsafe.Pointer(stats)), 0)
}

// ReadGlobalStats get global statistics
func ReadGlobalStats(stats *GlobalStats) {
	cgo.NonBlocking((*byte)(C.do_rpmalloc_global_statistics), uintptr(unsafe.Pointer(stats)), 0)
}

// Malloc allocate a memory block of at least the given size
func StdMalloc(size uintptr) uintptr {
	args := struct {
		size uintptr
		ptr  uintptr
	}{size: size}
	ptr := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_malloc), ptr, 0)
	return args.ptr
}

// Free the given memory block
func StdFree(ptr uintptr) {
	cgo.NonBlocking((*byte)(C.do_free), ptr, 0)
}

// Malloc allocate a memory block of at least the given size
func Malloc(size uintptr) uintptr {
	args := struct {
		size uintptr
		ptr  uintptr
	}{size: size}
	ptr := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpmalloc), ptr, 0)
	return args.ptr
}

// Zero clears n bytes starting at ptr.
//
// Usually you should use typedmemclr. memclrNoHeapPointers should be
// used only when the caller knows that *ptr contains no heap pointers
// because either:
//
// *ptr is initialized memory and its type is pointer-free, or
//
// *ptr is uninitialized memory (e.g., memory that's being reused
// for a new allocation) and hence contains only "junk".
//
// memclrNoHeapPointers ensures that if ptr is pointer-aligned, and n
// is a multiple of the pointer size, then any pointer-aligned,
// pointer-sized portion is cleared atomically. Despite the function
// name, this is necessary because this function is the underlying
// implementation of typedmemclr and memclrHasPointers. See the doc of
// Memmove for more details.
//
// The (CPU-specific) implementations of this function are in memclr_*.s.
//
//go:noescape
//go:linkname memclrNoHeapPointers runtime.memclrNoHeapPointers
func memclrNoHeapPointers(ptr unsafe.Pointer, n uintptr)

// Malloc allocate a memory block of at least the given size
func MallocZeroed(size uintptr) uintptr {
	args := struct {
		size uintptr
		ptr  uintptr
	}{size: size}
	ptr := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpmalloc), ptr, 0)
	// This is faster than memset in C and calloc(1, size)
	memclrNoHeapPointers(unsafe.Pointer(args.ptr), size)
	return args.ptr
}

// MallocCap allocate a memory block of at least the given size
func MallocCap(size uintptr) (uintptr, uintptr) {
	args := struct {
		size uintptr
		ptr  uintptr
		cap  uintptr
	}{size: size}
	ptr := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpmalloc_cap), ptr, 0)
	return args.ptr, args.cap
}

// MallocZeroedCap allocate a memory block of at least the given size
func MallocZeroedCap(size uintptr) (uintptr, uintptr) {
	args := struct {
		size uintptr
		ptr  uintptr
		cap  uintptr
	}{size: size}
	ptr := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpmalloc_zero_cap), ptr, 0)
	return args.ptr, args.cap
}

// Calloc Allocates a memory block of at least the given size and zero initialize it.
func Calloc(num, size uintptr) uintptr {
	args := struct {
		num  uintptr
		size uintptr
		ptr  uintptr
	}{
		num:  num,
		size: size,
	}
	ptr := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpcalloc), ptr, 0)
	return args.ptr
}

// Calloc Allocates a memory block of at least the given size and zero initialize it.
func CallocCap(num, size uintptr) (uintptr, uintptr) {
	args := struct {
		num  uintptr
		size uintptr
		ptr  uintptr
		cap  uintptr
	}{
		num:  num,
		size: size,
	}
	ptr := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpcalloc_cap), ptr, 0)
	return args.ptr, args.cap
}

// Realloc the given block to at least the given size
func Realloc(ptr, size uintptr) uintptr {
	args := struct {
		ptr    uintptr
		size   uintptr
		newptr uintptr
	}{
		ptr:  ptr,
		size: size,
	}
	p := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rprealloc), p, 0)
	return args.newptr
}

// Realloc the given block to at least the given size
func ReallocCap(ptr, size uintptr) (uintptr, uintptr) {
	args := struct {
		ptr    uintptr
		size   uintptr
		newptr uintptr
		cap    uintptr
	}{
		ptr:  ptr,
		size: size,
	}
	p := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rprealloc_cap), p, 0)
	return args.newptr, args.cap
}

// UsableSize Query the usable size of the given memory block (from given pointer to the end of block)
func UsableSize(ptr uintptr) uintptr {
	args := struct {
		ptr uintptr
		ret uintptr
	}{ptr: ptr}
	p := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpmalloc_usable_size), p, 0)
	return args.ret
}

// Free the given memory block
func Free(ptr uintptr) {
	cgo.NonBlocking((*byte)(C.do_rpfree), ptr, 0)
}

func InitThread() {
	C.rpmalloc_thread_initialize()
}

func AcquireHeap() *Heap {
	args := struct {
		ptr uintptr
	}{}
	ptr := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpmalloc_heap_acquire), ptr, 0)
	return (*Heap)(unsafe.Pointer(args.ptr))
}

func (h *Heap) Release() {
	args := struct {
		heap uintptr
	}{heap: uintptr(unsafe.Pointer(h))}
	ptr := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpmalloc_heap_release), ptr, 0)
}

// Alloc Allocate a memory block of at least the given size using the given heap.
func (h *Heap) Alloc(size uintptr) uintptr {
	args := struct {
		heap uintptr
		size uintptr
		ptr  uintptr
	}{heap: uintptr(unsafe.Pointer(h)), size: size}
	ptr := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpmalloc_heap_alloc), ptr, 0)
	return args.ptr
}

// AllocCap Allocate a memory block of at least the given size using the given heap.
func (h *Heap) AllocCap(size uintptr) (uintptr, uintptr) {
	args := struct {
		heap uintptr
		size uintptr
		ptr  uintptr
		cap  uintptr
	}{heap: uintptr(unsafe.Pointer(h)), size: size}
	ptr := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpmalloc_heap_alloc), ptr, 0)
	return args.ptr, args.cap
}

// Calloc Allocate a memory block of at least the given size using the given heap and zero initialize it.
func (h *Heap) Calloc(num, size uintptr) uintptr {
	args := struct {
		heap uintptr
		num  uintptr
		size uintptr
		ptr  uintptr
	}{heap: uintptr(unsafe.Pointer(h)), num: num, size: size}
	ptr := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpmalloc_heap_calloc), ptr, 0)
	return args.ptr
}

// Calloc Allocate a memory block of at least the given size using the given heap and zero initialize it.
func (h *Heap) CallocCap(num, size uintptr) (uintptr, uintptr) {
	args := struct {
		heap uintptr
		num  uintptr
		size uintptr
		ptr  uintptr
		cap  uintptr
	}{heap: uintptr(unsafe.Pointer(h)), num: num, size: size}
	ptr := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpmalloc_heap_calloc_cap), ptr, 0)
	return args.ptr, args.cap
}

// Realloc Reallocate the given block to at least the given size. The memory block MUST be allocated
// by the same heap given to this function.
func (h *Heap) Realloc(ptr, size uintptr, flags int32) uintptr {
	args := struct {
		heap   uintptr
		ptr    uintptr
		size   uintptr
		newptr uintptr
		flags  int32
	}{heap: uintptr(unsafe.Pointer(h)), ptr: ptr, size: size, flags: flags}
	p := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpmalloc_heap_realloc), p, 0)
	return args.newptr
}

// ReallocCap Reallocate the given block to at least the given size. The memory block MUST be allocated
// by the same heap given to this function.
func (h *Heap) ReallocCap(ptr, size uintptr, flags int32) (uintptr, uintptr) {
	args := struct {
		heap   uintptr
		ptr    uintptr
		size   uintptr
		newptr uintptr
		cap    uintptr
		flags  int32
	}{heap: uintptr(unsafe.Pointer(h)), ptr: ptr, size: size, flags: flags}
	p := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpmalloc_heap_realloc_cap), p, 0)
	return args.newptr, args.cap
}

// Free the given memory block from the given heap. The memory block MUST be allocated
// by the same heap given to this function.
func (h *Heap) Free(ptr uintptr) {
	args := struct {
		heap uintptr
		ptr  uintptr
	}{heap: uintptr(unsafe.Pointer(h)), ptr: ptr}
	p := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpmalloc_heap_free), p, 0)
}

// FreeAll memory allocated by the heap
func (h *Heap) FreeAll() {
	args := struct {
		heap uintptr
	}{heap: uintptr(unsafe.Pointer(h))}
	p := uintptr(unsafe.Pointer(&args))
	cgo.NonBlocking((*byte)(C.do_rpmalloc_heap_free_all), p, 0)
}
