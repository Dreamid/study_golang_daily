package november

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup
var stopFlag chan bool

func TestReflectDeep(t *testing.T) {
	str1 := "hello"
	str2 := "hello"

	t.Log(str1 == str2)

	s3 := []string{"hello", "golang"}
	s4 := []string{"hello", "golang"}

	t.Log(reflect.DeepEqual(s3, s4))
}

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go DoTask(i + 1)
	}
	wg.Wait()
	fmt.Println("All task Done")

}

func DoTask(n int) {
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Printf("Task %d Done\n", n)
	wg.Done()
}

// 使用select+chan 主动通知子协程停止
func TestSelectChan(t *testing.T) {
	stopFlag = make(chan bool)
	go DoSubTask("waiter1")
	time.Sleep(3 * time.Second)
	stopFlag <- true
	time.Sleep(3 * time.Second)

}

func DoSubTask(name string) {
	for {
		select {
		case <-stopFlag:
			fmt.Printf("%v stop \n", name)
			return
		default:
			fmt.Printf("%v send request \n", name)
			time.Sleep(time.Second * 1)
		}
	}
}

//使用context完成 select+ chan 的功能，即主动通知子协程退出。

func TestContextCancle(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go DoSubTaskCancle(ctx, "woker1")
	go DoSubTaskCancle(ctx, "woker2")
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}

func DoSubTaskCancle(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%v stop\n", name)
			return
		default:
			fmt.Printf("%v send request\n", name)
			time.Sleep(time.Second * 1)
		}
	}
}

type Options struct {
	Interval time.Duration
}

func TestContextWithValue(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "options", Options{1})
	go DoSUbTaskWithValue(ctx, "woker1")
	go DoSUbTaskWithValue(ctx, "woker2")
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func DoSUbTaskWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%v stop\n", name)
			return
		default:
			fmt.Printf("%v send request\n", name)
			opt := ctx.Value("options").(Options)
			time.Sleep(opt.Interval * time.Second)
		}
	}
}
