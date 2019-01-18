package main

import (
	"fmt"
)

func main() {
	var ntest int
	fmt.Scan(&ntest)
	for ntest > 0 {
		ntest -= 1
		var h, x, y int
		fmt.Scan(&h, &x, &y)
		snail(h, x, y)
	}

}

func snail(h, x, y int) {
	if y >= x {
		fmt.Println("LOSER")
		return
	}
	day := 0
	for h > 0 {

		day += 1
		if h-x <= 0 {
			break
		}
		h = h - x + y
		//if h-x <= 0 {
		//	break
		//}
	}
	fmt.Println(day)
}
