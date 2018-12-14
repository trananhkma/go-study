package main

import (
    "fmt"
    "strings"
)

func main() {
    var x int
    fmt.Scan(&x)
    Draw_X(x, "*")
}

func Draw_X(x int, blank_char string) {
    upside := x / 2
    n_upside := 0 - upside

    length := upside*2 + 1
    for i := n_upside; i <= upside; i++ {
        c_length := 2*Abs(i) + 1
        nblank := (length - c_length) / 2
        blanks := strings.Repeat(blank_char, nblank)
        output := blanks + "*" + blanks
        if c_length-2 >= 1 {
            output = blanks + "*" + strings.Repeat(" ", c_length-2) + "*" + blanks
        }

        fmt.Println(output)
        nblank += 1
    }
}

func Abs(x int) int {
    if x > 0 {
        return x
    }
    return -x
}
