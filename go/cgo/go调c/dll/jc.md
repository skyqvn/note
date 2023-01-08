# 动态链接库

[教程](https://blog.csdn.net/quicmous/article/details/80350484)

### 调用的两种方法

```go
package main

import (
	"fmt"
	"syscall"
)

//https://blog.csdn.net/quicmous/article/details/80350484
func main() {
	//方案1
	d, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer syscall.FreeLibrary(d)
	p, err := syscall.GetProcAddress(d, "GetVersion")
	if err != nil {
		fmt.Println(err)
		return
	}
	r, _, _ := syscall.SyscallN(p)
	printVersion(uint32(r))

	//方案2
	Compress := syscall.NewLazyDLL("kernel32.dll")
	Comp := Compress.NewProc("GetVersion")
	ret, _, err := Comp.Call()
	printVersion(uint32(ret))
}

func printVersion(v uint32) {
	major := uint8(v)
	minor := uint8(v >> 8)
	build := uint16(v >> 16)
	fmt.Printf("Version:%d.%d(Build:%d)\n", major, minor, build)
}

```
