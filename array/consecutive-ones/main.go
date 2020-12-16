package main

import (
	"fmt"
)

// Question: Find count of maximum consecutive 1s in a binary array.
// Example: In: {0, 0, 1, 1, 1, 0, 1}, Out: 3

// Solution: O(N)
/*
	Traverse the array from left to right. Whenever we see
	a zero we reset a consecutive count. Whenever we see one
	we increment the consecutive count and compare it with
	previous maximum consecutive count.
*/
func consecutiveOnes(arr []int) int {
	consecutiveCount := 0
	maxConsecutiveCount := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			consecutiveCount = 0
		} else {
			consecutiveCount++
			maxConsecutiveCount = max(maxConsecutiveCount, consecutiveCount)
		}
	}

	return maxConsecutiveCount
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(consecutiveOnes([]int{0, 1, 1, 0, 0, 0, 1, 1, 1, 0}))
}
