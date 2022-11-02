package goroutin_test

import (
	"fmt"
	"sync"
)

func CrossPrint() {
	number, letter := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	go func() {
		i := 0
		for {
			switch {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				letter <- true
			default:
				break

			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		i := 0
		for {
			select {
			case <-letter:
				if i > len(letters)-1 {
					wait.Done()
					return
				}
				fmt.Print(letters[i : i+1])
				i++
				fmt.Print(letters[i : i+1])
				i++
				number <- true
			default:
				break
			}
		}
	}(&wait)
	number <- true
	wait.Wait()
	fmt.Println()
	fmt.Println("end")
}
