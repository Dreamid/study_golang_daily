package gip

import (
	"strings"
	"unicode"
)

//不使用额外空间，翻转字符串。
func ReverseString(str string) string {

	slice1 := []rune(str)
	l := len(str)
	for i := 0; i < l/2; i++ {
		slice1[i], slice1[l-1-i] = slice1[l-1-i], slice1[i]
	}
	return string(slice1)
}

//判断字符串是否全部都不同
//要求：不使用额外存储
//前提条件：字符串的字符为 ASCII 字符，字符长度<=3000

//方法1：strings.Count() 函数判断
func JugeStringIsEqua1(str string) bool {

	if strings.Count(str, "") > 3000 {
		return false
	}

	for _, v := range str {
		if v > 127 {
			return false
		}
		if strings.Count(str, string(v)) > 1 {
			return false
		}
	}
	return true
}

//方法2，使用strings.index()函数判断
func JugeStringIsEqual2(str string) bool {
	if strings.Count(str, "") > 3000 {
		return false
	}
	for i, v := range str {
		if v > 127 {
			return false
		}
		if strings.Index(str, string(v)) != i {
			return false
		}
	}
	return true
}

//字符串替换及字符全部由大小写组成
//将字符串中的空格全部替换成 %20，并且判断字符串是否都是由英文大小写字母组成
//前提条件：该字符串有⾜够的空间存放新增的字符，并且知道字符串的真实⻓度(⼩于等于1000)
func ReplaceStrAndJugeLetter(s string) (string, bool) {
	if len(s) > 1000 {
		return s, false
	}
	for _, v := range s {
		//使用Unicode.IsLetter 判断是否是大小写字母。
		if string(v) != " " && !unicode.IsLetter(v) {
			return s, false
		}
	}
	return strings.Replace(s, " ", "%20", -1), true
}
