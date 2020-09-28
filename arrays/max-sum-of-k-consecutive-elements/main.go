package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

// O(N)
func maxSum(arr []int, k int) int {
	if k > len(arr) {
		return 0
	}

	sum := 0

	for i := 0; i < k; i++ {
		sum += arr[i]
	}

	currSum := sum

	for start, end := 0, k; end < len(arr); start, end = start+1, end+1 {
		currSum += arr[end] - arr[start]
		sum = max(sum, currSum)
	}

	return sum
}

func main() {
	fmt.Println(maxSum([]int{1, 8, 30, -5, 20}, 3))
}
