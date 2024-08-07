# GOVCL

## 常用属性

### 通用

```text
Caption//标题
Left//到左侧距离，用于设置控件位置
Top//到顶部距离，用于设置控件位置
Width//宽度
Height//高度
Constraints//极限大小
	MaxHeight
	MaxWidth
	MinHeight
	MinWidth
BorderSpacing//周围间距
	- ···
Hint//提示文字（鼠标悬浮或自动提示的提示文字），用A|B格式写时，A为鼠标悬浮时的提示文字，B为自动提示（多用在StatusBar中）文字
ShowHint//是否显示提示文字
Cursor//光标形状
Visible//是否显示
Enabled//是否启用
Color//颜色
Align//排列
	- alBottom
	- alClient
	- alCustom
	- alLeft
	- alNone
	- alRight
	- alTop
Anchors//锚点（相对位置）
AllowDropFile//是否允许拖入文件
ChildSizing//布局方式（横向、纵向等）
	ControlsPerLine//每行空间数
	EnlargeHorizontal//横向对齐选项
		- crsAnchorAligning//左端对齐
		- crsHomogenousChildResize//均匀拉伸对齐
		- crsHomogenousSpaceResize//均匀分散对齐（不拉伸）
		- crsScaleChilds//与crsHomogenousChildResize相似？
	EnlargeVertical//纵向对齐选项
		- //同上
	HorizontalSpacing//横向控件间间隔
	Layout
		- cclLeftToRightThenTopToBottom
		- cclNone
		- cclTopToBottomThenLeftToRight
	LeftRightSpacing//左右两侧空间
	ShrinkHorizontal//横向空间不足时
	ShrinkVertical//纵向空间不足时
	TopBottomSpacing//上下两侧空间
	VerticalSpacing//纵向控件间间隔
Default//是否默认得焦
DoubleBuffered//双缓冲，建议为True
Font//字体
	CharSet
	Color
	Height
	Name
	Orientation
	Pitch
	Quality
	Size
	Style//样式
		fsBold
		fsltalic
		fsStrikeOut
		fsUnderline
Menu//菜单
Name//名称
PopupMenu//弹出菜单
PopupMode//弹出模式？
PopupParent//弹出的父对象

```

### Form

```text
AlphaBlend//是否启用透明度
AlphaBlendValue//透明度值（数值越小越透明）
BorderIcons//边框按钮
	biHelp//帮助（仅在biMaximize、biMinimize均为False时显示）
	biMaximize（最大化按钮）
	biMinimize（最小化按钮）
	biSystemMenu（系统菜单，False时窗口只有标题，无图标与任何按钮）
BorderStyle//边框模式
	bsDialog//对话框（标题、叉号）
	bsNone（无）
	bsSingle（全部、窄边框）
	bsSizeable（全部）
	bsSizeToolWin（标题、方叉号、可变大小）
	bsToolWindow（标题、方叉号、不可变大小）
FormStyle//
	- fsNormal//既不是MDI父窗口，也不是一个MDI子窗口
	- fsMDIChild//是一个MDI子窗口
	- fsMDIForm//是一个MDI父窗口
	- fsSystemStayOnTop//窗体始终处在最前端
	- fsStayOnTop//窗体可以处在最前端
Left//到左侧距离，用于设置窗口位置
PixelsPerlnch//分辨率，PPI
Top//到顶部距离，用于设置窗口位置
ShowlnTaskBar//是否在任务栏显示
WindowState//窗口状态
	- wsNormal//普通状态
	- wsFullScreen//全屏状态
	- wsMaximized//最大化状态
	- wsMinimized//最小化状态

```

### Panel

```text
//边框样式
BevelColor
BevelInner
BevelOuter
BevelWidth

```

### ImageButton

```text
//将图片分割成几份，平时第一份，鼠标悬浮第二份，按下第三份
ImageCount
//图片分割方向
Orientation
	- ioHorizontal
	- ioVertical


```

### Shape

