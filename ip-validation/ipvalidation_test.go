package ipvalidation_test

import (
	"strconv"
	s "strings"
	"testing"
)

func Is_valid_ip(ip string) bool {
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

func TestIs_valid_ip(t *testing.T) {
	test1 := Is_valid_ip("1.2.3.4")
	if test1 == false {
		t.Error("Failed")
	}

	test2 := Is_valid_ip("0.2.3.4")
	if test2 == true {
		t.Error("Failed")
	}
}
