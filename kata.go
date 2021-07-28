package main

import "fmt"

func main() {
	fmt.Println(BreakChocolate(5, 5))
}

func BreakChocolate(n, m int) int {
	area := m * n
	if area > 1 {
		return area - 1
	} else {
		return 0
	}
}
