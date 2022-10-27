package singleton

import (
	"sync"
	"sync/atomic"
)

type singleton struct{

}
var ins *singleton
var one *sync.Once

var initailized uint32

//使用 sync.Once 来控制初始化
func GetIns()*singleton{
	one.Do(func() {
		ins = &singleton{}
	})
	return ins
}

//使用编译器优化检查存储状态
func GetInsByAotomic()*singleton{
	if atomic.LoadUint32(&initailized) == 1{
		return ins
	}
	mu.Lock()
	defer mu.Unlock()
	ins = &singleton{}
	atomic.StoreUint32(&initailized,1)
	return ins
}