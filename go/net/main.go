package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//根据域名返回ip地址
	addr, err := net.ResolveIPAddr("ip", "devops.miliantech.com")
	if err != nil {
		return
	}
	log.Println("devops.miliantech.com 对应地址ip地址---->", addr)

	// 查找DNS A记录
	iprecords, _ := net.LookupIP("mk.miliantech.com")
	for _, ip := range iprecords {
		fmt.Println("LookupIP ----->", ip)
	}
	// 查找DNS CNAME记录
	canme, _ := net.LookupCNAME("devops.miliantech.com")
	fmt.Println("LookupCNAME ----->", canme)

	// 查找DNS PTR记录
	ptr, e := net.LookupAddr("101,43,124.202")
	if e != nil {
		fmt.Println(e)
	}
	for _, ptrval := range ptr {
		fmt.Println(ptrval)
	}
	// 查找DNS NS记录 名字服务器记录
	nameserver, _ := net.LookupNS("baidu.com")
	for _, ns := range nameserver {
		fmt.Println("ns记录", ns)
	}
	// 查找DNS MX记录 邮件服务器记录
	mxrecods, _ := net.LookupMX("google.com")
	for _, mx := range mxrecods {
		fmt.Println("mx:", mx)
	}
	// 查找DNS TXT记录 域名对应的文本信息
	txtrecords, _ := net.LookupTXT("baidu.com")

	for _, txt := range txtrecords {
		fmt.Println("txt:", txt)
	}

	// 查看本地IP地址
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("本地地址 -- >", ipnet.IP.String())
			}
		}
	}
}
