# walk

## user32

### WS_
https://learn.microsoft.com/zh-cn/previous-versions/visualstudio/visual-studio-2012/czada357(v=vs.110)
https://learn.microsoft.com/zh-cn/previous-versions/visualstudio/visual-studio-2012/61fe4bte(v=vs.110)
```text
WS_BORDER 创建具有边框的窗口。
WS_CAPTION 创建具有标题栏的窗口 (即表示 WS_BORDER 样式)。不能使用 WS_DLGFRAME 样式。
WS_CHILD 创建子窗口。不能使用 WS_POPUP 样式。
WS_CHILDWINDOW 和 WS_CHILD 样式相同。
在父窗口中，绘制WS_CLIPCHILDREN 排除子窗口占用的区域。使用，在创建父窗口。
WS_CLIPSIBLINGS 剪辑相对的子窗口;也就是说，在特定子窗口接收绘制消息时， WS_CLIPSIBLINGS 样式剪裁其他重载重叠的子窗口在要更新的子窗口之外的区域。(如果未给出 WS_CLIPSIBLINGS ，以及窗口重叠的子级，那么，当您在子窗口的工作区中绘制，它是在有相邻的子窗口的工作区中绘制。)用于 WS_CHILD 样式中只使用。
WS_DISABLED 创建初始禁用的窗口。
WS_DLGFRAME 用一个双边框，但没有书名创建一个窗口。
WS_GROUP 指定一组控件的第一个控件用户可以从一个控件移动到下使用箭头键。所有控件定义与 WS_GROUP 样式 FALSE 在第一个控件后属于同一组。与 WS_GROUP 样式的下一个控件开始下一组 (即一组结束下开始) 的位置。
WS_HSCROLL 创建具有水平滚动条的窗口。
WS_ICONIC 创建初始最小化的窗口。和 WS_MINIMIZE 样式相同。
WS_MAXIMIZE 创建一个窗口最大大小。
WS_MAXIMIZEBOX 创建一个具有最大化按钮的窗口。
WS_MINIMIZE 创建初始最小化的窗口。用于 WS_OVERLAPPED 样式中只使用。
WS_MINIMIZEBOX 创建具有最小化按钮的窗口。
WS_OVERLAPPED 创建重叠的窗口。一个重叠的窗口通常具有说明和一个边框。
WS_OVERLAPPEDWINDOW 用 WS_OVERLAPPED、 WS_CAPTION、 WS_SYSMENU、 WS_THICKFRAME、 WS_MINIMIZEBOX和 WS_MAXIMIZEBOX 样式创建重叠的窗口。
WS_POPUP 创建一个弹出窗口。不能使用 WS_CHILD 样式。
WS_POPUPWINDOW 用 WS_BORDER、 WS_POPUP和 WS_SYSMENU 样式创建一个弹出窗口。必须合并 WS_CAPTION 样式。 WS_POPUPWINDOW 样式允许控制菜单可见。
WS_SIZEBOX 创建一个具有大小调整边框的窗口。和 WS_THICKFRAME 样式相同。
WS_SYSMENU 创建一个具有控件菜单框在其标题栏的窗口。仅用于具有标题栏的窗口。
WS_TABSTOP 指定用户可以移动任何数量的控件之一使用 tab 键，tab 键移动用户移到 WS_TABSTOP 样式指定的下一个控件。
WS_THICKFRAME 使用 (可用于调整窗口中粗的帧创建一个窗口。
WS_TILED 创建重叠的窗口。一个重叠的窗口的标题栏和一个边框。和 WS_OVERLAPPED 样式相同。
WS_TILEDWINDOW 用 WS_OVERLAPPED、 WS_CAPTION、 WS_SYSMENU、 WS_THICKFRAME、 WS_MINIMIZEBOX和 WS_MAXIMIZEBOX 样式创建重叠的窗口。和 WS_OVERLAPPEDWINDOW 样式相同。
WS_VISIBLE 创建初始可见的窗口。
WS_VSCROLL 创建具有垂直滚动条的窗口。
```

