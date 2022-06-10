package cgo

import (
	"github.com/pidato/unsafe/cgo/cgo"
	"runtime"
	"testing"
	"time"
)

func BenchmarkCall(b *testing.B) {
	b.Run("Assembly Trampoline Call", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NonBlocking((*byte)(cgo.Stub), 0, 0)
		}
	})
	b.Run("CGO Trampoline Call", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cgo.Blocking((*byte)(cgo.Stub), 0, 0)
		}
	})
	b.Run("CGO Standard", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cgo.CGO()
		}
	})
}

func TestSleep(t *testing.T) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	for i := 0; i < 10000; i++ {
		NonBlocking((*byte)(cgo.Sleep), uintptr(time.Second), 0)
		println(time.Now().UnixNano())
	}
}
