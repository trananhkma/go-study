package utils

import (
    "fmt"
    "strings"
)

func PrintString(s string, n int)  {
    fmt.Printf(strings.Repeat(s, n))
}

func PrintV(n, h int)  {
    lenght := h*2-1
    //line 1
    for i:=0; i<n; i++ {
        PrintString("*", 1)
        PrintString(" ", lenght-2)
    }
    PrintString("*", 1)
    fmt.Println()

    //line n
    sideBlanks := 1
    midBlanks := lenght-4
    midBlank2s := 0

    for hh:=2; hh<h; hh++ {
        for i:=0; i<n; i++ {
            PrintString(" ", sideBlanks)
            PrintString("*", 1)
            PrintString(" ", midBlanks)
            PrintString("*", 1)
            PrintString(" ", midBlank2s)
        }
        fmt.Println()
        sideBlanks++
        midBlanks -= 2
        midBlank2s += 1
    }

    //last line
    PrintString(" ", h-1)
    PrintString("*", 1)
    for i:=1; i<n; i++ {
        PrintString(" ", lenght-2)
        PrintString("*", 1)
    }
}