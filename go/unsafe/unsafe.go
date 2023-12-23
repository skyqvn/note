package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//Example1:
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
	
	//Example2:
	const len = 11
	var data = [len]byte{'h', 'e', 'l', 'l', 'o',' ','w','o','r','l','d'}
	//unsafe.String() 根据指针和长度创造String.注意，String不应被更改.
	str:=unsafe.String(&data[0],len)
	fmt.Println(str)
	
	//unsafe.Slice() 根据指针和长度创作Slice
	sli:=unsafe.Slice(&data[0],len)
	fmt.Println(sli)
	
	strPtr:=unsafe.StringData(str)
	fmt.Println(strPtr)
	
	sliPtr:=unsafe.SliceData(sli)
	fmt.Println(sliPtr)
}
