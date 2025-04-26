package main

import (
	"fmt"
	"time"
)

func speedtest() {
	n := 1000000
	testSlice := []int{}            // length 0, capacity 0
	testSlice2 := make([]int, 0, n) // length 0, capacity 10

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	t0 := time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
