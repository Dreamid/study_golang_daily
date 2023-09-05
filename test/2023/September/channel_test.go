package September

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestChannel(t *testing.T) {
	for i := 0; i < 10; i++ {
		//	1.声明多个channel
		channels := [3]chan int{
			make(chan int, 1),
			make(chan int, 1),
			make(chan int, 1),
		}
		//	随机向某个channel中写入数据
		num := rand.Intn(3)
		fmt.Printf("num = %d\n", num)
		channels[num] <- num
		//	从channel中取数据
		select {
		case <-channels[0]:
			fmt.Println("the first candidate is selected")
		case <-channels[1]:
			fmt.Println("the second candidate is selected")
		case <-channels[2]:
			fmt.Println("the third candidate is selected")
		}
	}

}
