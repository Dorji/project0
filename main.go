package main

import "fmt"

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you

	var arr []*int
	var arrSrc = []int{2, 3, 4, 1, 6, 10}

	for k := range arrSrc {
		arr = append(arr, &arrSrc[k])
	}

	for _, v := range arr {
		fmt.Println(*v)
	}
}
