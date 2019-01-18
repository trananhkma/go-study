package main

import (
	"fmt"
	"strconv"
)

func main() {
	var ntest int
	fmt.Scan(&ntest)
	for ntest > 0 {
		ntest -= 1
		var s string
		fmt.Scan(&s)
		fmt.Println(mini(s))
	}

}

func mini(s string) string {
	ss := 0
	for _, v := range s {
		i, _ := strconv.Atoi(string(v))
		ss += i
	}

	if ss < 10 {
		return strconv.Itoa(ss)
	}

	var result string

	for i := 1; i < 10; i++ {
		x := ss - i
		if x < 10 {
			result = strconv.Itoa(i) + strconv.Itoa(x) + result
			break
		}
	}
	result += "9"

	return result
}
