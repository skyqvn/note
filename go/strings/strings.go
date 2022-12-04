package main

import (
	"strings"
	"unicode"
)

func main() {
	s := "hello word"

	strings.Count(s, "o") //指定字符串的出现次数

	strings.Contains(s, "word")  //是否包含指定字符串
	strings.ContainsAny(s, "go") //是否包含指定字符串中的任意一个字符
	strings.ContainsRune(s, 'o') //是否包含指定的unicode码值

	strings.Index(s, "o")         //指定字符串首次出现的位置
	strings.LastIndex(s, "o")     //指定字符串最后一次出现的位置
	strings.IndexByte(s, 'o')     //指定字符首次出现的位置
	strings.LastIndexByte(s, 'o') //指定字符最后一次出现的位置
	strings.IndexRune(s, 'o')     //指定unicode码值首次出现的位置
	strings.IndexAny(s, "o")      //指定字符串中的任意一个unicode码值首次出现的位置
	strings.LastIndexAny(s, "o")  //指定字符串中的任意一个unicode码值最后一次出现的位置

	strings.HasPrefix(s, "hello") //是否有指定前缀
	strings.HasSuffix(s, "world") //是否有指定后缀

	s = "中国"

	unicode.Is(unicode.Han, '中') //是否是汉字

	isHan := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}

	strings.IndexFunc(s, isHan)     //第一个汉字位置
	strings.LastIndexFunc(s, isHan) //最后一个汉字位置

	unicode.Is(unicode.Han, '中') //判断指定unicode码值是否是指定文字
	unicode.In('中', unicode.Han) //判断指定unicode码值是否是指定文字

	unicode.IsLower('a') //判断指定unicode码值是不是unicode小写
	unicode.IsUpper('A') //判断指定unicode码值是不是unicode大写
	unicode.IsTitle('ǅ') //判断指定unicode码值是不是unicode标题大小写

	unicode.To(unicode.UpperCase /*unicode.LowerCase unicode.TitleCase*/, 'a')
	unicode.ToUpper('a')
	unicode.ToLower('A')
	unicode.ToTitle('Ǆ')

	unicode.IsOneOf(
		[]*unicode.RangeTable{
			unicode.Han,             //中文
			unicode.ASCII_Hex_Digit, //英文
			unicode.P,               //标点
		}, '中')

	unicode.IsDigit('1')    //判断是不是十进制数字字符
	unicode.IsNumber('1')   //判断是不是数字字符
	unicode.IsLetter('a')   //判断是不是字母字符
	unicode.IsSpace('a')    //判断是不是空格字符
	unicode.IsControl('\a') //判断是不是控制字符
	unicode.IsGraphic(',')  //判断是不是图形字符
	unicode.IsPrint('a')    //判断是不是可打印字符
	unicode.IsPunct('!')    //判断是不是标点字符
	unicode.IsSymbol('+')   //判断是不是符号字符
	unicode.IsSymbol('n')   //判断是不是标记字符
}
