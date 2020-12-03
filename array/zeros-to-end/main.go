package main

import "fmt"

// Question: Move zeros to end
// Example:
// arr: [10, 0, 0, 10, 10, 0]; zeros to end: [10, 10, 10, 0, 0, 0]
// arr: [1, 2, 3, 4]; zeros to end: [1, 2, 3, 4]

// Solution: 0(N)
/*
	Using two pointer approach we keep track of two indices.
	One to keep track of position where to insert non zero element, and
	other to traverse the whole array. We keep track of how many non
	zero element we have.

	In = {10, 0, 0, 10, 0}, Out: {10, 10, 0, 0, 0}
	I1: i = 0, j = 0, In: {10, 0, 0, 10, 0}
	I2: i = 1, j = 1, In: {10, 0, 0, 10, 0}
	I3: i = 2, j = 1, In: {10, 0, 0, 10, 0}
	I4: i = 3, j = 1, In: {10, 0, 0, 10, 0}
	I5: i = 4, j = 2, In: {10, 10, 0, 0, 0}
*/
func zerosToEnd(arr []int) {
	for i, j := 0, 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
}

func main() {
	arr := []int{1, 0, 3, 2, 0, 0, 10}
	zerosToEnd(arr)
	fmt.Println(arr)
}
