package main

import "C"

import (
	"fmt"
	"golang.org/x/sys/windows"
	"math"
	"syscall"
	"unsafe"
)

var FormatDLL *syscall.LazyDLL
var FormatFunc *syscall.LazyProc

func init() {
	FormatDLL = syscall.NewLazyDLL("format/format.dll")
	FormatFunc = FormatDLL.NewProc("format")
}

func Format(a int, b float32, c float64, d string) string {
	s, _ := syscall.UTF16PtrFromString(d)
	// NewCallback将Go函数转换为符合stdcall调用约定的函数指针。
	// 这在与需要回调的Windows代码进行互操作时非常有用。
	// 返回值应该是一个uintptr大小的值。
	// 该函数的每个参数不能大于uintptr类型的大小。
	// 在单个Go进程中只能创建有限数量的回调，并且为这些回调分配的任何内存永远不会释放。
	// 在NewCallback和NewCallbackCDecl之间，至少可以创建1024个回调。
	// syscall.NewCallback()

	// NewCallbackCDecl将Go函数转换为符合cdecl调用约定的函数指针。
	// 其他相同
	// syscall.NewCallbackCDecl()
	e := syscall.NewCallback(func(state int) uintptr {
		fmt.Printf("状态码：%d\n", state)
		return 0
	})

	result, _, _ := FormatFunc.Call(uintptr(a), uintptr(math.Float32bits(b)) /* GO的float格式就是IEEE 754二进制，所以也可以直接强制地址转换 */, uintptr(math.Float64bits(c)), uintptr(unsafe.Pointer(s)), e)
	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(result)))
}

func main() {
	fmt.Println(Format(1, 2.0, 3.0, "4"))
}
