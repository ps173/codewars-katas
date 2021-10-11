package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ToCamelCase("The_Stealth_Warrior"))
	fmt.Println(ToCamelCase("the-Stealth-Warrior"))
}

func ToCamelCase(s string) string {
	ans := make([]string, 2)
	if strings.Contains(s, "-") {
		splstr := strings.Split(s, "-")
		for x := 0; x < len(splstr); x++ {
			ans = append(ans, strings.Title(splstr[x]))
		}
	} else {
		splstr := strings.Split(s, "_")
		for x := 0; x < len(splstr); x++ {
			ans = append(ans, strings.Title(splstr[x]))
		}
	}

	return strings.Join(ans, "")
}
