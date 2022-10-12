package test

import (
	"fmt"
	"sync"
)

func main() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}

		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		i := 0
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for {
			select {
			case <-letter:
				if i > len(str)-1 {
					wait.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				number <- true
			default:
				break
			}
		}
	}(&wait)
	number <- true
	wait.Wait()

}
