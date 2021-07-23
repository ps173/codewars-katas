package ipvalidation

import (
	"fmt"
	"strconv"
	s "strings"
)

var p = fmt.Println

func IpValid(ip string) {
	p(is_valid_ip(ip))
}

func is_valid_ip(ip string) bool {
	var numArr []int
	strArr := s.Split(ip, ".")
	for i := 0; i < len(strArr); i++ {
		if string(strArr[i][0]) == "0" {
			break
		}
		numb, err := strconv.Atoi(strArr[i])
		if err != nil {
			break
		}
		if numb < 255 && numb > 0 {
			numArr = append(numArr, numb)
		}
	}
	return len(numArr) > 0
}
