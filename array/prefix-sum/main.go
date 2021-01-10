package main

import "fmt"

// Question: Given a fixed array and multiple queries of following
// types on the array, how to efficiently perform these queries.

// Solution: O(N)
/*
	In order for every query to take O(1) we need to precompute the prefix
	sum array. To build prefix sum array it takes O(N). Each prefix sum is
	equal to previous prefix sum plus current element. For example if we have
	array {2, 8, 3, 9}, prefix sum array will be {2, 10, 13, 22}. This way we
	can get each sum of subarray in O(1).
*/
func prefixSum(arr []int) func(start, end int) int {
	a := make([]int, len(arr))
	copy(a, arr)
	for i := 1; i < len(a); i++ {
		a[i] += a[i-1]
	}

	return func(start int, end int) int {
		if start == 0 {
			return a[end]
		}

		return a[end] - a[start-1]
	}
}

func main() {
	getSum := prefixSum([]int{2, 8, 3, 9, 6, 5, 4})
	fmt.Println(getSum(0, 2))
	fmt.Println(getSum(1, 3))
	fmt.Println(getSum(2, 6))
}
