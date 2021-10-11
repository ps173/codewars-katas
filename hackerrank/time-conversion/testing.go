package main

import (
	"time"
)

func main() {
	timeConversion("12:01:00AM")
}

func timeConversion(s string) string {
	// defining layouts for time this can have any value
	// you just have to keep the militaryLayout value correct :)
	normalLayout := "03:04:05PM"
	militaryLayout := "15:04:05"
	t, _ := time.Parse(normalLayout, s)
	return t.Format(militaryLayout)
}
