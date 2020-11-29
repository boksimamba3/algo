package main

import "fmt"

func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

// Naive solution 0(n^2)
/* func maxSumCircuralSubarray(arr []int) int {
	if len(arr) < 1 {
		return 0
	}

	maxSum := arr[0]

	for i := 0; i < len(arr); i++ {
		currentMax := arr[i]
		currentSum := arr[i]
		for j := 1; j < len(arr); j++ {
			index := (i + j) % len(arr)
			currentSum += arr[index]
			currentMax = max(currentMax, currentSum)
		}
		maxSum = max(maxSum, currentMax)
	}

	return maxSum
} */

func maxSumSubarray(arr []int) int {
	maxSum := arr[0]
	maxConsecutiveSum := arr[0]
	for i := 1; i < len(arr); i++ {
		maxConsecutiveSum = max(arr[i], maxConsecutiveSum+arr[i])
		maxSum = max(maxSum, maxConsecutiveSum)
	}

	return maxSum
}

func maxSumCircuralSubarray(arr []int) int {
	maxConsecutiveSum := maxSumSubarray(arr)
	if maxConsecutiveSum < 0 {
		return 0
	}

	arrSum := 0

	for i := 0; i < len(arr); i++ {
		arrSum += arr[i]
		arr[i] = -arr[i]
	}
	// Finding max sum in inverted array is like finding min in original
	maxCircuralSum := arrSum + maxSumSubarray(arr)

	return max(maxCircuralSum, maxConsecutiveSum)
}

func main() {
	fmt.Println(maxSumCircuralSubarray([]int{5, -2, 3, 4}))
	fmt.Println(maxSumCircuralSubarray([]int{8, -4, 3, -5, 4}))
}
