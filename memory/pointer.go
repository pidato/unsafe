package memory

import (
	"github.com/pidato/unsafe/memory/hash"
	"unsafe"
)

// PtrSize is the size of a pointer in bytes - unsafe.Sizeof(uintptr(0)) but as an ideal constant.
// It is also the size of the machine's native word size (that is, 4 on 32-bit systems, 8 on 64-bit).
const PtrSize = 4 << (^uintptr(0) >> 63)

type Bytes struct {
	Pointer
	Size uintptr
}

// Pointer is a wrapper around a raw pointer that is not unsafe.Pointer
// so Go won't confuse it for a potential GC managed pointer.
type Pointer uintptr

func (p Pointer) Size() int {
	return int(Sizeof(p))
}

func PointerOfString(s string) Pointer {
	h := *(*_string)(unsafe.Pointer(&s))
	return Pointer(h.ptr)
}

//goland:noinspection ALL
func (p Pointer) Unsafe() unsafe.Pointer {
	return unsafe.Pointer(p)
}

// Add is Pointer arithmetic.
func (p Pointer) Add(offset int) Pointer {
	return Pointer(uintptr(int(p) + offset))
}

// Free deallocates memory pointed by Pointer
func (p *Pointer) Free() {
	if p == nil || *p == 0 {
		return
	}
	Free(*p)
	*p = 0
}

// Sizeof returns the size of the allocation provided by the platform allocator.
func (p Pointer) SizeOf() uintptr {
	return Sizeof(p)
}

// Clone the memory starting at offset for size number of bytes and return the new Pointer.
func (p Pointer) Clone(offset, size int) Pointer {
	clone := Alloc(uintptr(size))
	p.Copy(offset, size, clone)
	return clone
}

// Zero zeroes out the entire allocation.
func (p Pointer) Zero(size uintptr) {
	Zero(p.Unsafe(), size)
}

// Move does a memmove

func (p Pointer) Move(offset, size int, to Pointer) {
	Move(to.Unsafe(), unsafe.Add(p.Unsafe(), offset), uintptr(size))
}

// Copy does a memcpy

func (p Pointer) Copy(offset, size int, to Pointer) {
	Copy(to.Unsafe(), unsafe.Add(p.Unsafe(), offset), uintptr(size))
}

// Equals does a memequal

func (p Pointer) Equals(offset, size int, to Pointer) bool {
	return Equals(to.Unsafe(), unsafe.Add(p.Unsafe(), offset), uintptr(size))
}

// Compare does a memcmp

