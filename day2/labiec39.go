package day2

import (
	"fmt"
	"go-study/pkg/utils"
)

func LaBiec39() {
	var n, height int
	fmt.Scan(&n, &height)
	printV(n, height)
}

func printV(n, h int) {
	lenght := h*2 - 1
	//line 1
	for i := 0; i < n; i++ {
		utils.PrintString("*", 1)
		utils.PrintString(" ", lenght-2)
	}
	utils.PrintString("*", 1)
	fmt.Println()

	//line n
	sideBlanks := 1
	midBlanks := lenght - 4
	midBlank2s := 0

	for hh := 2; hh < h; hh++ {
		for i := 0; i < n; i++ {
			utils.PrintString(" ", sideBlanks)
			utils.PrintString("*", 1)
			utils.PrintString(" ", midBlanks)
			utils.PrintString("*", 1)
			utils.PrintString(" ", midBlank2s)
		}
		fmt.Println()
		sideBlanks++
		midBlanks -= 2
		midBlank2s += 1
	}

	//last line
	utils.PrintString(" ", h-1)
	utils.PrintString("*", 1)
	for i := 1; i < n; i++ {
		utils.PrintString(" ", lenght-2)
		utils.PrintString("*", 1)
	}
}
