package november

import "testing"

var ch chan string

func TestForRange(t *testing.T) {
	ch = make(chan string)
	go func() {
		//ch <- "Go"
		//ch <- "语言"
		//ch <- "高性能"
		//ch <- "编程"
		close(ch)
	}()

	for i := range ch {
		t.Log(i)
	}

}
