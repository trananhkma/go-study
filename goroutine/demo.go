package main

import (
	"fmt"
	"mypkg/slices"
	"mypkg/timetrack"
	"sync"
	"time"
)

var SliceA = []int{1, 2, 5, 76, 23, 4, 5, 6, 7, 34}
var SliceB = []int{11, 35, 3, 5, 77, 4, 33, 76, 9}
var SliceC = []int{15, 56, 3, 5, 6, 7, 33, 66, 88}

var Result = []int{}
var wg sync.WaitGroup

func main() {
	defer timetrack.Track(time.Now(), "main")
	SliceA = append_range(SliceA, 100, 1000000)
	SliceB = append_range(SliceB, 100, 1000000)
	SliceC = append_range(SliceC, 100, 1000000)

	wg.Add(3)
	go GetResult(SliceA)
	go GetResult(SliceB)
	go GetResult(SliceC)

	wg.Wait()
	fmt.Println(Result)
}

func GetResult(s []int) {
	defer wg.Done()
	for _, v := range s {
		if v > 10 {
			Result = append(Result, v)
		}
	}
}

func append_range(s []int, a int, b int) []int {
	r := slices.MakeRange(a, b)
	x := append(s, r...)
	return x
}
