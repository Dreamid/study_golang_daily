package September

import (
	"fmt"
	"testing"
)

const s1 = "hello"

func TestConst(t *testing.T) {
	var s2 = "world"
	fmt.Printf("s2 address = %p\n", &s2)
	//Cannot take the address of 's1'
	//fmt.Sprintf("%p", &s1)
}
