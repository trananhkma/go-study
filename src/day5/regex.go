package main

import (
    "fmt"
    "regexp"
)

func main() {
    var ntest int
    fmt.Print("Ntest: ")
    fmt.Scan(&ntest)

    for i := 0; i < ntest; i++ {
        var text string
        fmt.Scan(&text)
        fmt.Printf("Test %d: %s\n", i+1, regex(text));
    }
}

func regex(s string) string {
    // phonePattern, err := regexp.Compile(`\b0\d{9,10}\b`)
    emailPattern, err := regexp.Compile(`^[a-z]([a-z0-9_\-.]+)@(([a-z0-9_\-]+)\.){1,2}([a-z]){2,3}$`)
    pattern := emailPattern

    if err != nil {
        fmt.Println(err)
        return "Wrong Pattern"
    }

    if pattern.Match([]byte(s)) {
        return "Matched"
    }

    return "Not Matched"
}