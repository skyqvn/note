package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	// Hostname 函数会返回内核提供的主机名。
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println(hostname)
	
	// Environ 函数会返回所有的环境变量，返回值格式为“key=value”的字符串的切片拷贝。
	env := os.Environ()
	fmt.Println(env)
	
	// Getenv 函数会检索并返回名为 key 的环境变量的值，如果不存在该环境变量则会返回空字符串。
	envPath := os.Getenv("path")
	fmt.Println(envPath)
	
	// Setenv 函数可以设置名为 key 的环境变量，如果出错会返回该错误。
	err = os.Setenv("GO111MODULE", "on")
	if err != nil {
		panic(err)
	}
	
	// Getuid 函数可以返回调用者的用户 ID。
	userid := os.Getuid()
	
	// Getgid 函数可以返回调用者的组 ID。
	groupid := os.Getgid()
	
	// Getpid 函数可以返回调用者所在进程的进程 ID。
	pidid := os.Getpid()
	
	fmt.Println(userid, groupid, pidid)
	
	// Getwd 函数可以返回一个对应当前工作目录的根路径。
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(workPath)
	
	// Chdir 将工作目录修改为dir
	err = os.Chdir("../")
	if err != nil {
		panic(err)
	}
	
	// Mkdir 函数可以使用指定的权限和名称创建一个目录。
	err = os.Mkdir("filePath", os.ModePerm)
	if err != nil {
		panic(err)
	}
	
	// MkdirAll 函数可以使用指定的权限和名称创建一个目录，包括任何必要的上级目录。
	err = os.MkdirAll("filePath/go", os.ModePerm)
	if err != nil {
		panic(err)
	}
	
	// Remove 函数会删除 name 指定的文件或目录。
	err = os.Remove("filePath/go")
	if err != nil {
		panic(err)
	}
	
	// RemoveAll 函数会删除 name 指定的文件或目录，并递归的删除所有子目录和文件。
	err = os.RemoveAll("filePath")
	if err != nil {
		panic(err)
	}
	
	// 文件
	
	// 创建文件
	newfile, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	newfile.Close()
	
	// 打开文件
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	file.Close()
	
	// 以指定模式和权限打开文件
	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	file.Close()
	
	// 文件描述符
	fmt.Println(file.Fd())
	
	// 文件权限
	// os.O_RDONLY 只读
	// os.O_WRONLY 只写
	// os.O_RDWR 读写
	// os.O_APPEND 往文件中添建（Append）
	// os.O_CREATE 如果文件不存在则先创建
	// os.O_TRUNC 文件打开时裁剪文件
	// os.O_EXCL 和O_CREATE一起使用，文件不能存在
	// os.O_SYNC 以同步I/O的方式打开
	
	// 读取文件
	file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	var fb = make([]byte, 256)
	for {
		byteReader, err := file.Read(fb)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println(byteReader, ": ", fb[:byteReader])
	}
	
	// 写入文件
	file, err = os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	byteWriter, err := file.Write([]byte{'h', 'e', 'l', 'l', 'o'})
	if err != nil {
		panic(err)
	}
	fmt.Println(byteWriter)
	
	byteWriter, err = file.WriteAt([]byte{'h', 'e', 'l', 'l', 'o'}, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(byteWriter)
	
	byteWriter, err = file.WriteString("hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(byteWriter)
	
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	
	data, err = os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	
	// 重命名和移动
	err = os.Rename("filePath", "newpath")
	if err != nil {
		panic(err)
	}
	
	// 裁剪文件,不足的字节以null字节填充,超过的字节被抛弃。
	err = os.Truncate("test.txt", 100)
	if err != nil {
		panic(err)
	}
	
	// 文件详情
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
	
	// 判断文件类型
	fmt.Println(fileStat.Mode()&fs.ModeDir == fs.ModeDir)
	fmt.Println(fileStat.Mode()&fs.ModeSymlink == fs.ModeSymlink)
	fmt.Println(fileStat.Mode()&fs.ModeSocket == fs.ModeSocket)
	// ……
	
	// 查看文件是否存在
	fileStat, err = os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			panic(fmt.Errorf("无法找到文件：test.txt"))
		}
	}
	
	// 查看文件是否有权限
	file, err = os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			panic(fmt.Errorf("无权限打开文件：test.txt"))
		}
	}
	file.Close()
	
	// 改变文件权限
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
	
	// 复制文件
	file, err = os.Create("test.txt")
	cpfile, err := os.Create("cptest.txt")
	byteWriter2, err := io.Copy(cpfile, file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("拷贝了%d个字节", byteWriter2)
	file.Close()
	
	// 跳转文件指定位置
	file, err = os.OpenFile("test.txt", os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	// whence的值，用来计算offset的初始位置
	// 0 = 文件开始位置
	// 1 = 当前位置
	// 2 = 文件结尾处
	position, err := file.Seek(8, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(position)
	file.Close()
	
	// 判断是否是绝对路径
	isAbs := path.IsAbs("hello.go")
	fmt.Println(isAbs)
	
	// 将路径分割为路径和文件名
	fileDir, fileName := path.Split("/xxx/hello.go")
	fmt.Println(fileDir, fileName)
	
	// 获取文件后缀名
	ext := path.Ext("hello.go")
	fmt.Println(ext)
	
	// 将多个路径合并
	filePath := path.Join("/xxx", "hello.go")
	fmt.Println(filePath)
	
	// 返回的最后一个元素
	base := path.Base("/xxx/hello.go")
	fmt.Println(base)
	
	// 返回路径所在目录
	fileDir = path.Dir("/xxx/hello.go")
	fmt.Println(fileDir)
	
	// 返回同目录的最短路径
	fileDir = path.Clean("/xxx/hello.go")
	
	// 创建硬链接
	os.Link("/xxx/hello.go", "./hello.go")
	
	// 创建软链接
	os.Symlink("/xxx/hello.go", "./hello.go")
	
	// 获取软链接真实位置
	s, _ := filepath.EvalSymlinks("/bin")
	fmt.Println(s)
	
	// Exit 函数可以让当前程序以给出的状态码 code 退出。
	os.Exit(0)
}
