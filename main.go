package main

import "fmt"

func main() {
	var arr = []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	arr = deleteElem(arr, 2) //не забывай добавлять присваивание
	fmt.Println(arr)
}
func deleteElem(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)

}
