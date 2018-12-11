package main

import "fmt"
import "strings"

// import "math"

func main() {
	la_biec_30()
}

func print_star(x int) {
	fmt.Println(strings.Repeat("*", x))
}

// La biec 7
func la_biec_7() {
	var x int
	fmt.Scan(&x)
	ve_tam_giac(x)
}

func ve_tam_giac(x int) {
	nstar := 1
	for i := 1; i <= x; i++ {
		nblank := x - i
		blanks := strings.Repeat(" ", nblank)
		stars := strings.Repeat("*", nstar)
		output := blanks + stars + blanks
		fmt.Println(output)
		nstar += 2
	}
}

// La biec 11
func la_biec_11() {
	var a, b int
	fmt.Scan(&a, &b)
	ve_chu_nhat(a, b)
}

func ve_chu_nhat(a, b int) {
	body := "*" + strings.Repeat(" ", a-2) + "*"

	print_star(a)
	for i := 2; i < b; i++ {
		fmt.Println(body)
	}
	print_star(a)
}

// La biec 15
func la_biec_15() {
	var x int
	fmt.Scan(&x)
	ve_tam_giac_rong(x)
}

func ve_tam_giac_rong(x int) {
	nstar := ve_chu_a(x)
	print_star(nstar)
}

func ve_chu_a(x int) int {
	nblank := x - 1
	blanks := strings.Repeat(" ", nblank)
	first := blanks + "*" + blanks

	fmt.Println(first)
	nstar := 3
	for i := 2; i < x; i++ {
		nblank := x - i
		blanks := strings.Repeat(" ", nblank)
		mblanks := strings.Repeat(" ", nstar-2)
		output := blanks + "*" + mblanks + "*" + blanks
		fmt.Println(output)
		nstar += 2
	}
	return nstar
}

// La biec 30
func la_biec_30() {
	var x int
	fmt.Scan(&x)
	ve_chu_x(x)
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func ve_chu_x(x int) {
	upside := x / 2
	n_upside := 0 - upside

	length := upside*2 + 1
	for i := n_upside; i <= upside; i++ {
		c_length := 2*Abs(i) + 1
		nblank := (length - c_length) / 2
		blanks := strings.Repeat(" ", nblank)
		output := blanks + "*" + blanks
		if c_length-2 >= 1 {
			output = blanks + "*" + strings.Repeat(" ", c_length-2) + "*" + blanks
		}

		fmt.Println(output)
		nblank += 1
	}

	// downsize := x - upside
	// ve_chu_v(upside)
	// ve_chu_a(downsize)
	// length := downsize*2 - 1
	// end := "*" + strings.Repeat(" ", length-2) + "*"
	// fmt.Println(end)
}

// func ve_chu_v(x int) {
// 	length := x*2 + 1
// 	nblank := 0
// 	for i := x; i >= 1; i-- {
// 		blanks := strings.Repeat(" ", nblank)
// 		output := blanks + "*" + strings.Repeat(" ", length-2) + "*" + blanks
// 		fmt.Println(output)
// 		length -= 2
// 		nblank += 1
// 	}
//
// }
