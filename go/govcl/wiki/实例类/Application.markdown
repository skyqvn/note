Lazarus中管理窗口程序的类实例。Application由Lazarus在单元initialization和finalization时自动构造和析构。  

常用方法：



* `Initialize`  
初始始应用程序相关设置，必要写的
```go
vcl.Application.Initialize()
```


* `Run`  
开始整个程序的消息循环,直到收到WM_QUIT或者主窗口收到WM_CLOSE退出循环，并结束程序
```govcl
vcl.Application.Run()
```

* `RunApp`  
一个简化的函数。
```go
vcl.RunApp(&form1, &from2)
```


* `CreateForm`  
创建并返回一个新的TGoForm实例
```go
//TApplication.CreateForm 一般不建议使用NewForm，而优先使用CreateForm或者使用对应的ＮewXXX()，如 NewForm1()

//-------------------- 用法一--------------------------------------
// 直接返回，不推荐使用
// mainForm := vcl.Application.CreateForm()    
mainForm := vcl.Application.CreateForm()
mainForm.SetOnClick(func(sender vcl.IObject) {
    vcl.ShowMessage("msg")
})


//-------------------- 用法二--------------------------------------
// 无资源加载，只会绑定窗口的事件，不会绑定子组件事件
//vcl.Application.CreateForm(&mainForm)       
type TMainForm struct {
    *vcl.TForm
}

var mainForm *TMainForm

vcl.Application.CreateForm(&mainForm)

func (f *TMainForm)OnFormCreate(sender vcl.IObject) {
    fmt.Println("FormCreate")
}

func (f *TMainForm)OnFormClick(sender vcl.IObject) {
    vcl.ShowMessage("click")
}


//-------------------- 用法三--------------------------------------
// 无资源加载时在OnFormCreate
//vcl.Application.CreateForm(&mainForm) 
type TMainForm struct {
    *vcl.TForm
    Btn1 *vcl.TButton
}

var mainForm *TMainForm
vcl.Application.CreateForm(&mainForm)

func (f *TMainForm)OnFormCreate(sender vcl.IObject) {
    fmt.Println("FormCreate")
    f.Btn1 = vcl.NewButton(f)
    f.Btn1.SetParent(f)
    f.Btn1.SetOnClick(f.OnBtn1Click)
}

func (f *TMainForm)OnFormClick(sender vcl.IObject) {
    vcl.ShowMessage("click")
}

func (f *TMainForm)OnBtn1Click(Sender vcl.IObject) {
    vcl.ShowMessage("Btn1 Click")
}


//-------------------- 用法四--------------------------------------
// 从资源文件中填充子组件，并绑定所有事件
// 不推荐使用了

//-------------------- 用法五--------------------------------------
// 从字节中填充子组件，并绑定所有事件
// 不推荐使用了 
```

* `MainFormOnTaskbar`、`SetMainFormOnTaskbar`  
获取或者设置主窗口是否显示在任务栏上，一般默认写True
```go
fmt.Println(vcl.Application.MainFormOnTaskbar())
vcl.Application.SetMainFormOnTaskbar(true)
```


* `ProcessMessages`  
处理消息，当某主线程中发生阻止时会造成整个消息循环阻止，UI无响应，所以在阻止中，比如循环中加上ProcessMessages可解决UI无响应问题。
```go
vcl.Application.ProcessMessages()
```


* `Terminate`  
结束应用程序，发出WM_QUIT消息
```go
vcl.Application.Terminate()
```


* `ExeName`  
获取当前运行的应用程序文件名，含路径
```go
fmt.Println(vcl.Application.ExeName())
```


* `Icon`, `SetIcon`  
获取或者设置应用程序的图标，一般默认会在程序资源中找名为MAINICON的ico资源作为应用程序图标。
```go
// 手动指定
ico := vcl.NewIcon()
defer ico.Free()
// LoadFromResourceID(rtl.MainInstance(), 3) // 3为资源中的id，具体要根据自己生成的.syso文件中定义的id
ico.LoadFromFile("xxx.ico")
vcl.Application.SetIcon(ico)
// 或者
vcl.Application.Icon().SetHandle(win.LoadIcon(rtl.MainInstance(), 3))
// 又或
vcl.Application.Icon().SetHandle(win.LoadIcon2(rtl.MainInstance(), "MAINICON"))
// 双或
vcl.Application.Icon().LoadFromFile("ico.ico")
```


* `Title`, `SetTitle`  
获取或者设置应用程序标题
```go
fmt.Println(vcl.Application.Title())
vcl.Application.SetTitle("我是标题")
```

* `SetScaled` ,`暂时废弃`  
有关dpi的绽放 
```go
vcl.Application.SetScaled(true)
```

* `SetIconResId`,`不推荐使用`    
设置App全局icon。从资源中加载，具体看rsrc后的id， 仅windows
```go
vcl.Application.SetIconResId(3)
```

* `ScaleForCurrentDpi`   
```go
//用于使用纯代码创建的窗口缩放，所有组件创建完后调用 
F.ScaleForCurrentDpi()
```