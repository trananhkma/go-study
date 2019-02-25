package main

import (
	"fmt"
	"sync"
)

type MyInfo struct {
	name  string
	age   int
	phone string
}

var info *MyInfo
var once sync.Once

func getInfo() *MyInfo {
	once.Do(func() {
		info = &MyInfo{name: "Anh Tran", age: 27, phone: "0986xxx686"}
	})
	return info
}

func main() {
	info1 := getInfo()
	info2 := getInfo()
	
	fmt.Println(info1)
	fmt.Println(info2)

	info1.age += 1
	fmt.Println(info1)
	fmt.Println(info2)

	info2.age += 1
	fmt.Println(info1)
	fmt.Println(info2)
}
