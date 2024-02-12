```go
//Arc方法在椭圆上画一段弧，椭圆由(x1,y1)、(x2,y2) 两点所确定的椭圆决定.弧的起点是椭圆圆周和椭圆中心与(x3,y3)连线的交点。弧矩形终点是椭圆圆周和椭圆中心与(x4,y4)连线的交点以逆时针方向画弧.
Arc(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32)
//ArcTo与Arc相似，但不提笔.
ArcTo(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32)
//AngleArc以(X,Y)为圆心，Radius为半径，StartAngle为起始角度，扫过SweepAngle作弧.不提笔.
AngleArc(X int32, Y int32, Radius uint32, StartAngle float32, SweepAngle float32)
//Chord作弓形，规则与Arc相似.
Chord(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32)
//Ellipse画椭圆，两点分别为左上和右下.
Ellipse(X1 int32, Y1 int32, X2 int32, Y2 int32)
//FloodFill洪水填充（以点为种子扩散填色）？
FloodFill(X int32, Y int32, Color types.TColor, FillStyle types.TFillStyle)
//LineTo线画到点.
LineTo(X int32, Y int32)
//MoveTo笔移到点.
MoveTo(X int32, Y int32)
//Pie画扇形.
Pie(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32)
//Rectangle画矩形.
Rectangle(X1 int32, Y1 int32, X2 int32, Y2 int32)
//RoundRect画圆角矩形，X3、Y3为圆角大小.
RoundRect(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32)
//StretchDraw拉伸绘制图片.
StretchDraw(Rect types.TRect, Graphic vcl.IGraphic)
//TextOut画文字.
TextOut(X int32, Y int32, Text string)
//SetPixels设置某个像素颜色.
SetPixels(X int32, Y int32, value types.TColor)
//BrushCopy方法把位图的一部分复制到画布的某个矩形区域，并用画笔的当前颜色替换位图的颜色.参数Dest定义画布的一个矩形区域，该矩形用以填充位图，Bitmap定义位图；Source定义位图中的矩形区域，该区域上的位图将被复制；Color定义画笔中用以替换位图的颜色.
BrushCopy(dest types.TRect, bitmap vcl.IObject, source types.TRect, color types.TColor)
//CopyRect从另一个画布对象上复制部分图像到该画布.Canvas表示源画布，Source是源画布上要复制的图像区域.Dest表示目标画布上将接受复制图像的矩形区域.
CopyRect(dest types.TRect, canvas vcl.IObject, source types.TRect)
//Draw绘制图片.
Draw(x int32, y int32, graphic vcl.IGraphic)
//DrawFocusRect绘制焦点矩形（控件得焦时外围的虚线框）.
DrawFocusRect(aRect types.TRect)
//FillRect填充矩形.
FillRect(aRect types.TRect)
//FrameRect描边矩形.
FrameRect(aRect types.TRect)
//TextRect绘制中心有文字的矩形.
TextRect(aRect types.TRect, x int32, y int32, text string)
//TextRect2与TextRect相似.
TextRect2(aRect *types.TRect, text string, textFormat types.TTextFormat)
//Polygon画多边形.
Polygon(points []types.TPoint)
//Polyline画多边形线.
Polyline(points []types.TPoint)
//PolyBezier画贝塞尔曲线.
PolyBezier(points []types.TPoint)
```




```text
Pixels(X int32, Y int32) types.TColor

SetOnChange(fn vcl.TNotifyEvent)
SetOnChanging(fn vcl.TNotifyEvent)

Brush() *vcl.TBrush
SetBrush(value *vcl.TBrush)
CopyMode() int32
SetCopyMode(value int32)
Font() *vcl.TFont
SetFont(value *vcl.TFont)
Pen() *vcl.TPen
SetPen(value *vcl.TPen)

TextExtent(Text string) types.TSize
TextHeight(Text string) int32
TextWidth(Text string) int32
```
