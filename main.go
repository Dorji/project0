package main

import "fmt"

// дан целочисленный массив и целое число (target). вернуть индексы двух чисел, сумма которых равна target
// Input: nums = [3,9,11,15], target = 14
// Output: [0,2]

func main() {

	nums := []int{3, 9, 11, 15}
	target := 14

	numsMap := make(map[int]int)

	for k := range nums {
		numsMap[nums[k]] = k
	}
LOOP:
	for k := range nums {
		if _, ok := numsMap[target-nums[k]]; ok {
			fmt.Println(k, numsMap[target-nums[k]])
			break LOOP
		}
	}

}

//1,2,3,9,11,13,14,15
//14 0

// O(n^2)
func main2() {

	nums := []int{9, 11, 3, 15}
	target := 14
	//mem:= make(map[int]int)

LOOP:
	for k := range nums {
		for key := range nums {
			if key == k {
				continue
			}
			if nums[k]+nums[key] == target {
				fmt.Println(k, key)
				break LOOP
			}
		}

	}

}
