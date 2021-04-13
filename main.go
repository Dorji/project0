package main

import "fmt"

func ArrayRotation(arr []int) string {

	start := arr[0]
	res := ""
	arr = append(arr[start:], arr[:start]...)

	for _, v := range arr {
		res = fmt.Sprintf("%s%v", res, v)
	}

	return res

}
func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you

	var arr = []int{2, 3, 4, 1, 6, 10}
	fmt.Println(ArrayRotation(arr))

}
