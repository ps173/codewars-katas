package main

import (
	"fmt"
)

func main() {
	PeakPos([]int{1, 2, 2, 3, 2, 4, 5, 5, 2, 2, 1})
}

func MaxPos(array []int) int {
	var maxpos int
	for index, value := range array {
		if array[0] < value {
			array[0] = value
			maxpos = index
		}
	}
	return maxpos
}

func PeakPos(arr []int) {
	max := MaxPos(arr)
	fmt.Println(max)
}
