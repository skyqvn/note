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
Left//？
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
Top//？
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

BorderSpacing//控件边距

