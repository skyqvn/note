package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const Addr = "/tmp/go_uds.sock"

var Stdin = bufio.NewReader(os.Stdin)

func Scanln(str *string) (int, error) {
	s, _, err := Stdin.ReadLine()
	*str = string(s)
	return len(s), err
}

func OnReceived(conn *net.UnixConn) {
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
	//unix:SOCK_STREAM，提供面向连接的稳定数据传输，即TCP协议。
	//unixgram:SOCK_DGRAM，使用不连续不可靠的数据包连接，即UDP协议。
	//unixpacket:SOCK_SEQPACKET，提供连续可靠的数据包连接。
	addr, _ := net.ResolveUnixAddr("unix", Addr)
	conn, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("连接成功")
	}
	defer conn.Close()
	OnReceived(conn)
}
