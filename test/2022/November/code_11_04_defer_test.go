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
}
