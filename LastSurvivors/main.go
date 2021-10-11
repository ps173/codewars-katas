package main

import "strconv"

func main() {
	equalLetters("zzzab")
}

func equalLetters(s string) {
	// alphabets := "abcdefghijklmnopqrstuvwxyz"
	for i, val := range s {
		if i+1 == 5 {
			break
		}
		println(strconv.QuoteRuneToASCII(val), (val))
	}
}
