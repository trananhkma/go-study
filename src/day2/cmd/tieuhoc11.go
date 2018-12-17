package main

import (
    "day2/pkg/utils"
    "fmt"
)

func main() {
    var ntest int

    fmt.Scan(&ntest)

    for i := 0; i<ntest; i++ {
        fmt.Printf("Test %d: %d\n", i+1, tieuhoc11())
    }
}


func tieuhoc11() int {
    var n int
    fmt.Scan(&n)

    s := utils.InputSlice(n)
    count := 0
    for i:=0;i<len(s)-1;i++ {
        for j:=i+1; j<len(s); j++ {
            if is_prime(s[i] + s[j]) {
                count += 1
            }
        }
    }
    return count
}


func is_prime (x int) bool {
    for i:=2; i<=x/2; i++ {
        if x % i == 0 {
            return false
        }
    }
    return true
}
