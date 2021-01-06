package main

import "fmt"

// Question: Find maximum circular subarray sum of a given array.

// Solution: O(N)
/*
	I. Find majority candidate
	II. Check if candidate is actual majority

	When processing an element or array, if the counter is zero,
	the algorithm stores current element as current majority and
	sets the counter to one. Otherwise, it compares current element
	to the current majority and either increments the counter (if they are equal)
	or decrements the counter (otherwise). At the end of this process, if the array
	has a majority, it will be the element stored by the algorithm as current majority.
	After we found majority candidate we need to check if it is actual majority element.
	This algorithm is also know as Boyer–Moore majority vote algorithm.
*/

func findMajority(arr []int) int {
	majorityIndex := 0
	count := 1

	// Find a candidate
	for i := 1; i < len(arr); i++ {
		if arr[majorityIndex] == arr[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			majorityIndex = i
			count = 1
		}
	}

	// Check if the candidate is majority
	if isMajority(arr, arr[majorityIndex]) {
		return majorityIndex
	}

	return -1
}

func isMajority(arr []int, n int) bool {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			count++
		}
	}

	if count <= len(arr)/2 {
		return false
	}

	return true
}

func main() {
	fmt.Println(findMajority([]int{8, 3, 4, 8, 8}))
}
