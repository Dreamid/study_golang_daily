package escape

type Student struct {
	name string
}

func EscapePtr() *Student {
	s := Student{name: "zhangsan"}
	return &s
}

func EscapeMap() *map[string]string {
	m1 := make(map[string]string)
	m1["test"] = "testValue"
	return &m1
}

func EscapeWithOutMap() *map[string]string {
	m1 := make(map[string]string, 5)
	m1["test"] = "testValue"
	return &m1
}
