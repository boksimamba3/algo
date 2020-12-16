package main

import (
	"fmt"
)

// Question: Find a maximum subarray sum of contiguous subarray of numbers.
// Example: In: {-5, 1, -2, 3, -1, 2, -2}

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
}
*/

// Optimized solution O(N) Kadane's algorithm
/*
	Traverse the array from left to right. We keep track of
	max sum ending at previous element, to compute max sum
	ending at current element.
	For i-th element max we compute max sum:
	maxEnding(i) = max(maxEnding(i-1) + arr[i], arr[i])

	In = {-3, 8, -2, 4, -5, 6}, Out: 11
	maxEnding = -3, maxSum = -3

	I1: i = 1, maxEnding = max(-3 + 8, 8) = 8, maxSum = max(-3, 8) = 8
	I2: i = 2, maxEnding = max(8 - 2, -2) = 6, maxSum = max(8, 6) = 8
	I3: i = 3, maxEnding = max(6 + 4, 4) = 10, maxSum = max(10, 8) = 10
	I4: i = 4, maxEnding = max(10 - 5, -5) = 5, maxSum = max(5, 10) = 10
	I5: i = 5, maxEnding = max(5 + 6) = 11, maxSum = max(11, 10) = 11
*/
func maxSumSubarray(arr []int) int {
	maxSum := arr[0]
	maxEnding := arr[0]

	for i := 1; i < len(arr); i++ {
		maxEnding = max(maxEnding+arr[i], arr[i])
		maxSum = max(maxSum, maxEnding)
	}

	return maxSum
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maxSumSubarray([]int{-3, 8, -2, 4, -5, 6}))
}
