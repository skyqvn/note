package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	fn := syscall.NewLazyDLL("add/add.dll")
	add := fn.NewProc("add")
	outRand := make([]byte, 2)
	a := 1
	b := 2
	sum, _, _ := add.Call(uintptr(a), uintptr(b), uintptr(unsafe.Pointer(&outRand[0])))
	fmt.Println(sum)
}