```text
Shape
	- stCircle
	- stDiamond
	- stEllipse
	- stRectangle
	- stRoundRect
	- stRoundSquare
	- stSquare
	- stSquaredDiamond
	- stStar
	- stStarDown
	- stTriangle
	- stTriangleDown
	- stTriangleLeft
	- stTriangleRight
Brush//填充笔刷
	Color//颜色
	Style//刷子样式
		- bsBDiagonalbsClear
		- bsCross
		- bsDiagCross
		- bsFDiagonal
		- bsHorizontall
		- bslmage
		- bsPattern
		- bsSolid
		- bsVertical
Pen
	Color
	Cosmeti
	EndCap//结束笔头
		- pecFlat
		- pecRound
		- pecSquare
	JoinStyle
		- pisBevell
		- pisMiter
		- pjsRound
	Mode
	Style//笔形
		- psClear
		- psDash
		- psDashDot
		- psDashDotDot
		- psDot
		- psinsideFrame
		- psPattern
		- psSolid
	Width//宽度

```

### 文本

```text
WordWrap//文本自动折行

```

### 选项

```text
ItemHeight//选项高度
ItemIndex//选中选项索引
Items//选项列表
ItemWidth//选项宽度

```

### ComboBox

```text
Style
	- csDropDown//可编辑，下拉
	- csDropDownList//不可编辑，下拉
	- csOwnerDrawEditableFixed
	- csOwnerDrawEditableVariable
	- csOwnerDrawFixed
	- csOwnerDrawVariable
	- csSimple

```

### ScrollBox

```text
HorzScrollBar//横向滚动条
VertScrollBar//纵向滚动条
	Increment//单击时移动
	Page
	Position//当前值
	Range//最大
	Smooth
	Tracking//平滑，就是鼠标拖动时页面移动，而不是等鼠标释放后再移动.
	Visible

```

## 窗口方法

```go
form.Show()//显示窗口
form.ShowModal()//将窗口以遮盖父窗口的形式弹出，与Show不重复使用
form.SetParent()//如果不是ShowModal，窗口将被嵌入父控件

```

## 鼠标

> 定义见govcl\types\cursors.go

### 自定义

```go
//定义光标
const MyCursor=1
i := vcl.NewIcon()
i.LoadFromFile("icon.ico")
vcl.Screen.SetCursors(MyCursor, i.Handle())//定义新光标，索引为1，形状是i的形状

//设置光标
vcl.Screen.SetCursor(MyCursor)//将光标设置为索引为1的光标
vcl.Screen.SetCursor(types.CrDefault)
```

### 常量

```text
CrHigh
CrDefault
CrNone
CrArrow
CrCross
CrIBeam
CrSize
CrSizeNESW
CrSizeNS
CrSizeNWSE
CrSizeWE
CrSizeNW
CrSizeN
CrSizeNE
CrSizeW
CrSizeE
CrSizeSW
CrSizeS
CrSizeSE
CrUpArrow
CrHourGlass
CrDrag
CrNoDrop
CrHSplit
CrVSplit
CrMultiDrag
CrSQLWait
CrNo
CrAppStart
CrHelp
CrHandPoint
CrSizeAll
CrLow
```



## 异常处理

```go
vcl.Application.SetOnException(func(sender vcl.IObject, e *vcl.Exception) {
	// 在这里自行处理VCL中的异常
})
```

## 接口

```text
基类接口 IObject

所有可视/非可视组件类接口 IComponent

所有可视控件类接口 IControl

所有容器控件类接口 IWinControl
```



## Event

### 关联事件

```go
type TMainForm struct {
	*vcl.TForm
	Button1 *vcl.TButton
	Button2 *vcl.TButton `events:"OnButton1Click"`//强制将OnButton1Click关联到Button2
	Button3 *vcl.TButton
}

func (f *TMainForm) OnButton1Click(sender vcl.IObject) {//已自动关联到Button1
	
}

f.Button3.SetOnClick(f.OnButton1Click)//动态将OnButton1Click关联到Button3
```

### 原理

```text
govcl会自动将有SetOn_EventType()的结构的On_StructName_EventType()绑定

见源代码autoBindEvents、addComponentNotifyEvent、addApplicationNotifyEvent、findAndSetEvent
```

