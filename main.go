package main

import (
	"fmt"
	"math"
)

func main() {

	var ar = []int32{1, 2, 1, 2, 1, 3, 2}

	fmt.Println(sockMerchant(7, ar))

}
func sockMerchant(n int32, ar []int32) int32 {
	//var m map[int32]int32
	colors := make(map[int32]int32)
	for _, v := range ar {
		if _, ok := colors[v]; !ok {
			colors[v] = 1
		} else {
			colors[v] = colors[v] + 1
		}
	}

	var res int32 = 0
	for _, v := range colors {
		res = res + int32(math.Floor(float64(v)/2))
	}
	return res
}
