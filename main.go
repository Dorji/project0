package main

import (
	"fmt"
	"strconv"
	"strings"
)

// RLE encoding
// Input: aaabbcaac
// Output: 3a 2b 1c 2a 1c
//-> pairs {counter, symbol}

func main() {

	var str = "aaabbcaac"
	foo := strings.Split(str, "")

	res := ""
	tmpStr := ""
	tmpInt := 0

	fmt.Println(foo)
	tmpRes := ""
	ln := len(foo) - 1
	for k, v := range foo {
		if tmpStr != v {
			if tmpInt != 0 {
				tmpRes = strconv.Itoa(tmpInt) + tmpStr
				res = res + tmpRes + " "
			}
			tmpStr = v
			tmpInt = 0

		}

		tmpStr = v
		tmpInt = tmpInt + 1

		if ln == k {
			tmpRes = strconv.Itoa(tmpInt) + tmpStr
			res = res + tmpRes + " "
		}
	}

	fmt.Println(res)
}
