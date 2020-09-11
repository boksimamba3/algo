package main

import "fmt"

func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

// Naive solution 0(n^2)
func maxSumCircuralSubarray(arr []int) int {
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
}

func main() {
	fmt.Println(maxSumCircuralSubarray([]int{5, -2, 3, 4}))
}
