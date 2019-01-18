package main

import (
	"fmt"
)

func main() {
	var ntest int
	fmt.Scan(&ntest)
	for ntest > 0 {
		ntest -= 1
		var n int
		fmt.Scan(&n)
		slice := input_slice(n)
		fmt.Println(zero(slice))
	}

}

func zero(s []int) int {
	ss := 1
	count := 0
	for _, v := range s {
		ss *= v
	}
	for true {
		if ss%10 == 0 {
			ss /= 10
			count += 1
		} else {
			break
		}
	}

	return count
}

func input_slice(n int) []int {
	var slice []int
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		slice = append(slice, x)
	}
	return slice
}
