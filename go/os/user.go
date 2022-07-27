package main

import (
	"fmt"
	"os/user"
)

func main() {
	u, _ := user.Current()
	fmt.Println("用户名：", u.Username)
	fmt.Println("用户id", u.Uid)
	fmt.Println("用户主目录：", u.HomeDir)
	fmt.Println("主组id：", u.Gid)
	
	// 用户所在的所有的组的id
	s, _ := u.GroupIds()
	fmt.Println("用户所在的所有组：", s)
}
