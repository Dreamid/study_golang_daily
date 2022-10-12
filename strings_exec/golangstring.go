package golang_strings

import "strings"

// 判断字符串中字符是否全都不同[使用golang内置方法，不使用额外的存储结构]
// 前提：字符串中的字符都是ASCII字符，字符串长<= 3000
// 题解：ASCII字符一共有256个，其中128个是常用字符，可以在键盘上输入；128之后的无法在键盘上找到。

// 思路1：使用strings.index() 下标 == index

func IsUniqueString1(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for i, k := range s {
		if strings.Index(s, string(k)) != i {
			return false
		}
	}
	return true
}

// 思路2：使用strings.Count() 计数 == 1

func IsUniqueString2(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}

	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}
