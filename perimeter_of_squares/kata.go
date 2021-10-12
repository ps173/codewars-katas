package main

import "fmt"

func main() {
	fmt.Println(Perimeter(5))
}

func Perimeter(n int) int {
	sum := 0
	// why this cause otherwise the result goes out of index
	fib := make([]int, n+2)

	// stop adding when sum n is zero
	if n <= 0 {
		sum += 0
	}

	// initialize the first 2 elements so as when we add them it does not goes
	// out of index
	fib[0] = 0
	fib[1] = 1

	// Initialize sum
	sum += fib[0] + fib[1]

	// Add remaining terms
	for i := 2; i <= n+1; i++ {
		fib[i] = fib[i-1] + fib[i-2]
		sum += fib[i]
	}

	// and give out the perimeter
	return sum * 4
}
