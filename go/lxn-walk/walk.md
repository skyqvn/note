# Walk



## rsrc

```bash
rsrc -manifest xxx.exe.manifest -o rsrc.syso
```


## window
```go
w, err := walk.NewMainWindow()
if err != nil {
	fmt.Println(err)
	return
}

window := MainWindow{
	AssignTo: &w,
	Title:    "test",
	MinSize:  Size{600, 400},
	Layout:   VBox{},
	Children: []Widget{
        //···
	},
}
window.Create()
//各种set操作
//...
w.Run()
```



### Set

```go
//设置坐标
w.SetX(0)
w.SetY(0)
//设置大小
w.SetHeight()
w.SetWidth()
w.SetSize(walk.Size{
	Width:  100,
	Height: 100,
})
//设置极限大小
w.SetMinMaxSize(walk.Size{
	Width:  100,
	Height: 100,
}, walk.Size{
	Width:  100,
	Height: 100,
})
//设置窗口标题
w.SetName("hello")
//设置窗口图标
i, err := walk.NewImageFromFileForDPI("icon.png", 96)
if err != nil {
	fmt.Println(err)
	return
}
w.SetIcon(i)
//全屏
w.SetFullscreen(true)
//得到焦点
w.SetFocus()
//设置是否禁用
w.SetEnabled(false)
//设置鼠标参数
w.SetCursor()
//设置字体
w.SetFont()
//不允许窗口大小调整
win.SetWindowLong(w.Handle(), win.GWL_STYLE,
	win.GetWindowLong(w.Handle(), win.GWL_STYLE)&^win.WS_MAXIMIZEBOX&^win.WS_THICKFRAME)
```

### 事件

```text
OnBoundsChanged    walk.EventHandler
OnKeyDown          walk.KeyEventHandler
OnKeyPress         walk.KeyEventHandler
OnKeyUp            walk.KeyEventHandler
OnMouseDown        walk.MouseEventHandler
OnMouseMove        walk.MouseEventHandler
OnMouseUp          walk.MouseEventHandler
OnSizeChanged      walk.EventHandler
```



## win包

```go
win.SetWindowLong(w.Handle(), win.GWL_STYLE,
	win.GetWindowLong(w.Handle(), win.GWL_STYLE)&^win.WS_MAXIMIZEBOX)//禁用最大化按钮
//···
win.RemoveMenu(win.GetSystemMenu(window.Handle(), false), win.SC_CLOSE, win.MF_BYCOMMAND)//禁用关闭按钮
//···
win.SetCursorPos(100, 100)//设置鼠标位置
win.GetWindowRect()//屏幕大小
win.ScreenToClient()//屏幕坐标=>窗口坐标
win.ClientToScreen()//窗口坐标=>屏幕坐标
```

### WS_

```text
win.WS_MAXIMIZEBOX 有最大化按钮
win.WS_MINIMIZEBOX 有最小化按钮
win.WS_THICKFRAME 允许拉伸窗口大小

```

### 其他函数

```go
walk.DriveNames()
```



## 组件

### PushButton

#### 事件

```text
//窗口事件···
OnClicked walk.EventHandler
```

### CheckBox

#### 事件

```text
//窗口事件···
OnCheckedChanged walk.EventHandler
OnClicked        walk.EventHandler
OnCheckStateChanged walk.EventHandler
```

### RadioButton

#### 事件

```text
//可配合RadioButtonGroupBox
//窗口事件···
OnClicked walk.EventHandler
```

### ProgressBar

#### 属性

```text
MaxValue//最大值
MinValue//最小值
Value//值
MarqueeMode//是否为未知模式（true时进度条不断移动，false时进度条稳定在Value位置）
```

### Combobox

#### 属性

```go
DisplayMember//显示的值
BindingMember//数据绑定的值
Model//模型
Value//值
CurrentIndex//默认选项的索引
Editable//是否可编辑
```

### 容器

```go
Composite//普通容器
GradientComposite
GroupBox//分组（有边框线）
RadioButtonGroupBox
HSplitter//横向分割（可鼠标调整宽度）
VSplitter//纵向分割（可鼠标调整宽度）
ScrollView//带进度条
TabPage
```

### 常用组件

```go
SplitButton//下拉按钮
Combobox//下拉菜单
CheckBox
DateEdit
DateLabel
ImageView
LineEdit
LinkLabel
ListBox
NumberEdit
NumberLabel
ProgressBar
TextEdit
TextLabel
ToolBar
ToolButton
TreeView
VSeparator//纵向分割（横线）
HSeparator//横向分割（竖线）
VSpacer//纵向填充
HSpacer//横向填充
```

## 字体

