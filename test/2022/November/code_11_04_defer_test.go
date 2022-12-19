package november

import "testing"

func getValue1() int {
	i := 1
	defer func() {
		i++
	}()
	return i
}

func getValue2() int {
	i := 1
	defer func(j int) {
		j++
	}(i)
	return 1
}

func getValue3() (i int) {
	i = 1
	defer func() {
		i++
	}()
	return
}

func TestDeferModifyReturenValue(t *testing.T) {
	t.Logf("value1=[%v]", getValue1())
	t.Logf("value2=[%v]", getValue2())
	t.Logf("value3=[%v]", getValue3())

	//go中，数组长度也是数组类型的一部分，所以 [5] int 和 [10] int 是属于不同类型的
	a1 := [3]int{1, 2, 3}
	a2 := [4]int{1, 2, 3}
	a3 := [4]int{1, 2, 3}
	t.Logf("a1=[%v]", a1)
	t.Logf("a2=[%v]", a2)
	if a2 == a3 {
		t.Log("a2 == a3")
	} else {
		t.Log("a2 != a3")
	}

	//if a1 == a2 {
	//
	//}

	//compilation error
}
