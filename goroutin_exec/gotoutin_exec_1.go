package goroutin_exec

import (
	"fmt"
	"sync"
)

func PrintCrosss() {
	// 1. 定义两个channel，用户控制goroutin
	number, letter := make(chan bool), make(chan bool)
	// 2. 定义 waitGroup，用户控制goroutin结束
	wait := sync.WaitGroup{}
	// 2. 需要输出的字符串
	strs := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 3. 启动goroutin
	go func() {
		i := 0
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()
	go func(wait *sync.WaitGroup) {
		i := 0
		for {
			select {
			case <-letter:
				if i > len(strs)-1 {
					wait.Done()
					return
				}
				fmt.Print(strs[i : i+1])
				i++
				fmt.Print(strs[i : i+1])
				i++
				number <- true
			}
		}
	}(&wait)
	wait.Add(1)
	number <- true
	wait.Wait()

}
