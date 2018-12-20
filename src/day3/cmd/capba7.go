package main

import (
    "fmt"
)

func main() {
    var ntest int

    fmt.Scan(&ntest)

    for i := 0; i < ntest; i++ {
        result, max := capba7()
        fmt.Printf("Test %d: %s %d\n", i+1, result, max);
    }
}

func capba7() (string, int) {
    var n int
    var result string
    var max int
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        var s string
        fmt.Scan(&s)

        chars := kytukhacnhau(s)
        if chars > max {
            max = chars
            result = s
        }
    }
    return result, max
}

func kytukhacnhau(s string) int {
    var slice [256]int
    var count int
    for _, v := range s {
        slice[v] += 1
    }

    for _, v := range slice {
        if v > 0 {
            count += 1
        }
    }

    return count
}
