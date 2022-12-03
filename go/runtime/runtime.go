package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

func main() {
	numcpu := runtime.GOMAXPROCS(4) //设置CPU核数，必须放在main的开头
	fmt.Println(numcpu)

	fmt.Println(runtime.NumCPU())       //CPU数量
	fmt.Println(runtime.NumGoroutine()) //goroutine数量
	fmt.Println(runtime.NumCgoCall())   //返回当前进程的cgo调用数

	fmt.Println(runtime.GOOS)     //系统
	fmt.Println(runtime.GOARCH)   //CPU架构
	fmt.Println(runtime.GOROOT()) //获取GOROOT

	//LockOSThread
	go func() {
		//获取系统线程锁，让此线程上只运行此协程
		//如果LockOSThread被多次调用，UnlockOSThread必须调用相同的次数才能解锁线程。
		runtime.LockOSThread()
		time.Sleep(time.Second)
		//释放系统线程锁
		runtime.UnlockOSThread()
	}()

	//Gosched
	go func() {
		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println("A")
			}
		}()
		for i := 0; i < 10; i++ {
			runtime.Gosched()
			fmt.Println("B")
		}
	}()

	//Goexit
	go func() {
		defer func() {
			fmt.Println("E")
		}()
		fmt.Println("A")
		goExit()
		fmt.Println("B")
		//Output:
		//A
		//C
		//E
	}()

	//KeepAlive
	var p1 unsafe.Pointer
	{
		a := 100
		p1 = unsafe.Pointer(&a)
		runtime.KeepAlive(a)
	}
	fmt.Println(*(*int)(p1))

	//SetFinalizer
	p2 := new(int)
	runtime.SetFinalizer(p2, func(n *int) {
		fmt.Println(*n)
	})
}

func goExit() {
	fmt.Println("C")
	runtime.Goexit()
	fmt.Println("D")
}
