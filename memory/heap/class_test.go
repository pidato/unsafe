package heap

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

const (
	SMALL_DIV  int = 4
	SMALL_MAX  int = 508
	MEDIUM_DIV int = 256
	MEDIUM_MAX int = 16380
	LARGE_MAX  int = 65536 * 256
	HUGE_MAX   int = 65536 * 256
)

// divRoundUp returns ceil(n / a).
func divRoundUpInt(n, a int) int {
	// a is generally a power of two. This will get inlined and
	// the compiler will optimize the division.
	return (n + a - 1) / a
}

func nextPowerOf2(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	return v + 1
}

var quark = []int{
	12, 28, 60, 124,
}

var nano = []int{
	12, 28, 44, 60, 124, 252, 508, 1020,
}
var nano16 = []int{
	14, 30, 46, 62, 126, 254, 510, 1020,
}

var micro = []int{
	8, 12, 16, 20, 28, 36, 44, 60, 76, 92,
	124, 160, 208, 252, 380, 508, 764, 1020,
	2044, 4092, 8188,
}

type Nano struct {
	Free [8]uint16
}

var small = []int{
	12, 16, 20, 24, 28, 36, 44, 60, 76, 92, 124, 156, 188, 220, 252, 300, 364, 428, 508, 764,
	1020, 1280, 1532, 2044, 4092, 8188, 16380,
}

var medium = []int{
	12, 16, 20, 24, 32, 36, 40, 44, 48, 52, 56, 60, 64, 68, 72, 76, 80, 84, 88, 92, 96, 104, 112, 120, 128, 136, 144, 152, 160, 168, 176, 192, 216, 232, 256, 384, 512, 768, 1024,
	2048, 4096, 8192, 16384, 32768, 65536,
}

var large = []int{
	12, 16, 20, 24, 28, 36, 44, 60, 76, 92, 124, 156, 188, 220, 252, 300, 364, 428, 508, 764,
	1020, 1280, 1532, 2044, 4092, 8188, 16380, 32764,
	65534, 131098, 262140, 524284, 1048572, 2097148, 4194300, 8388604, 16777212,
}

func TestGenerate(t *testing.T) {
	//{
	//	g := *NewSizeClassGen(
	//		4,
	//		4,
	//		SMALL_MAX,
	//		256,
	//		4096,
	//		4096,
	//		32768,
	//		small,
	//		//[]int{
	//		//	12, 16, 24, 32, 48, 64, 80, 96, 128, 256, 384, 512, 1024,
	//		//	2048, 4096,
	//		//	8192, 16384, 32768},
	//	)
	//
	//	runtime.SetFinalizer(&g, func(s *SizeClassGen) {
	//		fmt.Println("SetFinalizer")
	//	})
	//	runtime.GC()
	//}
	//runtime.GC()
	fmt.Println("Small Size class count:", len(small))
	g := NewSizeClassGen(
		4,
		SMALL_DIV,
		SMALL_MAX,
		MEDIUM_DIV,
		small[len(small)-1]+4,
		0,
		0,
		NewSizeClasses(small...),
		//[]int{
		//	12, 16, 24, 32, 48, 64, 80, 96, 128, 256, 384, 512, 1024,
		//	2048, 4096,
		//	8192, 16384, 32768},
	)

	a := divRoundUpInt(3976-g.Medium.Min, MEDIUM_DIV)
	fmt.Println(a)
	fmt.Println(a * MEDIUM_DIV)
	fmt.Println(a*MEDIUM_DIV + g.Medium.Min)

	fmt.Println("")
	PrintSizeClassesStructInitializer(g.Classes)
	fmt.Println("")
	g.Small.PrintSizeClassesStructInitializer()
	fmt.Println("")
	g.Medium.PrintSizeClassesStructInitializer()
	fmt.Println("")
	//g.Small.PrintIndexes()
	//g.Small.PrintClasses()
	//g.Medium.PrintIndexes()
	//g.Medium.PrintClasses()

	//g.printSizeClass(0)
	//g.printSizeClass(1)
	//g.printSizeClass(9)
	//g.printSizeClass(12)
	//g.printSizeClass(13)
	//g.printSizeClass(21)
	//g.printSizeClass(29)
	//g.printSizeClass(33)
	//g.printSizeClass(36)
	//g.printSizeClass(43)
	//g.printSizeClass(59)
	//g.printSizeClass(78)
	//g.printSizeClass(256)
	//g.printSizeClass(768)
	//g.printSizeClass(1024)
	//g.printSizeClass(1025)
	g.printSizeClass(3976)
}

func (s *SizeClassGen) printSizeClass(size int) {

	cls := s.Get(size)
	if cls == nil {
		fmt.Println("Size:", size, "  is out of bounds!")
		return
	}
	fmt.Println("["+strconv.Itoa(cls.Index)+"]"+" Size:", size, " Class:", cls.Size, " Shift:", cls.Shift)
}

