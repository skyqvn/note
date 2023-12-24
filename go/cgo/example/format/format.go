//go:generate go build -buildmode=c-shared -o format.dll
package main

/*
#include "run_callback.h"
*/
import "C"
import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

//export format
func format(a int, b float32, c float64 /* GO的float格式就是IEEE 754二进制，不必再转换 */, d *uint16, e uintptr /* callback */) uintptr {
	ds := windows.UTF16PtrToString(d)
	s, _ := syscall.UTF16PtrFromString(fmt.Sprintf("int:%d\nfloat32:%G\nfloat64:%G\nstring:%s", a, b, c, ds))
	C.run_callback(unsafe.Pointer(e), 0) //调用回调
	return uintptr(unsafe.Pointer(s))
}

func main() {}
