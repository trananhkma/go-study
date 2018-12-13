package main

import (
	"./pakage1"
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	pakage1.Draw_X(x, "*")
}