type SizeClass struct {
	Index int
	Shift int
	Size  int
}

func NewSizeClasses(sizes ...int) []SizeClass {
	sort.Ints(sizes)
	classes := make([]SizeClass, len(sizes))
	for i := 0; i < len(sizes); i++ {
		classes[i] = SizeClass{
			Index: i,
			Shift: 31 - i,
			Size:  sizes[i],
		}
	}
	return classes
}

type SizeClassGen struct {
	wordSize       int
	Classes        []SizeClass
	Small          *ClassRange
	Medium         *ClassRange
	Large          *ClassRange
	smallBoundary  int
	mediumBoundary int
	largeBoundary  int
}

func NewSizeClassGen(
	wordSize int,
	smallSizeDiv int,
	smallSizeMax int,
	mediumSizeDiv int,
	mediumSizeMax int,
	largeSizeDiv int,
	largeSizeMax int,
	classes []SizeClass) *SizeClassGen {

	if smallSizeMax < 8 || smallSizeDiv < 2 {
		return nil
	}

	s := &SizeClassGen{wordSize: wordSize, Classes: classes}
	s.Small = NewClassRange("small", 0, smallSizeMax, smallSizeDiv, classes)

	if mediumSizeMax > 0 && closestFit(smallSizeMax+mediumSizeDiv, classes).Size > 0 {
		s.Medium = NewClassRange("medium", smallSizeMax, mediumSizeMax, mediumSizeDiv, classes)
	}
	if mediumSizeMax > 0 && largeSizeMax > 0 && closestFit(mediumSizeMax+largeSizeDiv, classes).Size > 0 {
		s.Large = NewClassRange("large", mediumSizeMax, largeSizeMax, largeSizeDiv, classes)
	}
	s.smallBoundary = smallSizeMax + 1
	if s.Medium != nil {
		s.mediumBoundary = mediumSizeMax + 1
	} else {
		s.mediumBoundary = 0
	}
	if s.Large != nil {
		s.largeBoundary = largeSizeMax + 1
	} else {
		s.largeBoundary = 0
	}

	return s
}

func (s *SizeClassGen) Get(size int) *SizeClass {
	if size < s.smallBoundary {
		return s.Small.Get(size)
	} else if size < s.mediumBoundary {
		return s.Medium.Get(size)
	} else if size < s.largeBoundary {
		return s.Large.Get(size)
	}
	return nil
}

type ClassRange struct {
	Name string
	Div  int
	Min  int
	Max  int
	//Map     []int
	Classes []SizeClass
}

func NewClassRange(name string, min, max, div int, classes []SizeClass) *ClassRange {
	max = max / div * div
	r := &ClassRange{
		Name: name,
		Min:  min,
		Max:  max,
		Div:  div,
		//Map:     make([]int, (max-min)/div),
		Classes: make([]SizeClass, (max-min)/div),
	}
	for index := 0; index < len(r.Classes); index++ {
		size := min + (index * div)
		//if size == 0 {
		//	small[0] = 0
		//	continue
		//}
		//size = divRoundUpInt(size, div)*div + min

		// Find closes size
		classIndex := closestFit(size, classes)
		r.Classes[index] = classIndex
	}
	return r
}

func (c *ClassRange) Get(size int) *SizeClass {
	index := divRoundUpInt(size-c.Min, c.Div)
	return &c.Classes[index]
}

func (c *ClassRange) PrintSizeClassesStructInitializer() {
	PrintSizeClassesStructInitializer(c.Classes)
}

func PrintSizeClassesStructInitializer(classes []SizeClass) {
	b := strings.Builder{}
	for i, cls := range classes {
		if i%3 == 0 && i > 0 {
			b.WriteString("\n")
		}

		b.WriteString("{")
		b.WriteString(strconv.Itoa(cls.Index))
		b.WriteString(",")
		b.WriteString(strconv.Itoa(cls.Shift))
		b.WriteString(",")
		b.WriteString(strconv.Itoa(cls.Size))
		b.WriteString("},")

		//if i < 4 {
		//	if i == 2 {
		//		b.WriteString("\n")
		//	}
		//	//if i == 2 {
		//	//	b.WriteString("\n")
		//	//}
		//} else if i%3 == 0 {
		//	b.WriteString("\n")
		//}
	}
	fmt.Println(b.String())
}

func (c *ClassRange) PrintClasses() {
	PrintSizeClassesStructInitializer(c.Classes)
}

func closestFit(size int, classes []SizeClass) SizeClass {
	for i := 0; i < len(classes); i++ {
		cls := classes[i]
		if size <= cls.Size {
			return cls
		}
	}
	return SizeClass{}
}
