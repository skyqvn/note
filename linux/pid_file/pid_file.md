## PID文件

[教程1](https://wenku.baidu.com/view/5c9c8415598102d276a20029bd64783e08127d71.html?_wkts_=1668247217881&bdQuery=pid%E6%96%87%E4%BB%B6)  
[教程2](https://segmentfault.com/a/1190000019036070)  
[Golang操作PID教程1](https://blog.csdn.net/villare/article/details/85453489)  
[Golang操作PID教程2](https://www.bbsmax.com/A/E35pRKYgzv/)

PID文件作用：  
防止进程启动多个副本。只有获得pid文件(固定路径固定文件名)写入权限(F_WRLCK)的进程才能正常启动并把自身的PID写入该文件中。其它同一个程序的多余进程则自动退出。

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// PIDFile stored the process id.
type PIDFile struct {
	path string
}

// Check if the process exists.
func ProcessExists(pid string) bool {
	if _, err := os.Stat(filepath.Join("/proc", pid)); err == nil {
		return true
	}
	return false
}

// Check if the PID file already exists.
func CheckPidFileAlreadyExists(path string) error {
	if pidByte, err := os.ReadFile(path); err == nil {
		pid := strings.TrimSpace(string(pidByte))
		if ProcessExists(pid) {
			return fmt.Errorf("ensure the process:%s is not running pid file:%s", pid, path)
		}
	}
	return nil
}

// Create a PID file in the specified path. The file content is the PID of this program.
func NewPIDFile(path string) (*PIDFile, error) {
	if err := CheckPidFileAlreadyExists(path); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(filepath.Dir(path), os.FileMode(0755)); err != nil {
		return nil, err
	}
	if err := os.WriteFile(path, []byte(fmt.Sprintf("%d", os.Getpid())), 0644); err != nil {
		return nil, err
	}
	return &PIDFile{path: path}, nil
}

// Remove the PID file
func (file PIDFile) Remove() error {
	return os.Remove(file.path)
}

```
