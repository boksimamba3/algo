package main

import "fmt"

func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

// Naive solution O(N^2)
/* func maxSumSubarray(arr []int) int {
	maxSum := 0

	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			maxSum = max(maxSum, sum)
		}
	}

	return maxSum
} */

// Optimized solution O(N)
func maxSumSubarray(arr []int) int {
	maxSum := 0
	maxConsecutiveSum := 0

	for i := 0; i < len(arr); i++ {
		maxConsecutiveSum += arr[i]

		if maxConsecutiveSum < 0 {
			maxConsecutiveSum = 0
		} else if maxSum < maxConsecutiveSum {
			maxSum = maxConsecutiveSum
		}

	}

	return maxSum
}

func main() {
	fmt.Println(maxSumSubarray([]int{2, 3, -8, 7, -1, 2, 3, -1, 4}))
	fmt.Println(maxSumSubarray([]int{1, -2, 3, -1, 2}))
	fmt.Println(maxSumSubarray([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
}
