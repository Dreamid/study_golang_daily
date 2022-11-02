package main

import (
	"log"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
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