```go
Font{
	Family:    "",//字体家族
	PointSize: 0,//字号
	Bold:      false,//粗体
	Italic:    false,//斜体
	Underline: false,//下划线
	StrikeOut: false,//删除线
}
```



## 文件对话框

### 设置

```go
fd := walk.FileDialog{
	Title:          "打开文档",   //标题
	FilePath:       "",         //默认文件
	FilePaths:      nil,        //默认多个文件
	InitialDirPath: "./",       //初始目录
	Filter:         "Word文档 (*.doc;*.docx)|*.doc;*.docx|Excel表格 (*.xls;*.xlsx)|*.xls;*.xlsx|PPT演示文稿 (*.ppt;*.pptx)|*.ppt;*.pptx",   //过滤器
	FilterIndex:    2,          //默认过滤器，是一般意义上的第几项。如此处第2项是Excel，而不是PPT
	ShowReadOnlyCB: false,      //是否添加只读CheckBox
}
```

### 打开

```go
fd.ShowOpen(w)         //打开文件
fd.ShowOpenMultiple(w) //打开多个文件
fd.ShowBrowseFolder(w) //打开文件夹
fd.ShowSave(w)         //打开保存
```

### 结果

```go
fd.FilePath//结果文件名
fd.FilePaths//多个结果文件名，只有ShowOpenMultiple时会使用
fd.FilterIndex//使用的过滤器项数，是一般意义上的第几项。
fd.Flags//使用的过滤器中的后缀名项数，是一般意义上的第几项。如docx是第二项。
```

### 过滤器

*e.g.*

```text
Word文档 (*.doc;*.docx)|*.doc;*.docx|Excel表格 (*.xls;*.xlsx)|*.xls;*.xlsx|PPT演示文稿 (*.ppt;*.pptx)|*.ppt;*.pptx
```

## 菜单

```go
MainWindow{
	AssignTo: &w,
	Title:    "test",
	MinSize:  Size{600, 400},
	Layout:   VBox{},
	Children: []Widget{},
	MenuItems: []MenuItem{
		Menu{//<--- 主菜单1
			//···
			Items: []MenuItem{
				Action{//<--- 按键1
					//···
				},
				Separator{},//<--- 分隔符1
				Menu{//<--- 次级菜单1
					//···
					Items: []MenuItem{
						//···
					},
				},
			},
		},
		//···
	},
}
```

### 更改

```go
walk.Menu{}.Actions().Add()
walk.Menu{}.Actions().AddMenu()
//···
```

## 托盘

```go
//https://www.jianshu.com/p/bbe8ba2e925a
package main

import (
	"github.com/lxn/walk"
	"log"
)
func main()  {
	mw, err := walk.NewMainWindow()
	if err != nil {
		log.Fatal(err)
	}
	//托盘图标文件
	icon, err := walk.Resources.Icon("./icon.ico")
	if err != nil {
		log.Fatal(err)
	}
	ni, err := walk.NewNotifyIcon(mw)
	if err != nil {
		log.Fatal(err)
	}
	defer ni.Dispose()
	if err := ni.SetIcon(icon); err != nil {
		log.Fatal(err)
	}
	if err := ni.SetToolTip("鼠标在icon上悬浮的信息."); err != nil {
		log.Fatal(err)
	}
	ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}
		if err := ni.ShowCustom("Walk 任务栏通知标题","walk 任务栏通知内容",icon); err != nil {
			//ni.ShowCustom()
			//ni.ShowInfo()
			//ni.ShowMessage()
			//ni.ShowWarning()
			//ni.ShowError()
			log.Fatal(err)
		}
	})
	exitAction := walk.NewAction()
	if err := exitAction.SetText("右键icon的菜单按钮"); err != nil {
		log.Fatal(err)
	}
	//Exit 实现的功能
	exitAction.Triggered().Attach(func() { walk.App().Exit(0) })
	if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
		log.Fatal(err)
	}
	if err := ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}
	if err := ni.ShowInfo("Walk NotifyIcon Example", "Click the icon to show again."); err != nil {
		log.Fatal(err)
	}
	mw.Run()
}

```



## 官方示例

```text
actions 下拉菜单
clipboard 剪切板
databinding 数据绑定
drawing 画笔
dropfiles 拖动打开文件
externalwidgets 拓展组件
filebrowser 树状列表/【应用】文件浏览器
gradientcomposite 渐变
imageicon 图片动态生成
imageview 图片显示方式
imageviewer 【应用】图片查看器
linklabel 链接状文本
listbox 列表 【应用】环境变量查看器
listbox_ownerdrawing 双栏列表框
logview 日志
multiplepags 多页
notifyicon 消息提醒
progressindicator 图标进度
radiobutton 多选框
settings 设置表（可排序文字表格）
slider 滑块
statusbar 状态栏（窗口下部显示状态区）
tableview 表格
webview 【应用】IE浏览器
webview_events 【应用】浏览器事件
```

