package main

import "fmt"

// Question: Find maximum circular subarray sum of a given array.

// Solution: O(N)
/*
	First calculate maximum sum of subarray in normal (not circural) array.
	After that check if maximum sum of subarray in normal array is less then
	zero. If it is that means that all elements in array are negative.
	After that we need to find minimum subarray sum. In order to not modify
	standard Kadane algorithm we invert the whole array (change sign of each element).
	This why we can find minimum subarray sum.
	While traversing the array and inverting the elements we also calculate the sum
	of inverted array. At the end we return maximum of maximum circural sum and maximum of normal
	consecutive sum.

	In: {[8], -4, 3, -5, 4}, maxConsecutiveSum = 8
	In inverted: {-8, [4, -3, 5], -4}, maxSubarraySum = 6, maxCircuralSum = maxSubarraysum + invertedSum = 12
	Maximum sum in circural array is max(maxConsecutiveSum, maxCircuralSum)
*/

// Solution 0(n)
func maxSumCircuralSubarray(arr []int) int {
	// find maximum subarray sum in normal array
	maxConsecutiveSum := maxSumSubarray(arr)

	// check if all elements are negative
	if maxConsecutiveSum < 0 {
		return 0
	}

	arrSum := 0

	// invert the array and calculate the maximum sum
	for i := 0; i < len(arr); i++ {
		arrSum += arr[i]
		arr[i] = -arr[i]
	}
	// Finding maximum sum in inverted array is like finding minimum in original
	maxCircuralSum := arrSum + maxSumSubarray(arr)

	return max(maxCircuralSum, maxConsecutiveSum)
}

// Kadane algorithm
func maxSumSubarray(arr []int) int {
	maxSum := arr[0]
	maxConsecutiveSum := arr[0]
	for i := 1; i < len(arr); i++ {
		maxConsecutiveSum = max(arr[i], maxConsecutiveSum+arr[i])
		maxSum = max(maxSum, maxConsecutiveSum)
	}

	return maxSum
}

func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maxSumCircuralSubarray([]int{5, -2, 3, 4}))
}