### 速查表

```go
// 事件参数中的类型参考types包中的类型。  
// sender 对象源， 啥是对象源，举个例子， button1.click；响应onclick事件，事件参数中的sender即为对象源也就button1  

// 公共的通知事件方法，像单击、鼠标进入、鼠标离开、双击等等
//    sender 对象源
type TNotifyEvent func(sender IObject)

// TUpDown按钮的上、下按钮调节事件
//    sender 对象源
//    button 哪个按钮
type TUDClickEvent func(sender IObject, button TUDBtnType)

// TForm独有的关闭事件
//    sender 为对象源
//    action 可指定关闭的动作
type TCloseEvent func(sender IObject, action *TCloseAction)

// TForm关闭事件，会发生在OnClose之前
//    sender   对象源
//    canClose 可以决定阻止关闭窗口
type TCloseQueryEvent func(sender IObject, canClose *bool) 

// 菜单的改变事事件
//    sender 对象源
//    source 源菜单对
//    rebuild是否重构造
type TMenuChangeEvent func(sender IObject, source *TMenuItem, rebuild bool)

// TLinkLabel连接被点击事件
//    sender 对象源
//    link   当前点击的连接
//    linkType 点击的连接类型，限windows vista之上的系统
type TSysLinkEvent func(sender IObject, link string, linkType TSysLinkType)  

// TApplication异常捕获事件，可在此事件中捕获vcl中的所有异常
//    sender为对象源，
//    e为异常信息对象，一般只用e.Message()获取消息内容
type TExceptionEvent func(sender IObject, e *Exception)

// 键盘按键事件
//    sender为对象源
//    key为按键码
//    shift表示控制按键状态
type TKeyEvent func(sender IObject, key *Char, shift TShiftState)

// 键盘按键事件
//    sender为对象源
//    key为按键
type TKeyPressEvent func(sender IObject, key *Char)

// 鼠标按下或者抬起事件
//    sender为对象源
//    button中鼠标当前按下的按钮（左，中，右）键
//    shift为控制键状态
//    x,y为鼠标坐标，相对坐标，以当前控件0,0起
type TMouseEvent func(sender IObject, button TMouseButton, shift TShiftState, x, y int32)

// 鼠标移动事件
//    sender为对象源
//    shift为控制键状态
//    x,y为鼠标移动坐标，以当前控件0,0起
type TMouseMoveEvent func(sender IObject, shift TShiftState, x, y int32)

// 鼠标滚轮滚动事件
//   sender为对象源
//   shift为控制
//   wheelDelta 滚动长度，分正负
//   x, y 滚动时鼠标的位置
//   handled 是否处理
type TMouseWheelEvent func(sender IObject, shift TShiftState, wheelDelta, x, y int32, handled *bool)

// TListBox自绘事件
//   control 对象源
//   index  项目索引
//   aRect 绘制区域
//   state 绘制状态
type TDrawItemEvent func(control IControl, index int32, aRect TRect, state TOwnerDrawState)

// TMenuItem自绘事件
//   sender 对象源
//   aCanvas菜单对象画布
//   selected 是否选择
type TMenuDrawItemEvent func(sender IObject, aCanvas *TCanvas, aRect TRect, selected bool)
 
// TListView柱头单击事件
//   sender 对象源
//   column 当前点击的column对象
type TLVColumnClickEvent func(sender IObject, column *TListColumn)

// TListView柱头右键单击事件
//   sender 对象源
//   column 当前点击的column对象
//   point 右键点击的坐标
type TLVColumnRClickEvent func(sender IObject, column *TListColumn, point TPoint)

// TListView项目选择事件
//   sender 为对象源
//   item 选择的项目对象
//   selected 是否选中
type TLVSelectItemEvent func(sender IObject, item *TListItem, selected bool)

// TListView checkbox选中/取消选中事件
//  sender 对象源
//  item 选择的项目对象
type TLVCheckedItemEvent func(sender IObject, item *TListItem)

// TListView排序比较事件
//   sender 对象源
//   item1 待比较项目1
//   item2 待比较项目2
//   data 调用传入的数据，限windows，vcl
//   compare 反回比较数据，小于0，等于0，大于0的数据
type TLVCompareEvent func(sender IObject, item1, item2 *TListItem, data int32, compare *int32)

// TListView项目改变事件
//   sender 对象源
//   item 改变的项目对象
//   change 改变的类型，文字，图像，状态
type TLVChangeEvent func(sender IObject, item *TListItem, change TItemChange)

// TListView通知事件
//   sender 对象源
//   item 通知的对象
type TLVNotifyEvent func(sender IObject, item *TListItem)

// TListView高级自绘项目事件
//   sender 对象
//   aRect 绘制区域
//   stage
//   defaultDraw 是否默认绘制，默认为true
type TLVAdvancedCustomDrawEvent func(sender *TListView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)

// TListView高级自绘项目事件
//   sender 对象源
//   item 当前被绘制的项目对象
//   state 当前项目状态
//   stage 
//   defaultDraw 是否默认绘制，默认为true
type TLVAdvancedCustomDrawItemEvent func(sender *TListView, item *TListItem, state TCustomDrawState, Stage TCustomDrawStage, defaultDraw *bool)

// TListView高级自绘子项目事件
//   sender 对象源
//   item 被绘制的项目对象
//   subItem 绘制的子项目索引
//   state 项目绘制状态
//   stage
//   defaultDraw 是否默认绘制，默认为true
type TLVAdvancedCustomDrawSubItemEvent func(sender *TListView, item *TListItem, subItem int32, state TCustomDrawState, stage TCustomDrawStage, defaultDraw *bool)

// TTreeView排序事件
//    sender 对象源
//    node1 待比较的节点1
//    node2 待比较的节点2
//    data 为调用函数传入的，可为指针等
//    compare 返回比较大小， 大于0，等于0，小于0的整型
type TTVCompareEvent func(sender IObject, node1, node2 *TTreeNode, data int32, compare *int32)

// TTreeView节点展开事件
//   sender 对象源
//   node 当前展开的节点对象
type TTVExpandedEvent func(sender IObject, node *TTreeNode)

// TTreeView节点改变事件
//   sender 对象源
//   node 当前改变的节点对象
type TTVChangedEvent func(sender IObject, node *TTreeNode)

// TTreeView自绘事件
//    sender 对象源
//    aRect 绘制区域
//    stage 绘制状态
//    defaultDraw 是否默认绘制，默认为true
type TTVAdvancedCustomDrawEvent func(sender *TTreeView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)

// TTreeView高级自绘项目事件
//   sender 对象源
//   node 当前绘制节点
//   state 绘制状态
//   stage
//   paintImages 是否绘制图像
//   defaultDraw 是否默认绘制，默认为true
type TTVAdvancedCustomDrawItemEvent func(sender *TTreeView, node *TTreeNode, state TCustomDrawState, stage TCustomDrawStage, paintImages, defaultDraw *bool)

// TPageControl获取图像索引事件
//   sender 对象源
//   tabIndex 绘制的tab索引
//   imageIndex 返回的图像索引
type TTabGetImageEvent func(sender IObject, tabIndex int32, imageIndex *int32)
 
// ToolBar自绘事件
//   sender 对象源
//   aRect  绘制区域坐标
//   stage
//   defaultDraw 是否使用默认绘制，默认为true
type TTBAdvancedCustomDrawEvent func(sender *TToolBar, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)

// ToolBar高级自绘事件，限windows，
//    sender 对象源
//    button 当前绘制的button对象，
//    state  绘制状态
//    stage  
//    flags  绘制标识
//    defaultDraw 是否默认绘制，默认为true
type TTBAdvancedCustomDrawBtnEvent func(sender *TToolBar, button *TToolButton, state TCustomDrawState, stage TCustomDrawStage, flags *TTBCustomDrawFlags, defaultDraw *bool)

// TForm文件拖放事件 
// sender 对象源
// aFileNames 一个文件名数组
type TDropFilesEvent func(sender *TObject, aFileNames []string)

// 窗口约束调整大小事件
//  sender 对象源
//  minWidth 最小宽度
//  minHeight 最大高度
//  maxWidth 最大宽度
//  maxHeight 最大高度
type TConstrainedResizeEvent func(sender IObject, minWidth, minHeight, maxWidth, maxHeight *int32)

// 帮助 F1吧
// command 命令
// Data 
// callHelp 能否使用帮助
type THelpEvent func(command uint16, data THelpEventData, callhelp *bool) bool

// 快捷键事件
// msg 消息
// handled 是否响应
type TShortCutEvent func(msg *TWMKey, handled *bool)

// 上下文件弹出事件，也就一般右键
// sender 对象源
// mousePos 鼠标位置
// handled 是否响应
type TContextPopupEvent func(sender IObject, mousePos TPoint, handled *bool)

// 拖放完成
// sender 对象源
// source 拖放源
// x 坐标
// y 坐标
// state  状态
// accept 是否接受
type TDragOverEvent func(sender, source IObject, x, y int32, state TDragState, accept *bool)

// 拖放下落
// sender 对象源
// source 拖放源
// x 坐标
// y 坐标
type TDragDropEvent func(sender, source IObject, x, y int32)

// 启动拖放
// sender 对象源
// dragObject 拖放对象
type TStartDragEvent func(sender IObject, dragObject *TDragObject)

// 结束拖放
// sender 对象源
// target 目标源
// x 坐标
// y 坐标
type TEndDragEvent func(sender, target IObject, x, y int32)

// 停靠下落
// sender 对象源
// source 目标源
// x 坐标
// y 坐标
type TDockDropEvent func(sender IObject, source *TDragDockObject, x, y int32)

// 停靠完成
// sender 对象源
// source 停靠源
// x 坐标
// y 坐标 
// state 拖放状态
// accept 是否接受
type TDockOverEvent func(sender IObject, source *TDragDockObject, x, y int32, state TDragState, accept *bool)

// 解除停靠
// sender 对象源
// client 客户源
// newTarget 新的目标
// allow 是否允许 
type TUnDockEvent func(sender IObject, client *TControl, newTarget *TControl, allow *bool)

// 启动停靠
// sender 对象源
// dragObject 源
type TStartDockEvent func(sender IObject, dragObject *TDragDockObject)

// 获取位置信息，停靠相关
// sender 对象源
// dockClient
// influenceRect
// mousePos
// canDock
type TGetSiteInfoEvent func(sender IObject, dockClient *TControl, influenceRect *TRect, mousePos TPoint, canDock *bool)

// 鼠标滚轮键按下或者抬起事件
// sender 对象源
// shift 状态
// mousePos 鼠标位置
// handled 是否继承
type TMouseWheelUpDownEvent func(sender IObject, shift TShiftState, mousePos TPoint, handled *bool)

// 消息循环，仅windows, libvcl
// msg Windows消息
// handled 是否继承处理
type TMessageEvent func(msg *TMsg, handled *bool)

// ---- grid
// 表格单元移动事件
// sender 源对象
// FromIndex 源单元格索引
// ToIndex 目标单元格索引
type TMovedEvent func(sender IObject, fromIndex, toIndex int32)

// 表格自绘事件
// sender 源对象
// aCol 列索引
// aRow 行索引
// aRect 单元格矩形大小
// state 绘制状态
```

