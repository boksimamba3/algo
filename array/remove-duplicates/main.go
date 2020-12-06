package main

import "fmt"

// Question: Remove duplicates from sorted array
// Example:
// arr: [10, 10, 20, 20, 20, 30]; res: [10, 20, 30]
// arr: [1, 2, 3, 4]; reversed: [4, 3, 2, 1]

// Solution: O(N) Space: O(1)
/*
	Keeping distinct elements at the begining of the
	array in sorted order. We have two cursors. One
	to iterate over the all elements, and one to keep
	track of how many distinct elements we have. Everytime
	we find new distinct element we increment count of
	distinct elements.

	arr := {10, 10, 20, 20, 30} => {10, 20, 30, _, _}

	In = {10, 10, 20, 20, 30}, Out: {10, 20, 30}
	I1: i = 1, j = 1, In: {10, 10, 20, 20, 30}
	I2: i = 2, j = 2, In: {10, 20, 20, 20, 30}
	I3: i = 3, j = 2, In: {10, 20, 20, 20, 30}
	I4: i = 4, j = 3, In: {10, 20, 30, 20, 30}

	Out: In[0,j]
*/
func removeDuplicates(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	j := 1 // position of next distinct element
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[j-1] {
			arr[j] = arr[i]
			j++
		}
	}

	return arr[0:j]
}

func main() {
	fmt.Println(removeDuplicates([]int{10, 20, 20, 30, 30, 30, 30, 40, 40, 50, 60}))
}