## progress indicator

#### 状态

```text
walk.PINoProgress   无(!)
walk.PIIndeterminat 不定位置（不断运动）(!)
walk.PINormal       普通
walk.PIError        错误
walk.PIPaused       暂停
```

> 标 `(!)` 的需要在进度条不变化时设置

```go
package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"time"
)

func main() {
	window, err := walk.NewMainWindow()
	if err != nil {
		log.Fatalln(err)
	}
	MainWindow{
		AssignTo: &window,
		Title:    "pi",
		Layout:   VBox{},
		Children: []Widget{
			PushButton{
				Text: "Start",
				OnClicked: func() {
					go func() {
						window.ProgressIndicator().SetTotal(100)
						var i uint32
						for i = 0; i < 100; i++ {
							fmt.Println("SetProgress", i)
							time.Sleep(100 * time.Millisecond)
							if err := window.ProgressIndicator().SetCompleted(i); err != nil {
								log.Print(err)
							}
						}
					}()
				},
			},
			//walk.PINoProgress   无
			//walk.PIIndeterminat 不定位置（不断运动）
			//walk.PINormal       普通
			//walk.PIError        错误
			//walk.PIPaused       暂停
			
			PushButton{
				Text: "NoProgress",
				OnClicked: func() {
					window.ProgressIndicator().SetState(walk.PINoProgress)
				},
			},
			PushButton{
				Text: "Normal",
				OnClicked: func() {
					window.ProgressIndicator().SetState(walk.PINormal)
				},
			},
			PushButton{
				Text: "Indeterminate",
				OnClicked: func() {
					window.ProgressIndicator().SetState(walk.PIIndeterminate)
				},
			},
			PushButton{
				Text: "Error",
				OnClicked: func() {
					window.ProgressIndicator().SetState(walk.PIError)
				},
			},
			PushButton{
				Text: "Paused",
				OnClicked: func() {
					window.ProgressIndicator().SetState(walk.PIPaused)
				},
			},
		},
	}.Create()
	
	window.Run()
}

```

## Resource

```go
walk.Resources.SetRootDirPath("../img")

walk.Resources.Image()
walk.Resources.Icon()
walk.Resources.BitmapForDPI()
```

## Image

#### 显示方式

```go
ImageViewModeIdeal 
ImageViewModeCorner 
ImageViewModeCenter 
ImageViewModeShrink 
ImageViewModeZoom 
ImageViewModeStretch 
```

![图片显示方式](D:\BaiduSyncdisk\code\note\go\lxn-walk\图片显示方式.png)

## MsgBox

### MsgBoxStyle

```text
MsgBoxOK
MsgBoxOKCancel
MsgBoxAbortRetryIgnore
MsgBoxYesNoCancel
MsgBoxYesNo
MsgBoxRetryCancel
MsgBoxCancelTryContinue
MsgBoxIconHand
MsgBoxIconQuestion
MsgBoxIconExclamation
MsgBoxIconAsterisk
MsgBoxUserIcon
MsgBoxIconWarning
MsgBoxIconError
MsgBoxIconInformation
MsgBoxIconStop
MsgBoxDefButton1
MsgBoxDefButton2
MsgBoxDefButton3
MsgBoxDefButton4
MsgBoxApplModal
MsgBoxSystemModal
MsgBoxTaskModal
MsgBoxHelp
MsgBoxSetForeground
MsgBoxDefaultDesktopOnly
MsgBoxTopMost
MsgBoxRight
MsgBoxRTLReading
MsgBoxServiceNotification
```

### DlgCmd

```text
DlgCmdNone
DlgCmdOK
DlgCmdCancel
DlgCmdAbort
DlgCmdRetry
DlgCmdIgnore
DlgCmdYes
DlgCmdNo
DlgCmdClose
DlgCmdHelp
DlgCmdTryAgai
DlgCmdContinu
DlgCmdTimeout
```



## Example

### 动态图片生成

```go
package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"image"
	"image/color"
)

func main() {
	im := image.NewRGBA(image.Rect(0, 0, 16, 16))
	for i := 0; i < 16; i++ {
		im.Set(i, i, color.Black)
	}
	for i := 0; i < 16; i++ {
		im.Set(15-i, i, color.Black)
	}
	w, _ := walk.NewMainWindow()
	i, err := walk.NewIconFromImageForDPI(im, 96)
	if err != nil {
		fmt.Println(err)
		return
	}
	wi := MainWindow{
		AssignTo: &w,
		Layout:   HBox{},
		Children: []Widget{},
	}
	wi.Create()
	w.SetIcon(i)
	w.Run()
}

```