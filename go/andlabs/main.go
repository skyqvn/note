package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func main() {
	err := ui.Main(func() {
		//生成输入框
		input := ui.NewEntry()
		input.SetText("这是一个输入框")
		input.LibuiControl()
		
		//生成范围框
		spinbox := ui.NewSpinbox(50, 150)
		spinbox.SetValue(55)
		
		//生成滑片
		slider := ui.NewSlider(0, 100)
		slider.SetValue(10)
		
		//生成进度条
		processbar := ui.NewProgressBar()
		processbar.SetValue(10)
		
		//生成多选框
		combobox := ui.NewCombobox()
		combobox.Append("选择1")
		combobox.Append("选择2")
		combobox.Append("选择3")
		combobox.SetSelected(2)
		
		//生成单选框
		checkbox1 := ui.NewCheckbox("Go语言")
		checkbox1.SetChecked(true)
		
		//生成下拉框
		checkbox2 := ui.NewCheckbox("qqhd.me")
		checkbox3 := ui.NewCheckbox("Python")
		checkbox4 := ui.NewCheckbox("Other")
		
		checkbox_div := ui.NewHorizontalBox()
		checkbox_div.Append(checkbox1, true)
		checkbox_div.Append(checkbox2, true)
		checkbox_div.Append(checkbox3, true)
		checkbox_div.Append(checkbox4, true)
		
		radio := ui.NewRadioButtons()
		radio.Append("Go语言")
		radio.Append("qqhd.me")
		radio.Append("Python")
		radio.Append("Other")
		
		checkbox_div.SetPadded(true)
		
		Separator := ui.NewHorizontalSeparator()
		Separator_label_l := ui.NewLabel("left")
		Separator_label_r := ui.NewLabel("right")
		
		Separator_div := ui.NewHorizontalBox()
		Separator_div.Append(Separator_label_l, true)
		Separator_div.Append(Separator, false)
		Separator_div.Append(Separator_label_r, true)
		Separator_div.SetPadded(true)
		datetimepicker := ui.NewDateTimePicker()
		
		//---------------------将单个子项设置为新组-----------------
		conT1 := ui.NewGroup("输入框")
		conT1.SetChild(input)
		
		conT2 := ui.NewGroup("设值范围框，只能通过箭头控制值，不能手动输入")
		conT2.SetChild(spinbox)
		
		conT3 := ui.NewGroup("滑片")
		conT3.SetChild(slider)
		
		conT4 := ui.NewGroup("进度条")
		conT4.SetChild(processbar)
		
		conT5 := ui.NewGroup("多选框")
		conT5.SetChild(checkbox_div)
		
		cont6 := ui.NewGroup("单选框")
		cont6.SetChild(radio)
		
		cont7 := ui.NewGroup("下拉框")
		cont7.SetChild(combobox)
		
		cont8 := ui.NewGroup("分隔符")
		cont8.SetChild(Separator_div)
		
		cont9 := ui.NewGroup("时间选取器")
		cont9.SetChild(datetimepicker)
		
		//------垂直排列的容器---------
		div := ui.NewVerticalBox()
		
		//------水平排列的容器
		boxs_1 := ui.NewHorizontalBox()
		boxs_1.Append(conT1, true)
		boxs_1.Append(conT2, true)
		
		boxs_1.SetPadded(false)
		
		boxs_2 := ui.NewHorizontalBox()
		boxs_2.Append(conT3, true)
		boxs_2.Append(conT4, true)
		
		boxs_3 := ui.NewHorizontalBox()
		boxs_3.Append(conT5, true)
		boxs_3.Append(cont6, true)
		
		boxs_4 := ui.NewHorizontalBox()
		boxs_4.Append(cont7, true)
		boxs_4.Append(cont8, true)
		
		div.Append(boxs_1, true)
		div.Append(boxs_2, true)
		div.Append(boxs_3, true)
		div.Append(boxs_4, true)
		div.Append(cont9, true)
		div.SetPadded(false)
		
		window := ui.NewWindow("测试软件", 600, 600, true)
		
		window.SetChild(div)
		window.SetMargined(true)
		
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	
	if err != nil {
		panic(err)
	}
}
