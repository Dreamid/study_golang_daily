package lock

import (
	"fmt"
	"time"
)

func DeadLock() {

	ch1 := make(chan string)
	//错误使用方法
	//go 语句后面的函数调用，其参数会先求值
	//go fmt.Println(<-ch1)

	//正确使用方法
	go func() {
		fmt.Println(<-ch1)
	}()

	ch1 <- "hello"
	time.Sleep(10)
}