## 实例类

### Application

> 详见[./wiki/实例类/Application.markdown](./wiki/实例类/Application.markdown)

* `Initialize`  
初始始应用程序相关设置，必要写的

* `Run`  
开始整个程序的消息循环,直到收到WM_QUIT或者主窗口收到WM_CLOSE退出循环，并结束程序


* `RunApp`
一个简化的函数。

* `CreateForm`
创建并返回一个新的TGoForm实例

* `MainFormOnTaskbar`、`SetMainFormOnTaskbar`

* `ProcessMessages`
处理消息，当某主线程中发生阻止时会造成整个消息循环阻止，UI无响应，所以在阻止中，比如循环中加上ProcessMessages可解决UI无响应问题。


* `Terminate`
结束应用程序，发出WM_QUIT消息

* `ExeName`
获取当前运行的应用程序文件名，含路径

* `Icon`, `SetIcon`
获取或者设置应用程序的图标，一般默认会在程序资源中找名为MAINICON的ico资源作为应用程序图标。

* `Title`, `SetTitle`
获取或者设置应用程序标题
* `SetScaled` ,`暂时废弃`
有关dpi的绽放 
* `SetIconResId`,`不推荐使用`
设置App全局icon。从资源中加载，具体看rsrc后的id， 仅windows
* `ScaleForCurrentDpi`

	用于使用纯代码创建的窗口缩放，所有组件创建完后调用 

