package main

import (
	"fmt"
	"sync"
)

//Надо поправить код, чтобы 10 раз вывело FooBar на новой строке. Нельзя убирать два вызова горутин и циклы в функциях, сигнатуры менять можно:
var mu1 = sync.Mutex{}
var mu2 = sync.Mutex{}

func foo(n int, mu1 sync.Mutex, mu2 sync.Mutex) {
	for i := 0; i < n; i++ {
		mu1.Lock()
		fmt.Println("Foo")
		mu2.Unlock()
	}
}

func bar(n int, mu1 sync.Mutex, mu2 sync.Mutex) {
	for i := 0; i < n; i++ {
		mu2.Lock()
		fmt.Println("Bar")
		mu1.Unlock()
	}
}

func main() {
	n := 10
	mu2.Lock()
	go foo(n, mu1, mu2)
	go bar(n, mu1, mu2)
}
