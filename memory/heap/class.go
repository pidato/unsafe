package heap

import "math/bits"

type Class struct {
	Index uint16
	Shift uint16
	Size  uint32
}

func (cls Class) Set(bitmap uint32) uint32 {
	return bitmap | (1 << cls.Shift)
}

func (cls Class) Set64(bitmap uint64) uint64 {
	return bitmap | (1 << cls.Shift)
}

func (cls Class) Unset(bitmap uint32) uint32 {
	return bitmap & ^(1 << cls.Shift)
}

func (cls Class) Unset64(bitmap uint64) uint64 {
	return bitmap & ^(1 << cls.Shift)
}

func (cls Class) Search(bitmap uint32) int {
	return bits.LeadingZeros32(bitmap<<cls.Index) + int(cls.Index)
}

func (cls Class) Search64(bitmap uint64) int {
	return bits.LeadingZeros64(bitmap<<cls.Index) + int(cls.Index)
}

// alignUp rounds n up to a multiple of a. a must be a power of 2.
func alignUp(n, a uintptr) uintptr {
	return (n + a - 1) &^ (a - 1)
}

// alignDown rounds n down to a multiple of a. a must be a power of 2.
func alignDown(n, a uintptr) uintptr {
	return n &^ (a - 1)
}

// divRoundUp returns ceil(n / a).
func divRoundUp(n, a uintptr) uintptr {
	// a is generally a power of two. This will get inlined and
	// the compiler will optimize the division.
	return (n + a - 1) / a
}
