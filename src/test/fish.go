package main

import (
	"fmt"
	"strings"
)

func main() {
	var ntest int
	fmt.Scan(&ntest)
	for i := 0; i < ntest; i++ {
		var n int
		fmt.Scan(&n)
		fish(n)
	}
}

func fish(n int) {
	//first
	hieght := n / 2
	nstar := 1
	nblank := n - 2
	sideStar := 0
	for i := 0; i < hieght; i++ {
		PrintString("*", nstar)
		PrintString(" ", nblank)
		PrintString("*", nstar)
		PrintString("*", sideStar)
		fmt.Println()
		nstar += 1
		nblank -= 2
		sideStar += 1
	}

	PrintString("*", n+n/2)
	fmt.Println()

	for i := hieght; i > 0; i-- {
		nstar -= 1
		nblank += 2
		sideStar -= 1

		PrintString("*", nstar)
		PrintString(" ", nblank)
		PrintString("*", nstar)
		PrintString("*", sideStar)
		fmt.Println()
	}

	// PrintString("*", 1)
	// PrintString(" ", n-2)
	// PrintString("*", 1)

	// PrintString("*", n)
}

func PrintString(s string, n int) {
	fmt.Printf(strings.Repeat(s, n))
}
