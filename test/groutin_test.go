package test

import (
	"golang_demo/golang_study/goroutin_exec"
	"testing"
)

func TestGoroutin(t *testing.T) {
	goroutin_exec.GoroutinPrint()
	t.Log("---------end---------")
}
