package test

import (
	"fmt"
	"testing"
)

//新类型
type MyInt1 int

//别名
type MyInt2 = int

func TestTypeAlias(t *testing.T) {

	var i int = 1
	//Go是强类型语言，需要强制转换。
	//var i1 MyInt1 = i //错误示范
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)

}
