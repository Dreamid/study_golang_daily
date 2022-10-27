package ge

import (
	"flag"
	"fmt"
)

//定义命令行参数对应的变量，这三个变量都是指针类型。
var cliName = flag.String("name", "nick", "Input Your Name")
var cliAge = flag.Int("age", 28, "Input Your Age")
var cliGender = flag.String("gender", "male", "Input Your Gender")

//定义值类型的命令行参数变量，在Init（）函数对其初始化。
//即：命令参数对应变量的 定义和初始化，是可以分开的。
var cliFlag int

func Init() {
	flag.IntVar(&cliFlag, "flagname", 1234, "Just for demo")
}

func InitFlag() {
	Init()
	flag.Parse()
	// flag.Args() 函数返回没有被解析的命令行参数
	// func NArg() 函数返回没有被解析的命令行参数的个数
	fmt.Printf("args=%s,num=%d\n", flag.Args(), flag.NArg())
	fmt.Println("-------------------")
	for i := 0; i < flag.NFlag(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("-------------------")

	//	输出命令行参数
	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", *cliGender)
	fmt.Println("cliFlag=", cliFlag)
}
