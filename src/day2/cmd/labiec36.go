package main

import (
    "day2/pkg/utils"
    "fmt"
)
func main()  {
    var n, thick int
    fmt.Scan(&n, &thick)
    LaBiec36(n, thick)
}

func LaBiec36(n , thick int)  {
    for i:=0;i<n;i++ {
        print7(thick)
    }
    utils.PrintString("*", thick)
}

func print7(thick int){
    utils.PrintString("*", thick)
    fmt.Println()

    //line n
    for i:=2;i<thick;i++ {
        utils.PrintString(" ", thick-i)
        utils.PrintString("*", 1)
        fmt.Println()
    }

}