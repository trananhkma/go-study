package main

import (
    "fmt"
)

func main() {
    var ntest int

    fmt.Scan(&ntest)

    for i := 0; i < ntest; i++ {
        a, b, greatest := trunghoc8()
        fmt.Printf("Test %d: %d %d %d\n", i+1, a, b, greatest);
    }
}

func trunghoc8() (int, int, int) {
    var a, b int
    fmt.Scan(&a, &b)

    greatest := 0
    for i:=a; i<=b; i++ {
        l := lengh_cycle(i)
        if l > greatest {
            greatest = l
        }
    }

    return a, b, greatest
}

func lengh_cycle(x int) int {
    count := 0
    for true {
        count ++
        if x == 1{
            break
        }

        if x % 2 == 0 {
            x /= 2
        } else {
            x = 3*x + 1
        }
    }
    return count
}