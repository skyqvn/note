package main

import (
	"fmt"
)

// 颜色
// 前景色   背景色
// 30  	40	  黑色
// 31  	41	  红色
// 32  	42	  绿色
// 33  	43    黄色
// 34  	44    蓝色
// 35  	45 	  紫色
// 36  	46 	  青色
// 37  	47	  白色

// 文本模式
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

func main() {
	str := "hello world"
	
	// 按前景、背景、模式打印
	const (
		fgBlack = iota + 30
		fgRed
		fgGreen
		fgYellow
		fgBlue
		fgPurple
		fgCyan
		fgWhite
	)
	
	const (
		bgBlack = iota + 40
		bgRed
		bgGreen
		bgYellow
		bgBlue
		bgPurple
		bgCyan
		bgWhite
	)
	
	const (
		mdDefault           = 0
		mdHighLight         = 1
		mdUnderline         = 4
		mdFlicker           = 5
		mdBackManifestation = 7
		mdInvisible         = 8
	)
	fmt.Printf("\x1b[%d;%d;%dm%s\x1b[0m", mdUnderline, bgWhite, fgYellow, str)
	// 按256色打印
	const (
		fg = 38
		bg = 48
	)
	// 只更改一项
	fmt.Printf("\x1b[%d;5;%dm%s\x1b[0m", bg, 241, str)
	// 同时更改bg和fg
	fmt.Printf("\x1b[38;5;%dm\x1b[48;5;%dm%s\x1b[0m", 120, 241, str)
}
