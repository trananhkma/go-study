package main

import (
    "fmt"
)

func main() {
    var ntest int

    fmt.Scan(&ntest)

    for i := 0; i < ntest; i++ {
        fmt.Printf("Test %d: %d\n", i+1, _test());
    }
}

func _test() int {
    return 0
}