package main

import (
	"fmt"
)

// Question: Find maximun difference of arr, arr[j] - arr[i] where j > i.

// Solution: O(N)
/*
  Traverse the whole array, and while doing so
  keep track of minimum element on the left:

	In: {3, 4, 2, 8, 10, 1}, Out: 8
	maxDiff => arr[1] - arr[0] = 1
	minVal => min(3, 4) => 3
	I1: j = 2, maxDiff = max(1, -2) = 1, minVal = min(3, 2) = 2
	I2: j = 3, maxDiff = max(1, 6) = 6, minVal = min(2, 8) = 2
	I3: j = 4, maxDiff = max(6, 8) = 8, minVal = min(2, 10) = 2
	I4: j = 5, maxDiff = max(8, -9) = 8, minVal = min(2, 1) = 1
	Out: 8
*/
func maxDiff(arr []int) int {
	maxDiff := arr[1] - arr[0]
	minVal := min(arr[0], arr[1])

	for j := 2; j < len(arr); j++ {
		maxDiff = max(maxDiff, arr[j]-minVal)
		minVal = min(minVal, arr[j])
	}

	return maxDiff
}

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

func main() {
	fmt.Println(maxDiff([]int{2, 3, 10, 6, 4, 17}))
}
