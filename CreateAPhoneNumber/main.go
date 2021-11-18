package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
}

func CreatePhoneNumber(numbers [10]uint) string {
	var arr []string
	arr = append(arr, "(")
	for index, i := range numbers {
		convertedChar := strconv.FormatUint(uint64(i), 10)
		if index == 3 {
			arr = append(arr, ") ")
		}
		if index == 6 {
			arr = append(arr, "-")
		}
		arr = append(arr, convertedChar)
	}
	reqString := strings.Join(arr, "")
	return fmt.Sprintln(reqString)
}

// Some really good answers
// https://www.codewars.com/kata/reviews/5cb39991f180cc00014b582f/groups/5cb5215f8a5a8800013dc599
// https://www.codewars.com/kata/reviews/5cb39991f180cc00014b582f/groups/5cdf97cad0d63b0001e95091
