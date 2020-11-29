package main

import (
	"fmt"
)

// Maximun value of arr[j] - arr[i] such that j > i
/* func maxDiff(arr []int) int {
	maxDiff := 0

	for j := 1; j < len(arr); j++ {
		for i := 0; i < j; i++ {
			if arr[j]-arr[i] > maxDiff {
				maxDiff = arr[j] - arr[i]
			}
		}
	}

	return maxDiff
} */

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxDiff(arr []int) int {
	maxDiff := arr[1] - arr[0]
	minVal := arr[0]
	for j := 1; j < len(arr); j++ {
		maxDiff = max(maxDiff, arr[j]-minVal)
		minVal = min(minVal, arr[j])
	}
	return maxDiff
}

func main() {
	arr := []int{2, 3, 10, 6, 4, 17}
	fmt.Println(maxDiff(arr))
}
