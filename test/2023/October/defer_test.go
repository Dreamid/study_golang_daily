package October

import (
	"fmt"
	"testing"
)

func TestDeferAndReturn(t *testing.T) {
	fmt.Printf("Func1()=%v\n", Func1())
	fmt.Printf("Func2()=%v\n", Func2())
	fmt.Printf("Func3()=%v\n", Func3())

}

func Func1() (r int) {
	defer func() {
		r = r + 1
	}()
	return 0
}

func Func2() (r int) {
	t := 1
	defer func() {
		t = t + 1
	}()
	return t
}

func Func3() (r int) {
	defer func(r int) {
		r = r + 1
	}(r)
	return 1
}
