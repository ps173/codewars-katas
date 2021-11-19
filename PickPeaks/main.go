package main

import (
	"fmt"
)

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func main() {
	PickPeaks([]int{1, 2, 2, 3, 2, 4, 5, 5, 2, 2, 1})
	PickPeaks([]int{5, 2, 2, 3, 2, 4, 5, 5, 2, 2, 5})
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

func PickPeaks(arr []int) PosPeaks {
	pos := MaxPos(arr)
	max := arr[pos]
	ret := PosPeaks{Pos: []int{pos}, Peaks: []int{max}}
	if pos == 0 || pos == len(arr)-1 {
		ret = PosPeaks{}
	}
	fmt.Println(ret)
	return ret
}
