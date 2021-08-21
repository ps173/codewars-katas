package main

import (
	"fmt"
	"sort"
)

var p = fmt.Println

func main() {
	HighestRank([]int{1, 2, 3, 4, 1, 2, 2, 10, 10})
	// p(HighestRank([]int{1, 2, 3, 4, 1, 2, 1, 2}))
}

func HighestRank(nums []int) int {
	counts := 0
	m := make(map[int]int)

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i-1] == nums[i] {
			counts++
			m[nums[i]] = counts + 1
		}
		if i != 0 && nums[i-1] != nums[i] {
			counts = 0
		}
	}

	p(FindMaxVal(m))

	return -1
}

func FindMaxVal(numbers map[int]int) int {
	var maxNumber int
	for maxNumber = range numbers {
		break
	}
	for k := range numbers {
		if k > maxNumber {
			maxNumber = k
		}
	}
	return maxNumber
}
