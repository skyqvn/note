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
}
