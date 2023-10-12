package September

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	tk := time.NewTicker(time.Second * 2)
	channel := make(chan int)
	go func() {
		time.Sleep(time.Second * 7)
		channel <- 1
	}()
	for {
		select {
		case <-channel:
			fmt.Println("complete....")
			return
		case <-tk.C:
			fmt.Println("wait...")
		}
	}
}
