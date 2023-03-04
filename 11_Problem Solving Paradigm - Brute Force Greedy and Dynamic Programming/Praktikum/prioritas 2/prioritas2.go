package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getAbsolut(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func frog(jumps []int) int {
	n := len(jumps)
	result := make([]int, n)

	result[0] = 0
	result[1] = getAbsolut(jumps[1] - jumps[0])

	for i := 2; i < n; i++ {
		result[i] = min(result[i-1]+getAbsolut(jumps[i]-jumps[i-1]), result[i-2]+getAbsolut(jumps[i]-jumps[i-2]))
	}

	return result[n-1]
}

func main() {
	fmt.Println(frog([]int{10, 30, 40, 20}))
	fmt.Println(frog([]int{30, 10, 60, 10, 60, 50}))
}
