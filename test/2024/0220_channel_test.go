package test

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	c1 := make(chan int, 10)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second * time.Duration(rand.Intn(1)))
			c1 <- i
		}(i)
	}
	wg.Wait()
	close(c1)

	for v := range c1 {
		t.Logf("v = %v", v)
	}

	//for {
	//	v, ok := <-c1
	//	t.Logf("v = %v,ok=%v", v, ok)
	//	if !ok {
	//		t.Log("channel is closed")
	//		break
	//	}
	//}
	t.Logf("=========end============")
}

func TestClosetChannelWithContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go RunTaskFunc(&wg, ctx, "A", func(ctx context.Context) {
		wg.Add(1)
		go RunTaskFunc(&wg, ctx, "B", func(ctx context.Context) {
			wg.Add(1)
			go RunTaskFunc(&wg, ctx, "C", func(ctx context.Context) {
			})
		})
	})
	wg.Add(1)
	go RunTaskFunc(&wg, ctx, "D", func(ctx context.Context) {
		wg.Add(1)
		go RunTaskFunc(&wg, ctx, "E", func(ctx context.Context) {
			wg.Add(1)
			go RunTaskFunc(&wg, ctx, "F", func(ctx context.Context) {
			})
		})
	})
	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()
	wg.Wait()
}

type TaskFunc func(ctx context.Context)

func RunTaskFunc(wg *sync.WaitGroup, ctx context.Context, name string, taskFunc TaskFunc) {
	defer wg.Done()
	fmt.Printf("%v started...\n", name)
	taskFunc(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%v exit...\n", name)
			return
		}
	}
}
