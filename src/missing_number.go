package main

import "fmt"

// find missing number
// Example 1
// input [1, 2, 3, 5, 6]
// output 4

// Example 2
// input [12, 13, 14, 15, 17]
// output 16

func solution(request []int) int {
	size := len(request)
	interim := make([]int, size+2)
	for _, v := range request {
		interim[v] += 1
	}
	for i, v := range interim {
		if v == 0 && i > 0 {
			return i
		}
	}
	return -1
}

func missing_number() {
	input_request := []int{1, 2, 6, 3, 5}
	result := solution(input_request)
	fmt.Printf("input=%v output=%v\n", input_request, result)
}
