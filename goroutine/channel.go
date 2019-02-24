package main

import (
	"fmt"
	"mypkg/slices"
	"mypkg/timetrack"
	"time"
)

var Result = []int{}

func main() {
	jobs := make_slice()

	defer timetrack.Track(time.Now(), "main")
	nWorker := 3
	cjob := make(chan int, 10000)

	for i := 1; i <= nWorker; i++ {
		go GetResult(cjob)
	}

	for _, j := range jobs {
		cjob <- j
	}
	close(cjob)
	fmt.Println(Result)
}

func GetResult(c chan int) {
	for {
		v, more := <-c
		if !more {
			return
		}
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

func make_slice() []int {
	var SliceA = []int{1, 2, 5, 76, 23, 4, 5, 6, 7, 34}
	var SliceB = []int{11, 35, 3, 5, 77, 4, 33, 76, 9}
	var SliceC = []int{15, 56, 3, 5, 6, 7, 33, 66, 88}

	SliceA = append_range(SliceA, 100, 1000000)
	SliceB = append_range(SliceB, 100, 1000000)
	SliceC = append_range(SliceC, 100, 1000000)

	s := SliceA
	s = append(s, SliceB...)
	s = append(s, SliceC...)

	return s
}
