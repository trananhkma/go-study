package day2

import (
	"fmt"
	"go-study/pkg/utils"
	"strconv"
)

func Trunghoc9() {
	var ntest int

	fmt.Scan(&ntest)

	for i := 0; i < ntest; i++ {
		s := ""
		result := trunghoc9()
		for _, v := range result {
			s += strconv.Itoa(v) + " "
		}
		fmt.Printf("Test %d: %s\n", i+1, s)
	}
}

func trunghoc9() []int {
	var n, m int
	fmt.Scan(&n, &m)

	closed := []int{}
	for i := 1; i <= n; i++ {
		s := utils.InputSlice(m)
		if is_closed(s) {
			closed = append(closed, i)
		}
	}
	if len(closed) == 0 {
		return []int{0}
	}
	return closed
}

func is_closed(s []int) bool {
	go_down := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			go_down += 1
		} else {
			go_down = 0
		}
		if go_down == 3 {
			return true
		}
	}
	return false
}
