package September

import (
	"fmt"
	"testing"
)

func TestStringAppend(t *testing.T) {

	str := "abc" + "123"
	fmt.Printf(str + "\n")

	//The value of ''abc' + '123'' (type rune) cannot be represented by the type string
	//str = 'abc' + '123'

	str = fmt.Sprintf("abc%d", 123)
	fmt.Printf(str + "\n")

	//fmt.Sprintf format %s has arg 123 of wrong type int
	//str = fmt.Sprintf("abc%s", 123)
	//fmt.Printf(str + "\n")

	str = fmt.Sprintf("abc%v", 123)
	fmt.Printf(str + "\n")
	//The value of ''abc' + 'hello'' (type rune) cannot be represented by the type string
	//str = 'abc' + 'hello'

}
