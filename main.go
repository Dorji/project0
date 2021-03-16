package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(isAnagram("abracadabra", "abracadaarb"))
	fmt.Println(isAnagram("abracadabra", "abracadaarc"))
	fmt.Println(isAnagram("abrac", "dsfdf"))

}
func isAnagram(str0 string, str1 string) bool {

	vocab0 := make(map[string]int)
	vocab1 := make(map[string]int)

	arr0 := strings.Split(str0, "")
	arr1 := strings.Split(str1, "")

	if len(arr0) != len(arr1) {
		return false
	}

	for _, v := range arr0 {
		if _, ok := vocab0[v]; !ok {
			vocab0[v] = 1
		} else {
			vocab0[v]++
		}
	}
	for _, v := range arr1 {
		if _, ok := vocab1[v]; !ok {
			vocab1[v] = 1
		} else {
			vocab1[v]++
		}
	}
	fmt.Println(vocab1)
	fmt.Println(vocab0)
	if reflect.DeepEqual(vocab0, vocab1) {
		return true
	} else {

		return false
	}

}
