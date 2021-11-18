package main

import (
	"fmt"
)

func main() {
	fmt.Println(ValidParentheses("())")) // false

	fmt.Println("------------")
	fmt.Println(ValidParentheses("(())")) // true

	fmt.Println("------------")
	fmt.Println(ValidParentheses("())(()"))
}

func ValidParentheses(parens string) bool {
	stack := make([]string, 0)

	if len(parens)%2 != 0 {
		return false
	}

	for _, i := range parens {
		if string(i) == "(" {
			stack = append(stack, "(")
		}
		if len(stack) > 0 && string(i) == ")" {
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}

// other smart solution was this
// https://www.codewars.com/kata/reviews/58c860e71f38971cad0002b4/groups/58c947a7b85ea0d8f1000f26
