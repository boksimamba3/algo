package main

import "fmt"

// Question: Find all leader elements in unsorted array.
// An element is called the leader of an array if there
// is no element greater than it on the right side.
// Example:
// arr: [1, 2, 3, 4, 5]; leaders: [5]
// arr: [5, 4, 3, 2, 1]; leaders: [5, 4, 3, 2, 1]
// arr: [7, 10, 5, 6, 2]; leaders: [10, 6, 2]

// Solution: Naive O(Nˆ2)
/* func getLeaders(arr []int) []int {
	n := len(arr)
	leaders := []int{}

	for i := 0; i < n; i++ {
		isLeader := true
		for j := i + 1; j < n; j++ {
			if arr[i] < arr[j] {
				isLeader = false
				break
			}
		}
		if isLeader {
				leaders = append(leaders, arr[i])
		}
	}

	return leaders
} */

// Solution: Optimized O(n)
// Note: Leaders are in reverse order.
/*
	We know that last element is always a leader.
	We start traversing from second last element,
	and we check if current element is greater than
	current leader if so current element is also a leader.
	We know that cause we start traversing from right to left
	and we always keep track of current leader on right.

	In = {7, 10, 4, 5}, Out: {10, 5}
	I1: {7, 10, 4, 5}, currentLeader: 5
	I2: {7, 10, 4, 5}, i = 2, currentLeader: 5
	I3: {7, 10, 4, 5}, i = 1, currentLeader: 10
	I4: {7, 10, 4, 5}, i = 0, currentLeader: 10
*/
func getLeaders(arr []int) []int {
	n := len(arr)
	current := arr[n-1]
	leaders := []int{current}
	for i := n - 2; i >= 0; i-- {
		if arr[i] > current {
			current = arr[i]
			leaders = append(leaders, current)
		}
	}

	return leaders
}

func main() {
	fmt.Println(getLeaders([]int{7, 10, 4, 3, 6, 5, 2}))
}
