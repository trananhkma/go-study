package main

import (
    "fmt"
)

func main() {
    var ntest int

    fmt.Scan(&ntest)

    for i:=0; i<ntest; i++ {
        var size, k int;
        fmt.Scan(&size, &k);

        result := "xxx";
        fmt.Printf("Test %d: %s", i+1, result);
    }
}
