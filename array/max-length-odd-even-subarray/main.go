package main

import "fmt"

// Question: Find the length of the longest subarray that has alternating even odd elements.
// Note: Alterating: (two or more things) occurring in turn repeatedly.

// Solution: O(N)
/*
	Traverse the whole array, and while doing so keep track
	of maximum consecutive count of alternating elements.
	In each step check current element parity against previous element parity.
	If they are different increment the count by one, otherwise reset the count.

	In: {10, 12, 14, 7, 8}, Out: 3
	consecutiveMaxLength = 1
	consecutiveLength = 1
	I1: i = 1, 12 % 2 == 10 % 2, consecutiveLength = 1, consecutiveMaxLength = 1
	I2: i = 2, 14 % 2 == 12 % 2, consecutiveLength = 1, consecutiveMaxLength = 1
	I3: i = 3, 7 % 2 != 14 % 2, consecutiveLength = 2, consecutiveMaxLength = 2
	I4: i = 4, 8 % 2 != 7 % 2, consecutiveLength = 3, consecutiveMaxLength = 3
*/

// Solution 0(n)
func maxLenOddEvenSubarray(arr []int) int {
	if len(arr) < 1 {
		return 0
	}

	consecutiveMaxLength := 1
	consecutiveLength := 1

	for i := 1; i < len(arr); i++ {
		if (remainder(arr[i]) == 0 && remainder(arr[i-1]) != 0) || (remainder(arr[i]) != 0 && remainder(arr[i-1]) == 0) {
			consecutiveLength++
			consecutiveMaxLength = max(consecutiveMaxLength, consecutiveLength)
		} else {
			consecutiveLength = 1
		}
	}

	return consecutiveMaxLength
}

func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

func remainder(x int) int {
	return x % 2
}

func main() {
	fmt.Println(maxLenOddEvenSubarray([]int{10, 12, 14, 7, 8}))
}