### Screen

> 详见[./wiki/实例类/Screen.markdown](./wiki/实例类/Screen.markdown)

与屏幕相关的API。Screen由Lazarus在单元initialization和finalization时自动构造和析构。  

常用方法：  

**以像素为单位**  

* `Width`
	当前屏幕宽度


* `Height`
	当前屏幕高度 


* `MonitorCount`
	监视器总数，多显示器情况


* `Monitors`
	获取指定索引监视器

* `FormCount`
	获当前app的Form总数

* `Forms`
	获当前app的Form索引


**桌面相关api，包含任务栏**  

* `DesktopHeight`
	桌面高度


* `DesktopLeft`
	桌面左边位置


* `DesktopTop`
	桌面顶边位置


* `DesktopWidth`
	桌面的宽度


**工作区域，不包含任务栏**  

* `WorkAreaHeight`
	工作区域高度


* `WorkAreaLeft`
	工作区域左边位置


* `WorkAreaTop`
	工作区域顶边位置


* `WorkAreaWidth`
	工作区域的宽度


* `Fonts`
	获取当前系统字体列表


* `Cursor`、`SetCursor`
	获取或者设置当前屏幕光标样式

### Mouse

> 详见[./wiki/实例类/Mouse.markdown](./wiki/实例类/Mouse.markdown)

