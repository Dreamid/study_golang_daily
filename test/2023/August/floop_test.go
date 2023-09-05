package August

import (
	"fmt"
	"testing"
)

func TestFloop(t *testing.T) {
	type Foo struct {
		name string
	}
	s1 := []Foo{{"A"}, {"B"}, {"C"}}
	fmt.Printf("%v,%v,%v", s1[0], s1[1], s1[2])
	s2 := make([]*Foo, len(s1))
	for i, v := range s1 {
		s2[i] = &v
	}
	fmt.Printf("%v,%v,%v", s2[0], s2[1], s2[2])

}
