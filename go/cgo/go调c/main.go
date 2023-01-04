package main

//#cgo CXXFLAGS: -std=c++14
//#include "sum.h"
import "C"
import "fmt"

func main() {
	s := C.s()
	fmt.Println(s)
}
