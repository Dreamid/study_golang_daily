package test

import (
	"fmt"
	"testing"
)

/*
*
Go 提供的语法糖...，可以将 slice 传进可变函数，不会创建新的切片。第一次调用 change() 时，append() 操作使切片底层数组发生了扩容，原 slice 的底层数组不会改变； 第二次调用change() 函数时，使用了操作符[i,j]获得一个新的切片，假定为 slice1，
它的底层数组和原切片底层数组是重合的，不过 slice1 的长度、容量分别是 2、5，所以在 change() 函数中对 slice1 底层数组的修改会影响到原切片。
*/
func TestSliceAppend(t *testing.T) {

	s := make([]int, 5, 5)
	s[0] = 1
	s[1] = 2
	t.Logf("%v\n", s)

	t.Log("------change-----\n")
	Change(s...)
	t.Logf("s = %v", s)

	t.Log("------change-----\n")
	Change(s[:2]...)
	t.Logf("s=%v", s)
}

func Change(s ...int) {
	s = append(s, 3)
	fmt.Printf("s = %v\n", s)
}
