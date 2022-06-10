package main

import (
	"github.com/pidato/unsafe/cgo"
)

func main() {
	//cgo.CGO()
	cgo.NonBlocking((*byte)(nil), 0, 0)
}
