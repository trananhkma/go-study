package day2

import "fmt"

func Trunghoc10() {
	trunghoc10()
}

func trunghoc10() {
	var n int
	fmt.Scan(&n)
	fmt.Println(1)

	if n <= 1 {
		return
	}

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		for j := i - 1; j >= 1; j-- {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
