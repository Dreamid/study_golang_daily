package test

import (
	"testing"
)

func Test0226(t *testing.T) {

	p := f()

	t.Log(p.fi)
}

type S struct {
	fi string
}

func f() *S {
	return &S{fi: "hello"}
}

func TestStr(t *testing.T) {
	t.Logf("hello \r\n world")
}