鼠标相关API。Mouse由Lazarus在单元initialization和finalization时自动构造和析构。  

> 注意，获取或者设置当前屏幕光标样式在Screen中.

常用方法：  

* `CursorPos`
	获取当前光标位置（即鼠标位置，相对屏幕，以左上角为原点，即（0,0））

### Printer

> 详见./samples/printer

### Clipboard
剪切板操作。 Clipboard实例由Lazarus在单元initialization和finalization时自动构造和析构。

常用的几个方法：

| 方法      | 说明                             | 例                                          |
| --------- | -------------------------------- | ------------------------------------------- |
| AsText    | 获取剪切板文字                   | ``` fmt.Println(vcl.Clipboard.AsText()) ``` |
| SetAsText | 设置剪切板文字                   | ``` vcl.Clipboard.SetAsText("1111") ```     |
| Clear     | 清空剪切板                       | ``` vcl.Clipboard.Clear() ```               |
| SetAsHtml | 设置一段html格式到剪切板中       |                                             |
| GetAsHtml | 从剪切板中获取一段html格式字符串 |                                             |
| Open      | 打开剪切板                       | ``` vcl.Clipboard.Open() ```                |
| Close     | 关闭剪切板                       | ``` vcl.Clipboard.Close() ```               |

## 包

* `vcl`
	包含Lazarus标准组件中的大部分
	* `i18n`
		一个简单的多国语言扩展包   

	* `api`
		DLL函数申明与重新包装,一般情况下尽量不要去直接调用api内的函数。
		 * `memorydll`
			Windows下的32bit内存dll加载包。   

	* `rtl`
		包含Lazarus中Set类型操作、内存操作等其它函数   
		* `version`
			包含一个跨平台的系统版本信息获取。   

	* `bitmap`
		Go的image与LCL image对象相互转换。

	* `locales`
		本地化库中的默认资源包。
		* `zh_CN`
			默认的中文资源包。

---

