package main

import "fmt"

// 前景色   背景色
// 30  	40	  黑色
// 31  	41	  红色
// 32  	42	  绿色
// 33  	43    黄色
// 34  	44    蓝色
// 35  	45 	  紫色
// 36  	46 	  青色
// 37  	47	  白色

const (
	textBlack = iota + 30
	textRed
	textGreen
	textYellow
	textBlue
	textPurple
	textCyan
	textWhite
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

func main()  {
	str:="hello world"
	fmt.Printf("\x1b[0;%dm%s\x1b[0m", bgWhite, str)
}
