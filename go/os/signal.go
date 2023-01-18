package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal)
	//signal.Stop(c)
	signal.Notify(c)
	fmt.Println(<-c)
	fmt.Println("quit")
}
