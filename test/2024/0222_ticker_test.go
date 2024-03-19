package test

import (
	"context"
	"testing"
	"time"
)

func TestGoroutin1(t *testing.T) {
	ctx, calcel := context.WithTimeout(context.Background(), time.Second*5)
	tk := time.Tick(time.Second)
	go func() {
		time.Sleep(time.Second * 2)
		calcel()
	}()
	for {
		select {
		case <-tk:
			t.Log("111")
		case <-ctx.Done():
			t.Log("ctx down")
			//break //无法退出 for 循环
			return
		}
	}
}
