//go:build tinygo.wasm || 386 || amd64 || amd64p32 || arm || arm64 || loong64 || mips64le || mips64p32 || mips64p32le || mipsle || ppc64le || riscv || riscv64 || wasm

package memory

import (
	"math/bits"
	"unsafe"
)

///////////////////////////////////////////////////////////////////////////////////////////////
// Int16 Little Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) AsInt16LE() int16 {
	return *(*int16)(p.Unsafe())
}

func Int16LE(p unsafe.Pointer) int16 {
	return *(*int16)(p)
}

func (p Pointer) Int16LE(offset int) int16 {
	return *(*int16)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetInt16LE(offset int, v int16) {
	*(*int16)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int16 Big Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) AsInt16BE() int16 {
	return int16(bits.ReverseBytes16(*(*uint16)(p.Unsafe())))
}

func (p Pointer) Int16BE(offset int) int16 {
	return int16(bits.ReverseBytes16(*(*uint16)(unsafe.Add(p.Unsafe(), offset))))
}

func (p Pointer) SetInt16BE(offset int, v int16) {
	*(*int16)(unsafe.Add(p.Unsafe(), offset)) = int16(bits.ReverseBytes16(uint16(v)))
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint16 Little Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) AsUint16LE() uint16 {
	return *(*uint16)(p.Unsafe())
}

func Uint16LE(p unsafe.Pointer) uint16 {
	return *(*uint16)(p)
}

func (p Pointer) Uint16LE(offset int) uint16 {
	return *(*uint16)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetUint16LE(offset int, v uint16) {
	*(*uint16)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint16 Big Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) AsUint16BE() uint16 {
	return bits.ReverseBytes16(*(*uint16)(p.Unsafe()))
}

func (p Pointer) Uint16BE(offset int) uint16 {
	return bits.ReverseBytes16(*(*uint16)(unsafe.Add(p.Unsafe(), offset)))
}

func (p Pointer) SetUint16BE(offset int, v uint16) {
	*(*uint16)(unsafe.Add(p.Unsafe(), offset)) = bits.ReverseBytes16(v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int32 Little Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) AsInt32LE() int32 {
	return *(*int32)(p.Unsafe())
}

func (p Pointer) Int32LE(offset int) int32 {
	return *(*int32)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetInt32LE(offset int, v int32) {
	*(*int32)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int32 Big Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) AsInt32BE() int32 {
	return int32(bits.ReverseBytes32(*(*uint32)(p.Unsafe())))
}

func (p Pointer) Int32BE(offset int) int32 {
	return int32(bits.ReverseBytes32(*(*uint32)(unsafe.Pointer(p + Pointer(offset)))))
}

func (p Pointer) SetInt32BE(offset int, v int32) {
	*(*int32)(unsafe.Add(p.Unsafe(), offset)) = int32(bits.ReverseBytes32(uint32(v)))
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint32 Little Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) AsUint32LE() uint32 {
	return *(*uint32)(p.Unsafe())
}

func (p Pointer) Uint32LE(offset int) uint32 {
	return *(*uint32)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetUint32LE(offset int, v uint32) {
	*(*uint32)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint32 Big Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint32BESlow() uint32 {
	return uint32(*(*byte)(unsafe.Add(p.Unsafe(), 3))) |
		uint32(*(*byte)(unsafe.Add(p.Unsafe(), 2)))<<8 |
		uint32(*(*byte)(unsafe.Add(p.Unsafe(), 1)))<<16 |
		uint32(*(*byte)(p.Unsafe()))<<24
}

func (p Pointer) AsUint32BE() uint32 {
	return bits.ReverseBytes32(*(*uint32)(p.Unsafe()))
}

func (p Pointer) Uint32BE(offset int) uint32 {
	return bits.ReverseBytes32(*(*uint32)(unsafe.Add(p.Unsafe(), offset)))
}

func (p Pointer) SetUint32BE(offset int, v uint32) {
	*(*uint32)(unsafe.Add(p.Unsafe(), offset)) = bits.ReverseBytes32(v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int64 Little Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) AsInt64LE() int64 {
	return *(*int64)(p.Unsafe())
}

func (p Pointer) Int64LE(offset int) int64 {
	return *(*int64)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetInt64LE(offset int, v int64) {
	*(*int64)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int64 Big Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) AsInt64BE() int64 {
	return int64(bits.ReverseBytes64(*(*uint64)(p.Unsafe())))
}

func (p Pointer) Int64BE(offset int) int64 {
	return int64(bits.ReverseBytes64(*(*uint64)(unsafe.Add(p.Unsafe(), offset))))
}

func (p Pointer) SetInt64BE(offset int, v int64) {
	*(*int64)(unsafe.Add(p.Unsafe(), offset)) = int64(bits.ReverseBytes64(uint64(v)))
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint64 Little Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) AsUint64LE() uint64 {
	return *(*uint64)(p.Unsafe())
}

func (p Pointer) Uint64LE(offset int) uint64 {
	return *(*uint64)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetUint64LE(offset int, v uint64) {
	*(*uint64)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint64 Big Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) AsUint64BE() uint64 {
	return bits.ReverseBytes64(*(*uint64)(p.Unsafe()))
}

func (p Pointer) Uint64BE(offset int) uint64 {
	return bits.ReverseBytes64(*(*uint64)(unsafe.Add(p.Unsafe(), offset)))
}

func (p Pointer) SetUint64BE(offset int, v uint64) {
	*(*uint64)(unsafe.Add(p.Unsafe(), offset)) = bits.ReverseBytes64(v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Float32 Little Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) AsFloat32LE() float32 {
	return *(*float32)(p.Unsafe())
}

func (p Pointer) Float32LE(offset int) float32 {
	return *(*float32)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetFloat32LE(offset int, v float32) {
	*(*float32)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Float32 Big Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) AsFloat32BE() float32 {
	return float32(bits.ReverseBytes32(*(*uint32)(p.Unsafe())))
}

func (p Pointer) Float32BE(offset int) float32 {
	return float32(bits.ReverseBytes32(*(*uint32)(unsafe.Add(p.Unsafe(), offset))))
}

func (p Pointer) SetFloat32BE(offset int, v float32) {
	*(*float32)(unsafe.Add(p.Unsafe(), offset)) = float32(bits.ReverseBytes32(uint32(v)))
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Float64 Little Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) AsFloat64LE() float64 {
	return *(*float64)(p.Unsafe())
}

func (p Pointer) Float64LE(offset int) float64 {
	return *(*float64)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetFloat64LE(offset int, v float64) {
	*(*float64)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Float64 Big Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) AsFloat64BE() float64 {
	return float64(bits.ReverseBytes64(*(*uint64)(p.Unsafe())))
}

func (p Pointer) Float64BE(offset int) float64 {
	return float64(bits.ReverseBytes64(*(*uint64)(unsafe.Add(p.Unsafe(), offset))))
}

func (p Pointer) SetFloat64BE(offset int, v float64) {
	*(*float64)(unsafe.Add(p.Unsafe(), offset)) = float64(bits.ReverseBytes64(uint64(v)))
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int24 Native Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int24(offset int) int32 {
	return p.Int24LE(offset)
}

func (p Pointer) SetInt24(offset int, v int32) {
	p.SetInt24LE(offset, v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint24 Native Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint24(offset int) uint32 {
	return p.Uint24LE(offset)
}

func (p Pointer) SetUint24(offset int, v uint32) {
	p.SetUint24LE(offset, v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int40 Native Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int40(offset int) int64 {
	return p.Int40LE(offset)
}

func (p Pointer) SetInt40(offset int, v int64) {
	p.SetInt40LE(offset, v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint40 Native Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint40(offset int) uint64 {
	return p.Uint40LE(offset)
}

func (p Pointer) SetUint40(offset int, v uint64) {
	p.SetUint40LE(offset, v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int48 Native Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int48(offset int) int64 {
	return p.Int48LE(offset)
}

func (p Pointer) SetInt48(offset int, v int64) {
	p.SetInt48LE(offset, v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint48 Native Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint48(offset int) uint64 {
	return p.Uint48LE(offset)
}

func (p Pointer) SetUint48(offset int, v uint64) {
	p.SetUint48LE(offset, v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int56 Native Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int56(offset int) int64 {
	return p.Int56LE(offset)
}

func (p Pointer) SetInt56(offset int, v int64) {
	p.SetInt56LE(offset, v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint56 Native Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint56(offset int) uint64 {
	return p.Uint56LE(offset)
}

func (p Pointer) SetUint56(offset int, v uint64) {
	p.SetUint56LE(offset, v)
}
