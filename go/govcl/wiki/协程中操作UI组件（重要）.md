Lazarus中的所有`可视的UI组件`都**非**`线程`或`协程`安全的。

所以在协程中操作UI组件都`必须`使用`vcl.ThreadSync`来同步到主线程（UI线程）中执行。否则会随机产生`AV`错误或者其他莫名的错误。   

### 正确的操作方式

* 示例一  

```go
  go func() {
     vcl.ThreadSync(func(){
        f.Button.SetCaption("Hello.")
        fmt.Println(f.Button.Caption())
        vcl.ShowMessage("11111")
     })
  }()
```

* 示例二 

```go
go func() {
    // 协程中进行http请求，
    resp, err := http.Get("https://github.com")
    if err != nil {
       return
    }
    defer resp.Body.Close()
    bs, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return 
    }
    // 整个操作完成后，显示到UI上时使用vcl.ThreadSync切换到主线程中执行。
    vcl.ThreadSync(func(){
       f.Memo1.SetText(string(bs))
    })
}()
```

### ~~错误的操作方式~~  

* ~~示例一~~ 

```go
  go func() {
        // 这样不定时的造成AV错误，也会内存泄漏
        f.Button.SetCaption("Hello.")
        fmt.Println(f.Button.Caption())
  }()
```

* ~~示例二~~

```go
go func() {
    resp, err := http.Get("https://github.com")
    if err != nil {
       return
    }
    defer resp.Body.Close()
    bs, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return 
    }
    // 这样就会造成AV错误，而且不定时的，也会内存泄漏
    f.Memo1.SetText(string(bs))
}()
```

注：AV错误指的是
![输入图片说明](https://images.gitee.com/uploads/images/2020/0513/142434_05009bb1_118989.png "屏幕截图.png")