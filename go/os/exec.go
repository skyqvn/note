package main

import (
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	cmd := exec.Command("cmd")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	time.Sleep(1 * time.Second)
	cmd.Process.Signal(syscall.SIGKILL)
}
