package gip

import (
	"fmt"
	"sync"
)

//交替打印字母和数字
func CrossPrintNumStr() {
	//要打印的字符串
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//声明两个channel，用于控制信号量传输。
	num, letter := make(chan bool), make(chan bool)
	//goroutin管理组
	wg := sync.WaitGroup{}
	//打印数组
	go func() {
		i := 0
		for {
			select {
			case <-num:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			default:
				//continue
				break
			}
		}
	}()
	go func(group *sync.WaitGroup) {
		i := 0
		for {
			select {
			case <-letter:
				if i >= len(str)-1 {
					wg.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				fmt.Print(str[i : i+1])
				i++
				num <- true
			default:
				//continue
				break
			}
		}
	}(&wg)
	wg.Add(1)
	num <- true
	wg.Wait()
	fmt.Println("\n----end----")
}
