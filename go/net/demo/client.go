package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

var Stdin = bufio.NewReader(os.Stdin)

func Scanln(str *string) (int, error) {
	s, _, err := Stdin.ReadLine()
	*str = string(s)
	return len(s), err
}

func OnReceived(conn *net.TCPConn) {
	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(time.Second * 20)
	r := bufio.NewReader(conn)
	var s string
	for {
		Scanln(&s)
		_, err := conn.Write([]byte(s + "\n"))
		if err != nil {
			panic("连接断开")
		}
		s, err = r.ReadString('\n')
		if err != nil {
			panic("连接断开")
		}
		fmt.Println("收到消息：", s)
	}
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("连接成功")
	}
	defer conn.Close()
	OnReceived(conn)
}
