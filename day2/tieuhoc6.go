package day2

import (
	"fmt"
	"go-study/pkg/utils"
)

func Tieuhoc6() {
	var ntest int

	fmt.Scan(&ntest)

	for i := 0; i < ntest; i++ {
		var n int
		fmt.Scan(&n)
		slice := utils.InputSlice(n)
		item, count := findMostAppear(slice)
		fmt.Printf("Test %d: %d %d \n", i+1, item, count)
	}
}

func findMostAppear(s []int) (int, int) {
	item := 0
	most := 0
	set := flatSlice(s)

	for _, v := range set {
		count := countAppear(v, s)
		if (count > most) || (count == most && v > item) {
			most = count
			item = v
		}
	}

	return item, most
}

func countAppear(x int, s []int) int {
	count := 0
	for _, v := range s {
		if v == x {
			count += 1
		}
	}
	return count
}

func contain(s []int, x int) bool {
	for _, value := range s {
		if x == value {
			return true
		}
	}
	return false
}

func flatSlice(s []int) []int {
	var slice []int
	for _, value := range s {
		if !contain(slice, value) {
			slice = append(slice, value)
		}
	}
	return slice
}
