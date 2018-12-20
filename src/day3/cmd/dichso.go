package main

import (
    "fmt"
)


var phatAm, donVi = initMap()

func main() {
    var ntest int

    fmt.Scan(&ntest)

    for i := 0; i < ntest; i++ {
        var x string
        fmt.Scan(&x)
        x = cut0(x)
        fmt.Printf("Test %d: %s\n", i+1, dichso(x));
    }
}

func cut0(x string) string {
    result := x
    for i:=0; i<len(x)-1;i++ {
        if string(x[i]) == "0" {
            result = x[i+1:]
        } else {
            break
        }
    }

    return result
}

func initMap() (map[string]map[string]string, map[int]string) {
    var phatAm = map[string]map[string]string{
        "normal": map[string]string{},
        "right": map[string]string{},
    }


    normal := phatAm["normal"]

    normal["0"] = "không"
    normal["1"] = "một"
    normal["2"] = "hai"
    normal["3"] = "ba"
    normal["4"] = "bốn"
    normal["5"] = "năm"
    normal["6"] = "sáu"
    normal["7"] = "bảy"
    normal["8"] = "tám"
    normal["9"] = "chín"
    normal["10"] = "mười"

    right := phatAm["right"]
    right["1"] = "mốt"
    right["4"] = "tư"
    right["5"] = "lăm"
    right["0"] = ""

    var donVi = map[int]string{
        0: "",
        1: "nghìn",
        2: "triệu",
        3: "tỉ",
    }

    return phatAm, donVi
}


func dichso (s string) string {
    var first = len(s)%3
    var level = len(s)/3

    dichFirst := false
    if first == 0 {
        level -= 1
        dichFirst = true
    }

    var result string

    for level >= 0 {
        var block string
        var subresult string
        if !dichFirst {
            block = s[:first]
            dichFirst = true
        } else {
            var star, end = first, first+3
            block = s[star:end]
            first += 3
        }
        if _, ok := phatAm["normal"][block]; ok {
            subresult = phatAm["normal"][block]
        }

        if len(block) == 2 {
            subresult = dichBlock2(block)
        }

        if len(block) == 3 {
            subresult = dichBlock3(block)
        }
        result += subresult + " " + donVi[level] + " "
        level--
    }
    return result
}

func dichBlock2(s string) string {
    var first = string(s[0])
    var last = string(s[1])

    if first == "1" {
        first = "mười "
        if last == "5" {
            last = "lăm"
        } else {
            last = phatAm["normal"][last]
        }
    } else {
        first = phatAm["normal"][first] + " mươi "
        if _, ok := phatAm["right"][last]; ok {
            last = phatAm["right"][last]
        } else {
            last = phatAm["normal"][last]
        }
    }
    return first + last
}

func dichBlock3(s string) string {
    if s == "000" {
        return ""
    }

    var first = string(s[0])
    first = phatAm["normal"][first] + " trăm"
    if string(s[1:]) == "00" {
        return first
    }

    if string(s[1]) == "0" {
        var mid = " linh "
        var last = phatAm["normal"][string(s[2])]
        return first + mid + last
    } else {
        var last = dichBlock2(s[1:])
        return first + " " + last
    }
}