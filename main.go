package main

import (
	"github.com/casbin/casbin/v2"
	"log"
)

func main() {
	a, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalln(err)
	}
	//a.AddFunction("my_func", KeyMatchFunc)
	Check(a, "dajun", "data1", "read")
	Check(a, "dajun", "data1", "write")
	Check(a, "dajun", "data2", "read")
	Check(a, "lizi", "data2", "write")
}

func Check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		log.Printf("%s can %s %s \n", sub, act, obj)
	} else {
		log.Printf("%s can't %s %s \n", sub, act, obj)
	}
}

//
//func KeyMatch(key1 string, key2 string) bool {
//	i := strings.Index(key2, "*")
//	if i == -1 {
//		return key1 == key2
//	}
//	if len(key1) > i {
//		return key1[:i] == key2[:i]
//	}
//	return key1 == key2[:i]
//}
//
//func KeyMatchFunc(args ...interface{}) (interface{}, error) {
//	name1 := args[0].(string)
//	name2 := args[1].(string)
//	return (bool)(KeyMatch(name1, name2)), nil
//}
