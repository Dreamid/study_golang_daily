package September

import "testing"

type People interface {
	Say()
}
type Man struct{}

func (stu *Man) Say() {}

func Live() People {
	var man Man
	return &man
}

func TestInterface(t *testing.T) {
	obj := Live()
	if obj == nil {
		t.Log("nil")
	} else {
		//ob = (September.People|*September.Man)
		t.Log("not nil")
	}
}
