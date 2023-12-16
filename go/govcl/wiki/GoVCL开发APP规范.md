<待继续补充...>  
----

主要是提供一个规范的写法，samples目录中的后期的大多数都是按照规范去写的，只有一些较早的没有规范化。  

#### 前言

* 1、作为一个UI程序，逻辑上`main`包为主，其他包为辅的方式去编写代码。不要想着在其他包中直接调用main包中的东西。
* 2、UI必须在主线程中运行，不能使用任何其他方式将UI的初始和创建放在其他协程中运行。换句话说，所有UI的初始和创建都只能在main函数中完成，并且不能用协程去操作这块。
* 3、在控件事件上应尽量避免使用匿名函数，不要贪图方便，应尽量使用类方法的形式。
* 4、TForm的方法函数，self变量命名不应该与全局命名相同，应该尽量使用简短的命名，比如：`(f *TForm1)`。
* 5、当其他包想要传递数据给UI时，应该在其他包中导出相关接口，然后在main包中的相应位置初始他们(在项目例子的rproxy中就有相关的)。
  ```go
     // 比如
     // pkg1 包
     package pkg1

     var (
        Callback1 func(msg string)
     )

     func xxx() {
        if Callback1 == nil {
           return 
        }
        for {
           Callback1("message")
        }
     }

     // -----------------------------------------------
     // main 包
     package main

     import "pkg1"


     func main() {
         // 这里只是说明方法。
         // 在合适的位置初始化这个pkg1.Callback1
         pkg1.Callback1 = xxx            
     }

     func xxx(msg string) {
        // 这里接收数据，根据实际情况传递到某处。
     }
  ```

#### 规范  

注：**main.go一般采用自动生成的即可，无特殊要求不要在里面写其它代码。一些初始化的代码都应该放在相应的TForm初始化事件OnFormCreate中。整个main.go保持简洁少量的代码。**  

#### 一、main.go

> 如（一般无特殊要求整个main.go文件内容就如下代码）： 

```go
package main

import (
    "github.com/ying32/govcl/vcl"
    _ "github.com/ying32/govcl/pkgs/winappres"
)

func main() {
    vcl.Application.Initialize()
    vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateForm(&Form1)
    vcl.Application.CreateForm(&Form2)
    vcl.Application.Run()
}
```

或者：
```go
package main

import (
    "github.com/ying32/govcl/vcl"
    _ "github.com/ying32/govcl/pkgs/winappres"
)

func main() {
    vcl.RunApp(&Form1, &Form2)
}
```

再或者：
```go
package main

import (
    "github.com/ying32/govcl/vcl"
    _ "github.com/ying32/govcl/pkgs/winappres"
)

func main() {
    vcl.RunApp(&Form1, func(){ 
       // do something... 
    }, &Form2, func(){ 
       // do something... 
    })
}
```

#### 二、TForm对应的impl文件   

> impl(impl指的implementation，即实现部分代码)后缀的文件名一般是首次自动生成的Form事件操作或者自定义的内容。   
> 如需要在TForm初始后初始一些自定义的变量或者方法之类的。   
> `TForm1Fields` 一般是对应的私有变量放置处，名字是与当前类对应的，命名规则为`类名+Fields`。   

```go
package main

import "github.com/ying32/govcl/vcl"

//::private::
type TForm1Fields struct {
    // 在这里定义你的变量
    isInit bool
    title string
}


func (f *TForm1) OnFormCreate(sender vcl.IObject) {
   // 一般在这里初始化自己的东东
   f.isInit = false
   f.title = "我是一个标题"

   f.SetCaption(f.title)
}

func (f *TForm1) OnFormClick(sender vcl.IObject) {
    
   f.isInit = true
   f.title = "222222"

   f.SetCaption(f.title)
}

```