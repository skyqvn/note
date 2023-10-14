package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	var cmd *exec.Cmd
	cmd = exec.Command("bash")
	//or
	cmd = &exec.Cmd{
		Path: "bash",
		Args: []string{"bash", "-c"},//注意，应该把path内容放在切片第一位
		Dir:  "/usr/bin",
		Env: []string{
			"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
		},
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	fmt.Println(cmd.Environ())
	time.Sleep(1 * time.Second)
	cmd.Process.Signal(syscall.SIGKILL)
	
	c1 := exec.Command("grep", "ERROR", "/var/log/messages")
	c2 := exec.Command("wc", "-l")
	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = os.Stdout
	_ = c2.Start()
	_ = c1.Run()
	_ = c2.Wait()
}
