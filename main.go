package main

import "fmt"

func main() {
	Solution("abc")                 //should return ["ab", "c_"]
	fmt.Println(Solution("abcdef")) //should return ["ab", "cd", "ef"]
}

func Solution(str string) []string {
	res := make([]string, 0, len(str))
	pair := ""
	cnt := len(str)
	if cnt%2 > 0 {
		str = str + "_"
	}

	for i := 0; i < cnt; {
		fmt.Println(i)
		if i+2 <= len(str) {
			pair = str[i : i+2]
			fmt.Println(pair)
		}
		res = append(res, pair)
		i = i + 2
	}
	return res
}
