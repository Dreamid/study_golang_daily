package October

import (
	"fmt"
	"testing"
)

/**
%d 表示输出十进制数字
+ 表示输出数值的符号
*/

func TestPlaceHolder(t *testing.T) {

	a := -1
	b := 1
	fmt.Printf("%+d\n", a)
	fmt.Printf("%+d\n", b)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", b)
}
