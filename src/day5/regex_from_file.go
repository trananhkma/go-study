package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

func main()  {
    f, err := os.Open("file_name.txt")
    if err != nil {
        fmt.Println(err)
    }

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanLines)

    phonePattern, err2 := regexp.Compile(`\b0\d{9,10}\b`)
    if err2 != nil {
        fmt.Println(err2)
    }

    for scanner.Scan() {
        allStr := phonePattern.FindAllString(scanner.Text(), -1)
        for _, v := range allStr {
            fmt.Println(v)
        }
    }


}