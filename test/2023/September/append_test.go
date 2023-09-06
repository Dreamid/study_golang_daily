package September

import (
	"fmt"
	"testing"
)

/*
*
append 函数，不能接受指针类型。
*/
func TestAppend(t *testing.T) {

	s1 := new([]int)
	//0xc0000080d8
	fmt.Printf("%p\n", s1)
	//s1的类型为 ： &[]
	fmt.Printf("%v\n", s1)

	//Cannot use 's1' (type *[]int) as the type []Type
	//s1 = append(s1, 1)

	var b1 bool
	//%t 布尔占位符
	fmt.Printf("%t\n", b1)

}