### WS_EX_
```text
WS_EX_ACCEPTFILES 指定用此样式创建的窗口接受拖放文件。
当窗口可见时，WS_EX_APPWINDOW 强制在任务栏上的顶级窗口。
WS_EX_CLIENTEDGE 指定是窗口查看三维)，使用一个凹下转到的边缘的一个边框。
WS_EX_CONTEXTHELP 在窗口的标题栏包含一个问号。用户单击问号时，光标变为一个带有指针的问号。如果用户然后单击子窗口，子 WM_HELP 接收消息。
使用 tab 键，WS_EX_CONTROLPARENT 允许用户在窗口的子窗口之间导航。
WS_EX_DLGMODALFRAME 指定包含可的一个双边框的窗口 (可选) 在标题栏创建，当您在 dwStyle 参数中指定 WS_CAPTION 样式标志。
WS_EX_LAYERED 窗口是 a。分层窗口此样式，则窗口具有 CS_OWNDC 或 CS_CLASSDC，则 选件类样式 不能使用。但是，Windows 8 支持子窗口的 WS_EX_LAYERED 样式，前面的 windows 版本支持仅对顶级窗口中。

WS_EX_LEFT 为窗口泛型左对齐的属性。这是默认设置。
WS_EX_LEFTSCROLLBAR 在客户端区域左边将垂直滚动条。
使用从左向右读取 orders 属性命令，WS_EX_LTRREADING 显示窗口文本。这是默认设置。
WS_EX_MDICHILD 创建 MDI 子窗口。
WS_EX_NOPARENTNOTIFY 指定用此样式创建的子窗口不会发送 WM_PARENTNOTIFY 信息到其父窗口，在子窗口创建或销毁时。
WS_EX_OVERLAPPEDWINDOW 合并 WS_EX_CLIENTEDGE 和 WS_EX_WINDOWEDGE 样式
WS_EX_PALETTEWINDOW 合并 WS_EX_WINDOWEDGE 和 WS_EX_TOPMOST 样式。
WS_EX_RIGHT 为窗口泛型右对齐的属性。这取决于窗口选件类。
WS_EX_RIGHTSCROLLBAR 在客户端区右边将垂直滚动条 (如果有)。这是默认设置。
使用从右向左的读取 orders 属性命令，WS_EX_RTLREADING 显示窗口文本。
WS_EX_STATICEDGE 用预期的一个三维边框样式创建一个窗口为不接受用户输入的项目。
WS_EX_TOOLWINDOW 创建一个工具窗口，是预期的窗口用作浮动工具栏。工具窗口具有使用较小的字体，比普通标题栏短的标题栏，并且，窗口标题绘制。工具窗口未显示在任务栏或于显示的窗口在用户按 ALT+TAB。
WS_EX_TOPMOST 指定用此样式创建的窗口应是放置的第 nontopmost 窗口和保持在其中发生，即使当停用窗口。应用程序可以使用 SetWindowPos 成员函数添加或移除此特性。
WS_EX_TRANSPARENT 指定用此样式创建的窗口是透明的。即在 windows 下的任何窗口未由 windows 遮盖。只有在更新后，用此样式创建的窗口接收消息 WM_PAINT 其下方的所有同级窗口。
WS_EX_WINDOWEDGE 指定窗口一个凸出的边缘的一个边框。
```

### HWND_
```text
HWND_TOP = 0 在前面
HWND_BOTTOM = 1 在后面
HWND_TOPMOST = -1 在前面, 位于任何顶部窗口的前面
HWND_NOTOPMOST = -2 在前面, 位于其他顶部窗口的后面
HWND_DESKTOP   = 0
HWND_MESSAGE   = -3
HWND_BROADCAST = 0xFFFF
```

