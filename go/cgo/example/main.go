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

func init() {
	fn := syscall.NewLazyDLL("format/format.dll")
	FormatFunc = fn.NewProc("format")
}

func Format(a int, b float32, c float64, d string) string {
	s, _ := syscall.UTF16PtrFromString(d)
	//NewCallback将Go函数转换为符合stdcall调用约定的函数指针。
	//这在与需要回调的Windows代码进行互操作时非常有用。
	//参数应该是一个uintptrsize结果的函数。
	//该函数的参数不能大于uintptr类型的大小。
	//在单个Go进程中只能创建有限数量的回调，并且为这些回调分配的任何内存永远不会释放。
	//在NewCallback和NewCallbackCDecl之间，至少可以创建1024个回调。
	//syscall.NewCallback()

	//NewCallbackCDecl将Go函数转换为符合cdecl调用约定的函数指针。
	//其他相同
	//syscall.NewCallbackCDecl()

	result, _, _ := FormatFunc.Call(uintptr(a), uintptr(math.Float32bits(b)), uintptr(math.Float64bits(c)), uintptr(unsafe.Pointer(s)))
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(result)))
}

func main() {
	fmt.Println(Format(1, 2.0, 3.0, "4"))
}
