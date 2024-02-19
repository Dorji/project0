package main

import "fmt"

type Obj struct {
	foo    string
	bar    int
	bararr []string
}

func main() {
	val := Obj{"sdsd", 23, []string{"sss", "ddd"}}
	fmt.Println(val)
	foo(val)
	fmt.Println(val)
}

func foo(obj Obj) Obj {
	obj.foo = "qwe"
	obj.bararr[0] = "qwe"
	return obj
}
