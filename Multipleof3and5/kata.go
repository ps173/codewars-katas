package main

import "fmt"

func main() {
	fmt.Println(Multipleof3or5(10))
}

func Multipleof3or5(number int) int {
	sum := 0
	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
