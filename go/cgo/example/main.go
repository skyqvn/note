package main

import (
	"fmt"
	"math"
	"syscall"
	"unsafe"
)

import "C"

var FormatFunc *syscall.LazyProc

func Format(a int, b float32, c float64, d string) string {
	sum, _, _ := FormatFunc.Call(uintptr(a), uintptr(math.Float32bits(b)), uintptr(math.Float64bits(c)), uintptr(unsafe.Pointer(&d)))
	return *(*string)(unsafe.Pointer(sum))
}

func init() {
	fn := syscall.NewLazyDLL("format/format.dll")
	FormatFunc = fn.NewProc("format")
}

func main() {
	fmt.Println(Format(1, 2.0, 3.0, "4"))
}
