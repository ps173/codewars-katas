package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(VowelCount("abracadabra"))

}

func VowelCount(str string) int {
	r, _ := regexp.Compile("[aeiou]")
	return len(r.FindAllStringSubmatchIndex(str, -1))
}
