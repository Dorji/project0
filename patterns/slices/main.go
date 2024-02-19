package main

import "fmt"

type MyError struct {
	Value string
}

func main() {
	arr := make([]string, 0, 4)
	arr = append(arr, "1", "2", "3", "4")
	fmt.Println(arr)
	arraySlice(arr)
	fmt.Println(arr, cap(arr))
}

func arraySlice(arr []string) {
	arr[0] = "5"

	arr = append(arr, "5")
	fmt.Println("inside:", arr, cap(arr))
}
