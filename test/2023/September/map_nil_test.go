package September

import "testing"

func TestMapWithNilKey(t *testing.T) {

	type Person struct {
		name string
	}
	m := make(map[Person]int)
	m2 := make(map[Person]string)

	p1 := Person{"Jacy"}

	//m.[p1]=0
	t.Logf("m.[p1]=%v\n", m[p1])
	//m.[p1]=
	t.Logf("m.[p1]=%v\n", m2[p1])
}
