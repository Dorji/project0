package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//	()
	//	(()) ()()
	//	()()() ()(()) (())() ((())) (()())
	generateParentheses("", 0, 0, 3)

}

func generateParentheses(curr string, open int, closed int, n int) string {

	if utf8.RuneCountInString(curr) == 2*n {
		fmt.Println(curr)
		return curr
	}
	if open < n {

		generateParentheses(curr+"(", open+1, closed, n)
	}
	if closed < open {

		generateParentheses(curr+")", open, closed+1, n)
	}
	return curr
}
