//go:generate go build -buildmode=c-shared -o format.dll
package main

import "C"
import (
	"fmt"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

//export format
func format(a int, b float32, c float64, d *uint16) uintptr {
	ds := windows.UTF16PtrToString(d)
	s, _ := syscall.UTF16PtrFromString(fmt.Sprintf("int:%d\nfloat32:%G\nfloat64:%G\nstring:%s", a, b, c, ds))
	return uintptr(unsafe.Pointer(s))
}

func main() {}
