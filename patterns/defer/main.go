package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer func")
		if res := recover(); res != nil {
			fmt.Println(res)
		}
	}()
	func() {
		panic(365)
	}()
	fmt.Println("123")
}
