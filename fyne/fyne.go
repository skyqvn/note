package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"

	"image/color"
)

func main() {
	//建立应用
	myapp := app.New()
	//建立窗口
	mywindow := myapp.NewWindow("windowname")
	//显示窗口
	mywindow.Show()
	//循环
	myapp.Run()
	//数据
	mypos := fyne.NewPos(200, 200)
	mysize := fyne.NewSize(200, 200)
	mypic := theme.FyneLogo()
	//窗口方法
	mywindow.Title()
	mywindow.SetTitle("test")
	mywindow.Show()
	mywindow.Hide()
	mywindow.Close()
	mywindow.FullScreen()
	mywindow.SetFullScreen(true)
	mywindow.Resize(mysize)
	mywindow.FixedSize()
	mywindow.Icon()
	mywindow.SetIcon(mypic)
	mywindow.CenterOnScreen()
	mywindow.SetOnClosed(func() { fmt.Println("close") })
	mywindow.SetCloseIntercept(func() { fmt.Println("close") })
	mywindow.SetFixedSize(true)
	mywindow.Clipboard()
	//剪切板结构
	//type Clipboard interface {
	//	Content() string
	//	SetContent(content string)
	//}
	//颜色
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	black := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	red := color.NRGBA{R: 180, G: 0, B: 0, A: 255}
	white := color.White
	//画布
	mycanvas := mywindow.Canvas()
	//画布对象
	mytext := canvas.NewText("text", green)
	myrect := canvas.NewRectangle(black)
	myline := canvas.NewLine(red)
	mycircle := canvas.NewCircle(white)
	//画布方法
	mycanvas.Content()
	mycanvas.SetContent(mytext)
	mycanvas.Refresh(mycircle)
	mycanvas.Size()
	mycanvas.SetOnTypedKey(func(ke *fyne.KeyEvent) { fmt.Println(ke) })
	//画布对象方法
	myline.Move(mypos)
	myrect.Position()
	mytext.Resize(mysize)
	mytext.Size()
	mycircle.Refresh()
	mytext.Hide()
	mytext.Show()
	mytext.Visible()

	mycircle.StrokeWidth = 4
	mycircle.StrokeColor = red
}
