package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

func main() {
	fmt.Println(utf8.RuneCountInString(strconv.Itoa(144444)))
	fmt.Printf("sum is: %v", DigitalRoot(195))
}

func DigitalRoot(n int) int {
	var sum int = 0
	var slc = IntToSlice(n)
	for _, v := range slc {
		sum += v
	}
	if sum/10 > 0 {
		sum = DigitalRoot(sum)
	}

	return sum
}
func IntToSlice(n int) []int {
	var number = n
	var lenN = utf8.RuneCountInString(strconv.Itoa(n))
	res := make([]int, 0, lenN)
	var elem0 int
	var elem1 int

	for i := 1; i < lenN+1; i++ {
		elem0 = number % int(math.Pow(float64(10), float64(i)))
		if i != 1 {

			elem1 = elem0 / int(math.Pow(float64(10), float64(i-1)))
		} else {
			elem1 = elem0
		}
		res = append(res, elem1)
	}
	return res
}
