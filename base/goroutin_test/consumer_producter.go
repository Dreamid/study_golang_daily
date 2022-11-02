package goroutin_test

import (
	"fmt"
	"testing"
)

func Consumer(chan int) {

	for i := 0; i < 10; i++ {
		//v := <-chan
	}

}
func Producter(chan bool) {
	//data_chan <- 1
}

func TestConsumerProducter(t *testing.T) {

	//	data_chan、	info_chan
	data_chan, info_chan := make(chan int), make(chan bool)

	Consumer(data_chan)
	Producter(info_chan)

	<-info_chan

	fmt.Println("---end---")
}
