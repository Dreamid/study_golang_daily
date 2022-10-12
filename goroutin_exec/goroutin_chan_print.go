package goroutin_exec

import (
	"fmt"
	"sync"
)

func GoroutinPrint() {
	letter, number := make(chan bool), make(chan bool)
	str1 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	waitGroup := sync.WaitGroup{}
	//执行goroutin，打印数字。
	go func() {
		i := 0
		for {
			switch {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			default:
				break
			}
		}
	}()
	waitGroup.Add(1)
	go func(group *sync.WaitGroup) {
		i := 0
		for {
			switch {
			case <-letter:
				if i >= len(str1)-1 {
					waitGroup.Done()
					return
				}
				fmt.Print(str1[i : i+1])
				i++
				fmt.Print(str1[i : i+1])
				number <- true
			default:
				break
			}
		}
	}(&waitGroup)
	number <- true
	waitGroup.Wait()

}
