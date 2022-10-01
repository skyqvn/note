package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "    hello   "
	strings.TrimSpace(str)
	fmt.Println(str)
}
