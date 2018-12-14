package main

import (
    "day2/pkg/utils"
    "fmt"
)

func main()  {
    var n, height int
    fmt.Scan(&n, &height)
    vHeight := height/2 + 1
    printA(n, vHeight)
    utils.PrintV(n, vHeight)
}


func printA(n, h int) {
    lenght := h*2
    utils.PrintString(" ", h-1)
    utils.PrintString("*", 1)

    for i:=1; i<n; i++ {
       utils.PrintString(" ", lenght-3)
       utils.PrintString("*", 1)
    }
    fmt.Println()

    //line n
    sideBlanks := h-2
    blanks := 1
    midBlanks := lenght-5

    for midBlanks > 0 {
        utils.PrintString(" ", sideBlanks)
        for i:=0; i<n; i++ {
            utils.PrintString("*", 1)
            utils.PrintString(" ", blanks)
            utils.PrintString("*", 1)
            utils.PrintString(" ", midBlanks)

        }
        fmt.Println()
        sideBlanks--
        blanks += 2
        midBlanks -= 2
    }

}