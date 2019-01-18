package day3

import (
	"fmt"
	"strconv"
)

func Bignum() {
	var ntest int

	fmt.Scan(&ntest)

	for i := 0; i < ntest; i++ {
		var s1, s2 string
		fmt.Scan(&s1, &s2)
		fmt.Printf("Test %d: %s\n", i+1, big_sum(s1, s2))
	}
}

func big_sum(s1, s2 string) string {
	len1 := len(s1)
	len2 := len(s2)

	if len1 > len2 {
		s2 = insert0tostring(s2, len1-len2)
	} else if len1 < len2 {
		s1 = insert0tostring(s1, len2-len1)
	}

	index := len1 - 1

	result := ""
	remain := 0

	for index >= 0 {
		int1, _ := strconv.Atoi(string(s1[index]))
		int2, _ := strconv.Atoi(string(s2[index]))
		sum := int1 + int2 + remain

		var strSum string

		if sum >= 10 {
			remain = 1
			strSum = strconv.Itoa(sum - 10)
		} else {
			remain = 0
			strSum = strconv.Itoa(sum)
		}

		result = strSum + result
		index--
	}

	if remain > 0 {
		result = strconv.Itoa(remain) + result
	}

	return result
}

func insert0tostring(s string, n int) string {
	result := s
	for i := 0; i < n; i++ {
		result = "0" + result
	}
	return result
}
