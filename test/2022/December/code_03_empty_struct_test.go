package november

import (
	"fmt"
	"testing"
	"unsafe"
)

// 空结构体不占据任何内存空间
func TestEmptyStructSize(t *testing.T) {
	//0
	fmt.Println(unsafe.Sizeof(struct{}{}))
	//1
	fmt.Println(unsafe.Sizeof(true))
	//8
	fmt.Println(unsafe.Sizeof(1))
	//16
	fmt.Println(unsafe.Sizeof(""))

}

// Go中没有提供Set数据结构，但是map配合空struct可以实现Set数据结构的效果
type Set map[string]struct{}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}
func (s Set) Del(key string) {
	delete(s, key)
}
func (s Set) IsExist(key string) bool {
	_, ok := s[key]
	return ok
}

func TestSet(t *testing.T) {
	set := make(Set)
	set.Add("tom")
	t.Log(set.IsExist("tom"))
	set.Del("tom")
	t.Log(set.IsExist("tom"))
}

// 使用struct{}作为不发送数据的信道，channel不需要发送任何数据，只用来通知子协程执行任务，或只用来控制协程并发度，这种情况下，使用空结构体作为占位符就非常合适了。
func TestChannelWithStruct(t *testing.T) {
	t.Log("--start main---")
	flag := make(chan struct{})
	go DoWork("worker1", flag)
	flag <- struct{}{}
	t.Log("---end main---")
}

func DoWork(name string, flag chan struct{}) {
	fmt.Printf("***%v start***\n", name)
	<-flag
	fmt.Println("***close channel***")
	close(flag)
}

// 使用struct{}用来表示仅包含方法的结构体
// 事实上，Door 可用任何数据结构替代，例如 type Door int； type Door bool
// 无论是int or bool，都会浪费额外的内存，在这种情况下，声明为空结构体是最合适的。
type Door struct{}

func (d Door) Open() {
	fmt.Println("open the door")
}
