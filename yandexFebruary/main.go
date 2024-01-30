package main

import (
	"fmt"
	"strconv"
)

//массив
//RLUPD

//найти индексы в начале и в конце петли

func main() {
	moves := "RLRUUUDDD"
	//0  1  2  3  4  5  6  7  8
	//10 00 10 11 12 13 12 11 10
	fmt.Println(findLoop(moves))
}

func findLoop(moves string) (int, int) {

	minmap := make(map[string]int)
	x := 0
	y := 0
	//resX := 0
	//resY := 0

	maxIndex := 0
	minIndex := 0
	for k, v := range moves {
		if v == 'L' {
			x = x - 1
		}
		if v == 'R' {
			x = x + 1
		}
		if v == 'U' {
			y = y - 1
		}
		if v == 'D' {
			y = y - 1
		}
		strX := strconv.Itoa(x)
		strY := strconv.Itoa(y)
		key := strX + ":" + strY

		if _, ok := minmap[key]; !ok {
			minmap[key] = k
		} else {
			if k > minmap[key] {
				maxIndex = k
				minIndex = minmap[key]
			}
		}
	}

	return minIndex + 1, maxIndex
}
