# Fyne

## 初始化

### 设置字体

```go
os.Setenv("FYNE_FONT", "C:\\Windows\\Fonts\\simhei.ttf")
```

## 数据

```go
mypos := fyne.NewPos(200, 200)
mysize := fyne.NewSize(200, 200)
mypic := theme.FyneLogo()

```

## 应用

### 建立应用

```go
myapp := app.New()
```

### 明亮主题

```go
myapp.Settings().SetTheme(theme.LightTheme())
```

### 打开链接

```go
myapp.OpenURL()
```

### 循环

```go
myapp.Run()
```

## 窗口

### 建立窗口

```go
mywindow := myapp.NewWindow("windowname")
```

### 显示窗口

```go
mywindow.Show()
```

### 居中

```go
mywindow.CenterOnScreen()
```

### 常用方法

```go
mywindow.Title()
mywindow.SetTitle("test")
mywindow.Show()
mywindow.Hide()
mywindow.Close()
mywindow.FullScreen()
mywindow.SetFullScreen(true)
mywindow.Resize(mysize)
mywindow.FixedSize()
mywindow.SetFixedSize(true)
mywindow.Icon()
mywindow.SetIcon(mypic)
mywindow.CenterOnScreen()
mywindow.SetOnClosed(func () { fmt.Println("close") })
mywindow.SetCloseIntercept(func () { fmt.Println("close") })
mywindow.Clipboard()
//请求焦点
mywindow.RequestFocus()
//窗口关闭时退出程序
mywindow.SetMaster()

```

## 提示

```go
fyne.NewNotification("hello", "nice to meet you")
```

## 剪切板

```go
type Clipboard interface {
    Content() string
    SetContent(content string)
}

```

## 颜色

```go
green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
black := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
red := color.NRGBA{R: 180, G: 0, B: 0, A: 255}
white := color.White

```

## 画布

```go
mycanvas := mywindow.Canvas()
```

### 画布对象

```go
mytext := canvas.NewText("text", green)
myrect := canvas.NewRectangle(black)
myline := canvas.NewLine(red)
mycircle := canvas.NewCircle(white)
myimage := canvas.NewImageFromResource(theme.FyneLogo())

```

### 画布方法

```go
mycanvas.Content()
mycanvas.SetContent(mytext)
mycanvas.Refresh(mycircle)
mycanvas.Size()
mycanvas.SetOnTypedKey(func(ke *fyne.KeyEvent) { fmt.Println(ke) })

```

### 画布对象方法

```go
myline.Move(mypos)
myrect.Position()
mytext.Resize(mysize)
mytext.Size()
mycircle.Refresh()
mytext.Hide()
mytext.Show()
mytext.Visible()
myimage.Move(mypos)

```

### 设置描边

```go
mycircle.StrokeWidth = 4
mycircle.StrokeColor = red

```

## 文件打开

### 不立即显示

```go
dialog.NewFileOpen()
dialog.NewFileSave()
dialog.NewFolderOpen()

```

### 立即显示

```go
dialog.ShowFileOpen()
dialog.ShowFileSave()
dialog.ShowFolderOpen()

```

### 选择框大小

> 一定要在窗口**显示**并**刷新**后调整！！！

```go
myfs.Resize(fyne.NewSize(1000,700))
```

### 过滤器

#### 扩展名型过滤器

> Example:`.txt`, `.png` ,`.docx`

```go
storage.NewExtensionFileFilter()
```

#### Mime型过滤器

> Example:`image/*`, `audio/mp3`, `text/plain`, `application/*`

```go
storage.NewMimeTypeFileFilter()
```

## 对话框

```go
dialog.ShowInformation()
dialog.NewInformation()
dialog.ShowError()
dialog.NewError()
dialog.ShowConfirm()
dialog.NewConfirm()
dialog.ShowEntryDialog()
dialog.NewEntryDialog()

```

## 布局

```text
HBox（水平方向单行排列）
VBox（垂直方向单列排列）
Center（居中）
Grid（网格）
GridWrap（自适应的网格）
Border（各边外的剩余空间）
Max（最大填充）
```

## 键盘消息

```go
if mycanvas, ok := mywindow.Canvas().(desktop.Canvas); ok {
	mycanvas.SetOnKeyDown(func(event *fyne.KeyEvent) {
		fmt.Println(event.Name)
	})
	mycanvas.SetOnKeyUp(func(event *fyne.KeyEvent) {
		fmt.Println(event.Name)
	})
}

```

## 组件

### Label

```go
mylab:= widget.NewLabel("hello")
mylab.SetText("ok")

```

### Entry

```go
myent := widget.NewMultiLineEntry()
myent.SetText("abc")
myent.SelectedText()
myent.Text
myent.CursorColumn
myent.CursorRow
//底衬文字
myent.SetPlaceHolder("abc")
myent.Cursor()
myent.OnCursorChanged = func() {
    fmt.Println("change:",
        myent.CursorColumn,
        myent.CursorRow,
    )
}
myent.OnChanged= func(s string) {
    fmt.Println(s)
}

```

### Button

```go
mybut.SetIcon(theme.FyneLogo())
mybut.SetText("ok")
mybut.Text
mybut.OnTapped

&widget.Button{
    Alignment:     widget.ButtonAlignTrailing,
    IconPlacement: widget.ButtonIconTrailingText,
    Text:          "hello",
    Icon:          theme.FyneLogo(),
    OnTapped:      func() { fmt.Println("hello") },
}

//ButtonIconLeadingText 图标在文字左侧
//ButtonIconTrailingText 图标在文字右侧

//ButtonAlignCenter 文字在按钮中间
//ButtonAlignLeading 文字在按钮左侧
//ButtonAlignTrailing 文字在按钮右侧

```

