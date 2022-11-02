package november

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person struct {
	Age int `json:"age"`
	//属性首字母小写，不能使用json反序列化赋值
	name  string `json:"name"`
	Child []int  `json:"child"`
}

func TestJsonSlice(t *testing.T) {

	jsonstr := `{"age":18,"name":"julia","child":[1]}`

	a := Person{}
	json.Unmarshal([]byte(jsonstr), &a)
	//append涉及到底层数组扩容，当容量不足的时候， newcap := v.Cap() + v.Cap()/2 ,当newcap小于4时，newcat = 4
	aa := a.Child

	fmt.Printf("aa=[%v],address=[%p]\n", aa, aa)

	jsonstr2 := `{"age":18,"name":"jack","child":[7,8,9,0,1,2,3,4]}`
	json.Unmarshal([]byte(jsonstr2), &a)
	fmt.Printf("aa=[%v],address=[%p]\n", aa, aa)

}
