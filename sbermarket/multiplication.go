package main

import (
	"fmt"
	"math"
)

func main() {
	var array = []int{0, 0, 0, 1, 1, 10, -9}
	positiveCnt := -1
	hasZero := 0
	for _, v := range array {
		if v > 0 {
			positiveCnt = v
		} else if v == 0 {
			hasZero = 1
		}
	}
	if positiveCnt > 0 {
		fmt.Println(hasPositive(array, hasZero, 1))
	} else {
		fmt.Println(allNegative(array))
	}

}

func hasPositive(array []int, hasZero int, hasPositive int) int {
	arrayPointer := &array
	mini := array[0]
	for _, v := range array {
		if mini >= v {
			mini = v
		}
	}
	positiveMax := -1
	if hasPositive == 1 {
		for _, v := range array {
			if positiveMax <= v {
				positiveMax = v
			}
		}
	}

	max := make([]int, 0, 0)
	maxKey := -1
	for i := 0; i < 3; i++ {

		for key, v := range *arrayPointer {
			if v != 0 {
				if len(max) == i {
					max = append(max, v)
				}
				if math.Abs(float64(max[i])) <= math.Abs(float64(v)) {
					max[i] = v
					maxKey = key
				}
			} else {
				hasZero = 1
			}
		}
		if maxKey != len(*arrayPointer)-1 {
			*arrayPointer = append(array[:maxKey], array[maxKey+1:]...)
		} else if maxKey != -1 {
			*arrayPointer = array[:maxKey]
		}
	}

	singleNegativeKey := 0
	negativeCnt := 0
	for k, v := range max {
		if v < 0 {
			negativeCnt = negativeCnt + 1
			singleNegativeKey = k
		}
	}

	if negativeCnt == 1 {
		maxPositiveForSingleNegative := 0
		for _, v := range array {
			if maxPositiveForSingleNegative < v {
				maxPositiveForSingleNegative = v
			}
		}
		max[singleNegativeKey] = maxPositiveForSingleNegative
	}

	if max[0] != positiveMax && max[1] != positiveMax && max[2] != positiveMax {
		minInMax := max[0]
		minInMaxKey := -1
		for k, v := range max {
			if math.Abs(float64(minInMax)) >= math.Abs(float64(v)) {
				minInMax = v
				minInMaxKey = k
			}
		}
		max[minInMaxKey] = positiveMax
	}

	if hasZero == 0 {
		if max[0]*max[1]*max[2] > 0 {
			return max[0] * max[1] * max[2]
		} else {
			return 0
		}
	}
	return max[0] * max[1] * max[2]
}

func allNegative(array []int) int {
	mini := array[0]
	for _, v := range array {
		if mini >= v {
			mini = v
		}
	}
	var min = []int{mini, mini, mini}
	minK := -1
	for i := 0; i < 3; i++ {
		for k, v := range array {
			if math.Abs(float64(min[i])) >= math.Abs(float64(v)) {
				min[i] = v
				minK = k
			}
		}

		if minK != len(array)-1 {
			array = append(array[:minK], array[minK+1:]...)
		} else {
			array = array[:minK]
		}

	}

	return min[0] * min[1] * min[2]
}
