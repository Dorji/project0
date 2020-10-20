package main

import "fmt"

func main() {

	//fmt.Println(maxIntExcept([]int{6,5,83,5,3,18}))
	fmt.Println(TwoOldestAges([]int{61, 11, 33, 79, 81, 27, 79, 83, 9, 95}))
}
func TwoOldestAges(ages []int) [2]int {
	var res [2]int
	slice := make([]int, len(ages), len(ages))
	slice = ages
	maxOne := maxIntExcept(slice)

	slice = remove(slice, maxOne[1])

	maxTwo := maxIntExcept(slice)
	res[0] = maxTwo[0]
	res[1] = maxOne[0]
	return res
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func maxIntExcept(slc []int) [2]int {
	var max int
	var maxKey int
	max = slc[0]
	for k, v := range slc {

		if v > max {

			max = v
			maxKey = k

		}
	}
	return [2]int{max, maxKey}
}
