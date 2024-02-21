package test

import (
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {

	type T1 int

	var a T1 = 1

	at := reflect.TypeOf(a)
	t.Logf("name = %v\n", at.Name())
	t.Logf("kind = %v\n", at.Kind())

	t.Log("================")

	type user struct {
		name string
		age  uint32
	}

	type Student struct {
		user
		school string
	}

	s1 := Student{}

	sT := reflect.TypeOf(s1)

	for i := 0; i < sT.NumField(); i++ {
		f := sT.Field(i)
		t.Logf("name=%v,type=%v,offset=%v,index=%v\n", f.Name, f.Type, f.Offset, f.Index)

		if f.Anonymous {
			//匿名结构体
			for y := 0; y < f.Type.NumField(); y++ {
				f1 := f.Type.Field(y)
				t.Logf("name=%v,type=%v,offset=%v,index=%v\n", f1.Name, f1.Type, f1.Offset, f1.Index)
			}
		}
	}

}
