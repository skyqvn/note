# GOVCL



## Form窗口属性

```text
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
AlphaBlend//透明色混合？
AlphaBlendValue//透明色混合值？
AutoScroll//自动滚动？
AutoSize//自动大小？
BiDiMode//文字排列方向
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
BorderWidth//？
Caption//标题
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
Color//颜色
Constraints//极限大小
	MaxHeight
	MaxWidth
	MinHeight
	MinWidth
Cursor//光标形状
DefaultMonitor//？
DesignTimePPI//？
DockSite//？
DoubleBuffered//双缓冲，建议为True
DragKind//拖动类型？
DragMode//拖动模式？
Enabled//是否启用
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
FormStyle//
	- fsNormal//既不是MDI父窗口，也不是一个MDI子窗口
	- fsMDIChild//是一个MDI子窗口
	- fsMDIForm//是一个MDI父窗口
	- fsStayOnTop//窗体始终处在最前端
Height//高度
HelpContext//？
HelpFile//？
HelpKeyword//？
HelpType//？
Hint//提示文字（鼠标悬浮时的提示文字）
HorzScrollBar//横向滚动条？
	Increment
	Page
	Position
	Range
	Smooth
	Tracking
	Visible
Icon//图标
KeyPreview//？
Left//到左侧距离，用于设置窗口位置
Menu//菜单
Name//名称
ParentBiDiMode//？
ParentDoubleBuffered//？
ParentFont//？
PixelsPerlnch//分辨率，PPI
PopupMenu//弹出菜单
PopupMode//弹出模式？
PopupParent//弹出的父对象
Position//？
Scaled//？
SessionProperties//会话属性
ShowHint//是否显示提示文字
ShowlnTaskBar//是否在任务栏显示
Tag//？
Top//到顶部距离，用于设置窗口位置
UseDockManager//？
VertScrollBar//纵向滚动条？
	Increment
	Page
	Position
	Range
	Smooth
	Tracking
	Visible
Visible//是否显示
Width//宽度
WindowState//窗口状态
	- wsNormal//普通状态
	- wsFullScreen//全屏状态
	- wsMaximized//最大化状态
	- wsMinimized//最小化状态
```

## 常用属性

```text
//通用
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
Hint//提示文字（鼠标悬浮时的提示文字）
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

//Form
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
	- fsStayOnTop//窗体始终处在最前端
Left//到左侧距离，用于设置窗口位置
PixelsPerlnch//分辨率，PPI
Top//到顶部距离，用于设置窗口位置
ShowlnTaskBar//是否在任务栏显示
WindowState//窗口状态
	- wsNormal//普通状态
	- wsFullScreen//全屏状态
	- wsMaximized//最大化状态
	- wsMinimized//最小化状态

//Panel

//边框样式
BevelColor
BevelInner
BevelOuter
BevelWidth

//Shape

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



## Application

> 详见./wiki/实例类/Application.markdown

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