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
