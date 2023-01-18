package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s string
	var i int
	var i64 int64
	var pi64 *int64
	var a2i [2]int
	type t struct {
		A int32
		B int
	}
	var st t
	
	fmt.Println(unsafe.Sizeof(i64))
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(a2i))
	
	fmt.Println(unsafe.Offsetof(st.A))
	fmt.Println(unsafe.Offsetof(st.B))
	
	pi64 = (*int64)(unsafe.Pointer(&i))
	*pi64 = i64
	
	var v t
	ip := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&v)) + uintptr(8)))
	*ip = 20
	fmt.Println(v.B)
}
