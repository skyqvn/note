### FPC/Delphi都为**手动管理内存**，所以习惯于自动管理内存的朋友请特别留意。每个类都带有Free方法，用于释放Create(Go中已经将Create变为以New开头的函数了，比如：NewButton)所分配的内存。这里举几个例子进行下说明：  

* 1、继承自`TComponent`的, 当`Owner不为nil`时，其`Owner`会在销毁时同时销毁它。  

```go
// 指定Owner, Owner可为自己的Parent或者Root，
// 一般Delphi/Lazarus从dfm/lfm中构建UI时指定的为Root，即TForm实例。
// button的Owner指定了mainForm，
// 那么此时在mainForm销毁前会自动调用button.Free()
button := vcl.NewButton(mainForm)

// Application.CreateForm 因为CreateForm方法为内部所封装，
// Owner已经自动指定了TApplication的实例，所不需要手动调用Free方法, 
// 除非需求中提前销毁某个窗口。
mainForm := Application.CreateForm()
```

* 2、继承自`TComponent`的, 当`Owner为nil`时   

```go
// 不指定Owner
// button的Owner为nil，需要手动调用button.Free()方法进行释放。
// 一般只有继承自TComponent的才拥有Owner。
button := vcl.NewButton(nil)
button.Free()
```

* 3、子控件(`TControl`)位于容器控件(`TWinControl`(也就控件的Parent不为nil时))

```go
  // 当pnl主动调用`Free`时，`btn`控件也会被销毁
  pnl := NewPanel(mainForm)
  pnl.SetParent(mainForm)
  btn := NewButton(mainForm)
  btn.SetParent(pnl)
  pnl.Free() // 此时不管btn的owner是谁，都会被销毁 
```

* 4、`不是`继承自`TComponent` ，v2.0.6开始可尝试开启`-tags finalizerOn`，这样就不需要调用`Free`了

```go

// 此类全都需要手动调用实例的Free方法。
ico := vcl.NewIcon() 
ico.Free()

bmp := vcl.NewBitmap()
bmp.Free() 

mem := vcl.NewMemoryStream()
mem.Free()

```

注：以上都为自己手动创建的对象才需要维护生命周期。原有对象中的属性或者方法返回的一个对象，一般情况下由原对象负责生命周期。 