package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	//Hostname 函数会返回内核提供的主机名。
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println(hostname)
	
	//Environ 函数会返回所有的环境变量，返回值格式为“key=value”的字符串的切片拷贝。
	env := os.Environ()
	fmt.Println(env)
	
	//Getenv 函数会检索并返回名为 key 的环境变量的值，如果不存在该环境变量则会返回空字符串。
	path := os.Getenv("path")
	fmt.Println(path)
	
	//Setenv 函数可以设置名为 key 的环境变量，如果出错会返回该错误。
	err = os.Setenv("GO111MODULE", "on")
	if err != nil {
		panic(err)
	}
	
	//Getuid 函数可以返回调用者的用户 ID。
	userid := os.Getuid()
	
	//Getgid 函数可以返回调用者的组 ID。
	groupid := os.Getgid()
	
	//Getpid 函数可以返回调用者所在进程的进程 ID。
	pidid := os.Getpid()
	
	fmt.Println(userid, groupid, pidid)
	
	//Getwd 函数可以返回一个对应当前工作目录的根路径。
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(workPath)
	
	//Chdir 将工作目录修改为dir
	err = os.Chdir("../")
	if err != nil {
		panic(err)
	}
	
	//Mkdir 函数可以使用指定的权限和名称创建一个目录。
	err = os.Mkdir("path", os.ModePerm)
	if err != nil {
		panic(err)
	}
	
	//MkdirAll 函数可以使用指定的权限和名称创建一个目录，包括任何必要的上级目录。
	err = os.MkdirAll("path/go", os.ModePerm)
	if err != nil {
		panic(err)
	}
	
	//Remove 函数会删除 name 指定的文件或目录。
	err = os.Remove("path/go")
	if err != nil {
		panic(err)
	}
	
	//RemoveAll 函数会删除 name 指定的文件或目录，并递归的删除所有子目录和文件。
	err = os.RemoveAll("path")
	if err != nil {
		panic(err)
	}
	
	//文件
	
	//创建文件
	newfile, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	newfile.Close()
	
	//打开文件
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	file.Close()
	
	//以指定模式和权限打开文件
	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	file.Close()
	
	//文件权限
	// os.O_RDONLY 只读
	// os.O_WRONLY 只写
	// os.O_RDWR 读写
	// os.O_APPEND 往文件中添建（Append）
	// os.O_CREATE 如果文件不存在则先创建
	// os.O_TRUNC 文件打开时裁剪文件
	// os.O_EXCL 和O_CREATE一起使用，文件不能存在
	// os.O_SYNC 以同步I/O的方式打开
	
	//重命名和移动
	err = os.Rename("path", "newpath")
	if err != nil {
		panic(err)
	}
	
	//裁剪文件,不足的字节以null字节填充,超过的字节被抛弃。
	err = os.Truncate("test.txt", 100)
	if err != nil {
		panic(err)
	}
	
	//文件详情
	fileStat, err := os.Stat("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File name:", fileStat.Name())
	fmt.Println("Size in bytes:", fileStat.Size())
	fmt.Println("Permissions:", fileStat.Mode())
	fmt.Println("Last modified:", fileStat.ModTime())
	fmt.Println("Is Directory: ", fileStat.IsDir())
	fmt.Printf("System interface type: %T\n", fileStat.Sys())
	fmt.Printf("System info: %+v\n", fileStat.Sys())
	
	//查看文件是否存在
	fileStat, err = os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			panic(fmt.Errorf("无法找到文件：test.txt"))
		}
	}
	
	//查看文件是否有权限
	file, err = os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			panic(fmt.Errorf("无权限打开文件：test.txt"))
		}
	}
	
	//改变文件权限
	err = os.Chmod("test.txt", 0777)
	if err != nil {
		panic(err)
	}
	
	// 改变文件所有者
	err = os.Chown("test.txt", os.Getuid(), os.Getgid())
	if err != nil {
		panic(err)
	}
	
	// 改变时间戳
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		panic(err)
	}
	
	//复制文件
	file, err = os.Create("test.txt")
	cpfile, err := os.Create("cptest.txt")
	byteWriter, err := io.Copy(cpfile, file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("拷贝了%d个字节", byteWriter)
	file.Close()
	
	//跳转文件指定位置
	file, err = os.OpenFile("test.txt", os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	// 用来计算offset的初始位置
	// 0 = 文件开始位置
	// 1 = 当前位置
	// 2 = 文件结尾处
	position, err := file.Seek(1, 5)
	if err != nil {
		panic(err)
	}
	fmt.Println(position)
	
	//Exit 函数可以让当前程序以给出的状态码 code 退出。
	os.Exit(0)
}
