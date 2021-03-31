package main

import (
	"fmt"
	"strings"
)

// Задана строка S из малых латинских букв, требуется узнать длину
// наибольшей подстроки состоящей не более чем из двух различных символов.

/*
"aaaa" - 4
"asasas" - 6
"asc" - 2


"asdddddasas" 4

0 1           [a]
1 2           [a,s]
2 1           [s,d]
3 2           [d,a]
4 1           [a,s]
5 2           [a,s]
6 3           [a,s]
3+1

"abccccccdeffffggfgfgfgfgfgffggggg"
0 1     [a=>[0]]
1 1     [a=>[0],b=>[1]]
2 1     [b=>[1],c=>[2]]
3 1     [c=>[2],d=>[3]]
*/

// "abcbcbcbcbd"
func main() {
	fmt.Println(foo("01234567891"))
	fmt.Println(" 0 1 2 3 4 5 6 7 8 9 10")
	fmt.Println(" a s d d d d d a s a s")
	fmt.Println(foo("asdddddasas"))
}

func foo(str string) int {
	tmp := make(map[string]int)
	arr := strings.Split(str, "")
	cnt := make([]int, len(arr))
	var counter int
	counter = 0
	for k, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = k

			if len(tmp) > 2 {

				counter = counter - cnt[tmp[minInMap(tmp)]] - 1
				delete(tmp, minInMap(tmp))
			} else {
				counter = counter - 1
			}

			if counter < 1 {
				counter = 1
			}
			cnt[k] = counter

		} else {
			counter = counter + 1
			cnt[k] = counter
			tmp[v] = k
		}

	}
	fmt.Println(cnt)
	return maxInSlice(cnt)
}

//return key
func minInMap(mapa map[string]int) string {
	min := 0
	minKey := ""
	for k, v := range mapa {
		if min == 0 {
			min = v
			minKey = k
		}
		if v < min {
			min = v
			minKey = k
		}
	}
	return minKey
}

// return value
func maxInSlice(slc []int) int {
	max := 0
	for _, v := range slc {
		if v > max {
			max = v
		}
	}
	return max
}
