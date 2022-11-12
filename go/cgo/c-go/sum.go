package main

import "C"
import "fmt"

//export printNum
func printNum(i int) {
	fmt.Println(i)
}

//export printString
func printString(s string) {
	fmt.Println(s)
}

func main() {

}
