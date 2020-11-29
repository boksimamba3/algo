package main

import (
	"fmt"
)

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

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

func main() {
	fmt.Println(consecutiveOnes([]int{0, 1, 1, 0, 0, 0, 1, 1, 1, 0}))
	fmt.Println(consecutiveOnes([]int{1, 1, 1, 1}))
	fmt.Println(consecutiveOnes([]int{0, 0, 0, 0}))
	fmt.Println(consecutiveOnes([]int{0, 0, 0, 1}))
	fmt.Println(consecutiveOnes([]int{1, 0, 0}))
}
