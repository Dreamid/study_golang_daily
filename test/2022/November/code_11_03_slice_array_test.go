package november

import (
	"log"
	"testing"
)

func TestSliceArray(t *testing.T) {
	a1 := [5]int{1, 2, 3, 4, 5}

	//	slice[low:high:max]
	//	长度：high-low 容量：max-low

	//下表从3开始，长度为1，容量为1。所以s1=[4]
	s1 := a1[3:5:5]

	log.Printf("s1=[%v]", s1)

}

func TestChannel(t *testing.T) {

	n := 10
	channels := make([]chan int, n)
	for i := 0; i < n; i++ {
		channels[i] = make(chan int)
		go func(i int, c1 chan int) {
			log.Println(i)
			c1 <- i
		}(i, channels[i])
	}

	for _, c := range channels {
		num := <-c
		log.Printf("recive :[%v]", num)
	}
}

func TestChannel1(t *testing.T) {

	c1 := make(chan int)
	go func(c1 chan int) {
		c1 <- 1
	}(c1)
	num := <-c1
	log.Printf("recive :[%v]", num)
}
