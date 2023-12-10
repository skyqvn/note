package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"math"
	"syscall"
	"unsafe"
)

import "C"

var FormatFunc *syscall.LazyProc

func Format(a int, b float32, c float64, d string) string {
	s, _ := syscall.UTF16PtrFromString(d)
	sum, _, _ := FormatFunc.Call(uintptr(a), uintptr(math.Float32bits(b)), uintptr(math.Float64bits(c)), uintptr(unsafe.Pointer(s)))
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(sum)))
}

func init() {
	fn := syscall.NewLazyDLL("format/format.dll")
	FormatFunc = fn.NewProc("format")
}

func main() {
	fmt.Println(Format(1, 2.0, 3.0, "4"))
}
