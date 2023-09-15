package September

import (
	"fmt"
	"testing"
)

func TestRuneAppend(t *testing.T) {

	r := []rune("hello,world!")
	s := []byte("hello,world!")

	fmt.Println(r) //[104 101 108 108 111 44 119 111 114 108 100 33]
	fmt.Println(s) //[104 101 108 108 111 44 119 111 114 108 100 33]
	fmt.Println("---------------")
	txt := "你好，世界！"
	fmt.Println([]rune(txt)) //[20320 22909 65292 19990 30028 65281]
	fmt.Println([]byte(txt)) //[228 189 160 229 165 189 239 188 140 228 184 150 231 149 140 239 188 129]
	fmt.Println("---------------")

	//	如果截取中文字符串？
	//中文利用 [] rune 转换成 unicode 码点 []rune(txt)
	//再利用 string 转化回去 string()
	fmt.Println(string([]rune(txt)[:1])) //你
	fmt.Println(string([]byte(txt)[:1])) //�

}
