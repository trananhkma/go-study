package main

import "fmt"

func main(){
    var n int
    fmt.Scan(&n)

    s := tieuhoc10(n)
    fmt.Println(s)
}

func tieuhoc10(n int) []int {
    result := []int{6}
    if n < 7 {
        return []int{}
    }
    if n < 28 {
        return result
    }
    for i:=28; i<n; i++ {
        if isPerfectNumber(i) {
            result = append(result, i)
        }
    }
    return result

}

func isPerfectNumber(x int) bool {
    sum := 1
    for i:=2; i<=x/2; i++ {
        if x % i == 0 {
            sum += i
        }
    }

    if sum == x {
        return true
    }
    return false

}