* `win`
	包含windows下的常量、函数、类型定义    
	* `errcode`
		包含windows下GetLastError返回的错误代码含义      

---


* `types`
	包含 类型定义、枚举定义、常量  
	* `colors`
		颜色定义   

	* `keys`
		虚拟键定义  

	* `messages`
		窗口消息常量定义    

---

* `pkgs` 扩展包
	* `macapp`
		MacOS下app打包工具  
	* `libname`
		自定义加载动态连接库   
	* `skinh`
		skinsharp皮肤扩展接口  
	* `winappres`
		包含了一个默认的syso文件
	* `wke`
		一个简单的wke浏览器封封装，只支持windows 32bit
	* `miniblink`
		一个miniblink浏览器组件的包装（有些还有问题）
	* 

* `Move` 
	内存操作，相当于c中的memcpy  


* `MainInstance` 
	Exe自身实例，比如在获取资源中的数据时就需要用到，限windows  

```go
win.LoadIconW(rtl.MainInstance(), 1)
```

### rtl
* `TextToShortCut` 
	将字符串转为ShortCut类型

```go
rtl.TextToShortCut("Ctrl+A")
```


* `ShortCutToText` 
	将ShortCut类型转为字符串  


* `SysOpen` 
	打开，windows下调用ShellExecute  

```go
// windows
rtl.SysOpen("http://www.xxx.com")
rtl.SysOpen("c:\")
rtl.SysOpen("c:\xxx.exe")

// linux or macOS
rtl.SysOpen("https://wwww.xxx.com")
rtl.SysOpen("file:///xxx.png");
```

* `ExtractFilePath` 
	提取文件名的路径，带“\”的  

```go
rtl.ExtractFilePath("C:\\aaa\\bbb\\aaa.text")
output: C:\aaa\bbbb\
```

* `FileExists` 
	判断文件是否存在  

* `ExtractFileExt` 
	获取文件扩展名   

* `ExtractFileName` 
	获取一个文件名   

* `GetFileNameWithoutExt` 
	获取一个无扩展的文件名  

* `Combine`
	合并   

* `LibVersion`
	共8位，2位2位的，如：$01020100 表示 1.2.1.0   

* `LibAbout`
	2.0.2版本之后的api，liblcl的关于信息。   

* `MainThreadId`
	2.0.2版本之后的api，返回主线程ID   

* `CurrentThreadId`
	2.0.2版本之后的api，返回当前线程ID

* `CreateURLShortCut`
	限Windows，创建一个url的快捷方式

```go
rtl.CreateURLShortCut("C:\\aaa\\bbb\\", "govcl", "https://github.com/ying32/govcl")
```

* `CreateShortCut`
	限Windows，创建一个快捷方式

```go
rtl.CreateShortCut("C:\\Users\\administrator\\Desktop\\", "govcl", os.Args[0], "", "", "")
// or
rtl.CreateShortCut("C:\\Users\\administrator\\Desktop\\", "govcl", os.Args[0], "", "Description text", "-a -b")
```

* `InitGoDll`
	2.0.3版本之后的api，用于go的dll中使用govcl，具体参考`samples/Windows/nppPlugins`   

### vcl

* `DEBUG`
	全局变量，默认为false,在CreateForm之前设置值方可生效，开启后，一般在OnFormCreate事件上发生错误后可以打印完整的堆栈信息。  


---

* `ShowMessage`
	显示一个消息框

----

* `ShowMessageFmt`
	显示一个消息框

---

* `MessageDlg`
	显示一个消息框，消息框，Buttons为按钮样式，祥见types.TMsgDlgButtons

---

* `SelectDirectory1`
	选择目录，弹出一个选择目录对话框，老版本样式

---

* `SelectDirectory2`
	选择目录，弹出一个选择目录对话框

* `ThreadSync`
	切换至主线程中运行指定代码，主要用于协程中UI的访问

* `InputBox`
	输入对话框，参考例程sysdialog

* `InputQuery`
	输入对话框，参考例程sysdialog

* `RunApp` 
	简化运行。

* `EqualsObject`
	比较两个对象