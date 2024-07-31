package main

import (
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/process"
)

func main() {
	// 全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	// 创建应用
	app := cef.NewApplication()
	// 指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		println("simple 这里只执行一次")
	})
	println("IsMain:", process.Args.IsMain()) // 主进程: true
	Init2()                                   // main 函数手动调用 被执行多次
	// 运行应用
	cef.Run(app)
}

// 由于CEF是多进程, 该函数会被执行多次
func init() {
	// Go在初始化时执行
	println("simple init")
}

func Init2() {
	println("simple Init2")
}
