package cond

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 使用cond来处理 通知|等待 场景下的并发问题
func TestCond(t *testing.T) {

	c := sync.NewCond(&sync.Mutex{})
	finishNum := 0
	notifyNum := 0
	for i := 1; i <= 100; i++ {
		go func(index int) {
			time.Sleep(time.Millisecond * time.Duration(rand.Int63n(1000)))
			c.L.Lock()
			finishNum++
			t.Logf("第%v位运动员完成了准备...", index)
			c.L.Unlock()
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for finishNum != 100 {
		notifyNum++
		t.Logf("裁判员被唤醒%v次", notifyNum)
		c.Wait()
	}
	c.L.Unlock()
	t.Log("所有运动员完成了准备...")
}
