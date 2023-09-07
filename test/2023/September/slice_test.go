package September

import (
	"fmt"
	"testing"
)

// s1 is nil
// s2 is not nil
func TestSlice(t *testing.T) {
	var s1 []int
	var s2 = []int{}

	//s1 is nil
	//s2 is not nil
	if s1 == nil {
		fmt.Println("s1 is nil")
	} else {
		fmt.Println("s1 is not nil")
	}
	if s2 == nil {
		fmt.Println("s2 is nil")
	} else {
		fmt.Println("s2 is not nil")
	}
}

/*
*

	golang 中的 ... 有2种作用

1. 作为函数的可变参数
2. 将slice的元素打散传递
*/
func TestPointPointPoint(t *testing.T) {

	s1 := []string{"hello", "world", "golang"}
	s2 := [3]string{"hello", "world", "golang"}

	fmt.Printf("%T\n", s1)
	fmt.Printf("%T\n", s2)

	Print(s1...)
	fmt.Println("----------")
	//Cannot use 's2' (type [3]string) as the type []string
	//Print(s2...)

}

func Print(args ...string) {
	for i, v := range args {
		fmt.Printf("i = %v,v = %v\n", i, v)
	}
}