func (p Pointer) Compare(offset, size int, to Pointer) int {
	return Compare(to.Unsafe(), unsafe.Add(p.Unsafe(), offset), uintptr(size))
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Byte
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int8(offset int) int8 {
	return *(*int8)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) Uint8(offset int) uint8 {
	return *(*uint8)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) Byte(offset int) byte {
	return *(*byte)(unsafe.Add(p.Unsafe(), offset))
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Put Byte
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) SetInt8(offset int, v int8) {
	*(*int8)(unsafe.Add(p.Unsafe(), offset)) = v
}

func (p Pointer) SetUint8(offset int, v uint8) {
	*(*uint8)(unsafe.Add(p.Unsafe(), offset)) = v
}

func (p Pointer) SetByte(offset int, v byte) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int16 Native Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int16(offset int) int16 {
	return *(*int16)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetInt16(offset int, v int16) {
	*(*int16)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint16 Native Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint16(offset int) uint16 {
	return *(*uint16)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetUint16(offset int, v uint16) {
	*(*uint16)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int32 Native Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int32(offset int) int32 {
	return *(*int32)(unsafe.Add(p.Unsafe(), offset))
}
func (p Pointer) Int32Alt(offset uintptr) int32 {
	return *(*int32)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetInt32(offset int, v int32) {
	*(*int32)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint32 Native Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint32(offset int) uint32 {
	return *(*uint32)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetUint32(offset int, v uint32) {
	*(*uint32)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int64 Native Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int64(offset int) int64 {
	return *(*int64)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetInt64(offset int, v int64) {
	*(*int64)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint64 Native Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint64(offset int) uint64 {
	return *(*uint64)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetUint64(offset int, v uint64) {
	*(*uint64)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Float32 Native Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Float32(offset int) float32 {
	return *(*float32)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetFloat32(offset int, v float32) {
	*(*float32)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Float64 Native Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Float64(offset int) float64 {
	return *(*float64)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetFloat64(offset int, v float64) {
	*(*float64)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// int
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int(offset int) int {
	return *(*int)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetInt(offset int, v int) {
	*(*int)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// uint
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint(offset int) uint {
	return *(*uint)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetUint(offset int, v uint) {
	*(*uint)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// uintptr
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uintptr(offset int) uintptr {
	return *(*uintptr)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetUintptr(offset int, v uintptr) {
	*(*uintptr)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Pointer
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Pointer(offset int) Pointer {
	return *(*Pointer)(unsafe.Add(p.Unsafe(), offset))
}

func (p Pointer) SetPointer(offset int, v Pointer) {
	*(*Pointer)(unsafe.Add(p.Unsafe(), offset)) = v
}

///////////////////////////////////////////////////////////////////////////////////////////////
// String
///////////////////////////////////////////////////////////////////////////////////////////////

type _string struct {
	ptr uintptr
	len int
}

func (p Pointer) String(offset, size int) string {
	return *(*string)(unsafe.Pointer(&_string{
		ptr: uintptr(int(p) + offset),
		len: size,
	}))
}

func (p Pointer) SetString(offset int, value string) {
	dst := *(*[]byte)(unsafe.Pointer(&_bytes{
		Data: uintptr(int(p) + offset),
		Len:  len(value),
		Cap:  len(value),
	}))
	copy(dst, value)
}

func (p Pointer) SetBytes(offset int, value []byte) {
	dst := *(*[]byte)(unsafe.Pointer(&_bytes{
		Data: uintptr(int(p) + offset),
		Len:  len(value),
		Cap:  len(value),
	}))
	copy(dst, value)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Byte Slice
///////////////////////////////////////////////////////////////////////////////////////////////

type _bytes struct {
	Data uintptr
	Len  int
	Cap  int
}

func (p Pointer) Bytes(offset, length, capacity int) []byte {
	return *(*[]byte)(unsafe.Pointer(&_bytes{
		Data: uintptr(int(p) + offset),
		Len:  length,
		Cap:  capacity,
	}))
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int24 Little Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int24LE(offset int) int32 {
	return int32(*(*byte)(unsafe.Add(p.Unsafe(), offset))) |
		int32(*(*byte)(unsafe.Add(p.Unsafe(), offset+1)))<<8 |
		int32(*(*byte)(unsafe.Add(p.Unsafe(), offset+2)))<<16
}

func (p Pointer) SetInt24LE(offset int, v int32) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = byte(v)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+1)) = byte(v >> 8)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+2)) = byte(v >> 16)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int24 Big Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int24BE(offset int) int32 {
	return int32(*(*byte)(unsafe.Add(p.Unsafe(), offset+2))) |
		int32(*(*byte)(unsafe.Add(p.Unsafe(), offset+1)))<<8 |
		int32(*(*byte)(unsafe.Add(p.Unsafe(), offset)))<<16
}

func (p Pointer) SetInt24BE(offset int, v int32) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = byte(v >> 16)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+1)) = byte(v >> 8)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+2)) = byte(v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint24 Little Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint24LE(offset int) uint32 {
	return uint32(*(*byte)(unsafe.Add(p.Unsafe(), offset))) |
		uint32(*(*byte)(unsafe.Add(p.Unsafe(), offset+1)))<<8 |
		uint32(*(*byte)(unsafe.Add(p.Unsafe(), offset+2)))<<16
}

func (p Pointer) SetUint24LE(offset int, v uint32) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = byte(v)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+1)) = byte(v >> 8)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+2)) = byte(v >> 16)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint24 Big Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint24BE(offset int) uint32 {
	return uint32(*(*byte)(unsafe.Add(p.Unsafe(), offset+2))) |
		uint32(*(*byte)(unsafe.Add(p.Unsafe(), offset+1)))<<8 |
		uint32(*(*byte)(unsafe.Add(p.Unsafe(), offset)))<<16
}

func (p Pointer) SetUint24BE(offset int, v uint32) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = byte(v >> 16)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+1)) = byte(v >> 8)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+2)) = byte(v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int40 Little Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int40LE(offset int) int64 {
	return int64(*(*byte)(unsafe.Add(p.Unsafe(), offset))) |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+1)))<<8 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+2)))<<16 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+3)))<<24 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+4)))<<32
}

func (p Pointer) SetInt40LE(offset int, v int64) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = byte(v)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+1)) = byte(v >> 8)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+2)) = byte(v >> 16)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+3)) = byte(v >> 24)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+4)) = byte(v >> 32)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int40 Big Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int40BE(offset int) int64 {
	return int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+4))) |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+3)))<<8 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+2)))<<16 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+1)))<<24 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset)))<<32
}

func (p Pointer) SetInt40BE(offset int, v int64) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = byte(v >> 32)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+1)) = byte(v >> 24)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+2)) = byte(v >> 16)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+3)) = byte(v >> 8)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+4)) = byte(v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint40 Little Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint40LE(offset int) uint64 {
	return uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset))) |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+1)))<<8 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+2)))<<16 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+3)))<<24 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+4)))<<32
}

func (p Pointer) SetUint40LE(offset int, v uint64) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = byte(v)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+1)) = byte(v >> 8)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+2)) = byte(v >> 16)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+3)) = byte(v >> 24)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+4)) = byte(v >> 32)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint40 Big Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint40BE(offset int) uint64 {
	return uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+4))) |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+3)))<<8 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+2)))<<16 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+1)))<<24 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset)))<<32
}

