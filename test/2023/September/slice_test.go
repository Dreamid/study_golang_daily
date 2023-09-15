package September

import (
	"fmt"
	"sort"
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

func TestSortAndFind(t *testing.T) {

	target := "golang"

	src := []string{"hello", "Golang", "yes", "GO"}

	ok := SortAndFind(target, src)
	if ok {
		t.Log("contains \n")
		return
	}
	t.Log("not contains\n")

}

// sort.Strings() : 对字符串数组进行排序
// sort.SearchStrings : 使用二分法在一个有序字符串数组中寻找目标字符串索引
func SortAndFind(target string, src []string) bool {
	sort.Strings(src)
	index := sort.SearchStrings(src, target)
	if index <= len(src) && src[index] == target {
		return true
	}
	return false
}

type Student struct {
	name string
}

func TestSliceAndList(t *testing.T) {

	s1 := Student{"zhangsan"}

	slice1 := make(map[string]Student)
	slice2 := make(map[string]*Student)

	fmt.Printf("%v\n", slice1)
	fmt.Printf("%t\n", slice2)

	slice1["s1"] = s1
	slice2["s1"] = &s1

	//Cannot assign to slice1["s1"].name
	//slice1["s1"].name = "lisi"

	slice2["s1"].name = "wangwu"

}
