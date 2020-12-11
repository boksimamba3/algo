package main

import "fmt"

// Question: Given n non-negative integers representing an elevation map
// where the width of each bar is 1, compute how much water it can trap after raining.
// Example: In: {3, 0, 1, 2, 5}, Out: 6
/* 								 ---
									|   |
 	 ---	~		~		~ |   |
	|   | ~   ~  ---|   |
	|   | ~  ---|   |   |
	|_3_|_0_|_1_|_2_|_5_|
*/

// Solution: O(N) Space O(N)
/*
	The idea is to precompute right and left maximum for each bar.
	To precompute left maximum we set the first element as maximum and
	we traverse to right while keeping track of maximum element on left. We do same
	for right maximum. We start from last element and we traverse to
	left while keeping track of maximum element on right.
	When we precompute the left and right, to get water for middle bars we need
	to find minimum of left and right maximum and subtract the current bar height.

	In: {3, 0, 1, 2, 5}, Out: 6
	res = 0
	lmax: {3, 3, 3, 3, 5}
	rmax: {5, 5, 5, 5, 5}
	I1: i = 1; res := 0 + min(3, 5) - 0 = 0 + 3 - 0 = 3
	I1: i = 2; res := 3 + min(3, 5) - 0 = 3 + 3 - 1 = 5
	I1: i = 3; res := 5 + min(3, 5) - 0 = 5 + 3 - 2 = 6
*/
func getWater(arr []int) int {
	res := 0
	lMax := make([]int, len(arr))
	rMax := make([]int, len(arr))

	lMax[0] = arr[0]

	// Precompute left maximums
	for i := 1; i < len(arr); i++ {
		lMax[i] = max(arr[i], lMax[i-1])
	}

	rMax[len(arr)-1] = arr[len(arr)-1]

	// Precompute right maximums
	for i := len(arr) - 2; i >= 0; i-- {
		rMax[i] = max(arr[i], rMax[i+1])
	}

	for i := 0; i < len(arr); i++ {
		res += min(lMax[i], rMax[i]) - arr[i]
	}

	return res
}

func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

func min(a int, b int) int {
	if a <= b {
		return a
	}

	return b
}

func main() {
	arr := []int{2, 1, 3, 1, 2, 0, 1}
	fmt.Println(getWater(arr))
}
