# Walk

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
		PushButton{
			Text: "SCREAM",
			OnClicked: func() {
				outTE.SetText(strings.ToUpper(inTE.Text()))
			},
		},
	},
}
window.Run()
```

### Set初始化

```go
window.Create()
//w.SetIcon(i)
//set操作
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
//窗口事件···
OnClicked walk.EventHandler
```

### 官方示例

```text
actions 下拉菜单
clipboard 剪切板
databinding 数据绑定
drawing 画笔
dropfiles 拖动打开文件
externalwidgets 拓展组件
filebrowser 文件浏览器
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

### progress indicator

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

### Resource

```go
walk.Resources.SetRootDirPath("../img")
```

### Image

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

### MsgBox

#### MsgBoxStyle

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

#### DlgCmd

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