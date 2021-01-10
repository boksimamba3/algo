package main

import "fmt"

// Question: Given an array of integers and a number k, find the
// maximum sum of k consecutive elements.

// Solution: O(N)
/*
	Using a sliding window technique we can in linear time find max sum
	of k consecutive elements. First we calculate the sum for first window (0...k).
	We initialize the max sum to sum of first k elements. To calculate the sum of
	remaning windows we traverse the remaining array starting from k. To calculate the
	sum for current window we need to add current element and subtract first element of
	previous window. We check if current sum is bigger then previous max sum.
*/
func maxSum(arr []int, k int) int {
	if k > len(arr) {
		return 0
	}

	sum := 0

	for i := 0; i < k; i++ {
		sum += arr[i]
	}

	currSum := sum

	for i := k; i < len(arr); i++ {
		currSum += arr[i] - arr[i-k]
		sum = max(sum, currSum)
	}

	return sum
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maxSum([]int{1, 8, 30, -5, 20}, 1))
	fmt.Println(maxSum([]int{1, 8, 30, -5, 20}, 3))
}