### Menu

```go
//弹出菜单
widget.ShowPopUpMenuAtPosition()

```

## 常用事件接口

### 焦点

```go
// Focusable describes any CanvasObject that can respond to being focused.
// It will receive the FocusGained and FocusLost events appropriately.
// When focused it will also have TypedRune called as text is input and
// TypedKey called when other keys are pressed.
//
// Note: You must not change canvas state (including overlays or focus) in FocusGained or FocusLost
// or you would end up with a dead-lock.
type Focusable interface {
	// FocusGained is a hook called by the focus handling logic after this object gained the focus.
	FocusGained()
	// FocusLost is a hook called by the focus handling logic after this object lost the focus.
	FocusLost()
	// Deprecated: this is an internal detail, canvas tracks current focused object
	Focused() bool

	// TypedRune is a hook called by the input handling logic on text input events if this object is focused.
	TypedRune(rune)
	// TypedKey is a hook called by the input handling logic on key events if this object is focused.
	TypedKey(*KeyEvent)
}

```

### 鼠标移动

```go
// Hoverable is used when a canvas object wishes to know if a pointer device moves over it.
type Hoverable interface {
	// MouseIn is a hook that is called if the mouse pointer enters the element.
	MouseIn(*MouseEvent)
	// MouseMoved is a hook that is called if the mouse pointer moved over the element.
	MouseMoved(*MouseEvent)
	// MouseOut is a hook that is called if the mouse pointer leaves the element.
	MouseOut()
}

```

### 鼠标点击

```go
// Mouseable represents desktop mouse events that can be sent to CanvasObjects
type Mouseable interface {
	MouseDown(*MouseEvent)
	MouseUp(*MouseEvent)
}

// Tappable describes any CanvasObject that can also be tapped.
// This should be implemented by buttons etc that wish to handle pointer interactions.
type Tappable interface {
    Tapped(*PointEvent)
}

// SecondaryTappable describes a CanvasObject that can be right-clicked or long-tapped.
type SecondaryTappable interface {
    TappedSecondary(*PointEvent)
}

// DoubleTappable describes any CanvasObject that can also be double tapped.
type DoubleTappable interface {
    DoubleTapped(*PointEvent)
}

```

### 鼠标滚动

```go
// Scrollable describes any CanvasObject that can also be scrolled.
// This is mostly used to implement the widget.ScrollContainer.
type Scrollable interface {
	Scrolled(*ScrollEvent)
}

```

### 鼠标拖动

```go
// Draggable indicates that a CanvasObject can be dragged.
// This is used for any item that the user has indicated should be moved across the screen.
type Draggable interface {
	Dragged(*DragEvent)
	DragEnd()
}

```

### 键盘

```go
// Keyable describes any focusable canvas object that can accept desktop key events.
// This is the traditional key down and up event that is not applicable to all devices.
type Keyable interface {
	fyne.Focusable

	KeyDown(*fyne.KeyEvent)
	KeyUp(*fyne.KeyEvent)
}

```

## 自定义组件

### 菜单弹出按钮

```go
type contextMenuButton struct {
	widget.Button
	menu *fyne.Menu
}

func (b *contextMenuButton) Tapped(e *fyne.PointEvent) {
	widget.ShowPopUpMenuAtPosition(b.menu, fyne.CurrentApp().Driver().CanvasForObject(b), e.AbsolutePosition)
}

func newContextMenuButton(label string, menu *fyne.Menu) *contextMenuButton {
	b := &contextMenuButton{menu: menu}
	b.Text = label

	b.ExtendBaseWidget(b)
	return b
}

```


### 填充

```go
package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"image/color"
)

var ma = app.New()
var mw = ma.NewWindow("test")

type Rec struct {
	widget.BaseWidget
	r       *canvas.Rectangle
	Color   color.Color
	render  *RecRender
	minSize fyne.Size
}

type RecRender struct {
	Rec *Rec
}

func (rr *RecRender) Destroy() {}

func (rr *RecRender) Layout(size fyne.Size) {
	rr.Rec.r.Resize(size)
	inset := fyne.NewPos(theme.Padding()/2, theme.Padding()/2)
	rr.Rec.r.Move(inset)
	rr.Rec.r.FillColor = rr.Rec.Color
}

func (rr *RecRender) MinSize() fyne.Size {
	return rr.Rec.minSize
}

func (rr *RecRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{rr.Rec.r}
}

func (rr *RecRender) Refresh() {
	rr.Layout(rr.Rec.Size())
	rr.Rec.r.Refresh()
	canvas.Refresh(rr.Rec)
}

func (rr *RecRender) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *Rec) CreateRenderer() fyne.WidgetRenderer {
	if r.render == nil {
		rr := new(RecRender)
		rr.Rec = r
		r.render = rr
		return rr
	} else {
		return r.render
	}
}

func (r *Rec) SetMinSize(size fyne.Size) {
	r.minSize = size
}

func NewRec(clo color.Color) *Rec {
	var f = new(Rec)
	f.r = canvas.NewRectangle(f.Color)
	f.Color = clo
	f.minSize = fyne.NewSize(100, 40)
	f.Resize(fyne.NewSize(100, 40))
	f.BaseWidget.ExtendBaseWidget(f)
	f.CreateRenderer()
	f.Refresh()
	return f
}

func main() {
	mw.SetContent(container.NewVBox(
		widget.NewLabel("left"),
		NewRec(color.Black),
		widget.NewLabel("right"),
	))
	mw.ShowAndRun()
}

```
