/*
	Example: Using an object method stored within a struct

	https://stackoverflow.com/questions/43591047/converting-a-pointer-to-a-byte-slice
 */


package main

import (
	"encoding/hex"
	"fmt"
	"unsafe"
)

type OneType struct {
	x int
}

func (o *OneType) write() {
	fmt.Printf("\tOneType::write(): %d\n", o.x)
}

type TwoType struct {
	write	func()
}

type SliceFake struct {
	addr	uintptr
	len		int
	cap		int
}

func main() {
	var sl1		*SliceFake
	var sl2		*SliceFake

	fmt.Println("Test...")
	
	one1 := &OneType{1}
	one2 := &OneType{2}
	two1 := TwoType{one1.write}
	two1.write()
	two1.write = one2.write
	two1.write()
	fmt.Printf("\t%T\n", two1.write)
	fmt.Printf("\tsizeof(TwoType): %d\n", unsafe.Sizeof(TwoType{}))
	fmt.Printf("\tsizeof(&func): %d\n", unsafe.Sizeof(func(){}))
	fmt.Printf("\taddr(two1.write): %x\n", uintptr(unsafe.Pointer(&two1)))
	pp := (*int32)(unsafe.Pointer(&two1))
	fmt.Printf("\taddr(*two1.write): %x\n", *pp)
	sl1 = &SliceFake{uintptr(unsafe.Pointer(&two1)), 32, 32}
	fmt.Printf("\ttwo1:\n%s\n", hex.Dump(*(*[]byte)(unsafe.Pointer(sl1))))
	sl2 = &SliceFake{uintptr(unsafe.Pointer(sl1.addr)), 32, 32}
	fmt.Printf("\ttwo1:\n%s\n", hex.Dump(*(*[]byte)(unsafe.Pointer(sl2))))

	
	fmt.Println("...end of Test")
}