### SM_
```text
SM_ARRANGE

56

指定系统如何排列最小化窗口的标志。 有关更多信息，请参阅本主题中的"备注"部分

SM_CLEANBOOT

67

指示系统的启动方式：

0 Normal boot
1 Fail-safe boot
2 Fail-safe with network boot
故障安全启动（也称为SafeBoot，安全模式或干净启动）会绕过用户启动文件

SM_CMONITORS

80

桌面上的显示器数量。 有关更多信息，请参阅本主题中的"备注"部分。

SM_CMOUSEBUTTONS

43

鼠标上的按钮数量，如果没有安装鼠标，则为零。

SM_CONVERTIBLESLATEMODE

0x2003

反映笔记本电脑或平板模式的状态，0表示平板模式，否则为非零。 当此系统指标发生变化时，系统通过WM_SETTINGCHANGE在LPARAM中发送带有"ConvertibleSlateMode"的广播消息。 请注意，此系统指标不适用于台式PC。 在这种情况下，请使用GetAutoRotationState。

SM_CXBORDER

5

窗口边框的宽度（以像素为单位）。 这相当于具有3-D外观的窗口的SM_CXEDGE值。

SM_CXCURSOR

13

光标的宽度，以像素为单位。 系统无法创建其他大小的游标。

SM_CXDLGFRAME

7

同 SM_CXFIXEDFRAME.

SM_CXDOUBLECLK

36

双击序列中第一次单击位置周围的矩形宽度（以像素为单位）。 第二次单击必须出现在由SM_CXDOUBLECLK和SM_CYDOUBLECLK定义的矩形内，以便系统考虑双击两次单击。 两次单击也必须在指定时间内发生。

要设置双击矩形的宽度，请使用SPI_SETDOUBLECLKWIDTH调用SystemParametersInfo。

SM_CXDRAG

68

鼠标指针在拖动操作开始之前可以移动的鼠标指向任意一侧的像素数。 这允许用户轻松地单击和释放鼠标按钮而不会无意中开始拖动操作。 如果此值为负，则从鼠标按下点的左侧减去该值并将其添加到其右侧。

SM_CXEDGE

45

三维边框的宽度（以像素为单位）。 该指标是SM_CXBORDER的3-D对应物。

SM_CXFIXEDFRAME

7

围绕窗口周边的框架的厚度，其具有标题但不是相当大的像素。 SM_CXFIXEDFRAME是水平边框的高度，SM_CYFIXEDFRAME是垂直边框的宽度。

该值与SM_CXDLGFRAME相同。

SM_CXFOCUSBORDER

83

DrawFocusRect绘制的焦点矩形的左右边缘的宽度。 该值以像素为单位。

Windows 2000：不支持此值。

SM_CXFRAME

32

同 SM_CXSIZEFRAME.

SM_CXFULLSCREEN

16

主显示器监视器上全屏窗口的客户端区域宽度（以像素为单位）。 要获取未被系统任务栏或应用程序桌面工具栏遮挡的屏幕部分的坐标，请使用SPI_GETWORKAREA值调用SystemParametersInfo函数。

SM_CXHSCROLL

21

水平滚动条上箭头位图的宽度（以像素为单位）。

SM_CXHTHUMB

10

水平滚动条中拇指框的宽度，以像素为单位。

SM_CXICON

11

图标的默认宽度（以像素为单位）。 LoadIcon函数只能加载具有SM_CXICON和SM_CYICON指定维度的图标。

SM_CXICONSPACING

38

大图标视图中项目的网格单元格的宽度，以像素为单位。 安排时，每个项目都适合SM_CYICONSPACING的SM_CXICONSPACING大小的矩形。 该值始终大于或等于SM_CXICON。

SM_CXMAXIMIZED

61

主显示器上最大化顶级窗口的默认宽度（以像素为单位）。

SM_CXMAXTRACK

59

具有标题和大小调整边框的窗口的默认最大宽度（以像素为单位）。 此指标指的是整个桌面。 用户无法将窗口框架拖动到大于这些尺寸的尺寸。 窗口可以通过处理WM_GETMINMAXINFO消息来覆盖此值。

SM_CXMENUCHECK

71

默认菜单复选标记位图的宽度（以像素为单位）。

SM_CXMENUSIZE

54

菜单栏按钮的宽度，例如多文档界面中使用的子窗口关闭按钮，以像素为单位。

SM_CXMIN

28

窗口的最小宽度（以像素为单位）。

SM_CXMINIMIZED

57

最小化窗口的宽度（以像素为单位）。

SM_CXMINSPACING

47

最小化窗口的网格单元格的宽度（以像素为单位）。 每个最小化的窗口在排列时都适合这个尺寸的矩形。 该值始终大于或等于SM_CXMINIMIZED。

SM_CXMINTRACK

34

窗口的最小跟踪宽度（以像素为单位）。 用户无法将窗口框架拖动到小于这些尺寸的尺寸。 窗口可以通过处理WM_GETMINMAXINFO消息来覆盖此值。

SM_CXPADDEDBORDER

92

标题窗口的边框填充量（以像素为单位）。

Windows XP / 2000：不支持此值。

SM_CXSCREEN

0

主显示器屏幕的宽度，以像素为单位。 这是通过调用GetDeviceCaps获得的相同值，如下所示：GetDeviceCaps（hdcPrimaryMonitor，HORZRES）。

SM_CXSIZE

30

窗口标题或标题栏中按钮的宽度（以像素为单位）。

SM_CXSIZEFRAME

32

围绕窗口周边的大小调整边框的粗细，可以调整大小，以像素为单位。 SM_CXSIZEFRAME是水平边框的宽度，SM_CYSIZEFRAME是垂直边框的高度。

该值与SM_CXFRAME相同。

SM_CXSMICON

49

建议的小图标宽度（以像素为单位）。 小图标通常出现在窗口标题和小图标视图中。

SM_CXSMSIZE

52

小字幕按钮的宽度，以像素为单位。

SM_CXVIRTUALSCREEN

78

虚拟屏幕的宽度（以像素为单位）。 虚拟屏幕是所有显示器的边界矩形。 SM_XVIRTUALSCREEN度量标准是虚拟屏幕左侧的坐标。

SM_CXVSCROLL

2

垂直滚动条的宽度（以像素为单位）。

SM_CYBORDER

6

窗口边框的高度（以像素为单位）。 这相当于具有3-D外观的窗口的SM_CYEDGE值。

SM_CYCAPTION

4

标题区域的高度，以像素为单位。

SM_CYCURSOR

14

光标的高度，以像素为单位。 系统无法创建其他大小的游标。

SM_CYDLGFRAME

8

同 SM_CYFIXEDFRAME.

SM_CYDOUBLECLK

37

双击序列中第一次单击位置周围的矩形高度，以像素为单位。 第二次单击必须在SM_CXDOUBLECLK和SM_CYDOUBLECLK定义的矩形内进行，系统才能考虑双击两次单击。 两次单击也必须在指定时间内发生。

要设置双击矩形的高度，请使用SPI_SETDOUBLECLKHEIGHT调用SystemParametersInfo。

SM_CYDRAG

69

在拖动操作开始之前鼠标指针可以移动的鼠标按下点上方和下方的像素数。 这允许用户轻松地单击和释放鼠标按钮而不会无意中开始拖动操作。 如果此值为负值，则从鼠标按下点上方减去它并在其下方添加。

SM_CYEDGE

46

三维边框的高度（以像素为单位）。 这是SM_CYBORDER的3D对应物。

SM_CYFIXEDFRAME

8

围绕窗口周边的框架的厚度，其具有标题但不是相当大的像素。 SM_CXFIXEDFRAME是水平边框的高度，SM_CYFIXEDFRAME是垂直边框的宽度。

该值与SM_CYDLGFRAME相同。

SM_CYFOCUSBORDER

84

DrawFocusRect绘制的焦点矩形的顶部和底部边缘的高度。 该值以像素为单位。

Windows 2000：不支持此值。

SM_CYFRAME

33

同 SM_CYSIZEFRAME.

SM_CYFULLSCREEN

17

主显示器监视器上全屏窗口的客户端区域高度（以像素为单位）。 要使屏幕部分的坐标不被系统任务栏或应用程序桌面工具栏遮挡，请使用SPI_GETWORKAREA值调用SystemParametersInfo函数。

SM_CYHSCROLL

3

水平滚动条的高度，以像素为单位。

SM_CYICON

12

图标的默认高度（以像素为单位）。 LoadIcon函数只能加载尺寸为SM_CXICON和SM_CYICON的图标。

SM_CYICONSPACING

39

大图标视图中项目的网格单元格的高度，以像素为单位。 安排时，每个项目都适合SM_CYICONSPACING的SM_CXICONSPACING大小的矩形。 该值始终大于或等于SM_CYICON。

SM_CYKANJIWINDOW

18

对于系统的双字节字符集版本，这是屏幕底部的汉字窗口的高度，以像素为单位。

SM_CYMAXIMIZED

62

主显示监视器上最大化顶级窗口的默认高度（以像素为单位）。

SM_CYMAXTRACK

60

具有标题和大小调整边框的窗口的默认最大高度（以像素为单位）。 此指标指的是整个桌面。 用户无法将窗口框架拖动到大于这些尺寸的尺寸。 窗口可以通过处理WM_GETMINMAXINFO消息来覆盖此值。

SM_CYMENU

15

单行菜单栏的高度（以像素为单位）。

SM_CYMENUCHECK

72

默认菜单复选标记位图的高度（以像素为单位）。

SM_CYMENUSIZE

55

菜单栏按钮的高度，例如多文档界面中使用的子窗口关闭按钮，以像素为单位。

SM_CYMIN

29

窗口的最小高度（以像素为单位）。

SM_CYMINIMIZED

58

最小化窗口的高度（以像素为单位）。

SM_CYMINSPACING

48

最小化窗口的网格单元格的高度，以像素为单位。 每个最小化的窗口在排列时都适合这个尺寸的矩形。 此值始终大于或等于SM_CYMINIMIZED。

SM_CYMINTRACK

35

窗口的最小跟踪高度（以像素为单位）。 用户无法将窗口框架拖动到小于这些尺寸的尺寸。 窗口可以通过处理WM_GETMINMAXINFO消息来覆盖此值。

SM_CYSCREEN

1

主显示器屏幕的高度，以像素为单位。 这是通过调用GetDeviceCaps获得的相同值，如下所示：GetDeviceCaps（hdcPrimaryMonitor，VERTRES）。

SM_CYSIZE

31

窗口标题或标题栏中按钮的高度（以像素为单位）。

SM_CYSIZEFRAME

33

围绕窗口周边的大小调整边框的粗细，可以调整大小，以像素为单位。 SM_CXSIZEFRAME是水平边框的宽度，SM_CYSIZEFRAME是垂直边框的高度。

该值与SM_CYFRAME相同。

SM_CYSMCAPTION

51

小标题的高度，以像素为单位。

SM_CYSMICON

50

建议的小图标高度（以像素为单位）。 小图标通常出现在窗口标题和小图标视图中。

SM_CYSMSIZE

53

小字幕按钮的高度，以像素为单位。

SM_CYVIRTUALSCREEN

79

虚拟屏幕的高度，以像素为单位。 虚拟屏幕是所有显示器的边界矩形。 SM_YVIRTUALSCREEN指标是虚拟屏幕顶部的坐标。

SM_CYVSCROLL

20

垂直滚动条上箭头位图的高度，以像素为单位。

SM_CYVTHUMB

9

垂直滚动条中拇指框的高度，以像素为单位。

SM_DBCSENABLED

42

如果User32.dll支持DBCS，则为非零; 否则，0。

SM_DEBUG

22

如果安装了User.exe的调试版本，则为非零; 否则，0。

SM_DIGITIZER

94

如果当前操作系统是Windows 7或Windows Server 2008 R2且Tablet PC输入服务已启动，则为非零; 返回值是一个位掩码，指定设备支持的数字化仪输入的类型。 有关更多信息，请参阅备注。

Windows Server 2008，Windows Vista和Windows XP / 2000：不支持此值。

SM_IMMENABLED

82

如果启用了输入法管理器/输入法编辑器功能，则为非零; 否则，0。

SM_IMMENABLED指示系统是否已准备好在Unicode应用程序上使用基于Unicode的IME。 要确保依赖于语言的IME正常工作，请检查SM_DBCSENABLED和系统ANSI代码页。 否则，可能无法正确执行ANSI到Unicode转换，或者可能不存在某些组件（如字体或注册表设置）。

SM_MAXIMUMTOUCHES

95

如果系统中有数字化仪，则为非零; 否则，0。

SM_MAXIMUMTOUCHES返回系统中每个数字化仪支持的最大联系数的最大值。 如果系统只有单触式数字化仪，则返回值为1.如果系统具有多点触摸数字化仪，则返回值是硬件可以提供的同时触点数。

Windows Server 2008，Windows Vista和Windows XP / 2000：不支持此值。

SM_MEDIACENTER

87

如果当前操作系统是Windows XP，Media Center Edition则为非零，否则为0。

SM_MENUDROPALIGNMENT

40

如果下拉菜单与相应的菜单栏项右对齐，则为非零; 如果菜单左对齐，则为0。

SM_MIDEASTENABLED

74

如果系统启用了希伯来语和阿拉伯语，则为非零，否则为0。

SM_MOUSEPRESENT

19

如果安装了鼠标，则为非零; 否则，0。该值很少为零，因为支持虚拟鼠标，并且因为某些系统检测到端口的存在而不是鼠标的存在。

SM_MOUSEHORIZONTALWHEELPRESENT

91

如果安装了带水平滚轮的鼠标，则为非零; 否则为0。

SM_MOUSEWHEELPRESENT

75

如果安装了带垂直滚轮的鼠标，则为非零; 否则为0。

SM_NETWORK

63

如果存在网络，则设置最低有效位; 否则，它被清除。 其他位保留供将来使用。

SM_PENWINDOWS

41

如果安装了Microsoft Windows for Pen计算扩展，则为非零; 否则为零。

SM_REMOTECONTROL

0x2001

此系统度量标准用于终端服务环境，以确定是否正在远程控制当前的终端服务器会话。 如果当前会话是远程控制的，它的值是非零的; 否则，0。

您可以使用终端服务管理工具（如终端服务管理器（tsadmin.msc）和shadow.exe）来控制远程会话。 当远程控制会话时，另一个用户可以查看该会话的内容并可能与其进行交互。

SM_REMOTESESSION

0x1000

此系统度量标准用于终端服务环境中。 如果调用进程与终端服务客户端会话关联，则返回值为非零。 如果调用进程与终端服务控制台会话关联，则返回值为0。

Windows Server 2003和Windows XP：控制台会话不一定是物理控制台。 有关更多信息，请参阅WTSGetActiveConsoleSessionId。

SM_SAMEDISPLAYFORMAT

81

如果所有显示器具有相同的颜色格式，则为非零，否则为0.两个显示器可以具有相同的位深度，但颜色格式不同。 例如，红色，绿色和蓝色像素可以用不同数量的比特编码，或者那些比特可以位于像素颜色值中的不同位置。

SM_SECURE

44

应忽略此系统指标; 它总是返回0。

SM_SERVERR2

89

如果系统是Windows Server 2003 R2，则为内部版本号; 否则，0。

SM_SHOWSOUNDS

70

非零，如果用户要求应用程序以可视方式呈现信息，否则它将仅以可听形式呈现信息; 否则，0。

SM_SHUTTINGDOWN

0x2000

如果当前会话正在关闭，则为非零; 否则，0。

Windows 2000：不支持此值。

SM_SLOWMACHINE

73

如果计算机具有低端（慢）处理器，则为非零; 否则，0。

SM_STARTER

88

如果当前操作系统是Windows 7 Starter Edition，Windows Vista Starter或Windows XP Starter Edition，则为非零; 否则，0。

SM_SWAPBUTTON

23

如果交换了鼠标左键和右键的含义，则为非零; 否则，0。

SM_SYSTEMDOCKED

0x2004

反映停靠模式的状态，0表示未锁定模式，否则为非零。 当此系统指标发生变化时，系统通过WM_SETTINGCHANGE通过LPARAM中的"SystemDockMode"发送广播消息。

SM_TABLETPC

86

如果当前操作系统是Windows XP Tablet PC版本，或者当前操作系统是Windows Vista或Windows 7且Tablet PC输入服务已启动，则为非零; SM_DIGITIZER设置指示运行Windows 7或Windows Server 2008 R2的设备支持的数字化仪输入类型。 有关更多信息，请参阅备注。

SM_XVIRTUALSCREEN

76

虚拟屏幕左侧的坐标。 虚拟屏幕是所有显示器的边界矩形。 SM_CXVIRTUALSCREEN指标是虚拟屏幕的宽度。

SM_YVIRTUALSCREEN

77

虚拟屏幕顶部的坐标。 虚拟屏幕是所有显示器的边界矩形。 SM_CYVIRTUALSCREEN指标是虚拟屏幕的高度。
```
