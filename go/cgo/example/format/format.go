package main

import "C"
import (
	"fmt"
	"unsafe"
)

////export format
//func format(aRaw, bRaw, cRaw, dRaw uintptr) string {
//	a := int(aRaw)
//	b := math.Float32frombits(uint32(bRaw))
//	c := math.Float64frombits(uint64(cRaw))
//	d := string(*(*[]byte)(unsafe.Pointer(dRaw)))
//	return fmt.Sprintf("int:%d\nfloat32:%G\nfloat64:%G\nstring:%s", a, b, c, d)
//}

//export format
func format(a int, b float32, c float64, d uintptr) uintptr {
	var s = fmt.Sprintf("int:%d\nfloat32:%G\nfloat64:%G\nstring:%s", a, b, c, *(*string)(unsafe.Pointer(d)))
	return uintptr(unsafe.Pointer(&s))
}

//export add
func add(a, b int) int {
	return a + b
}

func main() {}
