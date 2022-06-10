package cgo

/*
#cgo CXXFLAGS: -std=c++20 -I./
#cgo LDFLAGS: -lstdc++
#include "trampoline.h"
*/
import "C"
import "unsafe"

var (
	Stub  = C.pidato_stub
	Sleep = C.pidato_sleep // pidato_sleep(u64 nanoseconds)
)

func CGO() {
	C.pidato_stub()
}

func NonBlocking(fn *byte, arg0, arg1 uintptr) {
	Blocking(fn, arg0, arg1)
}

func Blocking(fn *byte, arg0, arg1 uintptr) {
	C.pidato_trampoline((C.size_t)(uintptr(unsafe.Pointer(fn))), (C.size_t)(arg0), (C.size_t)(arg1))
}