func (p Pointer) SetUint40BE(offset int, v uint64) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = byte(v >> 32)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+1)) = byte(v >> 24)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+2)) = byte(v >> 16)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+3)) = byte(v >> 8)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+4)) = byte(v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int48 Little Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int48LE(offset int) int64 {
	return int64(*(*byte)(unsafe.Add(p.Unsafe(), offset))) |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+1)))<<8 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+2)))<<16 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+3)))<<24 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+4)))<<32 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+5)))<<40
}

func (p Pointer) SetInt48LE(offset int, v int64) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = byte(v)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+1)) = byte(v >> 8)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+2)) = byte(v >> 16)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+3)) = byte(v >> 24)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+4)) = byte(v >> 32)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+5)) = byte(v >> 40)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int48 Big Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int48BE(offset int) int64 {
	return int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+5))) |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+4)))<<8 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+3)))<<16 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+2)))<<24 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+1)))<<32 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset)))<<40
}

func (p Pointer) SetInt48BE(offset int, v int64) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = byte(v >> 40)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+1)) = byte(v >> 32)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+2)) = byte(v >> 24)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+3)) = byte(v >> 16)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+4)) = byte(v >> 8)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+5)) = byte(v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint48 Little Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint48LE(offset int) uint64 {
	return uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset))) |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+1)))<<8 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+2)))<<16 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+3)))<<24 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+4)))<<32 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+5)))<<40
}

func (p Pointer) SetUint48LE(offset int, v uint64) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = byte(v)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+1)) = byte(v >> 8)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+2)) = byte(v >> 16)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+3)) = byte(v >> 24)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+4)) = byte(v >> 32)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+5)) = byte(v >> 40)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint48 Big Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint48BE(offset int) uint64 {
	return uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+5))) |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+4)))<<8 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+3)))<<16 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+2)))<<24 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+1)))<<32 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset)))<<40
}

func (p Pointer) SetUint48BE(offset int, v uint64) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = byte(v >> 40)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+1)) = byte(v >> 32)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+2)) = byte(v >> 24)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+3)) = byte(v >> 16)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+4)) = byte(v >> 8)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+5)) = byte(v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int56 Little Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int56LE(offset int) int64 {
	return int64(*(*byte)(unsafe.Add(p.Unsafe(), offset))) |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+1)))<<8 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+2)))<<16 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+3)))<<24 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+4)))<<32 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+5)))<<40 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+6)))<<48
}

func (p Pointer) SetInt56LE(offset int, v int64) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = byte(v)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+1)) = byte(v >> 8)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+2)) = byte(v >> 16)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+3)) = byte(v >> 24)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+4)) = byte(v >> 32)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+5)) = byte(v >> 40)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+6)) = byte(v >> 48)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Int56 Big Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Int56BE(offset int) int64 {
	return int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+6))) |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+5)))<<8 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+4)))<<16 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+3)))<<24 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+2)))<<32 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset+1)))<<40 |
		int64(*(*byte)(unsafe.Add(p.Unsafe(), offset)))<<48
}

func (p Pointer) SetInt56BE(offset int, v int64) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = byte(v >> 48)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+1)) = byte(v >> 40)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+2)) = byte(v >> 32)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+3)) = byte(v >> 24)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+4)) = byte(v >> 16)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+5)) = byte(v >> 8)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+6)) = byte(v)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint56 Little Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint56LE(offset int) uint64 {
	return uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset))) |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+1)))<<8 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+2)))<<16 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+3)))<<24 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+4)))<<32 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+5)))<<40 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+6)))<<48
}

func (p Pointer) SetUint56LE(offset int, v uint64) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = byte(v)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+1)) = byte(v >> 8)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+2)) = byte(v >> 16)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+3)) = byte(v >> 24)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+4)) = byte(v >> 32)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+5)) = byte(v >> 40)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+6)) = byte(v >> 48)
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Uint56 Big Endian
///////////////////////////////////////////////////////////////////////////////////////////////

func (p Pointer) Uint56BE(offset int) uint64 {
	return uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+6))) |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+5)))<<8 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+4)))<<16 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+3)))<<24 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+2)))<<32 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset+1)))<<40 |
		uint64(*(*byte)(unsafe.Add(p.Unsafe(), offset)))<<48
}

func (p Pointer) SetUint56BE(offset int, v uint64) {
	*(*byte)(unsafe.Add(p.Unsafe(), offset)) = byte(v >> 48)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+1)) = byte(v >> 40)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+2)) = byte(v >> 32)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+3)) = byte(v >> 24)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+4)) = byte(v >> 16)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+5)) = byte(v >> 8)
	*(*byte)(unsafe.Add(p.Unsafe(), offset+6)) = byte(v)
}

func (p Pointer) Hash64(length int) uint64 {
	return hash.Hash(p.Unsafe(), uint64(length), hash.DefaultSeed)
}

func (p Pointer) Hash64At(offset, length int) uint64 {
	return hash.Hash(p.Pointer(offset).Unsafe(), uint64(length), hash.DefaultSeed)
}
