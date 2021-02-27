package main

import (
	"fmt"
)

func main() {

	var mass = [10]int{2, 3, 42, 5, 6, 8, 9, 333, 55, 6}

	var elem int
	for k := 0; k <= 9; k++ {
		for i := 1; i <= 9; i++ {
			if mass[i] < mass[i-1] {
				elem = mass[i]
				mass[i] = mass[i-1]
				mass[i-1] = elem
			}
		}
	}
	fmt.Println(mass)

}
