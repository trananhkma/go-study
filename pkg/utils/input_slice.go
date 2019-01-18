package utils

import "fmt"

func InputSlice(n int) []int {
    var slice []int
    for i:=0;i<n;i++ {
        var x int
        fmt.Scan(&x)
        slice = append(slice, x)
    }
    return slice
}
