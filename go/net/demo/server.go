package main

import (
	"bufio"
	"fmt"
	"net"
)

func Pipe(conn *net.TCPConn) {
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
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("服务运行在", listener.Addr().String())

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}
		go Pipe(conn)
	}
}
