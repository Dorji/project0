package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var mass = []int{2, 3, 42, 5, 6, 8, 9, 333, 55, 6}

	fmt.Println(quicksort(mass))

}

func quicksort(mass []int) []int {
	if len(mass) < 2 {
		return mass
	}
	left := 0
	right := len(mass) - 1

	pivot := rand.Int() % len(mass)

	tmp := 0
	tmp = mass[pivot]
	mass[pivot] = mass[right]
	mass[right] = tmp

	for i, _ := range mass {
		if mass[i] < mass[right] {
			tmp = mass[left]
			mass[left] = mass[i]
			mass[i] = tmp
			left++
		}
	}
	tmp = mass[left]
	mass[left] = mass[right]
	mass[right] = tmp

	quicksort(mass[:left])
	quicksort(mass[left+1:])
	return mass
}
