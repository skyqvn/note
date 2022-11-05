package main

import (
	_"unsafe"
	_"hello"
)

//go:linkname sayhello hello.hello
func sayhello()

func main(){
	sayhello()
}

