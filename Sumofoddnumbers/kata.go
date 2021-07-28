package main

import "fmt"

func main() {
	RowSumOddNumber(3)
	RowSumOddNumber(13)
}

func RowSumOddNumber(n int) {
	finalnum := n*n + n - 1
	arr := GenerateOddNumberTriangle(finalnum)
	final := len(arr)
	initial := final - n
	answerarr := arr[initial:final]
	answer := sum(answerarr)
	fmt.Println(answer)
}

func GenerateOddNumberTriangle(n int) []int {
	a := make([]int, 0)
	for i := 0; i <= n; i++ {
		if i%2 != 0 {
			a = append(a, i)
		}
	}
	return a
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

// Note please don't try my solution it is highly overengineered the better
// one is this
// func RowSumOddNumbers(n int) int {
//     return n * n * n
// }
