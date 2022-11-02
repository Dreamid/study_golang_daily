package november

import (
	"log"
	"sync"
	"sync/atomic"
	"testing"
)

var wg sync.WaitGroup

func TestWaitGroup(t *testing.T) {
	total, sum := int64(0), 0
	n := 100000
	wg.Add(n)

	for i := 1; i <= n; i++ {
		sum += i
		go func(v int) {
			atomic.AddInt64(&total, int64(v))
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	log.Printf("total=[%v],suml=[%v]\n", total, sum)
}
