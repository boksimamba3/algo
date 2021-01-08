package main

import "fmt"

// Question: Find if array has equilibrium.
// Equilibrium point of an array is an element such
// that the sum of elements at lower indexes is equal
// to the sum of elements at higher indexes.

// Solution: O(N)
/*
	First get total sum of array. Total sum of array will
	represent rightSum. We iterate through the array and for each
	element we update the rightSum -= a[i]. If leftSum is equal
	to rightSum we found equilibrium point. Otherwise update the
	leftSum += a[i].

	Example: In: {-7, 1, 5, 2, -4, 3, 0}: Out: true
	Index at position 3 is equilibrium point because:
	a[0] + a[1] + a[2] = a[4] + a[5] + a[6]

	  2        6
		∑ a[i] = ∑ a[i]
	 i=0      i=4

*/
func hasEquilibriumPoint(arr []int) bool {
	rightSum := 0

	for i := 0; i < len(arr); i++ {
		rightSum += arr[i]
	}

	leftSum := 0

	for i := 0; i < len(arr); i++ {
		rightSum -= arr[i]

		if leftSum == rightSum {
			return true
		}

		leftSum += arr[i]
	}

	return false
}

func main() {
	fmt.Println(hasEquilibriumPoint([]int{3, 4, 8, -9, 20, 6}))
}
