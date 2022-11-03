package november

import (
	"log"
	"testing"
)

func TestSliceArray(t *testing.T) {
	a1 := [5]int{1, 2, 3, 4, 5}

	//	slice[low:high:max]
	//	长度：high-low 容量：max-low

	//下表从3开始，长度为1，容量为1。所以s1=[4]
	s1 := a1[3:5:5]

	log.Printf("s1=[%v]", s1)

}
