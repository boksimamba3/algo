package main

import "fmt"

// Question: Given a binary array, we need to find the minimum
// of number of group flips to make all array elements same.
// In a group flip, we can flip any set of consecutive 1s or 0s.

// Solution: O(N)
/*
	Solution to this problem is insteresting property of binary array.
	We have only two possibilities. We can flip either zeros or ones.
	If we make a rule that we will always fill second group of elements
	we will always have at least equal number of flips or less number of
	flips if started with first group.

	In: {1, 1, 0, 0, 1, 0, 1}, Out: 2
	Number of flips if we flipped ones: 3
	Number of flips if we flipped zeros: 2
	Array begins and ends with same element so number of flips is one less than
	if we flipped the first group.

	In: {0, 0, 0, 1}, Out: 1
	Number of flips if we flipped ones: 1
	Number of flips if we flipped zeros: 1
	Array begins with different element compared to the end element, so number of
	flips is equal to number of flips if started with first group.

	In: {1, 1, 1, 1}, Out: 1
	Array has all element the same so number of flips is zero.
*/
func minConsecutiveFlipsCount(arr []int) int {
	flips := 0
	first := arr[0]
	consecutive := false
	for i := 1; i < len(arr); i++ {
		if arr[i] != first {
			if !consecutive {
				flips++
			}
			consecutive = true
		} else {
			consecutive = false
		}
	}

	return flips
}

func minConsecutiveFlips(arr []int) {

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			if arr[i] != arr[0] {
				fmt.Printf("From %d to", i)
			} else {
				fmt.Printf(" %d\n", i-1)
			}
		}
	}
	if arr[len(arr)-1] != arr[0] {
		fmt.Printf(" %d\n", len(arr)-1)
	}
}

func main() {
	minConsecutiveFlips([]int{1, 0, 0, 0, 1, 0, 1, 0, 0, 0})
	fmt.Println("Min consecutive flips count: ", minConsecutiveFlipsCount([]int{1, 0, 0, 0, 1, 0, 1, 0, 0, 0}))
}
