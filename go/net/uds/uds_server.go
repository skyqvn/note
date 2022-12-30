package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const Addr = "/tmp/go_uds.sock"

func Pipe(conn *net.UnixConn) {
	cliIP := conn.RemoteAddr().String()
	fmt.Println("新连接:", cliIP)
	r := bufio.NewReader(conn)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("连接 " + cliIP + " 断开")
			break
		}
		fmt.Printf("收到%s消息：%s", cliIP, s)
		conn.Write([]byte("OK\n"))
	}
}

func main() {
	os.Remove(Addr)
	addr, _ := net.ResolveUnixAddr("unix", Addr)
	listener, err := net.ListenUnix("unix", addr)
	if err != nil {
		panic(err)
	}
	defer func() {
		listener.Close()
		err := os.Remove("/tmp/go_uds.sock")
		fmt.Println(err)
	}()
	fmt.Println("服务运行在", listener.Addr().String())
	q := make(chan os.Signal)
	signal.Notify(q, syscall.SIGINT)
	go func() {
		fmt.Printf("收到退出信号，信号%s", (<-q).String())
		os.Exit(0)
	}()

	for {
		conn, err := listener.AcceptUnix()
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}
		go Pipe(conn)
	}
}
