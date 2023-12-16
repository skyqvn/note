Lazarus中可视控件的布局大多数都默认的是绝对定位，一般没有像QT那种横向，竖向的布局思想（但是Lazarus中有些容器控件带有`ChildSizing`属性，可以对子控件做出一些类似横向、竖向的布局）。  

所有的可视控件都拥有一个Align属性，用来自动调整控件大小的，可以说是Lazarus的一种布局方式吧。更多请前往例程中的"layout"查看。   

与控件位置大小相关的方法有Left、SetLeft、Top、SetTop、Width、SetWidth、Height、SetHeight、SetBounds、ClientRect。

* Align  （layout例子有演示）

![Align](https://images.gitee.com/uploads/images/2020/0424/131235_e7c25696_118989.png "屏幕截图.png")   

Align是一个枚举类型，所有的值我在go中我按Lazarus中的定义为：  
```go
// 见 types 包
const (
	AlNone = iota + 0   // 不进行自动调整, 大多数控件默认值
	AlTop               // 居于顶部， 仅SetHeight有效
	AlBottom            // 居于底部， 仅SetHeight有效
	AlLeft              // 居于左边， 仅SetWidth有效
	AlRight             // 居于右边， 仅SetWidth有效
	AlClient            // 填充客户区域，无法手动调整控件大小
	AlCustom            // 自定义，这个很少用, 需要的可以用OnAlignPosition事件，详细用法见layout例子。
)
```  
如：使用时直接调用Control.SetAlign(types.AlTop)，即可使用控件居于父控件的顶部，并自动调整Top, Left, Right的值。  

* BorderSpacing  （layout例子有演示）

通过此属性可以另外调整控件的边距   

![BorderSpacing](https://images.gitee.com/uploads/images/2020/0424/130451_1b665d48_118989.png "屏幕截图.png")   


* Anchors  （layout例子有演示）

锚点属性，相对于某一控件的位置  

![Anchors](https://images.gitee.com/uploads/images/2020/0424/131314_4c5c6fc9_118989.png "屏幕截图.png")  

Anchors属性为一个集合类型，有akLeft, akTop, akRight, akBottom，可多选，一般默认的为akLeft, akTop。他的作用是，比如一个Button当Align设置为alNone时，我想这个控件居于右边  
并且不想改变控件的大小，这时可以将其设置为akTop, akRight。当父窗口尺寸发生变化时会自动
调整到设计模式的位置。   


与其相关的属性和方法：  

```go
// methods
func AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) // 当前控件的位置以ASide位于ASibling控件ASide偏移ASpace(右边)。
func AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) // 当前控件位置以ASide以ASibling的ASide位置偏移（左边）ASpace。
func AnchorHorizontalCenterTo(ASibling IControl) // 置于指定控件的横向中心
func AnchorVerticalCenterTo(ASibling IControl)  // 置于指定控件的纵向中心。
func AnchorSame(ASide TAnchorKind, ASibling IControl) //2.0.6版本开始支持此接口。
func AnchorAsAlign(ATheAlign TAlign, ASpace int32)
func AnchorClient(ASpace int32)

// propertys
func AnchorSide(AKind TAnchorKind) *TAnchorSide
func AnchorSideLeft() *TAnchorSide 
func AnchorSideTop() *TAnchorSide
func AnchorSideRight() *TAnchorSide
func AnchorSideBottom() *TAnchorSide
```
 
* ChildSizing  （simaplecalc例子有演示）

容器类控件所拥有的，用于对子控件的一些布局方面的调整。。。比如，横向，坚向排列等等。。。。     

![ChildSizing](https://images.gitee.com/uploads/images/2020/0424/130904_7e77a1b7_118989.png "屏幕截图.png")   

### 使用TPanel进行布局

一般情况下布局方面都选择使用TPanel作为布局的容器，比如说想实现一个“上-左-客户区”的布局，这里就需要2-3个TPanel，这里使用3个TPanel来演示，将3个TPanel的Align属性分别设置为alLeft, alTop, alClient就可实现。注意控件的创建顺序哦，详情看例程中的“layout”。  