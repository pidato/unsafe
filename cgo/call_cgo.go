//go:build !amd64 && !arm64 && !tinygo

package cgo

import (
	"github.com/pidato/unsafe/cgo/cgo"
)

func NonBlocking(fn *byte, arg0, arg1 uintptr) {
	cgo.Blocking(fn, arg0, arg1)
}

func Blocking(fn *byte, arg0, arg1 uintptr) {
	cgo.Blocking(fn, arg0, arg1)
}
