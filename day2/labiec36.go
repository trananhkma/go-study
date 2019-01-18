package day2

import (
	"fmt"
	"go-study/pkg/utils"
)

func LaBiec36() {
	var n, thick int
	fmt.Scan(&n, &thick)
	laBiec36(n, thick)
}

func laBiec36(n, thick int) {
	for i := 0; i < n; i++ {
		print7(thick)
	}
	utils.PrintString("*", thick)
}

func print7(thick int) {
	utils.PrintString("*", thick)
	fmt.Println()

	//line n
	for i := 2; i < thick; i++ {
		utils.PrintString(" ", thick-i)
		utils.PrintString("*", 1)
		fmt.Println()
	}

}
