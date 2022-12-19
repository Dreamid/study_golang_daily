package november

import (
	"fmt"
	"testing"
)

//接口是对非接口值（struct、指针）的封装，其内部包含2个字段，类型T和值V。一个接口值==nil的前提条件是：当且仅当T和V处于unset状态（T=nil，V is unset）
//-两个接口值比较的时，先比较T，在比较V
//-接口值与非接口值比较时，会先将非接口值转换成接口值，再比较。

func TestNil(t *testing.T) {

	var p *int = nil
	//i的内部字段为 T=*int，V=nil
	var i interface{} = p

	//false，i与nil做比较时，会先把nil转换成接口（T=nil，V=nil），与i（T=*int，V=nil）不相等，因此i!=nil。
	fmt.Println(i == nil)
	//ture
	fmt.Println(p == nil)
	//ture
	fmt.Println(i == p)
}
