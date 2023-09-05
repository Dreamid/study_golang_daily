package September

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 模拟100个运动员准备
func TestCond(t *testing.T) {
	con := sync.NewCond(&sync.Mutex{})
	finishedNum := 0
	notifyNum := 0
	for i := 0; i < 100; i++ {
		go func(index int) {
			time.Sleep(time.Millisecond * time.Duration(rand.Int63n(10000)))
			con.L.Lock()
			finishedNum++
			fmt.Printf("第 %d 名运动员完成了准备\n", index)
			con.L.Unlock()
			con.Broadcast()
		}(i)
	}
	con.L.Lock()
	for finishedNum != 100 {
		notifyNum++
		fmt.Printf("裁判员被唤醒 %d 次\n", notifyNum)
		con.Wait()
	}
	con.L.Unlock()
}
