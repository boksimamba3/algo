package main

import "fmt"

// Question: Given a sorted array find frequencies of all the elements in array.

// Solution: O(N)
/*
 We set count to 1 since first element will appear at least one time.
 Start traversing from second element until the end. On each iteration
 check if element is equal to previous. If it is increment the count.
 If current element is different from previous, we push that element
 together with the count to frequencies array and we reset count to 1.
*/
func frequenciesCounter(arr []int) [][]int {
	frequencies := [][]int{}

	count := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			count++
		}

		if arr[i] != arr[i-1] || i == len(arr)-1 {
			frequencies = append(frequencies, []int{arr[i-1], count})
			count = 1
		}
	}

	return frequencies
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(frequenciesCounter([]int{1, 2, 3, 3, 3}))
}
